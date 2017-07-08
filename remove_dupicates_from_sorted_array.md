###[Remove Duplicates from Sorted Array][qurl]

##Analysis
简单的一题，关键是理解题目中的`Do not allocate extra space for another array, you must do this in place with constant memory`,只需要将下一个与它不同的值交换一下就行.

##Solutions
```c++
class Solution {
public:
    int removeDuplicates(int A[], int n) {
        if(n<2) return n;
        int curpos =0;
        int pswapos = 0;
        int flag = A[curpos];
        int ncount = 0;
        for(int i=curpos+1;i<n;++i){
            if( A[i]==flag ) ncount++;
            else{
                //swap A[i] with A[pswapos]
                pswapos++;
                A[pswapos] = A[i];
                curpos = i;
                flag = A[i];
            }
        }
        return n-ncount;
    }
};
```

[qurl]:https://oj.leetcode.com/problems/remove-duplicates-from-sorted-array/
