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

	"github.com/gregdel/pushover"
)

// Settings struct for config.json
type ConfigJSON struct {
	TVlogfile   string  `json:"TVlogfile"`
	ConnLogLine float64 `json:"LastConnLogLine"`
	DiscLogLine float64 `json:"LastDiscLogLine"`
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

func verifyNewAlarm(s ConfigJSON, linematch float64, kind string) ConfigJSON {

	switch kind {
	case "connect":
		if s.ConnLogLine < linematch {
			fmt.Println("Launch alarm ", s.ConnLogLine, " & linematch", linematch)
			pushoverNotification("New connection", "CONNECT title")
			//only alarm is sent, update the struct to avoid double alarm
			s.ConnLogLine = linematch
		}
	case "disconnect":
		if s.DiscLogLine < linematch {
			fmt.Println("Launch alarm ", s.DiscLogLine, " & linematch", linematch)
			pushoverNotification("New connection", "DisCONNECT title")
			//only alarm is sent, update the struct to avoid double alarm
			s.DiscLogLine = linematch
		}
	}
	return s
}

func pushoverNotification(messageString string, title string) {
	// Create a new pushover app with a token
	app := pushover.New("az6p3exoje7usacvfgov2i2o84ba37")

	// Create a new recipient
	recipient := pushover.NewRecipient("u8Uc4AH3Z3ixWCMrSZUSenrSHmj3Fp")

	// Create the message to send
	message := &pushover.Message{
		Message:   messageString,
		Title:     title,
		Priority:  pushover.PriorityNormal,
		Timestamp: time.Now().Unix(),
		Sound:     pushover.SoundCosmic,
	}

	// Send the message to the recipient
	response, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}

	// Print the response if you want
	log.Println(response)
}

func main() {

	//s := readConfig("config.json")
	s := readConfig("test.json")

	//check if last run was longer than 5 minutes ago
	//t, err := time.Parse(time.RFC3339Nano, s.LastRun)
	//if err != nil {
	//	log.Panic("unable to parse time: %v", err)
	//}
	//fmt.Println("The time was parsed... ", t)

	fmt.Println("The output was... ", s)
	fmt.Printf("%T", s)

	fmt.Println("The path is... ", s.TVlogfile)
	file, err := os.Open("TeamViewer12_Logfile.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// fun search

	var newConnLineMatch float64 = 0
	var newDiscLineMatch float64 = 0
	var line float64 = 1

	// scanner goes line by line. add counter
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// new connection check string
		if strings.Contains(scanner.Text(), "Starting desktop process for") {
			newConnLineMatch = line
			s = verifyNewAlarm(s, newConnLineMatch, "connect")
		}
		// end session search string = "EndSession(): Session to"
		if strings.Contains(scanner.Text(), "EndSession(): Session to") {
			newDiscLineMatch = line
			s = verifyNewAlarm(s, newDiscLineMatch, "disconnect")
		}
		line++
	}

	// update the time playtime
	nowTime := time.Now()
	s.LastRun = nowTime.Format("2006-01-02 15:04:05")
	//s.LastRun = time.Now().String()

	fmt.Println(s)
	s.saveConfig("test.json")
}
