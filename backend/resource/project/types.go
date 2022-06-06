package project

import "fmt"

type ReqData struct {
	UsrProjectName string `json:"projectName"`
	Plan 		 string `json:"plan"`
	Username 	 string `json:"username"`
	Email 		 string `json:"email"`
}

func (r *ReqData) Validate() error {
	if r.UsrProjectName == "" {
		return fmt.Errorf("projectName is required")
	}
	if r.Plan == "" {
		return fmt.Errorf("plan is required")
	}
	if r.Username == "" {
		return fmt.Errorf("username is required")
	}
	if r.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}

type RespData struct {
	ProjectId string `json:"id"`
}

type Kubeconfig struct {
	BaseType string `json:"baseType"`
	Config 	string `json:"config"`
	Type 	string `json:"type"`
}