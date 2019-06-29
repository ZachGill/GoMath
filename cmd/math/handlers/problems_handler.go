package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Problems responds to GET requests with an array of the problems sent to the
// API so far
type Problems struct {
	ProblemsAccessor *Accessor
}

// ProblemsResponse is a JSON object containing
type ProblemsResponse struct {
	Problems []MathResponse `json:"problems"`
}

// ServeHTTP returns a list of all problems and their results
func (handler *Problems) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var (
		httpStatusCode = http.StatusOK
		problems       []MathResponse
		problemsMap    map[string]MathResponse
		message        ProblemsResponse
		jsonBody       []byte
	)

	problemsMap = handler.ProblemsAccessor.GetProblems()

	for _, problem := range problemsMap {
		problems = append(problems, problem)
	}

	message = ProblemsResponse{
		Problems: problems,
	}

	writer.WriteHeader(httpStatusCode)
	jsonBody, _ = json.Marshal(message)
	writer.Write(jsonBody)
}

// Problem returns a single problem response by problem ID
type Problem struct {
	ProblemsAccessor *Accessor
}

// ServeHTTP receives a problem ID, then responds with the problem results
// attached to the ID.
func (handler *Problem) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var (
		httpStatusCode = http.StatusOK
		muxVars        map[string]string
		found          bool
		problemID      string
		problem        MathResponse
		problemsMap    map[string]MathResponse
		jsonBody       []byte
	)

	muxVars = mux.Vars(request)
	problemID = muxVars["problem_id"]

	problemsMap = handler.ProblemsAccessor.GetProblems()

	if problem, found = problemsMap[problemID]; !found {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("404 - Problem not found"))
		return
	}

	writer.WriteHeader(httpStatusCode)
	jsonBody, _ = json.Marshal(problem)
	writer.Write(jsonBody)
}
