package main

import (
	"strings"
)

func reformat(s string) string {
	a := []string{s[0:1]}
	numberEmptyIndex := make([]int, 0)
	letterEmptyIndex := make([]int, 0)
	for i := 1; i < len(s); i++ {
		s1 := s[i : i+1]
		if isLeeter(s1) {
			a, letterEmptyIndex, numberEmptyIndex = convA(a, s1, letterEmptyIndex, numberEmptyIndex, isNumber)
		} else {
			a, numberEmptyIndex, letterEmptyIndex = convA(a, s1, numberEmptyIndex, letterEmptyIndex, isLeeter)

		}
	}
	if len(numberEmptyIndex) > 0 || len(letterEmptyIndex) > 0 {
		return ""
	}

	return strings.Join(a, "")
}

func isLeeter(s string) bool {
	lowCases := "abcdefghijklmnopqrstuvwxyz"
	return strings.Index(lowCases, s) > -1
}

func isNumber(s string) bool {
	return !isLeeter(s)
}

func convA(a []string, s1 string, indexa []int, indexb []int, checkStr func(string) bool) ([]string, []int, []int) {
	if len(indexa) > 0 {
		a[indexa[0]] = s1
		indexa = indexa[1:]
	} else {
		if checkStr(a[len(a)-1]) {
			a = append(a, s1)
		} else if checkStr(a[0]) {
			a = append([]string{s1}, a...)
		} else {
			indexb = append(indexb, len(a))
			a = append(a, "")
			a = append(a, s1)
		}
	}
	return a, indexa, indexb
}
