const std = @import("std");

pub fn containsDupe(list: []const i32) !bool {
    const allocator = std.heap.page_allocator;
    var mp = std.AutoHashMap(i32, bool).init(allocator);

    try mp.ensureTotalCapacity(@intCast(list.len));

    for (list) |n| {
        if (mp.get(n)) |_| {
            return true;
        }
        try mp.put(n, true);
    }

    return false;
}
