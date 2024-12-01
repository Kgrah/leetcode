#include <string>
#include <vector>
using namespace std;

bool is_anagram(string s, string t) {
    vector<int> char_set_s;
    vector<int> char_set_t;

    for (char c : s) {
        char_set_s[c - 'A']++;
    }

    for (char c : t) {
        char_set_t[c - 'A']++;
    }

    for (int i = 0; i < char_set_s.size(); i++) {
        int count_s = char_set_s[i];

        if (char_set_t[i] != count_s) {
            return false;
        }
    }

    return true;
}