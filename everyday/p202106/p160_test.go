package p202106

import (
	"reflect"
	"testing"
)

//[4,1,8,4,5], listB = [5,0,1,8,4,5],
func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}

	intersectionNode := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	headA := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: intersectionNode,
			},
		},
	}
	headB := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: intersectionNode,
		},
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"c1",
			args{
				headA: headA,
				headB: headB,
			},
			intersectionNode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
