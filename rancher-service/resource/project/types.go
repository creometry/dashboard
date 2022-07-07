package project

import (
	"fmt"
)

type ReqData struct {
	UsrProjectName string `json:"projectName"`
	UserId         string `json:"userId"`
	Plan           string `json:"plan"`
	GitRepoName    string `json:"gitRepoName"`
	GitRepoBranch  string `json:"gitRepoBranch"`
	GitRepoUrl     string `json:"gitRepoUrl"`
}

func (r *ReqData) Validate() error {
	if r.UsrProjectName == "" {
		return fmt.Errorf("projectName is required")
	}
	if r.Plan == "" {
		return fmt.Errorf("plan is required")
	}
	if r.UserId == "" {
		return fmt.Errorf("user id is required")
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

type CreateNsRespData struct {
	Error  string `json:"error"`
	NsName string `json:"ns_name"`
}

type RespData struct {
	ProjectId string `json:"id"`
}

type RespDataCreateProjectAndRepo struct {
	User_token string `json:"user_token"`
	Namespace  string `json:"namespace"`
	ProjectId  string `json:"projectId"`
}

type RespDataRoleBinding struct {
	RoleTemplateId string `json:"roleTemplateId"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Code           string `json:"code"`
}

type RespDataCreateUser struct {
	Id           string   `json:"id"`
	PrincipalIds []string `json:"principalIds"`
}

type Kubeconfig struct {
	BaseType string `json:"baseType"`
	Config   string `json:"config"`
	Type     string `json:"type"`
}

type RespDataLogin struct {
	AuthProvider string `json:"authProvider"`
	Token        string `json:"token"`
	Name         string `json:"name"`
	Id           string `json:"id"`
}

type RespDataCreateGitRepo struct {
	Id string `json:"id"`
}

type ReqDataKubeconfig struct {
	Token string `json:"token"`
}

type RespDataUser struct {
	Namespace    string   `json:"namespace"`
	Id           string   `json:"id"`
	Token        string   `json:"token"`
	PrincipalIds []string `json:"principalIds"`
	ProjectId    string   `json:"projectId"`
}

type UserData struct {
	Data []RespDataUser `json:"data"`
}

type RespDataNs struct {
	Data []NsData `json:"data"`
}

type NsData struct {
	Id       string        `json:"id"`
	Metadata MetadDataData `json:"metadata"`
}
type MetadDataData struct {
	Annotations map[string]string `json:"annotations"`
}

type RespDataProjectsByUser struct {
	Data []ProjectRoliBindingsData `json:"data"`
}

type ProjectRoliBindingsData struct {
	Id string `json:"id"`
}

type FindUserData struct {
	Data []Data `json:"data"`
}

type Data struct {
	Id           string   `json:"id"`
	PrincipalIds []string `json:"principalIds"`
}

type RespDataUserByUserId struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Id       string `json:"id"`
	Type     string `json:"type"`
}

type RespDataTeamMembers struct {
	Data []TeamMemberData `json:"data"`
}

type TeamMemberData struct {
	UserId string `json:"userId"`
}