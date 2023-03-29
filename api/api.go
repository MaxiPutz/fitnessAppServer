package api

import (
	"encoding/json"
	"fitnessApp/key"
	"fitnessApp/structs"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const fitnessAppAPIUrl = "https://www.strava.com/api/v3"

func GetWorkoutData(id int) structs.WorkoutData {
	fitnessAppUrl := fitnessAppAPIUrl + "/activities/" + strconv.Itoa(id) + "/streams?keys=latlng,watts,heartrate,time,velocity_smooth&resolution=low&key_by_type=true"
	fmt.Printf("fitnessAppUrl: %v\n", fitnessAppUrl)

	req, _ := http.NewRequest("GET", fitnessAppUrl, nil)

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + key.ResponseKey().AccessToken},
	}

	client := http.Client{}

	res, _ := client.Do(req)
	defer res.Body.Close()

	dataByte, _ := ioutil.ReadAll(res.Body)

	workoutData := structs.WorkoutData{}

	json.Unmarshal(dataByte, &workoutData)

	return workoutData
}
