
package model

import (
	"encoding/json"
	"fmt"
)

type SafeProcedureType struct {
	InstantiatesURI *string `json:"instantiatesUri,omitempty"`
	ID *string `json:"id,omitempty"`
	Language *string `json:"language,omitempty"`
	PartOf TypedObject `json:"partOf"`
	OccurrenceRange *Range `json:"occurrenceRange,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`
	ReportedReference TypedObject `json:"reportedReference"`
	Report TypedObject `json:"report"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	Outcome *CodeableConcept `json:"outcome,omitempty"`
	ReportedBoolean *string `json:"reportedBoolean,omitempty"`
	BodySite []*CodeableConcept `json:"bodySite,omitempty"`
	SupportingInfo TypedObject `json:"supportingInfo"`
	Extension []*Extension `json:"extension,omitempty"`
	Used []*CodeableReference `json:"used,omitempty"`
	OccurrenceString *string `json:"occurrenceString,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	FocalDevice []*ProcedureFocalDevice `json:"focalDevice,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`
	FollowUp []*CodeableConcept `json:"followUp,omitempty"`
	OccurrenceAge *Age `json:"occurrenceAge,omitempty"`
	OccurrenceTiming *Timing `json:"occurrenceTiming,omitempty"`
	Subject TypedObject `json:"subject"`
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Recorder TypedObject `json:"recorder"`
	Status *string `json:"status,omitempty"`
	Complication []*CodeableReference `json:"complication,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	StatusReason *CodeableConcept `json:"statusReason,omitempty"`
	Focus TypedObject `json:"focus"`
	Performer []*ProcedurePerformer `json:"performer,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Recorded *string `json:"recorded,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ProcedureType) UnmarshalJSON(b []byte) error {
	var safe SafeProcedureType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ProcedureType{
		InstantiatesURI: safe.InstantiatesURI,
		ID: safe.ID,
		Language: safe.Language,
		OccurrenceRange: safe.OccurrenceRange,
		Note: safe.Note,
		Category: safe.Category,
		OccurrencePeriod: safe.OccurrencePeriod,
		Reason: safe.Reason,
		Outcome: safe.Outcome,
		ReportedBoolean: safe.ReportedBoolean,
		BodySite: safe.BodySite,
		Extension: safe.Extension,
		Used: safe.Used,
		OccurrenceString: safe.OccurrenceString,
		FocalDevice: safe.FocalDevice,
		Identifier: safe.Identifier,
		InstantiatesCanonical: safe.InstantiatesCanonical,
		FollowUp: safe.FollowUp,
		OccurrenceAge: safe.OccurrenceAge,
		OccurrenceTiming: safe.OccurrenceTiming,
		OccurrenceDateTime: safe.OccurrenceDateTime,
		ImplicitRules: safe.ImplicitRules,
		Status: safe.Status,
		Complication: safe.Complication,
		Meta: safe.Meta,
		ResourceType: safe.ResourceType,
		StatusReason: safe.StatusReason,
		Performer: safe.Performer,
		Text: safe.Text,
		Code: safe.Code,
		ModifierExtension: safe.ModifierExtension,
		Recorded: safe.Recorded,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "partOf", safe.PartOf.Typename, &o.PartOf); err != nil {
		return fmt.Errorf("failed to unmarshal PartOf: %w", err)
	}
	if err := unmarshalUnion(b, "reportedReference", safe.ReportedReference.Typename, &o.ReportedReference); err != nil {
		return fmt.Errorf("failed to unmarshal ReportedReference: %w", err)
	}
	if err := unmarshalUnion(b, "report", safe.Report.Typename, &o.Report); err != nil {
		return fmt.Errorf("failed to unmarshal Report: %w", err)
	}
	if err := unmarshalUnion(b, "supportingInfo", safe.SupportingInfo.Typename, &o.SupportingInfo); err != nil {
		return fmt.Errorf("failed to unmarshal SupportingInfo: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "recorder", safe.Recorder.Typename, &o.Recorder); err != nil {
		return fmt.Errorf("failed to unmarshal Recorder: %w", err)
	}
	if err := unmarshalUnion(b, "focus", safe.Focus.Typename, &o.Focus); err != nil {
		return fmt.Errorf("failed to unmarshal Focus: %w", err)
	}

	return nil
}
