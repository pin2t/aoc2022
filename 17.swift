import Foundation

let shapes: [[String]] = [
    ["####"],
    [".#.", "###", ".#."],
    ["..#", "..#", "###"],
    ["#", "#", "#", "#"],
    ["##", "##"]
]
let jets = readLine() ?? ""
let dx: [Character: Int] = ["<": -1, ">": 1]
var tops: [String: [Int64]] = [:]
var maxheight: Int = 0
var njet: Int = 0
var stones = Set<Pos>()

struct Pos: Hashable { var x, y: Int }

func overlap(_ rock: [String], _ pp: Pos) -> Bool {
    for x in 0..<rock.count {
        for (y, c) in rock[x].enumerated() {
            if c == "#" && stones.contains(Pos(x: pp.x - x - 1, y: y + pp.y)) {
                return true
            }
        }
    }
    return false
}

func drop(_ rock: [String]) {
    var pp = Pos(x: maxheight + 3 + rock[0].count, y: 2)
    repeat {
        let jet = jets[jets.index(jets.startIndex, offsetBy: njet % jets.count)]
        njet += 1
        let dy = dx[jet]!
        if (dy > 0 && pp.y + rock[0].count < 7 || dy < 0 && pp.y > 0) &&
            !overlap(rock, Pos(x: pp.x - 1, y: pp.y + dy)) {
            let jet = jets[jets.index(jets.startIndex, offsetBy: (njet - 1) % jets.count)]
            pp.y += dx[jet]!
        }
        pp.x -= 1
    } while (pp.x > 0) && (rock.count - pp.x - 1 > 0) && !overlap(rock, pp)
    for x in 0..<rock.count {
        for (y, char) in rock[x].enumerated() {
            if char == "#" {
                stones.insert(Pos(x: pp.x - x, y: y + pp.y))
                maxheight = max(maxheight, pp.x - x)
            }
        }
    }
}

func height(_ threshold: Int64) -> Int64 {
    maxheight = -1
    for r in 0..<threshold {
        drop(shapes[Int(r % Int64(shapes.count))])
        if maxheight > 100 {
            var key = ""
            for x in stride(from: maxheight, to: maxheight - 100, by: -1) {
                for y in 0..<7 {
                    key.append(stones.contains(Pos(x: x, y: y)) ? "#" : ".")
                }
            }
            if let prev = tops[key] {
                print("cycle found at", r, "length", r - prev[0])
                let idiff = r - prev[0]
                let remaining = threshold - r
                if remaining % idiff == 0 {
                    let height = Int64(maxheight + 1)
                    let hdiff = height - prev[1]
                    let remcycles = (remaining / idiff) + 1
                    return prev[1] + remcycles * hdiff
                }
            }
            tops[key] = [r, Int64(maxheight + 1)]
        }
    }
    return Int64(maxheight + 1)
}

print(height(2022), height(1000000000000))
