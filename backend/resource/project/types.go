package project

import "fmt"

type ReqData struct {
	UsrProjectName string `json:"projectName"`
}

func (r *ReqData) Validate() error {
	if r.UsrProjectName == "" {
		return fmt.Errorf("projectName is required")
	}
	return nil
}

type RespData struct {
	ProjectId string `json:"id"`
}
