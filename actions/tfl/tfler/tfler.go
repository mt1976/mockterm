package tfler

import (
	"errors"
	"strings"
)

// Stations - GetTubeLines
func GetTubeLines() ([]Line, error) {

	var lines []Line

	lines = append(lines, Line{Name: "Station1", Code: "Code1", Status: "Status1", Type: "Type1", Line: "Line1", StatusCode: "StatusCode1"})
	lines = append(lines, Line{Name: "Station2", Code: "Code2", Status: "Status2", Type: "Type2", Line: "Line2", StatusCode: "StatusCode2"})
	lines = append(lines, Line{Name: "Station3", Code: "Code3", Status: "Status3", Type: "Type3", Line: "Line3", StatusCode: "StatusCode3"})
	lines = append(lines, Line{Name: "Station4", Code: "Code4", Status: "Status4", Type: "Type4", Line: "Line4", StatusCode: "StatusCode4"})
	lines = append(lines, Line{Name: "Station5", Code: "Code5", Status: "Status5", Type: "Type5", Line: "Line5", StatusCode: "StatusCode5"})
	lines = append(lines, Line{Name: "Station6", Code: "Code6", Status: "Status6", Type: "Type6", Line: "Line6", StatusCode: "StatusCode6"})
	lines = append(lines, Line{Name: "Station7", Code: "Code7", Status: "Status7", Type: "Type7", Line: "Line7", StatusCode: "StatusCode7"})
	lines = append(lines, Line{Name: "Station8", Code: "Code8", Status: "Status8", Type: "Type8", Line: "Line8", StatusCode: "StatusCode8"})
	lines = append(lines, Line{Name: "Station9", Code: "Code9", Status: "Status9", Type: "Type9", Line: "Line9", StatusCode: "StatusCode9"})
	lines = append(lines, Line{Name: "Station10", Code: "Code10", Status: "Status10", Type: "Type10", Line: "Line10", StatusCode: "StatusCode10"})
	return lines, nil
}

func IsValidLineName(lineIn string) (bool, error) {
	return isValidLine(lineIn, "NAME")

}

func IsValidLineCode(lineIn string) (bool, error) {
	return isValidLine(lineIn, "CODE")
}

func isValidLine(lineIn, poo string) (bool, error) {
	lines, _ := GetTubeLines()

	for _, line := range lines {
		switch upcase(poo) {
		case "NAME":
			if upcase(line.Name) == upcase(lineIn) {
				return true, nil
			}
		case "CODE":
			if upcase(line.Code) == upcase(lineIn) {
				return true, nil
			}
		default:
			return false, errors.New("invalid field")
		}
	}
	return false, nil
}

func upcase(line string) string {
	return strings.ToUpper(line)
}
