const std = @import("std");
const power_of_2 = @import("power_of_2.zig");
const contains_dupe = @import("contains_duplicate.zig");
const is_anagram = @import("anagram.zig");
const two_sum = @import("two_sum.zig");

pub fn main() !void {
    // Prints to stderr (it's a shortcut based on `std.io.getStdErr()`)
    std.debug.print("All your {s} are belong to us.\n", .{"codebase"});

    // stdout is for the actual output of your application, for example if you
    // are implementing gzip, then only the compressed bytes should be sent to
    // stdout, not any debugging messages.
    const stdout_file = std.io.getStdOut().writer();
    var bw = std.io.bufferedWriter(stdout_file);
    const stdout = bw.writer();
    try stdout.print("Run `zig build test` to run the tests.\n", .{});
    try bw.flush(); // don't forget to flush!
}

test "power of 2" {
    const result = power_of_2.powerOf2(10);
    const result_string = if (result) "true" else "false";
    std.debug.print("{s}\n", .{result_string});
    try std.testing.expect(result == false);
}

test "contains duplicate" {
    const list = [_]i32{ 10, 10, 20, 40, 60 };
    const result = try contains_dupe.containsDupe(list[0..]);
    const result_string = if (result) "true" else "false";
    std.debug.print("{s}\n", .{result_string});
    try std.testing.expect(result == true);
}

test "is_anagram" {
    const s = "test";
    const t = "ttes";

    const result = try is_anagram.isAnagram(s, t);
    const result_string = if (result) "true" else "false";
    std.debug.print("{s}\n", .{result_string});
    try std.testing.expect(result);
}

test "two_sum" {
    const target = 9;
    const numbers: [6]i32 = .{ 10, 5, 3, 6, 5, 5 };
    const numbers_slice: []const i32 = numbers[0..];
    const result = two_sum.twoSum(numbers_slice, target);
    try std.testing.expect(result == [2]i32{ 2, 3 });
}
