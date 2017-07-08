###[4Sum][qurl]

##Analysis
见2年前写的[blog][blogurl]

##Solution
代码依旧很Ugly
```c++
class Solution {
public:
    vector<vector<int> > fourSum(vector<int> &num, int target) {
        // Note: The Solution object is instantiated only once and is reused by each test case.
        vector<vector<int> > result;
    if(num.size()<4)
    {
        return result;
    }
    sort(num.begin(),num.end());
    if(num.size()==4)
    {
        if(num[0]+num[1]+num[2]+num[3] == target)
        {
            result.push_back(num);
        }
        return result;
    }
    //handle more than 4 elements
    int i,j,k,l;
    for(i=0; i<=num.size()-4;)
    {
        l = num.size()-1;
        for(j=i+1;j<l-1;)
        {
            l=num.size()-1;
            for(k=j+1;k<l;)
            {
                if(num[i]+num[j]+num[k]+num[l] < target )
                {
                    k++;
                }
                else if(num[i]+num[j]+num[k]+num[l] > target )
                {
                    l--;
                }
                else
                {
                    vector<int> v;
                    v.push_back(num[i]);
                    v.push_back(num[j]);
                    v.push_back(num[k]);
                    v.push_back(num[l]);
                    result.push_back(v);
                    //update k,l
                    if(k+1 == l || k+1==l-1)
                    {
                        break;
                    }
                    int temp1=k,temp2=l;
                    int count1=1,count2=1;
                    while(num[temp1+1] == num[temp1] && temp1 < temp2)
                    {
                        count1++;
                        temp1++;
                    }
                    while(num[temp2-1] == num[temp2] && temp2>temp1)
                    {
                        count2++;
                        temp2--;
                    }
                    k += count1;
                    l -= count2;
                }
            }
            //update j
            int count =1;
            int temp = j;
            while(num[temp+1]==num[temp])
            {
                count++;
                temp++;
            }
            j+=count;
        }
        //update i
        int count =1;
        int temp = i;
        while(num[temp+1]==num[temp])
        {
            count++;
            temp++;
        }
        i+=count;
    }
    return result;
    }
};
```

[qurl]:https://oj.leetcode.com/problems/4sum/
[blogurl]:http://blog.catorpilor.com/2012/07/24/mulitiple-sum-problem.html
