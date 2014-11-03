###[Maximum Product Subarray][qurl]

##Analysis
- DP catagory, so...using dp.
- The solution below comes from the [disscuss][disurl]
- And the comments share the whole idea. So i just paste it.
- Further learning <Programming Pearls II> related problems.

##Solution
```c++
class Solution {
public:
    int maxProduct(int A[], int n) {
         // store the result that is the max we have found so far
         int r = A[0];
        // imax/imin stores the max/min product of
        // subarray that ends with the current number A[i]
        for (int i = 1, imax = r, imin = r; i < n; i++) {
            // multiplied by a negative makes big number smaller, small number bigger
            // so we redefine the extremums by swapping them
            if (A[i] < 0)
                swap(imax, imin);
            // max/min product for the current number is either the current number itself
            // or the max/min by the previous number times the current one
            imax = max(A[i], imax * A[i]);
            imin = min(A[i], imin * A[i]);
            // the newly computed max value is a candidate for our global result
            r = max(r, imax);
        }
        return r;
    }
};
```

[qurl]:https://oj.leetcode.com/problems/maximum-product-subarray/
[disurl]:https://oj.leetcode.com/discuss/14235/possibly-simplest-solution-with-o-n-time-complexity
