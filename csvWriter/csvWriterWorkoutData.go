package csvwriter

import (
	"fitnessApp/fn"
	"fitnessApp/structs"
	"os"
	"strconv"
	"strings"
)

func WriteWorkoutData(id int, workoutData structs.WorkoutData, path string) {
	headerString := "id,lat,long,timerTime,power,speed,heartRate"
	header := strings.Split(headerString, ",")

	writeHeader := false
	if _, err := os.Stat(path); os.IsNotExist(err) {
		writeHeader = true
	}

	f, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	defer f.Close()

	if writeHeader {
		f.WriteString(headerString)
	}

	heartrate := workoutData.Heartrate.Get
	lat := fn.Map[[]float64, float64](workoutData.Latlng.Get,
		func(ele []float64) float64 {
			return ele[0]
		})
	long := fn.Map[[]float64, float64](workoutData.Latlng.Get,
		func(ele []float64) float64 {
			return ele[1]
		})

	speed := workoutData.Speed.Get
	watts := workoutData.Watts.Get
	timerTime := workoutData.TimerTime.Get

	arr := make([]string, len(lat))

	for i := 0; i < len(arr); i++ {
		data := fn.Map[string, string](header,
			func(ele string) string {
				if ele == "id" {
					return strconv.Itoa(id)
				} else if ele == "lat" {
					return strconv.FormatFloat(lat[i], 'f', 14, 64)
				} else if ele == "long" {
					return strconv.FormatFloat(long[i], 'f', 14, 64)
				} else if ele == "timerTime" {
					if len(timerTime) != len(lat) {
						return ""
					}
					return strconv.FormatFloat(timerTime[i], 'f', 1, 64)
				} else if ele == "power" {
					if len(watts) != len(lat) {
						return ""
					}
					return strconv.FormatFloat(watts[i], 'f', 1, 64)
				} else if ele == "speed" {
					if len(speed) != len(lat) {
						return ""
					}
					kmH := meterPerSecountToKmPerHour(speed[i])
					return strconv.FormatFloat(kmH, 'f', 1, 64)
				} else if ele == "heartRate" {
					if len(heartrate) != len(lat) {
						return ""
					}
					return strconv.FormatFloat(heartrate[i], 'f', 1, 64)
				}
				return ""
			})

		arr[i] = strings.Join(data, ",")

		f.WriteString("\n" + arr[i])
	}

}
