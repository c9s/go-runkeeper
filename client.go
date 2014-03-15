package nikeplus

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
)

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

var baseUrl string = "https://api.nike.com"

type ErrorResult struct {
	Result       string `json:"result"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

/**
result should be a struct pointer
*/
func parseJsonResponse(resp *http.Response, result interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if bytes.Contains(body, []byte("errorCode")) {
		result := ErrorResult{}
		if err := json.Unmarshal(body, &result); err != nil {
			return err
		}
		return errors.New(result.ErrorMessage)
	}
	return json.Unmarshal(body, result)
}

/*
Login to Nike+ developer and get the access token

@param string email
@param string password

$this->requestPost('https://developer.nike.com/login',
	array( 'email' => $email, 'password' => $password ));
*/
func (self *Client) Login(email string, password string) error {
	params := url.Values{}
	params.Set("email", email)
	params.Set("password", password)
	params.Set("continue_url", "/categories")

	buf := bytes.NewBuffer([]byte(params.Encode()))
	resp, err := self.Post("https://developer.nike.com/login", "application/x-www-form-urlencoded", buf)
	if err != nil {
		return err
	}
	if resp.Request.URL != nil {
		redirectUrlStr := resp.Request.URL.String()
		url, err := url.Parse(redirectUrlStr)
		if err != nil {
			return err
		}
		if strings.Contains(url.RawQuery, "error=") {
			return errors.New("Login return: " + url.RawQuery)
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func (self *Client) AskAccessToken() (string, error) {
	req, err := http.NewRequest("POST", "https://developer.nike.com/get_auth_token", nil)
	req.Header.Add("Accept", "application/json")
	resp, err := self.Do(req)
	if err != nil {
		return "", err
	}
	var retval map[string]interface{}
	if err := parseJsonResponse(resp, &retval); err != nil {
		return "", err
	}
	if val, ok := retval["auth_token"].(string); ok {
		self.AccessToken = val
		return val, nil
	}
	return "", errors.New("Can't get access token")
}

/*
curl -H "Accept: application/json" "https://api.nike.com/me/sport/activities/c8f65c19-6fe6-43fe-9393-90f52246e111?access_token=dee6ce5e936434ca7275d678d4104f30"
*/
func (self *Client) GetActivityDetails(activityId string) (*Activity, error) {
	url := fmt.Sprintf("%s/me/sport/activities/%s?access_token=%s", baseUrl, activityId, self.AccessToken)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := self.Do(req)
	if err != nil {
		return nil, err
	}
	var activity = Activity{}
	defer resp.Body.Close()
	if err := parseJsonResponse(resp, &activity); err != nil {
		return nil, err
	}
	return &activity, nil
}

type Params map[string]interface{}

func (self *Client) GetRequestParams(userParams *Params) url.Values {
	params := url.Values{"access_token": {self.AccessToken}}
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

func (self *Client) getActivitiesFromRequest(req *http.Request) (*Activities, error) {
	resp, err := self.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var activities = Activities{}
	if err := parseJsonResponse(resp, &activities); err != nil {
		return nil, err
	}
	return &activities, nil
}

func (self *Client) GetActivities(userParams *Params) (*Activities, error) {
	params := self.GetRequestParams(userParams)
	req, err := http.NewRequest("GET", baseUrl+"/me/sport/activities?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	return self.getActivitiesFromRequest(req)
}

func (self *Client) GetActivitiesByType(actType string, userParams *Params) (*Activities, error) {
	params := self.GetRequestParams(userParams)
	req, err := http.NewRequest("GET", baseUrl+"/me/sport/activities/"+actType+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	return self.getActivitiesFromRequest(req)
}
