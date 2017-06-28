方面在golang里能用上跟php一样的函数

例如:in_array等

使用：

	package main
	
	import (
		"fmt"
		"github.com/zouhuigang/php"
	)
	
	func main() {
		fmt.Println("\ntest1 .....")
		names := []string{"Mary", "Anna", "Beth", "Johnny", "Beth"}
		fmt.Println(php.In_array2("Anna", names))
		fmt.Println(php.In_array2("Jon", names))
	
		ints := []int{1, 4, 3, 2, 6}
		fmt.Println(php.In_array2(3, ints))
		fmt.Println(php.In_array2(95, ints))
	
		//测试2
		fmt.Println("\ntest2 .....")
		names2 := []string{"Mary", "Anna", "Beth", "Johnny", "Beth"}
		fmt.Println(php.In_array("Anna", names2))
		fmt.Println(php.In_array("Jon", names2))
	
		ints2 := [...]int{1, 4, 3, 2, 6}
		fmt.Println(php.In_array(3, ints2))
		fmt.Println(php.In_array(95, ints2))
	
		mix := [...]interface{}{1331, 44.5, "test"}
		fmt.Println(php.In_array(1331, mix))
		fmt.Println(php.In_array(44.5, mix))
		fmt.Println(php.In_array("bad", mix))
	}

输出：

	test1 ...
	true 1
	false -1
	true 2
	false -1
	
	test2 ...
	true 1
	false 5
	true 2
	false 5
	true 0
	true 1
	false 3