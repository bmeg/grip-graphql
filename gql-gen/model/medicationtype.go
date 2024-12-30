
package model

import (
	"encoding/json"
	"fmt"
)

type SafeMedicationType struct {
	Meta *Meta `json:"meta,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	MarketingAuthorizationHolder *OrganizationType `json:"marketingAuthorizationHolder"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	TotalVolume *Quantity `json:"totalVolume,omitempty"`
	Language *string `json:"language,omitempty"`
	DoseForm *CodeableConcept `json:"doseForm,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Ingredient []*MedicationIngredient `json:"ingredient,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ID *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	Batch *MedicationBatch `json:"batch,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
}

func (o *MedicationType) UnmarshalJSON(b []byte) error {
	var safe SafeMedicationType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = MedicationType{
		Meta: safe.Meta,
		Text: safe.Text,
		MarketingAuthorizationHolder: safe.MarketingAuthorizationHolder,
		ModifierExtension: safe.ModifierExtension,
		TotalVolume: safe.TotalVolume,
		Language: safe.Language,
		DoseForm: safe.DoseForm,
		ResourceType: safe.ResourceType,
		ImplicitRules: safe.ImplicitRules,
		Ingredient: safe.Ingredient,
		Extension: safe.Extension,
		Identifier: safe.Identifier,
		ID: safe.ID,
		Status: safe.Status,
		Batch: safe.Batch,
		Code: safe.Code,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
