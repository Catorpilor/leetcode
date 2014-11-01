###[Unique Paths][qurl]

##Analysis
Dynamic Programming using Top-Down.
- Edge check. for m = 1 or n = 1 only one way out so return 1;
- Let A[i,j] stands for From (0,0) to (i,j) 's unique paths.
- There are only 2 ways to A[i,j] : A[i-1,j] and A[i,j-1] those we've already caculated before. So A[i,j] = A[i-1,j] + A[i,j-1];

##Solutions
```c++
class Solution {
public:
    int uniquePaths(int m, int n) {
        //top-down
        if( m < 0 || n<0 ) return 0;
        if( m == 1 || n == 1) return 1;
        vector< vector<int> > result(m, vector<int>(n,1));
        for( int i = 1; i<m; i++){
            for(int j = 1; j<n;j++){
                result[i][j] = result[i][j-1] + result[i-1][j];
            }
        }
        return result[m-1][n-1];
    }
};
```


[qurl]:https://oj.leetcode.com/problems/unique-paths/
