package idg

import "math/rand"

type RandomizedSet struct {
	// key->val value: pos in res
	hashset map[int]int
	res     []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{hashset: make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.hashset[val]; exists {
		return false
	}
	this.res = append(this.res, val)
	this.hashset[val] = len(this.res) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, exists := this.hashset[val]; !exists {
		return false
	}
	n := len(this.res)
	pos, last := this.hashset[val], n-1
	// swap with the last element in res
	this.res[pos], this.res[last] = this.res[last], this.res[pos]
	this.hashset[this.res[pos]] = pos
	// delete from hashmap
	delete(this.hashset, val)
	// shrink this.res
	this.res = this.res[:n-1]
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.res[rand.Intn(len(this.res))]
}
