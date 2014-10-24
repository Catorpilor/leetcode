###[Best Time to Buy and Sell Stock III][qurl]

##Analysis    
    First consider the case that when we are only allowed to make one transaction. we can handle this easily with DP. If we move forward, any new price we meet will only affect our history result by two ways:will it be so low that it beats our previous lowest price? will it be so high that we should instead sell on this time to gain a higher profit (than the history record)? Similarly, we can move backward with the highest price and profit in record. Either way would take O(n) time. Now consider the two transaction case. Since there will be no overlaps, we are actually dividing the whole time into two intervals.We want to maximize the profit in each of them so the same method above will apply here. We are actually trying to break the day at each time instance, by adding the potential max profit before and after it together. By recording history and future for each time point, we can again do this within O(n) time.

##Solutions
```c++
class Solution {
public:
    int maxProfit(vector<int> &prices) {
        // null check
        int len = prices.size();
        if (len==0) return 0;
        vector<int> historyProfit;
        vector<int> futureProfit;
        historyProfit.assign(len,0);
        futureProfit.assign(len,0);
        int valley = prices[0];
        int peak = prices[len-1];
        int maxProfit = 0;
        // forward, calculate max profit until this time
        for (int i = 0; i<len; ++i)
        {
            valley = min(valley,prices[i]);
            if(i>0)
            {
                historyProfit[i]=max(historyProfit[i-1],prices[i]-valley);
            }
        }
        // backward, calculate max profit from now, and the sum with history
        for (int i = len-1; i>=0; --i)
        {
            peak = max(peak, prices[i]);
            if (i<len-1)
            {
                futureProfit[i]=max(futureProfit[i+1],peak-prices[i]);
            }
            maxProfit = max(maxProfit,historyProfit[i]+futureProfit[i]);
        }
        return maxProfit;
    }
};
```

[qurl]:https://oj.leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
