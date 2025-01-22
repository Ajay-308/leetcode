class LFUCache {
public:
    int cap, n;
    unordered_map<int, list<vector<int>>::iterator> mp; // key -> iterator in freq
    map<int, list<vector<int>>> freq; // freq_count -> {key, value, freq}
    
    LFUCache(int capacity) {
        cap = capacity;
        n = 0;
    }

    void makeItMostRecentlyUsed(int key) {
        auto &vec = *(mp[key]); //passing it by refrence kyuki hume changes bhi to karna hai idher
        int val = vec[1];
        int f = vec[2];

        freq[f].erase(mp[key]);// adress humara kaha store hai mp mai isliye mp[key];
        //agar f pr ab koi or Doubly linked list hai hi nhi to f ko hi hata do na 
        if (freq[f].empty()) freq.erase(f); 

        f++; // Increase frequency
        freq[f].push_front({key, val, f}); // Add to the new frequency list
        mp[key] = freq[f].begin(); // Update iterator in mp
    }
    
    int get(int key) {
        if (mp.find(key) == mp.end()) return -1;  // mtlb nhi mila hume
        int val = (*(mp[key]))[1]; // mp mai to bas adress hai hume us address pr jo store hai vo chiye
        makeItMostRecentlyUsed(key);
        return val;
    }
    
    void put(int key, int value) {
        if (cap == 0) return;
        // ab agar jo banda aaya vo phele se hi exsit kar raha hai to 
        if (mp.find(key) != mp.end()) {
            (*(mp[key]))[1] = value;
            makeItMostRecentlyUsed(key);
        } 
        else {
            if (n == cap) { // Cache is full, remove LFU item
                auto &min_freq_list = freq.begin()->second;
                int key_to_del = min_freq_list.back()[0]; // Get the least recently used key
                min_freq_list.pop_back();
                if (min_freq_list.empty()) freq.erase(freq.begin()->first); // Remove empty list
                mp.erase(key_to_del); // Remove from map
            } 
            else {
                n++; // Increase size if not full
            }

            // Insert new element
            freq[1].push_front({key, value, 1});
            mp[key] = freq[1].begin();
        }
    }
};

/**
 * Example usage:
 * LFUCache* obj = new LFUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key, value);
 */
