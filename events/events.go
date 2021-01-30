package events

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/nikunicke/ftapi"
)

type Event struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Location       string    `json:"location"`
	Kind           string    `json:"kind"`
	MaxPeople      int       `json:"max_people"`
	NbrSubscribers int       `json:"nbr_subscribers"`
	BeginAt        time.Time `json:"begin_at"`
	EndAt          time.Time `json:"end_at"`
	CampusIds      []int     `json:"campus_ids"`
	CursusIds      []int     `json:"cursus_ids"`
	Themes         []struct {
		CreatedAt time.Time `json:"created_at"`
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"themes"`
	Waitlist                  interface{} `json:"waitlist"`
	ProhibitionOfCancellation int         `json:"prohibition_of_cancellation"`
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt                 time.Time   `json:"updated_at"`
}

// EventsService communicates with the 42 API event endpoint
type EventsService struct {
	s *ftapi.FtClient
}

// Service creates an *EventsService
func Service(s *ftapi.FtClient) *EventsService {
	return &EventsService{s: s}
}

type ListEventsResponse struct {
	Events        []*Event
	Size          int
	NextPageToken string
}

// EventsGetCall method ID: ftapi.events.get:
type EventsGetCall struct {
	s         *ftapi.FtClient
	id        string
	urlParams map[string]string
	header    http.Header
}

// Get : Gets an event specified by an ID
func (s *EventsService) Get(ID string) *EventsGetCall {
	return &EventsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes the "ftapi.events.get" call. Exactly one of *Event or error will be non-nil.
func (c *EventsGetCall) Do() (*Event, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.intra.42.fr/v2/events/"+c.id, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	ret := &Event{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

type EventsListCall struct {
	s         *ftapi.FtClient
	urlParams map[string][]string
	header    http.Header
}

func (s *EventsService) List() *EventsListCall {
	return &EventsListCall{s: s.s, urlParams: make(map[string][]string)}
}

func (c *EventsListCall) Q(key string, values []string) *EventsListCall {
	c.urlParams[key] = values
	return c
}

func (c *EventsListCall) Do() (*ListEventsResponse, error) {
	urls := "https://api.intra.42.fr/v2/events" + "?" + url.Values(c.urlParams).Encode()
	fmt.Println(urls)
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	ret := &ListEventsResponse{}
	ret.Events = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Events); err != nil {
		return nil, err
	}
	ret.Size = len(ret.Events)
	return ret, nil
}
