package main

import "fmt"

type Roman struct {
	Val      int
	Priority int
}

var rep = map[rune]*Roman{
	'I': &Roman{1, 1},
	'V': &Roman{5, 2},
	'X': &Roman{10, 3},
	'L': &Roman{50, 4},
	'C': &Roman{100, 5},
	'D': &Roman{500, 6},
	'M': &Roman{1000, 7},
}

func romanToInt(s string) int {
	ret := 0
	var last *Roman
	// s := "DCXXI"
	for i, c := range s {
		cur, ok := rep[c]
		if !ok {
			panic("not a valid roman character")
		}
		if i == 0 {
			ret += cur.Val
		} else {
			if cur.Priority > last.Priority {
				ret -= last.Val
				ret += cur.Val - last.Val
			} else {
				ret += cur.Val
			}
		}
		last = cur
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(romanToInt("MCMXCVI"))
}
