package ransom

func canConstruct(ransom, magazine string) bool {
    return useHashmap(ransom, magazine)
}

func useHashmap(ransomNote, magazine string) bool {
    n := len(magazine)
    hset := make(map[byte]int, n)
    for i := range magazine {
        hset[magazine[i]]++
    }
    for i := range ransomNote {
        if v, exists := hset[ransomNote[i]]; !exists {
            return false
        } else {
            if v < 1 {
                return false
            }
            hset[ransomNote[i]]--
        }
    }
    return true
}
