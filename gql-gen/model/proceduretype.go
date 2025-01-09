
package model

import (
	"encoding/json"
	"fmt"
)

type SafeProcedureType struct {
	Contained TypedObject `json:"contained,omitempty"`
	FocalDevice []*ProcedureFocalDevice `json:"focalDevice,omitempty"`
	Recorded *string `json:"recorded,omitempty"`
	InstantiatesURI *string `json:"instantiatesUri,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Status *string `json:"status,omitempty"`
	ReportedReference TypedObject `json:"reportedReference"`
	PartOf TypedObject `json:"partOf"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	OccurrenceTiming *Timing `json:"occurrenceTiming,omitempty"`
	StatusReason *CodeableConcept `json:"statusReason,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Used []*CodeableReference `json:"used,omitempty"`
	OccurrenceString *string `json:"occurrenceString,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	ID *string `json:"id,omitempty"`
	Outcome *CodeableConcept `json:"outcome,omitempty"`
	OccurrenceAge *Age `json:"occurrenceAge,omitempty"`
	OccurrenceRange *Range `json:"occurrenceRange,omitempty"`
	ReportedBoolean *bool `json:"reportedBoolean,omitempty"`
	Recorder TypedObject `json:"recorder"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Subject TypedObject `json:"subject"`
	Note []*Annotation `json:"note,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	SupportingInfo TypedObject `json:"supportingInfo"`
	Report TypedObject `json:"report"`
	FollowUp []*CodeableConcept `json:"followUp,omitempty"`
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`
	Focus TypedObject `json:"focus"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	Complication []*CodeableReference `json:"complication,omitempty"`
	Performer []*ProcedurePerformer `json:"performer,omitempty"`
	Language *string `json:"language,omitempty"`
	BodySite []*CodeableConcept `json:"bodySite,omitempty"`
	OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ProcedureType) UnmarshalJSON(b []byte) error {
	var safe SafeProcedureType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ProcedureType{
		FocalDevice: safe.FocalDevice,
		Recorded: safe.Recorded,
		InstantiatesURI: safe.InstantiatesURI,
		Code: safe.Code,
		Status: safe.Status,
		ImplicitRules: safe.ImplicitRules,
		ModifierExtension: safe.ModifierExtension,
		OccurrenceTiming: safe.OccurrenceTiming,
		StatusReason: safe.StatusReason,
		ResourceType: safe.ResourceType,
		Used: safe.Used,
		OccurrenceString: safe.OccurrenceString,
		Meta: safe.Meta,
		ID: safe.ID,
		Outcome: safe.Outcome,
		OccurrenceAge: safe.OccurrenceAge,
		OccurrenceRange: safe.OccurrenceRange,
		ReportedBoolean: safe.ReportedBoolean,
		Category: safe.Category,
		Note: safe.Note,
		Text: safe.Text,
		Identifier: safe.Identifier,
		Extension: safe.Extension,
		FollowUp: safe.FollowUp,
		OccurrenceDateTime: safe.OccurrenceDateTime,
		Reason: safe.Reason,
		Complication: safe.Complication,
		Performer: safe.Performer,
		Language: safe.Language,
		BodySite: safe.BodySite,
		OccurrencePeriod: safe.OccurrencePeriod,
		InstantiatesCanonical: safe.InstantiatesCanonical,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "reportedReference", safe.ReportedReference.Typename, &o.ReportedReference); err != nil {
		return fmt.Errorf("failed to unmarshal ReportedReference: %w", err)
	}
	if err := unmarshalUnion(b, "partOf", safe.PartOf.Typename, &o.PartOf); err != nil {
		return fmt.Errorf("failed to unmarshal PartOf: %w", err)
	}
	if err := unmarshalUnion(b, "recorder", safe.Recorder.Typename, &o.Recorder); err != nil {
		return fmt.Errorf("failed to unmarshal Recorder: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "supportingInfo", safe.SupportingInfo.Typename, &o.SupportingInfo); err != nil {
		return fmt.Errorf("failed to unmarshal SupportingInfo: %w", err)
	}
	if err := unmarshalUnion(b, "report", safe.Report.Typename, &o.Report); err != nil {
		return fmt.Errorf("failed to unmarshal Report: %w", err)
	}
	if err := unmarshalUnion(b, "focus", safe.Focus.Typename, &o.Focus); err != nil {
		return fmt.Errorf("failed to unmarshal Focus: %w", err)
	}

	return nil
}
