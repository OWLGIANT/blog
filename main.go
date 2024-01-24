package main

import (
	"fmt"
	ezip "github.com/alexmullins/zip"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {

		//file, header, err := c.Request.FormFile("file")

		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("获取表单失败: %s", err.Error()))
			return
		}

		files := form.File["files"]

		// 创建一个临时目录用于存放上传的文件
		tempDir := "./temp"
		err = os.MkdirAll(tempDir, os.ModePerm)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("创建临时目录失败: %s", err.Error()))
			return
		}

		// 保存上传的文件到临时目录
		for _, file := range files {
			src, oerr := file.Open()
			if oerr != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("打开文件失败: %s", oerr.Error()))
				return
			}
			defer src.Close()

			dstPath := filepath.Join(tempDir, file.Filename)
			dst, cerr := os.Create(dstPath)
			if cerr != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("创建文件失败: %s", cerr.Error()))
				return
			}
			defer dst.Close()

			_, err = io.Copy(dst, src)
			if err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("保存文件失败: %s", err.Error()))
				return
			}
		}

		// 创建一个ZIP文件并设置密码
		zipFilePath := "./uploads/uploaded_files.zip"
		password := "your_password"
		err = EncryptAndZip(tempDir, zipFilePath, password)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("创建并加密ZIP文件失败: %s", err.Error()))
			return
		}

		// 删除临时目录
		err = os.RemoveAll(tempDir)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("删除临时目录失败: %s", err.Error()))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("文件上传并压缩成功，ZIP文件路径：%s", zipFilePath))
	})

	router.Run(":8080")
}

// EncryptAndZip 加密压缩文件夹
func EncryptAndZip(src, dest, password string) error {
	zipFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := ezip.NewWriter(zipFile)
	defer archive.Close()

	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 相对路径
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		header, err := ezip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// 设置密码
		header.SetPassword(password)

		header.Name = strings.Replace(relPath, string(filepath.Separator), "/", -1)

		// 如果是目录，则只创建目录，不写入文件
		if info.IsDir() {
			header.Name += "/"
			_, err = archive.CreateHeader(header)
			return err
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		// 如果是文件，则写入文件内容
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
