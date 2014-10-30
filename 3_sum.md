###[3 Sum][qurl]

##Analysis
2年前写了关于TA的[blog][blogurl]


##Solutions
    代码比较ugly...
```c++
class Solution {
public:
    vector<vector<int> > threeSum(vector<int> &num) {
        // Note: The Solution object is instantiated only once and is reused by each test case.
        vector<vector<int>> vresult;
        if(num.size()<3) return vresult;
        sort(num.begin(),num.end());
        //vector<vector<int>> vresult;
    if(num.at(0)>0 || num.at(num.size()-1)<0){
        return vresult;
    }
    int i,j,k=num.size()-1;
    if(num.size()==3)
    {
        if(num[0]+num[1]+num[2]==0)
        {
            vresult.push_back(num);
            return vresult;
        }
    }
    else if(num.empty() || num.size()<3)
    {
        return vresult;
    }
    else if(num[0]>=0 || num[num.size()-1]<=0)
    {
        if(num[0]==0 && num[num.size()-1]==0)
        {
            vector<int> v(3);
            vresult.push_back(v);
        }
        return vresult;
    }
    for(i=0;num[i]<=0;)
    {
        /*for(j=i+1;j<k;)
        {
        }*/
        j = i+1;
        k=num.size()-1;
        for(;k>j;)
        {
            if(num[i]+num[j]+num[k] < 0 )
            {
                j++;
            }
            else if(num[i]+num[j]+num[k] > 0 )
            {
                k--;
            }
            else
            {
                int count1=1;
                int count2=1;
                vector<int> v;
                v.push_back(num[i]);
                v.push_back(num[j]);
                v.push_back(num[k]);
                vresult.push_back(v);
                int temp1=j,temp2=k;
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
                j += count1;
                k -= count2;
            }
        }
        int count=1;
        int temp =i;
        while(num[temp+1] == num[temp] )
        {
            temp++;
            count++;
        }
        i+= count;
    }
    return vresult;
    }
};
```


[qurl]:https://oj.leetcode.com/problems/3sum/
[blogurl]:http://blog.catorpilor.com/2012/07/24/mulitiple-sum-problem.html
