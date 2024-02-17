package utils

//
//import (
//	"bytes"
//	"github.com/sirupsen/logrus"
//	"golang.org/x/crypto/ssh"
//	"strings"
//)
//
//var sshConfig *ssh.ClientConfig // ssh连接配置
//func init() {
//	var err error
//	sshConfig, err = newSshConfig() //ssh连接初始化
//	if err != nil {
//		logrus.Error(err)
//		return
//	}
//}
//
//// 堡垒机  用的是堡垒机上的 ssh公钥 配置
//func newSshConfig() (conf *ssh.ClientConfig, err error) {
//	//privateKeyBytes, err := ioutil.ReadFile(`/root/.ssh/id_rsa`)
//	//if err != nil {
//	//	logrus.Error("Failed to open id_rsa:", err)
//	//	return
//	//}
//	//privateKey, err := ssh.ParsePrivateKey(privateKeyBytes)
//	//if err != nil {
//	//	logrus.Error("Failed to Parse private key: ", err)
//	//	return
//	//}
//	conf = &ssh.ClientConfig{
//		User: "root", // 你的SSH用户名
//		Auth: []ssh.AuthMethod{
//			ssh.Password("1325384144@qq"),
//		},
//		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
//	}
//	return
//}
//
//func SSHCmd(ip, cmd string) (outPut string, err error) {
//	client, err := ssh.Dial("tcp", ip+":22", sshConfig)
//	if err != nil {
//		return
//	}
//	defer client.Close()
//	session, err := client.NewSession()
//	if err != nil {
//		return
//	}
//	defer session.Close()
//	var b bytes.Buffer
//	session.Stdout = &b
//	if err = session.Run(cmd); err != nil {
//		return
//	}
//	return strings.Trim(b.String(), "\n"), err
//}
