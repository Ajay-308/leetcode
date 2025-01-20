// TLE SOLUTION 

class LRUCache {
public:
    
    vector<pair<int, int>> cache;
    int n;
    
    LRUCache(int capacity) {
        n = capacity;
    }
    
    int get(int key) {
        
        for(int i = 0; i<cache.size(); i++) {
            
            if(cache[i].first == key) {
                int val = cache[i].second;
                
                pair<int, int> temp = cache[i];
                cache.erase(cache.begin()+i);
                cache.push_back(temp);
                
                return val;
            }
            
        }
        
        return -1;
        
    }
    
    void put(int key, int value) {
        
        for(int i = 0; i<cache.size(); i++) {
            
            if(cache[i].first == key) {
                cache.erase(cache.begin()+i);
                cache.push_back({key, value});
                return;
            }
        }
        
        if(cache.size() == n) {
            cache.erase(cache.begin());
            cache.push_back({key, value});
        } else {
            cache.push_back({key, value});
        }
        
    }
};

// SUBMITEED SOLUTION 
class LRUCache {
public:
    list<int> dll; // using doubly linked list 
    // pair<adress,int> use karna hoga
    // adress ko access karte hai by ::iterator
    unordered_map<int,pair<list<int>::iterator,int>> mp;
    int n;
    LRUCache(int capacity) {
        n = capacity;
        
    }
    void makeItrecentlyUsed(int key){
        dll.erase(mp[key].first);// phele use delete karo
        dll.push_front(key); // suru mai push kardo
        mp[key].first = dll.begin(); // mp key ka first ko dll ke suru mai dal do 

    }
    
    int get(int key) {
        if(mp.find(key)==mp.end()) return -1; // mtlb ki nhi mila hume
        //or agar mil gya hai to use make it to recently used one
        makeItrecentlyUsed(key);
        return mp[key].second;
        
    }
    
    void put(int key, int value) {
        // agar key mil gya to
        if(mp.find(key)!=mp.end()){
            mp[key].second= value;
            makeItrecentlyUsed(key);
        }//agar nhi mila to
        else{
            dll.push_front(key);
            mp[key]= {dll.begin(),value};
            n--;
        }
        //agar humne extra element dal diya to 
        if(n<0){
            int ele_tobe_del = dll.back();
            mp.erase(ele_tobe_del);
            dll.pop_back();
            n++;
        }
    }
};
