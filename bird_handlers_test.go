package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetBirdsHandler(t *testing.T) {
	// initialize mock store
	mockStore := InitMockStore()

	// define mock data
	mockStore.On("GetBirds").Return([]*Bird{{"sparrow", "A small harmless bird"}}, nil).Once()

	// setup new get request and test recorder
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(getBirdHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// setup expected results and get body from response
	expected := Bird{"sparrow", "A small harmless bird"}
	b := []Bird{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

	// assert expectations defined in the On method above.
	mockStore.AssertExpectations(t)
}
func TestCreateBirdsHandler(t *testing.T) {
	// initialize mock store
	mockStore := InitMockStore()

	// define mock data
	mockStore.On("CreateBird", &Bird{"eagle", "A bird of prey"}).Return(nil)

	form := newCreateBirdForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(createBirdHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// assert expectations defined in the On method above.
	mockStore.AssertExpectations(t)
}

func newCreateBirdForm() *url.Values {
	form := url.Values{}
	form.Set("species", "eagle")
	form.Set("description", "A bird of prey")
	return &form
}
