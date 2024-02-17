package utils

import (
	"fmt"
	"github.com/pkg/sftp"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"time"
)

func ScpFile(localFilePath string) bool {
	logrus.Infof("开始上传文件 %v", localFilePath)
	remoteFilePath := fmt.Sprintf("/root/upload/mkd/%v", localFilePath)

	// 建立SSH连接
	sshConfig := &ssh.ClientConfig{
		User: "root", // 你的SSH用户名
		Auth: []ssh.AuthMethod{
			ssh.Password("1325384144@qq"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	sshClient, err := ssh.Dial("tcp", "114.55.134.26:22", sshConfig)
	if err != nil {
		logrus.Warnf("Error establishing SSH connection: %v", err)
		return false
	}
	defer sshClient.Close()

	// 打开SFTP会话
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		logrus.Warnf("Error creating SFTP client: %v", err)
		return false
	}
	defer sftpClient.Close()

	// 打开本地文件
	localFile, err := os.Open(localFilePath)
	if err != nil {
		logrus.Warnf("Error opening local file: %v", err)
		return false
	}
	defer localFile.Close()

	// 创建远程文件
	remoteFile, err := sftpClient.Create(remoteFilePath)
	if err != nil {
		logrus.Warnf("Error creating remote file: %v", err)
		return false
	}
	defer remoteFile.Close()

	fmt.Println("==============time.Sleep========")
	time.Sleep(20 * time.Second)
	// 将本地文件内容拷贝到远程文件
	bytesCopied, err := io.Copy(remoteFile, localFile)
	if err != nil {
		logrus.Warnf("Error copying file: %v", err)
		return false
	}

	return fileUploadCheck(localFilePath, remoteFilePath, bytesCopied, localFile, sftpClient)
}

func fileUploadCheck(localFilePath, remoteFilePath string, bytesCopied int64, localFile *os.File, sftpClient *sftp.Client) bool {

	localFileInfo, err := localFile.Stat()
	if err != nil {
		logrus.Warnf("localFile.Stat() err: %v", err)
		return false
	}

	remoteFileInfo, err := sftpClient.Stat(remoteFilePath)
	if err != nil {
		logrus.Warnf("sftpClient.Stat(%s) err: %v", remoteFilePath, err)
		return false
	}

	//if localFileInfo.ModTime().After(remoteFileInfo.ModTime()) { //本地文件修改时间需 小于远程文件修改时间
	//	logrus.Warnf("%v File copy failed: localFileModeTime After remoteFileModeTime", localFilePath)
	//	return false
	//}

	if !(bytesCopied == localFileInfo.Size() && bytesCopied == remoteFileInfo.Size()) { //文件大小需要一致
		logrus.Warnf("%v File copy failed: Sizes do not match.", localFilePath)
		return false
	}

	logrus.Infof("%v File copied successfully.", localFilePath)
	return true
}
