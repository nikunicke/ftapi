package ftapi

// This is generated code. DO NOT EDIT!!!
// This is generated code. DO NOT EDIT!!!
// This is generated code. DO NOT EDIT!!!
// This is generated code. DO NOT EDIT!!!
// This is generated code. DO NOT EDIT!!!

// Package ftapi provides access to the 42 Api.
// This is the second version of the 42‘s API.

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

// AccreditationsService handles 42 API Accreditations.
type AccreditationsService struct {
	s *Service
}

// AccreditationsItem is a 42 API type
type AccreditationsItem struct {
	CursusID  int    `json:"cursus_id"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	UserID    int    `json:"user_id"`
	Validated bool   `json:"validated"`
}

// AccreditationsListCall description:
type AccreditationsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListAccreditationsResponse represents a list response.
type ListAccreditationsResponse struct {
	Accreditations []*AccreditationsItem
}

// List returns a AccreditationsListCall which can request data from the 42 API.
func (s *AccreditationsService) List() *AccreditationsListCall {
	return &AccreditationsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Accreditations matching the specified setting.
func (c *AccreditationsListCall) P(key string, values []string) *AccreditationsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a AccreditationsListCall request call. Exactly one of *ListAccreditationsResponse or error will be non-nil.
func (c *AccreditationsListCall) Do() (*ListAccreditationsResponse, error) {
	urls := c.s.baseURL + "/accreditations/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListAccreditationsResponse{}
	ret.Accreditations = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Accreditations); err != nil {
		return nil, err
	}
	return ret, nil
}

// AccreditationsGetCall description:
type AccreditationsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified AccreditationsItem.
func (s *AccreditationsService) Get(ID string) *AccreditationsGetCall {
	return &AccreditationsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a AccreditationsGetCall request call. Exactly one of *AccreditationsItem or error will be non-nil.
func (c *AccreditationsGetCall) Do() (*AccreditationsItem, error) {
	urls := c.s.baseURL + "/accreditations/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &AccreditationsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Accreditations description: Accreditations
func Accreditations(s *Service) *AccreditationsService {
	return &AccreditationsService{s: s}
}

// AchievementsService handles 42 API Achievements.
type AchievementsService struct {
	s *Service
}

// AchievementsItem is a 42 API type
type AchievementsItem struct {
	Achievements []interface{} `json:"achievements"`
	Description  string        `json:"description"`
	ID           int           `json:"id"`
	Image        string        `json:"image"`
	Kind         string        `json:"kind"`
	Name         string        `json:"name"`
	NbrOfSuccess int           `json:"nbr_of_success"`
	Parent       interface{}   `json:"parent"`
	Tier         string        `json:"tier"`
	Title        interface{}   `json:"title"`
	UsersURL     string        `json:"users_url"`
	Visible      bool          `json:"visible"`
}

// AchievementsListCall description:
type AchievementsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListAchievementsResponse represents a list response.
type ListAchievementsResponse struct {
	Achievements []*AchievementsItem
}

// List returns a AchievementsListCall which can request data from the 42 API.
func (s *AchievementsService) List() *AchievementsListCall {
	return &AchievementsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Achievements matching the specified setting.
func (c *AchievementsListCall) P(key string, values []string) *AchievementsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a AchievementsListCall request call. Exactly one of *ListAchievementsResponse or error will be non-nil.
func (c *AchievementsListCall) Do() (*ListAchievementsResponse, error) {
	urls := c.s.baseURL + "/achievements/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListAchievementsResponse{}
	ret.Achievements = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Achievements); err != nil {
		return nil, err
	}
	return ret, nil
}

// AchievementsGetCall description:
type AchievementsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified AchievementsItem.
func (s *AchievementsService) Get(ID string) *AchievementsGetCall {
	return &AchievementsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a AchievementsGetCall request call. Exactly one of *AchievementsItem or error will be non-nil.
func (c *AchievementsGetCall) Do() (*AchievementsItem, error) {
	urls := c.s.baseURL + "/achievements/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &AchievementsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Achievements description: Meta-goals earned by users all along their progression.
func Achievements(s *Service) *AchievementsService {
	return &AchievementsService{s: s}
}

// AchievementsUsersService handles 42 API AchievementsUsers.
type AchievementsUsersService struct {
	s *Service
}

// AchievementsUsersItem is a 42 API type
type AchievementsUsersItem struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
	Login     string    `json:"login"`
	URL       string    `json:"url"`
	UserID    int       `json:"user_id"`
}

// AchievementsUsersListCall description:
type AchievementsUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListAchievementsUsersResponse represents a list response.
type ListAchievementsUsersResponse struct {
	AchievementsUsers []*AchievementsUsersItem
}

// List returns a AchievementsUsersListCall which can request data from the 42 API.
func (s *AchievementsUsersService) List() *AchievementsUsersListCall {
	return &AchievementsUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return AchievementsUsers matching the specified setting.
func (c *AchievementsUsersListCall) P(key string, values []string) *AchievementsUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a AchievementsUsersListCall request call. Exactly one of *ListAchievementsUsersResponse or error will be non-nil.
func (c *AchievementsUsersListCall) Do() (*ListAchievementsUsersResponse, error) {
	urls := c.s.baseURL + "/achievements_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListAchievementsUsersResponse{}
	ret.AchievementsUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).AchievementsUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// AchievementsUsersGetCall description:
type AchievementsUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified AchievementsUsersItem.
func (s *AchievementsUsersService) Get(ID string) *AchievementsUsersGetCall {
	return &AchievementsUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a AchievementsUsersGetCall request call. Exactly one of *AchievementsUsersItem or error will be non-nil.
func (c *AchievementsUsersGetCall) Do() (*AchievementsUsersItem, error) {
	urls := c.s.baseURL + "/achievements_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &AchievementsUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// AchievementsUsers description: Users which earned an achievement
func AchievementsUsers(s *Service) *AchievementsUsersService {
	return &AchievementsUsersService{s: s}
}

// AnnouncementsService handles 42 API Announcements.
type AnnouncementsService struct {
	s *Service
}

// AnnouncementsItem is a 42 API type
type AnnouncementsItem struct {
	Author    string      `json:"author"`
	CreatedAt time.Time   `json:"created_at"`
	ExpireAt  time.Time   `json:"expire_at"`
	ID        int         `json:"id"`
	Kind      string      `json:"kind"`
	Link      interface{} `json:"link"`
	Text      string      `json:"text"`
	Title     string      `json:"title"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// AnnouncementsGetCall description:
type AnnouncementsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified AnnouncementsItem.
func (s *AnnouncementsService) Get(ID string) *AnnouncementsGetCall {
	return &AnnouncementsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a AnnouncementsGetCall request call. Exactly one of *AnnouncementsItem or error will be non-nil.
func (c *AnnouncementsGetCall) Do() (*AnnouncementsItem, error) {
	urls := c.s.baseURL + "/announcements/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &AnnouncementsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Announcements description: An announcement made to users in a cursus on their homepage.
func Announcements(s *Service) *AnnouncementsService {
	return &AnnouncementsService{s: s}
}

// AntiGravUnitsService handles 42 API AntiGravUnits.
type AntiGravUnitsService struct {
	s *Service
}

// AntiGravUnitsItem is a 42 API type
type AntiGravUnitsItem struct {
}

// AntiGravUnitsListCall description:
type AntiGravUnitsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListAntiGravUnitsResponse represents a list response.
type ListAntiGravUnitsResponse struct {
	AntiGravUnits []*AntiGravUnitsItem
}

// List returns a AntiGravUnitsListCall which can request data from the 42 API.
func (s *AntiGravUnitsService) List() *AntiGravUnitsListCall {
	return &AntiGravUnitsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return AntiGravUnits matching the specified setting.
func (c *AntiGravUnitsListCall) P(key string, values []string) *AntiGravUnitsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a AntiGravUnitsListCall request call. Exactly one of *ListAntiGravUnitsResponse or error will be non-nil.
func (c *AntiGravUnitsListCall) Do() (*ListAntiGravUnitsResponse, error) {
	urls := c.s.baseURL + "/anti_grav_units/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListAntiGravUnitsResponse{}
	ret.AntiGravUnits = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).AntiGravUnits); err != nil {
		return nil, err
	}
	return ret, nil
}

// AntiGravUnitsGetCall description:
type AntiGravUnitsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified AntiGravUnitsItem.
func (s *AntiGravUnitsService) Get(ID string) *AntiGravUnitsGetCall {
	return &AntiGravUnitsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a AntiGravUnitsGetCall request call. Exactly one of *AntiGravUnitsItem or error will be non-nil.
func (c *AntiGravUnitsGetCall) Do() (*AntiGravUnitsItem, error) {
	urls := c.s.baseURL + "/anti_grav_units/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &AntiGravUnitsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// AntiGravUnits description:
func AntiGravUnits(s *Service) *AntiGravUnitsService {
	return &AntiGravUnitsService{s: s}
}

// AntiGravUnitsUsersService handles 42 API AntiGravUnitsUsers.
type AntiGravUnitsUsersService struct {
	s *Service
}

// AntiGravUnitsUsersItem is a 42 API type
type AntiGravUnitsUsersItem struct {
}

// AntiGravUnitsUsersListCall description:
type AntiGravUnitsUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListAntiGravUnitsUsersResponse represents a list response.
type ListAntiGravUnitsUsersResponse struct {
	AntiGravUnitsUsers []*AntiGravUnitsUsersItem
}

// List returns a AntiGravUnitsUsersListCall which can request data from the 42 API.
func (s *AntiGravUnitsUsersService) List() *AntiGravUnitsUsersListCall {
	return &AntiGravUnitsUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return AntiGravUnitsUsers matching the specified setting.
func (c *AntiGravUnitsUsersListCall) P(key string, values []string) *AntiGravUnitsUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a AntiGravUnitsUsersListCall request call. Exactly one of *ListAntiGravUnitsUsersResponse or error will be non-nil.
func (c *AntiGravUnitsUsersListCall) Do() (*ListAntiGravUnitsUsersResponse, error) {
	urls := c.s.baseURL + "/anti_grav_units_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListAntiGravUnitsUsersResponse{}
	ret.AntiGravUnitsUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).AntiGravUnitsUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// AntiGravUnitsUsersGetCall description:
type AntiGravUnitsUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified AntiGravUnitsUsersItem.
func (s *AntiGravUnitsUsersService) Get(ID string) *AntiGravUnitsUsersGetCall {
	return &AntiGravUnitsUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a AntiGravUnitsUsersGetCall request call. Exactly one of *AntiGravUnitsUsersItem or error will be non-nil.
func (c *AntiGravUnitsUsersGetCall) Do() (*AntiGravUnitsUsersItem, error) {
	urls := c.s.baseURL + "/anti_grav_units_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &AntiGravUnitsUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// AntiGravUnitsUsers description:
func AntiGravUnitsUsers(s *Service) *AntiGravUnitsUsersService {
	return &AntiGravUnitsUsersService{s: s}
}

// AppsService handles 42 API Apps.
type AppsService struct {
	s *Service
}

// AppsItem is a 42 API type
type AppsItem struct {
	CreatedAt   time.Time   `json:"created_at"`
	Description interface{} `json:"description"`
	ID          int         `json:"id"`
	Image       interface{} `json:"image"`
	Name        string      `json:"name"`
	Owner       struct {
		URL   string `json:"url"`
		ID    int    `json:"id"`
		Login string `json:"login"`
	} `json:"owner"`
	Public    bool `json:"public"`
	RateLimit int  `json:"rate_limit"`
	Roles     []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"roles"`
	Scopes    []interface{} `json:"scopes"`
	UpdatedAt time.Time     `json:"updated_at"`
	Website   interface{}   `json:"website"`
}

// AppsListCall description:
type AppsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListAppsResponse represents a list response.
type ListAppsResponse struct {
	Apps []*AppsItem
}

// List returns a AppsListCall which can request data from the 42 API.
func (s *AppsService) List() *AppsListCall {
	return &AppsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Apps matching the specified setting.
func (c *AppsListCall) P(key string, values []string) *AppsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a AppsListCall request call. Exactly one of *ListAppsResponse or error will be non-nil.
func (c *AppsListCall) Do() (*ListAppsResponse, error) {
	urls := c.s.baseURL + "/apps/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListAppsResponse{}
	ret.Apps = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Apps); err != nil {
		return nil, err
	}
	return ret, nil
}

// AppsGetCall description:
type AppsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified AppsItem.
func (s *AppsService) Get(ID string) *AppsGetCall {
	return &AppsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a AppsGetCall request call. Exactly one of *AppsItem or error will be non-nil.
func (c *AppsGetCall) Do() (*AppsItem, error) {
	urls := c.s.baseURL + "/apps/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &AppsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// Apps description: Applications for the API v2
func Apps(s *Service) *AppsService {
	return &AppsService{s: s}
}

// AttachmentsService handles 42 API Attachments.
type AttachmentsService struct {
	s *Service
}

// AttachmentsItem is a 42 API type
type AttachmentsItem struct {
	BaseID    int       `json:"base_id"`
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
	Language  struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		IDentifier string `json:"identifier"`
	} `json:"language"`
	Name      string `json:"name"`
	PageCount int    `json:"page_count"`
	Pdf       struct {
		Pdf struct {
			URL   interface{} `json:"url"`
			Thumb struct {
				URL interface{} `json:"url"`
			} `json:"thumb"`
		} `json:"pdf"`
	} `json:"pdf"`
	PdfProcessing bool        `json:"pdf_processing"`
	Slug          string      `json:"slug"`
	ThumbURL      interface{} `json:"thumb_url"`
	Type          string      `json:"type"`
	URL           interface{} `json:"url"`
}

// AttachmentsListCall description:
type AttachmentsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListAttachmentsResponse represents a list response.
type ListAttachmentsResponse struct {
	Attachments []*AttachmentsItem
}

// List returns a AttachmentsListCall which can request data from the 42 API.
func (s *AttachmentsService) List() *AttachmentsListCall {
	return &AttachmentsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Attachments matching the specified setting.
func (c *AttachmentsListCall) P(key string, values []string) *AttachmentsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a AttachmentsListCall request call. Exactly one of *ListAttachmentsResponse or error will be non-nil.
func (c *AttachmentsListCall) Do() (*ListAttachmentsResponse, error) {
	urls := c.s.baseURL + "/attachments/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListAttachmentsResponse{}
	ret.Attachments = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Attachments); err != nil {
		return nil, err
	}
	return ret, nil
}

// AttachmentsGetCall description:
type AttachmentsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified AttachmentsItem.
func (s *AttachmentsService) Get(ID string) *AttachmentsGetCall {
	return &AttachmentsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a AttachmentsGetCall request call. Exactly one of *AttachmentsItem or error will be non-nil.
func (c *AttachmentsGetCall) Do() (*AttachmentsItem, error) {
	urls := c.s.baseURL + "/attachments/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &AttachmentsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Attachments description: All data which can be linked, like videos, pdfs, or links.
func Attachments(s *Service) *AttachmentsService {
	return &AttachmentsService{s: s}
}

// BalancesService handles 42 API Balances.
type BalancesService struct {
	s *Service
}

// BalancesItem is a 42 API type
type BalancesItem struct {
	BeginAt string `json:"begin_at"`
	EndAt   string `json:"end_at"`
	ID      int    `json:"id"`
	PoolID  int    `json:"pool_id"`
}

// BalancesListCall description:
type BalancesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListBalancesResponse represents a list response.
type ListBalancesResponse struct {
	Balances []*BalancesItem
}

// List returns a BalancesListCall which can request data from the 42 API.
func (s *BalancesService) List() *BalancesListCall {
	return &BalancesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Balances matching the specified setting.
func (c *BalancesListCall) P(key string, values []string) *BalancesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a BalancesListCall request call. Exactly one of *ListBalancesResponse or error will be non-nil.
func (c *BalancesListCall) Do() (*ListBalancesResponse, error) {
	urls := c.s.baseURL + "/balances/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListBalancesResponse{}
	ret.Balances = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Balances); err != nil {
		return nil, err
	}
	return ret, nil
}

// BalancesGetCall description:
type BalancesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified BalancesItem.
func (s *BalancesService) Get(ID string) *BalancesGetCall {
	return &BalancesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a BalancesGetCall request call. Exactly one of *BalancesItem or error will be non-nil.
func (c *BalancesGetCall) Do() (*BalancesItem, error) {
	urls := c.s.baseURL + "/balances/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &BalancesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// Balances description: The balance of a pool
func Balances(s *Service) *BalancesService {
	return &BalancesService{s: s}
}

// BlocDeadlinesService handles 42 API BlocDeadlines.
type BlocDeadlinesService struct {
	s *Service
}

// BlocDeadlinesItem is a 42 API type
type BlocDeadlinesItem struct {
	BeginAt     time.Time `json:"begin_at"`
	BlocID      int       `json:"bloc_id"`
	CoalitionID int       `json:"coalition_id"`
	CreatedAt   time.Time `json:"created_at"`
	EndAt       time.Time `json:"end_at"`
	ID          int       `json:"id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// BlocDeadlinesListCall description:
type BlocDeadlinesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListBlocDeadlinesResponse represents a list response.
type ListBlocDeadlinesResponse struct {
	BlocDeadlines []*BlocDeadlinesItem
}

// List returns a BlocDeadlinesListCall which can request data from the 42 API.
func (s *BlocDeadlinesService) List() *BlocDeadlinesListCall {
	return &BlocDeadlinesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return BlocDeadlines matching the specified setting.
func (c *BlocDeadlinesListCall) P(key string, values []string) *BlocDeadlinesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a BlocDeadlinesListCall request call. Exactly one of *ListBlocDeadlinesResponse or error will be non-nil.
func (c *BlocDeadlinesListCall) Do() (*ListBlocDeadlinesResponse, error) {
	urls := c.s.baseURL + "/bloc_deadlines/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListBlocDeadlinesResponse{}
	ret.BlocDeadlines = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).BlocDeadlines); err != nil {
		return nil, err
	}
	return ret, nil
}

// BlocDeadlinesGetCall description:
type BlocDeadlinesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified BlocDeadlinesItem.
func (s *BlocDeadlinesService) Get(ID string) *BlocDeadlinesGetCall {
	return &BlocDeadlinesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a BlocDeadlinesGetCall request call. Exactly one of *BlocDeadlinesItem or error will be non-nil.
func (c *BlocDeadlinesGetCall) Do() (*BlocDeadlinesItem, error) {
	urls := c.s.baseURL + "/bloc_deadlines/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &BlocDeadlinesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// BlocDeadlines description: A bloc
func BlocDeadlines(s *Service) *BlocDeadlinesService {
	return &BlocDeadlinesService{s: s}
}

// BlocsService handles 42 API Blocs.
type BlocsService struct {
	s *Service
}

// BlocsItem is a 42 API type
type BlocsItem struct {
	CampusID   int `json:"campus_id"`
	Coalitions []struct {
		ImageURL string `json:"image_url"`
		Color    string `json:"color"`
		Score    int    `json:"score"`
		UserID   int    `json:"user_id"`
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Slug     string `json:"slug"`
	} `json:"coalitions"`
	CreatedAt time.Time `json:"created_at"`
	CursusID  int       `json:"cursus_id"`
	ID        int       `json:"id"`
	SquadSize int       `json:"squad_size"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BlocsListCall description:
type BlocsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListBlocsResponse represents a list response.
type ListBlocsResponse struct {
	Blocs []*BlocsItem
}

// List returns a BlocsListCall which can request data from the 42 API.
func (s *BlocsService) List() *BlocsListCall {
	return &BlocsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Blocs matching the specified setting.
func (c *BlocsListCall) P(key string, values []string) *BlocsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a BlocsListCall request call. Exactly one of *ListBlocsResponse or error will be non-nil.
func (c *BlocsListCall) Do() (*ListBlocsResponse, error) {
	urls := c.s.baseURL + "/blocs/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListBlocsResponse{}
	ret.Blocs = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Blocs); err != nil {
		return nil, err
	}
	return ret, nil
}

// BlocsGetCall description:
type BlocsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified BlocsItem.
func (s *BlocsService) Get(ID string) *BlocsGetCall {
	return &BlocsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a BlocsGetCall request call. Exactly one of *BlocsItem or error will be non-nil.
func (c *BlocsGetCall) Do() (*BlocsItem, error) {
	urls := c.s.baseURL + "/blocs/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &BlocsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// Blocs description: A bloc is the managing container of coalitions.
func Blocs(s *Service) *BlocsService {
	return &BlocsService{s: s}
}

// BroadcastsService handles 42 API Broadcasts.
type BroadcastsService struct {
	s *Service
}

// BroadcastsItem is a 42 API type
type BroadcastsItem struct {
}

// BroadcastsListCall description:
type BroadcastsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListBroadcastsResponse represents a list response.
type ListBroadcastsResponse struct {
	Broadcasts []*BroadcastsItem
}

// List returns a BroadcastsListCall which can request data from the 42 API.
func (s *BroadcastsService) List() *BroadcastsListCall {
	return &BroadcastsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Broadcasts matching the specified setting.
func (c *BroadcastsListCall) P(key string, values []string) *BroadcastsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a BroadcastsListCall request call. Exactly one of *ListBroadcastsResponse or error will be non-nil.
func (c *BroadcastsListCall) Do() (*ListBroadcastsResponse, error) {
	urls := c.s.baseURL + "/broadcasts/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListBroadcastsResponse{}
	ret.Broadcasts = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Broadcasts); err != nil {
		return nil, err
	}
	return ret, nil
}

// Broadcasts description: Broadcasts publicated on a campus
func Broadcasts(s *Service) *BroadcastsService {
	return &BroadcastsService{s: s}
}

// CampusService handles 42 API Campus.
type CampusService struct {
	s *Service
}

// CampusItem is a 42 API type
type CampusItem struct {
	Endpoint interface{} `json:"endpoint"`
	ID       int         `json:"id"`
	Language struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		IDentifier string `json:"identifier"`
	} `json:"language"`
	Name        string `json:"name"`
	TimeZone    string `json:"time_zone"`
	UsersCount  int    `json:"users_count"`
	VogsphereID int    `json:"vogsphere_id"`
}

// CampusListCall description:
type CampusListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListCampusResponse represents a list response.
type ListCampusResponse struct {
	Campus []*CampusItem
}

// List returns a CampusListCall which can request data from the 42 API.
func (s *CampusService) List() *CampusListCall {
	return &CampusListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Campus matching the specified setting.
func (c *CampusListCall) P(key string, values []string) *CampusListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a CampusListCall request call. Exactly one of *ListCampusResponse or error will be non-nil.
func (c *CampusListCall) Do() (*ListCampusResponse, error) {
	urls := c.s.baseURL + "/campus/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListCampusResponse{}
	ret.Campus = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Campus); err != nil {
		return nil, err
	}
	return ret, nil
}

// CampusGetCall description:
type CampusGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified CampusItem.
func (s *CampusService) Get(ID string) *CampusGetCall {
	return &CampusGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a CampusGetCall request call. Exactly one of *CampusItem or error will be non-nil.
func (c *CampusGetCall) Do() (*CampusItem, error) {
	urls := c.s.baseURL + "/campus/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &CampusItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// Campus description: Places where 42 users works
func Campus(s *Service) *CampusService {
	return &CampusService{s: s}
}

// CampusUsersService handles 42 API CampusUsers.
type CampusUsersService struct {
	s *Service
}

// CampusUsersItem is a 42 API type
type CampusUsersItem struct {
	CampusID  int  `json:"campus_id"`
	ID        int  `json:"id"`
	IsPrimary bool `json:"is_primary"`
	UserID    int  `json:"user_id"`
}

// CampusUsersListCall description:
type CampusUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListCampusUsersResponse represents a list response.
type ListCampusUsersResponse struct {
	CampusUsers []*CampusUsersItem
}

// List returns a CampusUsersListCall which can request data from the 42 API.
func (s *CampusUsersService) List() *CampusUsersListCall {
	return &CampusUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return CampusUsers matching the specified setting.
func (c *CampusUsersListCall) P(key string, values []string) *CampusUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a CampusUsersListCall request call. Exactly one of *ListCampusUsersResponse or error will be non-nil.
func (c *CampusUsersListCall) Do() (*ListCampusUsersResponse, error) {
	urls := c.s.baseURL + "/campus_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListCampusUsersResponse{}
	ret.CampusUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).CampusUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// CampusUsersGetCall description:
type CampusUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified CampusUsersItem.
func (s *CampusUsersService) Get(ID string) *CampusUsersGetCall {
	return &CampusUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a CampusUsersGetCall request call. Exactly one of *CampusUsersItem or error will be non-nil.
func (c *CampusUsersGetCall) Do() (*CampusUsersItem, error) {
	urls := c.s.baseURL + "/campus_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &CampusUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// CampusUsers description: The users wich are in a campus
func CampusUsers(s *Service) *CampusUsersService {
	return &CampusUsersService{s: s}
}

// CampusUsersActivitiesService handles 42 API CampusUsersActivities.
type CampusUsersActivitiesService struct {
	s *Service
}

// CampusUsersActivitiesItem is a 42 API type
type CampusUsersActivitiesItem struct {
}

// CampusUsersActivitiesListCall description:
type CampusUsersActivitiesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListCampusUsersActivitiesResponse represents a list response.
type ListCampusUsersActivitiesResponse struct {
	CampusUsersActivities []*CampusUsersActivitiesItem
}

// List returns a CampusUsersActivitiesListCall which can request data from the 42 API.
func (s *CampusUsersActivitiesService) List() *CampusUsersActivitiesListCall {
	return &CampusUsersActivitiesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return CampusUsersActivities matching the specified setting.
func (c *CampusUsersActivitiesListCall) P(key string, values []string) *CampusUsersActivitiesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a CampusUsersActivitiesListCall request call. Exactly one of *ListCampusUsersActivitiesResponse or error will be non-nil.
func (c *CampusUsersActivitiesListCall) Do() (*ListCampusUsersActivitiesResponse, error) {
	urls := c.s.baseURL + "/campus_users_activities/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListCampusUsersActivitiesResponse{}
	ret.CampusUsersActivities = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).CampusUsersActivities); err != nil {
		return nil, err
	}
	return ret, nil
}

// CampusUsersActivities description: Get all the users activities for a campus
func CampusUsersActivities(s *Service) *CampusUsersActivitiesService {
	return &CampusUsersActivitiesService{s: s}
}

// CertificatesService handles 42 API Certificates.
type CertificatesService struct {
	s *Service
}

// CertificatesItem is a 42 API type
type CertificatesItem struct {
}

// CertificatesListCall description:
type CertificatesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListCertificatesResponse represents a list response.
type ListCertificatesResponse struct {
	Certificates []*CertificatesItem
}

// List returns a CertificatesListCall which can request data from the 42 API.
func (s *CertificatesService) List() *CertificatesListCall {
	return &CertificatesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Certificates matching the specified setting.
func (c *CertificatesListCall) P(key string, values []string) *CertificatesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a CertificatesListCall request call. Exactly one of *ListCertificatesResponse or error will be non-nil.
func (c *CertificatesListCall) Do() (*ListCertificatesResponse, error) {
	urls := c.s.baseURL + "/certificates/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListCertificatesResponse{}
	ret.Certificates = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Certificates); err != nil {
		return nil, err
	}
	return ret, nil
}

// CertificatesGetCall description:
type CertificatesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified CertificatesItem.
func (s *CertificatesService) Get(ID string) *CertificatesGetCall {
	return &CertificatesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a CertificatesGetCall request call. Exactly one of *CertificatesItem or error will be non-nil.
func (c *CertificatesGetCall) Do() (*CertificatesItem, error) {
	urls := c.s.baseURL + "/certificates/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &CertificatesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// Certificates description: certificates
func Certificates(s *Service) *CertificatesService {
	return &CertificatesService{s: s}
}

// CertificatesUsersService handles 42 API CertificatesUsers.
type CertificatesUsersService struct {
	s *Service
}

// CertificatesUsersItem is a 42 API type
type CertificatesUsersItem struct {
}

// CertificatesUsersListCall description:
type CertificatesUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListCertificatesUsersResponse represents a list response.
type ListCertificatesUsersResponse struct {
	CertificatesUsers []*CertificatesUsersItem
}

// List returns a CertificatesUsersListCall which can request data from the 42 API.
func (s *CertificatesUsersService) List() *CertificatesUsersListCall {
	return &CertificatesUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return CertificatesUsers matching the specified setting.
func (c *CertificatesUsersListCall) P(key string, values []string) *CertificatesUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a CertificatesUsersListCall request call. Exactly one of *ListCertificatesUsersResponse or error will be non-nil.
func (c *CertificatesUsersListCall) Do() (*ListCertificatesUsersResponse, error) {
	urls := c.s.baseURL + "/certificates_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListCertificatesUsersResponse{}
	ret.CertificatesUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).CertificatesUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// CertificatesUsersGetCall description:
type CertificatesUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified CertificatesUsersItem.
func (s *CertificatesUsersService) Get(ID string) *CertificatesUsersGetCall {
	return &CertificatesUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a CertificatesUsersGetCall request call. Exactly one of *CertificatesUsersItem or error will be non-nil.
func (c *CertificatesUsersGetCall) Do() (*CertificatesUsersItem, error) {
	urls := c.s.baseURL + "/certificates_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &CertificatesUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// CertificatesUsers description: User belonging to a certificate.
func CertificatesUsers(s *Service) *CertificatesUsersService {
	return &CertificatesUsersService{s: s}
}

// ClosesService handles 42 API Closes.
type ClosesService struct {
	s *Service
}

// ClosesItem is a 42 API type
type ClosesItem struct {
	Closer struct {
		URL   string `json:"url"`
		ID    int    `json:"id"`
		Login string `json:"login"`
	} `json:"closer"`
	CommunityServices []struct {
		ID         int       `json:"id"`
		Duration   int       `json:"duration"`
		ScheduleAt time.Time `json:"schedule_at"`
		Occupation string    `json:"occupation"`
		State      string    `json:"state"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	} `json:"community_services"`
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
	Reason    string    `json:"reason"`
	State     string    `json:"state"`
	UpdatedAt time.Time `json:"updated_at"`
	User      struct {
		Login string `json:"login"`
		URL   string `json:"url"`
		ID    int    `json:"id"`
	} `json:"user"`
}

// ClosesListCall description:
type ClosesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListClosesResponse represents a list response.
type ListClosesResponse struct {
	Closes []*ClosesItem
}

// List returns a ClosesListCall which can request data from the 42 API.
func (s *ClosesService) List() *ClosesListCall {
	return &ClosesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Closes matching the specified setting.
func (c *ClosesListCall) P(key string, values []string) *ClosesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ClosesListCall request call. Exactly one of *ListClosesResponse or error will be non-nil.
func (c *ClosesListCall) Do() (*ListClosesResponse, error) {
	urls := c.s.baseURL + "/closes/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListClosesResponse{}
	ret.Closes = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Closes); err != nil {
		return nil, err
	}
	return ret, nil
}

// ClosesGetCall description:
type ClosesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified ClosesItem.
func (s *ClosesService) Get(ID string) *ClosesGetCall {
	return &ClosesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a ClosesGetCall request call. Exactly one of *ClosesItem or error will be non-nil.
func (c *ClosesGetCall) Do() (*ClosesItem, error) {
	urls := c.s.baseURL + "/closes/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ClosesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Closes description: The closing of a 42 account
func Closes(s *Service) *ClosesService {
	return &ClosesService{s: s}
}

// CoalitionsService handles 42 API Coalitions.
type CoalitionsService struct {
	s *Service
}

// CoalitionsItem is a 42 API type
type CoalitionsItem struct {
	Color    string `json:"color"`
	ID       int    `json:"id"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
	Score    int    `json:"score"`
	Slug     string `json:"slug"`
	UserID   int    `json:"user_id"`
}

// CoalitionsListCall description:
type CoalitionsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListCoalitionsResponse represents a list response.
type ListCoalitionsResponse struct {
	Coalitions []*CoalitionsItem
}

// List returns a CoalitionsListCall which can request data from the 42 API.
func (s *CoalitionsService) List() *CoalitionsListCall {
	return &CoalitionsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Coalitions matching the specified setting.
func (c *CoalitionsListCall) P(key string, values []string) *CoalitionsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a CoalitionsListCall request call. Exactly one of *ListCoalitionsResponse or error will be non-nil.
func (c *CoalitionsListCall) Do() (*ListCoalitionsResponse, error) {
	urls := c.s.baseURL + "/coalitions/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListCoalitionsResponse{}
	ret.Coalitions = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Coalitions); err != nil {
		return nil, err
	}
	return ret, nil
}

// CoalitionsGetCall description:
type CoalitionsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified CoalitionsItem.
func (s *CoalitionsService) Get(ID string) *CoalitionsGetCall {
	return &CoalitionsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a CoalitionsGetCall request call. Exactly one of *CoalitionsItem or error will be non-nil.
func (c *CoalitionsGetCall) Do() (*CoalitionsItem, error) {
	urls := c.s.baseURL + "/coalitions/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &CoalitionsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// Coalitions description: A users competing inside of a bloc.
func Coalitions(s *Service) *CoalitionsService {
	return &CoalitionsService{s: s}
}

// CoalitionsUsersService handles 42 API CoalitionsUsers.
type CoalitionsUsersService struct {
	s *Service
}

// CoalitionsUsersItem is a 42 API type
type CoalitionsUsersItem struct {
	CoalitionID int       `json:"coalition_id"`
	CreatedAt   time.Time `json:"created_at"`
	ID          int       `json:"id"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int       `json:"user_id"`
}

// CoalitionsUsersListCall description:
type CoalitionsUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListCoalitionsUsersResponse represents a list response.
type ListCoalitionsUsersResponse struct {
	CoalitionsUsers []*CoalitionsUsersItem
}

// List returns a CoalitionsUsersListCall which can request data from the 42 API.
func (s *CoalitionsUsersService) List() *CoalitionsUsersListCall {
	return &CoalitionsUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return CoalitionsUsers matching the specified setting.
func (c *CoalitionsUsersListCall) P(key string, values []string) *CoalitionsUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a CoalitionsUsersListCall request call. Exactly one of *ListCoalitionsUsersResponse or error will be non-nil.
func (c *CoalitionsUsersListCall) Do() (*ListCoalitionsUsersResponse, error) {
	urls := c.s.baseURL + "/coalitions_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListCoalitionsUsersResponse{}
	ret.CoalitionsUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).CoalitionsUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// CoalitionsUsersGetCall description:
type CoalitionsUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified CoalitionsUsersItem.
func (s *CoalitionsUsersService) Get(ID string) *CoalitionsUsersGetCall {
	return &CoalitionsUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a CoalitionsUsersGetCall request call. Exactly one of *CoalitionsUsersItem or error will be non-nil.
func (c *CoalitionsUsersGetCall) Do() (*CoalitionsUsersItem, error) {
	urls := c.s.baseURL + "/coalitions_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &CoalitionsUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// CoalitionsUsers description: coalition.
func CoalitionsUsers(s *Service) *CoalitionsUsersService {
	return &CoalitionsUsersService{s: s}
}

// CommandsService handles 42 API Commands.
type CommandsService struct {
	s *Service
}

// CommandsItem is a 42 API type
type CommandsItem struct {
}

// CommandsListCall description:
type CommandsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListCommandsResponse represents a list response.
type ListCommandsResponse struct {
	Commands []*CommandsItem
}

// List returns a CommandsListCall which can request data from the 42 API.
func (s *CommandsService) List() *CommandsListCall {
	return &CommandsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Commands matching the specified setting.
func (c *CommandsListCall) P(key string, values []string) *CommandsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a CommandsListCall request call. Exactly one of *ListCommandsResponse or error will be non-nil.
func (c *CommandsListCall) Do() (*ListCommandsResponse, error) {
	urls := c.s.baseURL + "/commands/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListCommandsResponse{}
	ret.Commands = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Commands); err != nil {
		return nil, err
	}
	return ret, nil
}

// CommandsGetCall description:
type CommandsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified CommandsItem.
func (s *CommandsService) Get(ID string) *CommandsGetCall {
	return &CommandsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a CommandsGetCall request call. Exactly one of *CommandsItem or error will be non-nil.
func (c *CommandsGetCall) Do() (*CommandsItem, error) {
	urls := c.s.baseURL + "/commands/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &CommandsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Commands description: Products are sold on the intranet shop, here are commands
func Commands(s *Service) *CommandsService {
	return &CommandsService{s: s}
}

// CommunityServicesService handles 42 API CommunityServices.
type CommunityServicesService struct {
	s *Service
}

// CommunityServicesItem is a 42 API type
type CommunityServicesItem struct {
	Close struct {
		ID        int       `json:"id"`
		Reason    string    `json:"reason"`
		State     string    `json:"state"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"close"`
	CreatedAt  time.Time `json:"created_at"`
	Duration   int       `json:"duration"`
	ID         int       `json:"id"`
	Occupation string    `json:"occupation"`
	ScheduleAt time.Time `json:"schedule_at"`
	State      string    `json:"state"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// CommunityServicesListCall description:
type CommunityServicesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListCommunityServicesResponse represents a list response.
type ListCommunityServicesResponse struct {
	CommunityServices []*CommunityServicesItem
}

// List returns a CommunityServicesListCall which can request data from the 42 API.
func (s *CommunityServicesService) List() *CommunityServicesListCall {
	return &CommunityServicesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return CommunityServices matching the specified setting.
func (c *CommunityServicesListCall) P(key string, values []string) *CommunityServicesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a CommunityServicesListCall request call. Exactly one of *ListCommunityServicesResponse or error will be non-nil.
func (c *CommunityServicesListCall) Do() (*ListCommunityServicesResponse, error) {
	urls := c.s.baseURL + "/community_services/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListCommunityServicesResponse{}
	ret.CommunityServices = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).CommunityServices); err != nil {
		return nil, err
	}
	return ret, nil
}

// CommunityServicesGetCall description:
type CommunityServicesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified CommunityServicesItem.
func (s *CommunityServicesService) Get(ID string) *CommunityServicesGetCall {
	return &CommunityServicesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a CommunityServicesGetCall request call. Exactly one of *CommunityServicesItem or error will be non-nil.
func (c *CommunityServicesGetCall) Do() (*CommunityServicesItem, error) {
	urls := c.s.baseURL + "/community_services/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &CommunityServicesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// CommunityServices description: A task that an user have to do for the community. Usually linked with a close.
func CommunityServices(s *Service) *CommunityServicesService {
	return &CommunityServicesService{s: s}
}

// CursusService handles 42 API Cursus.
type CursusService struct {
	s *Service
}

// CursusItem is a 42 API type
type CursusItem struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
}

// CursusListCall description:
type CursusListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListCursusResponse represents a list response.
type ListCursusResponse struct {
	Cursus []*CursusItem
}

// List returns a CursusListCall which can request data from the 42 API.
func (s *CursusService) List() *CursusListCall {
	return &CursusListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Cursus matching the specified setting.
func (c *CursusListCall) P(key string, values []string) *CursusListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a CursusListCall request call. Exactly one of *ListCursusResponse or error will be non-nil.
func (c *CursusListCall) Do() (*ListCursusResponse, error) {
	urls := c.s.baseURL + "/cursus/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListCursusResponse{}
	ret.Cursus = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Cursus); err != nil {
		return nil, err
	}
	return ret, nil
}

// CursusGetCall description:
type CursusGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified CursusItem.
func (s *CursusService) Get(ID string) *CursusGetCall {
	return &CursusGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a CursusGetCall request call. Exactly one of *CursusItem or error will be non-nil.
func (c *CursusGetCall) Do() (*CursusItem, error) {
	urls := c.s.baseURL + "/cursus/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &CursusItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Cursus description: An educational cycle in 42
func Cursus(s *Service) *CursusService {
	return &CursusService{s: s}
}

// CursusUsersService handles 42 API CursusUsers.
type CursusUsersService struct {
	s *Service
}

// CursusUsersItem is a 42 API type
type CursusUsersItem struct {
	BeginAt time.Time `json:"begin_at"`
	Cursus  struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		Name      string    `json:"name"`
		Slug      string    `json:"slug"`
	} `json:"cursus"`
	CursusID     int           `json:"cursus_id"`
	EndAt        interface{}   `json:"end_at"`
	Grade        string        `json:"grade"`
	HasCoalition bool          `json:"has_coalition"`
	ID           int           `json:"id"`
	Level        int           `json:"level"`
	Skills       []interface{} `json:"skills"`
	User         struct {
		URL   string `json:"url"`
		ID    int    `json:"id"`
		Login string `json:"login"`
	} `json:"user"`
}

// CursusUsersListCall description:
type CursusUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListCursusUsersResponse represents a list response.
type ListCursusUsersResponse struct {
	CursusUsers []*CursusUsersItem
}

// List returns a CursusUsersListCall which can request data from the 42 API.
func (s *CursusUsersService) List() *CursusUsersListCall {
	return &CursusUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return CursusUsers matching the specified setting.
func (c *CursusUsersListCall) P(key string, values []string) *CursusUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a CursusUsersListCall request call. Exactly one of *ListCursusUsersResponse or error will be non-nil.
func (c *CursusUsersListCall) Do() (*ListCursusUsersResponse, error) {
	urls := c.s.baseURL + "/cursus_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListCursusUsersResponse{}
	ret.CursusUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).CursusUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// CursusUsersGetCall description:
type CursusUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified CursusUsersItem.
func (s *CursusUsersService) Get(ID string) *CursusUsersGetCall {
	return &CursusUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a CursusUsersGetCall request call. Exactly one of *CursusUsersItem or error will be non-nil.
func (c *CursusUsersGetCall) Do() (*CursusUsersItem, error) {
	urls := c.s.baseURL + "/cursus_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &CursusUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// CursusUsers description: The users wich are in a cursus
func CursusUsers(s *Service) *CursusUsersService {
	return &CursusUsersService{s: s}
}

// DashesService handles 42 API Dashes.
type DashesService struct {
	s *Service
}

// DashesItem is a 42 API type
type DashesItem struct {
}

// DashesListCall description:
type DashesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListDashesResponse represents a list response.
type ListDashesResponse struct {
	Dashes []*DashesItem
}

// List returns a DashesListCall which can request data from the 42 API.
func (s *DashesService) List() *DashesListCall {
	return &DashesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Dashes matching the specified setting.
func (c *DashesListCall) P(key string, values []string) *DashesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a DashesListCall request call. Exactly one of *ListDashesResponse or error will be non-nil.
func (c *DashesListCall) Do() (*ListDashesResponse, error) {
	urls := c.s.baseURL + "/dashes/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListDashesResponse{}
	ret.Dashes = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Dashes); err != nil {
		return nil, err
	}
	return ret, nil
}

// DashesGetCall description:
type DashesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified DashesItem.
func (s *DashesService) Get(ID string) *DashesGetCall {
	return &DashesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a DashesGetCall request call. Exactly one of *DashesItem or error will be non-nil.
func (c *DashesGetCall) Do() (*DashesItem, error) {
	urls := c.s.baseURL + "/dashes/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &DashesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Dashes description: The Dash is a short-time project
func Dashes(s *Service) *DashesService {
	return &DashesService{s: s}
}

// DashesUsersService handles 42 API DashesUsers.
type DashesUsersService struct {
	s *Service
}

// DashesUsersItem is a 42 API type
type DashesUsersItem struct {
	DashID    int         `json:"dash_id"`
	FinalMark interface{} `json:"final_mark"`
	ID        int         `json:"id"`
	RepoURL   interface{} `json:"repo_url"`
	RepoUUID  interface{} `json:"repo_uuid"`
	UserID    int         `json:"user_id"`
}

// DashesUsersListCall description:
type DashesUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListDashesUsersResponse represents a list response.
type ListDashesUsersResponse struct {
	DashesUsers []*DashesUsersItem
}

// List returns a DashesUsersListCall which can request data from the 42 API.
func (s *DashesUsersService) List() *DashesUsersListCall {
	return &DashesUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return DashesUsers matching the specified setting.
func (c *DashesUsersListCall) P(key string, values []string) *DashesUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a DashesUsersListCall request call. Exactly one of *ListDashesUsersResponse or error will be non-nil.
func (c *DashesUsersListCall) Do() (*ListDashesUsersResponse, error) {
	urls := c.s.baseURL + "/dashes_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListDashesUsersResponse{}
	ret.DashesUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).DashesUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// DashesUsersGetCall description:
type DashesUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified DashesUsersItem.
func (s *DashesUsersService) Get(ID string) *DashesUsersGetCall {
	return &DashesUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a DashesUsersGetCall request call. Exactly one of *DashesUsersItem or error will be non-nil.
func (c *DashesUsersGetCall) Do() (*DashesUsersItem, error) {
	urls := c.s.baseURL + "/dashes_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &DashesUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// DashesUsers description: The dash of a user
func DashesUsers(s *Service) *DashesUsersService {
	return &DashesUsersService{s: s}
}

// EndpointsService handles 42 API Endpoints.
type EndpointsService struct {
	s *Service
}

// EndpointsItem is a 42 API type
type EndpointsItem struct {
	Campus []struct {
		Language struct {
			CreatedAt  time.Time `json:"created_at"`
			UpdatedAt  time.Time `json:"updated_at"`
			ID         int       `json:"id"`
			Name       string    `json:"name"`
			IDentifier string    `json:"identifier"`
		} `json:"language"`
		UsersCount int    `json:"users_count"`
		ID         int    `json:"id"`
		Name       string `json:"name"`
		TimeZone   string `json:"time_zone"`
	} `json:"campus"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	ID          int       `json:"id"`
	UpdatedAt   time.Time `json:"updated_at"`
	URL         string    `json:"url"`
}

// EndpointsListCall description:
type EndpointsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListEndpointsResponse represents a list response.
type ListEndpointsResponse struct {
	Endpoints []*EndpointsItem
}

// List returns a EndpointsListCall which can request data from the 42 API.
func (s *EndpointsService) List() *EndpointsListCall {
	return &EndpointsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Endpoints matching the specified setting.
func (c *EndpointsListCall) P(key string, values []string) *EndpointsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a EndpointsListCall request call. Exactly one of *ListEndpointsResponse or error will be non-nil.
func (c *EndpointsListCall) Do() (*ListEndpointsResponse, error) {
	urls := c.s.baseURL + "/endpoints/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListEndpointsResponse{}
	ret.Endpoints = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Endpoints); err != nil {
		return nil, err
	}
	return ret, nil
}

// EndpointsGetCall description:
type EndpointsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified EndpointsItem.
func (s *EndpointsService) Get(ID string) *EndpointsGetCall {
	return &EndpointsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a EndpointsGetCall request call. Exactly one of *EndpointsItem or error will be non-nil.
func (c *EndpointsGetCall) Do() (*EndpointsItem, error) {
	urls := c.s.baseURL + "/endpoints/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &EndpointsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Endpoints description: A endpoint for a campus
func Endpoints(s *Service) *EndpointsService {
	return &EndpointsService{s: s}
}

// EvaluationsService handles 42 API Evaluations.
type EvaluationsService struct {
	s *Service
}

// EvaluationsItem is a 42 API type
type EvaluationsItem struct {
}

// EvaluationsListCall description:
type EvaluationsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListEvaluationsResponse represents a list response.
type ListEvaluationsResponse struct {
	Evaluations []*EvaluationsItem
}

// List returns a EvaluationsListCall which can request data from the 42 API.
func (s *EvaluationsService) List() *EvaluationsListCall {
	return &EvaluationsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Evaluations matching the specified setting.
func (c *EvaluationsListCall) P(key string, values []string) *EvaluationsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a EvaluationsListCall request call. Exactly one of *ListEvaluationsResponse or error will be non-nil.
func (c *EvaluationsListCall) Do() (*ListEvaluationsResponse, error) {
	urls := c.s.baseURL + "/evaluations/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListEvaluationsResponse{}
	ret.Evaluations = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Evaluations); err != nil {
		return nil, err
	}
	return ret, nil
}

// EvaluationsGetCall description:
type EvaluationsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified EvaluationsItem.
func (s *EvaluationsService) Get(ID string) *EvaluationsGetCall {
	return &EvaluationsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a EvaluationsGetCall request call. Exactly one of *EvaluationsItem or error will be non-nil.
func (c *EvaluationsGetCall) Do() (*EvaluationsItem, error) {
	urls := c.s.baseURL + "/evaluations/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &EvaluationsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Evaluations description: The Evaluation of a project
func Evaluations(s *Service) *EvaluationsService {
	return &EvaluationsService{s: s}
}

// EventsService handles 42 API Events.
type EventsService struct {
	s *Service
}

// EventsItem is a 42 API type
type EventsItem struct {
	BeginAt                   time.Time `json:"begin_at"`
	CampusIDs                 []int     `json:"campus_ids"`
	CreatedAt                 time.Time `json:"created_at"`
	CursusIDs                 []int     `json:"cursus_ids"`
	Description               string    `json:"description"`
	EndAt                     time.Time `json:"end_at"`
	ID                        int       `json:"id"`
	Kind                      string    `json:"kind"`
	Location                  string    `json:"location"`
	MaxPeople                 int       `json:"max_people"`
	Name                      string    `json:"name"`
	NbrSubscribers            int       `json:"nbr_subscribers"`
	ProhibitionOfCancellation int       `json:"prohibition_of_cancellation"`
	Themes                    []struct {
		CreatedAt time.Time `json:"created_at"`
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"themes"`
	UpdatedAt time.Time `json:"updated_at"`
	Waitlist  struct {
		CreatedAt        time.Time `json:"created_at"`
		ID               int       `json:"id"`
		UpdatedAt        time.Time `json:"updated_at"`
		WaitlistableID   int       `json:"waitlistable_id"`
		WaitlistableType string    `json:"waitlistable_type"`
	} `json:"waitlist"`
}

// EventsListCall description:
type EventsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListEventsResponse represents a list response.
type ListEventsResponse struct {
	Events []*EventsItem
}

// List returns a EventsListCall which can request data from the 42 API.
func (s *EventsService) List() *EventsListCall {
	return &EventsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Events matching the specified setting.
func (c *EventsListCall) P(key string, values []string) *EventsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a EventsListCall request call. Exactly one of *ListEventsResponse or error will be non-nil.
func (c *EventsListCall) Do() (*ListEventsResponse, error) {
	urls := c.s.baseURL + "/events/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListEventsResponse{}
	ret.Events = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Events); err != nil {
		return nil, err
	}
	return ret, nil
}

// EventsGetCall description:
type EventsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified EventsItem.
func (s *EventsService) Get(ID string) *EventsGetCall {
	return &EventsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a EventsGetCall request call. Exactly one of *EventsItem or error will be non-nil.
func (c *EventsGetCall) Do() (*EventsItem, error) {
	urls := c.s.baseURL + "/events/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &EventsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Events description: The events in a campus or a cursus
func Events(s *Service) *EventsService {
	return &EventsService{s: s}
}

// EventsUsersService handles 42 API EventsUsers.
type EventsUsersService struct {
	s *Service
}

// EventsUsersItem is a 42 API type
type EventsUsersItem struct {
}

// EventsUsersListCall description:
type EventsUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListEventsUsersResponse represents a list response.
type ListEventsUsersResponse struct {
	EventsUsers []*EventsUsersItem
}

// List returns a EventsUsersListCall which can request data from the 42 API.
func (s *EventsUsersService) List() *EventsUsersListCall {
	return &EventsUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return EventsUsers matching the specified setting.
func (c *EventsUsersListCall) P(key string, values []string) *EventsUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a EventsUsersListCall request call. Exactly one of *ListEventsUsersResponse or error will be non-nil.
func (c *EventsUsersListCall) Do() (*ListEventsUsersResponse, error) {
	urls := c.s.baseURL + "/events_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListEventsUsersResponse{}
	ret.EventsUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).EventsUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// EventsUsersGetCall description:
type EventsUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified EventsUsersItem.
func (s *EventsUsersService) Get(ID string) *EventsUsersGetCall {
	return &EventsUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a EventsUsersGetCall request call. Exactly one of *EventsUsersItem or error will be non-nil.
func (c *EventsUsersGetCall) Do() (*EventsUsersItem, error) {
	urls := c.s.baseURL + "/events_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &EventsUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// EventsUsers description: Users registered to an event
func EventsUsers(s *Service) *EventsUsersService {
	return &EventsUsersService{s: s}
}

// ExamsService handles 42 API Exams.
type ExamsService struct {
	s *Service
}

// ExamsItem is a 42 API type
type ExamsItem struct {
	BeginAt time.Time `json:"begin_at"`
	Campus  []struct {
		UsersCount  int    `json:"users_count"`
		VogsphereID int    `json:"vogsphere_id"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
		TimeZone    string `json:"time_zone"`
		Language    struct {
			ID         int       `json:"id"`
			Name       string    `json:"name"`
			IDentifier string    `json:"identifier"`
			CreatedAt  time.Time `json:"created_at"`
			UpdatedAt  time.Time `json:"updated_at"`
		} `json:"language"`
	} `json:"campus"`
	Cursus []struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		Name      string    `json:"name"`
		Slug      string    `json:"slug"`
	} `json:"cursus"`
	EndAt          time.Time `json:"end_at"`
	ID             int       `json:"id"`
	IPRange        string    `json:"ip_range"`
	Location       string    `json:"location"`
	MaxPeople      int       `json:"max_people"`
	Name           string    `json:"name"`
	NbrSubscribers int       `json:"nbr_subscribers"`
	Projects       []struct {
		Slug        string        `json:"slug"`
		Attachments []interface{} `json:"attachments"`
		Exam        bool          `json:"exam"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		Parent      interface{}   `json:"parent"`
		Children    []interface{} `json:"children"`
		Objectives  []string      `json:"objectives"`
		Tier        int           `json:"tier"`
		CreatedAt   time.Time     `json:"created_at"`
		UpdatedAt   time.Time     `json:"updated_at"`
		ID          int           `json:"id"`
	} `json:"projects"`
}

// ExamsListCall description:
type ExamsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListExamsResponse represents a list response.
type ListExamsResponse struct {
	Exams []*ExamsItem
}

// List returns a ExamsListCall which can request data from the 42 API.
func (s *ExamsService) List() *ExamsListCall {
	return &ExamsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Exams matching the specified setting.
func (c *ExamsListCall) P(key string, values []string) *ExamsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ExamsListCall request call. Exactly one of *ListExamsResponse or error will be non-nil.
func (c *ExamsListCall) Do() (*ListExamsResponse, error) {
	urls := c.s.baseURL + "/exams/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListExamsResponse{}
	ret.Exams = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Exams); err != nil {
		return nil, err
	}
	return ret, nil
}

// ExamsGetCall description:
type ExamsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified ExamsItem.
func (s *ExamsService) Get(ID string) *ExamsGetCall {
	return &ExamsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a ExamsGetCall request call. Exactly one of *ExamsItem or error will be non-nil.
func (c *ExamsGetCall) Do() (*ExamsItem, error) {
	urls := c.s.baseURL + "/exams/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ExamsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Exams description: The exam in a campus or a cursus
func Exams(s *Service) *ExamsService {
	return &ExamsService{s: s}
}

// ExamsUsersService handles 42 API ExamsUsers.
type ExamsUsersService struct {
	s *Service
}

// ExamsUsersItem is a 42 API type
type ExamsUsersItem struct {
}

// ExamsUsersListCall description:
type ExamsUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListExamsUsersResponse represents a list response.
type ListExamsUsersResponse struct {
	ExamsUsers []*ExamsUsersItem
}

// List returns a ExamsUsersListCall which can request data from the 42 API.
func (s *ExamsUsersService) List() *ExamsUsersListCall {
	return &ExamsUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return ExamsUsers matching the specified setting.
func (c *ExamsUsersListCall) P(key string, values []string) *ExamsUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ExamsUsersListCall request call. Exactly one of *ListExamsUsersResponse or error will be non-nil.
func (c *ExamsUsersListCall) Do() (*ListExamsUsersResponse, error) {
	urls := c.s.baseURL + "/exams_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListExamsUsersResponse{}
	ret.ExamsUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).ExamsUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// ExamsUsers description:
func ExamsUsers(s *Service) *ExamsUsersService {
	return &ExamsUsersService{s: s}
}

// ExperiencesService handles 42 API Experiences.
type ExperiencesService struct {
	s *Service
}

// ExperiencesItem is a 42 API type
type ExperiencesItem struct {
}

// ExperiencesListCall description:
type ExperiencesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListExperiencesResponse represents a list response.
type ListExperiencesResponse struct {
	Experiences []*ExperiencesItem
}

// List returns a ExperiencesListCall which can request data from the 42 API.
func (s *ExperiencesService) List() *ExperiencesListCall {
	return &ExperiencesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Experiences matching the specified setting.
func (c *ExperiencesListCall) P(key string, values []string) *ExperiencesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ExperiencesListCall request call. Exactly one of *ListExperiencesResponse or error will be non-nil.
func (c *ExperiencesListCall) Do() (*ListExperiencesResponse, error) {
	urls := c.s.baseURL + "/experiences/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListExperiencesResponse{}
	ret.Experiences = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Experiences); err != nil {
		return nil, err
	}
	return ret, nil
}

// ExperiencesGetCall description:
type ExperiencesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified ExperiencesItem.
func (s *ExperiencesService) Get(ID string) *ExperiencesGetCall {
	return &ExperiencesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a ExperiencesGetCall request call. Exactly one of *ExperiencesItem or error will be non-nil.
func (c *ExperiencesGetCall) Do() (*ExperiencesItem, error) {
	urls := c.s.baseURL + "/experiences/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ExperiencesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Experiences description: An experience gained by an user in a particular skill.
func Experiences(s *Service) *ExperiencesService {
	return &ExperiencesService{s: s}
}

// ExpertisesService handles 42 API Expertises.
type ExpertisesService struct {
	s *Service
}

// ExpertisesItem is a 42 API type
type ExpertisesItem struct {
	CreatedAt          time.Time `json:"created_at"`
	ExpertisesUsersURL string    `json:"expertises_users_url"`
	ID                 int       `json:"id"`
	Kind               string    `json:"kind"`
	Name               string    `json:"name"`
	Slug               string    `json:"slug"`
	URL                string    `json:"url"`
}

// ExpertisesListCall description:
type ExpertisesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListExpertisesResponse represents a list response.
type ListExpertisesResponse struct {
	Expertises []*ExpertisesItem
}

// List returns a ExpertisesListCall which can request data from the 42 API.
func (s *ExpertisesService) List() *ExpertisesListCall {
	return &ExpertisesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Expertises matching the specified setting.
func (c *ExpertisesListCall) P(key string, values []string) *ExpertisesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ExpertisesListCall request call. Exactly one of *ListExpertisesResponse or error will be non-nil.
func (c *ExpertisesListCall) Do() (*ListExpertisesResponse, error) {
	urls := c.s.baseURL + "/expertises/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListExpertisesResponse{}
	ret.Expertises = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Expertises); err != nil {
		return nil, err
	}
	return ret, nil
}

// ExpertisesGetCall description:
type ExpertisesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified ExpertisesItem.
func (s *ExpertisesService) Get(ID string) *ExpertisesGetCall {
	return &ExpertisesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a ExpertisesGetCall request call. Exactly one of *ExpertisesItem or error will be non-nil.
func (c *ExpertisesGetCall) Do() (*ExpertisesItem, error) {
	urls := c.s.baseURL + "/expertises/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ExpertisesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Expertises description: Pedagogic expertises
func Expertises(s *Service) *ExpertisesService {
	return &ExpertisesService{s: s}
}

// ExpertisesUsersService handles 42 API ExpertisesUsers.
type ExpertisesUsersService struct {
	s *Service
}

// ExpertisesUsersItem is a 42 API type
type ExpertisesUsersItem struct {
	ContactMe bool      `json:"contact_me"`
	CreatedAt time.Time `json:"created_at"`
	Expertise struct {
		ExpertisesUsersURL string    `json:"expertises_users_url"`
		ID                 int       `json:"id"`
		Name               string    `json:"name"`
		Slug               string    `json:"slug"`
		URL                string    `json:"url"`
		Kind               string    `json:"kind"`
		CreatedAt          time.Time `json:"created_at"`
	} `json:"expertise"`
	ExpertiseID int  `json:"expertise_id"`
	ID          int  `json:"id"`
	Interested  bool `json:"interested"`
	User        struct {
		URL   string `json:"url"`
		ID    int    `json:"id"`
		Login string `json:"login"`
	} `json:"user"`
	UserID int `json:"user_id"`
	Value  int `json:"value"`
}

// ExpertisesUsersListCall description:
type ExpertisesUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListExpertisesUsersResponse represents a list response.
type ListExpertisesUsersResponse struct {
	ExpertisesUsers []*ExpertisesUsersItem
}

// List returns a ExpertisesUsersListCall which can request data from the 42 API.
func (s *ExpertisesUsersService) List() *ExpertisesUsersListCall {
	return &ExpertisesUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return ExpertisesUsers matching the specified setting.
func (c *ExpertisesUsersListCall) P(key string, values []string) *ExpertisesUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ExpertisesUsersListCall request call. Exactly one of *ListExpertisesUsersResponse or error will be non-nil.
func (c *ExpertisesUsersListCall) Do() (*ListExpertisesUsersResponse, error) {
	urls := c.s.baseURL + "/expertises_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListExpertisesUsersResponse{}
	ret.ExpertisesUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).ExpertisesUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// ExpertisesUsersGetCall description:
type ExpertisesUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified ExpertisesUsersItem.
func (s *ExpertisesUsersService) Get(ID string) *ExpertisesUsersGetCall {
	return &ExpertisesUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a ExpertisesUsersGetCall request call. Exactly one of *ExpertisesUsersItem or error will be non-nil.
func (c *ExpertisesUsersGetCall) Do() (*ExpertisesUsersItem, error) {
	urls := c.s.baseURL + "/expertises_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ExpertisesUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// ExpertisesUsers description: Users which have an expertise
func ExpertisesUsers(s *Service) *ExpertisesUsersService {
	return &ExpertisesUsersService{s: s}
}

// FeedbacksService handles 42 API Feedbacks.
type FeedbacksService struct {
	s *Service
}

// FeedbacksItem is a 42 API type
type FeedbacksItem struct {
}

// FeedbacksListCall description:
type FeedbacksListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListFeedbacksResponse represents a list response.
type ListFeedbacksResponse struct {
	Feedbacks []*FeedbacksItem
}

// List returns a FeedbacksListCall which can request data from the 42 API.
func (s *FeedbacksService) List() *FeedbacksListCall {
	return &FeedbacksListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Feedbacks matching the specified setting.
func (c *FeedbacksListCall) P(key string, values []string) *FeedbacksListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a FeedbacksListCall request call. Exactly one of *ListFeedbacksResponse or error will be non-nil.
func (c *FeedbacksListCall) Do() (*ListFeedbacksResponse, error) {
	urls := c.s.baseURL + "/feedbacks/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListFeedbacksResponse{}
	ret.Feedbacks = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Feedbacks); err != nil {
		return nil, err
	}
	return ret, nil
}

// FeedbacksGetCall description:
type FeedbacksGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified FeedbacksItem.
func (s *FeedbacksService) Get(ID string) *FeedbacksGetCall {
	return &FeedbacksGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a FeedbacksGetCall request call. Exactly one of *FeedbacksItem or error will be non-nil.
func (c *FeedbacksGetCall) Do() (*FeedbacksItem, error) {
	urls := c.s.baseURL + "/feedbacks/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &FeedbacksItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Feedbacks description: The feedback of a ScaleTeam or an Event
func Feedbacks(s *Service) *FeedbacksService {
	return &FeedbacksService{s: s}
}

// FlagsService handles 42 API Flags.
type FlagsService struct {
	s *Service
}

// FlagsItem is a 42 API type
type FlagsItem struct {
}

// FlagsListCall description:
type FlagsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListFlagsResponse represents a list response.
type ListFlagsResponse struct {
	Flags []*FlagsItem
}

// List returns a FlagsListCall which can request data from the 42 API.
func (s *FlagsService) List() *FlagsListCall {
	return &FlagsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Flags matching the specified setting.
func (c *FlagsListCall) P(key string, values []string) *FlagsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a FlagsListCall request call. Exactly one of *ListFlagsResponse or error will be non-nil.
func (c *FlagsListCall) Do() (*ListFlagsResponse, error) {
	urls := c.s.baseURL + "/flags/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListFlagsResponse{}
	ret.Flags = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Flags); err != nil {
		return nil, err
	}
	return ret, nil
}

// Flags description: Flags from scales
func Flags(s *Service) *FlagsService {
	return &FlagsService{s: s}
}

// FlashUsersService handles 42 API FlashUsers.
type FlashUsersService struct {
	s *Service
}

// FlashUsersItem is a 42 API type
type FlashUsersItem struct {
	EndAt   time.Time `json:"end_at"`
	FlashID int       `json:"flash_id"`
	ID      int       `json:"id"`
	Seen    bool      `json:"seen"`
	User    struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"user"`
}

// FlashUsersListCall description:
type FlashUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListFlashUsersResponse represents a list response.
type ListFlashUsersResponse struct {
	FlashUsers []*FlashUsersItem
}

// List returns a FlashUsersListCall which can request data from the 42 API.
func (s *FlashUsersService) List() *FlashUsersListCall {
	return &FlashUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return FlashUsers matching the specified setting.
func (c *FlashUsersListCall) P(key string, values []string) *FlashUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a FlashUsersListCall request call. Exactly one of *ListFlashUsersResponse or error will be non-nil.
func (c *FlashUsersListCall) Do() (*ListFlashUsersResponse, error) {
	urls := c.s.baseURL + "/flash_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListFlashUsersResponse{}
	ret.FlashUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).FlashUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// FlashUsersGetCall description:
type FlashUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified FlashUsersItem.
func (s *FlashUsersService) Get(ID string) *FlashUsersGetCall {
	return &FlashUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a FlashUsersGetCall request call. Exactly one of *FlashUsersItem or error will be non-nil.
func (c *FlashUsersGetCall) Do() (*FlashUsersItem, error) {
	urls := c.s.baseURL + "/flash_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &FlashUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// FlashUsers description: The Flash Users
func FlashUsers(s *Service) *FlashUsersService {
	return &FlashUsersService{s: s}
}

// FlashesService handles 42 API Flashes.
type FlashesService struct {
	s *Service
}

// FlashesItem is a 42 API type
type FlashesItem struct {
	Content    string `json:"content"`
	Duration   int    `json:"duration"`
	ID         int    `json:"id"`
	IDentifier string `json:"identifier"`
	Selector   string `json:"selector"`
	Title      string `json:"title"`
}

// FlashesListCall description:
type FlashesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListFlashesResponse represents a list response.
type ListFlashesResponse struct {
	Flashes []*FlashesItem
}

// List returns a FlashesListCall which can request data from the 42 API.
func (s *FlashesService) List() *FlashesListCall {
	return &FlashesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Flashes matching the specified setting.
func (c *FlashesListCall) P(key string, values []string) *FlashesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a FlashesListCall request call. Exactly one of *ListFlashesResponse or error will be non-nil.
func (c *FlashesListCall) Do() (*ListFlashesResponse, error) {
	urls := c.s.baseURL + "/flashes/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListFlashesResponse{}
	ret.Flashes = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Flashes); err != nil {
		return nil, err
	}
	return ret, nil
}

// FlashesGetCall description:
type FlashesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified FlashesItem.
func (s *FlashesService) Get(ID string) *FlashesGetCall {
	return &FlashesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a FlashesGetCall request call. Exactly one of *FlashesItem or error will be non-nil.
func (c *FlashesGetCall) Do() (*FlashesItem, error) {
	urls := c.s.baseURL + "/flashes/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &FlashesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// Flashes description: The Flash
func Flashes(s *Service) *FlashesService {
	return &FlashesService{s: s}
}

// GroupsService handles 42 API Groups.
type GroupsService struct {
	s *Service
}

// GroupsItem is a 42 API type
type GroupsItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GroupsListCall description:
type GroupsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListGroupsResponse represents a list response.
type ListGroupsResponse struct {
	Groups []*GroupsItem
}

// List returns a GroupsListCall which can request data from the 42 API.
func (s *GroupsService) List() *GroupsListCall {
	return &GroupsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Groups matching the specified setting.
func (c *GroupsListCall) P(key string, values []string) *GroupsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a GroupsListCall request call. Exactly one of *ListGroupsResponse or error will be non-nil.
func (c *GroupsListCall) Do() (*ListGroupsResponse, error) {
	urls := c.s.baseURL + "/groups/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListGroupsResponse{}
	ret.Groups = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Groups); err != nil {
		return nil, err
	}
	return ret, nil
}

// GroupsGetCall description:
type GroupsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified GroupsItem.
func (s *GroupsService) Get(ID string) *GroupsGetCall {
	return &GroupsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a GroupsGetCall request call. Exactly one of *GroupsItem or error will be non-nil.
func (c *GroupsGetCall) Do() (*GroupsItem, error) {
	urls := c.s.baseURL + "/groups/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &GroupsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Groups description: Groups in which users belong to. It will display a label on their profile and on the forum.
func Groups(s *Service) *GroupsService {
	return &GroupsService{s: s}
}

// GroupsUsersService handles 42 API GroupsUsers.
type GroupsUsersService struct {
	s *Service
}

// GroupsUsersItem is a 42 API type
type GroupsUsersItem struct {
	CreatedAt time.Time `json:"created_at"`
	GroupID   int       `json:"group_id"`
	ID        int       `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    int       `json:"user_id"`
}

// GroupsUsersListCall description:
type GroupsUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListGroupsUsersResponse represents a list response.
type ListGroupsUsersResponse struct {
	GroupsUsers []*GroupsUsersItem
}

// List returns a GroupsUsersListCall which can request data from the 42 API.
func (s *GroupsUsersService) List() *GroupsUsersListCall {
	return &GroupsUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return GroupsUsers matching the specified setting.
func (c *GroupsUsersListCall) P(key string, values []string) *GroupsUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a GroupsUsersListCall request call. Exactly one of *ListGroupsUsersResponse or error will be non-nil.
func (c *GroupsUsersListCall) Do() (*ListGroupsUsersResponse, error) {
	urls := c.s.baseURL + "/groups_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListGroupsUsersResponse{}
	ret.GroupsUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).GroupsUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// GroupsUsersGetCall description:
type GroupsUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified GroupsUsersItem.
func (s *GroupsUsersService) Get(ID string) *GroupsUsersGetCall {
	return &GroupsUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a GroupsUsersGetCall request call. Exactly one of *GroupsUsersItem or error will be non-nil.
func (c *GroupsUsersGetCall) Do() (*GroupsUsersItem, error) {
	urls := c.s.baseURL + "/groups_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &GroupsUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// GroupsUsers description: Users who are in a group.
func GroupsUsers(s *Service) *GroupsUsersService {
	return &GroupsUsersService{s: s}
}

// InternshipsService handles 42 API Internships.
type InternshipsService struct {
	s *Service
}

// InternshipsItem is a 42 API type
type InternshipsItem struct {
	AdministrationID         int         `json:"administration_id"`
	BreachAt                 interface{} `json:"breach_at"`
	CompanyAddress           string      `json:"company_address"`
	CompanyBossUserEmail     string      `json:"company_boss_user_email"`
	CompanyBossUserFirstName string      `json:"company_boss_user_first_name"`
	CompanyBossUserLastName  string      `json:"company_boss_user_last_name"`
	CompanyBossUserPhone     string      `json:"company_boss_user_phone"`
	CompanyCity              string      `json:"company_city"`
	CompanyCountry           string      `json:"company_country"`
	CompanyName              string      `json:"company_name"`
	CompanyPostal            string      `json:"company_postal"`
	CompanySiret             string      `json:"company_siret"`
	CompanyUserEmail         string      `json:"company_user_email"`
	CompanyUserFirstName     string      `json:"company_user_first_name"`
	CompanyUserLastName      string      `json:"company_user_last_name"`
	CompanyUserPhone         string      `json:"company_user_phone"`
	CompanyUserPost          string      `json:"company_user_post"`
	ContractType             string      `json:"contract_type"`
	Convention               struct {
		Convention struct {
			URL string `json:"url"`
		} `json:"convention"`
	} `json:"convention"`
	Currency          string      `json:"currency"`
	Days              string      `json:"days"`
	Duration          int         `json:"duration"`
	EndAt             time.Time   `json:"end_at"`
	ID                int         `json:"id"`
	InternshipAddress string      `json:"internship_address"`
	InternshipCity    string      `json:"internship_city"`
	InternshipCountry string      `json:"internship_country"`
	InternshipPostal  string      `json:"internship_postal"`
	LanguageID        int         `json:"language_id"`
	Movement          interface{} `json:"movement"`
	NbDays            int         `json:"nb_days"`
	NbHours           int         `json:"nb_hours"`
	OfferID           interface{} `json:"offer_id"`
	Salary            int         `json:"salary"`
	StartAt           time.Time   `json:"start_at"`
	State             string      `json:"state"`
	Subject           string      `json:"subject"`
	User              struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"user"`
	UserAddress string `json:"user_address"`
	UserCity    string `json:"user_city"`
	UserCountry string `json:"user_country"`
	UserPostal  string `json:"user_postal"`
}

// InternshipsListCall description:
type InternshipsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListInternshipsResponse represents a list response.
type ListInternshipsResponse struct {
	Internships []*InternshipsItem
}

// List returns a InternshipsListCall which can request data from the 42 API.
func (s *InternshipsService) List() *InternshipsListCall {
	return &InternshipsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Internships matching the specified setting.
func (c *InternshipsListCall) P(key string, values []string) *InternshipsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a InternshipsListCall request call. Exactly one of *ListInternshipsResponse or error will be non-nil.
func (c *InternshipsListCall) Do() (*ListInternshipsResponse, error) {
	urls := c.s.baseURL + "/internships/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListInternshipsResponse{}
	ret.Internships = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Internships); err != nil {
		return nil, err
	}
	return ret, nil
}

// InternshipsGetCall description:
type InternshipsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified InternshipsItem.
func (s *InternshipsService) Get(ID string) *InternshipsGetCall {
	return &InternshipsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a InternshipsGetCall request call. Exactly one of *InternshipsItem or error will be non-nil.
func (c *InternshipsGetCall) Do() (*InternshipsItem, error) {
	urls := c.s.baseURL + "/internships/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &InternshipsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// Internships description: The internship
func Internships(s *Service) *InternshipsService {
	return &InternshipsService{s: s}
}

// LanguagesService handles 42 API Languages.
type LanguagesService struct {
	s *Service
}

// LanguagesItem is a 42 API type
type LanguagesItem struct {
	ID         int    `json:"id"`
	IDentifier string `json:"identifier"`
	Name       string `json:"name"`
}

// LanguagesListCall description:
type LanguagesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListLanguagesResponse represents a list response.
type ListLanguagesResponse struct {
	Languages []*LanguagesItem
}

// List returns a LanguagesListCall which can request data from the 42 API.
func (s *LanguagesService) List() *LanguagesListCall {
	return &LanguagesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Languages matching the specified setting.
func (c *LanguagesListCall) P(key string, values []string) *LanguagesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a LanguagesListCall request call. Exactly one of *ListLanguagesResponse or error will be non-nil.
func (c *LanguagesListCall) Do() (*ListLanguagesResponse, error) {
	urls := c.s.baseURL + "/languages/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListLanguagesResponse{}
	ret.Languages = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Languages); err != nil {
		return nil, err
	}
	return ret, nil
}

// LanguagesGetCall description:
type LanguagesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified LanguagesItem.
func (s *LanguagesService) Get(ID string) *LanguagesGetCall {
	return &LanguagesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a LanguagesGetCall request call. Exactly one of *LanguagesItem or error will be non-nil.
func (c *LanguagesGetCall) Do() (*LanguagesItem, error) {
	urls := c.s.baseURL + "/languages/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &LanguagesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Languages description: The language
func Languages(s *Service) *LanguagesService {
	return &LanguagesService{s: s}
}

// LanguagesUsersService handles 42 API LanguagesUsers.
type LanguagesUsersService struct {
	s *Service
}

// LanguagesUsersItem is a 42 API type
type LanguagesUsersItem struct {
	CreatedAt  time.Time `json:"created_at"`
	ID         int       `json:"id"`
	LanguageID int       `json:"language_id"`
	Position   int       `json:"position"`
	UserID     int       `json:"user_id"`
}

// LanguagesUsersListCall description:
type LanguagesUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListLanguagesUsersResponse represents a list response.
type ListLanguagesUsersResponse struct {
	LanguagesUsers []*LanguagesUsersItem
}

// List returns a LanguagesUsersListCall which can request data from the 42 API.
func (s *LanguagesUsersService) List() *LanguagesUsersListCall {
	return &LanguagesUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return LanguagesUsers matching the specified setting.
func (c *LanguagesUsersListCall) P(key string, values []string) *LanguagesUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a LanguagesUsersListCall request call. Exactly one of *ListLanguagesUsersResponse or error will be non-nil.
func (c *LanguagesUsersListCall) Do() (*ListLanguagesUsersResponse, error) {
	urls := c.s.baseURL + "/languages_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListLanguagesUsersResponse{}
	ret.LanguagesUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).LanguagesUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// LanguagesUsersGetCall description:
type LanguagesUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified LanguagesUsersItem.
func (s *LanguagesUsersService) Get(ID string) *LanguagesUsersGetCall {
	return &LanguagesUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a LanguagesUsersGetCall request call. Exactly one of *LanguagesUsersItem or error will be non-nil.
func (c *LanguagesUsersGetCall) Do() (*LanguagesUsersItem, error) {
	urls := c.s.baseURL + "/languages_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &LanguagesUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// LanguagesUsers description: The languages of a user
func LanguagesUsers(s *Service) *LanguagesUsersService {
	return &LanguagesUsersService{s: s}
}

// LevelsService handles 42 API Levels.
type LevelsService struct {
	s *Service
}

// LevelsItem is a 42 API type
type LevelsItem struct {
}

// LevelsListCall description:
type LevelsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListLevelsResponse represents a list response.
type ListLevelsResponse struct {
	Levels []*LevelsItem
}

// List returns a LevelsListCall which can request data from the 42 API.
func (s *LevelsService) List() *LevelsListCall {
	return &LevelsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Levels matching the specified setting.
func (c *LevelsListCall) P(key string, values []string) *LevelsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a LevelsListCall request call. Exactly one of *ListLevelsResponse or error will be non-nil.
func (c *LevelsListCall) Do() (*ListLevelsResponse, error) {
	urls := c.s.baseURL + "/levels/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListLevelsResponse{}
	ret.Levels = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Levels); err != nil {
		return nil, err
	}
	return ret, nil
}

// Levels description: A level indicator for a cursus.
func Levels(s *Service) *LevelsService {
	return &LevelsService{s: s}
}

// LocationsService handles 42 API Locations.
type LocationsService struct {
	s *Service
}

// LocationsItem is a 42 API type
type LocationsItem struct {
	BeginAt  time.Time   `json:"begin_at"`
	CampusID int         `json:"campus_id"`
	EndAt    time.Time   `json:"end_at"`
	Floor    interface{} `json:"floor"`
	Host     string      `json:"host"`
	ID       int         `json:"id"`
	Post     interface{} `json:"post"`
	Primary  bool        `json:"primary"`
	Row      interface{} `json:"row"`
	User     struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"user"`
}

// LocationsListCall description:
type LocationsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListLocationsResponse represents a list response.
type ListLocationsResponse struct {
	Locations []*LocationsItem
}

// List returns a LocationsListCall which can request data from the 42 API.
func (s *LocationsService) List() *LocationsListCall {
	return &LocationsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Locations matching the specified setting.
func (c *LocationsListCall) P(key string, values []string) *LocationsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a LocationsListCall request call. Exactly one of *ListLocationsResponse or error will be non-nil.
func (c *LocationsListCall) Do() (*ListLocationsResponse, error) {
	urls := c.s.baseURL + "/locations/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListLocationsResponse{}
	ret.Locations = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Locations); err != nil {
		return nil, err
	}
	return ret, nil
}

// LocationsGetCall description:
type LocationsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified LocationsItem.
func (s *LocationsService) Get(ID string) *LocationsGetCall {
	return &LocationsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a LocationsGetCall request call. Exactly one of *LocationsItem or error will be non-nil.
func (c *LocationsGetCall) Do() (*LocationsItem, error) {
	urls := c.s.baseURL + "/locations/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &LocationsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Locations description: The location of an user in a campus
func Locations(s *Service) *LocationsService {
	return &LocationsService{s: s}
}

// MailingsService handles 42 API Mailings.
type MailingsService struct {
	s *Service
}

// MailingsItem is a 42 API type
type MailingsItem struct {
	Attachment  interface{}   `json:"attachment"`
	Attachments interface{}   `json:"attachments"`
	Bcc         []interface{} `json:"bcc"`
	Cc          []interface{} `json:"cc"`
	Content     string        `json:"content"`
	CreatedAt   time.Time     `json:"created_at"`
	From        string        `json:"from"`
	HTMLContent string        `json:"html_content"`
	ID          int           `json:"id"`
	IDentifier  string        `json:"identifier"`
	Meta        struct {
		ID        int       `json:"id"`
		Date      time.Time `json:"date"`
		Name      string    `json:"name"`
		User      string    `json:"user"`
		OtherUser string    `json:"other_user"`
	} `json:"meta"`
	Subject   string      `json:"subject"`
	Subtitle  interface{} `json:"subtitle"`
	Title     string      `json:"title"`
	To        []string    `json:"to"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// MailingsListCall description:
type MailingsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListMailingsResponse represents a list response.
type ListMailingsResponse struct {
	Mailings []*MailingsItem
}

// List returns a MailingsListCall which can request data from the 42 API.
func (s *MailingsService) List() *MailingsListCall {
	return &MailingsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Mailings matching the specified setting.
func (c *MailingsListCall) P(key string, values []string) *MailingsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a MailingsListCall request call. Exactly one of *ListMailingsResponse or error will be non-nil.
func (c *MailingsListCall) Do() (*ListMailingsResponse, error) {
	urls := c.s.baseURL + "/mailings/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListMailingsResponse{}
	ret.Mailings = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Mailings); err != nil {
		return nil, err
	}
	return ret, nil
}

// MailingsGetCall description:
type MailingsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified MailingsItem.
func (s *MailingsService) Get(ID string) *MailingsGetCall {
	return &MailingsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a MailingsGetCall request call. Exactly one of *MailingsItem or error will be non-nil.
func (c *MailingsGetCall) Do() (*MailingsItem, error) {
	urls := c.s.baseURL + "/mailings/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &MailingsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Mailings description: Mails from and between 42 entities
func Mailings(s *Service) *MailingsService {
	return &MailingsService{s: s}
}

// NotesService handles 42 API Notes.
type NotesService struct {
	s *Service
}

// NotesItem is a 42 API type
type NotesItem struct {
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	FromUser  struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"from_user"`
	ID      int    `json:"id"`
	Subject string `json:"subject"`
	User    struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"user"`
}

// NotesListCall description:
type NotesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListNotesResponse represents a list response.
type ListNotesResponse struct {
	Notes []*NotesItem
}

// List returns a NotesListCall which can request data from the 42 API.
func (s *NotesService) List() *NotesListCall {
	return &NotesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Notes matching the specified setting.
func (c *NotesListCall) P(key string, values []string) *NotesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a NotesListCall request call. Exactly one of *ListNotesResponse or error will be non-nil.
func (c *NotesListCall) Do() (*ListNotesResponse, error) {
	urls := c.s.baseURL + "/notes/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListNotesResponse{}
	ret.Notes = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Notes); err != nil {
		return nil, err
	}
	return ret, nil
}

// NotesGetCall description:
type NotesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified NotesItem.
func (s *NotesService) Get(ID string) *NotesGetCall {
	return &NotesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a NotesGetCall request call. Exactly one of *NotesItem or error will be non-nil.
func (c *NotesGetCall) Do() (*NotesItem, error) {
	urls := c.s.baseURL + "/notes/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &NotesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Notes description: A note for an user
func Notes(s *Service) *NotesService {
	return &NotesService{s: s}
}

// NotionsService handles 42 API Notions.
type NotionsService struct {
	s *Service
}

// NotionsItem is a 42 API type
type NotionsItem struct {
	CreatedAt time.Time `json:"created_at"`
	Cursus    []struct {
		Name      string    `json:"name"`
		Slug      string    `json:"slug"`
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"cursus"`
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Slug       string        `json:"slug"`
	Subnotions []interface{} `json:"subnotions"`
	Tags       []struct {
		Name string `json:"name"`
		Kind string `json:"kind"`
		ID   int    `json:"id"`
	} `json:"tags"`
}

// NotionsListCall description:
type NotionsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListNotionsResponse represents a list response.
type ListNotionsResponse struct {
	Notions []*NotionsItem
}

// List returns a NotionsListCall which can request data from the 42 API.
func (s *NotionsService) List() *NotionsListCall {
	return &NotionsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Notions matching the specified setting.
func (c *NotionsListCall) P(key string, values []string) *NotionsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a NotionsListCall request call. Exactly one of *ListNotionsResponse or error will be non-nil.
func (c *NotionsListCall) Do() (*ListNotionsResponse, error) {
	urls := c.s.baseURL + "/notions/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListNotionsResponse{}
	ret.Notions = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Notions); err != nil {
		return nil, err
	}
	return ret, nil
}

// NotionsGetCall description:
type NotionsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified NotionsItem.
func (s *NotionsService) Get(ID string) *NotionsGetCall {
	return &NotionsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a NotionsGetCall request call. Exactly one of *NotionsItem or error will be non-nil.
func (c *NotionsGetCall) Do() (*NotionsItem, error) {
	urls := c.s.baseURL + "/notions/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &NotionsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Notions description: The elearning notion in a cursus
func Notions(s *Service) *NotionsService {
	return &NotionsService{s: s}
}

// ParamsProjectSessionsRulesService handles 42 API ParamsProjectSessionsRules.
type ParamsProjectSessionsRulesService struct {
	s *Service
}

// ParamsProjectSessionsRulesItem is a 42 API type
type ParamsProjectSessionsRulesItem struct {
}

// ParamsProjectSessionsRulesListCall description:
type ParamsProjectSessionsRulesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListParamsProjectSessionsRulesResponse represents a list response.
type ListParamsProjectSessionsRulesResponse struct {
	ParamsProjectSessionsRules []*ParamsProjectSessionsRulesItem
}

// List returns a ParamsProjectSessionsRulesListCall which can request data from the 42 API.
func (s *ParamsProjectSessionsRulesService) List() *ParamsProjectSessionsRulesListCall {
	return &ParamsProjectSessionsRulesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return ParamsProjectSessionsRules matching the specified setting.
func (c *ParamsProjectSessionsRulesListCall) P(key string, values []string) *ParamsProjectSessionsRulesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ParamsProjectSessionsRulesListCall request call. Exactly one of *ListParamsProjectSessionsRulesResponse or error will be non-nil.
func (c *ParamsProjectSessionsRulesListCall) Do() (*ListParamsProjectSessionsRulesResponse, error) {
	urls := c.s.baseURL + "/params_project_sessions_rules/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListParamsProjectSessionsRulesResponse{}
	ret.ParamsProjectSessionsRules = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).ParamsProjectSessionsRules); err != nil {
		return nil, err
	}
	return ret, nil
}

// ParamsProjectSessionsRulesGetCall description:
type ParamsProjectSessionsRulesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified ParamsProjectSessionsRulesItem.
func (s *ParamsProjectSessionsRulesService) Get(ID string) *ParamsProjectSessionsRulesGetCall {
	return &ParamsProjectSessionsRulesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a ParamsProjectSessionsRulesGetCall request call. Exactly one of *ParamsProjectSessionsRulesItem or error will be non-nil.
func (c *ParamsProjectSessionsRulesGetCall) Do() (*ParamsProjectSessionsRulesItem, error) {
	urls := c.s.baseURL + "/params_project_sessions_rules/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ParamsProjectSessionsRulesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// ParamsProjectSessionsRules description: The value of a parameter for a project sessions rule.
func ParamsProjectSessionsRules(s *Service) *ParamsProjectSessionsRulesService {
	return &ParamsProjectSessionsRulesService{s: s}
}

// PartnershipsService handles 42 API Partnerships.
type PartnershipsService struct {
	s *Service
}

// PartnershipsItem is a 42 API type
type PartnershipsItem struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	PartnershipsUsersURL string `json:"partnerships_users_url"`
	Slug                 string `json:"slug"`
	Tier                 int    `json:"tier"`
	URL                  string `json:"url"`
}

// PartnershipsListCall description:
type PartnershipsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListPartnershipsResponse represents a list response.
type ListPartnershipsResponse struct {
	Partnerships []*PartnershipsItem
}

// List returns a PartnershipsListCall which can request data from the 42 API.
func (s *PartnershipsService) List() *PartnershipsListCall {
	return &PartnershipsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Partnerships matching the specified setting.
func (c *PartnershipsListCall) P(key string, values []string) *PartnershipsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a PartnershipsListCall request call. Exactly one of *ListPartnershipsResponse or error will be non-nil.
func (c *PartnershipsListCall) Do() (*ListPartnershipsResponse, error) {
	urls := c.s.baseURL + "/partnerships/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListPartnershipsResponse{}
	ret.Partnerships = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Partnerships); err != nil {
		return nil, err
	}
	return ret, nil
}

// PartnershipsGetCall description:
type PartnershipsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified PartnershipsItem.
func (s *PartnershipsService) Get(ID string) *PartnershipsGetCall {
	return &PartnershipsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a PartnershipsGetCall request call. Exactly one of *PartnershipsItem or error will be non-nil.
func (c *PartnershipsGetCall) Do() (*PartnershipsItem, error) {
	urls := c.s.baseURL + "/partnerships/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &PartnershipsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Partnerships description: Pedagogic partnerships
func Partnerships(s *Service) *PartnershipsService {
	return &PartnershipsService{s: s}
}

// PartnershipsUsersService handles 42 API PartnershipsUsers.
type PartnershipsUsersService struct {
	s *Service
}

// PartnershipsUsersItem is a 42 API type
type PartnershipsUsersItem struct {
	FinalMark     interface{} `json:"final_mark"`
	ID            int         `json:"id"`
	PartnershipID int         `json:"partnership_id"`
	User          struct {
		Login string `json:"login"`
		URL   string `json:"url"`
		ID    int    `json:"id"`
	} `json:"user"`
}

// PartnershipsUsersListCall description:
type PartnershipsUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListPartnershipsUsersResponse represents a list response.
type ListPartnershipsUsersResponse struct {
	PartnershipsUsers []*PartnershipsUsersItem
}

// List returns a PartnershipsUsersListCall which can request data from the 42 API.
func (s *PartnershipsUsersService) List() *PartnershipsUsersListCall {
	return &PartnershipsUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return PartnershipsUsers matching the specified setting.
func (c *PartnershipsUsersListCall) P(key string, values []string) *PartnershipsUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a PartnershipsUsersListCall request call. Exactly one of *ListPartnershipsUsersResponse or error will be non-nil.
func (c *PartnershipsUsersListCall) Do() (*ListPartnershipsUsersResponse, error) {
	urls := c.s.baseURL + "/partnerships_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListPartnershipsUsersResponse{}
	ret.PartnershipsUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).PartnershipsUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// PartnershipsUsersGetCall description:
type PartnershipsUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified PartnershipsUsersItem.
func (s *PartnershipsUsersService) Get(ID string) *PartnershipsUsersGetCall {
	return &PartnershipsUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a PartnershipsUsersGetCall request call. Exactly one of *PartnershipsUsersItem or error will be non-nil.
func (c *PartnershipsUsersGetCall) Do() (*PartnershipsUsersItem, error) {
	urls := c.s.baseURL + "/partnerships_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &PartnershipsUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// PartnershipsUsers description: Users doing a partnership
func PartnershipsUsers(s *Service) *PartnershipsUsersService {
	return &PartnershipsUsersService{s: s}
}

// PatronagesService handles 42 API Patronages.
type PatronagesService struct {
	s *Service
}

// PatronagesItem is a 42 API type
type PatronagesItem struct {
	CreatedAt time.Time `json:"created_at"`
	Godfather struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"godfather"`
	GodfatherID int       `json:"godfather_id"`
	ID          int       `json:"id"`
	Ongoing     bool      `json:"ongoing"`
	UpdatedAt   time.Time `json:"updated_at"`
	User        struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"user"`
	UserID int `json:"user_id"`
}

// PatronagesListCall description:
type PatronagesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListPatronagesResponse represents a list response.
type ListPatronagesResponse struct {
	Patronages []*PatronagesItem
}

// List returns a PatronagesListCall which can request data from the 42 API.
func (s *PatronagesService) List() *PatronagesListCall {
	return &PatronagesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Patronages matching the specified setting.
func (c *PatronagesListCall) P(key string, values []string) *PatronagesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a PatronagesListCall request call. Exactly one of *ListPatronagesResponse or error will be non-nil.
func (c *PatronagesListCall) Do() (*ListPatronagesResponse, error) {
	urls := c.s.baseURL + "/patronages/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListPatronagesResponse{}
	ret.Patronages = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Patronages); err != nil {
		return nil, err
	}
	return ret, nil
}

// PatronagesGetCall description:
type PatronagesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified PatronagesItem.
func (s *PatronagesService) Get(ID string) *PatronagesGetCall {
	return &PatronagesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a PatronagesGetCall request call. Exactly one of *PatronagesItem or error will be non-nil.
func (c *PatronagesGetCall) Do() (*PatronagesItem, error) {
	urls := c.s.baseURL + "/patronages/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &PatronagesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Patronages description: A patronage between two users
func Patronages(s *Service) *PatronagesService {
	return &PatronagesService{s: s}
}

// PatronagesReportsService handles 42 API PatronagesReports.
type PatronagesReportsService struct {
	s *Service
}

// PatronagesReportsItem is a 42 API type
type PatronagesReportsItem struct {
	Answers   []interface{} `json:"answers"`
	BeginAt   time.Time     `json:"begin_at"`
	CreatedAt time.Time     `json:"created_at"`
	ID        int           `json:"id"`
	Patronage struct {
		GodfatherID int       `json:"godfather_id"`
		Ongoing     bool      `json:"ongoing"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		ID          int       `json:"id"`
		UserID      int       `json:"user_id"`
	} `json:"patronage"`
	PatronageID int `json:"patronage_id"`
	Report      struct {
		Comment        string    `json:"comment"`
		UpdatedAt      time.Time `json:"updated_at"`
		Slug           string    `json:"slug"`
		DelayDays      int       `json:"delay_days"`
		CreatedAt      time.Time `json:"created_at"`
		ID             int       `json:"id"`
		Name           string    `json:"name"`
		IntroductionMd string    `json:"introduction_md"`
		DisclaimerMd   string    `json:"disclaimer_md"`
		GuidelinesMd   string    `json:"guidelines_md"`
	} `json:"report"`
	ReportID  int       `json:"report_id"`
	UpdatedAt time.Time `json:"updated_at"`
	User      struct {
		Login string `json:"login"`
		URL   string `json:"url"`
		ID    int    `json:"id"`
	} `json:"user"`
	UserID      int         `json:"user_id"`
	ValidatedAt interface{} `json:"validated_at"`
}

// PatronagesReportsListCall description:
type PatronagesReportsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListPatronagesReportsResponse represents a list response.
type ListPatronagesReportsResponse struct {
	PatronagesReports []*PatronagesReportsItem
}

// List returns a PatronagesReportsListCall which can request data from the 42 API.
func (s *PatronagesReportsService) List() *PatronagesReportsListCall {
	return &PatronagesReportsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return PatronagesReports matching the specified setting.
func (c *PatronagesReportsListCall) P(key string, values []string) *PatronagesReportsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a PatronagesReportsListCall request call. Exactly one of *ListPatronagesReportsResponse or error will be non-nil.
func (c *PatronagesReportsListCall) Do() (*ListPatronagesReportsResponse, error) {
	urls := c.s.baseURL + "/patronages_reports/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListPatronagesReportsResponse{}
	ret.PatronagesReports = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).PatronagesReports); err != nil {
		return nil, err
	}
	return ret, nil
}

// PatronagesReportsGetCall description:
type PatronagesReportsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified PatronagesReportsItem.
func (s *PatronagesReportsService) Get(ID string) *PatronagesReportsGetCall {
	return &PatronagesReportsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a PatronagesReportsGetCall request call. Exactly one of *PatronagesReportsItem or error will be non-nil.
func (c *PatronagesReportsGetCall) Do() (*PatronagesReportsItem, error) {
	urls := c.s.baseURL + "/patronages_reports/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &PatronagesReportsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// PatronagesReports description: A report for a patronage
func PatronagesReports(s *Service) *PatronagesReportsService {
	return &PatronagesReportsService{s: s}
}

// PoolsService handles 42 API Pools.
type PoolsService struct {
	s *Service
}

// PoolsItem is a 42 API type
type PoolsItem struct {
	CampusID      int `json:"campus_id"`
	CurrentPoints int `json:"current_points"`
	CursusID      int `json:"cursus_id"`
	ID            int `json:"id"`
	MaxPoints     int `json:"max_points"`
}

// PoolsListCall description:
type PoolsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListPoolsResponse represents a list response.
type ListPoolsResponse struct {
	Pools []*PoolsItem
}

// List returns a PoolsListCall which can request data from the 42 API.
func (s *PoolsService) List() *PoolsListCall {
	return &PoolsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Pools matching the specified setting.
func (c *PoolsListCall) P(key string, values []string) *PoolsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a PoolsListCall request call. Exactly one of *ListPoolsResponse or error will be non-nil.
func (c *PoolsListCall) Do() (*ListPoolsResponse, error) {
	urls := c.s.baseURL + "/pools/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListPoolsResponse{}
	ret.Pools = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Pools); err != nil {
		return nil, err
	}
	return ret, nil
}

// PoolsGetCall description:
type PoolsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified PoolsItem.
func (s *PoolsService) Get(ID string) *PoolsGetCall {
	return &PoolsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a PoolsGetCall request call. Exactly one of *PoolsItem or error will be non-nil.
func (c *PoolsGetCall) Do() (*PoolsItem, error) {
	urls := c.s.baseURL + "/pools/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &PoolsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// Pools description: The pool of evaluation points.
func Pools(s *Service) *PoolsService {
	return &PoolsService{s: s}
}

// ProductsService handles 42 API Products.
type ProductsService struct {
	s *Service
}

// ProductsItem is a 42 API type
type ProductsItem struct {
	BeginAt     interface{} `json:"begin_at"`
	CategoryID  int         `json:"category_id"`
	CreatedAt   time.Time   `json:"created_at"`
	Description string      `json:"description"`
	EndAt       interface{} `json:"end_at"`
	ID          int         `json:"id"`
	Image       struct {
		Thumb struct {
			URL string `json:"url"`
		} `json:"thumb"`
		URL string `json:"url"`
	} `json:"image"`
	IsUniq          bool      `json:"is_uniq"`
	Kind            string    `json:"kind"`
	Name            string    `json:"name"`
	OneTimePurchase bool      `json:"one_time_purchase"`
	Price           int       `json:"price"`
	Quantity        int       `json:"quantity"`
	Slug            string    `json:"slug"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// ProductsListCall description:
type ProductsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListProductsResponse represents a list response.
type ListProductsResponse struct {
	Products []*ProductsItem
}

// List returns a ProductsListCall which can request data from the 42 API.
func (s *ProductsService) List() *ProductsListCall {
	return &ProductsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Products matching the specified setting.
func (c *ProductsListCall) P(key string, values []string) *ProductsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ProductsListCall request call. Exactly one of *ListProductsResponse or error will be non-nil.
func (c *ProductsListCall) Do() (*ListProductsResponse, error) {
	urls := c.s.baseURL + "/products/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListProductsResponse{}
	ret.Products = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Products); err != nil {
		return nil, err
	}
	return ret, nil
}

// ProductsGetCall description:
type ProductsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified ProductsItem.
func (s *ProductsService) Get(ID string) *ProductsGetCall {
	return &ProductsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a ProductsGetCall request call. Exactly one of *ProductsItem or error will be non-nil.
func (c *ProductsGetCall) Do() (*ProductsItem, error) {
	urls := c.s.baseURL + "/products/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ProductsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Products description: Products are sold on the intranet shop
func Products(s *Service) *ProductsService {
	return &ProductsService{s: s}
}

// ProjectDataService handles 42 API ProjectData.
type ProjectDataService struct {
	s *Service
}

// ProjectDataItem is a 42 API type
type ProjectDataItem struct {
	By               []interface{} `json:"by"`
	Coordinates      []int         `json:"coordinates"`
	ID               int           `json:"id"`
	Kind             string        `json:"kind"`
	ProjectSessionID int           `json:"project_session_id"`
}

// ProjectDataListCall description:
type ProjectDataListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListProjectDataResponse represents a list response.
type ListProjectDataResponse struct {
	ProjectData []*ProjectDataItem
}

// List returns a ProjectDataListCall which can request data from the 42 API.
func (s *ProjectDataService) List() *ProjectDataListCall {
	return &ProjectDataListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return ProjectData matching the specified setting.
func (c *ProjectDataListCall) P(key string, values []string) *ProjectDataListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ProjectDataListCall request call. Exactly one of *ListProjectDataResponse or error will be non-nil.
func (c *ProjectDataListCall) Do() (*ListProjectDataResponse, error) {
	urls := c.s.baseURL + "/project_data/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListProjectDataResponse{}
	ret.ProjectData = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).ProjectData); err != nil {
		return nil, err
	}
	return ret, nil
}

// ProjectDataGetCall description:
type ProjectDataGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified ProjectDataItem.
func (s *ProjectDataService) Get(ID string) *ProjectDataGetCall {
	return &ProjectDataGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a ProjectDataGetCall request call. Exactly one of *ProjectDataItem or error will be non-nil.
func (c *ProjectDataGetCall) Do() (*ProjectDataItem, error) {
	urls := c.s.baseURL + "/project_data/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ProjectDataItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// ProjectData description: Project data for the graph
func ProjectData(s *Service) *ProjectDataService {
	return &ProjectDataService{s: s}
}

// ProjectSessionsService handles 42 API ProjectSessions.
type ProjectSessionsService struct {
	s *Service
}

// ProjectSessionsItem is a 42 API type
type ProjectSessionsItem struct {
	BeginAt      interface{} `json:"begin_at"`
	Campus       interface{} `json:"campus"`
	CampusID     interface{} `json:"campus_id"`
	CreatedAt    time.Time   `json:"created_at"`
	Cursus       interface{} `json:"cursus"`
	CursusID     interface{} `json:"cursus_id"`
	DurationDays interface{} `json:"duration_days"`
	EndAt        interface{} `json:"end_at"`
	EstimateTime int         `json:"estimate_time"`
	Evaluations  []struct {
		ID   int    `json:"id"`
		Kind string `json:"kind"`
	} `json:"evaluations"`
	ID              int         `json:"id"`
	IsSubscriptable bool        `json:"is_subscriptable"`
	MaxPeople       interface{} `json:"max_people"`
	Project         struct {
		Name        string        `json:"name"`
		Description string        `json:"description"`
		Children    []interface{} `json:"children"`
		Tier        int           `json:"tier"`
		Exam        bool          `json:"exam"`
		ID          int           `json:"id"`
		Slug        string        `json:"slug"`
		Parent      interface{}   `json:"parent"`
		Objectives  []string      `json:"objectives"`
		Attachments []interface{} `json:"attachments"`
		CreatedAt   time.Time     `json:"created_at"`
		UpdatedAt   time.Time     `json:"updated_at"`
	} `json:"project"`
	ProjectID int `json:"project_id"`
	Scales    []struct {
		ID               int  `json:"id"`
		CorrectionNumber int  `json:"correction_number"`
		IsPrimary        bool `json:"is_primary"`
	} `json:"scales"`
	Solo             bool          `json:"solo"`
	TeamBehaviour    string        `json:"team_behaviour"`
	TerminatingAfter interface{}   `json:"terminating_after"`
	UpdatedAt        time.Time     `json:"updated_at"`
	Uploads          []interface{} `json:"uploads"`
}

// ProjectSessionsListCall description:
type ProjectSessionsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListProjectSessionsResponse represents a list response.
type ListProjectSessionsResponse struct {
	ProjectSessions []*ProjectSessionsItem
}

// List returns a ProjectSessionsListCall which can request data from the 42 API.
func (s *ProjectSessionsService) List() *ProjectSessionsListCall {
	return &ProjectSessionsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return ProjectSessions matching the specified setting.
func (c *ProjectSessionsListCall) P(key string, values []string) *ProjectSessionsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ProjectSessionsListCall request call. Exactly one of *ListProjectSessionsResponse or error will be non-nil.
func (c *ProjectSessionsListCall) Do() (*ListProjectSessionsResponse, error) {
	urls := c.s.baseURL + "/project_sessions/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListProjectSessionsResponse{}
	ret.ProjectSessions = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).ProjectSessions); err != nil {
		return nil, err
	}
	return ret, nil
}

// ProjectSessionsGetCall description:
type ProjectSessionsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified ProjectSessionsItem.
func (s *ProjectSessionsService) Get(ID string) *ProjectSessionsGetCall {
	return &ProjectSessionsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a ProjectSessionsGetCall request call. Exactly one of *ProjectSessionsItem or error will be non-nil.
func (c *ProjectSessionsGetCall) Do() (*ProjectSessionsItem, error) {
	urls := c.s.baseURL + "/project_sessions/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ProjectSessionsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// ProjectSessions description: A project session defines a particular behaviour for a project, based on the cursus and / or the campus .
func ProjectSessions(s *Service) *ProjectSessionsService {
	return &ProjectSessionsService{s: s}
}

// ProjectSessionsRulesService handles 42 API ProjectSessionsRules.
type ProjectSessionsRulesService struct {
	s *Service
}

// ProjectSessionsRulesItem is a 42 API type
type ProjectSessionsRulesItem struct {
}

// ProjectSessionsRulesListCall description:
type ProjectSessionsRulesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListProjectSessionsRulesResponse represents a list response.
type ListProjectSessionsRulesResponse struct {
	ProjectSessionsRules []*ProjectSessionsRulesItem
}

// List returns a ProjectSessionsRulesListCall which can request data from the 42 API.
func (s *ProjectSessionsRulesService) List() *ProjectSessionsRulesListCall {
	return &ProjectSessionsRulesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return ProjectSessionsRules matching the specified setting.
func (c *ProjectSessionsRulesListCall) P(key string, values []string) *ProjectSessionsRulesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ProjectSessionsRulesListCall request call. Exactly one of *ListProjectSessionsRulesResponse or error will be non-nil.
func (c *ProjectSessionsRulesListCall) Do() (*ListProjectSessionsRulesResponse, error) {
	urls := c.s.baseURL + "/project_sessions_rules/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListProjectSessionsRulesResponse{}
	ret.ProjectSessionsRules = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).ProjectSessionsRules); err != nil {
		return nil, err
	}
	return ret, nil
}

// ProjectSessionsRulesGetCall description:
type ProjectSessionsRulesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified ProjectSessionsRulesItem.
func (s *ProjectSessionsRulesService) Get(ID string) *ProjectSessionsRulesGetCall {
	return &ProjectSessionsRulesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a ProjectSessionsRulesGetCall request call. Exactly one of *ProjectSessionsRulesItem or error will be non-nil.
func (c *ProjectSessionsRulesGetCall) Do() (*ProjectSessionsRulesItem, error) {
	urls := c.s.baseURL + "/project_sessions_rules/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ProjectSessionsRulesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// ProjectSessionsRules description: A rule linked to a project session.
func ProjectSessionsRules(s *Service) *ProjectSessionsRulesService {
	return &ProjectSessionsRulesService{s: s}
}

// ProjectSessionsSkillsService handles 42 API ProjectSessionsSkills.
type ProjectSessionsSkillsService struct {
	s *Service
}

// ProjectSessionsSkillsItem is a 42 API type
type ProjectSessionsSkillsItem struct {
	CreatedAt        time.Time `json:"created_at"`
	ID               int       `json:"id"`
	ProjectSessionID int       `json:"project_session_id"`
	SkillID          int       `json:"skill_id"`
	UpdatedAt        time.Time `json:"updated_at"`
	Value            int       `json:"value"`
}

// ProjectSessionsSkillsListCall description:
type ProjectSessionsSkillsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListProjectSessionsSkillsResponse represents a list response.
type ListProjectSessionsSkillsResponse struct {
	ProjectSessionsSkills []*ProjectSessionsSkillsItem
}

// List returns a ProjectSessionsSkillsListCall which can request data from the 42 API.
func (s *ProjectSessionsSkillsService) List() *ProjectSessionsSkillsListCall {
	return &ProjectSessionsSkillsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return ProjectSessionsSkills matching the specified setting.
func (c *ProjectSessionsSkillsListCall) P(key string, values []string) *ProjectSessionsSkillsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ProjectSessionsSkillsListCall request call. Exactly one of *ListProjectSessionsSkillsResponse or error will be non-nil.
func (c *ProjectSessionsSkillsListCall) Do() (*ListProjectSessionsSkillsResponse, error) {
	urls := c.s.baseURL + "/project_sessions_skills/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListProjectSessionsSkillsResponse{}
	ret.ProjectSessionsSkills = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).ProjectSessionsSkills); err != nil {
		return nil, err
	}
	return ret, nil
}

// ProjectSessionsSkillsGetCall description:
type ProjectSessionsSkillsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified ProjectSessionsSkillsItem.
func (s *ProjectSessionsSkillsService) Get(ID string) *ProjectSessionsSkillsGetCall {
	return &ProjectSessionsSkillsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a ProjectSessionsSkillsGetCall request call. Exactly one of *ProjectSessionsSkillsItem or error will be non-nil.
func (c *ProjectSessionsSkillsGetCall) Do() (*ProjectSessionsSkillsItem, error) {
	urls := c.s.baseURL + "/project_sessions_skills/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ProjectSessionsSkillsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// ProjectSessionsSkills description: A skill linked to a project session.
func ProjectSessionsSkills(s *Service) *ProjectSessionsSkillsService {
	return &ProjectSessionsSkillsService{s: s}
}

// ProjectsService handles 42 API Projects.
type ProjectsService struct {
	s *Service
}

// ProjectsItem is a 42 API type
type ProjectsItem struct {
	Attachments []interface{} `json:"attachments"`
	Campus      []struct {
		TimeZone string `json:"time_zone"`
		Language struct {
			ID         int       `json:"id"`
			Name       string    `json:"name"`
			IDentifier string    `json:"identifier"`
			CreatedAt  time.Time `json:"created_at"`
			UpdatedAt  time.Time `json:"updated_at"`
		} `json:"language"`
		UsersCount  int    `json:"users_count"`
		VogsphereID int    `json:"vogsphere_id"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
	} `json:"campus"`
	Children  []interface{} `json:"children"`
	CreatedAt time.Time     `json:"created_at"`
	Cursus    []struct {
		Slug      string    `json:"slug"`
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		Name      string    `json:"name"`
	} `json:"cursus"`
	Description     string      `json:"description"`
	Exam            bool        `json:"exam"`
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	Objectives      []string    `json:"objectives"`
	Parent          interface{} `json:"parent"`
	ProjectSessions []struct {
		Solo          bool        `json:"solo"`
		ProjectID     int         `json:"project_id"`
		TeamBehaviour string      `json:"team_behaviour"`
		ID            int         `json:"id"`
		CampusID      interface{} `json:"campus_id"`
		UpdatedAt     time.Time   `json:"updated_at"`
		Scales        []struct {
			ID               int  `json:"id"`
			CorrectionNumber int  `json:"correction_number"`
			IsPrimary        bool `json:"is_primary"`
		} `json:"scales"`
		DurationDays     interface{}   `json:"duration_days"`
		EstimateTime     int           `json:"estimate_time"`
		CursusID         interface{}   `json:"cursus_id"`
		CreatedAt        time.Time     `json:"created_at"`
		MaxPeople        interface{}   `json:"max_people"`
		Uploads          []interface{} `json:"uploads"`
		BeginAt          interface{}   `json:"begin_at"`
		TerminatingAfter interface{}   `json:"terminating_after"`
		IsSubscriptable  bool          `json:"is_subscriptable"`
		EndAt            interface{}   `json:"end_at"`
	} `json:"project_sessions"`
	Skills []struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"skills"`
	Slug string `json:"slug"`
	Tags []struct {
		Name string `json:"name"`
		Kind string `json:"kind"`
		ID   int    `json:"id"`
	} `json:"tags"`
	Tier      int           `json:"tier"`
	UpdatedAt time.Time     `json:"updated_at"`
	Videos    []interface{} `json:"videos"`
}

// ProjectsListCall description:
type ProjectsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListProjectsResponse represents a list response.
type ListProjectsResponse struct {
	Projects []*ProjectsItem
}

// List returns a ProjectsListCall which can request data from the 42 API.
func (s *ProjectsService) List() *ProjectsListCall {
	return &ProjectsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Projects matching the specified setting.
func (c *ProjectsListCall) P(key string, values []string) *ProjectsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ProjectsListCall request call. Exactly one of *ListProjectsResponse or error will be non-nil.
func (c *ProjectsListCall) Do() (*ListProjectsResponse, error) {
	urls := c.s.baseURL + "/projects/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListProjectsResponse{}
	ret.Projects = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Projects); err != nil {
		return nil, err
	}
	return ret, nil
}

// ProjectsGetCall description:
type ProjectsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified ProjectsItem.
func (s *ProjectsService) Get(ID string) *ProjectsGetCall {
	return &ProjectsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a ProjectsGetCall request call. Exactly one of *ProjectsItem or error will be non-nil.
func (c *ProjectsGetCall) Do() (*ProjectsItem, error) {
	urls := c.s.baseURL + "/projects/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ProjectsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Projects description: Pedagogic projects of a cursus
func Projects(s *Service) *ProjectsService {
	return &ProjectsService{s: s}
}

// ProjectsUsersService handles 42 API ProjectsUsers.
type ProjectsUsersService struct {
	s *Service
}

// ProjectsUsersItem is a 42 API type
type ProjectsUsersItem struct {
	CurrentTeamID int         `json:"current_team_id"`
	CursusIDs     []int       `json:"cursus_ids"`
	FinalMark     interface{} `json:"final_mark"`
	ID            int         `json:"id"`
	Occurrence    int         `json:"occurrence"`
	Project       struct {
		ID       int         `json:"id"`
		Name     string      `json:"name"`
		Slug     string      `json:"slug"`
		ParentID interface{} `json:"parent_id"`
	} `json:"project"`
	Status string `json:"status"`
	Teams  []struct {
		ProjectSessionID int         `json:"project_session_id"`
		Status           string      `json:"status"`
		Closed           bool        `json:"closed?"`
		RepoURL          interface{} `json:"repo_url"`
		ClosedAt         time.Time   `json:"closed_at"`
		LockedAt         time.Time   `json:"locked_at"`
		ID               int         `json:"id"`
		URL              string      `json:"url"`
		FinalMark        interface{} `json:"final_mark"`
		Locked           bool        `json:"locked?"`
		CreatedAt        time.Time   `json:"created_at"`
		TerminatingAt    interface{} `json:"terminating_at"`
		RepoUUID         string      `json:"repo_uuid"`
		Validated        interface{} `json:"validated?"`
		Name             string      `json:"name"`
		ProjectID        int         `json:"project_id"`
		UpdatedAt        time.Time   `json:"updated_at"`
		Users            []struct {
			Leader         bool   `json:"leader"`
			Occurrence     int    `json:"occurrence"`
			Validated      bool   `json:"validated"`
			ProjectsUserID int    `json:"projects_user_id"`
			ID             int    `json:"id"`
			Login          string `json:"login"`
			URL            string `json:"url"`
		} `json:"users"`
	} `json:"teams"`
	User struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"user"`
	Validated interface{} `json:"validated?"`
}

// ProjectsUsersListCall description:
type ProjectsUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListProjectsUsersResponse represents a list response.
type ListProjectsUsersResponse struct {
	ProjectsUsers []*ProjectsUsersItem
}

// List returns a ProjectsUsersListCall which can request data from the 42 API.
func (s *ProjectsUsersService) List() *ProjectsUsersListCall {
	return &ProjectsUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return ProjectsUsers matching the specified setting.
func (c *ProjectsUsersListCall) P(key string, values []string) *ProjectsUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ProjectsUsersListCall request call. Exactly one of *ListProjectsUsersResponse or error will be non-nil.
func (c *ProjectsUsersListCall) Do() (*ListProjectsUsersResponse, error) {
	urls := c.s.baseURL + "/projects_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListProjectsUsersResponse{}
	ret.ProjectsUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).ProjectsUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// ProjectsUsersGetCall description:
type ProjectsUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified ProjectsUsersItem.
func (s *ProjectsUsersService) Get(ID string) *ProjectsUsersGetCall {
	return &ProjectsUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a ProjectsUsersGetCall request call. Exactly one of *ProjectsUsersItem or error will be non-nil.
func (c *ProjectsUsersGetCall) Do() (*ProjectsUsersItem, error) {
	urls := c.s.baseURL + "/projects_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ProjectsUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// ProjectsUsers description: Users which did or are doing a project
func ProjectsUsers(s *Service) *ProjectsUsersService {
	return &ProjectsUsersService{s: s}
}

// QuestsService handles 42 API Quests.
type QuestsService struct {
	s *Service
}

// QuestsItem is a 42 API type
type QuestsItem struct {
	Campus    interface{} `json:"campus"`
	CampusID  interface{} `json:"campus_id"`
	CreatedAt time.Time   `json:"created_at"`
	Cursus    struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		Name      string    `json:"name"`
		Slug      string    `json:"slug"`
	} `json:"cursus"`
	CursusID     int         `json:"cursus_id"`
	Description  string      `json:"description"`
	Grade        interface{} `json:"grade"`
	GradeID      interface{} `json:"grade_id"`
	ID           int         `json:"id"`
	InternalName interface{} `json:"internal_name"`
	Kind         string      `json:"kind"`
	Name         string      `json:"name"`
	Position     int         `json:"position"`
	Slug         string      `json:"slug"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

// QuestsListCall description:
type QuestsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListQuestsResponse represents a list response.
type ListQuestsResponse struct {
	Quests []*QuestsItem
}

// List returns a QuestsListCall which can request data from the 42 API.
func (s *QuestsService) List() *QuestsListCall {
	return &QuestsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Quests matching the specified setting.
func (c *QuestsListCall) P(key string, values []string) *QuestsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a QuestsListCall request call. Exactly one of *ListQuestsResponse or error will be non-nil.
func (c *QuestsListCall) Do() (*ListQuestsResponse, error) {
	urls := c.s.baseURL + "/quests/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListQuestsResponse{}
	ret.Quests = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Quests); err != nil {
		return nil, err
	}
	return ret, nil
}

// QuestsGetCall description:
type QuestsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified QuestsItem.
func (s *QuestsService) Get(ID string) *QuestsGetCall {
	return &QuestsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a QuestsGetCall request call. Exactly one of *QuestsItem or error will be non-nil.
func (c *QuestsGetCall) Do() (*QuestsItem, error) {
	urls := c.s.baseURL + "/quests/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &QuestsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Quests description: Quests which can or must be done by users
func Quests(s *Service) *QuestsService {
	return &QuestsService{s: s}
}

// QuestsUsersService handles 42 API QuestsUsers.
type QuestsUsersService struct {
	s *Service
}

// QuestsUsersItem is a 42 API type
type QuestsUsersItem struct {
	Advancement interface{} `json:"advancement"`
	CreatedAt   time.Time   `json:"created_at"`
	EndAt       interface{} `json:"end_at"`
	ID          int         `json:"id"`
	Prct        interface{} `json:"prct"`
	Quest       struct {
		CampusID     interface{} `json:"campus_id"`
		CreatedAt    time.Time   `json:"created_at"`
		UpdatedAt    time.Time   `json:"updated_at"`
		GradeID      interface{} `json:"grade_id"`
		Position     int         `json:"position"`
		ID           int         `json:"id"`
		Slug         string      `json:"slug"`
		InternalName interface{} `json:"internal_name"`
		CursusID     int         `json:"cursus_id"`
		Name         string      `json:"name"`
		Kind         string      `json:"kind"`
		Description  string      `json:"description"`
	} `json:"quest"`
	QuestID   int       `json:"quest_id"`
	UpdatedAt time.Time `json:"updated_at"`
	User      struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"user"`
	ValidatedAt interface{} `json:"validated_at"`
}

// QuestsUsersListCall description:
type QuestsUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListQuestsUsersResponse represents a list response.
type ListQuestsUsersResponse struct {
	QuestsUsers []*QuestsUsersItem
}

// List returns a QuestsUsersListCall which can request data from the 42 API.
func (s *QuestsUsersService) List() *QuestsUsersListCall {
	return &QuestsUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return QuestsUsers matching the specified setting.
func (c *QuestsUsersListCall) P(key string, values []string) *QuestsUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a QuestsUsersListCall request call. Exactly one of *ListQuestsUsersResponse or error will be non-nil.
func (c *QuestsUsersListCall) Do() (*ListQuestsUsersResponse, error) {
	urls := c.s.baseURL + "/quests_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListQuestsUsersResponse{}
	ret.QuestsUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).QuestsUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// QuestsUsersGetCall description:
type QuestsUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified QuestsUsersItem.
func (s *QuestsUsersService) Get(ID string) *QuestsUsersGetCall {
	return &QuestsUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a QuestsUsersGetCall request call. Exactly one of *QuestsUsersItem or error will be non-nil.
func (c *QuestsUsersGetCall) Do() (*QuestsUsersItem, error) {
	urls := c.s.baseURL + "/quests_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &QuestsUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// QuestsUsers description: Users which earned an quest
func QuestsUsers(s *Service) *QuestsUsersService {
	return &QuestsUsersService{s: s}
}

// RolesService handles 42 API Roles.
type RolesService struct {
	s *Service
}

// RolesItem is a 42 API type
type RolesItem struct {
	Description string `json:"description"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
}

// RolesListCall description:
type RolesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListRolesResponse represents a list response.
type ListRolesResponse struct {
	Roles []*RolesItem
}

// List returns a RolesListCall which can request data from the 42 API.
func (s *RolesService) List() *RolesListCall {
	return &RolesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Roles matching the specified setting.
func (c *RolesListCall) P(key string, values []string) *RolesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a RolesListCall request call. Exactly one of *ListRolesResponse or error will be non-nil.
func (c *RolesListCall) Do() (*ListRolesResponse, error) {
	urls := c.s.baseURL + "/roles/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListRolesResponse{}
	ret.Roles = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Roles); err != nil {
		return nil, err
	}
	return ret, nil
}

// RolesGetCall description:
type RolesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified RolesItem.
func (s *RolesService) Get(ID string) *RolesGetCall {
	return &RolesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a RolesGetCall request call. Exactly one of *RolesItem or error will be non-nil.
func (c *RolesGetCall) Do() (*RolesItem, error) {
	urls := c.s.baseURL + "/roles/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &RolesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Roles description: Grants particular privileges to entities like users and applications
func Roles(s *Service) *RolesService {
	return &RolesService{s: s}
}

// RolesEntitiesService handles 42 API RolesEntities.
type RolesEntitiesService struct {
	s *Service
}

// RolesEntitiesItem is a 42 API type
type RolesEntitiesItem struct {
	CreatedAt time.Time `json:"created_at"`
	Entity    struct {
		RateLimit   int           `json:"rate_limit"`
		ID          int           `json:"id"`
		Description interface{}   `json:"description"`
		Image       interface{}   `json:"image"`
		Scopes      []interface{} `json:"scopes"`
		CreatedAt   time.Time     `json:"created_at"`
		Name        string        `json:"name"`
		Website     interface{}   `json:"website"`
		Public      bool          `json:"public"`
		UpdatedAt   time.Time     `json:"updated_at"`
		Owner       struct{}      `json:"owner"`
	} `json:"entity"`
	EntityID   int         `json:"entity_id"`
	EntityType string      `json:"entity_type"`
	ExpiresAt  interface{} `json:"expires_at"`
	ID         int         `json:"id"`
	Role       struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"role"`
}

// RolesEntitiesListCall description:
type RolesEntitiesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListRolesEntitiesResponse represents a list response.
type ListRolesEntitiesResponse struct {
	RolesEntities []*RolesEntitiesItem
}

// List returns a RolesEntitiesListCall which can request data from the 42 API.
func (s *RolesEntitiesService) List() *RolesEntitiesListCall {
	return &RolesEntitiesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return RolesEntities matching the specified setting.
func (c *RolesEntitiesListCall) P(key string, values []string) *RolesEntitiesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a RolesEntitiesListCall request call. Exactly one of *ListRolesEntitiesResponse or error will be non-nil.
func (c *RolesEntitiesListCall) Do() (*ListRolesEntitiesResponse, error) {
	urls := c.s.baseURL + "/roles_entities/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListRolesEntitiesResponse{}
	ret.RolesEntities = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).RolesEntities); err != nil {
		return nil, err
	}
	return ret, nil
}

// RolesEntitiesGetCall description:
type RolesEntitiesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified RolesEntitiesItem.
func (s *RolesEntitiesService) Get(ID string) *RolesEntitiesGetCall {
	return &RolesEntitiesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a RolesEntitiesGetCall request call. Exactly one of *RolesEntitiesItem or error will be non-nil.
func (c *RolesEntitiesGetCall) Do() (*RolesEntitiesItem, error) {
	urls := c.s.baseURL + "/roles_entities/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &RolesEntitiesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// RolesEntities description: The users and applications of a role
func RolesEntities(s *Service) *RolesEntitiesService {
	return &RolesEntitiesService{s: s}
}

// RulesService handles 42 API Rules.
type RulesService struct {
	s *Service
}

// RulesItem is a 42 API type
type RulesItem struct {
	CreatedAt    time.Time `json:"created_at"`
	Description  string    `json:"description"`
	ID           int       `json:"id"`
	InternalName string    `json:"internal_name"`
	Kind         string    `json:"kind"`
	Name         string    `json:"name"`
	Params       []struct {
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		DataType     string    `json:"data_type"`
		ID           int       `json:"id"`
		Name         string    `json:"name"`
		DefaultValue string    `json:"default_value"`
		RuleID       int       `json:"rule_id"`
	} `json:"params"`
	ProjectSessionsRules []interface{} `json:"project_sessions_rules"`
	Slug                 string        `json:"slug"`
	UpdatedAt            time.Time     `json:"updated_at"`
}

// RulesListCall description:
type RulesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListRulesResponse represents a list response.
type ListRulesResponse struct {
	Rules []*RulesItem
}

// List returns a RulesListCall which can request data from the 42 API.
func (s *RulesService) List() *RulesListCall {
	return &RulesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Rules matching the specified setting.
func (c *RulesListCall) P(key string, values []string) *RulesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a RulesListCall request call. Exactly one of *ListRulesResponse or error will be non-nil.
func (c *RulesListCall) Do() (*ListRulesResponse, error) {
	urls := c.s.baseURL + "/rules/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListRulesResponse{}
	ret.Rules = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Rules); err != nil {
		return nil, err
	}
	return ret, nil
}

// RulesGetCall description:
type RulesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified RulesItem.
func (s *RulesService) Get(ID string) *RulesGetCall {
	return &RulesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a RulesGetCall request call. Exactly one of *RulesItem or error will be non-nil.
func (c *RulesGetCall) Do() (*RulesItem, error) {
	urls := c.s.baseURL + "/rules/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &RulesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Rules description: A rule for a project
func Rules(s *Service) *RulesService {
	return &RulesService{s: s}
}

// ScaleTeamsService handles 42 API ScaleTeams.
type ScaleTeamsService struct {
	s *Service
}

// ScaleTeamsItem is a 42 API type
type ScaleTeamsItem struct {
	BeginAt        time.Time     `json:"begin_at"`
	Comment        interface{}   `json:"comment"`
	Correcteds     string        `json:"correcteds"`
	Corrector      string        `json:"corrector"`
	CreatedAt      time.Time     `json:"created_at"`
	Feedback       interface{}   `json:"feedback"`
	FeedbackRating interface{}   `json:"feedback_rating"`
	Feedbacks      []interface{} `json:"feedbacks"`
	FilledAt       interface{}   `json:"filled_at"`
	FinalMark      interface{}   `json:"final_mark"`
	Flag           struct {
		Icon      string    `json:"icon"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Positive  bool      `json:"positive"`
	} `json:"flag"`
	ID    int `json:"id"`
	Scale struct {
		Name           string    `json:"name"`
		IsPrimary      bool      `json:"is_primary"`
		IntroductionMd string    `json:"introduction_md"`
		GuidelinesMd   string    `json:"guidelines_md"`
		CreatedAt      time.Time `json:"created_at"`
		Languages      []struct {
			ID         int       `json:"id"`
			Name       string    `json:"name"`
			IDentifier string    `json:"identifier"`
			CreatedAt  time.Time `json:"created_at"`
			UpdatedAt  time.Time `json:"updated_at"`
		} `json:"languages"`
		ID                 int    `json:"id"`
		EvaluationID       int    `json:"evaluation_id"`
		Comment            string `json:"comment"`
		DisclaimerMd       string `json:"disclaimer_md"`
		CorrectionNumber   int    `json:"correction_number"`
		Duration           int    `json:"duration"`
		ManualSubscription bool   `json:"manual_subscription"`
	} `json:"scale"`
	ScaleID   int       `json:"scale_id"`
	Truant    struct{}  `json:"truant"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ScaleTeamsListCall description:
type ScaleTeamsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListScaleTeamsResponse represents a list response.
type ListScaleTeamsResponse struct {
	ScaleTeams []*ScaleTeamsItem
}

// List returns a ScaleTeamsListCall which can request data from the 42 API.
func (s *ScaleTeamsService) List() *ScaleTeamsListCall {
	return &ScaleTeamsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return ScaleTeams matching the specified setting.
func (c *ScaleTeamsListCall) P(key string, values []string) *ScaleTeamsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ScaleTeamsListCall request call. Exactly one of *ListScaleTeamsResponse or error will be non-nil.
func (c *ScaleTeamsListCall) Do() (*ListScaleTeamsResponse, error) {
	urls := c.s.baseURL + "/scale_teams/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListScaleTeamsResponse{}
	ret.ScaleTeams = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).ScaleTeams); err != nil {
		return nil, err
	}
	return ret, nil
}

// ScaleTeamsGetCall description:
type ScaleTeamsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified ScaleTeamsItem.
func (s *ScaleTeamsService) Get(ID string) *ScaleTeamsGetCall {
	return &ScaleTeamsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a ScaleTeamsGetCall request call. Exactly one of *ScaleTeamsItem or error will be non-nil.
func (c *ScaleTeamsGetCall) Do() (*ScaleTeamsItem, error) {
	urls := c.s.baseURL + "/scale_teams/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ScaleTeamsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// ScaleTeams description: A defence of a team (on a project), involving an evaluator
func ScaleTeams(s *Service) *ScaleTeamsService {
	return &ScaleTeamsService{s: s}
}

// ScalesService handles 42 API Scales.
type ScalesService struct {
	s *Service
}

// ScalesItem is a 42 API type
type ScalesItem struct {
	Comment          string    `json:"comment"`
	CorrectionNumber int       `json:"correction_number"`
	CreatedAt        time.Time `json:"created_at"`
	DisclaimerMd     string    `json:"disclaimer_md"`
	Duration         int       `json:"duration"`
	Evaluation       struct {
		ID   int    `json:"id"`
		Kind string `json:"kind"`
	} `json:"evaluation"`
	EvaluationID   int    `json:"evaluation_id"`
	GuidelinesMd   string `json:"guidelines_md"`
	ID             int    `json:"id"`
	IntroductionMd string `json:"introduction_md"`
	IsPrimary      bool   `json:"is_primary"`
	Languages      []struct {
		Name       string    `json:"name"`
		IDentifier string    `json:"identifier"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		ID         int       `json:"id"`
	} `json:"languages"`
	ManualSubscription bool   `json:"manual_subscription"`
	Name               string `json:"name"`
	Sections           []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Questions   []struct {
			CreatedAt  time.Time `json:"created_at"`
			ID         int       `json:"id"`
			Name       string    `json:"name"`
			Guidelines string    `json:"guidelines"`
			Rating     string    `json:"rating"`
			Kind       string    `json:"kind"`
		} `json:"questions"`
	} `json:"sections"`
}

// ScalesListCall description:
type ScalesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListScalesResponse represents a list response.
type ListScalesResponse struct {
	Scales []*ScalesItem
}

// List returns a ScalesListCall which can request data from the 42 API.
func (s *ScalesService) List() *ScalesListCall {
	return &ScalesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Scales matching the specified setting.
func (c *ScalesListCall) P(key string, values []string) *ScalesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ScalesListCall request call. Exactly one of *ListScalesResponse or error will be non-nil.
func (c *ScalesListCall) Do() (*ListScalesResponse, error) {
	urls := c.s.baseURL + "/scales/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListScalesResponse{}
	ret.Scales = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Scales); err != nil {
		return nil, err
	}
	return ret, nil
}

// ScalesGetCall description:
type ScalesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified ScalesItem.
func (s *ScalesService) Get(ID string) *ScalesGetCall {
	return &ScalesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a ScalesGetCall request call. Exactly one of *ScalesItem or error will be non-nil.
func (c *ScalesGetCall) Do() (*ScalesItem, error) {
	urls := c.s.baseURL + "/scales/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ScalesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Scales description: A scale is composed by questions which allows an users to rate the quality of a project .
func Scales(s *Service) *ScalesService {
	return &ScalesService{s: s}
}

// ScoresService handles 42 API Scores.
type ScoresService struct {
	s *Service
}

// ScoresItem is a 42 API type
type ScoresItem struct {
	CalculationID    int       `json:"calculation_id"`
	CoalitionID      int       `json:"coalition_id"`
	CoalitionsUserID int       `json:"coalitions_user_id"`
	CreatedAt        time.Time `json:"created_at"`
	ID               int       `json:"id"`
	Reason           string    `json:"reason"`
	ScoreableID      int       `json:"scoreable_id"`
	ScoreableType    string    `json:"scoreable_type"`
	UpdatedAt        time.Time `json:"updated_at"`
	Value            int       `json:"value"`
}

// ScoresListCall description:
type ScoresListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListScoresResponse represents a list response.
type ListScoresResponse struct {
	Scores []*ScoresItem
}

// List returns a ScoresListCall which can request data from the 42 API.
func (s *ScoresService) List() *ScoresListCall {
	return &ScoresListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Scores matching the specified setting.
func (c *ScoresListCall) P(key string, values []string) *ScoresListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a ScoresListCall request call. Exactly one of *ListScoresResponse or error will be non-nil.
func (c *ScoresListCall) Do() (*ListScoresResponse, error) {
	urls := c.s.baseURL + "/scores/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListScoresResponse{}
	ret.Scores = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Scores); err != nil {
		return nil, err
	}
	return ret, nil
}

// ScoresGetCall description:
type ScoresGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified ScoresItem.
func (s *ScoresService) Get(ID string) *ScoresGetCall {
	return &ScoresGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a ScoresGetCall request call. Exactly one of *ScoresItem or error will be non-nil.
func (c *ScoresGetCall) Do() (*ScoresItem, error) {
	urls := c.s.baseURL + "/scores/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ScoresItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// delete

// Scores description: Points given to a coalition.
func Scores(s *Service) *ScoresService {
	return &ScoresService{s: s}
}

// SkillsService handles 42 API Skills.
type SkillsService struct {
	s *Service
}

// SkillsItem is a 42 API type
type SkillsItem struct {
}

// SkillsListCall description:
type SkillsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListSkillsResponse represents a list response.
type ListSkillsResponse struct {
	Skills []*SkillsItem
}

// List returns a SkillsListCall which can request data from the 42 API.
func (s *SkillsService) List() *SkillsListCall {
	return &SkillsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Skills matching the specified setting.
func (c *SkillsListCall) P(key string, values []string) *SkillsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a SkillsListCall request call. Exactly one of *ListSkillsResponse or error will be non-nil.
func (c *SkillsListCall) Do() (*ListSkillsResponse, error) {
	urls := c.s.baseURL + "/skills/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListSkillsResponse{}
	ret.Skills = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Skills); err != nil {
		return nil, err
	}
	return ret, nil
}

// SkillsGetCall description:
type SkillsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified SkillsItem.
func (s *SkillsService) Get(ID string) *SkillsGetCall {
	return &SkillsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a SkillsGetCall request call. Exactly one of *SkillsItem or error will be non-nil.
func (c *SkillsGetCall) Do() (*SkillsItem, error) {
	urls := c.s.baseURL + "/skills/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &SkillsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Skills description: A particlar skill.
func Skills(s *Service) *SkillsService {
	return &SkillsService{s: s}
}

// SlotsService handles 42 API Slots.
type SlotsService struct {
	s *Service
}

// SlotsItem is a 42 API type
type SlotsItem struct {
}

// SlotsListCall description:
type SlotsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListSlotsResponse represents a list response.
type ListSlotsResponse struct {
	Slots []*SlotsItem
}

// List returns a SlotsListCall which can request data from the 42 API.
func (s *SlotsService) List() *SlotsListCall {
	return &SlotsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Slots matching the specified setting.
func (c *SlotsListCall) P(key string, values []string) *SlotsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a SlotsListCall request call. Exactly one of *ListSlotsResponse or error will be non-nil.
func (c *SlotsListCall) Do() (*ListSlotsResponse, error) {
	urls := c.s.baseURL + "/slots/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListSlotsResponse{}
	ret.Slots = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Slots); err != nil {
		return nil, err
	}
	return ret, nil
}

// SlotsGetCall description:
type SlotsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified SlotsItem.
func (s *SlotsService) Get(ID string) *SlotsGetCall {
	return &SlotsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a SlotsGetCall request call. Exactly one of *SlotsItem or error will be non-nil.
func (c *SlotsGetCall) Do() (*SlotsItem, error) {
	urls := c.s.baseURL + "/slots/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &SlotsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Slots description: The slots available to users for booking a project scale team.
func Slots(s *Service) *SlotsService {
	return &SlotsService{s: s}
}

// SquadsService handles 42 API Squads.
type SquadsService struct {
	s *Service
}

// SquadsItem is a 42 API type
type SquadsItem struct {
}

// create

// delete

// SquadsListCall description:
type SquadsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListSquadsResponse represents a list response.
type ListSquadsResponse struct {
	Squads []*SquadsItem
}

// List returns a SquadsListCall which can request data from the 42 API.
func (s *SquadsService) List() *SquadsListCall {
	return &SquadsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Squads matching the specified setting.
func (c *SquadsListCall) P(key string, values []string) *SquadsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a SquadsListCall request call. Exactly one of *ListSquadsResponse or error will be non-nil.
func (c *SquadsListCall) Do() (*ListSquadsResponse, error) {
	urls := c.s.baseURL + "/squads/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListSquadsResponse{}
	ret.Squads = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Squads); err != nil {
		return nil, err
	}
	return ret, nil
}

// SquadsGetCall description:
type SquadsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified SquadsItem.
func (s *SquadsService) Get(ID string) *SquadsGetCall {
	return &SquadsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a SquadsGetCall request call. Exactly one of *SquadsItem or error will be non-nil.
func (c *SquadsGetCall) Do() (*SquadsItem, error) {
	urls := c.s.baseURL + "/squads/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &SquadsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// Squads description: A squads is the managing container of squads_users.
func Squads(s *Service) *SquadsService {
	return &SquadsService{s: s}
}

// SquadsUsersService handles 42 API SquadsUsers.
type SquadsUsersService struct {
	s *Service
}

// SquadsUsersItem is a 42 API type
type SquadsUsersItem struct {
}

// create

// delete

// SquadsUsersListCall description:
type SquadsUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListSquadsUsersResponse represents a list response.
type ListSquadsUsersResponse struct {
	SquadsUsers []*SquadsUsersItem
}

// List returns a SquadsUsersListCall which can request data from the 42 API.
func (s *SquadsUsersService) List() *SquadsUsersListCall {
	return &SquadsUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return SquadsUsers matching the specified setting.
func (c *SquadsUsersListCall) P(key string, values []string) *SquadsUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a SquadsUsersListCall request call. Exactly one of *ListSquadsUsersResponse or error will be non-nil.
func (c *SquadsUsersListCall) Do() (*ListSquadsUsersResponse, error) {
	urls := c.s.baseURL + "/squads_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListSquadsUsersResponse{}
	ret.SquadsUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).SquadsUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// SquadsUsers description: A squads_users will group users inside a same coalition
func SquadsUsers(s *Service) *SquadsUsersService {
	return &SquadsUsersService{s: s}
}

// SubnotionsService handles 42 API Subnotions.
type SubnotionsService struct {
	s *Service
}

// SubnotionsItem is a 42 API type
type SubnotionsItem struct {
	Attachments []interface{} `json:"attachments"`
	CreatedAt   time.Time     `json:"created_at"`
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Notepad     interface{}   `json:"notepad"`
	Notion      interface{}   `json:"notion"`
	Slug        string        `json:"slug"`
}

// SubnotionsListCall description:
type SubnotionsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListSubnotionsResponse represents a list response.
type ListSubnotionsResponse struct {
	Subnotions []*SubnotionsItem
}

// List returns a SubnotionsListCall which can request data from the 42 API.
func (s *SubnotionsService) List() *SubnotionsListCall {
	return &SubnotionsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Subnotions matching the specified setting.
func (c *SubnotionsListCall) P(key string, values []string) *SubnotionsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a SubnotionsListCall request call. Exactly one of *ListSubnotionsResponse or error will be non-nil.
func (c *SubnotionsListCall) Do() (*ListSubnotionsResponse, error) {
	urls := c.s.baseURL + "/subnotions/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListSubnotionsResponse{}
	ret.Subnotions = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Subnotions); err != nil {
		return nil, err
	}
	return ret, nil
}

// SubnotionsGetCall description:
type SubnotionsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified SubnotionsItem.
func (s *SubnotionsService) Get(ID string) *SubnotionsGetCall {
	return &SubnotionsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a SubnotionsGetCall request call. Exactly one of *SubnotionsItem or error will be non-nil.
func (c *SubnotionsGetCall) Do() (*SubnotionsItem, error) {
	urls := c.s.baseURL + "/subnotions/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &SubnotionsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Subnotions description: The elearning subnotion in a notion
func Subnotions(s *Service) *SubnotionsService {
	return &SubnotionsService{s: s}
}

// TagsService handles 42 API Tags.
type TagsService struct {
	s *Service
}

// TagsItem is a 42 API type
type TagsItem struct {
	ID         int           `json:"id"`
	Kind       string        `json:"kind"`
	Name       string        `json:"name"`
	Subnotions []interface{} `json:"subnotions"`
	Users      []interface{} `json:"users"`
}

// TagsListCall description:
type TagsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListTagsResponse represents a list response.
type ListTagsResponse struct {
	Tags []*TagsItem
}

// List returns a TagsListCall which can request data from the 42 API.
func (s *TagsService) List() *TagsListCall {
	return &TagsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Tags matching the specified setting.
func (c *TagsListCall) P(key string, values []string) *TagsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a TagsListCall request call. Exactly one of *ListTagsResponse or error will be non-nil.
func (c *TagsListCall) Do() (*ListTagsResponse, error) {
	urls := c.s.baseURL + "/tags/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListTagsResponse{}
	ret.Tags = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Tags); err != nil {
		return nil, err
	}
	return ret, nil
}

// TagsGetCall description:
type TagsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified TagsItem.
func (s *TagsService) Get(ID string) *TagsGetCall {
	return &TagsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a TagsGetCall request call. Exactly one of *TagsItem or error will be non-nil.
func (c *TagsGetCall) Do() (*TagsItem, error) {
	urls := c.s.baseURL + "/tags/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &TagsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Tags description: Non-hierarchical keyword, acting as a meta-data and helping to describe entities.
func Tags(s *Service) *TagsService {
	return &TagsService{s: s}
}

// TeamsService handles 42 API Teams.
type TeamsService struct {
	s *Service
}

// TeamsItem is a 42 API type
type TeamsItem struct {
}

// TeamsListCall description:
type TeamsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListTeamsResponse represents a list response.
type ListTeamsResponse struct {
	Teams []*TeamsItem
}

// List returns a TeamsListCall which can request data from the 42 API.
func (s *TeamsService) List() *TeamsListCall {
	return &TeamsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Teams matching the specified setting.
func (c *TeamsListCall) P(key string, values []string) *TeamsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a TeamsListCall request call. Exactly one of *ListTeamsResponse or error will be non-nil.
func (c *TeamsListCall) Do() (*ListTeamsResponse, error) {
	urls := c.s.baseURL + "/teams/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListTeamsResponse{}
	ret.Teams = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Teams); err != nil {
		return nil, err
	}
	return ret, nil
}

// TeamsGetCall description:
type TeamsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified TeamsItem.
func (s *TeamsService) Get(ID string) *TeamsGetCall {
	return &TeamsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a TeamsGetCall request call. Exactly one of *TeamsItem or error will be non-nil.
func (c *TeamsGetCall) Do() (*TeamsItem, error) {
	urls := c.s.baseURL + "/teams/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &TeamsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Teams description: One or many users which have to finish a project together.
func Teams(s *Service) *TeamsService {
	return &TeamsService{s: s}
}

// TeamsUploadsService handles 42 API TeamsUploads.
type TeamsUploadsService struct {
	s *Service
}

// TeamsUploadsItem is a 42 API type
type TeamsUploadsItem struct {
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	FinalMark int       `json:"final_mark"`
	ID        int       `json:"id"`
	Upload    struct {
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		ID           int       `json:"id"`
		EvaluationID int       `json:"evaluation_id"`
		Name         string    `json:"name"`
		Description  string    `json:"description"`
	} `json:"upload"`
	UploadID int `json:"upload_id"`
}

// TeamsUploadsListCall description:
type TeamsUploadsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListTeamsUploadsResponse represents a list response.
type ListTeamsUploadsResponse struct {
	TeamsUploads []*TeamsUploadsItem
}

// List returns a TeamsUploadsListCall which can request data from the 42 API.
func (s *TeamsUploadsService) List() *TeamsUploadsListCall {
	return &TeamsUploadsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return TeamsUploads matching the specified setting.
func (c *TeamsUploadsListCall) P(key string, values []string) *TeamsUploadsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a TeamsUploadsListCall request call. Exactly one of *ListTeamsUploadsResponse or error will be non-nil.
func (c *TeamsUploadsListCall) Do() (*ListTeamsUploadsResponse, error) {
	urls := c.s.baseURL + "/teams_uploads/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListTeamsUploadsResponse{}
	ret.TeamsUploads = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).TeamsUploads); err != nil {
		return nil, err
	}
	return ret, nil
}

// TeamsUploadsGetCall description:
type TeamsUploadsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified TeamsUploadsItem.
func (s *TeamsUploadsService) Get(ID string) *TeamsUploadsGetCall {
	return &TeamsUploadsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a TeamsUploadsGetCall request call. Exactly one of *TeamsUploadsItem or error will be non-nil.
func (c *TeamsUploadsGetCall) Do() (*TeamsUploadsItem, error) {
	urls := c.s.baseURL + "/teams_uploads/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &TeamsUploadsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// TeamsUploads description: An uploaded mark for a team, given by a bot (like the Moulinette), without any defence.
func TeamsUploads(s *Service) *TeamsUploadsService {
	return &TeamsUploadsService{s: s}
}

// TeamsUsersService handles 42 API TeamsUsers.
type TeamsUsersService struct {
	s *Service
}

// TeamsUsersItem is a 42 API type
type TeamsUsersItem struct {
	CreatedAt  time.Time `json:"created_at"`
	ID         int       `json:"id"`
	Leader     bool      `json:"leader"`
	Occurrence int       `json:"occurrence"`
	Team       struct {
		UpdatedAt     time.Time   `json:"updated_at"`
		Closed        bool        `json:"closed?"`
		FinalMark     interface{} `json:"final_mark"`
		ProjectID     int         `json:"project_id"`
		Status        string      `json:"status"`
		TerminatingAt interface{} `json:"terminating_at"`
		Users         []struct {
			Validated      bool   `json:"validated"`
			ProjectsUserID int    `json:"projects_user_id"`
			ID             int    `json:"id"`
			Login          string `json:"login"`
			URL            string `json:"url"`
			Leader         bool   `json:"leader"`
			Occurrence     int    `json:"occurrence"`
		} `json:"users"`
		Locked           bool        `json:"locked?"`
		Name             string      `json:"name"`
		URL              string      `json:"url"`
		RepoUUID         string      `json:"repo_uuid"`
		LockedAt         interface{} `json:"locked_at"`
		Validated        interface{} `json:"validated?"`
		RepoURL          interface{} `json:"repo_url"`
		CreatedAt        time.Time   `json:"created_at"`
		ProjectSessionID int         `json:"project_session_id"`
		ID               int         `json:"id"`
		ClosedAt         interface{} `json:"closed_at"`
	} `json:"team"`
	TeamID int `json:"team_id"`
	User   struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"user"`
	UserID    int  `json:"user_id"`
	Validated bool `json:"validated"`
}

// TeamsUsersListCall description:
type TeamsUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListTeamsUsersResponse represents a list response.
type ListTeamsUsersResponse struct {
	TeamsUsers []*TeamsUsersItem
}

// List returns a TeamsUsersListCall which can request data from the 42 API.
func (s *TeamsUsersService) List() *TeamsUsersListCall {
	return &TeamsUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return TeamsUsers matching the specified setting.
func (c *TeamsUsersListCall) P(key string, values []string) *TeamsUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a TeamsUsersListCall request call. Exactly one of *ListTeamsUsersResponse or error will be non-nil.
func (c *TeamsUsersListCall) Do() (*ListTeamsUsersResponse, error) {
	urls := c.s.baseURL + "/teams_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListTeamsUsersResponse{}
	ret.TeamsUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).TeamsUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// TeamsUsersGetCall description:
type TeamsUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified TeamsUsersItem.
func (s *TeamsUsersService) Get(ID string) *TeamsUsersGetCall {
	return &TeamsUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a TeamsUsersGetCall request call. Exactly one of *TeamsUsersItem or error will be non-nil.
func (c *TeamsUsersGetCall) Do() (*TeamsUsersItem, error) {
	urls := c.s.baseURL + "/teams_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &TeamsUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// TeamsUsers description: Team composed of one User
func TeamsUsers(s *Service) *TeamsUsersService {
	return &TeamsUsersService{s: s}
}

// TitlesService handles 42 API Titles.
type TitlesService struct {
	s *Service
}

// TitlesItem is a 42 API type
type TitlesItem struct {
}

// TitlesListCall description:
type TitlesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListTitlesResponse represents a list response.
type ListTitlesResponse struct {
	Titles []*TitlesItem
}

// List returns a TitlesListCall which can request data from the 42 API.
func (s *TitlesService) List() *TitlesListCall {
	return &TitlesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Titles matching the specified setting.
func (c *TitlesListCall) P(key string, values []string) *TitlesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a TitlesListCall request call. Exactly one of *ListTitlesResponse or error will be non-nil.
func (c *TitlesListCall) Do() (*ListTitlesResponse, error) {
	urls := c.s.baseURL + "/titles/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListTitlesResponse{}
	ret.Titles = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Titles); err != nil {
		return nil, err
	}
	return ret, nil
}

// TitlesGetCall description:
type TitlesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified TitlesItem.
func (s *TitlesService) Get(ID string) *TitlesGetCall {
	return &TitlesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a TitlesGetCall request call. Exactly one of *TitlesItem or error will be non-nil.
func (c *TitlesGetCall) Do() (*TitlesItem, error) {
	urls := c.s.baseURL + "/titles/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &TitlesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Titles description: Titles a user can obtain, generally through achievements. It will be displayed on their profile and on the forum.
func Titles(s *Service) *TitlesService {
	return &TitlesService{s: s}
}

// TitlesUsersService handles 42 API TitlesUsers.
type TitlesUsersService struct {
	s *Service
}

// TitlesUsersItem is a 42 API type
type TitlesUsersItem struct {
}

// TitlesUsersListCall description:
type TitlesUsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListTitlesUsersResponse represents a list response.
type ListTitlesUsersResponse struct {
	TitlesUsers []*TitlesUsersItem
}

// List returns a TitlesUsersListCall which can request data from the 42 API.
func (s *TitlesUsersService) List() *TitlesUsersListCall {
	return &TitlesUsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return TitlesUsers matching the specified setting.
func (c *TitlesUsersListCall) P(key string, values []string) *TitlesUsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a TitlesUsersListCall request call. Exactly one of *ListTitlesUsersResponse or error will be non-nil.
func (c *TitlesUsersListCall) Do() (*ListTitlesUsersResponse, error) {
	urls := c.s.baseURL + "/titles_users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListTitlesUsersResponse{}
	ret.TitlesUsers = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).TitlesUsers); err != nil {
		return nil, err
	}
	return ret, nil
}

// TitlesUsersGetCall description:
type TitlesUsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified TitlesUsersItem.
func (s *TitlesUsersService) Get(ID string) *TitlesUsersGetCall {
	return &TitlesUsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a TitlesUsersGetCall request call. Exactly one of *TitlesUsersItem or error will be non-nil.
func (c *TitlesUsersGetCall) Do() (*TitlesUsersItem, error) {
	urls := c.s.baseURL + "/titles_users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &TitlesUsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// TitlesUsers description: Users who have a title.
func TitlesUsers(s *Service) *TitlesUsersService {
	return &TitlesUsersService{s: s}
}

// TransactionsService handles 42 API Transactions.
type TransactionsService struct {
	s *Service
}

// TransactionsItem is a 42 API type
type TransactionsItem struct {
	ID               int    `json:"id"`
	Reason           string `json:"reason"`
	TransactableID   int    `json:"transactable_id"`
	TransactableType string `json:"transactable_type"`
	User             struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"user"`
	UserID int `json:"user_id"`
	Value  int `json:"value"`
}

// TransactionsListCall description:
type TransactionsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListTransactionsResponse represents a list response.
type ListTransactionsResponse struct {
	Transactions []*TransactionsItem
}

// List returns a TransactionsListCall which can request data from the 42 API.
func (s *TransactionsService) List() *TransactionsListCall {
	return &TransactionsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Transactions matching the specified setting.
func (c *TransactionsListCall) P(key string, values []string) *TransactionsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a TransactionsListCall request call. Exactly one of *ListTransactionsResponse or error will be non-nil.
func (c *TransactionsListCall) Do() (*ListTransactionsResponse, error) {
	urls := c.s.baseURL + "/transactions/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListTransactionsResponse{}
	ret.Transactions = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Transactions); err != nil {
		return nil, err
	}
	return ret, nil
}

// TransactionsGetCall description:
type TransactionsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified TransactionsItem.
func (s *TransactionsService) Get(ID string) *TransactionsGetCall {
	return &TransactionsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a TransactionsGetCall request call. Exactly one of *TransactionsItem or error will be non-nil.
func (c *TransactionsGetCall) Do() (*TransactionsItem, error) {
	urls := c.s.baseURL + "/transactions/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &TransactionsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Transactions description: Transaction represents Altarian Dollars earned.
func Transactions(s *Service) *TransactionsService {
	return &TransactionsService{s: s}
}

// TranslationsService handles 42 API Translations.
type TranslationsService struct {
	s *Service
}

// TranslationsItem is a 42 API type
type TranslationsItem struct {
	CreatedAt time.Time `json:"created_at"`
	Default   bool      `json:"default"`
	Fields    struct {
		Name        string      `json:"name"`
		Description interface{} `json:"description"`
	} `json:"fields"`
	ID                    int    `json:"id"`
	LanguageID            int    `json:"language_id"`
	TranslatableID        int    `json:"translatable_id"`
	TranslatableType      string `json:"translatable_type"`
	TranslationsStructure struct {
		ID                 int       `json:"id"`
		SearchableBy       []string  `json:"searchable_by"`
		UpdatedAt          time.Time `json:"updated_at"`
		IDentifiedBy       []string  `json:"identified_by"`
		StructuresKind     string    `json:"structures_kind"`
		UpToDate           bool      `json:"up_to_date"`
		FieldsOrganisation struct {
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"fields_organisation"`
		TypeName  string    `json:"type_name"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"translations_structure"`
	TranslationsStructureID int         `json:"translations_structure_id"`
	UpToDate                bool        `json:"up_to_date"`
	UpdatedAt               time.Time   `json:"updated_at"`
	UserID                  interface{} `json:"user_id"`
}

// TranslationsListCall description:
type TranslationsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListTranslationsResponse represents a list response.
type ListTranslationsResponse struct {
	Translations []*TranslationsItem
}

// List returns a TranslationsListCall which can request data from the 42 API.
func (s *TranslationsService) List() *TranslationsListCall {
	return &TranslationsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Translations matching the specified setting.
func (c *TranslationsListCall) P(key string, values []string) *TranslationsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a TranslationsListCall request call. Exactly one of *ListTranslationsResponse or error will be non-nil.
func (c *TranslationsListCall) Do() (*ListTranslationsResponse, error) {
	urls := c.s.baseURL + "/translations/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListTranslationsResponse{}
	ret.Translations = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Translations); err != nil {
		return nil, err
	}
	return ret, nil
}

// TranslationsGetCall description:
type TranslationsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified TranslationsItem.
func (s *TranslationsService) Get(ID string) *TranslationsGetCall {
	return &TranslationsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a TranslationsGetCall request call. Exactly one of *TranslationsItem or error will be non-nil.
func (c *TranslationsGetCall) Do() (*TranslationsItem, error) {
	urls := c.s.baseURL + "/translations/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &TranslationsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// Translations description: Translations
func Translations(s *Service) *TranslationsService {
	return &TranslationsService{s: s}
}

// UserCandidaturesService handles 42 API UserCandidatures.
type UserCandidaturesService struct {
	s *Service
}

// UserCandidaturesItem is a 42 API type
type UserCandidaturesItem struct {
	BirthCity          string      `json:"birth_city"`
	BirthCountry       string      `json:"birth_country"`
	BirthDate          string      `json:"birth_date"`
	ContactAffiliation interface{} `json:"contact_affiliation"`
	ContactFirstName   interface{} `json:"contact_first_name"`
	ContactLastName    interface{} `json:"contact_last_name"`
	ContactPhone1      interface{} `json:"contact_phone1"`
	ContactPhone2      interface{} `json:"contact_phone2"`
	Country            string      `json:"country"`
	CreatedAt          time.Time   `json:"created_at"`
	Email              string      `json:"email"`
	Gender             string      `json:"gender"`
	HiddenPhone        interface{} `json:"hidden_phone"`
	ID                 int         `json:"id"`
	Language           interface{} `json:"language"`
	MaxLevelLogic      interface{} `json:"max_level_logic"`
	MaxLevelMemory     interface{} `json:"max_level_memory"`
	MeetingDate        interface{} `json:"meeting_date"`
	OtherInformation   interface{} `json:"other_information"`
	Phone              interface{} `json:"phone"`
	PhoneCountryCode   interface{} `json:"phone_country_code"`
	Pin                interface{} `json:"pin"`
	PiscineDate        interface{} `json:"piscine_date"`
	PostalCity         string      `json:"postal_city"`
	PostalComplement   interface{} `json:"postal_complement"`
	PostalCountry      interface{} `json:"postal_country"`
	PostalStreet       string      `json:"postal_street"`
	PostalZipCode      interface{} `json:"postal_zip_code"`
	UpdatedAt          time.Time   `json:"updated_at"`
	UserID             int         `json:"user_id"`
	ZipCode            string      `json:"zip_code"`
}

// UserCandidaturesListCall description:
type UserCandidaturesListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListUserCandidaturesResponse represents a list response.
type ListUserCandidaturesResponse struct {
	UserCandidatures []*UserCandidaturesItem
}

// List returns a UserCandidaturesListCall which can request data from the 42 API.
func (s *UserCandidaturesService) List() *UserCandidaturesListCall {
	return &UserCandidaturesListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return UserCandidatures matching the specified setting.
func (c *UserCandidaturesListCall) P(key string, values []string) *UserCandidaturesListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a UserCandidaturesListCall request call. Exactly one of *ListUserCandidaturesResponse or error will be non-nil.
func (c *UserCandidaturesListCall) Do() (*ListUserCandidaturesResponse, error) {
	urls := c.s.baseURL + "/user_candidatures/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListUserCandidaturesResponse{}
	ret.UserCandidatures = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).UserCandidatures); err != nil {
		return nil, err
	}
	return ret, nil
}

// UserCandidaturesGetCall description:
type UserCandidaturesGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified UserCandidaturesItem.
func (s *UserCandidaturesService) Get(ID string) *UserCandidaturesGetCall {
	return &UserCandidaturesGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a UserCandidaturesGetCall request call. Exactly one of *UserCandidaturesItem or error will be non-nil.
func (c *UserCandidaturesGetCall) Do() (*UserCandidaturesItem, error) {
	urls := c.s.baseURL + "/user_candidatures/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &UserCandidaturesItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// delete

// UserCandidatures description: The candidature of an user
func UserCandidatures(s *Service) *UserCandidaturesService {
	return &UserCandidaturesService{s: s}
}

// UsersService handles 42 API Users.
type UsersService struct {
	s *Service
}

// UsersItem is a 42 API type
type UsersItem struct {
	Achievements  []interface{} `json:"achievements"`
	AnonymizeDate time.Time     `json:"anonymize_date"`
	Campus        []struct {
		VogsphereID int    `json:"vogsphere_id"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
		TimeZone    string `json:"time_zone"`
		Language    struct {
			UpdatedAt  time.Time `json:"updated_at"`
			ID         int       `json:"id"`
			Name       string    `json:"name"`
			IDentifier string    `json:"identifier"`
			CreatedAt  time.Time `json:"created_at"`
		} `json:"language"`
		UsersCount int `json:"users_count"`
	} `json:"campus"`
	CampusUsers []struct {
		ID        int  `json:"id"`
		UserID    int  `json:"user_id"`
		CampusID  int  `json:"campus_id"`
		IsPrimary bool `json:"is_primary"`
	} `json:"campus_users"`
	CorrectionPoint int `json:"correction_point"`
	CursusUsers     []struct {
		User struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			URL   string `json:"url"`
		} `json:"user"`
		ID           int         `json:"id"`
		Grade        interface{} `json:"grade"`
		CursusID     int         `json:"cursus_id"`
		HasCoalition bool        `json:"has_coalition"`
		Cursus       struct {
			ID        int       `json:"id"`
			CreatedAt time.Time `json:"created_at"`
			Name      string    `json:"name"`
			Slug      string    `json:"slug"`
		} `json:"cursus"`
		BeginAt time.Time     `json:"begin_at"`
		EndAt   interface{}   `json:"end_at"`
		Level   float32       `json:"level"`
		Skills  []interface{} `json:"skills"`
	} `json:"cursus_users"`
	Displayname     string `json:"displayname"`
	Email           string `json:"email"`
	ExpertisesUsers []struct {
		Interested  bool      `json:"interested"`
		Value       int       `json:"value"`
		ContactMe   bool      `json:"contact_me"`
		CreatedAt   time.Time `json:"created_at"`
		UserID      int       `json:"user_id"`
		ID          int       `json:"id"`
		ExpertiseID int       `json:"expertise_id"`
	} `json:"expertises_users"`
	FirstName      string        `json:"first_name"`
	Groups         []interface{} `json:"groups"`
	ID             int           `json:"id"`
	ImageURL       string        `json:"image_url"`
	LanguagesUsers []struct {
		ID         int       `json:"id"`
		LanguageID int       `json:"language_id"`
		UserID     int       `json:"user_id"`
		Position   int       `json:"position"`
		CreatedAt  time.Time `json:"created_at"`
	} `json:"languages_users"`
	LastName     string        `json:"last_name"`
	Location     interface{}   `json:"location"`
	Login        string        `json:"login"`
	Partnerships []interface{} `json:"partnerships"`
	Patroned     []struct {
		UserID      int       `json:"user_id"`
		GodfatherID int       `json:"godfather_id"`
		Ongoing     bool      `json:"ongoing"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		ID          int       `json:"id"`
	} `json:"patroned"`
	Patroning      []interface{} `json:"patroning"`
	Phone          interface{}   `json:"phone"`
	PoolMonth      string        `json:"pool_month"`
	PoolYear       string        `json:"pool_year"`
	ProjectsUsers  []interface{} `json:"projects_users"`
	Staff          bool          `json:"staff?"`
	Titles         []interface{} `json:"titles"`
	TitlesUsers    []interface{} `json:"titles_users"`
	URL            string        `json:"url"`
	UsualFirstName string        `json:"usual_first_name"`
	UsualFullName  string        `json:"usual_full_name"`
	Wallet         int           `json:"wallet"`
}

// UsersListCall description:
type UsersListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListUsersResponse represents a list response.
type ListUsersResponse struct {
	Users []*UsersItem
}

// List returns a UsersListCall which can request data from the 42 API.
func (s *UsersService) List() *UsersListCall {
	return &UsersListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Users matching the specified setting.
func (c *UsersListCall) P(key string, values []string) *UsersListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a UsersListCall request call. Exactly one of *ListUsersResponse or error will be non-nil.
func (c *UsersListCall) Do() (*ListUsersResponse, error) {
	urls := c.s.baseURL + "/users/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListUsersResponse{}
	ret.Users = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Users); err != nil {
		return nil, err
	}
	return ret, nil
}

// UsersGetCall description:
type UsersGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified UsersItem.
func (s *UsersService) Get(ID string) *UsersGetCall {
	return &UsersGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a UsersGetCall request call. Exactly one of *UsersItem or error will be non-nil.
func (c *UsersGetCall) Do() (*UsersItem, error) {
	urls := c.s.baseURL + "/users/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &UsersItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// create

// update

// Users description: A 42 student, staff, or any entity with a 42 account.
func Users(s *Service) *UsersService {
	return &UsersService{s: s}
}

// WaitlistsService handles 42 API Waitlists.
type WaitlistsService struct {
	s *Service
}

// WaitlistsItem is a 42 API type
type WaitlistsItem struct {
	CreatedAt        time.Time `json:"created_at"`
	ID               int       `json:"id"`
	WaitlistableID   int       `json:"waitlistable_id"`
	WaitlistableType string    `json:"waitlistable_type"`
}

// WaitlistsListCall description:
type WaitlistsListCall struct {
	s         *Service
	urlParams map[string][]string
	header    http.Header
}

// ListWaitlistsResponse represents a list response.
type ListWaitlistsResponse struct {
	Waitlists []*WaitlistsItem
}

// List returns a WaitlistsListCall which can request data from the 42 API.
func (s *WaitlistsService) List() *WaitlistsListCall {
	return &WaitlistsListCall{
		s:         s.s,
		urlParams: make(map[string][]string),
	}
}

// P sets an optional parameter and its values: Only return Waitlists matching the specified setting.
func (c *WaitlistsListCall) P(key string, values []string) *WaitlistsListCall {
	c.urlParams[key] = values
	return c
}

// Do executes a WaitlistsListCall request call. Exactly one of *ListWaitlistsResponse or error will be non-nil.
func (c *WaitlistsListCall) Do() (*ListWaitlistsResponse, error) {
	urls := c.s.baseURL + "/waitlists/" + "?" + url.Values(c.urlParams).Encode()
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &ListWaitlistsResponse{}
	ret.Waitlists = nil
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(&(*target).Waitlists); err != nil {
		return nil, err
	}
	return ret, nil
}

// WaitlistsGetCall description:
type WaitlistsGetCall struct {
	s         *Service
	id        string
	urlParams map[string][]string
	header    http.Header
}

// Get gets the specified WaitlistsItem.
func (s *WaitlistsService) Get(ID string) *WaitlistsGetCall {
	return &WaitlistsGetCall{
		s:  s.s,
		id: ID,
	}
}

// Do executes a WaitlistsGetCall request call. Exactly one of *WaitlistsItem or error will be non-nil.
func (c *WaitlistsGetCall) Do() (*WaitlistsItem, error) {
	urls := c.s.baseURL + "/waitlists/" + c.id
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.s.Do(req)
	if err != nil {
		return nil, err
	}
	c.header = res.Header
	ret := &WaitlistsItem{}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
}

// delete

// Waitlists description: Waitlist for an event or an exam.
func Waitlists(s *Service) *WaitlistsService {
	return &WaitlistsService{s: s}
}
