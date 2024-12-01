def top_k_frequent(nums: list[int], k: int):
    freq_map = {}

    for n in nums:
        freq_map[n] = freq_map.get(n, 0) + 1

    buckets = [[] for _ in range(len(nums) + 1)]
    for n, freq in freq_map.items():
        buckets[freq].append(n)
    
    result = []
    for i in range(len(buckets) - 1, -1, -1):
        if len(result) == k:
            return result
        
        for n in buckets[i]:
            result.append(n)
            if len(result) == k:
                return result
    
    return result