package main

import (
	"fmt"
	"github.com/zouhuigang/php"
)

func main() {
	a, b := 2, 3
	max := php.If(a > b, a, b).(int)
	fmt.Println(max)

	str1 := "sdsadasddsads"
	s1 := php.If(str1 != "", "存在", "不存在").(string)
	fmt.Println(s1)

	str2 := ""
	s2 := php.If(str2 != "", "存在", "不存在").(string)
	fmt.Println(s2)
}
