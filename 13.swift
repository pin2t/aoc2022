import Foundation

var n = 0, i = 1, prev2 = 1, prev6 = 2
while let line = readLine() {
    let packet1 = try JSONSerialization.jsonObject(with: Data(line.utf8)) as? [Any]
    let packet2 = try JSONSerialization.jsonObject(with: Data(readLine()!.utf8)) as? [Any]
    _ = readLine()
    if compare(packet1 as Any, packet2 as Any) < 0 { n += i }
    i += 1
    if compare(packet1 as Any, [[2]]) < 0 { prev2 += 1 }
    if compare(packet2 as Any, [[2]]) < 0 { prev2 += 1 }
    if compare(packet1 as Any, [[6]]) < 0 { prev6 += 1 }
    if compare(packet2 as Any, [[6]]) < 0 { prev6 += 1 }
}
print(n, prev2 * prev6)

func compare(_ p1: Any, _ p2: Any) -> Int {
    if let n1 = p1 as? Int {
        if let n2 = p2 as? Int {
            return n1 - n2
        } else {
            return compare([n1], p2)
        }
    } else {
        if let n2 = p2 as? Int {
            return compare(p1, [n2])
        }
    }
    let a1 = p1 as? [Any]
    let a2 = p2 as? [Any]
    for i in 0..<min(a1!.count, a2!.count) {
        if compare(a1![i], a2![i]) != 0 { 
            return compare(a1![i], a2![i]) 
        }
    }
    return a1!.count - a2!.count
}