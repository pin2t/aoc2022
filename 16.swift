import Foundation

var valves = [String: Int]()
var tunnels = [String: [String]]()
while let line = readLine() {
    let v = line.allMatches("[A-Z]{2}")
    valves[v[0]] = Int(line.allMatches("[0-9]+")[0])!
    tunnels[v[0]] = Array(v.dropFirst())
}
var distances = [String: [String: Int]]()
var indices = [String: Int]()
var i = 0
for valve in valves {
    if valve.key != "AA" && valve.value == 0 {
        continue
    }
    if valve.key != "AA" {
        indices[valve.key] = i
        i += 1
        distances[valve.key] = [valve.key: 0, "AA": 0]
    } else {
        distances[valve.key] = ["AA": 0]
    }
    var visited = Set<String>()
    visited.insert(valve.key)
    var queue = [(0, valve.key)]
    while !queue.isEmpty {
        let (distance, position) = queue.removeFirst()
        for neighbor in tunnels[position]! {
            if visited.contains(neighbor) { continue }
            visited.insert(neighbor)
            if valves[neighbor] != nil {
                distances[valve.key]![neighbor] = distance + 1
            }
            queue.append((distance + 1, neighbor))
        }
    }
    distances[valve.key]!.removeValue(forKey: valve.key)
    if valve.key != "AA" {
        distances[valve.key]!.removeValue(forKey: "AA")
    }
}
var cache = [String: Int]()
func pressure(_ time: Int, _ valve: String, _ openmask: Int) -> Int {
    let key = "\(time),\(valve),\(openmask)"
    if let p = cache[key] { return p }
    var m = 0
    for (neighbor, dist) in distances[valve]! {
        guard let i = indices[neighbor] else { continue }
        let bit = 1 << i
        if openmask & bit != 0 { continue }
        let remaining = time - dist - 1
        if remaining <= 0 { continue }
        m = max(m, pressure(remaining, neighbor, openmask | bit) + valves[neighbor]! * remaining)
    }
    cache[key] = m
    return m
}
let n1 = pressure(30, "AA", 0)
var n2 = 0
for openmask in 0...(1 << indices.count - 1) {
    n2 = max(n2, pressure(26, "AA", openmask) + pressure(26, "AA", ~openmask))
}
print(n1, n2)
extension String {
    func allMatches(_ pattern: String) -> [String] {
        guard let regex = try? NSRegularExpression(pattern: pattern) else { return [] }
        let matches  = regex.matches(in: self, options: [], range: NSMakeRange(0, self.count))
        return matches.map { match in
            String(self[Range(match.range, in: self)!])
        }
    }
}

