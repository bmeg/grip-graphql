package graph

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"

	"github.com/bmeg/grip-graphql/gql-gen/generated"

	"github.com/bmeg/grip-graphql/gql-gen/model"
)

// Organization is the resolver for the organization field.
func (r *queryResolver) Organization(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.OrganizationType, error) {
	return nil, nil
}

// Group is the resolver for the group field.
func (r *queryResolver) Group(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.GroupType, error) {
	panic("not implemented")
}

// Practitioner is the resolver for the practitioner field.
func (r *queryResolver) Practitioner(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.PractitionerType, error) {
	panic("not implemented")
}

// PractitionerRole is the resolver for the practitionerRole field.
func (r *queryResolver) PractitionerRole(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.PractitionerRoleType, error) {
	panic("not implemented")
}

// ResearchStudy is the resolver for the researchStudy field.
func (r *queryResolver) ResearchStudy(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.ResearchStudyType, error) {
	panic("not implemented")
}

// Patient is the resolver for the patient field.
func (r *queryResolver) Patient(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.PatientType, error) {
	panic("not implemented")
}

// ResearchSubject is the resolver for the researchSubject field.
func (r *queryResolver) ResearchSubject(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.ResearchSubjectType, error) {
	panic("not implemented")
}

// Substance is the resolver for the substance field.
func (r *queryResolver) Substance(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.SubstanceType, error) {
	panic("not implemented")
}

// SubstanceDefinition is the resolver for the substanceDefinition field.
func (r *queryResolver) SubstanceDefinition(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.SubstanceDefinitionType, error) {
	panic("not implemented")
}

// Specimen is the resolver for the specimen field.
func (r *queryResolver) Specimen(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.SpecimenType, error) {
	panic("not implemented")
}

// Observation is the resolver for the observation field.
func (r *queryResolver) Observation(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.ObservationType, error) {
	//sourceType := "ObservationType"
	//fields := GetQueryFields(ctx, sourceType)
	//res := r.gripQuery(fields, sourceType)
	//fmt.Println("RES: ", res)
	r.GetSelectedFieldsAst(ctx, "ObservationType")

	/*for _, field := range fields {
		fmt.Println("PATH: ", field)
	}*/
	return nil, nil
}

// DiagnosticReport is the resolver for the diagnosticReport field.
func (r *queryResolver) DiagnosticReport(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.DiagnosticReportType, error) {
	panic("not implemented")
}

// Condition is the resolver for the condition field.
func (r *queryResolver) Condition(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.ConditionType, error) {
	panic("not implemented")
}

// Medication is the resolver for the medication field.
func (r *queryResolver) Medication(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.MedicationType, error) {
	panic("not implemented")
}

// MedicationAdministration is the resolver for the medicationAdministration field.
func (r *queryResolver) MedicationAdministration(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.MedicationAdministrationType, error) {
	panic("not implemented")
}

// MedicationStatement is the resolver for the medicationStatement field.
func (r *queryResolver) MedicationStatement(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.MedicationStatementType, error) {
	panic("not implemented")
}

// MedicationRequest is the resolver for the medicationRequest field.
func (r *queryResolver) MedicationRequest(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.MedicationRequestType, error) {
	panic("not implemented")
}

// Procedure is the resolver for the procedure field.
func (r *queryResolver) Procedure(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.ProcedureType, error) {
	panic("not implemented")
}

// DocumentReference is the resolver for the documentReference field.
func (r *queryResolver) DocumentReference(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.DocumentReferenceType, error) {
	panic("not implemented")
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.TaskType, error) {
	panic("not implemented")
}

// ImagingStudy is the resolver for the imagingStudy field.
func (r *queryResolver) ImagingStudy(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.ImagingStudyType, error) {
	panic("not implemented")
}

// FamilyMemberHistory is the resolver for the familyMemberHistory field.
func (r *queryResolver) FamilyMemberHistory(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.FamilyMemberHistoryType, error) {
	panic("not implemented")
}

// BodyStructure is the resolver for the bodyStructure field.
func (r *queryResolver) BodyStructure(ctx context.Context, offset *int, first *int, filter *string, sort *string, accessibility *model.Accessibility, format *model.Format) ([]*model.BodyStructureType, error) {
	panic("not implemented")
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	type Resolver struct{}
*/
