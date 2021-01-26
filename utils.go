package chuckjoke

// Credits
// code originally taken from a project
// https://github.com/farhaanbukhsh/chuck/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetNextJokeValue gets the next chuck joke and returns the text value of the joke
func GetNextJokeValue() (string, error) {
	joke, err := GetNextJoke()
	if err != nil {
		return "", err
	}

	return joke.Value, nil
}

// GetNextJoke gets the next chuck joke and returns the entire struct
func GetNextJoke() (Joke, error) {
	req, err := http.NewRequest("GET", chuckAPI, nil)
	var joke Joke
	if err != nil {
		return joke, fmt.Errorf("No request formed %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return joke, fmt.Errorf("No response: %v", err)
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return joke, fmt.Errorf("Read error")
	}

	if err = json.Unmarshal(respData, &joke); err != nil {
		return joke, fmt.Errorf("Error in unmarsheling, %v", err)
	}

	return joke, nil
}
