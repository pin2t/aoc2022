import Foundation

struct Pos: Hashable {
    var x, y: Int

    init(_ x: Int, _ y: Int) { self.x = x; self.y = y; }
}

var elves = Set<Pos>()
var y = 0
while let line = readLine() {
    for (x, c) in line.enumerated() {
        if c == "#" { elves.insert(Pos(x, y)) }
    }
    y += 1
}
var (n1, n2, i) = (0, 0, 0)
while true {
    var propositions = [Pos: Pos]()
    for e in elves {
        if elves.contains(where :{ $0 == Pos(e.x, e.y - 1) || $0 == Pos(e.x, e.y + 1) ||
                            $0 == Pos(e.x - 1, e.y) || $0 == Pos(e.x + 1, e.y) ||
                            $0 == Pos(e.x - 1, e.y - 1) || $0 == Pos(e.x + 1, e.y + 1) ||
                            $0 == Pos(e.x - 1, e.y + 1) || $0 == Pos(e.x + 1, e.y - 1) }) {
            let moves = [
                [Pos(e.x, e.y - 1), Pos(e.x, e.y - 1), Pos(e.x - 1, e.y - 1), Pos(e.x + 1, e.y - 1)],
    			[Pos(e.x, e.y + 1), Pos(e.x, e.y + 1), Pos(e.x - 1, e.y + 1), Pos(e.x + 1, e.y + 1)],
                [Pos(e.x - 1, e.y), Pos(e.x - 1, e.y), Pos(e.x - 1, e.y - 1), Pos(e.x - 1, e.y + 1)],
                [Pos(e.x + 1, e.y), Pos(e.x + 1, e.y), Pos(e.x + 1, e.y - 1), Pos(e.x + 1, e.y + 1)]
            ]
            for j in 0..<4 {
                let to = (i + j) % 4
                if !elves.contains(moves[to][1]) && !elves.contains(moves[to][2]) && !elves.contains(moves[to][3]) {
                    propositions[e] = moves[to][0]
                    break
                }
            }
        }
    }
    var npropositions = [Pos:Int]()
    for (from, to) in propositions {
        npropositions[to] = (npropositions[to] ?? 0) + 1
    }
    var moved = 0
    for (from, to) in propositions {
        if npropositions[to] == 1 {
            elves.remove(from)
            elves.insert(to)
            moved += 1
        }
    }
    if moved == 0 {
        n2 = i + 1
        break
    }
    if i == 9 {
        var top = 1000000, left = 1000000, bottom = -1000000, right = -1000000
        for e in elves {
            (left, right, top, bottom) = (min(left, e.x), max(right, e.x), min(top, e.y), max(bottom, e.y))
        }
        n1 = (bottom - top + 1) * (right - left + 1) - elves.count
    }
    i += 1
}
print(n1, n2)
