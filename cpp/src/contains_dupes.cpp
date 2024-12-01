#include <unordered_set>
#include <vector>
using namespace std;

bool contains_dupes(vector<int> nums) {
    unordered_set<int> seen;

    for (int n : nums) {
        if (seen.find(n) != seen.end()) {
            return true;
        }
        seen.insert(n);
    }

    return false;
}