var calories: [Int] = [0]
while let line = readLine() {
    if line == "" { calories.append(0) }
    else { calories[calories.count - 1] += Int(line) ?? 0 }
}
calories.sort(by: >)
print(calories[0], calories[0] + calories[1] + calories[2])