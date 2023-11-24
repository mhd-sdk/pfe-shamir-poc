package servicesManager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pfe-manager/config"
	"github.com/pfe-manager/pkg/models"
	"github.com/pfe-manager/pkg/shamir"
	"github.com/pfe-manager/pkg/statusTableUI"
	"github.com/rodaine/table"
)

var Services []models.Service

func UpdateServicesStatus() {
	type pingStatusResponse struct {
		Status string `json:"status"`
		Name   string `json:"name"`
	}

	for i := range Services {
		res, err := http.Get(Services[i].Host + "/status")
		if err != nil || res.StatusCode != http.StatusOK {
			Services[i].Status = models.ServiceDown
			Services[i].Name = "Unknown"
			continue
		}
		response := pingStatusResponse{}
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			Services[i].Status = models.ServiceDown
			Services[i].Name = "Unknown"
			continue
		}
		switch response.Status {
		case models.ServiceUp.String():
			Services[i].Status = models.ServiceUp
			Services[i].Name = response.Name
		case models.ServiceDown.String():
			Services[i].Status = models.ServiceDown
			Services[i].Name = "Unknown"
		}

	}
	if config.GetConfig().Mode == "prod" {
		statusTableUI.UpdateTable(Services)
	}

}

func Init() {
	Services = config.GetServices()
	UpdateServicesStatus()
	table.DefaultHeaderFormatter = func(format string, vals ...interface{}) string {
		return strings.ToUpper(fmt.Sprintf(format, vals...))
	}
}

func checkServicesAreAllUp() bool {
	for _, service := range Services {
		if service.Status == models.ServiceDown {
			return false
		}
	}
	return true
}

func checkRetrieveMinimumRequirements() bool {
	// check if enough services are up to retrieve the secret
	upServices := 0
	for _, service := range Services {
		if service.Status == models.ServiceUp {
			upServices++
		}
	}
	if upServices < config.GetConfig().ShamirThreshold {
		return false
	}
	return true

}

type SaveRequest struct {
	Part shamir.ShamirPart `json:"part"`
}

// dispatch key parts to services, return error if any service is down
func Dispatch(parts []shamir.ShamirPart) error {
	if !checkServicesAreAllUp() {
		return fmt.Errorf("some services are down")
	}
	for idx, part := range parts {
		// http.Get(Services[part.Key].Host + "/save/")
		// get request with part in body

		body := SaveRequest{Part: part}
		postBody, _ := json.Marshal(body)
		postBodyBuffer := bytes.NewBuffer(postBody)
		_, err := http.Post(Services[idx].Host+"/save", "application/json", postBodyBuffer)
		if err != nil {
			return err
		}

	}
	return nil
}

func Retrieve() (secret []byte, err error) {
	if !checkRetrieveMinimumRequirements() {
		return nil, fmt.Errorf("not enough services are up to retrieve the secret")
	}
	parts := make([]shamir.ShamirPart, 0)
	for _, service := range Services {
		if service.Status == models.ServiceUp {
			res, err := http.Get(service.Host + "/retrieve")
			if err != nil {
				return nil, err
			}
			response := SaveRequest{}
			err = json.NewDecoder(res.Body).Decode(&response)
			if err != nil {
				return nil, err
			}
			parts = append(parts, response.Part)
		}
	}
	secret, err = shamir.ReconstructSeed(parts)
	if err != nil {
		return nil, err
	}
	return secret, nil
}
