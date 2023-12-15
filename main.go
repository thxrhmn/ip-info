package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

type ASN struct {
	Network     string `json:"network"`
	ASN         int    `json:"asn"`
	Domain      string `json:"domain"`
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
}

type Company struct {
	Network string `json:"network"`
	Domain  string `json:"domain"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Privacy struct {
	VPN     bool `json:"vpn"`
	Proxy   bool `json:"proxy"`
	Tor     bool `json:"tor"`
	Hosting bool `json:"hosting"`
}

type Abuse struct {
	Network     string `json:"network"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	CountryCode string `json:"countryCode"`
}

type LocationInfo struct {
	IP         string  `json:"ip"`
	Latitude   string  `json:"latitude"`
	Longitude  string  `json:"longitude"`
	Country    string  `json:"country"`
	Region     string  `json:"region"`
	City       string  `json:"city"`
	PostalCode string  `json:"postalCode"`
	UTCOffset  string  `json:"utcOffset"`
	ASN        ASN     `json:"asn"`
	Company    Company `json:"company"`
	Privacy    Privacy `json:"privacy"`
	Abuse      Abuse   `json:"abuse"`
}

func RemoveText(text, word string) string {
	filteredText := strings.ReplaceAll(text, word, "")

	return filteredText
}

func IpInfo() {
	page := rod.New().MustConnect().MustPage("https://ipinfo.ai/")
	page.MustWindowFullscreen()
	time.Sleep(3 * time.Second)
	page.MustElement(`[type="text"]`).MustSelectAllText().MustInput("")
	page.MustElement("input").MustInput("167.94.145.52") // change to your/another ip address
	page.KeyActions().Press(input.ControlLeft).Type(input.Enter).MustDo()

	time.Sleep(3 * time.Second)
	container := page.MustElement("div.jv-braced")
	text := container.MustText()
	resultText := RemoveText(text, "collapse")

	// change json string to struct/object
	var locationInfo LocationInfo
	err := json.Unmarshal([]byte(resultText), &locationInfo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// change struct to json with indent
	jsonData, err := json.MarshalIndent(locationInfo, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// time now
	currentTime := time.Now()
	id := currentTime.Format("2006_01_02_15_04_05")

	// write json to file output.json
	err = ioutil.WriteFile(id+"_output.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("The JSON data has been saved into a file " + id + "_output.json")
	page.MustClose()
}

func main() {
	IpInfo()
	time.Sleep(time.Hour)
}
