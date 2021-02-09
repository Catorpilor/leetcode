packge iter

type Iterator struct{}

func (it *Iterator) next() int {
	return 0
}

func (it *Iterator) hasNext() bool {
	return false
}

type PeekIterator struct{
	iter *Iterator
	nextVal int
}

func Constructor(iter *Iterator) *PeekIterator{
	next := -1
	if iter.hasNext() {
		next = iter.next()
	}
	return &PeekIterator{
		iter: iter,
		nextVal: next,
	}
}

func (this *PeekIterator) peek() int {
	return this.nextVal
}

func (this *PeekIterator) hasNext() bool {
	return this.nextVal != -1
}

func (this *PeekIterator) next() int {
	ans := this.nextVal
	if this.iter.hasNext() {
		this.nextVal = this.iter.next()
	}else{
		this.nextVal = -1
	}
	return ans
}
