package main

import (
	"blog/utils"
	"fmt"
)

func main() {
	commod := fmt.Sprintf(`cd ~/CompressFiles;./debugStraMkd -v  | grep -o '"BuildTime":"[^"]*"' | awk -F'"' '{print $4}'`)
	outPut, err := utils.SSHCmd("114.55.134.26", commod)
	if err != nil {
		println(err)
	}
	fmt.Println("------114.55.134.26----------", outPut)
	//
}
