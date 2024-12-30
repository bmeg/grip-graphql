
package model

import (
	"encoding/json"
	"fmt"
)

type SafeGroupType struct {
	Language *string `json:"language,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Member TypedObject `json:"member"`
	ID *string `json:"id,omitempty"`
	Quantity *string `json:"quantity,omitempty"`
	Name *string `json:"name,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Active *string `json:"active,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Membership *string `json:"membership,omitempty"`
	Type *string `json:"type,omitempty"`
	Characteristic []*GroupCharacteristic `json:"characteristic,omitempty"`
	ManagingEntity TypedObject `json:"managingEntity"`
	Meta *Meta `json:"meta,omitempty"`
}

func (o *GroupType) UnmarshalJSON(b []byte) error {
	var safe SafeGroupType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = GroupType{
		Language: safe.Language,
		Extension: safe.Extension,
		ImplicitRules: safe.ImplicitRules,
		Identifier: safe.Identifier,
		ID: safe.ID,
		Quantity: safe.Quantity,
		Name: safe.Name,
		Text: safe.Text,
		Code: safe.Code,
		Description: safe.Description,
		ModifierExtension: safe.ModifierExtension,
		Active: safe.Active,
		ResourceType: safe.ResourceType,
		Membership: safe.Membership,
		Type: safe.Type,
		Characteristic: safe.Characteristic,
		Meta: safe.Meta,
	}
	if err := unmarshalUnion(b, "member", safe.Member.Typename, &o.Member); err != nil {
		return fmt.Errorf("failed to unmarshal Member: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "managingEntity", safe.ManagingEntity.Typename, &o.ManagingEntity); err != nil {
		return fmt.Errorf("failed to unmarshal ManagingEntity: %w", err)
	}

	return nil
}
