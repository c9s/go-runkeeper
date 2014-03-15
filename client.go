package runkeeper

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
)

const (
	ContentTypeBackgroundActivity       = "application/vnd.com.runkeeper.BackgroundActivity+json"
	ContentTypeBackgroundActivitySet    = "application/vnd.com.runkeeper.BackgroundActivitySet+json"
	ContentTypeComment                  = "application/vnd.com.runkeeper.Comment+json"
	ContentTypeDiabetesMeasurementSet   = "application/vnd.com.runkeeper.DiabetesMeasurementSet+json"
	ContentTypeFitnessActivity          = "application/vnd.com.runkeeper.FitnessActivity+json"
	ContentTypeFitnessActivityFeed      = "application/vnd.com.runkeeper.FitnessActivityFeed+json"
	ContentTypeGeneralMeasurementSet    = "application/vnd.com.runkeeper.GeneralMeasurementSet+json"
	ContentTypeNutritionSet             = "application/vnd.com.runkeeper.NutritionSet+json"
	ContentTypeProfile                  = "application/vnd.com.runkeeper.Profile+json"
	ContentTypeSleepSet                 = "application/vnd.com.runkeeper.SleepSet+json"
	ContentTypeStrengthTrainingActivity = "application/vnd.com.runkeeper.StrengthTrainingActivity+json"
	ContentTypeUser                     = "application/vnd.com.runkeeper.User+json"
	ContentTypeWeightSet                = "application/vnd.com.runkeeper.WeightSet+json"
)
const (
	ContentTypeNewSleep = "application/vnd.com.runkeeper.NewSleep+json"
)

const baseUrl = "https://api.runkeeper.com"

type Client struct {
	AccessToken string
	CookieJar   *cookiejar.Jar
	*http.Client
}

func NewClient(accessToken string) *Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	return &Client{accessToken, jar, &http.Client{Jar: jar}}
}

/**
result should be a struct pointer
*/
func parseJsonResponse(resp *http.Response, result interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, result)
}

func (self *Client) createBaseRequest(method string, url string, acceptContentType string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, baseUrl+url, body)
	req.Header.Add("Accept", acceptContentType)
	req.Header.Add("Authorization", "Bearer "+self.AccessToken)
	if err != nil {
		return nil, err
	}
	return req, nil
}

type Params map[string]interface{}

func (self *Client) GetRequestParams(userParams *Params) url.Values {
	params := url.Values{}
	if userParams != nil {
		for key, val := range *userParams {
			switch t := val.(type) {
			case int, int8, int16, int32, int64:
				params.Set(key, strconv.Itoa(t.(int)))
			case string:
				params.Set(key, t)
			case []byte:
				params.Set(key, string(t))
			default:
				params.Set(key, t.(string))
			}
		}
	}
	return params
}

/**
Valid params:
	page: (int)
	pageSize: (int)
*/
func (self *Client) GetFitnessActivityFeed(userParams *Params) (*FitnessActivityFeed, error) {
	params = self.GetRequestParams(userParams)
	req, err := self.createBaseRequest("GET", "/fitnessActivities?"+params.Encode(), ContentTypeFitnessActivityFeed, nil)
	if err != nil {
		return nil, err
	}

	resp, err := self.Do(req)
	if err != nil {
		return nil, err
	}
	var activities = FitnessActivityFeed{}
	defer resp.Body.Close()
	if err := parseJsonResponse(resp, &activities); err != nil {
		return nil, err
	}
	return &activities, nil
}
