package integration_test

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
	"testing"

	"github.com/bmeg/grip-graphql/tests/integration"
	"github.com/bmeg/grip/gripql"
)

func Test_Writer_Load_Ok(t *testing.T) {
	req := &integration.Request{
		Url:     "http://localhost:8201/writer/add-graph/TEST/ohsu-test",
		Method:  "POST",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, true, true)},
	}
	// Inlined subtest logic
	response_json, pass := integration.TemplateRequest(req, t)
	if !pass {
		t.Error("status is not 200")
	}
	t.Log("RESPONSE JSON: ", response_json, "STATUS: ", pass)

	err, responses := integration.BulkLoad("http://localhost:8201/writer/TEST/bulk-load/ohsu-test",
		"../fixtures/combio-examples-grip",
		integration.CreateToken(false, true, true),
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

func Test_Writer_Json_Schema_Load_Ok(t *testing.T) {
	req := &integration.Request{
		Url:     "http://localhost:8201/writer/add-graph/JSONTEST/ohsu-test",
		Method:  "POST",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, true, true)},
	}
	// Inlined subtest logic from "create_load_graph"
	response_json, pass := integration.TemplateRequest(req, t)
	if !pass {
		t.Error("status is not 200")
	}
	t.Log("RESPONSE JSON: ", response_json, "STATUS: ", pass)

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

	err, response := integration.MultipartFormTest("http://localhost:8201/writer/JSONTEST/add-json-schema/ohsu-test",
		tempFilePath,
		integration.CreateToken(false, true, true),
	)
	if err != nil {
		t.Error(err)
	}
	t.Log("RESP: ", response)
}

func Test_Writer_Json_Load_Ok(t *testing.T) {
	err, response := integration.MultipartFormTest("http://localhost:8201/writer/JSONTEST/bulk-load-raw/ohsu-test",
		"../fixtures/compbio-examples-fhir/DocumentReference.ndjson",
		integration.CreateToken(false, true, true),
	)
	if err != nil {
		t.Error(err)
	}
	if response["status"].(float64) != 200 {
		t.Errorf("Response %f != 200\n", response["status"].(float64))
	}
}

func Test_Writer_Json_Schema_Not_Found(t *testing.T) {
	err, response := integration.MultipartFormTest("http://localhost:8201/writer/JSONTEST/add-json-schema/ohsu-test",
		"../fixtures/not_found_schema.json",
		integration.CreateToken(false, true, true),
	)
	if err != nil {
		t.Error(err)
	}
	if response["status"].(float64) == 500 {
		t.Log("Expected outcome:", response["message"].(string))
	} else {
		t.Error("Status shouldn't equal 200", response["status"].(string))
	}
}

func Test_Writer_Json_Load_Validation_Errors(t *testing.T) {
	file, err := os.Open("../fixtures/compbio-examples-fhir/ValidationErrors.ndjson")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("file", filepath.Base("../fixtures/compbio-examples-fhir/ValidationErrors.ndjson"))
	io.Copy(part, file)
	writer.Close()

	req := &integration.Request{
		Url:     "http://localhost:8201/writer/JSONTEST/bulk-load-raw/ohsu-test",
		Method:  "POST",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, true, true)},
	}

	request, err := http.NewRequest(req.Method, req.Url, body)
	if err != nil {
		t.Error("Error creating request:", err)
		return
	}
	for key, val := range req.Headers {
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

func Test_Writer_Json_Load_Invalid_Json(t *testing.T) {
	file, err := os.Open("../fixtures/compbio-examples-fhir/invalid_json.ndjson")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("file", filepath.Base("../fixtures/compbio-examples-fhir/ValidationErrors.ndjson"))
	io.Copy(part, file)
	writer.Close()

	req := &integration.Request{
		Url:     "http://localhost:8201/writer/JSONTEST/bulk-load-raw/ohsu-test",
		Method:  "POST",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, true, true)},
	}

	request, err := http.NewRequest(req.Method, req.Url, body)
	if err != nil {
		t.Error("Error creating request:", err)
		return
	}
	for key, val := range req.Headers {
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

func Test_Writer_Load_Malformed_Token(t *testing.T) {
	/* Server returns a 400 given an unparsable token */
	err, responses := integration.BulkLoad("http://localhost:8201/writer/TEST/bulk-load/ohsu-test",
		"../fixtures/combio-examples-grip",
		integration.CreateToken(false, true, true)[2:50],
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

func Test_Writer_Load_Expired_Token(t *testing.T) {
	/* Server returns a 401 given an expired token */
	err, responses := integration.BulkLoad("http://localhost:8201/writer/TEST/bulk-load/ohsu-test",
		"../fixtures/combio-examples-grip",
		integration.CreateToken(true, true, true),
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

func Test_Writer_Load_No_Writer_Perms(t *testing.T) {
	/* Server returns a 403 given a token that represents a user with no writer perms on the specified project */
	err, responses := integration.BulkLoad("http://localhost:8201/writer/TEST/bulk-load/ohsu-test",
		"../fixtures/combio-examples-grip",
		integration.CreateToken(false, false, true),
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

func Test_Writer_Get_Vertex_No_Reader_Perms(t *testing.T) {
	/* Server returns a 403 given a token that represents a user with no reader perms on the specified project */
	req := &integration.Request{
		Url:     "http://localhost:8201/writer/TEST/get-vertex/cec32723-9ede-5f24-ba63-63cb8c6a02cf/ohsu-test",
		Method:  "GET",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, true, false)},
	}

	response, status := integration.TemplateRequest(req, t)
	if !(response["status"].(float64) == 403) || !status {
		t.Error(response)
	}
	t.Log(status, response)
}

func Test_Writer_Get_Vertex_Ok(t *testing.T) {
	/* Given appropriate perms, vertex should be retrieved */

	req := &integration.Request{
		Url:     "http://localhost:8201/writer/TEST/get-vertex/cec32723-9ede-5f24-ba63-63cb8c6a02cf/ohsu-test",
		Method:  "GET",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, false, true)},
	}
	response, status := integration.TemplateRequest(req, t)
	if !(response["status"].(float64) == 200) || !status {
		t.Error(response)
	}
	t.Log(status, response)
}

func Test_Writer_Get_Project_Vertices_Ok(t *testing.T) {
	req := &integration.Request{
		Url:     "http://localhost:8201/writer/TEST/get-vertices/ohsu-test",
		Method:  "GET",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, false, true)},
	}

	request, err := http.NewRequest(req.Method, req.Url, bytes.NewBuffer(req.Body))
	if err != nil {
		t.Error("Error creating request:", err)
		return
	}
	for key, val := range req.Headers {
		request.Header.Set(key, val.(string))
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		t.Error("Error sending request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Logf("Server responded with status: %s", resp.Status)
	}

	reader := bufio.NewReader(resp.Body)
	jum := gripql.NewFlattenMarshaler()
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Errorf("Error reading response: %s", err)
		}
		v := &gripql.Vertex{}
		err = jum.Unmarshal([]byte(line), v)
		if err != nil {
			t.Error(err)
		}
		if v.Id == "" {
			t.Error("Gid should be populated if unmarshal was successful")
		}
		mappedData := v.Data.AsMap()

		t.Logf("Received Vertex: %s\n", mappedData["resourceType"])

		if mappedData["auth_resource_path"] != "/programs/ohsu/projects/test" {
			t.Error("Returned data should have resource path: /programs/ohsu/projects/test")
		}
	}
}

func Test_Writer_Delete_Edge_Ok(t *testing.T) {
	/* Regular delete should return 200 */
	req := &integration.Request{
		Url:     "http://localhost:8201/writer/TEST/del-edge/e2867c6d-db7e-5d6e-87d8-b1c293f5b47e/ohsu-test",
		Method:  "DELETE",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, true, true)},
	}
	response, status := integration.TemplateRequest(req, t)
	if !(response["status"].(float64) == 200) || !status {
		t.Error(response)
	}
	t.Log(status, response)
}

func Test_Writer_Delete_Proj(t *testing.T) {
	/* Delete everything from test graph project ohsu-test. Should return 200 */
	req := &integration.Request{
		Url:     "http://localhost:8201/writer/TEST/proj-delete/ohsu-test",
		Method:  "DELETE",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, true, true)},
	}
	response, status := integration.TemplateRequest(req, t)
	if !(response["status"].(float64) == 200) {
		t.Error(response)
	}
	t.Log(status, response)

	// Look for the vertex that was retrieved earlier. It should be gone
	req = &integration.Request{
		Url:     "http://localhost:8201/writer/TEST/get-vertex/cec32723-9ede-5f24-ba63-63cb8c6a02cf/ohsu-test",
		Method:  "GET",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, false, true)},
	}
	response, status = integration.TemplateRequest(req, t)
	if !(response["status"].(float64) == 404) {
		t.Error(response)
	}
	t.Log("RESPONSE: ", response)
}
