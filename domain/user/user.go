package user

import "cuong-eav/domain"

type User struct {
	Credential     domain.UserEntity      `json:"credential"`
	AdditionalInfo map[string]interface{} `json:"additional_info"`
}
