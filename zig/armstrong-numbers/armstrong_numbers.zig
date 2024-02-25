const std = @import("std");

fn number_len(n: u128) usize {
    if (n == 0) return 1;
    return @as(usize, @intFromFloat(@floor(@log10(@as(f64, @floatFromInt(n))) + 1)));
}

pub fn isArmstrongNumber(n: u128) bool {
    const len = number_len(n);

    var digits: u128 = n;
    var sum: u128 = 0;

    while (digits != 0) : (digits /= 10) {
        const digit = digits % 10;

        const safeAdd = @addWithOverflow(sum, std.math.pow(u128, digit, len));
        if (safeAdd[1] != 0) return false;

        sum = safeAdd[0];
    }
    return sum == n;
}
