package runkeeper

import "time"

/*
{
	"size": 40,
	"items": [
		{
			"type": "Running",
			"start_time": "Tue, 1 Mar 2011 07:00:00",
			"total_distance": 70,
			"duration": 10,
			"source": "RunKeeper",
			"entry_mode": "API",
			"has_map": "true",
			"uri": "/activities/40"
		},
		{
		"type": "Running",
		"start_time": "Thu, 3 Mar 2011 07:00:00",
		"total_distance": 70,
		"duration": 10,
		"source": "RunKeeper",
		"entry_mode": "Web",
		"has_map": "true",
		"uri": "/activities/39"
		},
		{
		"type": "Running",
		"startTime": "Sat, 5 Mar 2011 11:00:00",
		"total_distance": 70,
		"duration": 10,
		"source": "RunKeeper",
		"entry_mode": "API",
		"has_map": "true",
		"uri": "/activities/38"
		},
		{
		"type": "Running",
		"startTime": "Mon, 7 Mar 2011 07:00:00",
		"total_distance": 70,
		"duration": 10,
		"source": "RunKeeper",
		"entry_mode": "API",
		"has_map": "false",
		"uri": "/activities/37"
		},
		⋮
	],
	"previous": "https://api.runkeeper.com/user/1234567890/activities?page=2"
}
*/

type FitnessActivityFeed struct {
	Size     int64             `json:"size"`
	Items    []FitnessActivity `json:"items"`
	Previous string            `json:"previous"`
}

/*
	{
		"type": "Running",
		"start_time": "Tue, 1 Mar 2011 07:00:00",
		"total_distance": 70,
		"duration": 10,
		"source": "RunKeeper",
		"entry_mode": "API",
		"has_map": "true",
		"uri": "/activities/40"
	}
*/
type FitnessActivity struct {
	Type          string  `json:"type"`
	StartTime     Time    `json:"start_time"`
	TotalDistance float64 `json:"total_distance"`
	Duration      float64 `json:"duration"`
	Source        string  `json:"source"`
	HasMap        bool    `json:"has_map"`
	HasPath       bool    `json:"has_path"`
	EntryMode     string  `json:"entry_mode"`
	Uri           string  `json:"uri"`

	// Details
	Climb         float64 `json:"climb"`
	Comment       string  `json:"comments"`       // "comments" : "/fitnessActivities/318671963/comments",
	UserID        int64   `json:"userID"`         // "userID" : 24207205,
	IsLive        bool    `json:"is_live"`        // "is_live" : false,
	Equipment     string  `json:"equipment"`      // "equipment" : "None",
	TotalCalories float64 `json:"total_calories"` // "total_calories" : 22,

	Share    string `json:"share"`     // "share" : "Everyone",
	ShareMap string `json:"share_map"` // "share_map" : "Friends",

	Distance []Distance `json:"distance"` // "distance" : [ { "distance" : 0, "timestamp" : 0 }, ... ]
	Path     []Path     `json:"path"`
}

type Distance struct {
	Distance  float64 `json:"distance"`  // "distance" : 0,
	Timestamp float64 `json:"timestamp"` // : 0
}

/*
   {
      "altitude" : 37,
      "longitude" : 121.371254,
      "type" : "gps",
      "timestamp" : 3.629,
      "latitude" : 24.942796
   },
*/
type Path struct {
	Altitude  float64 `json:"altitude"`
	Longitude float64 `json:"longitude"` // 121.37
	Type      string  `json:"type"`      // gps
	Latitude  float64 `json:"latitude"`
	Timestamp float64 `json:"timestamp"`
}

type Time time.Time

// Unmarshal "Tue, 1 Mar 2011 07:00:00"
func (self *Time) UnmarshalJSON(data []byte) (err error) {
	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		loc, _ := time.LoadLocation("Local")
		t, err := time.ParseInLocation("Mon, _2 Jan 2006 15:04:05", string(data[1:len(data)-1]), loc)
		if err != nil {
			return err
		}
		*self = Time(t)
	}
	return nil
}

/*
 */

/*
func (self *MetricValue) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		var val int64
		if _, err := fmt.Sscanf(string(data), "\"%d\"", &val); err != nil {
			return err
		}
		*self = MetricValue(val)
	} else {
		val, err := strconv.Atoi(string(data))
		if err != nil {
			return err
		}
		*self = MetricValue(val)
	}
	return nil
}
*/
/*
type Duration float64
func ParseDurationInSeconds(duration string) (Duration, error) {
	var hours, minutes int
	var seconds float64
	if _, err := fmt.Sscanf(duration, "\"%02d:%02d:%f\"", &hours, &minutes, &seconds); err != nil {
		return 0, err
	}
	return Duration(float64(hours)*60*60 + float64(minutes)*60 + seconds), nil
}
func (self *Duration) UnmarshalText(data []byte) (err error) {
	// Fractional seconds are handled implicitly by Parse.
	*self, err = ParseDurationInSeconds(string(data))
	return err
}
func (self *Duration) UnmarshalJSON(data []byte) (err error) {
	// Fractional seconds are handled implicitly by Parse.
	*self, err = ParseDurationInSeconds(string(data))
	return err
}
*/
