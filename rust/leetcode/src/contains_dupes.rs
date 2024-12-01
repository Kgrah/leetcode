use std::collections::HashSet;

fn contains_duplicate(nums: Vec<i32>) -> bool {
    let mut seen: HashSet<i32> = HashSet::new();

    for n in nums {
        if seen.contains(&n) {
            return true;
        }

        seen.insert(n);
    }

    false
}
