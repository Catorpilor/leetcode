###[Unique Binary Search Trees][qurl]
## Analysis
    *   Using DP
    *   Binary Search Trees have fixed format that is root is larger than its chile subtrees and smaller than its right child subtree;
    *   1...n given the fixed root k then its combiantion is T[k](its left child subtree combinations) * T[n-k](right subtree combinations)

## Solutions
```c++
class Solution {
public:
    int numTrees(int n) {
        vector<int> trees(n+1,0);
        trees[0] = 1;
        for( int i =1; i<=n; ++i){
            for( int j =0; j<i;++j) {
                trees[i] += trees[j] * trees[i-j-1];
            }
        }
        return trees[n];
    }
};
```


[qurl]:https://oj.leetcode.com/problems/unique-binary-search-trees/
