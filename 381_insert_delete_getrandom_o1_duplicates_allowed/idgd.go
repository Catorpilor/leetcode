package idgd

import (
	"fmt"
	"math/rand"
)

type RandomizedCollection struct {
	hashset map[int][]int
	res     []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{hashset: make(map[int][]int)}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	this.res = append(this.res, val)
	this.hashset[val] = append(this.hashset[val], len(this.res)-1)
	return true
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	if _, exists := this.hashset[val]; !exists {
		return false
	}
	sn, n := len(this.hashset[val]), len(this.res)
	v1 := this.hashset[val]
	pos := v1[sn-1]
	key := this.res[n-1]
	fmt.Printf("key:%d, pos:%d, v1:%v, this.hashset[key]:%v\n", key, pos, v1, this.hashset[key])
	// swap
	this.res[pos], this.res[n-1] = this.res[n-1], this.res[pos]
	v := this.hashset[key]
	v[len(v)-1] = pos
	this.hashset[key] = v
	this.hashset[val] = v1[:sn-1]
	this.res = this.res[:n-1]
	return true

}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.res[rand.Intn(len(this.res))]
}
