package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// Settings struct for config.json
type ConfigJSON struct {
	TVlogfile   string  `json:"TVlogfile"`
	ConnLogLine float64 `json:"LastConnLogLine"`
	LastRun     string  `json:"LastRun"`
}

// Function parses json file into Struct variables
func readConfig(settingFile string) ConfigJSON {
	raw, err := ioutil.ReadFile(settingFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var s ConfigJSON
	json.Unmarshal(raw, &s)
	return s
}

// Function saves json file into Struct variables
func (s *ConfigJSON) saveConfig(settingFile string) error {
	// let's write it back
	out, err := json.Marshal(s)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(settingFile, out, 0600)
	if err != nil {
		return err
	}
	return nil
}

func verifyNewAlarm(s ConfigJSON, linematch float64) {
	// check if new alert is needed or not
	if s.ConnLogLine != linematch {
		fmt.Println("Launch alarm ", s.ConnLogLine, " & linematch", linematch)
		//only alarm is sent, update the struct to avoid double alarm
		s.ConnLogLine = linematch
	}
}

func main() {
	//s := readConfig("config.json")
	s := readConfig("test.json")
	fmt.Println("The output was... ", s)
	fmt.Printf("%T", s)

	fmt.Println("The path is... ", s.TVlogfile)
	file, err := os.Open("TeamViewer12_Logfile.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// fun search
	var linematch float64 = 0
	var line float64 = 1
	// scanner goes line by line. add counter
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// new connection check string
		if strings.Contains(scanner.Text(), "firefox.exe") {
			linematch = line
			verifyNewAlarm(s, linematch)
		}
		line++
	}
	fmt.Println("Found matching last line...", linematch)

	// update the time playtime
	s.LastRun = (time.Now()).String()
	s.saveConfig("test.json")
}
