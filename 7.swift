import Foundation

class Directory {
    var subDirs: [String: Directory] = [:]
    var files: [String: Int] = [:]
    var parent: Directory?
    var size: Int {
        return files.values.reduce(0, +) + subDirs.values.map { $0.size }.reduce(0, +)
    }
    var allSizes: [Int] {
        var sizes: [Int] = []
        self.forEach { sizes.append($0.size) }
        return sizes
    }

    init() { }
    init(_ p: Directory) { parent = p }
    func forEach(_ f: (Directory) -> Void) {
        f(self)
        subDirs.forEach { $0.value.forEach(f) }
    }
}

var root = Directory()
var current = root
while let line = readLine() {
    let words = line.components(separatedBy: [" "])
    if words[0] == "$" { 
        if words[1] == "cd" {
            switch words[2] {
                case "/":  current = root
                case "..": current = current.parent!
                default:   current = current.subDirs[words[2]]!
            }
        }
    } else if words[0] == "dir" {
        current.subDirs[words[1]] = Directory(current)
    } else {
        current.files[words[1]] = Int(words[0])!
    }
}
print(root.allSizes.filter {$0 <= 100000}.reduce(0, +), 
      root.allSizes.filter {$0 >= 30000000-(70000000-root.size)}.min() ?? 0)