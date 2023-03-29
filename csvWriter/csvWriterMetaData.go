package csvwriter

import (
	"fitnessApp/fn"
	"fitnessApp/structs"
	"os"
	"strconv"
	"strings"
)

func WriteMetaData(metadataList []structs.MetaData, path string) {

	headerString := "id,sport,startTime,averagePower,averageSpeed,averageHearRate,totalTime,totalDistance"
	head := strings.Split(headerString, ",")

	writeHeader := false
	if _, err := os.Stat(path); os.IsNotExist(err) {
		writeHeader = true
	}

	f, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)

	defer f.Close()
	if writeHeader {
		f.WriteString(headerString)
	}

	for _, meta := range metadataList {
		data := fn.Map(head, func(ele string) string {
			if "id" == ele {
				return strconv.Itoa(meta.Id)
			} else if "startTime" == ele {
				return meta.StartTime.Format("2006-01-02T15:04:05Z")
			} else if "sport" == ele {
				return meta.Sport
			} else if "averagePower" == ele {
				return strconv.Itoa(int(meta.AveragePower))
			} else if "averageSpeed" == ele {
				speed := meterPerSecountToKmPerHour(meta.AverageSpeed)
				return strconv.FormatFloat(speed, 'f', 1, 64)
			} else if "averageHearRate" == ele {
				return strconv.FormatFloat(meta.AverageHearRate, 'f', 1, 64)
			} else if "totalTime" == ele {
				return strconv.Itoa(meta.TotalTime)
			} else if "totalDistance" == ele {
				return strconv.FormatFloat(meta.TotalDistance/1000, 'f', 1, 64)
			} else {
				return ""
			}
		})

		dataString := strings.Join(data, ",")

		f.WriteString("\n" + dataString)

	}

}
