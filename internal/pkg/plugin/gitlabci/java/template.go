package java

import (
	"bytes"
	"html/template"
)

var gitlabCITemplate = `
image: docker:stable

stages:
  - package
  - docker_build
  - k8s_deploy

{{if .Package.Enable}}
mvn_package_job:
  image: {{.Package.BaseOption.Image}}
  stage: package
  tags:
    - {{.Package.BaseOption.Tags}}
  script: 
    {{range .Package.ScriptCommand}}
    - {{.}}
    {{end}} 
  artifacts:
    paths:
      - target/*.jar
  only:
    {{range .Package.BaseOption.AllowedBranch}}
    - {{.}}
    {{end}}
{{end}}

{{if .Build.Enable}}
docker_build_job:
  image: {{.Build.BaseOption.Image}}
  stage: docker_build
  tags: 
    - {{.Build.BaseOption.Tags}}
  script:
    {{range .Build.ScriptCommand}}
    - {{.}}
    {{end}} 
  only:
    {{range .Build.BaseOption.AllowedBranch}}
    - {{.}}
    {{end}}
{{end}}

{{if .Deploy.Enable}}
k8s_deploy_job:
  image: 
    name: {{.Deploy.BaseOption.Image}}
    entrypoint: [""]
  stage: k8s_deploy
  tags: 
    - {{.Deploy.BaseOption.Tags}}
  script:
    {{range .Deploy.ScriptCommand}}
    - {{.}}
    {{end}} 
  only:
    {{range .Deploy.BaseOption.AllowedBranch}}
    - {{.}}
    {{end}}
{{end}}
`

// Render gitlab-ci.yml template with Options
func renderTmpl(Opts *Options) (string, error) {
	t := template.Must(template.New("gitlabci-java").Option("missingkey=error").Parse(gitlabCITemplate))

	var buf bytes.Buffer
	err := t.Execute(&buf, Opts)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
