package surv

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cothi/port/utils"
)

type answer	struct {
	Port string `survey:"port"` // or you can tag fields to match a specific name
}

type answers struct {
	Port []string `survey:"port"`
}

func SelectPort(ports []string) answer {
	// the answers will be written to this struct
	var ans answer 

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

func MultiSelectPort(ports []string) answers {
	var ans answers


	// create question
	//listPort := []string{}
	var qs = []*survey.Question {
		{
			Name: "Port",
			Prompt: &survey.MultiSelect{
				Message: "Choose a port:",
				Options: ports,
				Default: "0",
			},
		},
	}
	survey.Ask(qs, &ans)

	fmt.Println(ans.Port)

	return ans

}
