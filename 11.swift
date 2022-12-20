import Foundation

extension String {
    var numbers: [Int] {
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
            Int($0)! 
        }
    }
}

struct Monkey {
    var items: [Int] = [], destinations: [Int] = []
    var op: [String] = []
    var inspected: Int = 0, div: Int = 1

    mutating func inspect(_ ms: [Monkey], _ relief: Bool) {
        for item in self.items {
            let lcm = ms.map { $0.div }.reduce(1, *)
            var level: Int = 0, arg: Int = op[1] == "old" ? item : Int(op[1])!
            switch op[0] {
                case "+": level = item + arg
                case "*": level = item * arg
                default: break
            }
            if relief { level /= 3 }
            level %= lcm
            if level % self.div == 0 {
                to = lines[3]!.numbers[0]
            } else {
                to = lines[4]!.numbers[0]
            }
        }
        self.inspected += self.items.count
        self.items = []
    }
}

var monkeys: [Monkey] = [], monkeys2: [Monkey] = []
while let line = readLine() {
    if line.hasPrefix("Monkey ") {
        let lines = [readLine(), readLine(), readLine(), readLine(), readLine()]
        print("items", lines[0]!.numbers, 
            "formula", lines[1]!.components(separatedBy: " "),
            "div", lines[2]!.numbers[0], "throw", lines[3]!.numbers[0], lines[4]!.numbers[0])
        var monkey = Monkey() { item in 
            let lcm = monkeys.map { $0.div }.reduce(1, *)
            print("lcm", lcm)
            let op: [String] = lines[1]!.components(separatedBy: " ").suffix(2)
            var level: Int = 0, arg: Int = op[1] == "old" ? item : Int(op[1])!
            switch op[0] {
                case "+": level = (item + arg / 3) % lcm
                case "*": level = (item * arg / 3) % lcm
                default: break
            }
            var to: Int
            if level % lines[2]!.numbers[0] == 0 {
                to = lines[3]!.numbers[0]
            } else {
                to = lines[4]!.numbers[0]
            }
            monkeys[to].items.append(level)
        }
        monkey.div = lines[2]!.numbers[0]
        monkey.items = lines[0]!.numbers
        monkeys.append(monkey)
        // monkey = Monkey() { item in 
        //     let lcm = monkeys2.map { $0.div }.reduce(1, *)
        //     print("lcm", lcm)
        //     let op: [String] = lines[1]!.components(separatedBy: " ").suffix(2)
        //     var level: Int = 0, arg: Int = op[1] == "old" ? item : Int(op[1])!
        //     switch op[0] {
        //         case "+": level = (item + arg) % lcm
        //         case "*": level = (item * arg) % lcm
        //         default: break
        //     }
        //     var to: Int
        //     if level % lines[2]!.numbers[0] == 0 {
        //         to = lines[3]!.numbers[0]
        //     } else {
        //         to = lines[4]!.numbers[0]
        //     }
        //     monkeys2[to].items.append(level)
        // }
        // monkey.div = lines[2]!.numbers[0]
        // monkey.items = lines[0]!.numbers
        // monkeys2.append(monkey)
   }
}
for _ in 0..<20 {
    for i in 0..<monkeys.count { 
        monkeys[i].inspectAll() 
    }
}
for _ in 0..<10000 {
    for i in 0..<monkeys2.count { 
        monkeys2[i].inspectAll() 
    }
}
monkeys.sort { $1.inspected > $0.inspected }
monkeys2.sort { $1.inspected > $0.inspected }
print(monkeys.suffix(2)[0].inspected * monkeys.suffix(2)[1].inspected,
      monkeys2.suffix(2)[0].inspected * monkeys2.suffix(2)[1].inspected)