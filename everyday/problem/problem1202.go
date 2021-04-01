package problem

//给你一个字符串 s，以及该字符串中的一些「索引对」数组 pairs，其中 pairs[i] = [a, b] 表示字符串中的两个索引（编号从 0 开始）。
//
//你可以 任意多次交换 在 pairs 中任意一对索引处的字符。
//
//返回在经过若干次交换后，s 可以变成的按字典序最小的字符串。
//
// 
//
//示例 1:
//
//输入：s = "dcab", pairs = [[0,3],[1,2]]
//输出："bacd"
//解释：
//交换 s[0] 和 s[3], s = "bcad"
//交换 s[1] 和 s[2], s = "bacd"
//示例 2：
//
//输入：s = "dcab", pairs = [[0,3],[1,2],[0,2]]
//输出："abcd"
//解释：
//交换 s[0] 和 s[3], s = "bcad"
//交换 s[0] 和 s[2], s = "acbd"
//交换 s[1] 和 s[2], s = "abcd"
//示例 3：
//
//输入：s = "cba", pairs = [[0,1],[1,2]]
//输出："abc"
//解释：
//交换 s[0] 和 s[1], s = "bca"
//交换 s[1] 和 s[2], s = "bac"
//交换 s[0] 和 s[1], s = "abc"


func smallestStringWithSwaps(s string, pairs [][]int) string {
	//sort.Slice(pairs, func(i, j int) bool {
	//	return pairs[i][0]<pairs[j][0] ||  (pairs[i][0]==pairs[j][0]&& pairs[i][1]<pairs[j][1])
	//})
	n:=len(pairs)
	uf:=newUnionfind1(n)
	bitPairs :=make([]int,n)

	//查找可以连接在一起的索引对
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {

			xi,yi:=pairs[i][0],pairs[i][1]
			xj,yj:=pairs[j][0],pairs[j][1]

			if xi==xj||xi==yj||yi==xj||yi==yj{
				bitPairs[i]=1
				bitPairs[j]=1
				uf.union(i,j)
			}
		}
	}

	partsMap :=make(map[int][]string)
	distinctIndexMap:=make(map[int]map[int]int)

	//按根分组
	for i := 0; i <n; i++ {
		root:=uf.find(i)
		x,y:=pairs[i][0],pairs[i][1]
		if _,ok:= distinctIndexMap[root];!ok{
			distinctIndexMap[root]=make(map[int]int)
			partsMap[root]=make([]string,0)

		}else{
			if _,ok:=distinctIndexMap[root][x];ok{
				distinctIndexMap[root][x]+=1
			}else{
				distinctIndexMap[root][x]+=0
			}
			if x!=y {
				if _, ok := distinctIndexMap[root][y];  ok{
					distinctIndexMap[root][y]+=1
				} else {
					distinctIndexMap[root][y] += 0
				}
			}

		}
		if distinctIndexMap[root][x]==0{
			partsMap[root]=append(partsMap[root],s[x:x+1] )
		}
		if x!=y&& distinctIndexMap[root][y]==0{
			partsMap[root]=append(partsMap[root],s[y:y+1] )
		}
	}


	for i := 0; i <n ; i++ {
		x,y:=pairs[i][0],pairs[i][1]
		if x>y{
			x,y=y,x
		}
		if bitPairs[i]==0&&s[x]>s[y]{
			s=s[0:x]+s[y:y+1]+s[x+1:y]+s[x:x+1]+s[y+1:]
		}
	}


	return s
}



type unionfind1 struct{
	parent []int
	rank []int
}

func newUnionfind1(n int)*unionfind1  {
	parent:=make([]int,n)
	rank:=make([]int,n)

	for i := 0; i < n; i++ {
		parent[i]=i
		rank[i]=1
	}
	uf:=&unionfind1{
		parent: parent,
		rank:rank,
	}
	return uf

}

func (uf *unionfind1)find(x int)int{
	if uf.parent[x]!=x{
		uf.parent[x]=uf.find(uf.parent[x])
	}
	return  uf.parent[x]
}

func (uf *unionfind1)union(x,y int){
	fx,fy:=uf.find(x),uf.find(y)
	if fx==fy{
		return
	}
	if uf.rank[fx]<uf.rank[fy]{
		fx,fy=fy,fx
	}
	uf.rank[fy]+=uf.rank[fx]
	uf.parent[fy]= uf.parent[fx]

}

