package p202103

func clumsy(N int) int {
	ans := 0
	step := 0
	for N>0 {
		n:=N
		N--
		if N>0{
			n*=N
			N--
			if N>0{
				n/=N
			}
			N--
			if N>0{
				if step==0{
					n+=N
				}else{
					n-=N
				}
			}
		}
		if step==0{
			ans+=n
		}else{
			ans-=n
		}
		step++
		N--
	}

	return ans
}
