package students

import "testing"

func TestNumOfStudents(t *testing.T) {
    st := []struct {
        name      string
        startTime []int
        endTime   []int
        queryTime int
        exp       int
    }{
        {"just one student", []int{4}, []int{4}, 4, 1},
        {"just one student diff query time", []int{4}, []int{4}, 5, 0},
        {"testcase1", []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{10, 10, 10, 10, 10, 10, 10, 10, 10}, 5, 5},
        {"testcase2", []int{1, 2, 3}, []int{3, 2, 7}, 4, 1},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := numOfStudents(tt.startTime, tt.endTime, tt.queryTime)
            if out != tt.exp {
                t.Fatalf("with startTime: %v, endTime %v and querTime:%d, wanted %d but got %d", tt.startTime, tt.endTime, tt.queryTime, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
