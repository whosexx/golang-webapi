package utils

import (
	"reflect"
	"strings"
)

func Contains(array interface{}, val interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					index = i
					return
				}
			}
		}
	}
	return
}

// ContainsString Returns the index position of the string val in array
func ContainsString(array []string, val string, _case bool) (index int) {
	index = -1
	if _case {
		for i := 0; i < len(array); i++ {
			if array[i] == val {
				index = i
				return
			}
		}
	} else {
		for i := 0; i < len(array); i++ {
			if strings.EqualFold(array[i], val) {
				index = i
				return
			}
		}
	}
	return
}
