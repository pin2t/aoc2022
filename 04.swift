import Foundation
var n1 = 0, n2 = 0
while let line = readLine() {
    let pairs = line.components(separatedBy: ["-", ","]).map { Int($0)! }
    if pairs[0] <= pairs[2] && pairs[1] >= pairs[3] ||
       pairs[2] <= pairs[0] && pairs[3] >= pairs[1] { n1 += 1 }
    if pairs[1] >= pairs[2] && pairs[3] >= pairs[0] { n2 += 1 }  
}
print(n1, n2)