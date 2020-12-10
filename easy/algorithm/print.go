package algorithm


import "fmt"

func Print() {

	//moveZeroes(nums1)
	//moveZeroes(nums2)
	//
	//moveZeroes(nums3)

	head:=&ListNode{
		Val: 1,
		Next:&ListNode{
			Val: 2,
			Next:&ListNode{
				Val: 2,
				Next:&ListNode{
					Val: 3,
					Next:&ListNode{
						Val: 4,
					},
				},
			},

		},
	}
	fmt.Printf("%v\n", removeElements(head,2))

}
