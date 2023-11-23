package competive_programming

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/* ******************************************************************** */
/* ***************   UTILS: print / splits / read input   ************* */
/* ******************************************************************** */

func GetWriter(outputPath string) (*os.File, *bufio.Writer) {
	if len(outputPath) > 0 {
		file, err := os.Create(outputPath)
		CheckError(err)
		return file, bufio.NewWriterSize(file, 16*1024*1024)
	}
	return os.Stdout, bufio.NewWriterSize(os.Stdout, 16*1024*1024)
}

func ReadFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	CheckError(err)
	return string(data)
}

func OpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	CheckError(err)
	return file
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func FlushWriter(writer *bufio.Writer) {
	err := writer.Flush()
	CheckError(err)
}

func CloseFile(file *os.File) {
	err := file.Close()
	CheckError(err)
}

func GetLines(input io.Reader) []string {
	reader := bufio.NewReaderSize(input, 16*1024*1024)
	var lines = make([]string, 0, 128)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// Print prints a string using a Writer
func Print(writer *bufio.Writer, data string) {
	_, err := fmt.Fprintf(writer, data)
	CheckError(err)
}

// Out prints a right trimmed string using a writer (specific to competitive programming platforms - avoid extra spaces at the end)
func Out(writer *bufio.Writer, data string) {
	trimRight := strings.TrimRight(data, " ")
	Print(writer, trimRight)
}

func Split(line string, separator string) []string {
	return strings.Split(line, separator)
}

func SplitInts(line string, separator string) []int {
	items := Split(line, separator)
	list := make([]int, 0, len(items))
	for _, item := range items {
		list = append(list, ParseInt(item))
	}
	return list
}

func SplitLongs(line string, separator string) []int64 {
	items := Split(line, separator)
	list := make([]int64, 0, len(items))
	for _, item := range items {
		list = append(list, ParseLong(item))
	}
	return list
}

func SplitDoubles(line string, separator string) []float64 {
	items := Split(line, separator)
	list := make([]float64, 0, len(items))
	for _, item := range items {
		list = append(list, ParseDouble(item))
	}
	return list
}

func GetInt(line string, part int) int {
	split := Split(line, " ")
	return ParseInt(split[part])
}

func GetLong(line string, part int) int64 {
	split := Split(line, " ")
	return ParseLong(split[part])
}

func GetDouble(line string, part int) float64 {
	split := Split(line, " ")
	return ParseDouble(split[part])
}

func ParseInt(s string) int {
	value, err := strconv.ParseInt(s, 10, 32)
	CheckError(err)
	return int(value)
}

func ParseLong(s string) int64 {
	value, err := strconv.ParseInt(s, 10, 64)
	CheckError(err)
	return value
}

func ParseDouble(s string) float64 {
	value, err := strconv.ParseFloat(s, 64)
	CheckError(err)
	return value
}

/* ********************************************************************* */
/* *****************   ALGORITHMIC: commons operations   *************** */
/* ********************************************************************* */

func Frequences(list []int) ([]int, map[int]int) {
	// creating map of frequencies
	freqMap := make(map[int]int, len(list))
	for _, item := range list {
		freqMap[item] += 1
	}
	// sorted list of items without repetition
	sortedItems := make([]int, 0, len(freqMap))
	for key, _ := range freqMap {
		sortedItems = append(sortedItems, key)
	}
	sort.Ints(sortedItems)
	return sortedItems, freqMap
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxIntList(list []int) int {
	max := list[0]
	for _, item := range list {
		if item > max {
			max = item
		}
	}
	return max
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinIntList(list []int) int {
	min := list[0]
	for _, item := range list {
		if item < min {
			min = item
		}
	}
	return min
}
