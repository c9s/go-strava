package strava

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
)

const baseUrl = "https://www.strava.com/api/v3"

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

func (self *Client) createBaseRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, baseUrl+url, body)
	req.Header.Add("Accept", "application/json")
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
	per_page: (int)
*/
func (self *Client) GetActivities(userParams *Params) (*Activities, error) {
	params := self.GetRequestParams(userParams)
	req, err := self.createBaseRequest("GET", "/athlete/activities"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := self.Do(req)
	if err != nil {
		return nil, err
	}
	var activities = Activities{}
	defer resp.Body.Close()
	if err := parseJsonResponse(resp, &activities); err != nil {
		return nil, err
	}
	return &activities, nil
}
