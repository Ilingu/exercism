// Please implement the `ComputationError.IllegalArgument` error.

pub const ComputationError = error{IllegalArgument};

pub fn steps(n: usize) anyerror!usize {
    if (n == 0) return ComputationError.IllegalArgument;

    var i: usize = 0;
    var collatz: usize = n;

    while (collatz != 1) : (i += 1) {
        if (collatz % 2 == 0) {
            collatz /= 2;
        } else {
            collatz = 3 * collatz + 1;
        }
    }
    return i;
}
