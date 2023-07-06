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
var seen: [String: [Int64]] = [:]
var maxheight: Int = 0
var njet: Int = 0
var stones: [[Int]: Bool] = [:]

func overlap(_ rock: [String], _ pp: [Int]) -> Bool {
    for x in 0..<rock.count {
        for (y, c) in rock[x].enumerated() {
            if c == "#" && stones[[pp[0] - x - 1, y + pp[1]]] != nil {
                return true
            }
        }
    }
    return false
}

func drop(_ rock: [String]) {
    var pp = [maxheight + 3 + rock[0].count, 2]
    repeat {
        let ch = jets[jets.index(jets.startIndex, offsetBy: njet % jets.count)]
        njet += 1
        let dy = dx[ch]!
        if (dy > 0 && pp[1] + rock[0].count < 7 || dy < 0 && pp[1] > 0) &&
            !overlap(rock, [pp[0] - 1, pp[1] + dy]) {
            let ch = jets[jets.index(jets.startIndex, offsetBy: (njet - 1) % jets.count)]
            pp[1] += dx[ch]!
        }
        pp[0] -= 1
    } while (pp[0] > 0) && (rock.count - pp[0] - 1 > 0) && !overlap(rock, pp)
    for x in 0..<rock.count {
        for (y, char) in rock[x].enumerated() {
            if char == "#" {
                stones[[pp[0] - x, y + pp[1]]] = true
                maxheight = max(maxheight, pp[0] - x)
            }
        }
    }
}

func height(_ threshold: Int64) -> Int64 {
    maxheight = -1
    for r in 0..<threshold {
        drop(shapes[Int(r % Int64(shapes.count))])
        if r > 3000 {
            var key = ""
            for x in stride(from: maxheight, to: maxheight - 100, by: -1) {
                for y in 0..<7 {
                    key.append(stones[[x, y]] != nil ? "#" : ".")
                }
            }
            if let prev = seen[key] {
                let idiff = r - prev[0]
                let remaining = threshold - r
                if remaining % idiff == 0 {
                    let height = Int64(maxheight + 1)
                    let hdiff = height - prev[1]
                    let remcycles = (remaining / idiff) + 1
                    return prev[1] + remcycles * hdiff
                }
            }
            seen[key] = [r, Int64(maxheight + 1)]
        }
    }
    return Int64(maxheight + 1)
}

print(height(2022), height(1000000000000))
