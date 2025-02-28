package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/bmeg/grip/gripql"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/protobuf/encoding/protojson"
)

type Request struct {
	url     string
	method  string
	headers map[string]any
	body    []byte
}

func createToken(expired bool, writer bool, reader bool) string {
	var create string
	timeNow := time.Now()
	time_exp := timeNow
	if !expired {
		time_exp = timeNow.Add(time.Minute * 20)
	}
	if writer {
		create = "create"
	}
	if reader {
		create = "reader"
	}
	if writer && reader {
		create = "create-reader"
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":      "https://foobar-domain/user",
			"username": "foobar-user",
			"iat":      timeNow.Unix(),
			"exp":      time_exp.Unix(),
			"jti":      "foobbar-jti",
			"sub":      create,
		})
	tokenString, err := token.SignedString([]byte("foo-bar-signature"))
	if err != nil {
		fmt.Println("Error creating signed string: ", err)
	}
	return tokenString
}

func TemplateRequest(request *Request, t *testing.T) (response_json map[string]any, status bool) {
	/* A templating function that inserts all of the arguments that would are needed to do an http request */

	req, err := http.NewRequest(request.method, request.url, bytes.NewBuffer(request.body))
	if err != nil {
		t.Error("Error creating request:", err)
		return
	}
	for key, val := range request.headers {
		req.Header.Set(key, val.(string))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Error sending request:", err)
		return nil, false
	}
	defer resp.Body.Close()

	t.Log("Response Status:", resp.Status)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		t.Error("Error reading response:", err)
		return nil, false
	}

	var data map[string]interface{}
	errors := json.Unmarshal([]byte(buf.String()), &data)
	if errors != nil {
		t.Error("Error: ", errors)
		return nil, false
	}
	return data, true
}

func bulkLoad(url, directoryPath, accessToken string) (error, []map[string]any) {
	files, _ := os.ReadDir(directoryPath)
	var allData []map[string]any
	for _, file := range files {
		if !file.IsDir() && (filepath.Ext(file.Name()) == ".json" || filepath.Ext(file.Name()) == ".gz" || filepath.Ext(file.Name()) == ".ndjson") {
			filePath := filepath.Join(directoryPath, file.Name())
			graphComponent := "vertex"
			if strings.Contains(filePath, "edge") {
				graphComponent = "edge"
			}

			file, _ := os.Open(filePath)
			defer file.Close()

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			part, _ := writer.CreateFormFile("file", filepath.Base(filePath))
			io.Copy(part, file)
			writer.WriteField("types", graphComponent)
			writer.Close()

			req, _ := http.NewRequest("POST", url, body)
			req.Header.Set("Authorization", accessToken)
			req.Header.Set("Content-Type", writer.FormDataContentType())

			client := &http.Client{}
			resp, _ := client.Do(req)
			defer resp.Body.Close()

			var err error
			buf := new(bytes.Buffer)
			_, err = buf.ReadFrom(resp.Body)
			if err != nil {
				return err, nil
			}

			var data map[string]interface{}
			errors := json.Unmarshal([]byte(buf.String()), &data)
			if errors != nil {
				return errors, nil
			}
			allData = append(allData, data)
		}
	}
	return nil, allData
}

func multipartFormTest(url string, filePath string, accessToken string) (error, map[string]any) {
	file, err := os.Open(filePath)
	/*if err != nil {
		return err, nil
	}*/
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(filePath))
	io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Authorization", accessToken)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return err, nil
	}

	var data map[string]interface{}
	errors := json.Unmarshal([]byte(buf.String()), &data)
	if errors != nil {
		return errors, nil
	}
	return nil, data
}

func Test_Load_Ok(t *testing.T) {
	req := &Request{
		url:     "http://localhost:8201/writer/add-graph/TEST/ohsu-test",
		method:  "POST",
		headers: map[string]any{"Authorization": createToken(false, true, true)},
	}
	t.Run("create_load_graph", func(t *testing.T) {
		response_json, pass := TemplateRequest(req, t)
		if !pass {
			t.Error("status is not 200")
		}
		t.Log("RESPONSE JSON: ", response_json, "STATUS: ", pass)
	})

	err, responses := bulkLoad("http://localhost:8201/writer/TEST/bulk-load/ohsu-test",
		"fixtures/combio-examples-grip",
		createToken(false, true, true),
	)
	if err != nil {
		t.Error(err)
	}
	for _, resp := range responses {
		if resp["status"].(float64) != 200 {
			t.Error(resp)
		} else {
			t.Log(resp)
		}
	}
}

func Test_Json_Schema_Load_Ok(t *testing.T) {
	req := &Request{
		url:     "http://localhost:8201/writer/add-graph/JSONTEST/ohsu-test",
		method:  "POST",
		headers: map[string]any{"Authorization": createToken(false, true, true)},
	}
	t.Run("create_load_graph", func(t *testing.T) {
		response_json, pass := TemplateRequest(req, t)
		if !pass {
			t.Error("status is not 200")
		}
		t.Log("RESPONSE JSON: ", response_json, "STATUS: ", pass)
	})

	tempDir, err := os.MkdirTemp("", "jsonschema")
	if err != nil {
		t.Errorf("Error creating temp directory: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir)
	tempFilePath := fmt.Sprintf("%s/graph-fhir.json", tempDir)
	curlCmd := exec.Command("curl", "-o", tempFilePath, "https://raw.githubusercontent.com/bmeg/iceberg/refs/heads/main/schemas/graph/graph-fhir.json")
	if err := curlCmd.Run(); err != nil {
		t.Errorf("Error running curl command: %v\n", err)
		return
	}

	t.Logf("Downloaded schema to: %s\n", tempFilePath)

	err, response := multipartFormTest("http://localhost:8201/writer/JSONTEST/add-json-schema/ohsu-test",
		tempFilePath,
		createToken(false, true, true),
	)
	if err != nil {
		t.Error(err)
	}
	t.Log("RESP: ", response)
}

func Test_Json_Schema_Not_Found(t *testing.T) {
	err, response := multipartFormTest("http://localhost:8201/writer/JSONTEST/add-json-schema/ohsu-test",
		"fixtures/not_found_schema.json",
		createToken(false, true, true),
	)
	if err != nil {
		t.Error(err)
	}
	if response["status"].(float64) == 500 {
		t.Log("Expected outcome:", response["message"].(string))
	} else {
		t.Error("Status shouldnt' equal 200", response["status"].(string))
	}
}

func Test_Json_Load_Ok(t *testing.T) {
	err, response := multipartFormTest("http://localhost:8201/writer/JSONTEST/bulk-load-raw/ohsu-test",
		"fixtures/compbio-examples-fhir/DocumentReference.ndjson",
		createToken(false, true, true),
	)
	if err != nil {
		t.Error(err)
	}
	if response["status"].(float64) != 200 {
		t.Errorf("Response %f != 200\n", response["status"].(float64))
	}
}

func Test_Json_Load_Validation_Errors(t *testing.T) {
	file, err := os.Open("fixtures/compbio-examples-fhir/ValidationErrors.ndjson")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("file", filepath.Base("fixtures/compbio-examples-fhir/ValidationErrors.ndjson"))
	io.Copy(part, file)
	writer.Close()

	req := &Request{
		url:     "http://localhost:8201/writer/JSONTEST/bulk-load-raw/ohsu-test",
		method:  "POST",
		headers: map[string]any{"Authorization": createToken(false, true, true)},
	}

	request, err := http.NewRequest(req.method, req.url, body)
	if err != nil {
		t.Error("Error creating request:", err)
		return
	}
	for key, val := range req.headers {
		request.Header.Set(key, val.(string))
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		t.Error("Error sending request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Logf("server responded with status: %d", resp.StatusCode)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		t.Error("Error reading response:", err)
		return
	}

	var data map[string]interface{}
	errors := json.Unmarshal([]byte(buf.String()), &data)
	if errors != nil {
		t.Error("Error: ", errors)
		return
	}
	t.Log(data)

	if data["message"] == nil || len(data["message"].([]any)) != 2 {
		t.Error("Expected return message of length 2")
	}
	if data["status"].(float64) != 206 {
		t.Error()
	}

}

func Test_Json_Load_Invalid_Json(t *testing.T) {
	file, err := os.Open("fixtures/compbio-examples-fhir/invalid_json.ndjson")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("file", filepath.Base("fixtures/compbio-examples-fhir/ValidationErrors.ndjson"))
	io.Copy(part, file)
	writer.Close()

	req := &Request{
		url:     "http://localhost:8201/writer/JSONTEST/bulk-load-raw/ohsu-test",
		method:  "POST",
		headers: map[string]any{"Authorization": createToken(false, true, true)},
	}

	request, err := http.NewRequest(req.method, req.url, body)
	if err != nil {
		t.Error("Error creating request:", err)
		return
	}
	for key, val := range req.headers {
		request.Header.Set(key, val.(string))
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		t.Error("Error sending request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Logf("server responded with status: %d", resp.StatusCode)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		t.Error("Error reading response:", err)
		return
	}

	var data map[string]interface{}
	errors := json.Unmarshal([]byte(buf.String()), &data)
	if errors != nil {
		t.Error("Error: ", errors)
		return
	}
	t.Log(data)

	if data["status"].(float64) != 206 {
		t.Error()
	}

}

func Test_Load_Malformed_Token(t *testing.T) {
	/* Server returns a 400 given an unparsable token */
	err, responses := bulkLoad("http://localhost:8201/writer/TEST/bulk-load/ohsu-test",
		"fixtures/combio-examples-grip",
		createToken(false, true, true)[2:50],
	)
	if err != nil {
		t.Error(err)
	}
	for _, resp := range responses {
		if resp["status"].(float64) != 400 {
			t.Error(resp)
		} else {
			t.Log(resp)
		}
	}
}

func Test_Load_Expired_Token(t *testing.T) {
	/* Server returns a 401 given an expired token */
	err, responses := bulkLoad("http://localhost:8201/writer/TEST/bulk-load/ohsu-test",
		"fixtures/combio-examples-grip",
		createToken(true, true, true),
	)
	if err != nil {
		t.Error(err)
	}
	for _, resp := range responses {
		if resp["status"].(float64) != 401 {
			t.Error(resp)
		} else {
			t.Log(resp)
		}
	}
}

func Test_Load_No_Writer_Perms(t *testing.T) {
	/* Server returns a 403 given a token that respresents a user with no writer perms on the specified project */
	err, responses := bulkLoad("http://localhost:8201/writer/TEST/bulk-load/ohsu-test",
		"fixtures/combio-examples-grip",
		createToken(false, false, true),
	)
	if err != nil {
		t.Error(err)
	}
	for _, resp := range responses {
		if resp["status"].(float64) != 403 {
			t.Error(resp)
		} else {
			t.Log(resp)
		}
	}
}

func Test_Get_Vertex_No_Reader_Perms(t *testing.T) {
	/* Server returns a 403 given a token that respresents a user with no writer perms on the specified project */
	req := &Request{
		url:     "http://localhost:8201/writer/TEST/get-vertex/cec32723-9ede-5f24-ba63-63cb8c6a02cf/ohsu-test",
		method:  "GET",
		headers: map[string]any{"Authorization": createToken(false, true, false)},
	}

	response, status := TemplateRequest(req, t)
	if !(response["status"].(float64) == 403) || !status {
		t.Error(response)
	}
	t.Log(status, response)
}

func Test_Get_Vertex_Ok(t *testing.T) {
	/* given appropriate perms, edge should be retrieved */
	req := &Request{
		url:     "http://localhost:8201/writer/TEST/get-vertex/cec32723-9ede-5f24-ba63-63cb8c6a02cf/ohsu-test",
		method:  "GET",
		headers: map[string]any{"Authorization": createToken(false, false, true)},
	}
	response, status := TemplateRequest(req, t)
	if !(response["status"].(float64) == 200) || !status {
		t.Error(response)
	}
	t.Log(status, response)
}

func Test_Get_Project_Vertices_Ok(t *testing.T) {
	req := &Request{
		url:     "http://localhost:8201/writer/TEST/get-vertices/ohsu-test",
		method:  "GET",
		headers: map[string]any{"Authorization": createToken(false, false, true)},
	}

	request, err := http.NewRequest(req.method, req.url, bytes.NewBuffer(req.body))
	if err != nil {
		t.Error("Error creating request:", err)
		return
	}
	for key, val := range req.headers {
		request.Header.Set(key, val.(string))
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		t.Error("Error sending request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Logf("server responded with status: %s", resp.Status)
	}

	reader := bufio.NewReader(resp.Body)
	jum := protojson.UnmarshalOptions{DiscardUnknown: true}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Errorf("error reading response: %s", err)
		}
		v := &gripql.Vertex{}
		err = jum.Unmarshal([]byte(line), v)
		if err != nil {
			t.Error(err)
		}
		if v.Gid == "" {
			t.Error("Gid should be populated if unmarshal was successful")
		}
		mappedData := v.Data.AsMap()

		t.Logf("Received Vertex: %s\n", mappedData["resourceType"])

		if mappedData["auth_resource_path"] != "/programs/ohsu/projects/test" {
			t.Error("returned data should have resource path: /programs/ohsu/projects/test")
		}
	}
}

func Test_Delete_Edge_Ok(t *testing.T) {
	/* Regular delete should return 200 */
	req := &Request{
		url:     "http://localhost:8201/writer/TEST/del-edge/e2867c6d-db7e-5d6e-87d8-b1c293f5b47e/ohsu-test",
		method:  "DELETE",
		headers: map[string]any{"Authorization": createToken(false, true, true)},
	}
	response, status := TemplateRequest(req, t)
	if !(response["status"].(float64) == 200) || !status {
		t.Error(response)
	}
	t.Log(status, response)

}

func Test_Graphql_Query_Forbidden_Perms(t *testing.T) {
	/* Testing do a query, but with a token that doesn't have read perms */
	payload := []byte(`{
	    "query": "query PatientIdsWithSpecimenEdge { PatientIdsWithSpecimenEdge {    id  }}",
	    "variables": {
	        "limit": 1000
	    }
	}`)

	req := &Request{
		url:    "http://localhost:8201/reader/api",
		method: "POST",
		headers: map[string]any{
			"Authorization": createToken(false, false, false),
			"Content-Type":  "application/json"},
		body: payload,
	}
	response, status := TemplateRequest(req, t)
	if !(response["StatusCode"].(float64) == 403) {
		t.Error(response)
	}
	t.Log(status, response)
}
func Test_Graphql_Query_Proj(t *testing.T) {
	/* A basic test for Graphql style query with mock auth */
	payload := []byte(`{
	    "query": "query PatientIdsWithSpecimenEdge { PatientIdsWithSpecimenEdge {    id  }}",
	    "variables": {
	        "limit": 1000
	    }
	}`)

	req := &Request{
		url:    "http://localhost:8201/reader/api",
		method: "POST",
		headers: map[string]any{
			"Authorization": createToken(false, false, true),
			"Content-Type":  "application/json"},
		body: payload,
	}

	response, status := TemplateRequest(req, t)
	if !status {
		t.Error("Status returned false: ", response)
	}
	correct_response, ok := response["data"].(map[string]any)["PatientIdsWithSpecimenEdge"].([]any)
	if !ok {
		t.Error("Response not indexable for 'data' and/or 'PatientIdsWithSpecimenEdge' keys")
	}
	for _, resp := range correct_response {
		val, ok := resp.(map[string]any)["id"].(string)
		if !ok {
			t.Error("Response not indexable on 'id': ", resp)
		}
		t.Log("Return VAL: ", val)
	}
}

func Test_Basic_Graphql(t *testing.T) {
	/* A basic test for Graphql style query with mock auth */
	payload := []byte(`{
	    "query": "query { specimen {    id  }}",
	    "variables": {
	        "limit": 10
	    }
	}`)

	req := &Request{
		url:    "http://localhost:8201/graphql/query",
		method: "POST",
		headers: map[string]any{
			"Authorization": createToken(false, false, true),
			"Content-Type":  "application/json"},
		body: payload,
	}

	response, status := TemplateRequest(req, t)
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

func Test_Graphql_Bad_Filter_Logical_Operator(t *testing.T) {
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

	req := &Request{
		url:    "http://localhost:8201/graphql/query",
		method: "POST",
		headers: map[string]any{
			"Authorization": createToken(false, false, true),
			"Content-Type":  "application/json"},
		body: payload,
	}

	response, status := TemplateRequest(req, t)
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

func Test_Graphql_Filter_Ok(t *testing.T) {
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

	req := &Request{
		url:    "http://localhost:8201/graphql/query",
		method: "POST",
		headers: map[string]any{
			"Authorization": createToken(false, false, true),
			"Content-Type":  "application/json"},
		body: payload,
	}

	response, status := TemplateRequest(req, t)
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

func Test_Graphql_Edge_Traversal_Ok(t *testing.T) {
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
	req := &Request{
		url:    "http://localhost:8201/graphql/query",
		method: "POST",
		headers: map[string]any{
			"Authorization": createToken(false, false, true),
			"Content-Type":  "application/json"},
		body: payload,
	}

	response, status := TemplateRequest(req, t)
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

func Test_Nested_Edge_Traversal_Ok(t *testing.T) {
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
	req := &Request{
		url:    "http://localhost:8201/graphql/query",
		method: "POST",
		headers: map[string]any{
			"Authorization": createToken(false, false, true),
			"Content-Type":  "application/json"},
		body: payload,
	}

	response, status := TemplateRequest(req, t)
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

func Test_Nested_Edge_Traversal_Filter_Ok(t *testing.T) {
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
	req := &Request{
		url:    "http://localhost:8201/graphql/query",
		method: "POST",
		headers: map[string]any{
			"Authorization": createToken(false, false, true),
			"Content-Type":  "application/json"},
		body: []byte(payload),
	}

	response, status := TemplateRequest(req, t)
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

func Test_Delete_Proj(t *testing.T) {
	/* Delete Everything from test graph project ohsu-test. Should return 200 */
	req := &Request{
		url:     "http://localhost:8201/writer/TEST/proj-delete/ohsu-test",
		method:  "DELETE",
		headers: map[string]any{"Authorization": createToken(false, true, true)},
	}
	response, status := TemplateRequest(req, t)
	if !(response["status"].(float64) == 200) {
		t.Error(response)
	}
	t.Log(status, response)

	// Look for the vertex that was retrieved earlier. It should be gone
	req = &Request{
		url:     "http://localhost:8201/writer/TEST/get-vertex/cec32723-9ede-5f24-ba63-63cb8c6a02cf/ohsu-test",
		method:  "GET",
		headers: map[string]any{"Authorization": createToken(false, false, true)},
	}
	response, status = TemplateRequest(req, t)
	if !(response["status"].(float64) == 404) {
		t.Error(response)
	}
	t.Log("RESPONSE: ", response)
}
