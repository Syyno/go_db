package analyzer

import (
	domainErr "mydb/internal/domain/errors"
	"mydb/internal/domain/operations"

	"go.uber.org/zap"
)

type Analyzer struct {
	logger *zap.Logger
}

func New(log *zap.Logger) *Analyzer {
	return &Analyzer{logger: log}
}

func (a *Analyzer) Analyze(args []string) (operations.Operation, error) {
	opr, ok := operations.CommandToRulesMapping[operations.Command(args[0])]
	if !ok {
		return 0, domainErr.ErrUnknownOperation
	}

	if opr.ArgsCount != len(args) {
		return 0, domainErr.ErrArgsCountMismatch
	}

	return opr.Operation, nil
}
