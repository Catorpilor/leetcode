package ltp

func numOfPossibilities(tiles string) int {
    res := make(map[string]bool)
    helper(res, []byte(tiles), 0)
    return len(res)
}

func helper(res map[string]bool, st []byte, idx int) {
    if idx == len(st) {
        return
    }
    c := st[idx]
    curKeys := make([]string, 0, len(res))
    for k := range res {
        curKeys = append(curKeys, k)
    }
    for _, k := range curKeys {
        // fmt.Printf("key: %s, c:%s\n", k, string(c))
        n := len(k)
        sk := make([]byte, n+1)
        copy(sk, k)
        sk[n] = c
        for i := 0; i < n; i++ {
            sk[i], sk[n] = sk[n], sk[i]
            res[string(sk)] = true
            sk[i], sk[n] = sk[n], sk[i]
        }
        res[string(sk)] = true
    }
    res[string(c)] = true
    helper(res, st, idx+1)
}

func numOfPossibilities2(tiles string) int {
    count := make([]int, 26)
    for i := range tiles {
        count[tiles[i]-'A']++
    }
    return helper2(count)
}

/*

for example tiles = "ABC", count is [1,1,1,0,0,0...]
1. pick tile "A" res = 1, count = [0,1,1,0,0,0,0....]
                 A, B ,C
                 0, 1, 1
                  / \
      pick tile B/   \ pick tile C
                /     \
               /       \
            A,B,C     A,B,C
            0,0,1     0,1,0
              |          \
              |           \
            A,B,C        A,B,C
            0,0,0        0,0,0
So the valid tiles are ["A","AB","ABC", "AC","ACB"]
2. same as above
The valid tiles beings with B are ["B","BA","BAC","BC","BCA"]
The valid tiles beings with C are ["C","CA","CAB","CB","CBA"]
*/

func helper2(count []int) int {
    var res int
    for i := range count {
        if count[i] == 0 {
            continue
        }
        count[i]--
        // pick this tile
        res++
        // add the # of picks of it's children
        res += helper2(count)
        // backtrack, unpick this tile.
        count[i]++
    }
    return res
}
