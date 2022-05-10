package project

import "fmt"


type ReqData struct {
	Plan string `json:"plan"`
	UsrProjectName string `json:"projectName"` 
}

func (r *ReqData) Validate() error {
	if r.Plan == "" || r.UsrProjectName == "" {
		return fmt.Errorf("plan and projectName are required")
	}
	return nil
}


type RespData struct {
	ProjectId string `json:"id"`
}
