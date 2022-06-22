Select
    eav_attribute.attribute_id,
    eav_attribute.backend_gateway,
    eav_attribute.backend_field,
    eav_attribute.backend_type,
    eav_attribute.backend_table,
    eav_attribute.frontend_input,
    eav_attribute.response_field,
    eav_attribute.is_required,
    ent.entity_type_id,
    attrs.attribute_set_id
From
    eav_entity_attribute
    inner join eav_attribute on eav_entity_attribute.attribute_id = eav_attribute.attribute_id
    left join eav_entity_type ent on eav_entity_attribute.entity_type_id = ent.entity_type_id
    left join eav_attribute_set attrs on eav_entity_attribute.attribute_set_id = attrs.attribute_set_id
where
    ent.entity_type_code = 'catalog_product'
    and attrs.attribute_set_name = 'Simple Product';