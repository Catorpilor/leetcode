###[Minimum Path Sum][qurl]

##Analysis
- Same as [Triangle Problem][trurl]
- Using bottom up strategy. Extra space O(n).
- Pre-caculate the bottom min path. From right to left we add it up.
- We use one demension to store the result. So for grid[i,j] the result[j] = min(result[j],result[j+1])+grid[i,j]. 

##Solution
```c++
class Solution {
public:
    int minPathSum(vector<vector<int> > &grid) {
        int m = grid.size();
        if( m==0) return 0;
        int n = grid[0].size();
        //bottom-up strategy
        //so it goes left and up.
        vector<int> result(grid.back());
        for(int i = n-2;i>=0;i--){
            result[i] += result[i+1];//pre-caculate the bottom paths.
        }
        for(int i = m-2;i>=0;i--){
            for(int j = n-1;j>=0;j--){
                if(j==n-1) result[j]+=grid[i][j];//edge senario
                else{
                    result[j]=min(result[j],result[j+1])+grid[i][j];
                }
            }
        }
        return result[0];
    }
};
```

[trurl]:https://github.com/Catorpilor/LeetCode/blob/master/Triangle.md
[qurl]:https://oj.leetcode.com/problems/minimum-path-sum/
