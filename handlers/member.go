package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-mongo-rest-api/common"
	"go-mongo-rest-api/helpers"
	"go-mongo-rest-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetMember(w http.ResponseWriter, r *http.Request) {
	var response types.Response
	vars := mux.Vars(r)
	email := vars["email"]
	if email == "" {
		response.Status = 400
		response.ErrorMessage = "Bad Request"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	result, err := helpers.FindOne(os.Getenv("COLLECTION_NAME"), bson.M{"email": email})
	if err != nil {
		response.Status = 400
		response.ErrorMessage = "Member not found"
		json.NewEncoder(w).Encode(response)
		return
	}

	var member types.Member
	err = result.Decode(&member)

	response.Status = 200
	response.Message = "Here is the response"
	response.Data = member
	json.NewEncoder(w).Encode(response)
}

func DeleteMember(w http.ResponseWriter, r *http.Request) {
	var response types.Response
	vars := mux.Vars(r)
	email := vars["email"]
	deletedCount, status := helpers.DeleteOne(os.Getenv("COLLECTION_NAME"), bson.M{"email": email})
	if !status {
		response.Status = 400
		response.ErrorMessage = "Member could not be deleted"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Status = 200
	response.Message = strconv.FormatInt(deletedCount, 10) + " object deleted successfully"
	json.NewEncoder(w).Encode(response)
}

func MemberHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetMember(w, r)
	case http.MethodDelete:
		DeleteMember(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}