###[Add Two Numbers][qurl]

##Analysis
同样的2年前写的[Blog][blogurl]

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
    ListNode *addTwoNumbers(ListNode *l1, ListNode *l2) {
        // Note: The Solution object is instantiated only once and is reused by each test case.
        bool bret = false;
        ListNode* result = l1;
        ListNode* tmp =l1->next;
        while(l1!=NULL||l2!=NULL){
            if(l1!=NULL && l2!=NULL){
                l1->val += l2->val;
                if(l1->val >= 10){
                    bret=true;
                    l1->val %= 10;
                    if(l1->next) {
                        tmp = l1->next;
                        l1->next->val += 1;
                    }
                    else{
                        if(l2->next){
                            tmp = l2->next;
                            l2->next->val += 1;
                            l2->next = NULL;
                            l1->next = tmp;
                        }
                        else{
                            l1->next = new ListNode(1);
                        }
                    }
                }
                else{
                    if(l1->next) {tmp = l1->next;}
                    else{
                        if(l2->next){
                            tmp = l2->next;
                            l2->next = NULL;
                            l1->next = tmp;
                        }
                    }
                }
                l1 = tmp;
                l2 = l2->next;
            //curnode = curnode->next;
            }
            else{
                ListNode* temp = tmp;
                while(temp!=NULL){
                    if(temp->val == 10 ){
                        temp->val %=10;
                        if(temp->next==NULL){
                            temp->next = new ListNode(1);
                            break;
                        }
                        temp->next->val += 1;
                    }
                    temp = temp->next;
                }
                break;
            }
        }
    //result->next = tmp;
        return result;
    }
};
```

[qurl]:https://oj.leetcode.com/problems/add-two-numbers/
[blogurl]:http://blog.catorpilor.com/2012/07/20/interview-test-add-two-numbers.html
