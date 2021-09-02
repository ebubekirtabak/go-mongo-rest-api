package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-mongo-rest-api/helpers"
	"go-mongo-rest-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"os"
	"strings"
)

func getQuery(vars map[string]string) bson.M {
	query := bson.M{}
	if len(vars) == 0 {
		return query
	}

	skills := vars["skills"]
	title := vars["title"]

	var andQueries []bson.M

	if skills != "" {
		andQueries = append(andQueries, bson.M {
			"skills": bson.M{ "$in": strings.Split(skills, ",")},
		})
	}

	if title != "" {
		andQueries = append(andQueries, bson.M {
			"title": bson.M{ "$eq": title },
		})
	}

	query["$and"] = andQueries

	return query
}

func findMember(w http.ResponseWriter, r *http.Request) {
	var response types.Response
	var members []types.Member
	vars := mux.Vars(r)

	query := getQuery(vars)

	cursor, err := helpers.FindAll(os.Getenv("COLLECTION_NAME"), query)
	if err != nil {
		fmt.Println(err)
	}

	for cursor.Next(context.TODO()) {
		var member types.Member
		err := cursor.Decode(&member)
		if err != nil {
			fmt.Println(err)
		}

		members = append(members, member)
	}

	response.Status = 200
	response.Message = "Here is the response"
	response.Data = members
	json.NewEncoder(w).Encode(response)
}

func FindMemberHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		findMember(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
