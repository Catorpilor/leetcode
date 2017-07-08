###[Word Break][qurl]

##Analysis
 - To be honest,no clue at all
 - Check the [disscuss][durl] 
 - And thanks to [shibai86][dpauthor]'s elegant code and finanly figured it out.

##Solutions(not mine)
```c++
class Solution {
public:
    bool wordBreak(string s, unordered_set<string> &dict) {
        int len = s.length();
        vector<bool> dp(len + 1,false);
        dp[len] = true; //empty s
        for (int i = len - 1; i >= 0; i--) {
            for (int j = i; j < len; j++) {
                string str = s.substr(i,j - i + 1);
                if (dict.find(str) != dict.end() && dp[j + 1]) {
                    dp[i] = true; //s[i:j] segmented by dict
                    break;
                }
            }
        }
        return dp[0];
    }
};
```

[qurl]:https://oj.leetcode.com/problems/word-break/
[durl]:https://oj.leetcode.com/discuss/1523/who-can-show-me-a-dp-solution-thanks
[dpauthor]:https://oj.leetcode.com/discuss/user/shibai86
