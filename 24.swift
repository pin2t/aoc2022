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
    var minutes: Int
    var stage: Int // 0 - down, 1 - up, 2 - down again

    func next(to: Pos) -> State {
        return State(pos: to, minutes: minutes + 1, stage: stage)
    }
}

var grid = [Pos: Character]()
var blizzards = [Pos: Pos]()
var bottomRight = Pos(x: 0, y: 0)
var x = 0

while let line = readLine() {
    let chars = Array(line.utf8)
    for (y, c) in chars.enumerated() {
        let position = Pos(x: x, y: y)
        grid[position] = Character(UnicodeScalar(c))
        bottomRight = Pos(x: max(position.x, bottomRight.x),
            y: max(position.y, bottomRight.y))
        switch c {
        case UInt8(ascii: "<"): blizzards[position] = Pos(x: -1, y: 0)
        case UInt8(ascii: ">"): blizzards[position] = Pos(x: 1, y: 0)
        case UInt8(ascii: "^"): blizzards[position] = Pos(x: 0, y: -1)
        case UInt8(ascii: "v"): blizzards[position] = Pos(x: 0, y: 1)
        default: break
        }
    }
    x += 1
}

let deltas = [Pos(x: 0, y: 1), Pos(x: 0, y: -1), Pos(x: 1, y: 0), Pos(x: -1, y: 0)]
var queue = [State(pos: Pos(x: 0, y: 1), minutes: 0, stage: 0)]
var processed = Set<State>()
var n1 = 0

while !queue.isEmpty {
    let st = queue.removeFirst()
    if processed.contains(st) {
        continue
    }
    processed.insert(st)
    var stopped = false
    for (sp, d) in blizzards {
        if sp.x == st.pos.x {
            let sy = mod(sp.y + (d.y * st.minutes) - 1, bottomRight.y - 1) + 1
            if sy == st.pos.y { stopped = true }
        }
        if sp.y == st.pos.y {
            let sx = mod(sp.x + (d.x * st.minutes) - 1, bottomRight.x - 1) + 1
            if sx == st.pos.x { stopped = true }
        }
    }
    if stopped {
        continue
    }
    if st.pos.x == 0 && st.stage > 1 {
        var nextState = st
        nextState.stage += 1
        queue.append(nextState)
    }
    if st.pos.x == bottomRight.x {
        if st.stage == 0 {
            n1 = st.minutes
            var nextState = st
            nextState.stage += 1
            queue.append(nextState)
        }
        if st.stage > 1 {
            print(n1, st.minutes)
            break
        }
    }
    for d in deltas {
        let to = st.pos.move(delta: Pos(x: d.x, y: d.y))
        if grid[to] == "#" {
            continue
        }
        if grid[to] == nil {
            continue
        }
        var nextState = st
        nextState.pos = to
        queue.append(nextState.next(to: to))
    }
    var nextState = st
    nextState.pos = st.pos
    queue.append(nextState.next(to: st.pos))
}
