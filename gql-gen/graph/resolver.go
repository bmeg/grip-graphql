package graph

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"
	"fmt"

	"github.com/bmeg/grip-graphql/gql-gen/generated"
	"github.com/bmeg/grip-graphql/gql-gen/model"
)

// Organization is the resolver for the organization field.
func (r *queryResolver) Organization(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.OrganizationType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "OrganizationType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalOrganizationTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// Group is the resolver for the group field.
func (r *queryResolver) Group(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.GroupType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "GroupType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalGroupTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// Practitioner is the resolver for the practitioner field.
func (r *queryResolver) Practitioner(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.PractitionerType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "PractitionerType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalPractitionerTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// PractitionerRole is the resolver for the practitionerRole field.
func (r *queryResolver) PractitionerRole(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.PractitionerRoleType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "PractitionerRoleType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalPractitionerRoleTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// ResearchStudy is the resolver for the researchStudy field.
func (r *queryResolver) ResearchStudy(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.ResearchStudyType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "ResearchStudyType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalResearchStudyTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// Patient is the resolver for the patient field.
func (r *queryResolver) Patient(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.PatientType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "PatientType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalPatientTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// ResearchSubject is the resolver for the researchSubject field.
func (r *queryResolver) ResearchSubject(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.ResearchSubjectType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "ResearchSubjectType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalResearchSubjectTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// Substance is the resolver for the substance field.
func (r *queryResolver) Substance(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.SubstanceType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "SubstanceType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalSubstanceTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// SubstanceDefinition is the resolver for the substanceDefinition field.
func (r *queryResolver) SubstanceDefinition(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.SubstanceDefinitionType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "SubstanceDefinitionType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalSubstanceDefinitionTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// Specimen is the resolver for the specimen field.
func (r *queryResolver) Specimen(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.SpecimenType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "SpecimenType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalSpecimenTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// Observation is the resolver for the observation field.
func (r *queryResolver) Observation(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.ObservationType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "ObservationType")
	if err != nil {
		return nil, err
	}
	//fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalObservationTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}

	return slice, nil
}

// DiagnosticReport is the resolver for the diagnosticReport field.
func (r *queryResolver) DiagnosticReport(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.DiagnosticReportType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "DiagnosticReportType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalDiagnosticReportTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// Condition is the resolver for the condition field.
func (r *queryResolver) Condition(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.ConditionType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "ConditionType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalConditionTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// Medication is the resolver for the medication field.
func (r *queryResolver) Medication(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.MedicationType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "MedicationType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalMedicationTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// MedicationAdministration is the resolver for the medicationAdministration field.
func (r *queryResolver) MedicationAdministration(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.MedicationAdministrationType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "MedicationAdministrationType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalMedicationAdministrationTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// MedicationStatement is the resolver for the medicationStatement field.
func (r *queryResolver) MedicationStatement(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.MedicationStatementType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "MedicationStatementType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalMedicationStatementTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// MedicationRequest is the resolver for the medicationRequest field.
func (r *queryResolver) MedicationRequest(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.MedicationRequestType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "MedicationRequestType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalMedicationRequestTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// Procedure is the resolver for the procedure field.
func (r *queryResolver) Procedure(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.ProcedureType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "ProcedureType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalProcedureTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// DocumentReference is the resolver for the documentReference field.
func (r *queryResolver) DocumentReference(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.DocumentReferenceType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "DocumentReferenceType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalDocumentReferenceTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.TaskType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "TaskType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalTaskTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// ImagingStudy is the resolver for the imagingStudy field.
func (r *queryResolver) ImagingStudy(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.ImagingStudyType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "ImagingStudyType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalImagingStudyTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// FamilyMemberHistory is the resolver for the familyMemberHistory field.
func (r *queryResolver) FamilyMemberHistory(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.FamilyMemberHistoryType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "FamilyMemberHistoryType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalFamilyMemberHistoryTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// BodyStructure is the resolver for the bodyStructure field.
func (r *queryResolver) BodyStructure(ctx context.Context, offset *int, first *int, filter map[string]interface{}, sort map[string]interface{}, accessibility *model.Accessibility, format *model.Format) ([]model.BodyStructureType, error) {
	data, err := r.GetSelectedFieldsAst(ctx, "BodyStructureType")
	if err != nil {
		return nil, err
	}
	fmt.Println("DATA: ", data)

	slice, err := model.UnmarshalBodyStructureTypeSlice(data)
	if err != nil {
		fmt.Println("ERR: ", err)
		return nil, err
	}
	return slice, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
