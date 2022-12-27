var grid: [String] = []
var start = Pos(0, 0), end = Pos(0, 0)
var starts: [Pos] = []
while let line = readLine() { grid.append(line) }
for y in 0..<grid.count {
    for x in 0..<grid[y].count {
        let c = grid[y][String.Index(utf16Offset: x, in: grid[y])]
        switch c {
        case "S": start = Pos(x, y)
        case "E": end = Pos(x, y)
        case "a": starts.append(Pos(x, y))
        default: break
        }
    }
}
var m = 1000000000000
for s in starts { 
    let distance = path(s)
    if distance > 0 { m = min(m, distance) }
}
print(path(start), m)

struct Pos: Hashable {
    var x = 0, y = 0
    init(_ x: Int, _ y: Int) { 
        self.x = x
        self.y = y
    }
}

func height(_ item: Character) -> Int {
    let chars = "abcdefghijklmnopqrstuvwxyz"
    return chars.firstIndex(of: item == "S" ? "a" : (item == "E" ? "z" : item))!.utf16Offset(in: chars) + 1 
}

struct Pair { 
    var pos: Pos
    var distance: Int
}

func path(_ start: Pos) -> Int {
    var queue: [Pair] = []
    var visited: Set<Pos> = []
    queue.append(Pair(pos: start, distance: 0))
    while !queue.isEmpty {
        let p = queue.removeFirst()
        for step in [Pos(p.pos.x + 1, p.pos.y), Pos(p.pos.x - 1, p.pos.y), 
            Pos(p.pos.x, p.pos.y + 1), Pos(p.pos.x, p.pos.y - 1)] {
            if step.x < 0 || step.x >= grid[0].count { continue }
            if step.y < 0 || step.y >= grid.count { continue }
            if visited.contains(step) { continue }
            let from = grid[p.pos.y][String.Index(utf16Offset: p.pos.x, in: grid[p.pos.y])]
            let to = grid[step.y][String.Index(utf16Offset: step.x, in: grid[step.y])]
            if height(to) - height(from) > 1 { continue }
            if step == end {
                return p.distance + 1
            }
            visited.insert(step)
            queue.append(Pair(pos: step, distance: p.distance + 1))
        }
    }
    return 0
}
