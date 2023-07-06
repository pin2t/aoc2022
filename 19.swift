import Foundation

struct Blueprint {
    var n: Int
    var ore, clay: Int
    var obsidian: (ore: Int, clay: Int)
    var geode: (ore: Int, obsidian: Int)
}

struct Items: Hashable {
    var ore, clay, obsidian, geode: Int

    func add(_ ore: Int, _ clay: Int, _ obsidian: Int, _ geode: Int) -> Items {
        return Items(ore: self.ore + ore, clay: self.clay + clay,
            obsidian: self.obsidian + obsidian, geode: self.geode + geode)
    }
}

struct State: Hashable {
    var minute: Int
    var values: Items
    var robots: Items

    init(_ m: Int, _ v: Items, _ r: Items) {
        self.minute = m
        self.values = v
        self.robots = r
    }

    func val() -> Int {
        return robots.ore + robots.clay * 10 + robots.obsidian * 100 + robots.geode * 1000
    }
}

func maxgeodes(_ bp: Blueprint, _ minutes: Int, _ maxlen: Int) -> Int {
    var m = 0, maxmin = 0
    var processed = Set<State>()
    var queue = [State(0,
        Items(ore: 0, clay: 0, obsidian: 0, geode: 0),
        Items(ore: 1, clay: 0, obsidian: 0, geode: 0))]

    while !queue.isEmpty {
        if queue.first!.minute > maxmin {
            maxmin = queue.first!.minute
            processed = Set<State>()
            if queue.count > maxlen {
                queue.sort { $0.val() > $1.val() }
                queue.removeLast(queue.count - maxlen)
                print("cutting queue, fist", queue.first!.minute, queue.first!.val())
                continue
            }
        }
        let s = queue.removeFirst()
        if s.minute == minutes {
            m = max(m, s.values.geode)
            continue
        }
        if processed.contains(s) {
            continue
        }
        processed.insert(s)
        if s.values.ore >= bp.geode.ore && s.values.obsidian >= bp.geode.obsidian {
            queue.append(State(s.minute + 1,
               s.values.add(s.robots.ore - bp.geode.ore, s.robots.clay, s.robots.obsidian - bp.geode.obsidian, s.robots.geode),
               s.robots.add(0, 0, 0, 1)))
        }
        if s.values.ore >= bp.obsidian.ore && s.values.clay >= bp.obsidian.clay {
            queue.append(State(s.minute + 1,
               s.values.add(s.robots.ore - bp.obsidian.ore, s.robots.clay - bp.obsidian.clay, s.robots.obsidian, s.robots.geode),
               s.robots.add(0, 0, 1, 0)))
        }
        if s.values.ore >= bp.clay {
            queue.append(State(s.minute + 1,
               s.values.add(s.robots.ore - bp.clay, s.robots.clay, s.robots.obsidian, s.robots.geode),
               s.robots.add(0, 1, 0, 0)))
        }
        if s.values.ore >= bp.ore {
            queue.append(State(s.minute + 1,
               s.values.add(s.robots.ore - bp.ore, s.robots.clay, s.robots.obsidian, s.robots.geode),
               s.robots.add(1, 0, 0, 0)))
        }
        queue.append(State(s.minute + 1,
            s.values.add(s.robots.ore, s.robots.clay, s.robots.obsidian, s.robots.geode), s.robots))
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
    print(n1, bp.n)
}
var prod = 1
for bp in blueprints.prefix(3) {
    prod *= maxgeodes(bp, 32, 2000000)
    print(prod, bp.n)
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
