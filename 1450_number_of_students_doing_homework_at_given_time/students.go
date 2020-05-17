package students

func numOfStudents(startTime, endTime []int, queryTime int) int {
    return useOnePass(startTime, endTime, queryTime)
}

func useOnePass(startTime, endTime []int, queryTime int) int {
    var ans int
    for i := range startTime {
        if startTime[i] <= queryTime && endTime[i] >= queryTime {
            ans++
        }
    }
    return ans
}
