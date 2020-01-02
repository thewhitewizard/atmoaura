package atmoaura

import (
	"errors"
	"net/http"
	"fmt"
	"encoding/json"
)

 
const (
	DefaultHost = "http://api.atmo-aura.fr"
)

var (
	UnauthorizedErr = errors.New("request error: Token unauthorized")
	NoResourceFoundErr = errors.New("request error: no resource found")
)

type Atmoaura struct {
	Token  string
	Host string
}


func NewClient(token string) *Atmoaura { 
	return &Atmoaura{Token: token, Host: DefaultHost} 
}

func (a *Atmoaura) WithHost(h string) {
	a.Host = h
}

func  (a *Atmoaura) GetCurrentIndex(insee string) (*CurrentData, error) {
	return getCurrent(a.Host,a.Token, insee, "indices")
}

func  (a *Atmoaura) GetCurrentVigilance(insee string) (*CurrentData, error) {
	return getCurrent(a.Host,a.Token, insee, "vigilances")
}
 


func getCurrent(host string, token string, insee string, dataType string) (*CurrentData, error) {

	var (
		result CurrentData
		uri  = fmt.Sprintf("%s/communes/%s/%s?api_token=%s&date=now", host, insee,dataType, token)
	)

	err := makeRequest(uri, &result)
	return &result, err
}

func makeRequest(uri string, dest interface{}) error {


	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		return json.NewDecoder(resp.Body).Decode(dest)
	case 401:
		return UnauthorizedErr
	case 403:
		return UnauthorizedErr
	case 404:
		return NoResourceFoundErr
	default:
		return fmt.Errorf(
			"request error occured: recieved http status code (%d)",
			resp.StatusCode)
	}
}