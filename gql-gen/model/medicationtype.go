
package model

import (
	"encoding/json"
	"fmt"
)

type SafeMedicationType struct {
	Status *string `json:"status,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Batch *MedicationBatch `json:"batch,omitempty"`
	ID *string `json:"id,omitempty"`
	MarketingAuthorizationHolder *OrganizationType `json:"marketingAuthorizationHolder"`
	Extension []*Extension `json:"extension,omitempty"`
	DoseForm *CodeableConcept `json:"doseForm,omitempty"`
	Ingredient []*MedicationIngredient `json:"ingredient,omitempty"`
	Language *string `json:"language,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	TotalVolume *Quantity `json:"totalVolume,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *MedicationType) UnmarshalJSON(b []byte) error {
	var safe SafeMedicationType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = MedicationType{
		Status: safe.Status,
		Code: safe.Code,
		Batch: safe.Batch,
		ID: safe.ID,
		MarketingAuthorizationHolder: safe.MarketingAuthorizationHolder,
		Extension: safe.Extension,
		DoseForm: safe.DoseForm,
		Ingredient: safe.Ingredient,
		Language: safe.Language,
		Identifier: safe.Identifier,
		Meta: safe.Meta,
		TotalVolume: safe.TotalVolume,
		ImplicitRules: safe.ImplicitRules,
		ResourceType: safe.ResourceType,
		Text: safe.Text,
		ModifierExtension: safe.ModifierExtension,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
