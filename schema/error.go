package schema

import "strings"

const (
	ErrorAccountBalanceBoundary      ErrorClass = "AccountBalanceBoundaryException"
	ErrorArticleBarcodeAlreadyExists ErrorClass = "ArticleBarcodeAlreadyExistsException"
	ErrorArticleNotFound             ErrorClass = "ArticleNotFoundException"
	ErrorArticleInactive             ErrorClass = "ArticleInactiveException"
	ErrorParameterInvalid            ErrorClass = "ParameterInvalidException"
	ErrorParameterMissing            ErrorClass = "ParameterMissingException"
	ErrorParameterNotFound           ErrorClass = "ParameterNotFoundException"
	ErrorTransactionBoundary         ErrorClass = "TransactionBoundaryException"
	ErrorTransactionNotFound         ErrorClass = "TransactionNotFoundException"
	ErrorTransactionNotDeletable     ErrorClass = "TransactionNotDeletableException"
	ErrorUserAlreadyExists           ErrorClass = "UserAlreadyExistsException"
	ErrorUserNotFound                ErrorClass = "UserNotFoundException"
)

type (
	// Aliasing string allows us to implement a coustom JSON
	// Decoder that parses error paths like
	// "App\\Exception\\FooBar" as "FooBar"
	ErrorClass string

	// Structure of an API error
	ErrorResponse struct {
		Class   ErrorClass `json:"class"`   // type of error
		Message string     `json:"message"` // human-readable description of error

		// HTTP error code. Note that the server actually
		// returns 500 and only sets this code in its JSON
		// response. We omit this because of this behaviour,
		// and because Class already uniquely identifies an
		// error type
		code int `json:"code"`
	}

	SingleErrorResponse struct {
		Error ErrorResponse `json:"error"`
	}
)

func (e ErrorResponse) Error() string {
	return e.Message
}

func (e *ErrorClass) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	ss := strings.Split(s, "\\")
	*e = ErrorClass(ss[len(ss)-1])
	return nil
}
