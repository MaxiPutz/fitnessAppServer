package structs

import "encoding/xml"

type Athlete struct {
	Id            int         `json:"id"`
	Username      string      `json:"username"`
	ResourceState int         `json:"resource_state"`
	Firstname     string      `json:"firstname"`
	Lastname      string      `json:"lastname"`
	Bio           string      `json:"bio"`
	City          string      `json:"city"`
	State         string      `json:"state"`
	Country       string      `json:"country"`
	Sex           string      `json:"sex"`
	Premium       bool        `json:"premium"`
	Summit        bool        `json:"summit"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
	BadgeTypeId   int         `json:"badge_type_id"`
	Weight        float64     `json:"weight"`
	ProfileMedium string      `json:"profile_medium"`
	Profile       string      `json:"profile"`
	Friend        interface{} `json:"friend"`
	Follower      interface{} `json:"follower"`
}

type Activity struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	Distance       float64 `json:"distance"`
	ElapsedTime    int     `json:"elapsed_time"`
	StartDateLocal string  `json:"start_date_local"`
	Type           string  `json:"type"`
}

type GPX struct {
	XMLName xml.Name `xml:"gpx"`
	Version string   `xml:"version,attr"`
	Creator string   `xml:"creator,attr"`
	Tracks  []Track  `xml:"trk"`
}

type Track struct {
	Name   string    `xml:"name"`
	Points []TrackPt `xml:"trkseg>trkpt"`
}

type TrackPt struct {
	Lat  float64 `xml:"lat,attr"`
	Lon  float64 `xml:"lon,attr"`
	Ele  float64 `xml:"ele"`
	Time string  `xml:"time"`
}
