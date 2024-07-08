package main

import (
	"fmt"
	"strings"
)

func main() {
	reverseAlphabet("NEGIE1")
	longestSentence("Saya sangat senang mengerjakan soal algoritma")
	
	INPUT := []string{"xc", "dz", "bbb", "dz"}
	QUERY := []string{"bbb", "ac", "dz"}
	countSameWords(INPUT,QUERY)

	matrix := [][]int{{1, 2, 0}, {4, 5, 6}, {7, 8, 9}}
   	diagonalMatrix(matrix)
}

func reverseAlphabet(s string) {
	reverseVersion := make([]byte, len(s))
	for i := len(s) - 2; i >= 0; i-- {
		reverseVersion[len(s)-2-i] = s[i]
	}
	reverseVersion[len(s)-1] = s[len(s)-1]
	fmt.Println(string(reverseVersion))
}

func longestSentence(s string){
	words := strings.Fields(s)
	longest := ""
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	fmt.Printf("%s: %d character\n", longest, len(longest))
}

func countSameWords(INPUT []string, QUERY []string) {
	countMap := make(map[string]int)
	for _, word := range INPUT {
		countMap[word]++
	}
	var result []int
	for _, word := range QUERY {
		count, exists := countMap[word]
		if exists {
			result = append(result, count)
		} else {
			result = append(result, 0)
		}
	}

	fmt.Println("Output: ", result)

}


func diagonalMatrix(matrix [][]int) {
    n := len(matrix)
	var sumFirst, sumSec int

    for i := 0; i < n; i++ {
        sumFirst += matrix[i][i]
        sumSec += matrix[i][n-i-1]
    }
	
    difference := sumFirst - sumSec
	fmt.Println(sumFirst, sumSec)
	fmt.Println(difference)
}
