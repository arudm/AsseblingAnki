package dictionary

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Entry []struct {
	Word       string      `json:"word"`
	Phonetic   string      `json:"phonetic"`
	Phonetics  []Phonetics `json:"phonetics"`
	Meanings   []Meanings  `json:"meanings"`
	License    License     `json:"license"`
	SourceUrls []string    `json:"sourceUrls"`
}

type License struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Phonetics struct {
	Text      string  `json:"text"`
	Audio     string  `json:"audio"`
	SourceURL string  `json:"sourceUrl,omitempty"`
	License   License `json:"license,omitempty"`
}

type Definitions struct {
	Definition string   `json:"definition"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
	Example    string   `json:"example,omitempty"`
}

type Meanings struct {
	PartOfSpeech string        `json:"partOfSpeech"`
	Definitions  []Definitions `json:"definitions"`
	Synonyms     []string      `json:"synonyms"`
	Antonyms     []string      `json:"antonyms"`
}

type Error struct {
	title      string
	message    string
	resolution string
}

func GetDictionary(text string) (*Entry, error) {
	const op = "dictionary.GetDictionary"

	req, err := http.NewRequest("GET",
		"https://api.dictionaryapi.dev/api/v2/entries/en/"+text,
		nil)
	if err != nil {
		// TODO: logging
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// TODO: logging
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		// TODO: logging
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	println(string(respBody))
	var res Entry
	err = json.Unmarshal(respBody, &res)
	if err != nil {
		// TODO: logging
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &res, nil
}
