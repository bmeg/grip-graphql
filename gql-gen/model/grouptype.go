
package model

import (
	"encoding/json"
	"fmt"
)

type SafeGroupType struct {
	Contained TypedObject `json:"contained,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ID *string `json:"id,omitempty"`
	Membership *string `json:"membership,omitempty"`
	Quantity *int `json:"quantity,omitempty"`
	ManagingEntity TypedObject `json:"managingEntity"`
	Language *string `json:"language,omitempty"`
	Characteristic []*GroupCharacteristic `json:"characteristic,omitempty"`
	Active *bool `json:"active,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Member []*GroupMember `json:"member,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Type *string `json:"type,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *GroupType) UnmarshalJSON(b []byte) error {
	var safe SafeGroupType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = GroupType{
		Name: safe.Name,
		Description: safe.Description,
		ID: safe.ID,
		Membership: safe.Membership,
		Quantity: safe.Quantity,
		Language: safe.Language,
		Characteristic: safe.Characteristic,
		Active: safe.Active,
		ImplicitRules: safe.ImplicitRules,
		Extension: safe.Extension,
		Member: safe.Member,
		ResourceType: safe.ResourceType,
		ModifierExtension: safe.ModifierExtension,
		Text: safe.Text,
		Type: safe.Type,
		Identifier: safe.Identifier,
		Meta: safe.Meta,
		Code: safe.Code,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "managingEntity", safe.ManagingEntity.Typename, &o.ManagingEntity); err != nil {
		return fmt.Errorf("failed to unmarshal ManagingEntity: %w", err)
	}

	return nil
}
