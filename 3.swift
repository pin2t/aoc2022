func priority(_ item: Character) -> Int {
    if item >= "a" && item <= "z" { return Int(item.utf8.first! - "a".utf8.first!) + 1 }
    if item >= "A" && item <= "Z" { return Int(item.utf8.first! - "A".utf8.first!) + 27 }
    return 0
}
var rucksacks: [String] = []
var p1 = 0, p2 = 0
while let rucksack = readLine() {
    let first = rucksack.prefix(rucksack.count / 2), second = rucksack.suffix(rucksack.count / 2)
    for item in  first {
        if second.contains(item) {
            p1 += priority(item)
            break
        }
    }
    rucksacks.append(rucksack)
    if rucksacks.count == 3 {
        for item in rucksacks[0] {
            if rucksacks[1].contains(item) && rucksacks[2].contains(item) {
                p2 += priority(item)
                break
            }
        }
        rucksacks = []
    }
}
print(p1, p2)