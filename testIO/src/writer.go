package main

import (
	"bufio"  //缓存IO
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) (bool) {
	var exist = true;
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false;
	}
	return exist;
}

func main() {

	var filename = "C:\\UsefulSoftware\\Go_IDE\\Go_project\\go_test\\WriteMe.txt";
	var f *os.File
	var err error;
	if checkFileIsExist(filename) {  //如果文件存在
		f, err = os.OpenFile(filename, os.O_APPEND, 0666)  //打开文件
		fmt.Println("文件存在");
	}else {
		f, err = os.Create(filename)  //创建文件
		fmt.Println("文件不存在");
	}

	w := bufio.NewWriter(f)  //创建新的 Writer 对象
	n, err := w.WriteString("Nice to meet u ~" + "\n" + "HAHAHA")
	fmt.Printf("写入 %d 个字节n", n)
	w.Flush()
	f.Close()
	check(err)
}