package p202103

import (
	"strconv"
)

func calculate(s string) int {
	typeMap := map[string]string{
		"+": "op",
		"-": "op",
		"*": "op",
		"/": "op",
		"0": "num",
		"1": "num",
		"2": "num",
		"3": "num",
		"4": "num",
		"5": "num",
		"6": "num",
		"7": "num",
		"8": "num",
		"9": "num",
		"(": "left",
		")": "right",
	}
	stacks := make([]string, 0)
	n := len(s)
	for i := 0; i < n; i++ {
		s1 := s[i : i+1]
		t, _ := typeMap[s1]
		switch t {
		case "right":
			stacks = minCalculate(stacks)
			break
			//出栈 3个
			//计算
			//入栈
		case "left":
			break
		case "op":
			if len(stacks) > 3 {
				preOp := stacks[len(stacks)-2]
				if s1 == "+" || s1 == "-" {
					stacks = minCalculate(stacks)
				} else if preOp == "*" || preOp == "/" {
					stacks = minCalculate(stacks)
				}
			}
			stacks = append(stacks, s1)
			break
		case "num":
			l := len(stacks)
			if l > 0 {
				top := stacks[l-1]
				if tt, _ := typeMap[top[0:1]]; tt == "num" {
					stacks[l-1] = top + s1
				} else {
					stacks = append(stacks, s1)
				}
			} else {
				stacks = append(stacks, s1)
			}

			break
		}
	}
	for len(stacks) >= 3 {
		stacks = minCalculate(stacks)
	}
	ans, _ := strconv.Atoi(stacks[0])

	return ans
}

func minCalculate(stacks []string) []string {
	l := len(stacks)
	if l < 3 {
		return stacks
	}
	a, _ := strconv.Atoi(stacks[l-3])
	b, _ := strconv.Atoi(stacks[l-1])
	ans := 0
	switch stacks[l-2] {
	case "+":
		ans = a + b
	case "-":
		ans = a - b
	case "*":
		ans = a * b
	case "/":
		ans = a / b
	}
	return append(stacks[0:l-3], strconv.Itoa(ans))
}
