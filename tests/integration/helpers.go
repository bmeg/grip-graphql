package integration

import (
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

	"github.com/bmeg/grip/log"
	"github.com/golang-jwt/jwt/v5"
)

type Request struct {
	Url     string
	Method  string
	Headers map[string]any
	Body    []byte
}

func CreateToken(expired bool, writer bool, reader bool) string {
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

	req, err := http.NewRequest(request.Method, request.Url, bytes.NewBuffer(request.Body))
	if err != nil {
		t.Error("Error creating request:", err)
		return
	}
	for key, val := range request.Headers {
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

	var data map[string]any
	errors := json.Unmarshal([]byte(buf.String()), &data)
	if errors != nil {
		t.Error("Error: ", errors)
		return nil, false
	}
	return data, true
}

func BulkLoad(url, directoryPath, accessToken string) (error, []map[string]any) {
	files, _ := os.ReadDir(directoryPath)
	var allData []map[string]any
	for _, file := range files {
		if !file.IsDir() && (filepath.Ext(file.Name()) == ".json" || filepath.Ext(file.Name()) == ".gz" || filepath.Ext(file.Name()) == ".ndjson") {
			filePath := filepath.Join(directoryPath, file.Name())
			graphComponent := "vertex"
			if strings.Contains(filePath, "edge") {
				graphComponent = "edge"
			}

			file, fileErr := os.Open(filePath)
			defer file.Close()
			if fileErr != nil {
				log.Infoln("ERR: ", fileErr)
				continue
			}

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

			var data map[string]any
			errors := json.Unmarshal([]byte(buf.String()), &data)

			if errors != nil {
				return errors, nil
			}
			allData = append(allData, data)
		}
	}
	return nil, allData
}

func MultipartFormTest(url string, filePath string, accessToken string) (error, map[string]any) {
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

	var data map[string]any
	errors := json.Unmarshal([]byte(buf.String()), &data)
	if errors != nil {
		return errors, nil
	}
	return nil, data
}
