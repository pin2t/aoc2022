import Foundation

var sum = 0
let digits: [Character: Int] = ["=": -2, "-": -1, "0": 0, "1": 1, "2": 2]
let snafuDigits = ["0", "1", "2", "=", "-"]
while let snafu = readLine() {
    var n = 0
    for digit in snafu {
        if let value = digits[digit] {
            n = n * 5 + value
        }
    }
    sum += n
}
var snafu = ""
while sum > 0 {
    snafu = snafuDigits[sum % 5] + snafu
    if sum % 5 > 2 { sum += 5 }
    sum = sum / 5
}
print(snafu)