package main

import (
	"fmt"

	"github.com/biuboombiuboom/leetcode/hard/algorithm"
)

func main() {
	n1 := &algorithm.ListNode{
		Val: -6,
		Next: &algorithm.ListNode{
			Val: -3,
			Next: &algorithm.ListNode{
				Val: -1,
				Next: &algorithm.ListNode{
					Val: 1,
					Next: &algorithm.ListNode{
						Val: 2,
						Next: &algorithm.ListNode{
							Val: 2,
							Next: &algorithm.ListNode{
								Val:  2,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	n2 := &algorithm.ListNode{
		Val: -10,
		Next: &algorithm.ListNode{
			Val: -8,
			Next: &algorithm.ListNode{
				Val: -6,
				Next: &algorithm.ListNode{
					Val: -2,
					Next: &algorithm.ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
	}
	n3 := &algorithm.ListNode{
		Val: -2,
	}
	n4 := &algorithm.ListNode{
		Val: -8,
		Next: &algorithm.ListNode{
			Val: -4,
			Next: &algorithm.ListNode{
				Val: -3,
				Next: &algorithm.ListNode{
					Val: -3,
					Next: &algorithm.ListNode{
						Val: -2,
						Next: &algorithm.ListNode{
							Val: -1,
							Next: &algorithm.ListNode{
								Val: 1,
								Next: &algorithm.ListNode{
									Val: 2,
									Next: &algorithm.ListNode{
										Val:  3,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	n5 := &algorithm.ListNode{
		Val: -8,
		Next: &algorithm.ListNode{
			Val: -6,
			Next: &algorithm.ListNode{
				Val: -5,
				Next: &algorithm.ListNode{
					Val: -4,
					Next: &algorithm.ListNode{
						Val: -2,
						Next: &algorithm.ListNode{
							Val: -2,
							Next: &algorithm.ListNode{
								Val: 2,
								Next: &algorithm.ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
		},
	}

	algorithm.MergeKLists([]*algorithm.ListNode{n1, n2, n3, n4, n5})
	fmt.Print("")
}
