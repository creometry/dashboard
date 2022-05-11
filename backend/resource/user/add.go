package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)


func AddUserToProject(data ReqData) (RespData, error) {

	req, err := http.NewRequest("POST", os.Getenv("ADD_USER_TO_PROJECT_URL"), bytes.NewBuffer([]byte(fmt.Sprintf(`{"userId":"%s","projectId":"%s","roleTemplateId":"%s"}`, data.UserId, data.ProjectId,data.RoleTemplateId))))
	if err != nil {
		return RespData{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("RANCHER_TOKEN")))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return RespData{}, err
	}

	defer resp.Body.Close()
	// parse response body
	dt := RespData{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RespData{}, err
	}
	err = json.Unmarshal(body, &dt)
	if err != nil {
		return  RespData{}, err
	}

	return dt, nil

}