package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	fmt.Printf("确定清空当前目录下的所有文件内容吗？（该操作不可逆，请慎重考虑！！） y/N：\n")

	// 逐行扫描
	for input.Scan() {
		line := input.Text()

		if line == "y" || line == "Y" {
			deal(".")
			break
		} else {
			fmt.Println("取消操作，程序即将关闭...")
			break
		}
	}

	closeDelay := 5
	fmt.Printf("close in %d seconds\n", closeDelay)
	time.Sleep(time.Duration(closeDelay) * time.Second)
}
func deal(path string) {
	// 读取当前文件夹
	dir, err := ioutil.ReadDir(path)
	handleErr(err)
	// 改名
	for _, info := range dir {
		targetFile := path + "/" + info.Name()
		fmt.Println("found: ", targetFile)
		if !info.IsDir() {
			err := os.Truncate(targetFile, 0)
			handleErr(err)
			fmt.Println("truncate: " + targetFile)
		} else {
			deal(path + "/" + info.Name())
		}
	}
}
func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
