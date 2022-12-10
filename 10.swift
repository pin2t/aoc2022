import Foundation

var strength = 0, x = 1, cycle = 1
var row = ""
var crt: [String] = []
func tick() {
    if [20, 60, 100, 140, 180, 220].contains(cycle) { 
        strength += cycle * x 
    }
    if (cycle - 1) % 40 == 0 {
        crt.append(row)
        row = ""
    }
    let sprite = (cycle - 1) % 40
    row.append(sprite >= x - 1 && sprite <= x + 1 ? "#" : ".")
    cycle += 1
}
while let line = readLine() {
    let words = line.components(separatedBy: [" "])
    switch words[0] {
    case "noop": 
        tick()
    case "addx":    
        tick()
        tick()
        x += Int(words[1])!
    default: break
    }
}
crt.append(row)
print(strength, crt.joined(separator: "\n"))