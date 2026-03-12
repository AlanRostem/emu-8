const std = @import("std");

pub fn main() !void {
    // Prints to stderr, ignoring potential errors.
    std.debug.print("Hello, {s}!\n", .{"EMU-8"});
}
