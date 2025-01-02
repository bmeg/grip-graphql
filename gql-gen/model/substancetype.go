
package model

import (
	"encoding/json"
	"fmt"
)

type SafeSubstanceType struct {
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Instance *string `json:"instance,omitempty"`
	Quantity *Quantity `json:"quantity,omitempty"`
	ID *string `json:"id,omitempty"`
	Language *string `json:"language,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Expiry *string `json:"expiry,omitempty"`
	Description *string `json:"description,omitempty"`
	Code *CodeableReference `json:"code,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Status *string `json:"status,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Ingredient []*SubstanceIngredient `json:"ingredient,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *SubstanceType) UnmarshalJSON(b []byte) error {
	var safe SafeSubstanceType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = SubstanceType{
		ImplicitRules: safe.ImplicitRules,
		Instance: safe.Instance,
		Quantity: safe.Quantity,
		ID: safe.ID,
		Language: safe.Language,
		ModifierExtension: safe.ModifierExtension,
		Expiry: safe.Expiry,
		Description: safe.Description,
		Code: safe.Code,
		Extension: safe.Extension,
		Meta: safe.Meta,
		Identifier: safe.Identifier,
		Text: safe.Text,
		Status: safe.Status,
		Category: safe.Category,
		Ingredient: safe.Ingredient,
		ResourceType: safe.ResourceType,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
