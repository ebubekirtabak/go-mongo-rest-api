package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-mongo-rest-api/common"
	"go-mongo-rest-api/handlers"
	"go-mongo-rest-api/types"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetMemberBadRequest(t *testing.T) {
	common.Init()
	req, err := http.NewRequest("GET", "/member", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/member", handlers.MemberHandler)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := `{"status":400,"message":"","error_message":"Bad Request","data":null}`
	bodyString := strings.Trim(rr.Body.String(), "\n")
	if bodyString != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			bodyString, expected)
	}
}

func TestGetMemberHandler(t *testing.T) {
	common.Init()
	req, err := http.NewRequest("GET", "/member/alive.wance@mail.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/member/{email}", handlers.MemberHandler)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	member := types.Member{}
	var response types.Response
	json.Unmarshal(rr.Body.Bytes(), &response)
	jsonString, _ := json.Marshal(response.Data)
	json.Unmarshal(jsonString, &member)

	validationErrors := handlers.ValidateMember(member)
	if  len(validationErrors) > 0 {
		t.Errorf("response.data is not valid Member, Error: %v", validationErrors)
	}

}

func TestCreateMemberHandler(t *testing.T) {
	common.Init()
	email := "test.tester" + fmt.Sprint(rand.Int31()) + "_" + fmt.Sprint(rand.Int31()) + "@mail.com"
	member := map[string]interface{}{
		"type": "employee",
		"skills": []string{
			"React",
			"ReactNative",
			"graphql",
			"Javascript",
			"Html",
			"scss",
			"css",
			"rest",
		},
		"member_name": "Test Test",
		"title": "Test Developer",
		"email": email,
	}

	fmt.Println(email)
	body, _ := json.Marshal(member)
	req, err := http.NewRequest("POST", "/member", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/member", handlers.MemberHandler)
	router.HandleFunc("/member/{email}", handlers.MemberHandler)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("POST handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestDeleteMemberHandler(t *testing.T) {
	common.Init()
	email := "test.tester" + fmt.Sprint(rand.Int31()) + "_" + fmt.Sprint(rand.Int31()) + "@mail.com"
	req, err := http.NewRequest("DELETE", "/member/" + email, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/member/{email}", handlers.MemberHandler)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
