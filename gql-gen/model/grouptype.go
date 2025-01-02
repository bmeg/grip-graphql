
package model

import (
	"encoding/json"
	"fmt"
)

type SafeGroupType struct {
	Quantity *string `json:"quantity,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Type *string `json:"type,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Membership *string `json:"membership,omitempty"`
	Active *string `json:"active,omitempty"`
	ID *string `json:"id,omitempty"`
	Language *string `json:"language,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	ManagingEntity TypedObject `json:"managingEntity"`
	Extension []*Extension `json:"extension,omitempty"`
	Name *string `json:"name,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Member TypedObject `json:"member"`
	Characteristic []*GroupCharacteristic `json:"characteristic,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Description *string `json:"description,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *GroupType) UnmarshalJSON(b []byte) error {
	var safe SafeGroupType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = GroupType{
		Quantity: safe.Quantity,
		Meta: safe.Meta,
		Type: safe.Type,
		Membership: safe.Membership,
		Active: safe.Active,
		ID: safe.ID,
		Language: safe.Language,
		ModifierExtension: safe.ModifierExtension,
		ImplicitRules: safe.ImplicitRules,
		Extension: safe.Extension,
		Name: safe.Name,
		Text: safe.Text,
		Code: safe.Code,
		Characteristic: safe.Characteristic,
		Identifier: safe.Identifier,
		ResourceType: safe.ResourceType,
		Description: safe.Description,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "managingEntity", safe.ManagingEntity.Typename, &o.ManagingEntity); err != nil {
		return fmt.Errorf("failed to unmarshal ManagingEntity: %w", err)
	}
	if err := unmarshalUnion(b, "member", safe.Member.Typename, &o.Member); err != nil {
		return fmt.Errorf("failed to unmarshal Member: %w", err)
	}

	return nil
}
