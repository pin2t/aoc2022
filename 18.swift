import Foundation

struct Pos: Hashable {
    var x, y, z: Int
}

func max(_ a: Pos, _ b: Pos) -> Pos {
    return Pos(x: max(a.x, b.x), y: max(a.y, b.y), z: max(a.z, b.z))
}

func min(_ a: Pos, _ b: Pos) -> Pos {
    return Pos(x: min(a.x, b.x), y: min(a.y, b.y), z: min(a.z, b.z))
}

var cubes = Set<Pos>()
while let line = readLine() {
    let n = line.numbers
    cubes.insert(Pos(x: n[0], y: n[1], z: n[2]))
}
var surface = 0
var bounds = [Pos(x: 0, y: 0, z: 0), Pos(x: 0, y: 0, z: 0)]
for cube in cubes {
    if !cubes.contains(Pos(x: cube.x + 1, y: cube.y, z: cube.z)) { surface += 1 }
    if !cubes.contains(Pos(x: cube.x - 1, y: cube.y, z: cube.z)) { surface += 1 }
    if !cubes.contains(Pos(x: cube.x, y: cube.y - 1, z: cube.z)) { surface += 1 }
    if !cubes.contains(Pos(x: cube.x, y: cube.y + 1, z: cube.z)) { surface += 1 }
    if !cubes.contains(Pos(x: cube.x, y: cube.y, z: cube.z - 1)) { surface += 1 }
    if !cubes.contains(Pos(x: cube.x, y: cube.y, z: cube.z + 1)) { surface += 1 }
    bounds[0] = min(bounds[0], cube)
    bounds[1] = max(bounds[1], cube)
}
bounds[0].x -= 1; bounds[0].y -= 1; bounds[0].z -= 1
bounds[1].x += 1; bounds[1].y += 1; bounds[1].z += 1

var external = Set<Pos>()
var queue = [bounds[0]]
while !queue.isEmpty {
    let c = queue.removeFirst()
    if external.contains(c) || cubes.contains(c) ||
       c.x < bounds[0].x || c.y < bounds[0].y || c.z < bounds[0].z ||
       c.x > bounds[1].x || c.y > bounds[1].y || c.z > bounds[1].z {
        continue
    }
    external.insert(c)
    queue.append(Pos(x: c.x - 1, y: c.y, z: c.z))
    queue.append(Pos(x: c.x + 1, y: c.y, z: c.z))
    queue.append(Pos(x: c.x, y: c.y - 1, z: c.z))
    queue.append(Pos(x: c.x, y: c.y + 1, z: c.z))
    queue.append(Pos(x: c.x, y: c.y, z: c.z - 1))
    queue.append(Pos(x: c.x, y: c.y, z: c.z + 1))
}
var esurface = 0
for c in cubes {
    if external.contains(Pos(x: c.x + 1, y: c.y, z: c.z)) { esurface += 1 }
    if external.contains(Pos(x: c.x - 1, y: c.y, z: c.z)) { esurface += 1 }
    if external.contains(Pos(x: c.x, y: c.y + 1, z: c.z)) { esurface += 1 }
    if external.contains(Pos(x: c.x, y: c.y - 1, z: c.z)) { esurface += 1 }
    if external.contains(Pos(x: c.x, y: c.y, z: c.z + 1)) { esurface += 1 }
    if external.contains(Pos(x: c.x, y: c.y, z: c.z - 1)) { esurface += 1 }
}
print(surface, esurface)

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
