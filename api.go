package clickmeeting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const clickMeetingURL = "https://api.clickmeeting.com/v1/"

type api struct {
	apiKey string

	client *http.Client
}

func NewAPI(apiKey string) Client {
	return &api{apiKey: apiKey, client: &http.Client{}}
}

type encoder interface {
	Encode() string
}

func (api *api) send(req *http.Request, holder interface{}) error {
	req.Header.Add("X-Api-Key", api.apiKey)

	resp, err := api.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	//d, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Status)
	//fmt.Println(string(d))

	dec := json.NewDecoder(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		var apiErr APIError
		decErr := dec.Decode(&apiErr)
		if decErr != nil {
			return decErr
		}
		return apiErr
	}
	return dec.Decode(holder)
}
func (api *api) sendGet(path string, data encoder, holder interface{}) error {
	req, err := http.NewRequest(http.MethodGet, api.getURL(path)+"?"+data.Encode(), nil)
	if err != nil {
		return fmt.Errorf("failed to create new request: %w", err)
	}
	return api.send(req, holder)
}
func (api *api) sendPost(path string, data encoder, holder interface{}) error {
	req, err := http.NewRequest(http.MethodPost, api.getURL(path), bytes.NewBufferString(data.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create new request: %w", err)
	}
	return api.send(req, holder)
}
func (api *api) sendPut(path string, data encoder, holder interface{}) error {
	req, err := http.NewRequest(http.MethodPut, api.getURL(path), bytes.NewBufferString(data.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create new request: %w", err)
	}
	return api.send(req, holder)
}
func (api *api) sendDelete(path string, data encoder, holder interface{}) error {
	req, err := http.NewRequest(http.MethodDelete, api.getURL(path), bytes.NewBufferString(data.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create new request: %w", err)
	}
	return api.send(req, holder)
}

func (api *api) getURL(path string) string {
	return fmt.Sprintf("%s%s.json", clickMeetingURL, path)
}

func (api *api) ListRooms(status RoomStatus) ([]Room, error) {
	var rooms []Room
	err := api.sendGet("conferences/"+string(status), url.Values{}, &rooms)
	return rooms, err
}

func (api *api) CreateRoom(newRoom NewRoom, opts ...CreateRoomOption) (Room, error) {
	v := url.Values{}
	v.Add("name", newRoom.Name)
	v.Add("room_type", newRoom.RoomType.String())
	v.Add("permanent_room", "0")
	if newRoom.PermanentRoom {
		v.Set("permanent_room", "1")
	}

	v.Add("access_type", strconv.Itoa(int(newRoom.AccessType)))
	for _, opt := range opts {
		opt(v)
	}
	var resp struct {
		Room Room `json:"room"`
	}
	err := api.sendPost("conferences", v, &resp)

	return resp.Room, err
}

func (api *api) UpdateRoom(roomID int, opts ...UpdateRoomOption) (Room, error) {
	v := url.Values{}
	for _, opt := range opts {
		opt(v)
	}
	var resp struct {
		Room Room `json:"conference"`
	}
	err := api.sendPut(fmt.Sprintf("conferences/%d", roomID), v, &resp)

	return resp.Room, err
}

func (api *api) DeleteRoom(roomID int) error {
	var resp struct {
		Result string `json:"result"`
	}
	err := api.sendDelete(fmt.Sprintf("conferences/%d", roomID), url.Values{}, &resp)
	return err
}

func (api *api) GetSessions(roomID int) ([]SessionSummary, error) {
	panic("implement me")
}

func (api *api) GetSession(roomID int, sessionID int) (Session, error) {
	panic("implement me")
}

func (api *api) GenerateAccessTokens(roomID int, howMany int) ([]AccessToken, error) {
	panic("implement me")
}

func (api *api) GetAccessTokens(roomID int) ([]AccessToken, error) {
	panic("implement me")
}

func (api *api) AutoLoginHash(roomID int) (string, error) {
	panic("implement me")
}

func (api *api) SendInvitation(roomID int, language string, attendees []string, opts ...SendInvitationOption) error {
	v := url.Values{
		"attendees[][email]": attendees,
	}
	for _, opt := range opts {
		opt(v)
	}

	var empty interface{}
	return api.sendPost(fmt.Sprintf("conferences/%d/invitation/email/%s", roomID, language), v, &empty)
}

func (api *api) GetRegistrations(roomID int, status string) ([]Participant, error) {
	var participants []Participant
	err := api.sendGet(fmt.Sprintf("conferences/%d/registrations/%s", roomID, status), url.Values{}, &participants)
	return participants, err
}

func (api *api) RegisterParticipant(roomID int, participant NewParticipant, opts ...RegisterParticipantOption) (string, error) {
	v := url.Values{}
	v.Add("registration[1]", participant.FirstName)
	v.Add("registration[2]", participant.LastName)
	v.Add("registration[3]", participant.EmailAddress)

	for _, opt := range opts {
		opt(v)
	}

	var resp struct {
		Status string `json:"status"`
		URL    string `json:"url"`
	}
	err := api.sendPost(fmt.Sprintf("conferences/%d/registration", roomID), v, &resp)
	return resp.URL, err
}

func (api *api) GetParticipants(roomID int, sessionID int) ([]Participant, error) {
	panic("implement me")
}
