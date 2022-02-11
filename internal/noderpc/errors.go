package noderpc

import (
	"errors"
	"fmt"
	"strings"
)

// NodeUnavailiableError -
type NodeUnavailiableError struct {
	Node string
	Code int
}

// NewNodeUnavailiableError -
func NewNodeUnavailiableError(node string, code int) NodeUnavailiableError {
	return NodeUnavailiableError{
		Node: node,
		Code: code,
	}
}

// Error -
func (e NodeUnavailiableError) Error() string {
	return fmt.Sprintf("%s is unavailiable: %d", e.Node, e.Code)
}

// MaxRetryExceededError -
type MaxRetryExceededError struct {
	Node string
}

// NewMaxRetryExceededError -
func NewMaxRetryExceededError(node string) MaxRetryExceededError {
	return MaxRetryExceededError{
		Node: node,
	}
}

// Error -
func (e MaxRetryExceededError) Error() string {
	return fmt.Sprintf("%s: max HTTP request retry exceeded", e.Node)
}

// IsNodeUnavailiableError -
func IsNodeUnavailiableError(err error) bool {
	if _, ok := err.(MaxRetryExceededError); ok {
		return true
	}
	if _, ok := err.(NodeUnavailiableError); ok {
		return true
	}
	return false
}

// InvalidNodeResponse -
type InvalidNodeResponse struct {
	Errors []RunCodeError
	Raw    []byte
}

func newInvalidNodeResponse() InvalidNodeResponse {
	return InvalidNodeResponse{
		Errors: make([]RunCodeError, 0),
	}
}

// Error -
func (e InvalidNodeResponse) Error() string {
	var s strings.Builder
	if len(e.Errors) > 0 {
		for i := range e.Errors {
			if i > 0 {
				s.WriteByte('\n')
			}
			s.WriteString(e.Errors[i].ID)
		}
	} else {
		_, _ = s.Write(e.Raw)
	}
	return s.String()
}

// Is -
func (e InvalidNodeResponse) Is(target error) bool {
	_, ok := target.(InvalidNodeResponse)
	return ok
}

// Errors
var (
	ErrInvalidStatusCode = errors.New("invalid status code")
	ErrNodeRPCError      = "Node RPC error"
)
