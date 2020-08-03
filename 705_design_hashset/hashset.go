package hashset

type MyHashSet struct {
	// cheated with 1d array
	store [1000001]bool
}

func Constructor() *MyHashSet {
	return &MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	this.store[key] = true
}

func (this *MyHashSet) Contains(key int) bool {
	return this.store[key]
}

func (this *MyHashSet) Remove(key int) {
	this.store[key] = false
}
