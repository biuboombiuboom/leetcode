package problem

func partitionLabels(S string) []int {
	S1 := make([]int, 0)
	smap := make(map[uint8][]int)

	for i := 0; i < len(S); i++ {
		smap[S[i]] = append(smap[S[i]], i)
	}
	startS := S[0]
	endS := S[0]
	start := smap[startS][0]
	end := smap[endS][len(smap[endS])-1]

	for i := 0; i < len(S); i++ {
		curr := S[i]

		//start1 := smap[curr][0]
		end1 := smap[curr][len(smap[curr])-1]

		if end1 < end {
			continue
		}

		if end1 > end {
			endS = curr
			end = end1
			continue
		}

		if end1 == i {
			S1 = append(S1, end-start+1)
			if i < len(S)-1 {
				startS = S[i+1]
				endS = startS
				start = i + 1
				end = smap[endS][len(smap[endS])-1]
			}
		}

	}
	return S1
}
