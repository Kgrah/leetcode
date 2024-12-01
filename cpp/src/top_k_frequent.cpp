#include <unordered_map>
#include <vector>
using namespace std;

vector<int> top_k_frequent(vector<int>& nums, int k) {
    unordered_map<int, int> freq_map;

    for (int n : nums) {
        freq_map[n]++;
    }

    vector<vector<int>> buckets(nums.size() + 1);

    for (const auto& [key, value] : freq_map) {
        buckets[value].push_back(key);
    }

    vector<int> result;

    for (int i = buckets.size() - 1; i > 0 && result.size() < k; i--) {
        for (int n : buckets[i]) {
            result.push_back(n);
            if (result.size() == k) {
                return result;
            }
        }
    }

    return result;
}
