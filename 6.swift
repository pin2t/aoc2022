extension String {
    subscript(_ r: Range<Int>) -> Substring {
        let start = self.index(self.startIndex, offsetBy: r.lowerBound)
        let end = self.index(self.startIndex, offsetBy: r.upperBound)
        return self[start..<end]
    }
}

extension Substring {
    var unique: Bool {
        var chars: Set<Character> = []
        self.forEach { chars.insert($0) }
        return chars.count == self.count
    }
}

func firstUnique(_ s: String, _ n: Int) -> Int {
    for i in n-1..<s.count { 
        if s[i-n+1..<i+1].unique { 
            return i
        }
    }
    return 0
}

let chars = readLine()!
print(firstUnique(chars, 4) + 1, firstUnique(chars, 14) + 1)