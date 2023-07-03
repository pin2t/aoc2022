import Foundation

struct Blueprint {
    var n: Int
    var ore, clay: Int
    var obsidian: (ore: Int, clay: Int)
    var geode: (ore: Int, obsidian: Int)
}

func sum(_ items: [Int], _ a: [Int]) -> [Int] {
    return [items[0] + a[0], items[1] + a[1], items[2] + a[2], items[3] + a[3]]
}

struct State: Hashable {
    var minute: Int
    var values: [Int] // ore, clay, obsidian, geode
    var robots: [Int] // ore, clay, obsidian, geode

    func val() -> Int {
        return robots[0] + robots[1] * 10 + robots[2] * 100 + robots[3] * 1000
    }
}

func maxgeodes(_ bp: Blueprint, _ minutes: Int, _ maxlen: Int) -> Int {
    var m = 0, maxmin = 0
    var processed = Set<State>()
    var queue = [State(minute: 0, values: [0, 0, 0, 0], robots: [1, 0, 0, 0])]

    while !queue.isEmpty {
        if queue.first!.minute > maxmin {
            maxmin = queue.first!.minute
            processed = Set<State>()
            if queue.count > maxlen {
                queue.sort { $0.val() > $1.val() }
                queue.removeLast(queue.count - maxlen)
                continue
            }
        }
        let s = queue.removeFirst()
        if s.minute == minutes {
            m = max(m, s.values[3])
            continue
        }
        if processed.contains(s) {
            continue
        }
        processed.insert(s)
        if s.values[0] >= bp.geode.ore && s.values[2] >= bp.geode.obsidian {
            queue.append(State(minute: s.minute + 1,
                               values: sum(s.values, [s.robots[0] - bp.geode.ore, s.robots[1], s.robots[2] - bp.geode.obsidian, s.robots[3]]),
                               robots: sum(s.robots, [0, 0, 0, 1])))
        }
        if s.values[0] >= bp.obsidian.ore && s.values[1] >= bp.obsidian.clay {
            queue.append(State(minute: s.minute + 1,
                               values: sum(s.values, [s.robots[0] - bp.obsidian.ore, s.robots[1] - bp.obsidian.clay, s.robots[2], s.robots[3]]),
                               robots: sum(s.robots, [0, 0, 1, 0])))
        }
        if s.values[0] >= bp.clay {
            queue.append(State(minute: s.minute + 1,
                               values: sum(s.values, [s.robots[0] - bp.clay, s.robots[1], s.robots[2], s.robots[3]]),
                               robots: sum(s.robots, [0, 1, 0, 0])))
        }
        if s.values[0] >= bp.ore {
            queue.append(State(minute: s.minute + 1,
                               values: sum(s.values, [s.robots[0] - bp.ore, s.robots[1], s.robots[2], s.robots[3]]),
                               robots: sum(s.robots, [1, 0, 0, 0])))
        }
        queue.append(State(minute: s.minute + 1, values: sum(s.values, s.robots), robots: s.robots))
    }
    return m
}

var blueprints = [Blueprint]()
while let line = readLine() {
    let n = line.numbers
    blueprints.append(Blueprint(n: n[0], ore: n[1], clay: n[2], obsidian: (n[3], n[4]), geode: (n[5], n[6])))
}
var n1 = 0
for bp in blueprints {
    n1 += bp.n * maxgeodes(bp, 24, 2000000)
}
var prod = 1
for bp in blueprints.prefix(3) {
    prod *= maxgeodes(bp, 32, 2000000)
}
print(n1, prod)

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
