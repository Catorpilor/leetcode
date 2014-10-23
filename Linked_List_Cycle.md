###[Linked List Cycle][qurl]
##Analysis
     很久之前看过一篇[文章][articlelink]讲解关于如何在单链表中确定是否有环,里面提到了O(n)解法就是使用快慢指针遍历链表，如果有环则慢指针==快指针，否则没环。这种算法有个名字叫做"Floyd's Cycle-Finding Algorithm"由Robert W. Floyed 1967年发表在"Non-deterministic Algorithms".又名"The Tortoise and the Hare Algorithm"

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
    bool hasCycle(ListNode *head) {
        //use slow and faster pointer
       if(head==NULL||head->next==NULL) return false;
       ListNode* slow = head;
       ListNode* fast = head;
       while(fast!=NULL&&fast->next!=NULL){
           slow = slow->next;
           fast = fast->next->next;
            if(slow == fast ) return true;
       }
       return false;
    }
};
```
[qurl]:https://oj.leetcode.com/problems/linked-list-cycle/
[articlelink]:http://ostermiller.org/find_loop_singly_linked_list.html
