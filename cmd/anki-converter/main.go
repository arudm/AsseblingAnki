package main

import (
	"AnkiConverter/internal/libre-translate"
	"encoding/json"
	"fmt"
)

func main() {

	translated, err := libre_translate.LibreTranslate("Я молодой и красивый заяц",
		"ru",
		"en",
		"text",
		3,
		"")
	if err != nil {
		// TODO: logging
		fmt.Println(err)
	}

	res, _ := json.Marshal(translated)

	println(string(res))
}
