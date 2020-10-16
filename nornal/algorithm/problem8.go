package algorithm

import (
	"errors"
	"strconv"
)

func MyAtoi(s string) int {
	num := 0
	FirstNotEmpty := ""
	max := (2 << 30) - 1
	min := -1 * (2 << 30)
	for i := 0; i < len(s); i++ {
		s1 := s[i : i+1]
		if s1 == " " {
			if FirstNotEmpty == "" {
				continue
			} else {
				break
			}
		}
		if s1 != " " && FirstNotEmpty == "" {
			FirstNotEmpty = s1
			if FirstNotEmpty == "+" || FirstNotEmpty == "-" {
				continue
			}
		}
		n, err := getNum(s1)
		if err != nil {
			break
		}
		num = num*10 + n
		if num > max  {
			if FirstNotEmpty == "-" {
				num = min
			}else{
				num=max
			}
			return num
		}
	}
	if FirstNotEmpty=="-"{
		num=-1*num
	}

	return num
}

func getNum(s string) (int, error) {
	if len(s) != 1 {
		return 0, errors.New("")
	}
	for i := 0; i < 10; i++ {
		if strconv.Itoa(i) == s {
			return i, nil
		}
	}
	return 0, errors.New("")
}
