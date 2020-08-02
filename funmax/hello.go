package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grade struct{
	Code    int
	Name    string
}
type Classid struct {

	Code    int
	Name    string
}
type Report struct{
	Code    string
	Name    string
	Score   int
}

func main() {
	file,err := os.Open("D:/score.txt")
	if err != nil {
		fmt.Println("read file failed,",err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	str,err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string failed,",err)
	}
	fmt.Printf("str:%s\n",str)
}