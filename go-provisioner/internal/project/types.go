package project

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ReqData struct {
	// TODO: add billing account data and validte it
	UsrProjectName   string `json:"projectName"`
	BillingAccountId string `json:"billingAccountId"`
	PaymentToken     string `json:"paymentToken"`
	UserId           string `json:"userId"`
	UUID             string `json:"uuid"`
	Plan             string `json:"plan"`
	GitRepoName      string `json:"gitRepoName"`
	GitRepoBranch    string `json:"gitRepoBranch"`
	GitRepoUrl       string `json:"gitRepoUrl"`
	IsCompany        bool   `json:"isCompany"`
	CompanyName      string `json:"companyName"`
	TaxId            string `json:"taxId"`
	Phone            string `json:"phone"`
	Email            string `json:"email"`
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
	if r.UUID == "" {
		return fmt.Errorf("user uuid is required")
	}
	if r.BillingAccountId == "1" && r.Email == "" {
		return fmt.Errorf("email is required")
	}
	if r.IsCompany && (r.CompanyName == "" || r.TaxId == "") {
		return fmt.Errorf("company name and tax id are required")
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
	Created   string `json:"created"`
	CreatedTS int64  `json:"createdTS"`
	UUID      string `json:"uuid"`
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
	UUID         string   `json:"uuid"`
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
	UUID         string `json:"uuid"`
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

type CheckPaymeePaymentResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Code    int64  `json:"code"`
	Data    struct {
		PaymentStatus bool    `json:"payment_status"`
		Token         string  `json:"token"`
		Amount        float64 `json:"amount"`
		TransactionId int64   `json:"transaction_id"`
		BuyerId       int64   `json:"buyer_id"`
	}
}

type RespDataCreateBillingAccount struct {
	Id string `json:"uuid"`
}

type ReqDataCreateBillingAccount struct {
	BillingAdmins []Admin   `json:"billingAdmins"`
	Company       Company   `json:"company"`
	Projects      []Project `json:"projects"`
}

type Company struct {
	IsCompany bool   `json:"isCompany"`
	TaxId     string `json:"TaxId"`
	Name      string `json:"name"`
}

type Admin struct {
	UUID         string `json:"uuid"`
	Email        string `json:"email"`
	Phone_number string `json:"phone_number"`
}

type Project struct {
	ProjectId         string    `json:"projectId"`
	ClusterId         string    `json:"clusterId"`
	CreationTimeStamp time.Time `json:"creationTimeStamp"`
	State             string    `json:"State"`
	Plan              string    `json:"accountType"`
}

type ReqDataAddProjectToBillingAccount struct {
	BillingAccountUUID uuid.UUID `json:"billing_account_uuid"`
	ProjectId          string    `json:"project_id"`
	ClusterId          string    `json:"clusterId"`
	CreationTimeStamp  time.Time `json:"creationTimeStamp"`
	Plan               string    `json:"accountType"`
	State              string    `json:"state"`
}
