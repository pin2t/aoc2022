import Foundation

extension String {
    var numbers: [Int] {
        guard let regex = try? NSRegularExpression(pattern: "\\d+", options: [.caseInsensitive]) else { return [] }
        let matches  = regex.matches(in: self, options: [], range: NSMakeRange(0, self.count))
        return matches.map { match in
            String(self[Range(match.range, in: self)!])
        }.map { Int($0)! }
    }
}

class Stack {
    var crates: [Character] = []

    func move(to: Stack, n: Int) {
        return Stack
    }
}

var lines: [String] = []
while let line = readLine() {
    if line == "" { break }
    lines.append(line)
}
print(lines)
while let line = readLine() {
    print(line.numbers)
}
