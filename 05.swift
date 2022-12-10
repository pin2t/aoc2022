import Foundation

extension String {
    var numbers: [Int] {
        guard let regex = try? NSRegularExpression(pattern: "\\d+", options: [.caseInsensitive]) else { 
            return [] 
        }
        let matches  = regex.matches(in: self, options: [], range: NSMakeRange(0, self.count))
        return matches.map { match in
            String(self[Range(match.range, in: self)!])
        }.map { 
            Int($0)! 
        }
    }
}

class Stack: NSObject {
    var crates: [Character] = []
    override var description: String { "Stack \(crates)" }

    override init() {}
    init(_ lines: [String], _ column: Int) {
        var i = lines.count - 1
        let c = String.Index(utf16Offset: column, in: lines[0])
        while i >= 0 {
            if lines[i][c] != " " { crates.append(lines[i][c]) }
            i -= 1
        }
    }

    func move(_ to: Stack, _ n: Int) {
        for _ in 1...n {
            to.crates.append(self.crates.last!)
            self.crates.removeLast()
        }
    }
}

func tops(_ stacks: [Stack]) -> String {
    var result = ""
    for stack in stacks { result.append(stack.crates.last!) }
    return result
}

var lines: [String] = []
while let line = readLine() {
    if line == "" { break }
    lines.append(line)
}
lines.removeLast()
var stacks: [Stack] = [], stacks2: [Stack] = []
var i = 1
while i < lines.last!.count {
    stacks.append(Stack(lines, i))
    stacks2.append(Stack(lines, i))
    i += 4
}
while let line = readLine() {
    let ns = line.numbers
    stacks[ns[1] - 1].move(stacks[ns[2] - 1], ns[0])
    let tmp = Stack()
    stacks2[ns[1] - 1].move(tmp, ns[0])
    tmp.move(stacks2[ns[2] - 1], ns[0])    
}
print(tops(stacks), tops(stacks2))