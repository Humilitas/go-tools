package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var count = 0
var target string
var allFileFlag = "*"

func main() {
	path := "."
	defaultTarget := "node_modules"

	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("请输入想要删除的文件或文件夹（默认为 node_modules，“%s”表示删除所有）：\n", allFileFlag)

	// 逐行扫描
	for input.Scan() {
		target = input.Text()
		fmt.Println(target)
		if target == "" {
			target = defaultTarget
		}
		deal(path)
		break
	}
	if count == 0 {
		fmt.Println("\n\t未找到目标文件：" + target)
	}

	delay := 5
	fmt.Printf("%ds 后自动关闭", delay)
	time.Sleep(time.Second * time.Duration(delay))
}
func deal(path string) {
	// 读取当前文件夹
	dir, err := ioutil.ReadDir(path)
	handleErr(err)
	for _, info := range dir {
		targetPath := path + "/" + info.Name()
		if info.Name() == target || allFileFlag == target {
			fmt.Println("** found target: " + targetPath)
			removeSub(targetPath)
			handleErr(os.RemoveAll(targetPath))
			fmt.Println("** removed target: " + targetPath)
			count++
		} else if info.IsDir() {
			fmt.Println("found directory: " + targetPath)
			deal(path + "/" + info.Name())
		}
	}
}
func removeSub(path string) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		handleErr(os.RemoveAll(path))
	}
	for _, info := range dir {
		handleErr(os.RemoveAll(path + "/" + info.Name()))
		fmt.Println("\tremoved: " + info.Name())
	}
}
func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
