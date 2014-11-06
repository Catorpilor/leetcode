###[Decode Ways][qurl]

##Analysis
- Edge case.If s[0]='0' then return 0;
- As Fibinacci sequences T[n]=T[n-1]+T[n-2] 

##Solution1 (O(n) space)
```c++
class Solution{
public:
    int numDecodings(string s){
        int n = s.size();
        if(n==0 || s[0]=='0') return 0;
        vector<int> v(n+1,0);
        v[0] = 1;
        v[1] = 1;
        for(int i = 2;i<=n;i++){
            if(s[i-1]!='0') v[i]+=v[i-1];
            int last2 = atoi(s.substr(i-2,2).c_str());
            if(last2>9 && last2<=26 ) v[i]+=v[i-2];
            if(v[i]==0) break; //two adjancet 0s
        }
        return v[n];
    }
}
```

##Solution2 (O(1) extra space)
```c++
class Solution{
public:
    int numDecodings(string s){
        int n = s.size();
        if(n==0||s[0]=='0') returnr 0; //edge cases
        int prev1 = 1, prev2 = 1 , cur = 0; //same as v[1] ,v[0]
        for(int i = 1; i<=n;++i){
            cur = 0;
            if(s[i-1]!='0') cur += prev1;
            if(i>=2){
                int last2 = atoi(s.substr(i-2,2).c_str());
                if(last2>9 && last2<=26) cur+=prev2;
                prev2 = prev1;
            }
            prev1 = cur;
        }
        return cur;
    }
}
```

[qurl]:https://oj.leetcode.com/problems/decode-ways/
