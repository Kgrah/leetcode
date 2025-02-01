pub fn powerOf2(n: i32) bool {
    return n > 0 and (n & (n - 1)) == 0;
}
