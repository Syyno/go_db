package operations

type Operation int

const (
	Get Operation = iota
	Set Operation = iota
	Del Operation = iota
)

type Command string

const (
	GetCommand Command = "GET"
	SetCommand Command = "SET"
	DelCommand Command = "DEL"
)

const (
	ArgsMinCount = 2
	ArgsMaxCount = 3
)

type OperationArgsRules struct {
	Operation Operation
	ArgsCount int
}

var (
	CommandToRulesMapping = map[Command]OperationArgsRules{
		GetCommand: {
			Operation: Get,
			ArgsCount: 2,
		},
		SetCommand: {
			Operation: Set,
			ArgsCount: 3,
		},
		DelCommand: {
			Operation: Del,
			ArgsCount: 2,
		},
	}
)
