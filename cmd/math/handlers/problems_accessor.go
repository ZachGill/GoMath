package handlers

import "sync"

// Accessor contains all MathResponses created by the API and allows multiple
// functions to safely access and modify them.
type Accessor struct {
	Problems map[string]MathResponse
	Mutex    *sync.RWMutex
}

// GetProblems returns the Accessor's slice of MathResponses
func (accessor *Accessor) GetProblems() map[string]MathResponse {
	accessor.Mutex.RLock()
	defer accessor.Mutex.RUnlock()
	return accessor.Problems
}

// AddProblem adds a given problem to the Accessor's slice of MathResponses
func (accessor *Accessor) AddProblem(problem MathResponse) {
	accessor.Mutex.Lock()
	defer accessor.Mutex.Unlock()
	accessor.Problems[problem.ProblemID] = problem
}
