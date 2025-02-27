package jiragithub

import (
	"fmt"

	"github.com/mitchellh/mapstructure"

	"github.com/devstream-io/devstream/pkg/util/git"
	"github.com/devstream-io/devstream/pkg/util/github"
	"github.com/devstream-io/devstream/pkg/util/log"
)

// Read get jira-github-integ workflows.
func Read(options map[string]interface{}) (map[string]interface{}, error) {
	var opts Options
	err := mapstructure.Decode(options, &opts)
	if err != nil {
		return nil, err
	}

	if errs := validate(&opts); len(errs) != 0 {
		for _, e := range errs {
			log.Errorf("Options error: %s.", e)
		}
		return nil, fmt.Errorf("options are illegal")
	}

	ghOptions := &git.RepoInfo{
		Owner:    opts.Owner,
		Org:      opts.Org,
		Repo:     opts.Repo,
		NeedAuth: true,
	}
	ghClient, err := github.NewClient(ghOptions)
	if err != nil {
		return nil, err
	}

	path, err := ghClient.GetWorkflowPath()
	if err != nil {
		return nil, err
	}

	return BuildReadState(path), nil
}

func BuildReadState(path string) map[string]interface{} {
	res := make(map[string]interface{})
	res["workflowDir"] = path
	return res
}
