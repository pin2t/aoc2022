import Foundation

func mix(_ input: [Int], _ indices: [Int], _ multiplicator: Int) -> ([Int], [Int]) {
    var file = Array(input)
    var ind = Array(indices)
    for i in 0..<file.count {
        var from = 0
        while from < ind.count && ind[from] != i { from += 1 }
        var to = from + (file[from] * multiplicator) % (file.count - 1)
        while to <= 0 { to += file.count - 1 }
        while to >= file.count { to -= file.count - 1 }
        let n = file[from]
        let nidx = ind[from]
        if to > from {
            for i in from..<to {
                file[i] = file[i + 1]
                ind[i] = ind[i + 1]
            }
        } else if to != from {
            var j = from
            while j > to {
                file[j] = file[j - 1]
                ind[j] = ind[j - 1]
                j -= 1
            }
        }
        file[to] = n
        ind[to] = nidx
    }
    return (file, ind)
}

var file: [Int] = [], indices: [Int] = []
while let line = readLine() {
    indices.append(file.count)
    file.append(Int(line)!)
}
let (file1, _) = mix(file, indices, 1)
var n1 = 0
for j in 0..<file1.count {
    if file1[j] == 0 {
        n1 = file1[(j + 1000) % file1.count] + file1[(j + 2000) % file1.count] + file1[(j + 3000) % file1.count]
        break
    }
}
var (file2, indices2) = mix(file, indices, 811589153)
for _ in 0..<9 {
    (file2, indices2) = mix(file2, indices2, 811589153)
}
for j in 0..<file2.count {
    if file2[j] == 0 {
        print(n1, file2[(j + 1000) % file2.count] * 811589153 + file2[(j + 2000) % file2.count] * 811589153 + file2[(j + 3000) % file2.count] * 811589153)
        break
    }
}
