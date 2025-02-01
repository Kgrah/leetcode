const std = @import("std");

pub fn isAnagram(s: []const u8, t: []const u8) !bool {
    if (t.len != s.len) {
        return false;
    }

    var char_set_s: [26]i32 = undefined;
    var char_set_t: [26]i32 = undefined;

    for (s) |c| {
        const idx = c - 'a';
        char_set_s[idx] += 1;
    }

    for (t) |c| {
        const idx = c - 'a';
        char_set_t[idx] += 1;
    }

    for (char_set_s, 0..) |count, i| {
        if (count != char_set_t[i]) {
            return false;
        }
    }

    return true;
}
