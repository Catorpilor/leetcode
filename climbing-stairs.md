###[Climbing Staris][qurl]

##Solutions

```c++
class Solution {
public:
    int climbStairs(int n) {
        // Note: The Solution object is instantiated only once and is reused by each test case.
        if(n == 1 ) return 1;
        if(n == 2 ) return 2;
        int nret1 = 1,nret2=2;
        int ret;
        for(int i=3;i<=n;++i){
            ret = nret1+nret2;
            nret1 = nret2;
            nret2 = ret;
        }
        return ret;
    }
};
```


[qurl]:https://oj.leetcode.com/problems/climbing-stairs/
