package gitlabcedocker

import (
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller"
	dockerInstaller "github.com/devstream-io/devstream/internal/pkg/plugininstaller/docker"
	"github.com/devstream-io/devstream/pkg/util/log"
)

func Read(options map[string]interface{}) (map[string]interface{}, error) {
	// 1. create config and pre-handle operations
	opts, err := validateAndDefault(options)
	if err != nil {
		return nil, err
	}

	// 2. config read operations
	operator := &plugininstaller.Operator{
		PreExecuteOperations: plugininstaller.PreExecuteOperations{
			dockerInstaller.Validate,
		},
		GetStateOperation: dockerInstaller.GetRunningState,
	}

	// 3. get status
	rawOptions, err := buildDockerOptions(opts).Encode()
	if err != nil {
		return nil, err
	}
	status, err := operator.Execute(rawOptions)
	if err != nil {
		return nil, err
	}

	log.Debugf("Return map: %v", status)
	return status, nil
}
