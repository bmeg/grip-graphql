
package model

import (
	"encoding/json"
	"fmt"
)

type SafeSubstanceType struct {
	Status *string `json:"status,omitempty"`
	Quantity *Quantity `json:"quantity,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Language *string `json:"language,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Instance *string `json:"instance,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Expiry *string `json:"expiry,omitempty"`
	Ingredient []*SubstanceIngredient `json:"ingredient,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Code *CodeableReference `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	ID *string `json:"id,omitempty"`
}

func (o *SubstanceType) UnmarshalJSON(b []byte) error {
	var safe SafeSubstanceType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = SubstanceType{
		Status: safe.Status,
		Quantity: safe.Quantity,
		Text: safe.Text,
		Identifier: safe.Identifier,
		Language: safe.Language,
		ImplicitRules: safe.ImplicitRules,
		ResourceType: safe.ResourceType,
		Instance: safe.Instance,
		ModifierExtension: safe.ModifierExtension,
		Category: safe.Category,
		Expiry: safe.Expiry,
		Ingredient: safe.Ingredient,
		Meta: safe.Meta,
		Code: safe.Code,
		Description: safe.Description,
		Extension: safe.Extension,
		ID: safe.ID,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
