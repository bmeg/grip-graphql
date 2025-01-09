
package model

import (
	"encoding/json"
	"fmt"
)

type SafeResearchStudyType struct {
	Recruitment *ResearchStudyRecruitment `json:"recruitment,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Classifier []*CodeableConcept `json:"classifier,omitempty"`
	AssociatedParty []*ResearchStudyAssociatedParty `json:"associatedParty,omitempty"`
	Site TypedObject `json:"site"`
	Keyword []*CodeableConcept `json:"keyword,omitempty"`
	StudyDesign []*CodeableConcept `json:"studyDesign,omitempty"`
	Condition []*CodeableConcept `json:"condition,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Version *string `json:"version,omitempty"`
	Status *string `json:"status,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Period *Period `json:"period,omitempty"`
	ProgressStatus []*ResearchStudyProgressStatus `json:"progressStatus,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Phase *CodeableConcept `json:"phase,omitempty"`
	Date *string `json:"date,omitempty"`
	ComparisonGroup []*ResearchStudyComparisonGroup `json:"comparisonGroup,omitempty"`
	Title *string `json:"title,omitempty"`
	Name *string `json:"name,omitempty"`
	RelatedArtifact []*RelatedArtifact `json:"relatedArtifact,omitempty"`
	Label []*ResearchStudyLabel `json:"label,omitempty"`
	URL *string `json:"url,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Description *string `json:"description,omitempty"`
	OutcomeMeasure []*ResearchStudyOutcomeMeasure `json:"outcomeMeasure,omitempty"`
	Result TypedObject `json:"result"`
	Extension []*Extension `json:"extension,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	ID *string `json:"id,omitempty"`
	Objective []*ResearchStudyObjective `json:"objective,omitempty"`
	PartOf *ResearchStudyType `json:"partOf"`
	PrimaryPurposeType *CodeableConcept `json:"primaryPurposeType,omitempty"`
	Language *string `json:"language,omitempty"`
	DescriptionSummary *string `json:"descriptionSummary,omitempty"`
	WhyStopped *CodeableConcept `json:"whyStopped,omitempty"`
	Region []*CodeableConcept `json:"region,omitempty"`
	Focus []*CodeableReference `json:"focus,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ResearchStudyType) UnmarshalJSON(b []byte) error {
	var safe SafeResearchStudyType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ResearchStudyType{
		Recruitment: safe.Recruitment,
		Text: safe.Text,
		Classifier: safe.Classifier,
		AssociatedParty: safe.AssociatedParty,
		Keyword: safe.Keyword,
		StudyDesign: safe.StudyDesign,
		Condition: safe.Condition,
		Meta: safe.Meta,
		Version: safe.Version,
		Status: safe.Status,
		Period: safe.Period,
		ProgressStatus: safe.ProgressStatus,
		Identifier: safe.Identifier,
		ImplicitRules: safe.ImplicitRules,
		Phase: safe.Phase,
		Date: safe.Date,
		ComparisonGroup: safe.ComparisonGroup,
		Title: safe.Title,
		Name: safe.Name,
		RelatedArtifact: safe.RelatedArtifact,
		Label: safe.Label,
		URL: safe.URL,
		ModifierExtension: safe.ModifierExtension,
		ResourceType: safe.ResourceType,
		Description: safe.Description,
		OutcomeMeasure: safe.OutcomeMeasure,
		Extension: safe.Extension,
		Note: safe.Note,
		ID: safe.ID,
		Objective: safe.Objective,
		PartOf: safe.PartOf,
		PrimaryPurposeType: safe.PrimaryPurposeType,
		Language: safe.Language,
		DescriptionSummary: safe.DescriptionSummary,
		WhyStopped: safe.WhyStopped,
		Region: safe.Region,
		Focus: safe.Focus,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "site", safe.Site.Typename, &o.Site); err != nil {
		return fmt.Errorf("failed to unmarshal Site: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "result", safe.Result.Typename, &o.Result); err != nil {
		return fmt.Errorf("failed to unmarshal Result: %w", err)
	}

	return nil
}
