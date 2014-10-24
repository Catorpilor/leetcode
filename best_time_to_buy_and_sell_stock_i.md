###[Best Time to Buy and Sell Stock][qurl]

##Analysis
    解法1：找出相邻元素差值的和连续递增的sub-array
    解法2：用游标指向当前最小元素，遍历数组找出最大差值

##Solutions
```c++
//解法1
class Solution {
public:
    int maxProfit(vector<int> &prices) {
        if( prices.size() <= 1 ) return 0;
        int max=0;
        int current=0;
        for(int i=1;i<prices.size();i++){
            current+=prices[i]-prices[i-1];
            if(current<0) current=0;
            else if(current>max) max=current;
        }
        return max;
    }
};

//解法2
class Solution {
public:
    int maxProfit(vector<int> &prices) {
        if( prices.size() <= 1 ) return 0;
        int min=prices[0];
        int max=0;
        for(int i=1;i<prices.size();i++){
            if(prices[i]<min){
                min = prices[i];
            }
            if(prices[i] - min > max) {
                max = prices[i] - min;
            }
        }
        return max;
    }
};
```

[qurl]:https://oj.leetcode.com/problems/best-time-to-buy-and-sell-stock/
