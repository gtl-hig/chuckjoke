package chuckjoke

const chuckAPI = "http://api.chucknorris.io/jokes/random"

// Chucknorris is the struct used to  unmarshal the JSON response from the URL
type Joke struct {
	Category []string `json:"category"`
	IconURL  string   `json:"icon_url"`
	ID       string   `json:"id"`
	URL      string   `json:"url"`
	Value    string   `json:"value"`
}
