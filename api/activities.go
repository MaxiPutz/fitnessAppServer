package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strava/key"
	"strava/structs"
	"strconv"
)

func GetAllActivities() []structs.MetaData {
	allActivities := []structs.MetaData{}

	tmp := GetLatestActivities()

	pageNumber := 2
	for len(tmp) != 0 {
		allActivities = append(allActivities, tmp...)
		tmp = getActivities(pageNumber)

		pageNumber++
	}

	return allActivities
}

func GetAllActivitiesChan() <-chan []structs.MetaData {
	allActivities := make(chan []structs.MetaData)

	go func() {
		defer close(allActivities)

		tmp := GetLatestActivities()

		pageNumber := 2
		for len(tmp) != 0 {
			allActivities <- tmp
			tmp = getActivities(pageNumber)

			pageNumber++
		}

	}()

	return allActivities
}

func GetLatestActivities() []structs.MetaData {
	return getActivities(1)
}

func getActivities(pageIndex int) []structs.MetaData {
	stravaUrl := stravaAPIUrl + "/athlete/activities?unit_system=metric&page=" + strconv.Itoa(pageIndex)
	fmt.Println(stravaUrl)
	req, err := http.NewRequest("GET", stravaUrl, nil)
	if err != nil {
		fmt.Println(err)
		return []structs.MetaData{}
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + key.ResponseKey().AccessToken},
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []structs.MetaData{}
	}
	defer res.Body.Close()

	data := res.Body
	dataByte, _ := ioutil.ReadAll(data)

	var metadata []structs.MetaData

	json.Unmarshal(dataByte, &metadata)

	for i, _ := range metadata {
		metadata[i].Page = pageIndex
	}

	return metadata
}
