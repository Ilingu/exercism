const std = @import("std");

pub fn isPangram(str: []const u8) bool {
    if (str.len < 26) return false;
    var alphabet = [_]bool{false} ** 26;
    var sum: u8 = 0;
    for (str) |c| {
        const index = std.ascii.toLower(c) % 26;
        if (!alphabet[index]) {
            alphabet[index] = true;
            sum += 1;
        }
    }
    return sum == 26;
}
