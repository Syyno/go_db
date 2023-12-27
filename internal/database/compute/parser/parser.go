package parser

import (
	domainErr "mydb/internal/domain/errors"
	"mydb/internal/domain/operations"
	"regexp"

	"go.uber.org/zap"
)

type Parser struct {
	logger *zap.Logger
}

func New(log *zap.Logger) *Parser {
	return &Parser{logger: log}
}

// todo : rewrite to state machine
func (p *Parser) Parse(cmd string) ([]string, error) {
	r := regexp.MustCompile(`[\w|*|/]+`)
	matches := r.FindAllString(cmd, -1)
	if len(matches) < operations.ArgsMinCount || len(matches) > operations.ArgsMaxCount {
		return nil, domainErr.ErrArgsCountIsInvalid
	}

	return matches, nil
}
