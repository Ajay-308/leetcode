class RandomizedSet {
public:
    //vector --> insert o(1) check 
    //       --> getRandom o(1) check 
    //       --> remove o(n) wrong index nhi hai humare pass number ka 
    // for getting index we are using another ds which is map
    unordered_map<int,int> mp;
    vector<int> v;

    RandomizedSet() {  
    }
    bool search(int x){
        if(mp.find(x)!=mp.end()) return true;
        return false;
    }
    
    bool insert(int val) {
        if(!search(val)){
            v.push_back(val);
            mp[val]=v.size()-1;
            return true;
        }
        return false;
        
    }
    
    bool remove(int val) {
        if(!search(val)) return false;
        auto it = mp.find(val);
        v[it->second] = v.back();
        v.pop_back();
        mp[v[it->second]]= it->second;
        mp.erase(val);
        return true;
    }
    
    int getRandom() {
        return v[rand()%v.size()];
        
    }
};

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * RandomizedSet* obj = new RandomizedSet();
 * bool param_1 = obj->insert(val);
 * bool param_2 = obj->remove(val);
 * int param_3 = obj->getRandom();
 */