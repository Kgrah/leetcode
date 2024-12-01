from collections import defaultdict

def group_anagrams(strs: list[str]) -> list[list[str]]:
    freq_buckets = defaultdict(list)

    for s in strs:
        freq_count = [0] * 26

        for c in s:
            freq_count[ord(c.lower() - ord('a'))] += 1

        freq_buckets[tuple(freq_count)].append(s)

    
    return list(freq_buckets.values())