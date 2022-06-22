package service

import "cuong-eav/domain"

func (p *kycService) CreateAttribute(att *domain.EavAttribute) error {
	// Create Attribute
	if err := p.kycRepo.CreateAttribute(att); err != nil {
		return err
	}

	// Create eav_entity_attribute
	// TODO: remove in production
	eA := domain.EavEntityAttribute{
		AttributeGroupID: 1,
		AttributeID:      int(att.AttributeID),
		AttributeSetID:   1,
		EntityTypeID:     1,
	}
	if err := eA.Create(); err != nil {
		return err
	}

	return nil
}
func (p *kycService) GetAttributes(attributeSet, entityType string) ([]*domain.AttributeListResponse, error) {
	return p.kycRepo.GetAttributeList(attributeSet, entityType)
}
