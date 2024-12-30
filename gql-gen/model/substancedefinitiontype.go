
package model

import (
	"encoding/json"
	"fmt"
)

type SafeSubstanceDefinitionType struct {
	Version *string `json:"version,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Manufacturer *OrganizationType `json:"manufacturer"`
	ID *string `json:"id,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Classification []*CodeableConcept `json:"classification,omitempty"`
	Code []*SubstanceDefinitionCode `json:"code,omitempty"`
	Characterization []*SubstanceDefinitionCharacterization `json:"characterization,omitempty"`
	Supplier *OrganizationType `json:"supplier"`
	ResourceType *string `json:"resourceType,omitempty"`
	SourceMaterial *SubstanceDefinitionSourceMaterial `json:"sourceMaterial,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Status *CodeableConcept `json:"status,omitempty"`
	Property []*SubstanceDefinitionProperty `json:"property,omitempty"`
	Language *string `json:"language,omitempty"`
	Name []*SubstanceDefinitionName `json:"name,omitempty"`
	Structure *SubstanceDefinitionStructure `json:"structure,omitempty"`
	Description *string `json:"description,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	MolecularWeight []*SubstanceDefinitionMolecularWeight `json:"molecularWeight,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Relationship []*SubstanceDefinitionRelationship `json:"relationship,omitempty"`
	Domain *CodeableConcept `json:"domain,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Moiety []*SubstanceDefinitionMoiety `json:"moiety,omitempty"`
	Note TypedObject `json:"note"`
	Grade []*CodeableConcept `json:"grade,omitempty"`
}

func (o *SubstanceDefinitionType) UnmarshalJSON(b []byte) error {
	var safe SafeSubstanceDefinitionType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = SubstanceDefinitionType{
		Version: safe.Version,
		Extension: safe.Extension,
		Manufacturer: safe.Manufacturer,
		ID: safe.ID,
		Meta: safe.Meta,
		Classification: safe.Classification,
		Code: safe.Code,
		Characterization: safe.Characterization,
		Supplier: safe.Supplier,
		ResourceType: safe.ResourceType,
		SourceMaterial: safe.SourceMaterial,
		Status: safe.Status,
		Property: safe.Property,
		Language: safe.Language,
		Name: safe.Name,
		Structure: safe.Structure,
		Description: safe.Description,
		Text: safe.Text,
		MolecularWeight: safe.MolecularWeight,
		ImplicitRules: safe.ImplicitRules,
		Identifier: safe.Identifier,
		Relationship: safe.Relationship,
		Domain: safe.Domain,
		ModifierExtension: safe.ModifierExtension,
		Moiety: safe.Moiety,
		Grade: safe.Grade,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}

	return nil
}
