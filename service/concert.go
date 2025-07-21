package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Event struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetConcerts(artist string, lat, lon float32) ([]Event, error) {
	apiKey := os.Getenv("TICKETMASTER_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("TICKETMASTER_API_KEY not set")
	}

	endpoint := "https://app.ticketmaster.com/discovery/v2/events.json"
	params := url.Values{}
	params.Set("apikey", apiKey)
	params.Set("keyword", artist)
	params.Set("latlong", fmt.Sprintf("%.4f,%.4f", lat, lon))
	params.Set("radius", "50")

	resp, err := http.Get(endpoint + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}
	json.Unmarshal(body, &data)

	events := []Event{}
	if embedded, ok := data["_embedded"].(map[string]interface{}); ok {
		for _, ev := range embedded["events"].([]interface{}) {
			e := ev.(map[string]interface{})
			events = append(events, Event{
				Name: e["name"].(string),
				URL:  e["url"].(string),
			})
		}
	}
	return events, nil
}
