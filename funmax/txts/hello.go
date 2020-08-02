package main

import (
	"bufio"
	"fmt"
	"io"
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
	//打开文件
	inputFile, inputError := os.Open("D:/score.txt")
	if inputError != nil {
		return
	}
	defer inputFile.Close()

	//按行读取文件
	var s string
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break
		}
		s = s + inputString
	}
	fmt.Printf("all: \n%s", s)
	fmt.Println("-------------------------------------------------")



}