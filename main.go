package main

import "fitnessApp/web"

func main() {

	// login to the page ***************************************

	web.RunWeb()

	// create Meta data and workout data ***********************

	// metadataList := api.GetAllActivities()
	// csvwriter.WriteMetaData(metadataList)

	// var wg sync.WaitGroup
	// for _, ele := range metadataList {
	// 	wg.Add(1)
	// 	go func(id int) {
	// 		fmt.Println(id)
	// 		data := api.GetWorkoutData(id)
	// 		csvwriter.WriteWorkoutData(id, data)
	// 		defer wg.Done()
	// 	}(ele.Id)
	// }
	// wg.Wait()

	// get all Data *******************************************

	// metaData := api.GetAllActivities()

	// for _, data := range metaData {
	// 	fmt.Printf("data: %v\n", data.StartTime)
	// }
	// fmt.Println(len(metaData))

	// test chan

	// chanActi := api.GetAllActivitiesChan()

	// // var wg sync.WaitGroup
	// for activities := range chanActi {
	// 	fmt.Printf("chanActi: %v\n", activities)

	// 	for _, ele := range activities {
	// 		// wg.Add(1)
	// 		func(id int) {

	// 			workourData := api.GetWorkoutData(id)
	// 			fmt.Printf("workourData.Watts.Get[0]: %v\n", workourData.Watts.Get)
	// 			// defer wg.Done()
	// 		}(ele.Id)

	// 	}
	// }
	// // wg.Wait()
}
