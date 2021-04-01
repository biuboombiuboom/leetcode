package problem

func equalSubstring(s string, t string, maxCost int) int {
	abs:=func(x int)int{
		if x<0{
			return x* -1
		}
		return x
	}
	diff:=make([]int,0)
	for i := 0; i < len(s); i++ {
		if i<len(t){
			si:=int(s[i])
			ti:=int(t[i])
			diff=append(diff, abs(si-ti))
		}
	}

	max:=0
	for i := 0; i <len(diff) ; i++ {
		if max>len(diff)-i{
			break
		}
		cost:=0
		j:=i
		mc:=maxCost
		for j <len(diff)&&mc>=diff[j]  {
			mc-=diff[j]
			cost++
			j++
		}
		if max<cost{
			max=cost
		}
	}


	return max
}