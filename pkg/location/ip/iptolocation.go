package ip

import (
	"fmt"
	"sso/config/errors"
	"sso/config/errors/status"
	"sso/pkg/datatype/convert"
	"sso/pkg/http"
	"strings"
)

// Response 지역 데이터 결과 값
type Response struct {
	IP         string `json:"ip"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
}

// GetLocation 아이피를 이용한 지역 정보 조회
func GetLocation(ipAddress string) (*Response, error) {
	// 1차 조회
	resp, err := getLocationIpApi(ipAddress)
	if err == nil {
		return resp, nil
	}
	// 2차 조회
	resp, err = getLocationIpInfo(ipAddress)
	if err == nil {
		return resp, nil
	}
	// 3차 조회
	resp, err = getLocationGeoLocationDB(ipAddress)
	if err == nil {
		return resp, nil
	}
	return nil, errors.New(status.LocationRetrievalError)
}

// getLocationIpApi 지역 정보 조회
func getLocationIpApi(ipAddress string) (*Response, error) {
	/*
		https://ipapi.co
		하루 1000건 제한
	*/
	url := fmt.Sprintf("https://ipapi.co/%s/json/", ipAddress)
	resp, err := http.Request("GET", url, nil, &http.Header{
		Key:   "User-Agent",
		Value: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36",
	})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(status.NotReceivedLocationInfo)
	}
	location, _ := convert.StringToMap(&resp.Body)
	if (location["country_code"] == nil) || (location["postal"] == nil) || (location["latitude"] == nil) || (location["longitude"] == nil) {
		return nil, errors.New(status.NotFoundLocationData)
	}
	return &Response{
		IP:         ipAddress,
		Country:    location["country_code"].(string),
		PostalCode: location["postal"].(string),
		Latitude:   fmt.Sprintf("%f", location["latitude"].(float64)),
		Longitude:  fmt.Sprintf("%f", location["longitude"].(float64)),
	}, nil
}

// getLocationIpInfo 지역 정보 조회
func getLocationIpInfo(ipAddress string) (*Response, error) {
	/*
		http://ipinfo.io
		한 달 50000건 제한
	*/
	url := fmt.Sprintf("https://ipinfo.io/%s/json", ipAddress)
	resp, err := http.Request("GET", url, nil, &http.Header{
		Key:   "User-Agent",
		Value: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36",
	})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(status.NotReceivedLocationInfo)
	}
	location, _ := convert.StringToMap(&resp.Body)
	// 위치 정보 조회
	locations := strings.Split(location["loc"].(string), ",")
	if len(locations) < 2 {
		return nil, errors.New(status.InvalidInformationError)
	}
	if (location["country"] == nil) || (location["postal"] == nil) {
		return nil, errors.New(status.NotFoundLocationData)
	}
	return &Response{
		IP:         ipAddress,
		Country:    location["country"].(string),
		PostalCode: location["postal"].(string),
		Latitude:   locations[0],
		Longitude:  locations[1],
	}, nil
}

// getLocationGeoLocationDB 지역 정보 조회
func getLocationGeoLocationDB(ipAddress string) (*Response, error) {
	/*
		https://geolocation-db.com
	*/
	url := "https://geolocation-db.com/json/" + ipAddress
	resp, err := http.Request("GET", url, nil, &http.Header{
		Key:   "User-Agent",
		Value: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36",
	})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(status.NotReceivedLocationInfo)
	}
	location, _ := convert.StringToMap(&resp.Body)
	if (location["country_code"] == nil) || (location["postal"] == nil) || (location["latitude"] == nil) || (location["longitude"] == nil) {
		return nil, errors.New(status.NotFoundLocationData)
	}
	return &Response{
		IP:         ipAddress,
		Country:    location["country_code"].(string),
		PostalCode: location["postal"].(string),
		Latitude:   location["latitude"].(string),
		Longitude:  location["longitude"].(string),
	}, nil
}
