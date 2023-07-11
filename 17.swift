import Foundation

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
let jets = readLine() ?? ""
let directions: [Character: Int] = ["<": -1, ">": 1]

struct Key: Hashable { var ijet, irock: Int; var depths: [Int] }
struct Pos: Hashable { var x, y: Int }

var states: [Key: [Int64]] = [:]
var maxheight = 0
var ijet = 0
var stones = Set<Pos>()

func overlap(_ rock: [String], _ pp: Pos) -> Bool {
    for x in 0..<rock.count {
        for (y, c) in rock[x].enumerated() {
            if c == "#" && stones.contains(Pos(x: pp.x - x, y: y + pp.y)) {
                return true
            }
        }
    }
    return false
}

func drop(_ rock: [String]) {
    var pp = Pos(x: maxheight + 3 + rock.count, y: 2)
    repeat {
        let jet = jets[jets.index(jets.startIndex, offsetBy: ijet % jets.count)]
        ijet += 1
        let dy = directions[jet]!
        if (dy > 0 && pp.y + rock[0].count < 7 || dy < 0 && pp.y > 0) &&
            !overlap(rock, Pos(x: pp.x, y: pp.y + dy)) {
            pp.y += dy
        }
        pp.x -= 1
    } while (pp.x > 0 && rock.count - pp.x > 0 && !overlap(rock, pp))
    print(pp)
    for x in 0..<rock.count {
        for (y, char) in rock[x].enumerated() {
            if char == "#" {
                stones.insert(Pos(x: pp.x - x, y: y + pp.y))
                maxheight = max(maxheight, pp.x)
            }
        }
    }
}

func height(_ threshold: Int64) -> Int64 {
    maxheight = -1
    states.removeAll()
    ijet = 0
    stones.removeAll()
    var r = Int64(), skipRocks = Int64(), skipHeight = Int64()
    while r + skipRocks < threshold {
        drop(shapes[Int(r % Int64(shapes.count))])
        var depths = [-100, -100, -100, -100, -100, -100, -100]
        for y in 0..<7 {
            for x in stride(from: maxheight, to: maxheight - 100, by: -1) {
                if (stones.contains(Pos(x: x, y: y))) { depths[y] = x }
            }
        }
        let k = Key(ijet: ijet % jets.count, irock: Int(r % Int64(shapes.count)), depths: depths)
        if let prev = states[k] {
            let idiff = r - prev[0]
            let remaining = threshold - r
            let remcycles = remaining / idiff
            skipRocks = remcycles * idiff
            skipHeight = remcycles * (prev[1] - Int64(maxheight))
            states.removeAll()
        }
        states[k] = [r, Int64(maxheight + 1)]
        r += 1
    }
    return Int64(maxheight) + 1 + skipHeight
}
print(height(2))
//print(height(2022), height(1000000000000))
