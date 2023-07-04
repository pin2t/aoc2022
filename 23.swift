import Foundation

struct Position: Hashable { var x, y: Int }

var elves = Set<Position>()
var y = 0
while let line = readLine() {
    for (x, c) in line.enumerated() {
        if c == "#" { elves.insert(Position(x: x, y: y)) }
    }
    y += 1
}
var (n1, n2, i) = (0, 0, 0)
while true {
    var moves = [Position: Position]()
    for e in elves {
        var north = false
        var south = false
        var west = false
        var east = false
        for d in [Position(0, -1), Position(0, 1), Position(-1, 0), Position(1, 0)] {
            if elves.contains(Position(x: e.x + d.x, y: e.y + d.y)) {
                if d.y == -1 {
                    north = true
                }
                if d.y == 1 {
                    south = true
                }
                if d.x == -1 {
                    west = true
                }
                if d.x == 1 {
                    east = true
                }
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

    var moved = [Position: Bool]()
    for (p, to) in moves {
        if wills[to] == 1 {
            moved[to] = true
            elves.remove(p)
        }
    }

    if moved.isEmpty {
        n2 = i
        break
    }

    for (e, _) in elves {
        moved[e] = true
    }

    elves = moved
    if i == 9 {
        var topleft = Position(x: 1000000, y: 1000000)
        var bottomright = Position(x: -1000000, y: -1000000)
        for e in elves {
            topleft.x = min(topleft.x, e.x)
            topleft.y = min(topleft.y, e.y)
            bottomright.x = max(bottomright.x, e.x)
            bottomright.y = max(bottomright.y, e.y)
        }
        var field = 0
        for x in topleft.x...bottomright.x {
            for y in topleft.y...bottomright.y {
                if elves.contains(Position(x: x, y: y)) {
                    field += 1
                }
            }
        }
        n1 = field
    }
}
print(n1, n2)
