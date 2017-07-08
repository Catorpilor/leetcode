###[Remove Duplicates from Sorted List][qurl]

##Analysis
跟数组的除重一样，链表的更为简单些只需要处理好边界就行.

##Solutions
```c++
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode *deleteDuplicates(ListNode *head) {
        ListNode* pori = head;
        ListNode* newh = head->next;
        while(newh!=NULL){
            if( newh->val == pori->val ) {
                newh = newh->next;
                if(newh == NULL ) pori->next = NULL;
            }else{
                pori->next = newh;
                pori = newh;
                newh = newh->next;
            }
        }
        return head;
    }
};
```

[qurl]:https://oj.leetcode.com/problems/remove-duplicates-from-sorted-list/
