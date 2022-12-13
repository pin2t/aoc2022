struct Pos: Hashable {var x = 0, y = 0}

func height(_ item: Character) -> Int {
    let chars = "abcdefghijklmnopqrstuvwxyz"
    return chars.firstIndex(of: item == "S" ? "a" : (item == "E" ? "z" : item))!.utf16Offset(in: chars) + 1 
}

func path(graph: [Pos:[Pos:Int]], start: Pos, end: Pos) -> Int {
    
}

var graph: [Pos:[Pos:Int]] = [:]
var grid: [String] = []
var start = Pos(), end = Pos()
var starts: [Pos] = []
while let line = readLine() { grid.append(line) }
for y in 0..<grid.count {
    for x in 0..<grid[y].count {
        var h = grid[y][String.Index(utf16Offset: x, in: grid[y])]
        switch h {
        case "S": start = Pos(x: x, y: y)
        case "E": end = Pos(x: x, y: y)
        case "a": starts.append(Pos(x: x, y: y))
        default: break
        }
        if !graph.keys.contains(Pos(x: x, y: y)) {
            var vertex: [Pos:Int] = [:]
            let edge = { (x: Int, y: Int) -> Void in
                if x >= 0 && x < grid[0].count && 
                   y >= 0 && y < grid.count {
                    var c = grid[y][String.Index(utf16Offset: x, in: grid[y])]
                    if height(c) - height(h) < 2 {
                        vertex[Pos(x: x, y: y)] = 1
                    }
                } 
            }
            edge(x + 1, y)
            edge(x - 1, y)
            edge(x, y + 1)
            edge(x, y - 1)
            graph[Pos(x: x, y: y)] = vertex
        }
    }
}
print(graph)