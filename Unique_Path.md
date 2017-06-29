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
                

##Update
Check this out!
The amazing solution i ever seen in [discuss][discussurl]! Which actually i am still confused...The solution and explantions as below
```c++
class Solution {
    public:
        int uniquePaths(int m, int n) {
            int N = n + m - 2;// how much steps we need to do
            int k = m - 1; // number of steps that need to go down
            double res = 1;
            // here we calculate the total possible path number 
            // Combination(N, k) = Combination(N, N - k)
            for (int i = 1; i <= k; i++)
                res = res * (N - k + i) / i;
            return (int)res;
        }
    };

```

- First of all you should understand that we need to do n + m - 2 movements : m - 1 down, n - 1 right, because we start from cell (1, 1).

- Secondly, the path it is the sequence of movements( go down / go right), therefore we can say that two paths are different when there is i-th (1 .. m + n - 2) movement in path1 differ i-th movement in path2.

- So, how we can build paths. Let us choose (n - 1) movements(number of steps to the right) from (m + n - 2), and rest (m - 1) is (number of steps down).
I think now it is obvious that count of different paths are all combinations (n - 1) movements from (m + n-2).

##Another Update
Using one dimension vector array to save the caculates. Extra space O(min{m,n})
Solution as below:
```c++
    // Reduces the problem to m <= n, because internal loop is faster.
    if (m > n) swap(m, n);
    if (m <= 1) return m; // Corner cases, one border is 0 or 1.
    // Uses a map to memoize previous results:
    vector<int> map(n, 1); // The map is filled of 1, so that first step is already done.
    // For every horizontal step...
    for (int i = 1; i < m; ++i) {
        // For every vertical step...
        for (int j = 1; j < n; ++j) {
            map[j] += map[j - 1];
        }
    }
    // The last element contains our solution.
    return map[n - 1];
}

```

[qurl]:https://oj.leetcode.com/problems/unique-paths/
[discussurl]:https://oj.leetcode.com/discuss/9110/my-ac-solution-using-formula
