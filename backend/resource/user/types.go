package user

import "fmt"

type ReqData struct {
	RoleTemplateId string `json:"roleTemplateId"`
	UserId         string `json:"userId"`
	ProjectId      string `json:"projectId"`
}

func (r *ReqData) Validate() error {
	if r.RoleTemplateId == "" {
		return fmt.Errorf("roleTemplateId is required")
	}
	if r.UserId == "" {
		return fmt.Errorf("userId is required")
	}
	if r.ProjectId == "" {
		return fmt.Errorf("projectId is required")
	}

	return nil
}

type RespData struct {
	RoleTemplateId string `json:"roleTemplateId"`
	Name           string `json:"name"`
	Type           string `json:"type"`
}
