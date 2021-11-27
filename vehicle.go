package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type vehicle interface{}

type car struct {
	model       string
	make        string
	typeVehicle string
}

type truck struct {
	model       string
	make        string
	typeVehicle string
}

type bike struct {
	model string
	make  string
}

type feedbackResult struct {
	feedbackTotal    int
	feedbackPositive int
	feedbackNegative int
	feedbackNeutral  int
}

var vehicleResult map[string]feedbackResult
var inventory []vehicle

// Values array for the feedback.json file
type Values struct {
	Models   []Model  `json:"values"`
	Name     string   `json:"model"`
	Feedback []string `json:"feedback"`
}

// Model array for the feedback.json file
type Model struct {
	Name     string   `json:"model"`
	Feedback []string `json:"feedback"`
}

type rating float32

const (
	extraPositive rating = 1.2
	positive      rating = 0.6
	negative      rating = -0.6
	initial       rating = 5.0
	extraNegative rating = -1.2
)

func init() {

	inventory = []vehicle{
		bike{"ftr 1200", "indian"},
		bike{"iron 1200", "harley"},
		car{"sonata", "hyundai", "sedan"},
		car{"santafe", "hyundai", "suv"},
		car{"civic", "honda", "hatchback"},
		car{"a5", "audi", "coupe"},
		car{"mazda6", "mazda", "sedan"},
		car{"crv", "honda", "suv"},
		car{"camry", "toyota", "sedan"},
		truck{"f-150", "ford", "truck"},
		truck{"ram1500", "dodge", "truck"}}
	vehicleResult = make(map[string]feedbackResult)

}

func main() {

	// Generate ratings for the different vehicles
	generateRating()

	// Print ratings for the different vehicles
}

func readJSONFile() Values {
	jsonFile, err := os.Open("feedback.json")

	if err != nil {
		log.Fatal("File not found")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var content Values
	json.Unmarshal(byteValue, &content)

	return content
}

func generateRating() {
	f := readJSONFile()

	for _, v := range f.Models {
		var vehResult feedbackResult
		var vehRating rating
		for _, msg := range v.Feedback {
			if text := strings.Split(msg, " "); len(text) >= 5 {
				vehRating = 5.0
				vehResult.feedbackTotal++
				for _, word := range text {
					switch s := strings.Trim(strings.ToLower(word), " ,.,!,?,\t,\n,\r"); s {
					case "pleasure", "impressed", "wonderful", "fantastic", "splendid":
						vehRating += extraPositive
					case "help", "helpful", "thanks", "thank you", "happy":
						vehRating += positive
					case "not helpful", "sad", "angry", "improve", "annoy":
						vehRating += negative
					case "pathetic", "bad", "worse", "unfortunately", "agitated", "frustrated":
						vehRating += extraNegative
					}
				}
				switch {
				case vehRating > 8.0:
					vehResult.feedbackPositive++
				case vehRating >= 4.0 && vehRating <= 8.0:
					vehResult.feedbackNeutral++
				case vehRating < 4.0:
					vehResult.feedbackNegative++
				}
			}
		}
		vehicleResult[v.Name] = vehResult
	}
}
