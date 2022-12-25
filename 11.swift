import Foundation

extension String {
    var numbers: [UInt64] {
        guard let regex = try? NSRegularExpression(
            pattern: "\\d+", 
            options: [.caseInsensitive]) else { 
            return [] 
        }
        let matches  = regex.matches(in: self, options: [], 
            range: NSMakeRange(0, self.count))
        return matches.map { match in
            String(self[Range(match.range, in: self)!])
        }.map { 
            UInt64($0)! 
        }
    }
}

class Monkey {
    var items: [UInt64] = [], destinations: [Int] = []
    var op: [String] = []
    var inspected: UInt64 = 1, div: UInt64
    var relief: Bool = true

    init(items: [UInt64], dest: [Int], op: [String], div: UInt64, relief: Bool) {
        self.items = items
        self.destinations = dest
        self.op = op
        self.div = div
    }
    func inspect(_ ms: inout [Monkey]) {
        let lcm: UInt64 = ms.map {$0.div}.reduce(1, *)
        for item in items {
            let arg: UInt64 = op[7] == "old" ? item : UInt64(op[7])!
            var level: UInt64 = op[6] == "+" ? item + arg : item * arg
            if relief { level /= 3 }
            level %= lcm
            let to = level % self.div == 0 ? destinations[0] : destinations[1]
            ms[to].items.append(level)
        }
        inspected += UInt64(items.count)
        self.items = []
    }
}

var monkeys: [Monkey] = [], monkeys2: [Monkey] = []

while let line = readLine() {
    if line.hasPrefix("Monkey ") {
        let lines = [readLine(), readLine(), readLine(), readLine(), readLine()]
        monkeys.append(Monkey(
            items: lines[0]!.numbers,
            dest: [Int(lines[3]!.numbers[0]), Int(lines[4]!.numbers[0])],
            op: lines[1]!.components(separatedBy: [" "]),
            div: lines[2]!.numbers[0],
            relief: true
        ))
        monkeys2.append(Monkey(
            items: lines[0]!.numbers, 
            dest: [Int(lines[3]!.numbers[0]), Int(lines[4]!.numbers[0])],
            op: lines[1]!.components(separatedBy: [" "]),
            div: lines[2]!.numbers[0],
            relief: false
        ))
   }
}
for _ in 0..<20 {
    for m in monkeys {
        m.inspect(&monkeys) 
    }
}
for _ in 0..<10000 {
    for m in monkeys2 { 
        m.inspect(&monkeys2) 
    }
}
monkeys.sort { $1.inspected < $0.inspected }
monkeys2.sort { $1.inspected < $0.inspected }
print(monkeys[0].inspected * monkeys[1].inspected, monkeys2[0].inspected * monkeys2[1].inspected)