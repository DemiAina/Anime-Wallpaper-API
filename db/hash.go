package db

import (
	"fmt"
	"math"
	"strings"
)

var supportedFileTypes = []string{"jpg", "png"}

func GenerateHash(filename string) string {
	fileType := getFileType(filename)
	fileTypeRep := convertToASCII([]string{fileType})[0]
	fileTypeNumRep := convertToNumberRep(fileType)
	countBeforeType := countInFile(filename)
	fileBeforeHash := filename[:len(filename)-len(fileType)-1]
	fileAfterHash := convertToHash(fileBeforeHash, countBeforeType)
	completeHash := fmt.Sprintf("%s%s%s%d", fileTypeNumRep, fileTypeRep, fileAfterHash, countBeforeType)
	return completeHash
}
func convertToASCII(fileTypes []string) []string {
	convertedList := make([]string, len(fileTypes))
	var temp strings.Builder
	for _, fileType := range fileTypes {
		for _, char := range fileType {
			temp.WriteString(fmt.Sprintf("%d", int(char)))
		}
		convertedList = append(convertedList, temp.String())
		temp.Reset()
	}
	return convertedList
}


func convertToNumberRep(fileType string) string {
	var numRep strings.Builder
	for _, char := range fileType {
		numRep.WriteString(fmt.Sprintf("%d", int(char)))
	}
	return numRep.String()
}


func getFileType(filename string) string {
	splitted := strings.Split(filename, ".")
	if len(splitted) > 1 {
		return splitted[len(splitted)-1]
	}
	return ""
}

func countInFile(filename string) int {
	count := 0
	for _, char := range filename {
		if char == '.' {
			break
		}
		count++
	}
	return count
}

func convertToHash(fileToHash string, count int) string {
	temp := ""
	oneQuarter := int(math.Floor(float64(count) / 4))
	threeQuarters := int(math.Floor(float64(count) * 3 / 4))
	middle := int(math.Floor(float64(count) / 2))

	oneQuarterStr := string(fileToHash[oneQuarter])
	threeQuartersStr := string(fileToHash[threeQuarters])
	middleStr := string(fileToHash[middle])

	temp += oneQuarterStr
	temp += threeQuartersStr
	temp += middleStr

	return temp
}

