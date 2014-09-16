###[Population Next Right Pointers In Each Node II][qurl]

###Solutions
```c++
/**
 * Definition for binary tree with next pointer.
 * struct TreeLinkNode {
 *  int val;
 *  TreeLinkNode *left, *right, *next;
 *  TreeLinkNode(int x) : val(x), left(NULL), right(NULL), next(NULL) {}
 * };
 */
class Solution {
public:
    void connect(TreeLinkNode *root) {
        // Note: The Solution object is instantiated only once and is reused by each test case.
        TreeLinkNode* n = root;
         while (n) {
            TreeLinkNode * next = NULL; // the first node of next level
            TreeLinkNode * prev = NULL; // previous node on the same level
            for (; n; n=n->next) {
                if (!next) next = n->left?n->left:n->right;
                if (n->left) {
                    if (prev) prev->next = n->left;
                    prev = n->left;
                }
                if (n->right) {
                    if (prev) prev->next = n->right;
                    prev = n->right;
                }
            }
            n = next; // turn to next level
        }
    }
};
```

[qurl]:https://oj.leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
