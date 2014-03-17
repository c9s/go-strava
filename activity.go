package strava

import (
	"strconv"
	"strings"
	"time"
)

/*
   "id": 8529483,
    "resource_state": 2,
    "external_id": "2013-08-23-17-04-12.fit",
    "upload_id": 84130503,
    "athlete": {
      "id": 227615,
      "resource_state": 1
    },
    "name": "08/23/2013 Oakland, CA",
    "distance": 32486.1,
    "moving_time": 5241,
    "elapsed_time": 5427,
    "total_elevation_gain": 566.0,
    "type": "Ride",
    "start_date": "2013-08-24T00:04:12Z",
    "start_date_local": "2013-08-23T17:04:12Z",
    "timezone": "(GMT-08:00) America/Los_Angeles",
    "start_latlng": [
      37.793551,
      -122.2686
    ],
    "end_latlng": [
      37.792836,
      -122.268287
    ],
    "location_city": "Oakland",
    "location_state": "CA",
    "location_country": "United States",
    "start_latitude": 37.793551,
    "start_longitude": -122.2686,
    "achievement_count": 8,
    "kudos_count": 0,
    "comment_count": 0,
    "athlete_count": 1,
    "photo_count": 0,
    "map": {
      "id": "a77175935",
      "summary_polyline": "cetewLja@zYcGbe@",
      "resource_state": 2
    },
    "trainer": false,
    "commute": false,
    "manual": false,
    "private": false,
    "flagged": false,
    "average_speed": 3.4,
    "max_speed": 4.514,
    "average_watts": 163.6,
    "kilojoules": 857.6,
    "average_heartrate": 138.8,
    "max_heartrate": 179.0
*/
type Activity struct {
	Id             int64   `json:"id"`
	ResourceState  int64   `json:"resource_state"`
	UploadId       int64   `json:"upload_id"`
	Name           string  `json:"name"`
	Distance       float64 `json:"distance"`
	MovingTime     float64 `json:"moving_time"`
	ElapsedTime    float64 `json:"elapsed_time"`
	Type           string  `json:"type"` // "Ride"
	TimeZone       string  `json:"timezone"`
	StartDate      string  `json:"start_date"`
	StartDateLocal string  `json:"start_date_local"`

	LocationCity    string `json:"location_city"`
	LocationState   string `json:"location_state"`
	LocationCountry string `json:"location_country"`

	StartLatitude  float64 `json:"start_latitude"`  // : 37.793551,
	StartLongitude float64 `json:"start_longitude"` // : -122.2686,
	AverageSpeed   float64 `json:"average_speed"`   // : 3.4,
	MaxSpeed       float64 `json:"max_speed"`       // : 4.514,
	AverageWatts   float64 `json:"average_watts"`   // : 163.6,
	MaxHeartRate   float64 `json:"max_heartrate"`   // 179.0
}

func (a *Activity) GetTimeZone() *time.Location {
	args := strings.SplitN(a.TimeZone, " ", 2)
	tz, _ := time.LoadLocation(args[1])
	return tz
}

func (a *Activity) GetStartDate() time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05Z", a.StartDateLocal)
	return t
	// a.GetTimeZone()
}

/*
type Time time.Time
func (s *Time) UnmarshalJSON(data []byte) {

}
*/

func (a *Activity) UUID() string {
	str := strconv.Itoa(int(a.Id))
	return str
}

type Activities []Activity
