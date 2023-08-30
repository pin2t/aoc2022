import Foundation

struct Point: Hashable { var row, column: Int }

let shapes: [[String]] = [
    ["####"],
    [".#.",
     "###",
     ".#."],
    ["..#",
     "..#",
     "###"],
    ["#",
     "#",
     "#",
     "#"],
    ["##",
     "##"]
]
let directions: [Character: Int] = ["<": -1, ">": 1]
let jets = (readLine() ?? "").map{ directions[$0]! }

var nrows = 0
var ijet = 0
var stones = Set<Point>()

func overlap(_ rock: [String], _ p: Point) -> Bool {
    for row in 0..<rock.count {
        for (column, ch) in rock[row].enumerated() {
            if ch == "#" && stones.contains(Point(row: p.row - row, column: column + p.column)) {
                return true
            }
        }
    }
    return false
}

func drop(_ rock: [String]) {
    var p = Point(row: nrows + 2 + rock.count, column: 2)
    repeat {
        let jet = jets[ijet % jets.count]
        ijet += 1
        if (jet > 0 && p.column + rock[0].count < 7 ||
            jet < 0 && p.column > 0) &&
           !overlap(rock, Point(row: p.row, column: p.column + jet)) {
            p.column += jet
        }
        p.row -= 1
    } while p.row >= 0 && p.row + 1 - rock.count >= 0 && !overlap(rock, p)
    if p.row < 0 || p.row + 1 - rock.count < 0 || overlap(rock, p)  {
        p.row += 1
    }
    for row in 0..<rock.count {
        for (column, char) in rock[row].enumerated() {
            if char == "#" {
                stones.insert(Point(row: p.row - row, column: column + p.column))
            }
        }
        nrows = max(nrows, p.row - row + 1)
    }
}

func height(_ threshold: Int64) -> Int64 {
    struct Key: Hashable { var ijet, irock: Int; var depths: [Int] }
    var states: [Key: [Int64]] = [:]
    nrows = 0
    ijet = 0
    stones.removeAll()
    var r = Int64(), skipRocks = Int64(), skipHeight = Int64()
    while r + skipRocks < threshold {
        drop(shapes[Int(r % Int64(shapes.count))])
        var depths = [-100, -100, -100, -100, -100, -100, -100]
        for column in 0..<7 {
            for row in stride(from: nrows, to: nrows - 100, by: -1) {
                if (stones.contains(Point(row: row, column: column))) {
                    depths[column] = max(depths[column], nrows - row)
                }
            }
        }
        let key = Key(ijet: ijet % jets.count, irock: Int(r % Int64(shapes.count)), depths: depths)
        if let prev = states[key] {
            let length = r - prev[0]
            let cycles = (threshold - r) / length
            skipRocks = cycles * length
            skipHeight = cycles * (Int64(nrows) - prev[1])
            states.removeAll()
        } else {
            states[key] = [r, Int64(nrows)]
        }
        r += 1
    }
    return Int64(nrows) + skipHeight
}
print(height(2022), height(1000000000000))
