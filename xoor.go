package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// 随机生成key,后面用来解密的
func key(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 进行XOR加密
func XorEncrypt(origData, key []byte) []byte {
	crypted := make([]byte, len(origData))
	for i := range origData {
		crypted[i] = origData[i] ^ key[i%len(key)]
	}
	return crypted
}

// 主函数入口,对字符进行了处理
func main() {
	argsWithProg := os.Args
	if len(argsWithProg) < 2 {
		fmt.Println("usage : ", argsWithProg[0], " payload.c")
		return
	}
	confFile := os.Args[1]
	str2 := strings.Replace(confFile, "\\x", "", -1)
	data, _ := hex.DecodeString(str2)
	key1 := key(16)
	fmt.Println("Key:", key1)
	var key []byte = []byte(key1)
	xor := XorEncrypt(data, key)
	encoded := base64.StdEncoding.EncodeToString(xor)
	fmt.Println("Code:", encoded)
}
