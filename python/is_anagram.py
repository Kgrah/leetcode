def is_anagram(s: str, t: str):
    char_set_s = [26]
    char_set_t = [26]

    for c in s:
        char_set_s[ord(c.lower()) - ord('a')] += 1
    
    for c in t:
        char_set_t[ord(c.lower()) - ord('a')] += 1
    
    for i in range(len(char_set_s)):
        s_count = char_set_s[i]
        if char_set_t[i] != s_count:
            return False

    return True
