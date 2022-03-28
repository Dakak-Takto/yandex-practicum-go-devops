package sender

import (
	"fmt"
	"log"
	"net/http"
)

type Sender interface {
	SendGauge(map[string]float64)
	SendCounter(map[string]int64)
}

type sender struct {
	httpClient *http.Client
	baseUrl    string
}

var _ Sender = &sender{}

func NewSender(baseUrl string) *sender {
	s := sender{
		baseUrl: baseUrl,
	}
	s.httpClient = http.DefaultClient
	return &s
}

func (s *sender) SendGauge(gauges map[string]float64) {
	const (
		gauge = "gauge"
	)

	for name, value := range gauges {
		url := fmt.Sprintf("%s/%s/%s/%f", s.baseUrl, gauge, name, value)

		req, err := http.NewRequest(http.MethodPost, url, nil)
		if err != nil {
			log.Println(err)
			continue
		}
		response, err := s.httpClient.Do(req)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(response.StatusCode, url)
	}
}

func (s *sender) SendCounter(gauges map[string]int64) {
	const (
		counter = "counter"
	)

	for name, value := range gauges {
		url := fmt.Sprintf("%s/%s/%s/%d", s.baseUrl, counter, name, value)

		req, err := http.NewRequest(http.MethodPost, url, nil)
		if err != nil {
			log.Println(err)
			break
		}
		response, err := s.httpClient.Do(req)
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(response.StatusCode, url)
	}
}
