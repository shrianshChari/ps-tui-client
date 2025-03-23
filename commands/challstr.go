package commands

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	// Autoloads .env variables into os.Getenv
	_ "github.com/joho/godotenv/autoload"
)

type ChallStrData struct {
	Actionsuccess bool
	Assertion     string
	Curuser       struct {
		Loggedin bool
		Username string
		Userid   string
	}
}

func ChallStr(challStr string) (ChallStrData, error) {
	data := ChallStrData{}

	user := os.Getenv("PS_USERNAME")
	pass := os.Getenv("PS_PASSWORD")

	parameters := map[string]string{
		"name":     user,
		"pass":     pass,
		"challstr": challStr,
	}

	parameterJson, err := json.Marshal(parameters)
	if err != nil {
		return data, err
	}

	challStrUrl := url.URL{
		Scheme: "https",
		Host:   "play.pokemonshowdown.com",
		Path:   "api/login",
	}

	log.Printf("ChallStr URL: %s\n", challStrUrl.String())

	resp, err := http.Post(challStrUrl.String(), "application/json", bytes.NewBuffer(parameterJson))
	if err != nil {
		return data, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	body = body[1:]

	log.Printf("Body: %s\n", body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
