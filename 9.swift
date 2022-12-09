struct Pos: Hashable {
    var x = 0, y = 0
}

struct Rope {
    var knots: [Pos] = []
    var visited: Set<Pos> = []

    init(_ nknots: Int) {
        for _ in 1...nknots { self.knots.append(Pos()) }
    }

    mutating func move(_ dx: Int, _ dy: Int) {
        knots[0].x += dx
        knots[0].y += dy
        for i in 1..<knots.count {
            if abs(knots[i - 1].x - knots[i].x) > 1 ||
               abs(knots[i - 1].y - knots[i].y) > 1 {
                knots[i].x += direction(knots[i].x, knots[i - 1].x)
                knots[i].y += direction(knots[i].y, knots[i - 1].y)
            }
        }
        visited.insert(knots.last!)
    }

    func direction(_ from: Int, _ to: Int) -> Int {
        return from == to ? 0 : (from > to ? -1 : 1)
    }
}

var rope1 = Rope(2), rope2 = Rope(10)
while let line = readLine() {
    var dx = 0, dy = 0, n = Int(line.dropFirst(2))!
    switch line.prefix(1) {
        case "R": dx = 1
        case "L": dx = -1
        case "D": dy = 1
        case "U": dy = -1
        default: break
    }
    for _ in 1...n { 
        rope1.move(dx, dy) 
        rope2.move(dx, dy) 
    }
}
print(rope1.visited.count, rope2.visited.count)