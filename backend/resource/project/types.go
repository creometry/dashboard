package project

import "fmt"

type ReqData struct {
	UsrProjectName string `json:"projectName"`
	Namespace	  string `json:"namespace"`
	Username   string `json:"username"`
	Plan 		 string `json:"plan"`
	GitRepoName string `json:"gitRepoName"`
	GitRepoBranch string `json:"gitRepoBranch"`
	GitRepoUrl string `json:"gitRepoUrl"`
	Id_token 	 string `json:"id_token"`
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
	if r.Id_token == "" {
		return fmt.Errorf("id_token is required")
	}
	if r.Namespace == "" {
		return fmt.Errorf("namespace is required")
	}
	if r.GitRepoName == "" {
		return fmt.Errorf("gitRepoName is required")
	}
	if r.GitRepoUrl == "" {
		return fmt.Errorf("gitRepoUrl is required")
	}
	if r.GitRepoBranch == "" {
		r.GitRepoBranch = "master"
	}
	return nil
}

type RespData struct {
	ProjectId string `json:"id"`
}

type RespDataRoleBinding struct {
	RoleTemplateId string `json:"roleTemplateId"`
	Name           string `json:"name"`
	Type           string `json:"type"`
}

type RespDataCreateUser	struct {
	Id string `json:"id"`
	PrincipalIds []string `json:"principalIds"`
}

type Kubeconfig struct {
	BaseType string `json:"baseType"`
	Config 	string `json:"config"`
	Type 	string `json:"type"`
}

type RespDataLogin struct {
	AuthProvider string `json:"authProvider"`
	Token 		 string `json:"token"`
	Name 		 string `json:"name"`
	Id 			 string `json:"id"`
}

type RespDataCreateGitRepo struct {
	Id 			 string `json:"id"`
}