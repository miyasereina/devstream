package generic

// Options is the struct for configurations of the githubci-generic plugin.
type Options struct {
	PathWithNamespace string
	Branch            string
	TemplateURL       string
	TemplateVariables map[string]interface{}
}
