package calculation

import "errors"

var (
	ErrInvalidExpression  = errors.New("expression is not valid")
	ErrDivisionByZero     = errors.New("division by zero")
	ErrUnbalancedBrackets = errors.New("unbalanced brackets")
	ErrUnknownSymbol      = errors.New("unknown symbol")
	ErrInternalServer     = errors.New("internal server error")
)
