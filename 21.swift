import Foundation

var monkeys: [String: Monkey] = [:]

class Monkey {
    var n: Double
    var wait: [String]
    var op: String

    init(_ n: Double, _ wait: [String], _ op: String) {
        self.n = n
        self.wait = wait
        self.op = op
    }

    func yell() -> Double {
        if n > 0 { return n }
        let yell1 = monkeys[wait[0]]!.yell()
        let yell2 = monkeys[wait[1]]!.yell()
        switch op {
        case "+": return yell1 + yell2
        case "-": return yell1 - yell2
        case "*": return yell1 * yell2
        case "/": return yell1 / yell2
        default:  return 0;
        }
    }
}

while let line = readLine() {
    let parts = line.split(separator: " ")
    let name = String(String(parts[0]).prefix(4))
    if parts.count == 2 {
        let n = Double(parts[1])!
        monkeys[name] = Monkey(n, [], "")
    } else if parts.count == 4 {
        monkeys[name] = Monkey(0, [String(parts[1]), String(parts[3])], String(parts[2]))
    }
}
let n1 = monkeys["root"]!.yell()
var low: Int64 = 1
var high: Int64 = 1_000_000_000_000_000
monkeys["humn"] = Monkey(Double(low + (high - low) / 2), [], "")
let root = monkeys["root"]!
root.op = "-"
while true {
    let cmp = root.yell()
    if cmp > 0 { low += (high - low) / 2 }
    else if cmp < 0 { high -= (high - low) / 2 }
    else if cmp == 0 {
        print("\(Int64(n1)) \(Int64(monkeys["humn"]!.yell()))")
        break
    }
    monkeys["humn"] = Monkey(Double(low + (high - low) / 2), [], "")
}
