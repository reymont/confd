package main

import (
	"fmt"
	"os/exec"
)
//package tools

func main(){
	//执行【ls /】并输出返回文本
	//f, err := exec.Command("ls", "/").Output()
	f, err := exec.Command("echo `dmidecode -s baseboard-serial-number`-`dmidecode -s system-serial-number").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(f))
}
