package common

import (
	"errors"
	"fmt"
)

type SliceDispose struct {
}

func (s *SliceDispose)SameValue(a interface{}, b interface{}) (reply interface{}, err error){
	aType := fmt.Sprintf(`%T`, a)
	bType := fmt.Sprintf(`%T`, b)
	if aType != bType {
		return nil, errors.New(errorSliceType)
	}
	switch a.(type) {
	case []int:
		return sliceSameValueHandlerInt(a.([]int), b.([]int)), nil
	case []string:
		return sliceSameValueHandlerString(a.([]string), b.([]string)), nil
	case []uint:
		return sliceSameValueHandlerUint(a.([]uint), b.([]uint)), nil
	case []int64:
		return sliceSameValueHandlerInt64(a.([]int64), b.([]int64)), nil
	default:
		return nil, errors.New(errorSliceTypeCannotDispose)
	}
}

func sliceSameValueHandlerInt(a []int, b[]int) (reply []int){
	for _, v := range a {
		for _, v1 := range b {
			if v == v1 {
				reply = append(reply, v1)
			}
		}
	}
	return
}

func sliceSameValueHandlerUint(a []uint, b[]uint) (reply []uint){
	for _, v := range a {
		for _, v1 := range b {
			if v == v1 {
				reply = append(reply, v1)
			}
		}
	}
	return
}

func sliceSameValueHandlerInt64(a []int64, b[]int64) (reply []int64){
	for _, v := range a {
		for _, v1 := range b {
			if v == v1 {
				reply = append(reply, v1)
			}
		}
	}
	return
}

func sliceSameValueHandlerString(a []string, b[]string) (reply []string){
	for _, v := range a {
		for _, v1 := range b {
			if v == v1 {
				reply = append(reply, v1)
			}
		}
	}
	return
}

