package team

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Creometry/dashboard/go-provisioner/utils"
)

// Exportable function

func ListTeamMembers(projectId string) ([]RespDataUserByUserId, error) {

	rancherToken, rancherURL, err := getRancherTokenAndUrl()

	if err != nil {
		return []RespDataUserByUserId{}, err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s", rancherURL, "/v3/projectroletemplatebindings?projectId=", projectId), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rancherToken))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	dt := RespDataTeamMembers{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		return nil, err
	}
	var res []RespDataUserByUserId
	if len(dt.Data) > 0 {
		// loop through all the members, get their userId and get their names
		for _, user := range dt.Data {
			d, err := getUserById(strings.Split(user.UserId, "/")[0])
			if err != nil {
				continue
			} else {
				if d.Type != "error" {
					res = append(res, d)
				}
			}
		}
	}
	return res, nil

}

// Local functions

func getUserById(userId string) (RespDataUserByUserId, error) {

	rancherToken, rancherURL, err := getRancherTokenAndUrl()

	if err != nil {
		return RespDataUserByUserId{}, err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s", rancherURL, "/v3/users/", userId), nil)
	if err != nil {
		return RespDataUserByUserId{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", rancherToken))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return RespDataUserByUserId{}, err
	}

	defer resp.Body.Close()

	// parse response body
	dt := RespDataUserByUserId{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RespDataUserByUserId{}, err
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		return RespDataUserByUserId{}, err
	}

	return dt, nil

}

// Local functions

func getRancherTokenAndUrl() (string, string, error) {
	rancherURL, err := utils.GetVariable("config", "RANCHER_URL")
	if err != nil {
		return "", "", err
	}

	rancherToken, err := utils.GetVariable("secrets", "RANCHER_TOKEN")
	if err != nil {
		return "", "", err
	}

	return rancherToken, rancherURL, nil
}
