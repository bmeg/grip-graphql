
package model

import (
	"encoding/json"
	"fmt"
)

type SafeDocumentReferenceType struct {
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	SecurityLabel []*CodeableConcept `json:"securityLabel,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Date *string `json:"date,omitempty"`
	FacilityType *CodeableConcept `json:"facilityType,omitempty"`
	Attester []*DocumentReferenceAttester `json:"attester,omitempty"`
	Event []*CodeableReference `json:"event,omitempty"`
	DocStatus *string `json:"docStatus,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Period *Period `json:"period,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	Content []*DocumentReferenceContent `json:"content,omitempty"`
	Language *string `json:"language,omitempty"`
	RelatesTo []*DocumentReferenceRelatesTo `json:"relatesTo,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Custodian *OrganizationType `json:"custodian"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Modality []*CodeableConcept `json:"modality,omitempty"`
	BodySite []*CodeableReference `json:"bodySite,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Status *string `json:"status,omitempty"`
	Author TypedObject `json:"author"`
	Type *CodeableConcept `json:"type,omitempty"`
	ID *string `json:"id,omitempty"`
	Version *string `json:"version,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Subject TypedObject `json:"subject"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	PracticeSetting *CodeableConcept `json:"practiceSetting,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *DocumentReferenceType) UnmarshalJSON(b []byte) error {
	var safe SafeDocumentReferenceType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = DocumentReferenceType{
		ModifierExtension: safe.ModifierExtension,
		SecurityLabel: safe.SecurityLabel,
		Category: safe.Category,
		Date: safe.Date,
		FacilityType: safe.FacilityType,
		Attester: safe.Attester,
		Event: safe.Event,
		DocStatus: safe.DocStatus,
		Extension: safe.Extension,
		Text: safe.Text,
		Period: safe.Period,
		Content: safe.Content,
		Language: safe.Language,
		RelatesTo: safe.RelatesTo,
		Custodian: safe.Custodian,
		ImplicitRules: safe.ImplicitRules,
		Modality: safe.Modality,
		BodySite: safe.BodySite,
		Meta: safe.Meta,
		Status: safe.Status,
		Type: safe.Type,
		ID: safe.ID,
		Version: safe.Version,
		Description: safe.Description,
		ResourceType: safe.ResourceType,
		Identifier: safe.Identifier,
		PracticeSetting: safe.PracticeSetting,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "author", safe.Author.Typename, &o.Author); err != nil {
		return fmt.Errorf("failed to unmarshal Author: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}

	return nil
}
