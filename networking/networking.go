package networking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	NumRepos int    `json:"public_repos"`
}

func UserInfo(login string) (*User, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%v", login))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
