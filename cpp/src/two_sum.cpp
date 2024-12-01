#include <unordered_map>
#include <vector>
using namespace std;

vector<int> two_sum(vector<int> nums, int target) {
    unordered_map<int, int> comps;

    vector<int> result;
    for (int i = 0; i < nums.size(); i++) {
        int n = nums[i];
        comps[n] = i;

        if (comps.find(target - n) != comps.end()) {
            int idx = comps[target - n];

            result.push_back(i);
            result.push_back(idx);

            return result;
        }
    }

    return result;
}