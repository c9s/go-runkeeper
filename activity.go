package nikeplus

import "time"
import "fmt"
import "strconv"

/*
Activity Response struct
*/
type Activities struct {
	Data   []Activity `json:"data"`
	Paging Paging     `json:"paging"`
	Error  string     `json:"error"`
}

type Paging struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}

func (self *Activities) GetNextPage() string {
	return self.Paging.Next
}

func (self *Activities) GetPrevPage() string {
	return self.Paging.Prev
}

/*
	{
      "activityId": "263c1cde-552f-4c65-a943-7214691ec81e",
      "activityType": "ALL_DAY",
      "startTime": "2014-03-09T16:00:00Z",
      "activityTimeZone": "Asia/Taipei",
      "status": "IN_PROGRESS",
      "deviceType": "FUELBAND2",
      "metricSummary": {
        "calories": 47,
        "fuel": 180,
        "distance": 0.7881873846054077,
        "steps": 1001,
        "duration": "0:56:00.000"
      },
      "tags": [],
      "metrics": []
	}
*/
type Activity struct {
	Id            string        `json:"activityId"`
	Type          string        `json:"activityType"`
	StartTime     time.Time     `json:"startTime"`
	TimeZone      TimeZone      `json:"activityTimeZone"`
	Status        string        `json:"status"`
	DeviceType    string        `json:"deviceType"`
	MetricSummary MetricSummary `json:"metricSummary"`
	Tags          []Tag         `json:"tags"`
	Metrics       []Metric      `json:"metrics"`
}

func (self *Activity) Location() time.Location {
	return time.Location(self.TimeZone)
}

type TimeZone time.Location

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

type Tag struct {
	Type  string `json:"tagType"`
	Value string `json:"tagValue"`
}

/*
   "calories": 173,
   "fuel": 655,
   "distance": 0,
   "steps": 595,
   "duration": "0:00:00.000"
*/
type MetricSummary struct {
	Calories int64    `json:"calories"`
	Fuel     int64    `json:"fuel"`
	Distance float64  `json:"distance"`
	Steps    int64    `json:"steps"`
	Duration Duration `json:"duration"`
}

type Metric struct {
	IntervalMetric int           `json:"intervalMetric"` // 1 by default
	IntervalUnit   string        `json:"intervalUnit"`   // "MIN",
	Type           string        `json:"metricType"`     // "metricType": "STARS", "CALORIES", "STEPS", "FUEL"
	Values         []MetricValue `json:"values"`         // "values": [ "1","2","3","4" , .... ]
}

type MetricValue int64

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

/**
http://www.ostyn.com/standards/scorm/samples/ISOTimeForSCORM.htm
http://support.sas.com/documentation/cdl/en/lrdict/64316/HTML/default/viewer.htm#a003169814.htm
format: hh:mm:ss.ffffff
*/
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
