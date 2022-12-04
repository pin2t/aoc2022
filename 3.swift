func priority(_ item: Character) -> Int {
    let chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    return chars.firstIndex(of: item)!.utf16Offset(in: chars) + 1 
}
var rucksacks: [Set<Character>] = []
var p1 = 0, p2 = 0
while let rucksack = readLine() {
    var first: Set<Character> = [], second: Set<Character> = []
    rucksack.prefix(rucksack.count / 2).forEach { first.insert($0) }
    rucksack.suffix(rucksack.count / 2).forEach { second.insert($0) }
    p1 += priority(first.intersection(second).first!)
    var r: Set<Character> = []
    rucksack.forEach { r.insert($0) }
    rucksacks.append(r)
    if rucksacks.count == 3 {
        p2 += priority(rucksacks[0].intersection(rucksacks[1]).intersection(rucksacks[2]).first!)
        rucksacks = []
    }
}
print(p1, p2)