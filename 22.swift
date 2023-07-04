import Foundation

struct Point: Hashable {
    var x, y: Int
}

var walls = Set<Point>()
var tiles = Set<Point>()
var rows = 0
var size = 50

func bounds(x: Int, y: Int) -> (Int, Int, Int, Int) {
    var left = 10000, right = 0, top = 10000, bottom = 0
    for p in walls {
        if p.y == y || p.x == x {
            left, top = min(left, p.x), min(top, p.y)
            right, bottom = max(right, p.x), max(bottom, p.y)
        }
    }
    for p in tiles {
        if p.y == y || p.x == x {
            left, top = min(left, p.x), min(top, p.y)
            right, bottom = max(right, p.x), max(bottom, p.y)
        }
    }
    return (left, right, top, bottom)
}

var state = (pos: Point(x: 0, y: 0), dx: 0, dy: 0)

func step1() {
    var np = Point(x: state.pos.x + state.dx, y: state.pos.y + state.dy)
    let (left, right, top, bottom) = bounds(x: state.pos.x, y: state.pos.y)
    if np.x < left   { np.x = right }
    if np.x > right  { np.x = left }
    if np.y < top    { np.y = bottom }
    if np.y > bottom { np.y = top }
    if walls.contains(np) { return }
    if tiles.contains(np) {
        state.pos = np
        return
    }
    fatalError("wrong position \(state)")
}

func step2() {
    let (left, right, top, bottom) = bounds(x: state.pos.x, y: state.pos.y)
    var np = Point(x: state.pos.x + state.dx, y: state.pos.y + state.dy)
    var nd = Point(x: state.dx, y: state.dy)
    if np.x < left {
        if state.pos.x == 0 {
            if state.pos.y < 3 * size {
                np = Point(x: size, y: size - state.pos.y)
                nd = Point(x: 0, y: 1)
            } else {
                np = Point(x: 4 * size - 1, y: state.pos.y - 2 * size)
                nd = Point(x: -1, y: 0)
            }
        }
        if state.pos.x == size {
            np = Point(x: 0, y: 3 * size - state.pos.y)
            nd = Point(x: 0, y: 1)
        }
        if state.pos.x == 2 * size {
            np = Point(x: state.pos.y + size, y: size - 1)
            nd = Point(x: 0, y: 1)
        }
    }
    if np.x > right {
        if state.pos.x == size - 1 {
            np = Point(x: state.pos.y - size, y: 2 * size - 1)
            nd = Point(x: 0, y: -1)
        }
        if state.pos.x == 3 * size - 1 {
            np = Point(x: 4 * size - (state.pos.y - size) - 1, y: 2 * size)
            nd = Point(x: 0, y: 1)
        }
        if state.pos.x == 4 * size - 1 {
            np = Point(x: 0, y: state.pos.y + 2 * size)
            nd = Point(x: 1, y: 0)
        }
    }
    if np.y < top {
        if state.pos.y == size {
            if state.pos.x < size {
                np = Point(x: size - state.pos.x, y: 0)
                nd = Point(x: 0, y: 1)
            } else {
                np = Point(x: 2 * size, y: state.pos.x - size)
                nd = Point(x: 1, y: 0)
            }
        }
        if state.pos.y == 0 {
            if state.pos.x < size * 3 - 1 {
                np = Point(x: size - (state.pos.x - 2 * size), y: size - 1)
                nd = Point(x: 0, y: 1)
            } else {
                np = Point(x: 0, y: state.pos.x - 2 * size)
                nd = Point(x: 0, y: 1)
            }
        }
    }
    if np.y > bottom {
        if state.pos.y == 3 * size - 1 {
            np = Point(x: 3 * size - 1 - state.pos.x, y: 2 * size - 1)
            nd = Point(x: 0, y: -1)
        }
        if state.pos.y == 2 * size - 1 {
            if state.pos.x < 2 * size {
                np = Point(x: size - 1, y: state.pos.x + size)
                nd = Point(x: -1, y: 0)
            } else {
                np = Point(x: size - (state.pos.x - 2 * size), y: 3 * size - 1)
                nd = Point(x: 0, y: -1)
            }
        }
        if state.pos.y == size - 1 {
            np = Point(x: 3 * size - 1, y: size + (state.pos.x - 3 * size))
            nd = Point(x: -1, y: 0)
        }
    }
    if walls.contains(np) {
        return
    }
    if tiles.contains(np) {
        state.pos = np
        state.dx = nd.x
        state.dy = nd.y
        return
    }
    fatalError("wrong position \(np) state \(state)")
}

func password(commands: String, step: () -> Void) -> Int {
    let l = bounds(x: 50, y: 0).0
    state.pos = Point(x: l, y: 0)
    state.dx = 1
    state.dy = 0
    let re = try! NSRegularExpression(pattern: "L|R|[0-9]+")
    let matches = re.matches(in: commands, range: NSRange(commands.startIndex..., in: commands))
    for match in matches {
        let cmd = String(commands[Range(match.range, in: commands)!])
        switch cmd {
        case "R":
            let d = Point(x: state.dx, y: state.dy)
            switch d {
            case Point(x: 1, y: 0): state.dx, state.dy = 0, 1
            case Point(x: 0, y: 1): state.dx, state.dy = -1, 0
            case Point(x: -1, y: 0): state.dx, state.dy = 0, -1
            case Point(x: 0, y: -1): state.dx, state.dy = 1, 0
            default: fatalError("wrong direction")
            }
        case "L":
            let d = Point(x: state.dx, y: state.dy)
            switch d {
            case Point(x: 1, y: 0):
                state.dx = 0
                state.dy = -1
            case Point(x: 0, y: 1):
                state.dx = 1
                state.dy = 0
            case Point(x: -1, y: 0):
                state.dx = 0
                state.dy = 1
            case Point(x: 0, y: -1):
                state.dx = -1
                state.dy = 0
            default: fatalError("wrong direction")
            }
        default:
            let n = Int(cmd)!
            for _ in 0..<n { step() }
        }
    }
    let d = Point(x: state.dx, y: state.dy)
    var face = 0
    switch d {
    case Point(x: 1, y: 0): face = 0
    case Point(x: 0, y: 1): face = 1
    case Point(x: -1, y: 0): face = 2
    case Point(x: 0, y: -1): face = 3
    default: fatalError("wrong direction")
    }
    return (state.pos.y + 1) * 1000 + (state.pos.x + 1) * 4 + face
}

var line = readLine()!
while !line.isEmpty {
    for (x, c) in line.enumerated() {
        if c == "#" { walls.insert(Point(x: x, y: rows)) }
        if c == "." { tiles.insert(Point(x: x, y: rows)) }
    }
    rows += 1
    line = readLine()!
}
if rows == 12 { size = 4 }
line = readLine()!
print(password(commands: line, step: step1), password(commands: line, step: step2))
