package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-mongo-rest-api/common"
	"go-mongo-rest-api/helpers"
	"go-mongo-rest-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/go-playground/validator.v8"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func ValidateMember(member types.Member) validator.ValidationErrors {
	var validate *validator.Validate
	config := &validator.Config{TagName: "validate"}

	validate = validator.New(config)
	err := validate.Struct(member)

	if member.Type == "contractor" {
		err = validate.Field(&member.DurationOfContract, "required,numeric")
	}

	if err != nil {
		return err.(validator.ValidationErrors)
	}

	return validator.ValidationErrors{}
}

func createMember(member types.Member) (bool, string) {
	if helpers.IsExistsObject(os.Getenv("COLLECTION_NAME"), bson.M {"email" : member.Email}) {
		return false, "User already exists"
	}

	uuidString, _ := common.GetUuid()
	member.UUID = uuidString
	member.CreatedTime = time.Now().UTC().Format("2006-01-02 03:04:05")
	bsonDoc, err := common.ConvertToBsonM(member)
	if err != nil {
		return false, err.Error()
	}

	return helpers.InsertOne(os.Getenv("COLLECTION_NAME"), bsonDoc), uuidString
}

func updateMember(member types.Member) (bool, string) {
	uuidString, _ := common.GetUuid()
	member.UUID = uuidString
	member.CreatedTime = time.Now().UTC().Format("2006-01-02 03:04:05")
	bsonDoc, err := common.ConvertToBsonM(member)
	if err != nil {
		return false, err.Error()
	}

	_, status := helpers.UpdateOne(os.Getenv("COLLECTION_NAME"), bson.M {"email" : member.Email}, bson.M{ "$set": bsonDoc })

	return status, ""
}

func UpdateMemberHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var member types.Member
	var response types.Response
	response.Status = 400
	json.Unmarshal(reqBody, &member)
	validationErrors := ValidateMember(member)
	if  len(validationErrors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "ValidationErrors: "
		for err := range validationErrors {
			response.ErrorMessage += err
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if !helpers.IsExistsObject(os.Getenv("COLLECTION_NAME"), bson.M {"email" : member.Email}) {
		response.Status = 404
		response.Message = "Member does not exist"
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	status, err := updateMember(member)

	if status {
		response.Status = 201
		response.Message = "Member updated successfully"
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = err
	}

	json.NewEncoder(w).Encode(response)
}

func CreateMemberHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var member types.Member
	var response types.Response
	response.Status = 400
	json.Unmarshal(reqBody, &member)
	validationErrors := ValidateMember(member)
	if  len(validationErrors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "ValidationErrors: "
		for err := range validationErrors {
			response.ErrorMessage += err
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	status, err := createMember(member)

	if status {
		response.Status = 201
		response.Message = "Member created successfully"
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = err
	}

	json.NewEncoder(w).Encode(response)
}

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
	case http.MethodPost:
		CreateMemberHandler(w, r)
	case http.MethodPut:
		UpdateMemberHandler(w, r)
	case http.MethodDelete:
		DeleteMember(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}