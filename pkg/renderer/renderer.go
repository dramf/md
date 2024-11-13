package renderer

import "github.com/aquasecurity/trivy/pkg/types"

type Renderer interface {
	RenderTitle(title string) string
	RenderEmptyTitle() string
	RenderVulnerabilities(v []types.DetectedVulnerability) string
	RenderMisconfigurations(m []types.DetectedMisconfiguration) string
	RenderSecrets(s []types.DetectedSecret) string
}
