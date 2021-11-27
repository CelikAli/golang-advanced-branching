package main

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
	Name     string   `json:"model`
	Feedback []string `json:"feedback"`
}

// Model array for the feedback.json file
type Model struct{}

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

	// Print ratings for the different vehicles
}

/*
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
*/
