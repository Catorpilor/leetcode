package mtd

import (
	"sort"
	"strconv"
	"strings"

	"github.com/catorpilor/leetcode/utils"
)

func minDiff(tps []string) int {
	return bruteForce(tps)
}

// iterator time complexity is O(nlgn), space complexity is O(N)
func iterator(tps []string) int {
	hset := map[string]int{
		"00": 0,
		"01": 60,
		"02": 120,
		"03": 180,
		"04": 240,
		"05": 300,
		"06": 360,
		"07": 420,
		"08": 480,
		"09": 540,
		"10": 600,
		"11": 660,
		"12": 720,
		"13": 780,
		"14": 840,
		"15": 900,
		"16": 960,
		"17": 1020,
		"18": 1080,
		"19": 1140,
		"20": 1200,
		"21": 1260,
		"22": 1320,
		"23": 1380,
	}
	sort.Slice(tps, func(i, j int) bool { return tps[i] < tps[j] })
	n := len(tps)

	ret := diffs(hset, tps[0], tps[n-1])
	for i := 0; i < n-1; i++ {
		ret = utils.Min(ret, diffs(hset, tps[i], tps[i+1]))
	}
	return ret
}

func diffs(hset map[string]int, left, right string) int {
	f := func(c rune) bool {
		return c == ':'
	}
	lefts := strings.FieldsFunc(left, f)
	rights := strings.FieldsFunc(right, f)
	rm, _ := strconv.Atoi(rights[1])
	rt := hset[rights[0]] + rm
	lm, _ := strconv.Atoi(lefts[1])
	lt := hset[lefts[0]] + lm
	if rights[0] > "12" && lefts[0] < "06" {
		lt += 24 * 60
	}
	return utils.Abs(rt - lt)
}
