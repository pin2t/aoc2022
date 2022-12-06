extension String {
    func unique(_ r: Range<Int>) -> Bool {
        let start = self.index(self.startIndex, offsetBy: r.lowerBound)
        let end = self.index(self.startIndex, offsetBy: r.upperBound)
        var chars: Set<Character> = [], sub = self[start..<end]
        sub.forEach { chars.insert($0) }
        return chars.count == sub.count
    }
}

let chars = readLine()!
var n1 = 0, n2 = 0
for i in 3..<chars.count { 
    if n1 == 0 && chars.unique(i-3..<i+1) { 
        n1 = i + 1
        break 
    }
}
for i in 13..<chars.count {
    if n2 == 0 && chars.unique(i-13..<i+1) { 
        n2 = i + 1 
        break 
    }
}
print(n1, n2)