package php

/*
三元符运算
a, b := 2, 3
max := If(a > b, a, b).(int)
println(max)
*/
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
