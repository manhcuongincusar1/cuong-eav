package service

import (
	e "cuong-eav/constants/entity"
	"cuong-eav/domain"
	"cuong-eav/domain/user"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func (p *kycService) GetUsers() ([]*user.User, error) {
	result := []*user.User{}
	userEntities, err := p.kycRepo.GetUserEntity(domain.UserEntity{})
	if err != nil {
		return nil, err
	}
	for _, userEntity := range userEntities {
		att, err := p.getUserAttribute(userEntity.Id)
		if err != nil {
			return nil, err
		}

		result = append(result, &user.User{
			Credential:     *userEntity,
			AdditionalInfo: att,
		})
	}
	return result, nil
}

func (p *kycService) getUserAttribute(userId int) (map[string]interface{}, error) {
	var (
		eavAttributeList []*domain.AttributeListResponse
		err              error
	)
	eavAttributeList, err = p.kycRepo.GetAttributeList(e.AttributeSetDefault, e.EntityTypeUser)
	if err != nil {
		return nil, err
	}

	response := make(map[string]interface{}, len(eavAttributeList))

	for _, attribute := range eavAttributeList {
		// Get Value
		var responseValue interface{}
		switch attribute.BackendTable {
		case e.UserEntityValueInt:
			fmt.Println("Reading from : ", e.UserEntityValueInt)
			var value []*domain.UserEntityInt
			if value, err = p.kycRepo.GetUserDataInt(attribute.AttributeID, userId); err != nil {
				return nil, err
			}

			jsonS, _ := json.MarshalIndent(value, "", "\t")
			fmt.Printf("Read from database: %s\n", jsonS)

			if len(value) <= 0 {
				continue
			}
			if len(value) == 1 {
				switch attribute.BackendType {
				case reflect.Bool.String():
					responseValue = value[0].Value == 1
				case reflect.Int.String():
					responseValue = value[0].Value
				}
			}
			if len(value) > 1 {
				var slice []interface{}
				for _, valueStore := range value {
					switch attribute.BackendType {
					case reflect.Bool.String():
						slice = append(slice, valueStore.Value == 1)
					case reflect.Int.String():
						slice = append(slice, valueStore.Value)
					}
				}
				responseValue = slice
			}

		case e.UserEntityValueText:
			fmt.Println("Reading from : ", e.UserEntityValueText)
			var value []*domain.UserEntityText
			if value, err = p.kycRepo.GetUserDataText(attribute.AttributeID, userId); err != nil {
				return nil, err
			}

			jsonS, _ := json.MarshalIndent(value, "", "\t")
			fmt.Printf("Read from database: %s\n", jsonS)
			if len(value) <= 0 {
				continue
			}
			if len(value) == 1 {
				responseValue = value[0].Value
			}
			if len(value) > 1 {
				var slice []interface{}
				for _, valueStore := range value {
					slice = append(slice, valueStore.Value)
				}
				responseValue = slice
			}

		case e.UserEntityValueDecimal:
			var value []*domain.UserEntityDecimal
			if value, err = p.kycRepo.GetUserDataDecimal(attribute.AttributeID, userId); err != nil {
				return nil, err
			}
			if len(value) <= 0 {
				continue
			}
			if len(value) == 1 {
				responseValue = value[0].Value
			}
			if len(value) > 1 {
				var slice []interface{}
				for _, valueStore := range value {
					{
						slice = append(slice, struct {
							StoreCode string
							Value     float64
						}{StoreCode: strings.ToLower("null store code"), Value: valueStore.Value})
					}
				}
				responseValue = slice
			}
		}
		var responseField string
		fmt.Println("Resp Field: ", attribute.ResponseField)
		responseField = attribute.ResponseField
		response[responseField] = responseValue
	}
	fmt.Printf("Resp: %+v\n", response)
	return response, nil
}
