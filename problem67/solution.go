package problem67

//AddBinary a
func AddBinary(a string, b string) string {
	l2 := len(b)
	j := l2 - 1
	i := 0
	for j > -1 {
		b1 := b[j : j+1]
		a = add(a, b1, len(a)-1-i)
		j--
		i++
	}
	return a
}

func add(a string, b string, partition int) string {
	if partition < 0 {
		return b + a
	}
	a1 := a[partition : partition+1]

	if a1 == "1" && b == "1" {
		a = a[0:partition] + "0" + a[partition+1:]
		a = add(a, "1", partition-1)
	} else if a1 == "0" {
		a = a[0:partition] + b + a[partition+1:]
	}
	return a
}
