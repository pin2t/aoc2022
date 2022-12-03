func priority(c: Character) -> Int {
    if c >= "a" && c <= "z" { return Int(c.utf8.first! - "a".utf8.first!) + 1 }
    if c >= "A" && c <= "Z" { return Int(c.utf8.first! - "A".utf8.first!) + 27 }
    return 0
}
var rucksacks: [String] = []
var p1 = 0, p2 = 0
while let rucksack = readLine() {
    for item in rucksack.prefix(rucksack.count / 2) {
        if rucksack.suffix(rucksack.count / 2).contains(item) {
            p1 += priority(c: item)
            break
        }
    }
    rucksacks.append(rucksack)
    if rucksacks.count == 3 {
        for item in rucksacks[0] {
            if rucksacks[1].contains(item) && rucksacks[2].contains(item) {
                p2 += priority(c: item)
                break
            }
        }
        rucksacks.removeAll(keepingCapacity: true)
    }
}
print(p1, p2)