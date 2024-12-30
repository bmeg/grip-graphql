# Custom Golang unmarshallers

When unmarshalling union types on graphql query functions, a custom unmarshaller needs to be used in order to maintain type correctness with the auto generated
gql-gen types that are being used. In order to create these custom types, a codegen script must be run.

Ex from this dir:

go build -o gencli

./gencli union --classes OrganizationType,GroupType,PractitionerType,PractitionerRoleType,ResearchStudyType,PatientType,ResearchSubjectType,SubstanceType,SubstanceDefinitionType,SpecimenType,ObservationType,DiagnosticReportType,ConditionType,MedicationType,MedicationAdministrationType,MedicationStatementType,MedicationRequestType,ProcedureType,DocumentReferenceType,TaskType,ImagingStudyType,FamilyMemberHistoryType,BodyStructureType --model-file ../generated.go

./gencli unmarshal --names OrganizationType,GroupType,PractitionerType,PractitionerRoleType,ResearchStudyType,PatientType,ResearchSubjectType,SubstanceType,SubstanceDefinitionType,SpecimenType,ObservationType,DiagnosticReportType,ConditionType,MedicationType,MedicationAdministrationType,MedicationStatementType,MedicationRequestType,ProcedureType,DocumentReferenceType,TaskType,ImagingStudyType,FamilyMemberHistoryType,BodyStructureType
