package main

import (
	"blog/utils"
	"fmt"
)

func main() {
	outPut, err := utils.SSHCmd("114.55.134.26", "cd ~/CompressFiles;./test.sh")
	if err != nil {
		println(err)
	}
	fmt.Println("------114.55.134.26----------", outPut)
}
