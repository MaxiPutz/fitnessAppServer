package structs

type Data struct {
	Get []float64 `json:"data"`
}
type LatLng struct {
	Get [][]float64 `json:"data"`
}

type WorkoutData struct {
	Watts     Data   `json:"watts"`
	Heartrate Data   `json:"heartrate"`
	Speed     Data   `json:"velocity_smooth"`
	Latlng    LatLng `json:"latlng"`
	TimerTime Data   `json:"time"`
}
