package main

/*
Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names,
begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value
by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53,
is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type node struct {
	value uint
	total uint
	left  *node
	right *node
}

func main() {
	getNamesScore("p022_names.txt")
}

func getNamesScore(fileName string) (scoreTotal int64) {
	fmt.Println(fmt.Sprintf("calculating the names scores for file %s", fileName))
	names := parseFile(fileName)
	sort.Strings(names)
	scoreTotal = getTotalScore(names)
	fmt.Println(fmt.Sprintf("total of names scores: %v", scoreTotal))
	return
}

func getTotalScore(names []string) (totalScore int64) {
	for index, name := range names {
		thisScore := getScoreForName(int64(index+1), name)
		totalScore += thisScore
	}
	return
}

func getScoreForName(index int64, name string) (score int64) {
	//TODO: logic here
	alphaValue := getAlphaValue(name)
	score = alphaValue * index
	return
}

func getAlphaValue(name string) (alphaValue int64) {
	for _, char := range strings.Split(name, "") {
		if char != "\"" {
			alphaValue += getAlpha(char)
		}
	}
	return
}

func getAlpha(char string) (num int64) {
	origin := byte('A') - 1
	charNum := byte(strings.ToUpper(char)[0])
	num = int64(charNum - origin)
	return
}

func parseFile(filename string) (contents []string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = strings.Split(scanner.Text(), ",")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}
