
package model

import (
	"encoding/json"
	"fmt"
)

type SafeSubstanceType struct {
	Expiry *string `json:"expiry,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Code *CodeableReference `json:"code,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Status *string `json:"status,omitempty"`
	Quantity *Quantity `json:"quantity,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ID *string `json:"id,omitempty"`
	Ingredient []*SubstanceIngredient `json:"ingredient,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Language *string `json:"language,omitempty"`
	Instance *string `json:"instance,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *SubstanceType) UnmarshalJSON(b []byte) error {
	var safe SafeSubstanceType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = SubstanceType{
		Expiry: safe.Expiry,
		Description: safe.Description,
		ResourceType: safe.ResourceType,
		Extension: safe.Extension,
		Identifier: safe.Identifier,
		Code: safe.Code,
		ImplicitRules: safe.ImplicitRules,
		Status: safe.Status,
		Quantity: safe.Quantity,
		ModifierExtension: safe.ModifierExtension,
		ID: safe.ID,
		Ingredient: safe.Ingredient,
		Text: safe.Text,
		Language: safe.Language,
		Instance: safe.Instance,
		Meta: safe.Meta,
		Category: safe.Category,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
