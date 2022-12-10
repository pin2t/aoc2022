var scores = ["A X": [1 + 3, 3 + 0], "A Y": [2 + 6, 1 + 3], "A Z": [3 + 0, 2 + 6],
   		      "B X": [1 + 0, 1 + 0], "B Y": [2 + 3, 2 + 3], "B Z": [3 + 6, 3 + 6],
		      "C X": [1 + 6, 2 + 0], "C Y": [2 + 0, 3 + 3], "C Z": [3 + 3, 1 + 6]]
var totals = [0, 0]              
while let line = readLine() {
    totals[0] += scores[line]![0]
    totals[1] += scores[line]![1]
}
print(totals[0], totals[1])