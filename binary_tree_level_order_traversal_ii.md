###[Binary Tree Level Order Traversal II][qurl]

##Analysis
  - [BFS][bfsurl] Initial an int variable to track the node count in each level and print level by level. Using a vector to store every node.


##Solution
```c++
/**
 * BFS version
 */
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
    vector<vector<int> > levelOrderBottom(TreeNode *root) {
        // Note: The Solution object is instantiated only once and is reused by each test case.
        vector<vector<int> > result;
        if(!root) return result;
        vector<TreeNode*> itemp;
        itemp.push_back(root);
        vector<int> temp;
        int cur = 0, last=0,lval;
        while(cur<itemp.size()){
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
        //reverse
        // vector<vector<int> >ssresult;
        // for(int i=result.size()-1;i>=0;--i) ssresult.push_back(result[i]);
        reverse(result.begin(),result.end());
        return result;
        
    }
};
```
[bfsurl]:http://en.wikipedia.org/wiki/Breadth-first_search
[qurl]:https://oj.leetcode.com/problems/binary-tree-level-order-traversal-ii/
