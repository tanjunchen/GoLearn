/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//cmd.Execute()
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(strconv.Itoa(rand.Intn(10000)))
	//fmt.Print(rand.New(ran))
	// 实现思路
	// 数值转字符串
	// 逆序循环颠倒字符串
	// 补足 32 位
	// 字符串转数值
	num := 00000010100101000001111010011100
	str := strconv.FormatUint(uint64(num), 2)
	rev := ""
	for i := len(str) - 1; i >= 0; i-- {
		rev = rev + str[i:i+1]
	}
	if len(rev) < 32 {
		rev = rev + strings.Repeat("0", 32-len(rev))
	}
	n, _ := strconv.ParseUint(rev, 2, 64)
	//
	fmt.Print(uint32(n))
}
