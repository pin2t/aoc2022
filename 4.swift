import Foundation
var n1 = 0, n2 = 0
while let line = readLine() {
    let pairs = line.components(separatedBy: ["-", ","])
    if Int(pairs[0])! <= Int(pairs[2])! && Int(pairs[1])! >= Int(pairs[3])! ||
       Int(pairs[2])! <= Int(pairs[0])! && Int(pairs[3])! >= Int(pairs[1])! { n1 += 1 }
    if Int(pairs[1])! >= Int(pairs[2])! && Int(pairs[3])! >= Int(pairs[0])! { n2 += 1 }  
}
print(n1, n2)