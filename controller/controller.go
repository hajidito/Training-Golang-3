package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"scheduler/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ReadJSONFile(fileName string) (*model.Weather, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var data model.Weather
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func EditWeather(data *model.Weather, water int, wind int, infoWater string, infoWind string) model.Weather {
	data.Water = water
	data.Wind = wind
	data.InfoWater = infoWater
	data.InfoWind = infoWind
	return *data
}

func UpdateWeather(ctx echo.Context) error {
	water, err := strconv.Atoi(ctx.QueryParam("water"))
	if err != nil {
		fmt.Println("Error converting water to int:", err)
		return err
	}
	wind, err := strconv.Atoi(ctx.QueryParam("wind"))
	if err != nil {
		fmt.Println("Error converting wind to int:", err)
		return err
	}

	fileName := "weather.json"
	data, err := ReadJSONFile(fileName)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return err
	}

	// json.Marshal(data)
	var infoWater string
	var infoWind string

	switch {
	case water <= 5:
		infoWater = "status water aman"
	case water >= 6 && water <= 8:
		infoWater = "status water siaga"
	case water > 8:
		infoWater = "status water bahaya"
	}

	switch {
	case wind <= 6:
		infoWind = "status wind aman"
	case wind >= 7 && wind <= 15:
		infoWind = "status wind siaga"
	case wind > 15:
		infoWind = "status wind bahaya"
	}

	editedData := EditWeather(data, water, wind, infoWater, infoWind)

	err = WriteJSONFile(fileName, editedData)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return err
	}
	return ctx.JSON(http.StatusOK, data)
}

func WriteJSONFile(fileName string, data model.Weather) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
