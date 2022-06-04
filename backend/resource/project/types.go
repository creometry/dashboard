package project

import "fmt"

type ReqData struct {
	UsrProjectName string `json:"projectName"`
	Plan 		 string `json:"plan"`
}

func (r *ReqData) Validate() error {
	if r.UsrProjectName == "" {
		return fmt.Errorf("projectName is required")
	}
	if r.Plan == "" {
		return fmt.Errorf("plan is required")
	}
	return nil
}

type RespData struct {
	ProjectId string `json:"id"`
}
