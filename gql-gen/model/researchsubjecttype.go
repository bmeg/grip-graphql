
package model

import (
	"encoding/json"
	"fmt"
)

type SafeResearchSubjectType struct {
	Study *ResearchStudyType `json:"study"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Status *string `json:"status,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Language *string `json:"language,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ActualComparisonGroup *string `json:"actualComparisonGroup,omitempty"`
	AssignedComparisonGroup *string `json:"assignedComparisonGroup,omitempty"`
	Progress []*ResearchSubjectProgress `json:"progress,omitempty"`
	Period *Period `json:"period,omitempty"`
	Subject TypedObject `json:"subject"`
	ID *string `json:"id,omitempty"`
}

func (o *ResearchSubjectType) UnmarshalJSON(b []byte) error {
	var safe SafeResearchSubjectType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ResearchSubjectType{
		Study: safe.Study,
		ImplicitRules: safe.ImplicitRules,
		Status: safe.Status,
		Meta: safe.Meta,
		ModifierExtension: safe.ModifierExtension,
		ResourceType: safe.ResourceType,
		Language: safe.Language,
		Text: safe.Text,
		Identifier: safe.Identifier,
		Extension: safe.Extension,
		ActualComparisonGroup: safe.ActualComparisonGroup,
		AssignedComparisonGroup: safe.AssignedComparisonGroup,
		Progress: safe.Progress,
		Period: safe.Period,
		ID: safe.ID,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}

	return nil
}
