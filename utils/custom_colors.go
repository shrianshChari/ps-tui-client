package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var CustomColors map[string]string

func initializeCustomColors() {
	if len(CustomColors) != 0 {
		return
	}
	
	curTime := time.Now()
	fileInfo, err := os.Stat("colors.json")
	// If the file exists and is not 24 hours out of date
	if err != nil || curTime.Sub(fileInfo.ModTime()).Hours() >= 24 {
		fileURL := "https://play.pokemonshowdown.com/config/colors.json"
		err := downloadFile(fileURL)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	content, err := os.ReadFile("colors.json")
	if (err != nil) {
		// I don't want to import the log package here just yet
		fmt.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(content, &CustomColors)
	if err != nil {
		// I don't want to import the log package here just yet
		fmt.Println(err)
		os.Exit(1)
	}
}

func downloadFile(url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}


	defer resp.Body.Close()

	out, err := os.Create("colors.json")
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
