###[Minimum Depth of Binary Tree][qurl]

##Solutions

```c++
/**
 * Definition for binary tree
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    int minDepth(TreeNode *root) {
        // Note: The Solution object is instantiated only once and is reused by each test case.
        if(root==NULL) return 0;
        if(root->left==NULL && root->right==NULL ) return 1;
        int lmin=0,rmin=0;
        if(root->left) lmin = minDepth(root->left)+1;
        if(root->right) rmin = minDepth(root->right)+1;
        if(lmin==0 ) return rmin;
        if(rmin==0) return lmin;
        return lmin>rmin ? rmin : lmin;
    }
};
```


[qurl]:https://oj.leetcode.com/problems/minimum-depth-of-binary-tree/
