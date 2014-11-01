###[Unique Paths II][qurl]

##Analysis
- Same as Unique Paths I
- Edge check for A[0..m-1,0] if A[k,0] = 1 then A[k+1..m-1,0] = 0
- Edge check for A[0,0..n-1] same as above.

##Solutions
```c++
class Solution {
public:
    int uniquePathsWithObstacles(vector<vector<int> > &obstacleGrid) {
        int m = obstacleGrid.size();
        int n = obstacleGrid[0].size();
        if( m == 0 || n== 0 ) return 0;
        vector<vector<int> >result(m,vector<int>(n,1));
        //caculate result[i][0] for i=0..m-1
        bool bflag = false;
        for( int i = 0;i<m;i++){
            if(obstacleGrid[i][0]==1||bflag){
                result[i][0] = 0;
                bflag = true;
            }
        }
        //caculate result[0][i] for i=0..n-1
        bflag = false;
        for(int i = 0; i<n;i++){
            if(obstacleGrid[0][i]==1 || bflag ){
                result[0][i] = 0;
                bflag = true;
            }
        }
        //same as unique paths
        for( int i =1; i<m;i++){
            for(int j = 1;j<n;j++){
                if(obstacleGrid[i][j] == 1 ) result[i][j] = 0;
                else result[i][j] = result[i-1][j]+result[i][j-1];
            }
        }
        return result[m-1][n-1];
    }
};
```

[qurl]:https://oj.leetcode.com/problems/unique-paths-ii/
