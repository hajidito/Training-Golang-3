package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"scheduler/controller"
	"strconv"

	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()
	go func() {
		for {
			select {
			case <-ticker.C:
				GetData()
			}
		}
	}()
	e := echo.New()
	e.GET("/", controller.UpdateWeather)
	e.Logger.Fatal(e.Start(":9000"))
}

func GetData() {
	water := rand.Intn(100) + 1
	wind := rand.Intn(100) + 1
	url := "http://localhost:9000/?water=" + strconv.Itoa(water) + "&wind=" + strconv.Itoa(wind)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var result bytes.Buffer
	err = json.Indent(&result, body, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.String())
}
