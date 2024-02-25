fn to_lowercase(c: u8) u8 {
    return switch (c) {
        65...90 => c + 32,
        else => c,
    };
}

pub fn score(s: []const u8) u32 {
    var cscore: u32 = 0;
    for (s) |c| {
        cscore += switch (to_lowercase(c)) {
            'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't' => 1,
            'd', 'g' => 2,
            'b', 'c', 'm', 'p' => 3,
            'f', 'h', 'v', 'w', 'y' => 4,
            'k' => 5,
            'j', 'x' => 8,
            'q', 'z' => 10,
            else => 0,
        };
    }
    return cscore;
}
