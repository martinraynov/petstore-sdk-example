package handler

import (
	"encoding/json"
	"net/http"
)

// IDParam is used to identify a user
//
// swagger:parameters updateUser deleteUser logoutUser resetPassword
type IDParam struct {
	// The ID of a user
	//
	// in: path
	// required: true
	ID string `json:"id"`
}

// Status is an httpHandler for route GET /status
func Status(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /status Status status
	//
	// Status of the application.
	//
	// This will show a result if the application is up.
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: StatusResponse
	//       400: JsonError

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("OK"); err != nil {
		panic(err)
	}

	// return
}
