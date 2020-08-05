package wd

type WordDict struct{}

func Constructor() *WordDict {
	return &WordDict{}
}

func (this *WordDict) AddWord(word string) {
}

func (this *WordDict) Search(word string) bool {
	return false
}
