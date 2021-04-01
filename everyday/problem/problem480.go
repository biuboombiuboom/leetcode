package problem

func medianSlidingWindow(nums []int, k int) []float64 {
	medians := make([]float64, 0)
	//n := len(nums)
	//firstWindow := append([]int{}, nums[0:k]...)
	//sort.Ints(firstWindow)
	//medians = append(medians, getMedian(firstWindow))
	//for i := 1; i <= n-k; i++ {
	//	out := nums[i-1]
	//	input := nums[i]
	//}
	return medians
}

type window struct {
	k    int
	nums []int
}

func newWindow(k int) *window {
	return &window{
		k,
		make([]int, 0),
	}
}

func (window *window)remove(num int){

}

func (window *window)insert(num int){
}

func (window *window) getMedian() float64 {
	k := window.k
	if k%2 == 1 {
		return float64(window.nums[k/2])
	} else {
		sum := window.nums[k/2-1] + window.nums[k/2]
		return float64(sum) / 2
	}
}
