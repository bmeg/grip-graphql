
package model

import (
	"encoding/json"
	"fmt"
)

type SafeResearchStudyType struct {
	Language *string `json:"language,omitempty"`
	Region []*CodeableConcept `json:"region,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Note TypedObject `json:"note"`
	Focus []*CodeableReference `json:"focus,omitempty"`
	PrimaryPurposeType *CodeableConcept `json:"primaryPurposeType,omitempty"`
	ProgressStatus []*ResearchStudyProgressStatus `json:"progressStatus,omitempty"`
	WhyStopped *CodeableConcept `json:"whyStopped,omitempty"`
	Result TypedObject `json:"result"`
	Version *string `json:"version,omitempty"`
	URL *string `json:"url,omitempty"`
	RelatedArtifact []*RelatedArtifact `json:"relatedArtifact,omitempty"`
	Name *string `json:"name,omitempty"`
	StudyDesign []*CodeableConcept `json:"studyDesign,omitempty"`
	Objective []*ResearchStudyObjective `json:"objective,omitempty"`
	Recruitment TypedObject `json:"recruitment"`
	Title *string `json:"title,omitempty"`
	AssociatedParty TypedObject `json:"associatedParty"`
	ID *string `json:"id,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Description *string `json:"description,omitempty"`
	ComparisonGroup []*ResearchStudyComparisonGroup `json:"comparisonGroup,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	PartOf *ResearchStudyType `json:"partOf"`
	Classifier []*CodeableConcept `json:"classifier,omitempty"`
	Label []*ResearchStudyLabel `json:"label,omitempty"`
	Keyword []*CodeableConcept `json:"keyword,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Period *Period `json:"period,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Date *string `json:"date,omitempty"`
	DescriptionSummary *string `json:"descriptionSummary,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Phase *CodeableConcept `json:"phase,omitempty"`
	Condition []*CodeableConcept `json:"condition,omitempty"`
	Status *string `json:"status,omitempty"`
	OutcomeMeasure []*ResearchStudyOutcomeMeasure `json:"outcomeMeasure,omitempty"`
	Site TypedObject `json:"site"`
}

func (o *ResearchStudyType) UnmarshalJSON(b []byte) error {
	var safe SafeResearchStudyType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ResearchStudyType{
		Language: safe.Language,
		Region: safe.Region,
		Identifier: safe.Identifier,
		Focus: safe.Focus,
		PrimaryPurposeType: safe.PrimaryPurposeType,
		ProgressStatus: safe.ProgressStatus,
		WhyStopped: safe.WhyStopped,
		Version: safe.Version,
		URL: safe.URL,
		RelatedArtifact: safe.RelatedArtifact,
		Name: safe.Name,
		StudyDesign: safe.StudyDesign,
		Objective: safe.Objective,
		Title: safe.Title,
		ID: safe.ID,
		Description: safe.Description,
		ComparisonGroup: safe.ComparisonGroup,
		ImplicitRules: safe.ImplicitRules,
		PartOf: safe.PartOf,
		Classifier: safe.Classifier,
		Label: safe.Label,
		Keyword: safe.Keyword,
		ModifierExtension: safe.ModifierExtension,
		Period: safe.Period,
		ResourceType: safe.ResourceType,
		Date: safe.Date,
		DescriptionSummary: safe.DescriptionSummary,
		Extension: safe.Extension,
		Text: safe.Text,
		Meta: safe.Meta,
		Phase: safe.Phase,
		Condition: safe.Condition,
		Status: safe.Status,
		OutcomeMeasure: safe.OutcomeMeasure,
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "result", safe.Result.Typename, &o.Result); err != nil {
		return fmt.Errorf("failed to unmarshal Result: %w", err)
	}
	if err := unmarshalUnion(b, "recruitment", safe.Recruitment.Typename, &o.Recruitment); err != nil {
		return fmt.Errorf("failed to unmarshal Recruitment: %w", err)
	}
	if err := unmarshalUnion(b, "associatedParty", safe.AssociatedParty.Typename, &o.AssociatedParty); err != nil {
		return fmt.Errorf("failed to unmarshal AssociatedParty: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "site", safe.Site.Typename, &o.Site); err != nil {
		return fmt.Errorf("failed to unmarshal Site: %w", err)
	}

	return nil
}
