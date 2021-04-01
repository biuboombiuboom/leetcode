package p202103

type NestedIterator struct {
	nestedList [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{[][]*NestedInteger{nestedList}}
}

func (this *NestedIterator) Next() int {
	topNested := this.nestedList[0]
	top := topNested[0]
	this.nestedList[0]=	topNested[1:]
	return top.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	for len(this.nestedList)>0{
		topNested := this.nestedList[0]
		if len(topNested)==0{
			this.nestedList=this.nestedList[1:]
			continue
		}
		top := topNested[0]
		if top.IsInteger(){
			return true
		}
		this.nestedList[0]=	topNested[1:]
		this.nestedList= append([][]*NestedInteger{top.GetList()},this.nestedList...)

	}
	return len(this.nestedList)>0
}

type NestedInteger struct {
	Integer  int
	Integers []*NestedInteger
}

func (this NestedInteger) IsInteger() bool {
	return this.Integers == nil
}

func (this NestedInteger) GetInteger() int {
	return this.Integer
}

func (this *NestedInteger) SetInteger(value int) {
	this.Integer = value
}

func (this *NestedInteger) Add(elem NestedInteger) {
	this.Integers = append(this.Integers, &elem)
}

func (this NestedInteger) GetList() []*NestedInteger {
	return this.Integers
}
