package gitlab

import (
	"github.com/xanzy/go-gitlab"

	"github.com/devstream-io/devstream/pkg/util/git"
	"github.com/devstream-io/devstream/pkg/util/log"
)

// TODO(steinliber) gitlab does'nt support checkUpdate for now
func (c *Client) PushLocalFileToRepo(commitInfo *git.CommitInfo, checkUpdate bool) (bool, error) {
	createCommitOptions := c.CreateCommitInfo(gitlab.FileCreate, commitInfo)
	_, response, err := c.Commits.CreateCommit(c.GetRepoPath(), createCommitOptions)
	log.Debug(response.Body)
	if err != nil {
		return true, err
	}
	return false, nil
}

func (c *Client) CreateCommitInfo(action gitlab.FileActionValue, commitInfo *git.CommitInfo) *gitlab.CreateCommitOptions {
	var commitActionsOptions = make([]*gitlab.CommitActionOptions, 0, len(commitInfo.GitFileMap))

	for fileName, content := range commitInfo.GitFileMap {
		commitActionsOptions = append(commitActionsOptions, &gitlab.CommitActionOptions{
			Action:   gitlab.FileAction(action),
			FilePath: gitlab.String(fileName),
			Content:  gitlab.String(string(content)),
		})
	}
	return &gitlab.CreateCommitOptions{
		Branch:        gitlab.String(c.Branch),
		CommitMessage: gitlab.String(commitInfo.CommitMsg),
		Actions:       commitActionsOptions,
	}
}
