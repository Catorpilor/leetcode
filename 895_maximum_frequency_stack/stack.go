package stack

// FreqStack is a stack with frequences.
type FreqStack struct {
	freq    map[int]int
	store   map[int][]int
	maxFreq int
}

// Construct create a FreqStack
func Construct() *FreqStack {
	return &FreqStack{
		freq:  make(map[int]int),
		store: make(map[int][]int),
	}
}

// Push push x into the stack
func (fs *FreqStack) Push(x int) {
	fs.freq[x]++
	cf := fs.freq[x]
	if cf > fs.maxFreq {
		fs.maxFreq = cf
	}
	fs.store[cf] = append(fs.store[cf], x)
}

// Pop pop the most frequces one from the stack.
func (fs *FreqStack) Pop() int {
	var ans int
	// find the proper frequence
	for len(fs.store[fs.maxFreq]) == 0 {
		fs.maxFreq--
	}
	cands := fs.store[fs.maxFreq]
	n := len(cands)
	if n > 1 {
		ans = cands[n-1]
		cands = cands[:n-1]
		fs.store[fs.maxFreq] = cands
	} else {
		ans = cands[0]
		delete(fs.store, fs.maxFreq)
		fs.maxFreq--
	}
	// update ans's frequence
	fs.freq[ans]--
	return ans
}
