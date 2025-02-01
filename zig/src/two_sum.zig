const std = @import("std");

pub fn twoSum(numbers: []const i32, target: i32) ![2]i32 {
    const allocator = std.heap.page_allocator;
    var comps = std.AutoHashMap(i32, i32).init(allocator);
    try comps.ensureTotalCapacity(@intCast(numbers.len));

    for (numbers, 0..) |n, i| {
        if (comps.get(target - n)) |idx| {
            return [2]i32{ idx, @intCast(i) };
        }

        try comps.put(n, i);
    }

    return null;
}
