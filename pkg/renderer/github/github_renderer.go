package github

import (
	"fmt"
	"github.com/aquasecurity/trivy/pkg/types"
)

const maxReferences = 3

type MarkdownGithubRenderer struct{}

func (m MarkdownGithubRenderer) RenderTitle(target string) string {
	return fmt.Sprintf("## Target `%s`\n", target)
}

func (m MarkdownGithubRenderer) RenderEmptyTitle() string {
	return "### Trivy Returned Empty Report\n"
}

func (m MarkdownGithubRenderer) RenderVulnerabilities(v []types.DetectedVulnerability) string {
	if len(v) == 0 {
		return ""
	}
	tableTitle := "### Vulnerabilities\n" +
		"| Library | Status | Vulnerability ID | Severity | Installed Version | Fixed Version | Title | References |\n" +
		"|---------|--------|------------------|----------|-------------------|---------------|-------|------------|\n"
	var vulnerabilities string
	for _, vulnerability := range v {
		var references string
		r := vulnerability.References
		for i := range r {
			references += fmt.Sprintf("%s", r[i])
			if i != maxReferences-1 && i != len(r)-1 {
				references += fmt.Sprintf(", <br>")
			}
			if i == maxReferences-1 {
				break
			}
		}
		vulnerabilities += fmt.Sprintf("| `%s` | %s | %s | %s | %s | %s | %s | %s |\n",
			vulnerability.PkgName,
			vulnerability.Status.String(),
			vulnerability.VulnerabilityID,
			vulnerability.Severity,
			vulnerability.InstalledVersion,
			vulnerability.FixedVersion,
			vulnerability.Title,
			references)
	}
	return tableTitle + vulnerabilities
}

func (m MarkdownGithubRenderer) RenderMisconfigurations(conf []types.DetectedMisconfiguration) string {
	if len(conf) == 0 {
		return ""
	}
	tableTitle := "### Misconfigurations\n" +
		"| ID | Title | Description | Severity | Message | Code |\n" +
		"|----|-------|-------------|----------|---------|------|\n"
	var misconfigurations string
	for _, c := range conf {
		var code string
		for _, line := range c.CauseMetadata.Code.Lines {
			code += fmt.Sprintf("%d %s <br>", line.Number, line.Content)
		}
		if len(code) > 0 {
			code = fmt.Sprintf("<pre>%s</pre>", code)
		}
		misconfigurations += fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n",
			c.ID,
			c.Title,
			c.Description,
			c.Severity,
			fmt.Sprintf("%s <br> See %s", c.Message, c.PrimaryURL),
			code,
		)
	}
	return tableTitle + misconfigurations

}
func (m MarkdownGithubRenderer) RenderSecrets(s []types.DetectedSecret) string {
	if len(s) == 0 {
		return ""
	}
	tableTitle := "### Secrets\n" +
		"| RuleID | Category | Severity | Title |\n" +
		"|--------|----------|----------|-------|\n"

	var secrets string
	for _, secret := range s {
		secrets += fmt.Sprintf("| %s | %s | %s | %s |\n",
			secret.RuleID,
			secret.Category,
			secret.Severity,
			secret.Title)
	}
	return tableTitle + secrets
}
