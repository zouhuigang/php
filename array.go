//https://codereview.stackexchange.com/questions/60074/in-array-in-go
package php

import (
	"reflect"
)

/*
	<?php
$people = array("Bill", "Steve", "Mark", "David");

if (in_array("Mark", $people))
  {
  echo "匹配已找到";
  }
else
  {
  echo "匹配未找到";
  }
?>

*/

func In_array2(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func In_array(v interface{}, in interface{}) (ok bool, i int) {
	val := reflect.Indirect(reflect.ValueOf(in))
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for ; i < val.Len(); i++ {
			if ok = v == val.Index(i).Interface(); ok {
				return
			}
		}
	}
	return
}
