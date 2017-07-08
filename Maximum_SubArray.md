###[Maximum Subarray][qurl]

##Analysis
- Idea is very simple. Basically, keep adding each integer to the sequence until the sum drops below 0. If sum is negative, then should reset the sequence

##Solutions
```c++
class Solution {
public:
    int maxSubArray(int A[], int n) {
        int sum = 0,maxsofar=A[0];
        for(int i=0;i<n;i++){
            sum += A[i];
            maxsofar = max(maxsofar,sum);
            sum = max(sum,0);
        }
        return maxsofar;
    }
};
```



[qurl]:https://oj.leetcode.com/problems/maximum-subarray/
