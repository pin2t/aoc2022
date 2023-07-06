import Foundation

struct Position: Hashable {
    var x, y: Int

    init(_ x: Int, _ y: Int) { self.x = x; self.y = y; }
}

var elves = Set<Position>()
var y = 0
while let line = readLine() {
    for (x, c) in line.enumerated() {
        if c == "#" { elves.insert(Position(x, y)) }
    }
    y += 1
}
var (n1, n2, i) = (0, 0, 0)
while true {
    var moves = [Position: Position]()
    for e in elves {
        var north = false, south = false, west = false, east = false
        for d in [Position(0, -1), Position(0, 1), Position(-1, 0), Position(1, 0)] {
            if elves.contains(Position(e.x + d.x, e.y + d.y)) {
                if d.y == -1 { north = true }
                if d.y == 1 { south = true }
                if d.x == -1 { west = true }
                if d.x == 1 { east = true }
            }
        }
        if !north && !south && !west && !east {
            continue
        }
        let can = [(can: north, pos: Position(e.x, e.y - 1)),
                   (can: south, pos: Position(e.x, e.y + 1)),
                   (can: west, pos: Position(e.x - 1, e.y)),
                   (can: east, pos: Position(e.x + 1, e.y))]
        for j in 0..<4 {
            let to = (i + j) % 4
            if can[to].can {
                moves[e] = can[to].pos
                break
            }
        }
    }
    var wills = [Position: Int]()
    for (_, to) in moves {
        wills[to, default: 0] += 1
    }
    var moved = Set<Position>()
    for (p, to) in moves {
        if wills[to] == 1 {
            moved.insert(to)
            elves.remove(p)
        }
    }
    if moved.isEmpty {
        n2 = i
        break
    }
    for e in elves {
        moved.insert(e)
    }
    elves = moved
    if i == 9 {
        var top = 1000000, left = 1000000
        var bottom = -1000000, right = -1000000
        for e in elves {
            left = min(left, e.x);   top = min(top, e.y)
            right = max(right, e.x); bottom = max(bottom, e.y)
        }
        var field = 0
        for x in left...right {
            for y in top...bottom {
                if elves.contains(Position(x, y)) { field += 1 }
            }
        }
        n1 = field
    }
}
print(n1, n2)
