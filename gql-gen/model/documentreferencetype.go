
package model

import (
	"encoding/json"
	"fmt"
)

type SafeDocumentReferenceType struct {
	DocStatus *string `json:"docStatus,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Event []*CodeableReference `json:"event,omitempty"`
	BodySite []*CodeableReference `json:"bodySite,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Content []*DocumentReferenceContent `json:"content,omitempty"`
	Custodian *OrganizationType `json:"custodian"`
	SecurityLabel []*CodeableConcept `json:"securityLabel,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	Language *string `json:"language,omitempty"`
	FacilityType *CodeableConcept `json:"facilityType,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Subject TypedObject `json:"subject"`
	PracticeSetting *CodeableConcept `json:"practiceSetting,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Description *string `json:"description,omitempty"`
	RelatesTo []*DocumentReferenceRelatesTo `json:"relatesTo,omitempty"`
	Attester []*DocumentReferenceAttester `json:"attester,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Period *Period `json:"period,omitempty"`
	Status *string `json:"status,omitempty"`
	ID *string `json:"id,omitempty"`
	Modality []*CodeableConcept `json:"modality,omitempty"`
	Version *string `json:"version,omitempty"`
	Type *CodeableConcept `json:"type,omitempty"`
	Author TypedObject `json:"author"`
	Date *string `json:"date,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *DocumentReferenceType) UnmarshalJSON(b []byte) error {
	var safe SafeDocumentReferenceType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = DocumentReferenceType{
		DocStatus: safe.DocStatus,
		ModifierExtension: safe.ModifierExtension,
		Event: safe.Event,
		BodySite: safe.BodySite,
		Extension: safe.Extension,
		Content: safe.Content,
		Custodian: safe.Custodian,
		SecurityLabel: safe.SecurityLabel,
		Language: safe.Language,
		FacilityType: safe.FacilityType,
		Meta: safe.Meta,
		Category: safe.Category,
		PracticeSetting: safe.PracticeSetting,
		Text: safe.Text,
		ImplicitRules: safe.ImplicitRules,
		Description: safe.Description,
		RelatesTo: safe.RelatesTo,
		Attester: safe.Attester,
		ResourceType: safe.ResourceType,
		Period: safe.Period,
		Status: safe.Status,
		ID: safe.ID,
		Modality: safe.Modality,
		Version: safe.Version,
		Type: safe.Type,
		Date: safe.Date,
		Identifier: safe.Identifier,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "author", safe.Author.Typename, &o.Author); err != nil {
		return fmt.Errorf("failed to unmarshal Author: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
