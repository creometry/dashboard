package controllers

import "errors"


type CreateNsRequestBody struct {
	ProjectName string `json:"projectName"`
	ProjectId string `json:"projectId"`
}


// function to validate the request body
func (req *CreateNsRequestBody) Validate() error {
	if req.ProjectName == "" || req.ProjectId == "" {
		return errors.New("projectName and projectId are required")
	}
	return nil
}