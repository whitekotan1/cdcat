package api

import (
	"cdcat/services"
	"cdcat/types"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type R2Client struct {
	CloudflareCfg *s3.Client
}

func (client *R2Client) HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only post", http.StatusMethodNotAllowed)
		return
	}

	var request types.Request

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	/*
		var userProject types.UserProject = services.CreateUserProject(request)
		services.CloneUserProject(userProject)

		services.BuildUserProject(userProject)
	*/

	var userProject types.UserProject = services.BuildProjectPipeline(request)

	services.DeleteUserProject(userProject.DistPath)

	services.DeployPipeline(userProject.DistPath, "cdcat", strconv.Itoa(userProject.ID), client.CloudflareCfg)

	response := types.Response{

		Status:  "cat",
		Message: "cat received your repos" + request.RepoUrl,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
