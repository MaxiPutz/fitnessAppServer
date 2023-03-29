package gpx

import (
	"encoding/xml"
	"fitnessApp/fn"
	"fitnessApp/structs"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func RunGpx() {
	res, _ := os.ReadDir("export/activities_gpx")

	res2 := fn.Filter(res, func(ele fs.DirEntry) bool {
		return strings.Contains(ele.Name(), ".gpx")
	})

	for _, ele := range res2 {
		xml1, _ := os.Open("export/activities_gpx/" + ele.Name())
		// buf := bytes.NewBuffer(nil)
		// io.Copy(buf, xml1)
		defer xml1.Close()

		var gpx structs.GPX
		xml.NewDecoder(xml1).Decode(&gpx)

		// fmt.Println(buf.String())
		fmt.Println(gpx.Version)
	}
}
