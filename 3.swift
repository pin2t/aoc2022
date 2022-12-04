func priority(_ item: Character) -> Int {
    let chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    return chars.firstIndex(of: item)?.utf16Offset(in: chars) ?? 0 + 1 
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