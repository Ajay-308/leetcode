// whenever you see a problem like i , j and some condition on them,
// like i<j and some thing like that 
// i <j and j-i != nums[j]- nums[i]
// then use concept of same side of equation 

// nums[i] - i  = nums[j] - j
// jab bhi 2 pair mere same ho gya to vo good pair ho gye
// [ 4 , 0 , 1 ,0 ] or mai last ke 0 pr hu to total mai pair kitne bana sakta hu 
// 3 but isme good pair kitne hai last ke 0 hai or ek bich mai aa raha hai to 1 baar
// aaya last ke 0 se phele to ==> 3 - 1 =2 ;

long long countBadPairs(vector<int>& nums) {
    unordered_map<int,int>mp;
    int n = nums.size();
    ll total_pair = (1LL * n * (n - 1)) / 2;
    ll good=0;
    for(int i=0;i<n;i++){
        int x = nums[i]-i;
        good += mp[x];
        mp[x]++;
    }
    return total_pair-good;
    
}