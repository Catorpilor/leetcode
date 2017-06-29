###[Construct Binary Tree From Preorder and Inorder Traversal][qurl]

##Source code

```c++
class Solution {
public:
    TreeNode* buildTree(vector<int> &preorder,int lsi,int lei,vector<int> &inorder, int esi, int eei){
        if(lei - lsi < 0) return NULL;       
        TreeNode* pnode = new TreeNode(preorder[lsi]);
        int c = -1;
        for( c = esi ; c<=eei;++c){
            if(inorder[c] == pnode->val ) break;
        }
        //spilt into two part
        int lcount = c-esi;
        pnode->left = buildTree(preorder,lsi+1,lsi+lcount,inorder,esi,esi+lcount);
        pnode->right = buildTree(preorder,lsi+lcount+1,lei,inorder,c+1,eei);
        return pnode;        
    }
    TreeNode *buildTree(vector<int> &preorder, vector<int> &inorder) {
        // Note: The Solution object is instantiated only once and is reused by each test case.
        if(preorder.size() != inorder.size() || preorder.size() < 1 && inorder.size()<1 ) return NULL;
        return buildTree(preorder,0,preorder.size()-1,inorder,0,inorder.size()-1);
    }
};
```


[qurl]:https://oj.leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
