package project

import (
	"fmt"
)

type ReqData struct {
	UsrProjectName string `json:"projectName"`
	BillingAccountId string `json:"billingAccountId"`
	PaymentToken string `json:"paymentToken"`
	UserId         string `json:"userId"`
	Plan           string `json:"plan"`
	GitRepoName    string `json:"gitRepoName"`
	GitRepoBranch  string `json:"gitRepoBranch"`
	GitRepoUrl     string `json:"gitRepoUrl"`
}

type ReqDataNewUser struct {
	UsrProjectName string `json:"projectName"`
	Username       string `json:"username"`
	Plan           string `json:"plan"`
	GitRepoName    string `json:"gitRepoName"`
	GitRepoBranch  string `json:"gitRepoBranch"`
	GitRepoUrl     string `json:"gitRepoUrl"`
}

type ReqDataLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReqDataRegister struct {
	Username string `json:"username"`
}

func (r *ReqDataLogin) Validate() error {
	if r.Username == "" {
		return fmt.Errorf("username is required")
	}
	if r.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}

func (r *ReqDataNewUser) Validate() error {
	if r.UsrProjectName == "" {
		return fmt.Errorf("project name is required")
	}
	if r.Plan == "" {
		return fmt.Errorf("plan is required")
	}
	if r.Username == "" {
		return fmt.Errorf("username is required")
	}
	return nil
}

func (r *ReqData) Validate() error {
	if r.BillingAccountId == "" {
		return fmt.Errorf("billing account id is required")
	}
	if r.PaymentToken == "" {
		return fmt.Errorf("payment token is required")
	}
	if r.UsrProjectName == "" {
		return fmt.Errorf("projectName is required")
	}
	if r.Plan == "" {
		return fmt.Errorf("plan is required")
	}
	if r.UserId == "" {
		return fmt.Errorf("user id is required")
	}
	return nil
}

type TokenResponse struct {
	Token string `json:"token"`
}

type CreateNsRespData struct {
	Error  string `json:"error"`
	NsName string `json:"ns_name"`
}

type RespData struct {
	ProjectId string `json:"id"`
}

type RespDataProvisionProject struct {
	ProjectId string `json:"projectId"`
}

type RespDataProvisionProjectNewUser struct {
	ProjectId string `json:"projectId"`
	Token     string `json:"token"`
	Password  string `json:"password"`
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
	Id           string `json:"userId"`
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


type RespDataCreateBillingAccount struct {
	Id string `json:"id"`
}

type CheckPaymeePaymentResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Code   int64 `json:"code"`
	Data   struct {
		PaymentStatus string `json:"payment_status"`
		Token string `json:"token"`
		Amount float64 `json:"amount"`
		TransactionId int64 `json:"transaction_id"`
		BuyerId int64 `json:"buyer_id"`
	}
}