
package model

import (
	"encoding/json"
	"fmt"
)

type SafeSubstanceDefinitionType struct {
	Version *string `json:"version,omitempty"`
	Supplier *OrganizationType `json:"supplier"`
	Contained TypedObject `json:"contained,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Relationship []*SubstanceDefinitionRelationship `json:"relationship,omitempty"`
	Name []*SubstanceDefinitionName `json:"name,omitempty"`
	SourceMaterial *SubstanceDefinitionSourceMaterial `json:"sourceMaterial,omitempty"`
	Code []*SubstanceDefinitionCode `json:"code,omitempty"`
	Note TypedObject `json:"note"`
	MolecularWeight []*SubstanceDefinitionMolecularWeight `json:"molecularWeight,omitempty"`
	Property []*SubstanceDefinitionProperty `json:"property,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Grade []*CodeableConcept `json:"grade,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Manufacturer *OrganizationType `json:"manufacturer"`
	Moiety []*SubstanceDefinitionMoiety `json:"moiety,omitempty"`
	Status *CodeableConcept `json:"status,omitempty"`
	Characterization []*SubstanceDefinitionCharacterization `json:"characterization,omitempty"`
	Classification []*CodeableConcept `json:"classification,omitempty"`
	Domain *CodeableConcept `json:"domain,omitempty"`
	Structure *SubstanceDefinitionStructure `json:"structure,omitempty"`
	Description *string `json:"description,omitempty"`
	Language *string `json:"language,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	ID *string `json:"id,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *SubstanceDefinitionType) UnmarshalJSON(b []byte) error {
	var safe SafeSubstanceDefinitionType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = SubstanceDefinitionType{
		Version: safe.Version,
		Supplier: safe.Supplier,
		Identifier: safe.Identifier,
		Relationship: safe.Relationship,
		Name: safe.Name,
		SourceMaterial: safe.SourceMaterial,
		Code: safe.Code,
		MolecularWeight: safe.MolecularWeight,
		Property: safe.Property,
		ResourceType: safe.ResourceType,
		Grade: safe.Grade,
		Extension: safe.Extension,
		Manufacturer: safe.Manufacturer,
		Moiety: safe.Moiety,
		Status: safe.Status,
		Characterization: safe.Characterization,
		Classification: safe.Classification,
		Domain: safe.Domain,
		Structure: safe.Structure,
		Description: safe.Description,
		Language: safe.Language,
		ModifierExtension: safe.ModifierExtension,
		Meta: safe.Meta,
		Text: safe.Text,
		ImplicitRules: safe.ImplicitRules,
		ID: safe.ID,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}

	return nil
}
