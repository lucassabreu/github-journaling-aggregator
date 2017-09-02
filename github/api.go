package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type API struct {
	username string
	token    string
	client   http.Client
}

func NewAPI(username string, token string) API {
	return API{
		username: username,
		token:    token,
	}
}

func (a *API) newRequest(url string) (req *http.Request, err error) {
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.SetBasicAuth(a.username, a.token)
	return
}

func (a *API) getResponse(url string) (res *http.Response, err error) {
	req, err := a.newRequest(url)
	if err != nil {
		return
	}
	res, err = a.client.Do(req)
	return
}

func (a *API) GetUser() (u CurrentUser, err error) {
	r, err := a.getResponse(fmt.Sprintf("https://api.github.com/%s", "user"))
	if err != nil {
		return
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		return
	}
	u.api = a

	fmt.Println(u)

	repos, err := u.Repos()
	for _, r := range repos {
		fmt.Println(r)
	}

	return
}

type CurrentUser struct {
	api      *API
	Login    string
	URL      string
	ReposUrl string `json:"repos_url"`
}

func (u *CurrentUser) Repos() (repos []*Repository, err error) {
	r, err := u.api.getResponse(u.ReposUrl)
	if err != nil {
		return
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&repos)
	return
}

type Repository struct {
	FullName string `json:"full_name"`
	Name     string `json:"name"`
}
