package problem

func maxTurbulenceSize(arr []int) int {
	if len(arr)==1{
		return  1
	}
	res1:=make([]int,len(arr))
	res2:=make([]int,len(arr))
	for i,_:= range res2{
		res2[i]=1

	}
	for i,_:= range res1{
		res1[i]=1

	}
	big:=0
	small:=0
	for i:=0;i<len(arr)-1;i++{
		if arr[i]<arr[i+1]{
			small+=1
		}else if arr[i]>arr[i+1]{
			big+=1
		}
	}

	if big==0||small==0{
		if big==small{
			return 1
		}
		return 2
	}
	max:=0
	for i:=1;i<len(res1);i++{
		if arr[i]<arr[i-1]&&i%2!=0{
			res1[i]=res1[i-1]+1
		}
		if arr[i]<arr[i-1]&&i%2==0{
			res2[i]=res2[i-1]+1
		}
		if arr[i]>arr[i-1]&&i%2==0{
			res1[i]=res1[i-1]+1
		}
		if arr[i]>arr[i-1]&&i%2!=0{
			res2[i]=res2[i-1]+1
		}
	}

	for i,_:=range res1{
		if res1[i]>max{
			max=res1[i]
		}
	}
	for i,_:=range res2{
		if res2[i]>max{
			max=res2[i]
		}
	}
	return max
}
