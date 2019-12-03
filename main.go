package main

import (
	"fmt"

	"github.com/biuboombiuboom/leetcode/problem66"
)

func main() {

	// 717
	// inputs := []int{1, 1, 1, 0}
	// fmt.Println(strconv.FormatBool(problem717.IsOneBitCharacter(inputs)))

	// 1
	// nums := []int{3, 2, 4}
	// target := 6
	// result := problem1.TwoSum(nums, target)

	// 7
	// input := -123
	// result := problem7.Reverse(input)

	// input := 121
	// result := problem9.IsPalindrome(input)

	// input := "MCMXCIV"
	// result := problem13.RomanToInt(input)

	// input := []string{"aaa", "aa", "aaa"}
	// result := problem14.LongestCommonPrefix(input)

	// input := "()"
	// result := problem20.IsValid(input)

	// l1 := &problem21.ListNode{
	// 	Val: 1,
	// 	Next: &problem21.ListNode{
	// 		Val: 2,
	// 		Next: &problem21.ListNode{
	// 			Val:  4,
	// 			Next: nil,
	// 		},
	// 	},
	// }

	// l2 := &problem21.ListNode{
	// 	Val: 1,
	// 	Next: &problem21.ListNode{
	// 		Val: 3,
	// 		Next: &problem21.ListNode{
	// 			Val:  4,
	// 			Next: nil,
	// 		},
	// 	},
	// }

	// l1 = problem21.MergeTwoLists(l1, l2)
	// for {
	// 	fmt.Printf("%d,", l1.Val)
	// 	l1 = l1.Next
	// 	if l1 == nil {
	// 		break
	// 	}
	// }

	// input := []int{0, 1, 2, 2, 3, 0, 4, 2}
	// result := problem27.RemoveElement(input, 2)
	// for i := 0; i < result; i++ {
	// 	fmt.Print(input[i])
	// }

	// l1 := &problem2.ListNode{
	// 	Val: 2,
	// 	Next: &problem2.ListNode{
	// 		Val:  4,
	// 		Next: nil,
	// 		// Next: &problem2.ListNode{
	// 		// 	Val:  3,
	// 		// 	Next: nil,
	// 		// },
	// 	},
	// }

	// l2 := &problem2.ListNode{
	// 	Val: 5,
	// 	Next: &problem2.ListNode{
	// 		Val: 6,
	// 		Next: &problem2.ListNode{
	// 			Val:  4,
	// 			Next: nil,
	// 		},
	// 	},
	// }
	// result := problem2.AddTwoNumbers(l1, l2)

	// for result != nil {
	// 	fmt.Println(result.Val)
	// 	result = result.Next
	// }
	// fmt.Println(result)

	// fmt.Println(problem28.StrStr("abcaabababaa", "abab"))
	// fmt.Println(problem3.LengthOfLongestSubstring("acdfggseercghacthgfhfg"))
	// fmt.Println(problem35.SearchInsert1([]int{1, 2, 3, 4, 5, 6, 10, 15, 22}, 15))
	input := []int{9}
	fmt.Println(problem66.PlusOne(input))

}
