import Foundation

extension String {
    var numbers: [Int] {
        guard let regex = try? NSRegularExpression(pattern: "\\d+", options: [.caseInsensitive]) else { 
            return [] 
        }
        let matches  = regex.matches(in: self, options: [], range: NSMakeRange(0, self.count))
        return matches.map { match in
            String(self[Range(match.range, in: self)!])
        }.map { 
            Int($0)! 
        }
    }
}

while let line = readLine() {
    if line.hasPrefix("Monkey ") {
        
    }
}