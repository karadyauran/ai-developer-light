
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Tip struct {
	Message string `json:"message"`
}

func fetchSustainabilityTips() ([]Tip, error) {
	resp, err := http.Get("https://api.example.com/sustainability-tips")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var tips []Tip
	err = json.Unmarshal(body, &tips)
	if err != nil {
		return nil, err
	}
	return tips, nil
}

func main() {
	tips, err := fetchSustainabilityTips()
	if err != nil {
		log.Fatal(err)
	}
	for _, tip := range tips {
		log.Println("Tip: ", tip.Message)
	}