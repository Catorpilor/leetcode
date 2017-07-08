###[Binary Tree Level Order Traversal][qurl]

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
    vector<vector<int> > levelOrder(TreeNode *root) {
        // Note: The Solution object is instantiated only once and is reused by each test case.
        vector<vector<int> > result;
        if(!root) return result;
        vector<TreeNode*> itemp;
        itemp.push_back(root);
        vector<int> temp;
        int cur = 0, last=0,lval;
        while( cur<itemp.size() ){
            last = itemp.size();
            lval = cur;
            while(lval<last){
                temp.push_back(itemp[lval]->val);
                if(itemp[lval]->left) itemp.push_back(itemp[lval]->left);
                if(itemp[lval]->right) itemp.push_back(itemp[lval]->right);
                lval++;
            }
            vector<int> pp(temp.begin()+cur,temp.begin()+last);
            result.push_back(pp);
            cur = lval;
        }
        return result;
    }
};
```

[qurl]:https://oj.leetcode.com/problems/binary-tree-level-order-traversal/
