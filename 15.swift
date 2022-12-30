import Foundation

var sensors: [[Int]] = []
while let line = readLine() { sensors.append(line.numbers) }
var target2m: [Interval]
var beacons: Int
(target2m, beacons) = parse(2000000)
target2m.sort {$1.start > $0.start}
let merged2m = merge(target2m)
for y in 0...40000000 {
    var targety: [Interval]
    (targety, _) = parse(y)
    targety.sort {$1.start > $0.start}
    let merged = merge(targety)
    if merged.count > 1 {
        print(merged2m.first!.end - merged2m.first!.start + 1 - beacons,
              (merged.first!.end + 1) * 4000000 + y)
        break      
    }
}

func parse(_ targety: Int) -> ([Interval], Int) {
    var result: [Interval] = []
    var beacons: Set<Int> = []
    for s in sensors {
        let sx = s[0], sy = s[1]
        let bx = s[2], by = s[3]
        let distance = abs(bx - sx) + abs(by - sy)
        let offset = distance - abs(sy - targety)
        if offset >= 0 {
            result.append(Interval(sx - offset, sx + offset))
            if by == targety {
                beacons.insert(bx)
            }
        }
    }
    return (result, beacons.count)
}

func merge(_ intervals: [Interval]) -> [Interval] {
    var merged: [Interval] = []
    for interval in intervals {
        if merged.isEmpty { 
            merged.append(interval)
            continue
        }
        let i = merged.first! 
        if interval.start > i.end + 1 {
            merged.append(Interval(interval.start, interval.end))
        } else {
            merged[0] = Interval(i.start, max(i.end, interval.end))
        }
    }
    return merged
}

struct Interval {
    var start, end: Int
    init(_ s: Int, _ e: Int) { self.start = s; self.end = e }
}

extension String {
    var numbers: [Int] {
        guard let regex = try? NSRegularExpression(pattern: "-?\\d+", options: [.caseInsensitive]) else { 
            return [] 
        }
        let matches  = regex.matches(in: self, options: [], range: NSMakeRange(0, self.count))
        return matches.map { match in
            String(self[Range(match.range, in: self)!])
        }.map { 
            Int($0)! 
        }
    }
}
