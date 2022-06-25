package project

import "fmt"

type ReqData struct {
	UsrProjectName string `json:"projectName"`
	Username   string `json:"username"`
	Plan 		 string `json:"plan"`
	GitRepoName string `json:"gitRepoName"`
	GitRepoBranch string `json:"gitRepoBranch"`
	GitRepoUrl string `json:"gitRepoUrl"`
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

type RespDataCreateProjectAndRepo struct {
	User_token string `json:"user_token"`
	Namespace string `json:"namespace"`
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

type ReqDataKubeconfig struct {
	Token string `json:"token"`
}