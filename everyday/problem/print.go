package problem

import "fmt"

func Print() {
	//head := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		2,
	//			&ListNode{
	//				2,
	//				&ListNode{
	//					1,
	//					nil,
	//				},
	//
	//		},
	//	},
	//}
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			2,
			&ListNode{
				5,
				&ListNode{
					2,
					&ListNode{
						1,
						nil,
					},
				},
			},
		},
	}


		fmt.Printf("%v\n", isPalindrome1(head))
	}
