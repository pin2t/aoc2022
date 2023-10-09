import Foundation

enum FieldItem {
    case tile, wall
}

var field: [Point: FieldItem] = [:]

struct Point: Hashable {
    var x, y: Int

    func bounds() -> (Int, Int, Int, Int) {
        var left = 10000, right = 0, top = 10000, bottom = 0
        for (point, _) in field {
            if point.y == y || point.x == x {
                (left, right) = (min(left, point.x), max(right, point.x))
                (top, bottom) = (min(top, point.y), max(bottom, point.y))
            }
        }
        return (left, right, top, bottom)
    }

    func move(_ direction: Direction) -> Point {
        return Point(x: x + direction.dx, y: y + direction.dy)
    }
}

struct Direction: Equatable {
    static let up    = Direction(dx: 0,  dy: -1)
    static let down  = Direction(dx: 0,  dy: 1)
    static let left  = Direction(dx: -1, dy: 0)
    static let right = Direction(dx: 1,  dy: 0)

    var dx, dy: Int

    func left() -> Direction {
        switch self {
        case .up:    return .left
        case .down:  return .right
        case .left:  return .down
        case .right: return .up
        default: fatalError("wrong direction")
        }
    }

    func right() -> Direction {
        switch self {
        case .up:    return .right
        case .down:  return .left
        case .left:  return .up
        case .right: return .down
        default: fatalError("wrong direction")
        }
    }
}


class State {
    var pos: Point
    var direction: Direction

    init() {
        let left = Point(x: 50, y: 0).bounds().0
        pos = Point(x: left, y: 0)
        direction = Direction.right
    }

    func step() {
        var next = pos.move(direction)
        let (left, right, top, bottom) = pos.bounds()
        if next.x < left   { next.x = right }
        if next.x > right  { next.x = left }
        if next.y < top    { next.y = bottom }
        if next.y > bottom { next.y = top }
        if let item = field[next] {
            if item == .tile { pos = next }
            return
        }
        fatalError("wrong position \(pos) direction \(direction)")
    }

    func left() {
        direction = direction.left()
    }

    func right() {
        direction = direction.right()
    }

    func password() -> Int {
        var face = 0
        switch direction {
        case .right: face = 0
        case .down:  face = 1
        case .left:  face = 2
        case .up:    face = 3
        default: fatalError("wrong direction")
        }
        return (pos.y + 1) * 1000 + (pos.x + 1) * 4 + face
    }
}

class CubicState: State {
    override init() {
        super.init()
    }

    override func step() {
        var next = pos.move(direction)
        var ndir = Direction.up
        switch direction {
        case .down:
            if pos.y == 49 {
                if next.x >= 100 && next.x < 150 {
                    (next, ndir) = (Point(x: 99, y: pos.x - 50), .left)
                }
            } else if pos.y == 149 {
                if next.x >= 50 && next.x < 100 {
                    (next, ndir) = (Point(x: 49, y: next.x + 100), .left)
                }
            } else if next.y >= 200 && next.x >= 0 && next.x < 50 {
                (next, ndir) = (Point(x: next.x + 100, y: 0), .down)
            }
        case .up:
            if pos.y == 0 {
                if next.x >= 50 && next.x < 100 {
                    (next, ndir) = (Point(x: 0, y: next.x + 100), .right)
                } else if next.x >= 100 && next.x < 150 {
                    (next, ndir) = (Point(x: next.x - 100, y: 199), .up)
                }
            } else if pos.y == 100 {
                if next.x >= 0 && next.x < 50 {
                    (next, ndir) = (Point(x: 50, y: next.x + 50), .right)
                }
            }
        case .right:
            if next.x >= 150 && next.y >= 0 && next.y < 50 {
                (next, ndir) = (Point(x: 99, y: 149 - next.y), .left)
            } else if pos.x == 49 {
                if next.y >= 150 && next.y < 200 {
                    (next, ndir) = (Point(x: next.y - 100, y: 149), .up)
                }
            } else if pos.x == 99 {
                if next.y >= 50 && next.y < 100 {
                    (next, ndir) = (Point(x: next.y + 50, y: 49), .up)
                } else if next.y >= 100 && next.y < 150 {
                    (next, ndir) = (Point(x: 149, y: 149 - next.y), .left)
                }
            }
        case .left:
            if next.x < 0 && next.y >= 150 && next.y < 200 {
                (next, ndir) = (Point(x: next.y - 100, y: 0), .down)
            } else if next.x == 49 && next.y >= 50 && next.y < 100 {
                (next, ndir) = (Point(x: next.y - 50, y: 100), .down)
            } else if next.x == 49 && next.y >= 0 && next.y < 50 {
                (next, ndir) = (Point(x: 0, y: 149 - next.y), .right)
            } else if next.x < 0 && next.y >= 100 && next.y < 150 {
                (next, ndir) = (Point(x: 50, y: 149 - next.y), .right)
            }
        default: fatalError("Unknown dorection \(direction)")
        }
        if let item = field[next] {
            if item == .tile {
                pos = next
                direction = ndir
            }
            return
        }
        fatalError("wrong next position \(next), current position \(pos) direction \(direction)")
    }
}

func password(state: State, commands: String) -> Int {
    let re = try! NSRegularExpression(pattern: "L|R|[0-9]+")
    let matches = re.matches(in: commands, range: NSRange(commands.startIndex..., in: commands))
    for match in matches {
        let cmd = String(commands[Range(match.range, in: commands)!])
        switch cmd {
        case "R": state.right()
        case "L": state.left()
        default:
            let n = Int(cmd)!
            for _ in 0..<n { state.step() }
        }
    }
    return state.password()
}

var line = readLine()!
var row = 0
while !line.isEmpty {
    for (col, ch) in line.enumerated() {
        if ch == "#" { field[Point(x: col, y: row)] = .wall }
        if ch == "." { field[Point(x: col, y: row)] = .tile }
    }
    row += 1
    line = readLine()!
}
line = readLine()!
print(password(state: State(), commands: line), password(state: CubicState(), commands: line))

