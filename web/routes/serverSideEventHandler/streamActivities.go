package serversideeventhandler

import (
	"encoding/json"
	"fitnessApp/api"
	csvwriter "fitnessApp/csvWriter"
	"fmt"
	"net/http"
	"sync"
)

func StreamAvtivities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")

	activietyPage := api.GetLatestActivities()
	csvwriter.WriteMetaData(activietyPage, "./web/public/runningApp/static/media/metaData.990b13c34a76fe3d28a0.csv")

	var wg sync.WaitGroup

	for _, activitesMetadata := range activietyPage {

		metaDataByte, _ := json.Marshal(activitesMetadata)

		fmt.Fprintf(w, "event: activity\ndata: %s\n\n", metaDataByte)

		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}

		go func(id int) {
			wg.Add(1)
			workout := api.GetWorkoutData(id)

			type JsonEntries struct {
				Entries int `json:"entries"`
				ID      int `json:"id"`
			}

			jsonEntries := JsonEntries{}
			jsonEntries.Entries = len(workout.Latlng.Get)
			jsonEntries.ID = id
			byteEntries, _ := json.Marshal(jsonEntries)
			fmt.Fprintf(w, "event: workoutData\ndata: %s\n\n", (byteEntries))
			csvwriter.WriteWorkoutData(id, workout, "./web/public/runningApp/static/media/workoutData.94676a2c4beb32eed05c.csv")
			defer wg.Done()
		}(activitesMetadata.Id)

	}
	wg.Wait()
	endOfStreamEvent := "event: end-of-stream\ndata: END-OF-STREAM\n\n"
	endOfStreamBytes := []byte(endOfStreamEvent)
	_, err := w.Write(endOfStreamBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
