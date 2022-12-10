struct Pos: Hashable {
    var row, col: Int

    init(_ r: Int, _ c: Int) {
        row = r
        col = c
    }
}

var grid: [Pos:Int] = [:]
var row = 0, cols = 0
while let line = readLine() {
    for (i, c) in line.enumerated() { grid[Pos(row, i)] = Int(String(c))! }
    row += 1
    cols  = line.count
}
var highest = 0, n = row * 4 - 4
var p = Pos(1, 1)
while p.row < row - 1 {
    p.col = 1
    while p.col < cols - 1  {
        var pp  = Pos(p.row, p.col - 1), visible = false, score = 1
        while pp.col > 0 && grid[pp]! < grid[p]! { pp.col -= 1 }
        visible = visible || pp.col == 0 && grid[pp]! < grid[p]!
        score *= (p.col - pp.col)
        pp  = Pos(p.row - 1, p.col)
        while pp.row > 0 && grid[pp]! < grid[p]! { pp.row -= 1 }
        visible = visible || pp.row == 0 && grid[pp]! < grid[p]!
        score *= (p.row - pp.row)
        pp  = Pos(p.row, p.col + 1)
        while pp.col < cols - 1 && grid[pp]! < grid[p]! { pp.col += 1 }
        visible = visible || pp.col == cols - 1 && grid[pp]! < grid[p]!
        score *= (pp.col - p.col)
        pp  = Pos(p.row + 1, p.col)
        while pp.row < row - 1 && grid[pp]! < grid[p]! { pp.row += 1 }
        visible = visible || pp.row == row - 1 && grid[pp]! < grid[p]!
        score *= (pp.row - p.row)
        if score > highest { highest = score }
        if visible { n += 1 }
        p.col += 1
    }
    p.row += 1
}
print(n, highest)
