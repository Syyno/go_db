package compute

import (
	"mydb/internal/domain/operations"

	"go.uber.org/zap"
)

//go:generate minimock -i mydb/internal/database/compute.Parser -o parser/parser_mock.go -n ParserMock
type Parser interface {
	Parse(cmd string) ([]string, error)
}

//go:generate minimock -i mydb/internal/database/compute.Analyzer -o analyzer/analyzer_mock.go -n AnalyzerMock
type Analyzer interface {
	Analyze(args []string) (operations.Operation, error)
}

type Computer struct {
	analyzer Analyzer
	parser   Parser
	logger   *zap.Logger
}

func New(
	analyzer Analyzer,
	parser Parser,
	logger *zap.Logger) *Computer {
	return &Computer{
		analyzer: analyzer,
		parser:   parser,
		logger:   logger,
	}
}

// minimock -i Parser
type ComputingResult struct {
	OpId operations.Operation
	Args []string
}

func (c *Computer) Handle(cmd string) (*ComputingResult, error) {
	args, err := c.parser.Parse(cmd)
	if err != nil {
		return nil, err
	}

	op, err := c.analyzer.Analyze(args)
	if err != nil {
		return nil, err
	}

	return &ComputingResult{
		OpId: op,
		Args: args,
	}, nil
}
