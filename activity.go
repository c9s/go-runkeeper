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
	Type          string    `json:"type"`
	StartTime     time.Time `json:"start_time"`
	TotalDistance int64     `json:"total_distance"`
	Duration      int64     `json:"duration"`
	Source        string    `json:"source"`
	HasMap        string    `json:"has_map"`
	HasPath       string    `json:"has_path"`
	EntryMode     string    `json:"entry_mode"`
	Uri           string    `json:"uri"`

	HeartRate []int `json:"heart_rate"`

	Client
}

func (self *FitnessActivity) GetDetail() {

}

/*
func (self *TimeZone) UnmarshalJSON(data []byte) (err error) {
	// Fractional seconds are handled implicitly by Parse.
	loc, err := time.LoadLocation(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}
	if loc != nil {
		*self = TimeZone(*loc)
	}
	return nil
}
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
