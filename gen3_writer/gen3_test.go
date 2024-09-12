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

func Test_Load_Ok(t *testing.T) {
	req := &Request{
		url:     "http://localhost:8201/graphql/add-graph/TEST/ohsu-test",
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

	err, responses := bulkLoad("http://localhost:8201/graphql/TEST/bulk-load/ohsu-test",
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

func Test_Load_Malformed_Token(t *testing.T) {
	/* Server returns a 400 given an unparsable token */
	err, responses := bulkLoad("http://localhost:8201/graphql/TEST/bulk-load/ohsu-test",
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
	err, responses := bulkLoad("http://localhost:8201/graphql/TEST/bulk-load/ohsu-test",
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
	err, responses := bulkLoad("http://localhost:8201/graphql/TEST/bulk-load/ohsu-test",
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
		url:     "http://localhost:8201/graphql/TEST/get-vertex/cec32723-9ede-5f24-ba63-63cb8c6a02cf/ohsu-test",
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
		url:     "http://localhost:8201/graphql/TEST/get-vertex/cec32723-9ede-5f24-ba63-63cb8c6a02cf/ohsu-test",
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
		url:     "http://localhost:8201/graphql/TEST/get-vertices/ohsu-test",
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
	t.Log("RESP: ")
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
		url:     "http://localhost:8201/graphql/TEST/del-edge/e2867c6d-db7e-5d6e-87d8-b1c293f5b47e/ohsu-test",
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

func Test_Delete_Proj(t *testing.T) {
	/* Delete Everything from test graph project ohsu-test. Should return 200 */
	req := &Request{
		url:     "http://localhost:8201/graphql/TEST/proj-delete/ohsu-test",
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
		url:     "http://localhost:8201/graphql/TEST/get-vertex/cec32723-9ede-5f24-ba63-63cb8c6a02cf/ohsu-test",
		method:  "GET",
		headers: map[string]any{"Authorization": createToken(false, false, true)},
	}
	response, status = TemplateRequest(req, t)
	if !(response["status"].(float64) == 404) {
		t.Error(response)
	}
	t.Log("RESPONSE: ", response)
}
