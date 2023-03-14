package response

import (
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/status"
)

type GenericResponse struct {
	Success  bool        `json:"success"`
	Messages []string    `json:"messages"`
	Data     interface{} `json:"data"`
	Meta     Metadata    `json:"meta,omitempty"`
}

// ─────────────────────────────────────────────────────────────────────────────
type Metadata struct {
	//this is used for the list request if a pagination is needed
	NextPageToken string `json:"next_page_token,omitempty"`
}

// ────────────────────────────────────────────────────────────────────────────────
func (r *GenericResponse) HandleError(err error) bool {
	if err == nil {
		return false
	}
	e, _ := status.FromError(err)

	r.Messages = append(r.Messages, e.Message())
	r.Success = false
	// error happened
	return true
}

// ────────────────────────────────────────────────────────────────────────────────
func (r *GenericResponse) ResponseJson(w http.ResponseWriter, status int) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		return
	}
}
