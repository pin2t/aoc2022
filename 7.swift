import Foundation

class Directory {
    var subDirs: [String: Directory] = [:]
    var files: [String: Int] = [:]
    var parent: Directory?
    var size: Int {
        return files.values.reduce(0, +) + subDirs.values.map { $0.size }.reduce(0, +)
    }

    func forEach(_ f: (Directory) -> Void) {
        f(self)
        subDirs.forEach { $0.value.forEach(f) }
    }
}

var root = Directory()
var current = root
while let line = readLine() {
    let words = line.components(separatedBy: [" "])
    if words[0] == "$" && words[1] == "ls" { 
        continue 
    } else if words[0] == "$" && words[1] == "cd" {
        switch words[2] {
            case "/":  current = root
            case "..": current = current.parent!
            default:   current = current.subDirs[words[2]]!
        }
    } else if words[0] == "dir" {
        let dir = Directory()
        dir.parent = current
        current.subDirs[words[1]] = dir
    } else {
        current.files[words[1]] = Int(words[0])!
    }
}
var total = 0, smallest = 10000000000, rootSize = root.size
root.forEach {
    if $0.size <= 100000 { 
        total += $0.size 
    }
    if ($0.size >= 30000000-(70000000-rootSize)) && $0.size < smallest { 
        smallest = $0.size 
    }
}
print(total, smallest)