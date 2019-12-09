package verificationsvc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gemsorg/dispute/pkg/apierror"
)

type Validations struct {
	IDs []uint64 `json:"ids"`
}

func ValidateResponse(responseID, verifierID uint64, authToken string) error {
	validation := Validations{IDs: []uint64{responseID}}
	maxRetries := 3
	tries := 0
	url := fmt.Sprintf("admin/responses/verify")

	var err error
	var data []byte
	success := false

	for !success && tries < maxRetries {
		data, err = json.Marshal(validation)
		body := bytes.NewReader(data)
		if err != nil {
			tries = tries + 1
		} else {
			_, err = serviceRequest("POST", url, verifierID, body, authToken)
			if err != nil {
				tries = tries + 1
			} else {
				success = true
			}
		}
	}

	if !success {
		return apierror.New(500, "Couldn't validate response in backend", err)
	}

	return nil
}

func serviceRequest(action string, route string, userID uint64, data io.Reader, authToken string) ([]byte, error) {
	client := &http.Client{}
	serviceURL := fmt.Sprintf("%s/%s", os.Getenv("BACKEND_SVC_URL"), route)
	fmt.Println("D", data)
	req, err := http.NewRequest(action, serviceURL, data)
	if err != nil {
		return nil, errorResponse(err)
	}
	// req.AddCookie(&http.Cookie{Name: "JWT", Value: authToken})

	// req.Header.Add("Authorization", "Bearer "+authToken)
	r, err := client.Do(req)
	if err != nil {
		return nil, errorResponse(err)
	}

	if r.StatusCode != 200 {
		body, _ := ioutil.ReadAll(r.Body)
		return nil, errorResponse(errors.New(string(body)))
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errorResponse(err)
	}

	return body, nil
}

func errorResponse(err error) *apierror.APIError {
	return apierror.New(500, err.Error(), err)
}
