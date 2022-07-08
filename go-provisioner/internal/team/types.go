package team

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

type RespDataRoleBinding struct {
	RoleTemplateId string `json:"roleTemplateId"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Code           string `json:"code"`
}
