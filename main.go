package main

import "fmt"

func main() {

	eNameMap := make(map[string][]string)
	for i := 0; i < 10; i++ {
		eNameMap["www"] = append(eNameMap["www"], "eeeeeeeeeeee")
	}
	fmt.Println(eNameMap["www"])
}
