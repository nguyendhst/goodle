package goodle

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type (
	MoodleAPI struct {
		base  string
		token string
		fetch Fetcher
	}

	Functions struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}

	SiteInfo struct {
		Sitename  string      `json:"sitename"`
		Username  string      `json:"username"`
		Firstname string      `json:"firstname"`
		Lastname  string      `json:"lastname"`
		Fullname  string      `json:"fullname"`
		Functions []Functions `json:"functions"`
	}

	UnreadConvo struct {
		Favourites int `json:"favourites"`
		Types      struct {
			One   int `json:"1"`
			Two   int `json:"2"`
			Three int `json:"3"`
		}
	}
)

// NewClient creates a new MoodleAPI instance.
func NewClient(base, token string) *MoodleAPI {
	if base != "" {
		if !strings.HasSuffix(base, "/") {
			base = base + "/"
		}
	}

	return &MoodleAPI{
		base:  base,
		token: token,
		fetch: &DefaultFetcher{},
	}
}

// GetUnreadConversationsCount returns the number of unread conversations.
func (m *MoodleAPI) GetUnreadConversationsCount() (UnreadConvo, error) {
	url := fmt.Sprintf("%swebservice/rest/server.php?wstoken=%s&wsfunction=%s&moodlewsrestformat=json&moodlewssettingraw=true", m.base, m.token, "core_message_get_unread_conversation_counts")
	body, _, _, err := m.fetch.Fetch(url)
	if err != nil {
		return UnreadConvo{}, err
	}
	bodyStr := strings.TrimSpace(string(body))

	if strings.HasPrefix(bodyStr, "{\"exception\":\"") {
		return UnreadConvo{}, errors.New(bodyStr)
	}

	var data UnreadConvo

	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return UnreadConvo{}, err
	}
	return data, nil
}

// GetSiteInfo returns the site info.
func (m *MoodleAPI) GetSiteInfo() (SiteInfo, error) {
	url := fmt.Sprintf("%swebservice/rest/server.php?wstoken=%s&wsfunction=%s&moodlewsrestformat=json&moodlewssettingraw=true", m.base, m.token, "core_webservice_get_site_info")

	body, _, _, err := m.fetch.Fetch(url)
	if err != nil {
		return SiteInfo{}, err
	}

	bodyStr := strings.TrimSpace(string(body))

	if strings.HasPrefix(bodyStr, "{\"exception\":\"") {
		return SiteInfo{}, errors.New(bodyStr)
	}

	var data SiteInfo

	if err := json.Unmarshal([]byte(body), &data); err != nil {
		return SiteInfo{}, errors.New("Server returned unexpected response. " + err.Error())
	}

	return data, nil

}
