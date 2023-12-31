package cp

import "strings"

func Split(line string, separator string) []string {
	return strings.Split(line, separator)
}

func SplitInts(line string, separator string) []int {
	items := Split(line, separator)
	list := make([]int, 0, len(items))
	for i := range items {
		item := Trim(items[i])
		if item != "" {
			list = append(list, ParseInt(item))
		}
	}
	return list
}

func SplitLongs(line string, separator string) []int64 {
	items := Split(line, separator)
	list := make([]int64, 0, len(items))
	for i := range items {
		item := Trim(items[i])
		if item != "" {
			list = append(list, ParseLong(item))
		}
	}
	return list
}

func SplitDoubles(line string, separator string) []float64 {
	items := Split(line, separator)
	list := make([]float64, 0, len(items))
	for i := range items {
		item := Trim(items[i])
		if item != "" {
			list = append(list, ParseDouble(item))
		}
	}
	return list
}

func SplitGetAt(line string, separator string, index int) string {
	split := Split(line, separator)
	return split[index]
}

func SplitGetIntAt(line string, separator string, index int) int {
	split := Split(line, separator)
	return ParseInt(split[index])
}

func SplitGetLongAt(line string, separator string, index int) int64 {
	split := Split(line, separator)
	return ParseLong(split[index])
}

func SplitGetDoubleAt(line string, separator string, index int) float64 {
	split := Split(line, separator)
	return ParseDouble(split[index])
}
