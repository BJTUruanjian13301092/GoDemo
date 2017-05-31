package main

import(
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"bufio"
	"time"
	"strings"
)

//First way to read a file
func read1(path string)string{
	fi,err := os.Open(path)
	if err != nil{
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte,1024,1024)
	buf := make([]byte,1024)
	for{
		n,err := fi.Read(buf)
		if err != nil && err != io.EOF{panic(err)}
		if 0 ==n {break}
		chunks=append(chunks,buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	return string(chunks)
}

//Second way to read a file
func read2(path string)string{
	fi,err := os.Open(path)
	if err != nil{panic(err)}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte,1024,1024)

	buf := make([]byte,1024)
	for{
		n,err := r.Read(buf)
		if err != nil && err != io.EOF{panic(err)}
		if 0 ==n {break}
		chunks=append(chunks,buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	return string(chunks)
}

//Third way to read a file
func read3(path string)string{
	fi,err := os.Open(path)
	if err != nil{panic(err)}
	defer fi.Close()
	fd,err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}

//Fourth way to read a file( One line each time )
func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		//handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func Print(line string) {
	fmt.Println(line)
}

func main(){

	//This is my file path
	file := "C:\\UsefulSoftware\\Go_IDE\\Go_project\\go_test\\ReadMe.txt"
	f,err := ioutil.ReadFile(file)
	if err != nil{
		fmt.Printf("%s\n",err)
		panic(err)
	}
	fmt.Println(string(f))
	start := time.Now()
	read1(file)
	t1 := time.Now()
	fmt.Printf("Cost time %v\n",t1.Sub(start))
	read2(file)
	t2 := time.Now()
	fmt.Printf("Cost time %v\n",t2.Sub(t1))
	read3(file)
	t3 := time.Now()
	fmt.Printf("Cost time %v\n",t3.Sub(t2))
	ReadLine(file,Print)
	t4 := time.Now()
	fmt.Printf("Cost time %v\n",t4.Sub(t3))
}