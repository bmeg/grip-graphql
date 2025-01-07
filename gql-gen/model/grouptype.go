
package model

import (
	"encoding/json"
	"fmt"
)

type SafeGroupType struct {
	Text *Narrative `json:"text,omitempty"`
	Membership *string `json:"membership,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Characteristic []*GroupCharacteristic `json:"characteristic,omitempty"`
	ManagingEntity TypedObject `json:"managingEntity"`
	ResourceType *string `json:"resourceType,omitempty"`
	Member []*GroupMember `json:"member,omitempty"`
	Quantity *string `json:"quantity,omitempty"`
	Description *string `json:"description,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ID *string `json:"id,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Language *string `json:"language,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Type *string `json:"type,omitempty"`
	Active *string `json:"active,omitempty"`
	Name *string `json:"name,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *GroupType) UnmarshalJSON(b []byte) error {
	var safe SafeGroupType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = GroupType{
		Text: safe.Text,
		Membership: safe.Membership,
		Meta: safe.Meta,
		Characteristic: safe.Characteristic,
		ResourceType: safe.ResourceType,
		Member: safe.Member,
		Quantity: safe.Quantity,
		Description: safe.Description,
		Identifier: safe.Identifier,
		ID: safe.ID,
		ImplicitRules: safe.ImplicitRules,
		Extension: safe.Extension,
		Language: safe.Language,
		ModifierExtension: safe.ModifierExtension,
		Code: safe.Code,
		Type: safe.Type,
		Active: safe.Active,
		Name: safe.Name,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "managingEntity", safe.ManagingEntity.Typename, &o.ManagingEntity); err != nil {
		return fmt.Errorf("failed to unmarshal ManagingEntity: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
