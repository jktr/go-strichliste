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
	ErrorClass string

	ErrorResponse struct {
		Class   ErrorClass `json:"class"`
		Code    int        `json:"code"`
		Message string     `json:"message"`
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
