package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'getGroupedAnagrams' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING_ARRAY words as parameter.
 */

func getGroupedAnagrams(words []string) int32 {
	// Write your code here

	m := make(map[string]bool)
	for _, value := range words {
		set := make([]int, 0, len(value))

		for _, c := range value {
			set = append(set, int(c))
		}
		sort.Ints(set)

		key := ""
		for _, c := range set {
			key += strconv.Itoa(c)
		}

		m[key] = true
	}

	return int32(len(m))
}

func main() {
	set := []string{"cat", "listen", "silent", "kitten", "salient"}
	count := getGroupedAnagrams(set)
	fmt.Println(count)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
