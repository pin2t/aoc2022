import Foundation

struct Pos: Hashable {
    var x, y: Int

    func move(delta: Pos) -> Pos {
        return Pos(x: x + delta.x, y: y + delta.y)
    }
}

func mod(_ x: Int, _ m: Int) -> Int {
    return (x % m + m) % m
}

struct State: Hashable {
    var pos: Pos
    var time: Int
}

var grid = [Pos: Character]()
var blizzards = [Pos: Pos]()
var bottomRight = Pos(x: 0, y: 0)
var y = 0

while let line = readLine() {
    let chars = Array(line.utf8)
    for (x, c) in chars.enumerated() {
        let position = Pos(x: x, y: y)
        grid[position] = Character(UnicodeScalar(c))
        bottomRight = Pos(x: max(position.x, bottomRight.x), y: max(position.y, bottomRight.y))
        switch c {
        case UInt8(ascii: "<"): blizzards[position] = Pos(x: -1, y: 0)
        case UInt8(ascii: ">"): blizzards[position] = Pos(x: 1, y: 0)
        case UInt8(ascii: "^"): blizzards[position] = Pos(x: 0, y: -1)
        case UInt8(ascii: "v"): blizzards[position] = Pos(x: 0, y: 1)
        default: break
        }
    }
    y += 1
}

let deltas = [Pos(x: 0, y: 1), Pos(x: 0, y: -1), Pos(x: 1, y: 0), Pos(x: -1, y: 0)]
var queue = [State(pos: Pos(x: 0, y: 1), time: 0)]
var processed = Set<State>()

func minutes(to: Pos) -> Int {
    while true {
        let st = queue.removeFirst()
        if st.pos == to {
            return st.time
        }
        if processed.contains(st) {
            continue
        }
        processed.insert(st)
        var stopped = false
        for (sp, d) in blizzards {
            if sp.x == st.pos.x && d.y != 0 {
                let sy = mod(sp.y + (d.y * st.time) - 1, bottomRight.y - 1) + 1
                if sy == st.pos.y {
                    stopped = true
                    break
                }
            }
            if sp.y == st.pos.y && d.x != 0 {
                let sx = mod(sp.x + (d.x * st.time) - 1, bottomRight.x - 1) + 1
                if sx == st.pos.x {
                    stopped = true
                    break
                }
            }
        }
        if stopped {
            continue
        }
        for d in deltas {
            let to = st.pos.move(delta: Pos(x: d.x, y: d.y))
            if grid[to] != nil && grid[to] != "#" {
                queue.append(State(pos: to, time: st.time + 1))
            }
        }
        queue.append(State(pos: st.pos, time: st.time + 1))
    }
}

let n1 = minutes(to: Pos(x: bottomRight.x - 1, y: bottomRight.y))
print(n1, n1 + minutes(to: Pos(x: 1, y: 0)) + minutes(to: Pos(x: bottomRight.x - 1, y: bottomRight.y)))