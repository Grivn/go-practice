package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 * Complete the 'strokesRequired' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING_ARRAY picture as parameter.
 */

type pixel struct {
	val  rune
	done bool
}

func strokesRequired(picture []string) int32 {
	// Write your code here
	if len(picture) <= 1 {
		return int32(len(picture))
	}

	pic := make([][]*pixel, 0, len(picture))
	for _, l := range picture {
		line := make([]*pixel, 0, len(l))
		for _, c := range l {
			line = append(line, &pixel{rune(c), false})
		}
		pic = append(pic, line)
	}

	count := 0
	for i, line := range pic {
		for j, item := range line {

			item.done = true

			render(i, j, item.val, pic)

			fmt.Println(count)
			for _, line := range pic {
				for _, val := range line {
					fmt.Printf("%d %v | ", val.val, val.done)
				}
				fmt.Printf("\n")
			}
			fmt.Printf("\n")
		}
	}

	return int32(count)
}

func render(i, j int, val rune, pic [][]*pixel) {
	if i > 0 && pic[i-1][j].val == val {
		pic[i][j+1].done = true
		render(i-1, j, val, pic)
	}

	if j > 0 && pic[i][j-1].val == val {
		pic[i][j+1].done = true
		render(i, j-1, val, pic)
	}

	if i < len(pic) && pic[i+1][j].val == val {
		pic[i][j+1].done = true
		render(i+1, j, val, pic)
	}

	if j < len(pic[0]) && pic[i][j+1].val == val {
		pic[i][j+1].done = true
		render(i, j+1, val, pic)
	}
}

func main() {
	//h := 3
	sets := []string{"bbba", "abba", "acaa", "aaac"}

	count := strokesRequired(sets)
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
