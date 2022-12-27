import Foundation

let packet = try JSONSerialization.jsonObject(with: Data(readLine()!.utf8)) as? [Any]
print(packet!)

func compare(_ p1: [Any], _ p2: [Any]) -> Int {

}