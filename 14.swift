import Foundation

var field: [Pos:Character] = [:]

while let line = readLine() {
    var from = Pos(line.numbers[0], line.numbers[1])
    var i = 2
    while i < line.numbers.count {
        let to = Pos(line.numbers[i], line.numbers[i + 1])
        let d = from.direction(to)
        while from != to { field[from] = "#"; from = from.move(d) }
        field[to] = "#"
        i += 2
    }
}
var n1 = simulate()
var floor = 0
for p in field.keys { floor = max(floor, p.y) }
for x in 0..<1000 { field[Pos(x, floor + 2)] = "#" }
print(n1, n1 + simulate() + 1)

func simulate() -> Int {
    var n = 0
    while true {
        var p = Pos(500, 0)
        while p.y < 999 {
            if !field.keys.contains(Pos(p.x, p.y + 1)) { 
                p = Pos(p.x, p.y + 1); 
            } else if !field.keys.contains(Pos(p.x - 1, p.y + 1)) { 
                p = Pos(p.x - 1, p.y + 1); 
            } else if !field.keys.contains(Pos(p.x + 1, p.y + 1)) { 
                p = Pos(p.x + 1, p.y + 1); 
            } else {
                field[p] = "o"
                break
            }
        }
        if p.y == 999 || p.x == 500 && p.y == 0 {
            break
        }
        n += 1
    }
    return n
}

struct Pos: Hashable { 
    var x, y: Int 
    init(_ x: Int, _ y: Int) { self.x = x; self.y = y }

    func direction(_ to: Pos) -> Pos {
        return Pos(
            to.x == self.x ? 0 : (to.x > self.x ? 1 : -1), 
            to.y == self.y ? 0 : (to.y > self.y ? 1 : -1)
        )
    }

    func move(_ dir: Pos) -> Pos {
        return Pos(self.x + dir.x, self.y + dir.y)
    }
}

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
