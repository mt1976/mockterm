package tfler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// Stations - GetTubeLines
func GetTubeLines() ([]Line, error) {

	client := &http.Client{}

	//var lines []Line

	// lines = append(lines, Line{Name: "Bakerloo", Code: "BAK", Status: "Status1", StatusCode: "StatusCode1", Type: []string{"Type1"}})
	// lines = append(lines, Line{Name: "Central", Code: "CEN", Status: "Status3", StatusCode: "StatusCode3", Type: []string{"Type3"}})
	// lines = append(lines, Line{Name: "Circle", Code: "CIR", Status: "Status2", StatusCode: "StatusCode2", Type: []string{"Type2"}})
	// lines = append(lines, Line{Name: "District", Code: "DIS", Status: "Status4", StatusCode: "StatusCode4", Type: []string{"Type4"}})
	// lines = append(lines, Line{Name: "Hammersmith & City", Code: "HAM", Status: "Status5", StatusCode: "StatusCode5", Type: []string{"Type5"}})
	// lines = append(lines, Line{Name: "Jubilee", Code: "JUB", Status: "Status6", StatusCode: "StatusCode6", Type: []string{"Type6"}})
	// lines = append(lines, Line{Name: "Metropolitan", Code: "MET", Status: "Status7", StatusCode: "StatusCode7", Type: []string{"Type7"}})
	// lines = append(lines, Line{Name: "Northern", Code: "NOR", Status: "Status8", StatusCode: "StatusCode8", Type: []string{"Type8"}})
	// lines = append(lines, Line{Name: "Piccadilly", Code: "PIC", Status: "Status9", StatusCode: "StatusCode9", Type: []string{"Type9"}})
	// lines = append(lines, Line{Name: "Victoria", Code: "VIC", Status: "Status10", StatusCode: "StatusCode10", Type: []string{"Type10"}})
	// lines = append(lines, Line{Name: "Waterloo & City", Code: "WAT", Status: "Status11", StatusCode: "StatusCode11", Type: []string{"Type11"}})

	uri := "https://api.tfl.gov.uk/Line/Mode/tube?&app_key=%v"

	uri = fmt.Sprintf(uri, api_key)

	// fmt.Println("URL: ", uri)
	// fmt.Println("URL: ", uri)
	// fmt.Println("URL: ", uri)
	// fmt.Println("URL: ", uri)
	// fmt.Println("URL: ", uri)
	// fmt.Println("URL: ", uri)
	// fmt.Println("URL: ", uri)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Println("ERROR :" + err.Error())
		os.Exit(1)
	}

	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

	// Send the HTTP request to the TfL API
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return []Line{}, err
	}
	defer resp.Body.Close()

	// Check if the status code indicates success
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Received non-200 response code: %d\n", resp.StatusCode)
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println("Response body:", string(bodyBytes))
		return []Line{}, errors.New("non-200 response code received")
	}

	//spew.Dump(resp)
	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return []Line{}, err
	}
	//spew.Dump(body)
	//os.Exit(1)
	var lines LineImport
	err = json.Unmarshal(body, &lines)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return []Line{}, err
	}

	// Print out the line names
	for _, line := range lines {
		fmt.Printf("Line ID: %s, Name: %s\n", line.ID, line.Name)
	}

	//spew.Dump(lines)

	var lineArray []Line
	// loop thgouht the lines and populate the Line array
	for _, line := range lines {
		lineItem := Line{}
		lineItem.ID = line.ID
		lineItem.Code = line.ID
		lineItem.Name = line.Name
		for _, serviceType := range line.ServiceTypes {
			lineItem.Type = append(lineItem.Type, serviceType.Name)
		}
		if len(line.Disruptions) > 0 {
			lineItem.Status = "Issues" + line.Disruptions[0].(string)
		} else {
			lineItem.Status = "Good Service"
		}
		for _, lineDisrupt := range line.Disruptions {
			//lineItem.Disruptions = []string{line.Disruptions[0]}
			lineItem.Disruptions = append(lineItem.Disruptions, lineDisrupt.(string))
		}
		lineItem.LastUpdated = line.Modified
		lineArray = append(lineArray, lineItem)
		//spew.Dump(line)
	}

	return lineArray, nil
}

func IsValidLineName(lineIn string) (bool, error) {
	return isValidLine(lineIn, "NAME")
}

func IsValidLineCode(lineCode string) (bool, error) {
	return isValidLine(lineCode, "CODE")
}

func isValidLine(lineIn, validate string) (bool, error) {
	lines, _ := GetTubeLines()

	for _, line := range lines {
		switch upcase(validate) {
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

func GetTubeLineDetails(lineCode string) (LineDetail, error) {

	client := &http.Client{}

	ok, err := IsValidLineCode(lineCode)
	if !ok {
		return LineDetail{}, err
	}
	//spew.Dump(lineCode)
	baseUrl := "https://api.tfl.gov.uk"
	uri := fmt.Sprintf("%s/Line/%s/StopPoints?app_key=%s", baseUrl, lineCode, api_key)

	fmt.Println("URL: ", uri)
	fmt.Println("URL: ", uri)
	fmt.Println("URL: ", uri)
	fmt.Println("URL: ", uri)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Println("ERROR :" + err.Error())
		os.Exit(1)
	}

	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

	// Send the HTTP request to the TfL API
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return LineDetail{}, err
	}
	defer resp.Body.Close()

	// Check if the status code indicates success
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Received non-200 response code: %d\n", resp.StatusCode)
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println("Response body:", string(bodyBytes))
		return LineDetail{}, errors.New("non-200 response code received")
	}

	//spew.Dump(resp)
	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return LineDetail{}, err
	}
	//spew.Dump(body)
	//os.Exit(1)
	var lineDeets LineImport
	err = json.Unmarshal(body, &lineDeets)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return LineDetail{}, err
	}

	spew.Dump(lineDeets)

	var lineDetail LineDetail
	lineDetail.Code = lineCode
	lineDetail.Name = "DummyLine" + lineCode
	lineDetail.Status = "DummyStatus" + lineCode
	lineDetail.StatusCode = "DummyStatusCode" + lineCode
	lineDetail.Type = []string{"DummyType" + lineCode}
	lineDetail.Narrative = "DummyNarrative" + lineCode

	stationsList, err := GetStations(lineCode, lineDeets)
	if err != nil {
		return LineDetail{}, err
	}
	lineDetail.Stations = stationsList

	return lineDetail, nil
}

func GetStations(lineCode string, details LineImport) ([]Station, error) {
	ok, err := IsValidLineCode(lineCode)
	if !ok {
		return []Station{}, err
	}

	var stations []Station

	for _, line := range details {
		if line.ID == lineCode {
			station := Station{Name: line.Name, Code: line.ID, Status: "Good Service", Type: []string{"Underground"}, LineCode: lineCode}
			stations = append(stations, station)
		}
	}

	spew.Dump(stations)
	//	stations = append(stations, Station{Name: "Station1", Code: "Code1", Status: "Status1", Type: []string{"Type1"}, LineCode: lineCode})
	//	stations = append(stations, Station{Name: "Station2", Code: "Code2", Status: "Status2", Type: []string{"Type2"}, LineCode: lineCode})
	//	stations = append(stations, Station{Name: "Station3", Code: "Code3", Status: "Status3", Type: []string{"Type3"}, LineCode: lineCode})
	//	stations = append(stations, Station{Name: "Station4", Code: "Code4", Status: "Status4", Type: []string{"Type4"}, LineCode: lineCode})
	//	stations = append(stations, Station{Name: "Station5", Code: "Code5", Status: "Status5", Type: []string{"Type5"}, LineCode: lineCode})
	//	stations = append(stations, Station{Name: "Station6", Code: "Code6", Status: "Status6", Type: []string{"Type6"}, LineCode: lineCode})
	//	stations = append(stations, Station{Name: "Station7", Code: "Code7", Status: "Status7", Type: []string{"Type7"}, LineCode: lineCode})

	return stations, nil
}

func GetStationDetails(stationCode string) (StationDetail, error) {
	var stationDetail StationDetail
	stationDetail.Code = stationCode
	stationDetail.Name = "DummyStation" + stationCode
	stationDetail.Status = "DummyStatus" + stationCode
	stationDetail.Type = []string{"DummyType" + stationCode}
	stationDetail.LineCode = "DummyLineCode" + stationCode
	stationDetail.LineName = "DummyLineName" + stationCode
	lineName, err := GetTubeLineDetails(stationDetail.LineCode)
	if err != nil {
		return StationDetail{}, err
	}
	stationDetail.LineName = lineName.Name
	stationDetail.StatusCode = "DummyStatusCode" + stationCode
	stationDetail.Narrative = "DummyStationNarrative" + stationCode

	return stationDetail, nil
}

/////
// Function to fetch line details
// func getLineDetails(lineID string) ([]Station, LineStatus, error) {
//     baseUrl := "https://api.tfl.gov.uk"
//     url := fmt.Sprintf("%s/Line/%s/StopPoints?app_key=%s", baseUrl, lineID, api_key)

//     // Make the HTTP GET request
//     resp, err := http.Get(url)
//     if err != nil {
//         return nil, LineStatus{}, fmt.Errorf("HTTP request failed: %v", err)
//     }
//     defer resp.Body.Close()

//     // Read and parse the response body
//     body, err := ioutil.ReadAll(resp.Body)
//     if err != nil {
//         return nil, LineStatus{}, fmt.Errorf("Failed to read response: %v", err)
//     }

//     // Parse stations
//     stations := []Station{}
//     gjson.GetBytes(body, "#.commonName").ForEach(func(key, value gjson.Result) bool {
//         stations = append(stations, Station{Name: value.String()})
//         return true // keep iterating
//     })

//     // Fetch line status
//     statusUrl := fmt.Sprintf("%s/Line/%s/Status?app_key=%s", baseUrl, lineID, api_key)
//     respStatus, err := http.Get(statusUrl)
//     if err != nil {
//         return nil, LineStatus{}, fmt.Errorf("HTTP request for status failed: %v", err)
//     }
//     defer respStatus.Body.Close()

//     bodyStatus, err := ioutil.ReadAll(respStatus.Body)
//     if err != nil {
//         return nil, LineStatus{}, fmt.Errorf("Failed to read status response: %v", err)
//     }

//     status := LineStatus{}
//     statusResult := gjson.GetBytes(bodyStatus, "0.lineStatuses.0.statusSeverityDescription")
//     status.StatusSeverityDescription = statusResult.String()

//     return stations, status, nil
// }
