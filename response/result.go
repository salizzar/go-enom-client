package response

type CommandResult struct {
	Command		string	`xml:"Command"  json:"command"`
	Language	string	`xml:"Language" json:"language"`
	ErrCount	int	`xml:"ErrCount" json:"err_count"`
}

type Result struct {
	ResponseParser		`json:"-"`

	CommandResult		`json:"command_result"`
}

