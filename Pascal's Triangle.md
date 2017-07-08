###[Pascal's Triangle][qurl]

##Solutions
```c++
class Solution {
public:
    vector<vector<int> > generate(int numRows) {
        // Note: The Solution object is instantiated only once and is reused by each test case.
        vector<vector<int> > result;
        if (numRows <= 0) return result;
        result.push_back(vector<int>(1,1));
        for (int i = 1; i < numRows; i++) {
            vector <int> vec;
            for (int j = 0; j <= i; j++) {
                if (j == 0 || j == i) vec.push_back(1);
                else vec.push_back(result.back()[j-1] + result.back()[j]);
            }
            result.push_back(vec);            
        }
        return result;
    }
};
```

[qurl]:https://oj.leetcode.com/problems/pascals-triangle/
