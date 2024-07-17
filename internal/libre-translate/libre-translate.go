package libre_translate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Translate struct {
	Alternatives   []string `json:"alternatives"`
	TranslatedText string   `json:"translatedText"`
}

func LibreTranslate(q string,
	source string,
	target string,
	format string,
	alternatives int,
	apiKey string) (*Translate, error) {
	const op = "libre-translate.LibreTranslate"

	req, err := http.NewRequest("POST",
		"http://localhost:5000/translate",
		nil)
	if err != nil {
		// TODO: logging
		fmt.Println(err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	que := req.URL.Query()
	que.Add("q", q)
	que.Add("source", source)
	que.Add("target", target)
	que.Add("format", format)
	que.Add("alternatives", strconv.Itoa(alternatives))
	que.Add("api_key", apiKey)
	req.URL.RawQuery = que.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// TODO: logging
		fmt.Println(err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		// TODO: logging
		fmt.Println(err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	println(string(respBody))
	res := &Translate{}
	err = json.Unmarshal(respBody, res)
	if err != nil {
		// TODO: logging
		fmt.Println(err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return res, nil
}
