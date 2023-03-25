package structs

import "time"

type ResourceState struct {
	Athlete            AthleteStruct `json:"athlete"`
	Id                 int           `json:"id"`
	ResourceState      int           `json:"resource_state"`
	Name               string        `json:"name"`
	Distance           float64       `json:"distance"`
	MovingTime         int           `json:"moving_time"`
	ElapsedTime        int           `json:"elapsed_time"`
	TotalElevationGain float64       `json:"total_elevation_gain"`
	Type               string        `json:"type"`
	SportType          int           `json:"sport_type"`
	WorkoutType        int           `json:"workout_type"`
	StartDate          time.Time     `json:"start_date"`
	StartDateLocal     time.Time     `json:"start_date_local"`
	Timezone           string        `json:"timezone"`
	UtcOffset          float64       `json:"utc_offset"`
	LocationCity       string        `json:"location_city"`
	LocationState      string        `json:"location_state"`
	LocationCountry    string        `json:"location_country"`
	AchievementCount   int           `json:"achievement_count"`
	KudosCount         int           `json:"kudos_count"`
	CommentCount       int           `json:"comment_count"`
	AthleteCount       int           `json:"athlete_count"`
	PhotoCount         int           `json:"photo_count"`
	Map                Map           `json:"map"`
}

type AthleteStruct struct {
	Id int `json:"id"`
}

type Map struct {
	Id                         string    `json:"id"`
	SummaryPolyline            string    `json:"summary_polyline"`
	ResourceState              int       `json:"resource_state"`
	Trainer                    bool      `json:"trainer"`
	Commute                    bool      `json:"commute"`
	Manual                     bool      `json:"manual"`
	Private                    bool      `json:"private"`
	Visibility                 string    `json:"visibility"`
	Flagged                    bool      `json:"flagged"`
	GearId                     string    `json:"gear_id"`
	StartLatLng                []float64 `json:"start_latlng"`
	EndLatLng                  []float64 `json:"end_latlng"`
	AverageSpeed               float64   `json:"average_speed"`
	MaxSpeed                   float64   `json:"max_speed"`
	AverageCadence             float64   `json:"average_cadence"`
	AverageWatts               float64   `json:"average_watts"`
	MaxWatts                   int       `json:"max_watts"`
	WeightedAverageWatts       float64   `json:"weighted_average_watts"`
	Kilojoules                 float64   `json:"kilojoules"`
	DeviceWatts                bool      `json:"device_watts"`
	HasHeartrate               bool      `json:"has_heartrate"`
	AverageHeartrate           float64   `json:"average_heartrate"`
	MaxHeartrate               int       `json:"max_heartrate"`
	HeartrateOptOut            bool      `json:"heartrate_opt_out"`
	DisplayHideHeartrateOption bool      `json:"display_hide_heartrate_option"`
	ElevHigh                   float64   `json:"elev_high"`
	ElevLow                    float64   `json:"elev_low"`
	UploadId                   int       `json:"upload_id"`
	UploadIdStr                string    `json:"upload_id_str"`
	ExternalId                 string    `json:"external_id"`
	FromAcceptedTag            bool      `json:"fromAcceptedTag`
}
