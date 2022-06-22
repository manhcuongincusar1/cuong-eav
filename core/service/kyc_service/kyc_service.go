package service

import (
	e "cuong-eav/constants/entity"
	"cuong-eav/core/port"
	"cuong-eav/domain"
	"cuong-eav/domain/user"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

func NewKycService(kycRepo port.KycRepo) port.KycService {
	return &kycService{
		kycRepo: kycRepo,
	}
}

type MultipleLanguagesString struct {
	StoreCode string `json:"store_code" valid:"Required"`
	Value     string `json:"value"`
}

type kycService struct {
	kycRepo port.KycRepo
}

func (p *kycService) CreateUser(user *user.User, attributeSet, entityType string) error {
	var (
		eavAttributeList []*domain.AttributeListResponse
		err              error
	)

	if eavAttributeList, err = p.kycRepo.GetAttributeList(attributeSet, entityType); err != nil {
		return err
	}

	jsons, _ := json.MarshalIndent(eavAttributeList, "", "\t")
	fmt.Printf("Att list: %s\n", jsons)

	user.Credential.Uuid = uuid.NewString()
	if err = p.kycRepo.CreateUserEntity(&user.Credential); err != nil {
		return err
	}

	if err = p.structWriter(user.Credential, eavAttributeList, user.Credential.Id); err != nil {
		return err
	}

	if err = p.mapWriter(user.AdditionalInfo, eavAttributeList, user.Credential.Id); err != nil {
		return err
	}
	return nil
}

func (p *kycService) structWriter(entity interface{}, eavAttributeList []*domain.AttributeListResponse, entityId int) error {
	// Loop to check each attreibute:
	for _, attribute := range eavAttributeList {
		if attribute.BackendField == "" {
			continue
		}
		if !attribute.IsRequired {
			continue
		}
		var kind reflect.Kind
		var value interface{}
		//check backend field
		if regexp.MustCompile(`^\{\}`).MatchString(attribute.BackendField) {
			// if struct
			kind = reflect.Indirect(reflect.ValueOf(entity)).FieldByName(attribute.BackendField[2:strings.Index(attribute.BackendGateway, ".")]).FieldByName(attribute.BackendField[strings.Index(attribute.BackendGateway, ".")+1 : len(attribute.BackendField)]).Kind()
			value = reflect.Indirect(reflect.ValueOf(entity)).FieldByName(attribute.BackendField[2:strings.Index(attribute.BackendGateway, ".")]).FieldByName(attribute.BackendField[strings.Index(attribute.BackendGateway, ".")+1 : len(attribute.BackendField)]).Interface()
		} else {
			kind = reflect.Indirect(reflect.ValueOf(entity)).FieldByName(attribute.BackendField).Kind()
			value = reflect.Indirect(reflect.ValueOf(entity)).FieldByName(attribute.BackendField).Interface()
		}

		if kind == reflect.Slice {
			switch attribute.BackendType {
			case reflect.String.String():
				f := reflect.Indirect(reflect.ValueOf(entity)).FieldByName(attribute.BackendField).Interface().([]MultipleLanguagesString)
				for _, i := range f {
					if err := p.kycRepo.CreateUserDataText(&domain.UserEntityText{
						Value: i.Value,
					}); err != nil {
						return err
					}
				}
			}
		} else {
			p.insert(kind, value, entityId, attribute)
		}

	}
	return nil
}

func (p *kycService) mapWriter(data map[string]interface{}, eavAttributeList []*domain.AttributeListResponse, entityId int) error {
	if data == nil {
		return nil
	}
	for _, attribute := range eavAttributeList {
		if attribute.FrontendInput == "" {
			fmt.Printf("Frontend Input %s: ", attribute.FrontendInput)
			continue
		}
		if !attribute.IsRequired {
			continue
		}

		var kind reflect.Kind
		var value interface{}
		var ok bool

		kind = strToKind(attribute.BackendType)
		value, ok = data[attribute.FrontendInput]

		fmt.Printf("Value: \"%s\" With key \"%s\" and Kind is \"%v\" \n", value, attribute.FrontendInput, kind)
		if !ok {
			continue
		}
		if err := p.insert(kind, value, entityId, attribute); err != nil {
			return err
		}
	}
	return nil
}

func (p *kycService) insert(kind reflect.Kind, value interface{}, entityId int, attribute *domain.AttributeListResponse) error {
	fmt.Printf("Inserter input: %+v\n", value)
	switch kind {
	case reflect.Int:
		if reflect.TypeOf(value).String() == reflect.Float64.String() {
			value = int(value.(float64))
		}
		if err := p.kycRepo.CreateUserDataInt(&domain.UserEntityInt{
			AttributeID: attribute.AttributeID,
			EntityID:    entityId,
			StoreID:     e.NoUseStoreSetting,
			Value:       value.(int),
		}); err != nil {
			return err
		}
	case reflect.Bool:
		var valueInt int
		if value.(bool) {
			valueInt = 1
		}
		if err := p.kycRepo.CreateUserDataInt(&domain.UserEntityInt{
			AttributeID: attribute.AttributeID,
			EntityID:    entityId,
			StoreID:     e.NoUseStoreSetting,
			Value:       valueInt,
		}); err != nil {
			return err
		}

	case reflect.Float64:
		if err := p.kycRepo.CreateUserDataDecimal(&domain.UserEntityDecimal{
			AttributeID: attribute.AttributeID,
			EntityID:    entityId,
			StoreID:     e.NoUseStoreSetting,
			Value:       value.(float64),
		}); err != nil {
			return err
		}

	case reflect.String:
		if err := p.kycRepo.CreateUserDataText(&domain.UserEntityText{
			AttributeID: attribute.AttributeID,
			EntityID:    entityId,
			StoreID:     e.NoUseStoreSetting,
			Value:       value.(string),
		}); err != nil {
			return err
		}

	}

	return nil
}

func strToKind(s string) reflect.Kind {
	var strKindMap = map[string]reflect.Kind{}
	for k := reflect.Invalid; k <= reflect.UnsafePointer; k++ {
		strKindMap[k.String()] = k
	}
	k, ok := strKindMap[s]
	if !ok {
		return reflect.Invalid
	}
	return k
}

// func (p *kycService) read(kind reflect.Kind)
