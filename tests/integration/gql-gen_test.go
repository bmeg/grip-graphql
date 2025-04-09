package integration_test

import (
	"strings"
	"testing"

	"github.com/bmeg/grip-graphql/tests/integration"
)

func Test_GqlGen_Basic_Graphql(t *testing.T) {
	/* A basic test for Graphql style query with mock auth */
	payload := []byte(`{
	    "query": "query { specimen {    id  }}",
	    "variables": {
	        "limit": 10
	    }
	}`)

	req := &integration.Request{
		Url:    "http://localhost:8201/graphql/query",
		Method: "POST",
		Headers: map[string]any{
			"Authorization": integration.CreateToken(false, false, true),
			"Content-Type":  "application/json"},
		Body: payload,
	}

	response, status := integration.TemplateRequest(req, t)
	if !status {
		t.Error("Status returned false: ", response)
	}

	correct_response, ok := response["data"].(map[string]any)["specimen"]
	if !ok {
		t.Error("Response not indexable for 'data' and/or 'specimen' keys")
	}
	// Only expect to get 1 entry back
	for i, resp := range correct_response.([]any) {
		id, ok := resp.(map[string]any)["id"]
		if !ok {
			t.Error("Response not indexable for 'id' keys")
		}
		if i == 0 && id != "60c67a06-ea2d-4d24-9249-418dc77a16a9" {
			t.Error("Expected 60c67a06-ea2d-4d24-9249-418dc77a16a9, but got", correct_response, "instead")
		}
	}
}

func Test_GqlGen_Bad_Filter_Logical_Operator(t *testing.T) {
	/* A basic test for Graphql style query with mock auth */
	payload := []byte(`{
	    "query": "query($filter: JSON) { specimen(filter: $filter first:1){    id  }}",
	    "variables": {
						"filter": {
							"orr": [
								{
									"!=": {
										"Specimen.processing.method.coding.display": "some code"
									}
								},
								{
									">": {
										"Specimen.collection.bodySite.concept.coding.code": "123456"
									}
								}
							]
						}
	    }
	}`)

	req := &integration.Request{
		Url:    "http://localhost:8201/graphql/query",
		Method: "POST",
		Headers: map[string]any{
			"Authorization": integration.CreateToken(false, false, true),
			"Content-Type":  "application/json"},
		Body: payload,
	}

	response, status := integration.TemplateRequest(req, t)
	t.Log("RESP: ", response)
	if !status {
		t.Error("Status returned false: ", response)
	}

	correct_response, ok := response["data"]
	if !ok {
		t.Error("Response not indexable for 'data' keys")
	}
	if correct_response != nil {
		t.Errorf("Expecting nil but got %s instead\n", correct_response)
	}

	errors, ok := response["errors"].([]any)[0].(map[string]any)["message"]
	if !ok {
		t.Error("Response not indexable for 'errors', or 'message keys")
	}
	if errors != "invalid logical operator 'orr'" {
		t.Errorf("invalid logical operator 'orr' but got %s instead", errors)
	}
}

func Test_GqlGen_Filter_Ok(t *testing.T) {
	/* A basic test for Graphql style query with mock auth */
	payload := []byte(`{
	    "query": "query($filter: JSON){observation(filter: $filter first:10){id component{ valueInteger code{ coding{ code}}}}}",
	    "variables": {
						"filter": {
							"and": [
								{
									"and": [
										{
											">": {
												"Observation.component.valueInteger": "364"
											}
										},
										{
											"=": {
												"Observation.component.code.coding.code": "indexed_collection_date"
											}
										}
									]
								},
								{
									"!=": {
										"Observation.id": "21f3411d-89a4-4bcc-9ce7-b76edb1c745f"
									}
								}
							]
						}
					}}`)

	req := &integration.Request{
		Url:    "http://localhost:8201/graphql/query",
		Method: "POST",
		Headers: map[string]any{
			"Authorization": integration.CreateToken(false, false, true),
			"Content-Type":  "application/json"},
		Body: payload,
	}

	response, status := integration.TemplateRequest(req, t)
	t.Log("RESP: ", response)
	if !status {
		t.Error("Status returned false: ", response)
	}

	correct_response, ok := response["data"].(map[string]any)["observation"].([]any)
	if !ok {
		t.Error("Response not indexable for 'data' keys")
	}
	if len(correct_response) != 1 {
		t.Error("Expecting response to be of length 1 not ", len(correct_response))

	}
}

func Test_GqlGen_Edge_Traversal_Ok(t *testing.T) {
	/* A basic test for Graphql style query with mock auth */
	query := strings.ReplaceAll(`query {
		  observation(first: 10) {
		    id
		    focus {
		      ... on DocumentReferenceType {
		        id
		        content {
		          attachment {
		            extension {
		              url
		              valueString
		              valueUrl
		            }
		          }
		        }
		      }
		    }
		    subject {
		      ... on PatientType {
		        id
		        identifier {
		          system
		          value
		        }
		      }
		    }
		  }
	}`, "\n", "")
	query = strings.ReplaceAll(query, "\t", "")

	payload := []byte(`{ "query": "` + query + `" }`)
	req := &integration.Request{
		Url:    "http://localhost:8201/graphql/query",
		Method: "POST",
		Headers: map[string]any{
			"Authorization": integration.CreateToken(false, false, true),
			"Content-Type":  "application/json"},
		Body: payload,
	}

	response, status := integration.TemplateRequest(req, t)
	data, ok := response["data"].(map[string]any)["observation"].([]any)
	if !ok {
		t.Error("observation index not found in data")
	}
	if len(data) != 2 {
		t.Errorf("expected 2 observations, found %d\n", len(data))
	}
	focus_content_attachement_extension, ok := data[0].(map[string]any)["focus"].(map[string]any)["content"].([]any)[0].(map[string]any)["attachment"].(map[string]any)["extension"].([]any)
	if !ok {
		t.Error("observation focus docref attachment not found")
	}
	t.Log("focus content found: ", focus_content_attachement_extension)

	obs_id, ok := data[0].(map[string]any)["id"]
	if !ok {
		t.Error("observation id not found")
	}
	obs_focus_docref_id, ok := data[0].(map[string]any)["focus"].(map[string]any)["id"]
	if !ok {
		t.Error("observation focus docref id not found")
	}
	if obs_id == obs_focus_docref_id {
		t.Errorf("observation id: %s and observation focus docref vertex id: %s should be different unique ids\n", obs_id, obs_focus_docref_id)
	}
	if !status {
		t.Error("Status returned false: ", response)
	}
}

func Test_GqlGen_Nested_Edge_Traversal_Ok(t *testing.T) {
	/* This test also acts as an auth test because there contains extra fields in the conformance data that are not in the TEST project */
	query := strings.ReplaceAll(`query{
	  observation(first:10){
	    id
	    focus{
	      ... on DocumentReferenceType{
	        id
	        content{
	          attachment{
	            extension{
	              url
	              valueString
	              valueUrl
	            }
	          }
	        }
	        subject{
	          ... on SpecimenType{
				id
	            processing{
	              method{
	                coding{
	                  code
	                  system
	                  display
	                }
	              }
	            }
	          }
	        }
	      }
	    }
	  }
	}`, "\n", "")
	query = strings.ReplaceAll(query, "\t", "")

	payload := []byte(`{ "query": "` + query + `" }`)
	req := &integration.Request{
		Url:    "http://localhost:8201/graphql/query",
		Method: "POST",
		Headers: map[string]any{
			"Authorization": integration.CreateToken(false, false, true),
			"Content-Type":  "application/json"},
		Body: payload,
	}

	response, status := integration.TemplateRequest(req, t)
	data, ok := response["data"].(map[string]any)["observation"].([]any)
	if !ok {
		t.Error("observation index not found in data")
	}
	obs_docref_specimen_id, ok := data[0].(map[string]any)["focus"].(map[string]any)["subject"].(map[string]any)
	if !ok {
		t.Error("observation docref specimen not found in data")
	}
	specimen_processing_code, ok := obs_docref_specimen_id["processing"].([]any)[0].(map[string]any)["method"].(map[string]any)["coding"].([]any)
	if !ok {
		t.Error("observation docref specimen specimen_processing_code not found in data")
	}
	t.Log("2x edge traversal specimen_processing_code data", specimen_processing_code)

	if len(data) != 2 {
		t.Errorf("expected 2 observations, found %d\n", len(data))
	}
	if !status {
		t.Error("Status returned false: ", response)
	}

	for _, row := range data {
		t.Log("Row: ", row)
		id, ok := row.(map[string]any)["id"]
		if !ok {
			t.Error("id not indexable on type observation")
		}
		if id == "11f3411d-89a4-4bcc-9ce7-b76edb1c745f" {
			t.Error("Non test project row detected. Auth filtering failed")
		}
		obs_docref_id, ok := row.(map[string]any)["focus"].(map[string]any)["id"]
		if !ok {
			t.Error("id not indexable on type obs_docref id")
		}
		if obs_docref_id == "8ae7e542-767f-4b03-a854-7ceed17152cb" {
			t.Error("Non test project row detected. Auth filtering failed")
		}

		obs_docref_specimen_id, ok := row.(map[string]any)["focus"].(map[string]any)["subject"].(map[string]any)["id"]
		if !ok {
			t.Error("id not indexable on type obs_docref_specimen id")
		}
		if obs_docref_specimen_id == "50c67a06-ea2d-4d24-9249-418dc77a16a9" {
			t.Error("Non test project row detected. Auth filtering failed")
		}
	}
}

func Test_GqlGen_Nested_Edge_Traversal_Filter_Ok(t *testing.T) {
	payload := strings.ReplaceAll(`{
	    "query": "query($filter: JSON){observation(filter: $filter first:10){id component{ valueInteger code{ coding{ code}}} 	focus {... on DocumentReferenceType {id auth_resource_path content { attachment { extension { url valueString valueUrl}}}}}}}",
	    "variables": {
						"filter": {
							"or": [
								{
									"and": [
										{
											"=": {
												"Observation.component.code.coding.code": "indexed_collection_date"
											}
										}
									]
								},
								{
									"!=": {
										"DocumentReference.content.attachment.extension.valueString": "227f0a5379362d42eaa1814cfc0101b8"
									}
								}
							]
						}
					}}`, "\t", "")
	req := &integration.Request{
		Url:    "http://localhost:8201/graphql/query",
		Method: "POST",
		Headers: map[string]any{
			"Authorization": integration.CreateToken(false, false, true),
			"Content-Type":  "application/json"},
		Body: []byte(payload),
	}

	response, status := integration.TemplateRequest(req, t)
	t.Log("RESP: ", response)
	if !status {
		t.Error("Status returned false: ", response)
	}

	data, ok := response["data"].(map[string]any)["observation"].([]any)
	if !ok {
		t.Error("observation index not found in data")
	}
	for _, row := range data {
		valueString, ok := row.(map[string]any)["focus"].(map[string]any)["content"].([]any)[0].(map[string]any)["attachment"].(map[string]any)["extension"].([]any)[0].(map[string]any)["valueString"]
		if !ok {
			t.Error("error in traversal, docref doesn't contain valuestring")
		}
		if valueString == "227f0a5379362d42eaa1814cfc0101b8" {
			t.Error("filter asks for value string to not be 227f0a5379362d42eaa1814cfc0101b8 yet valuestring is 227f0a5379362d42eaa1814cfc0101b8")
		}
	}
}
