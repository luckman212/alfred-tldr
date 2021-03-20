package tldr

import (
	"bufio"
	"io"
	"strings"
)

// Page contents of a tldr page
type Page struct {
	CmdName         string
	CmdDescriptions []string
	CmdExamples     []*CmdExample
}

// CmdExample a command example in a tldr page
type CmdExample struct {
	Description string
	Cmd         string
}

func parsePage(s io.Reader) (*Page, error) {
	cmdExamples := make([]*CmdExample, 0)
	var cmdName, description, cmd string
	var cmdDescriptions []string
	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 0 {
			continue
		}
		if strings.HasPrefix(line, "#") {
			cmdName = strings.TrimSpace(strings.TrimLeft(line, "#"))
		}
		if strings.HasPrefix(line, ">") {
			trimedLine := strings.TrimSpace(strings.TrimLeft(line, ">"))
			cmdDescriptions = append(cmdDescriptions, trimedLine)
		}
		if strings.HasPrefix(line, "-") {
			description = strings.TrimSpace(strings.TrimLeft(line, "-"))
		}
		if strings.HasPrefix(line, "`") {
			cmd = strings.TrimSpace(strings.Trim(line, "`"))
			cmdExamples = append(cmdExamples, &CmdExample{
				Description: description,
				Cmd:         cmd,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &Page{
		CmdName:         cmdName,
		CmdDescriptions: cmdDescriptions,
		CmdExamples:     cmdExamples,
	}, nil
}
