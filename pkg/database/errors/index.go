package dbErrors

import (
	"database/sql"
	"github.com/lib/pq"
	httpErrors "github.com/martirosharutyunyan/axxon-test-task/pkg/common/http-errors"
	"strings"
)

func Parse(err error, messages ...string) error {
	if err == nil {
		return nil
	}

	if dberror, ok := err.(*pq.Error); ok {
		if strings.HasPrefix(dberror.Constraint, "unq") {
			return httpErrors.NewConflictError(constraintErrorsMap[dberror.Constraint])
		}
		if len(messages) != 0 {
			return httpErrors.NewNotFoundError(messages[0])
		}
		return httpErrors.NewInternalServerError(err.Error())
	}

	if err == sql.ErrNoRows {
		if len(messages) != 0 {
			return httpErrors.NewNotFoundError(messages[0])
		}

		return httpErrors.NewNotFoundError(err.Error())
	}

	if len(messages) != 0 {
		return httpErrors.NewInternalServerError(messages[0])
	}

	return httpErrors.NewInternalServerError(err.Error())
}
