package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	middleware "github.com/ZachGill/transaction-mw"
)

// Multiply performs multiplication on two input floats and returns their product
type Multiply struct {
	ProblemsAccessor *Accessor
}

// ServeHTTP multiplies two input floats and returns their product them as a
// JSON body in the standard MathResponse format
func (handler *Multiply) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	var (
		err error

		httpStatusCode = http.StatusOK
		inputs         []string

		input1  float64
		input2  float64
		product float64

		response MathResponse

		jsonBody []byte
	)

	inputs = request.URL.Query()["input"]

	if inputLength := len(inputs); inputLength < 2 || inputLength > 2 {
		httpStatusCode = http.StatusBadRequest
		writer.WriteHeader(httpStatusCode)
		writer.Write([]byte("400 - Request does not contain 2 numbers"))
		return
	}

	if input1, err = strconv.ParseFloat(inputs[0], 64); err != nil {
		httpStatusCode = http.StatusBadRequest
		writer.WriteHeader(httpStatusCode)
		writer.Write([]byte("400 - Request contains one or more non-number inputs"))
		return
	}

	if input2, err = strconv.ParseFloat(inputs[1], 64); err != nil {
		httpStatusCode = http.StatusBadRequest
		writer.WriteHeader(httpStatusCode)
		writer.Write([]byte("400 - Request contains one or more non-number inputs"))
		return
	}

	product = input1 * input2

	response = MathResponse{
		ProblemID: request.Header.Get(middleware.Key),
		Operation: "multiplication",
		Input1:    input1,
		Input2:    input2,
		Answer:    product,
	}

	if jsonBody, err = json.Marshal(response); err != nil {
		httpStatusCode = http.StatusInternalServerError
		writer.WriteHeader(httpStatusCode)
		writer.Write([]byte("500 - Error marshalling JSON response"))
		return
	}

	handler.ProblemsAccessor.AddProblem(response)

	writer.WriteHeader(httpStatusCode)
	writer.Write(jsonBody)
}
