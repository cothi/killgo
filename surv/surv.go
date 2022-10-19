package surv

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/cothi/port/utils"
)

type answers struct {
	Port string `survey:"port"` // or you can tag fields to match a specific name
}

func SelectPort(ports []string) answers {
	// the answers will be written to this struct
	var ans answers

	// create question
	var qs = []*survey.Question{
		{
			Name: "Port",
			Prompt: &survey.Select{
				Message: "Choose a port:",
				Options: ports,
				Default: "0",
			},
		},
	}
	err := survey.Ask(qs, &ans)
	utils.ErrorCheck(err)
	return ans
}
