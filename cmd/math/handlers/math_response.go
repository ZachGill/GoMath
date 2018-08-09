package handlers

// MathResponse defines a standard format for 2-input math problems
type MathResponse struct {
	ProblemID string  `json:"problem_id"`
	Operation string  `json:"operation"`
	Input1    float64 `json:"input_1"`
	Input2    float64 `json:"input_2"`
	Answer    float64 `json:"answer"`
}
