package main

const githubTpl = `
{{- if . }}
{{- range . }}
{{- if or .Vulnerabilities .Misconfigurations .Secrets}}
## Target:` + "`{{ .Target }}`" + `
{{- if .Vulnerabilities }}
### Vulnerabilities
| Status | Library | Vulnerability ID | Severity | Installed Version | Fixed Version | Title | References |
|--------|---------|------------------|----------|-------------------|---------------|-------|------------|
{{- range .Vulnerabilities }}
| ` + "`{{ .Library }}`" + `| {{ .Status }} | {{ .VulnerabilityID }} | {{ .Severity }} | {{ .InstalledVersion }} | {{ if .FixedVersion }}{{ .FixedVersion }}{{ else }} {{ end }} | {{ .Title }} | 
  {{- range $index, $ref := .References }}{{ if $index }}, {{ end }}[{{ $ref }}]({{ $ref }}){{- end }} |
{{- end }}
{{- end }}

{{- if .Misconfigurations }}
### Misconfigurations
| ID | Title | Description | Severity | Message | Code |
|----|-------|-------------|----------|---------|------|
{{- range .Misconfigurations }}
| {{ .ID }} |{{ .Title }} | {{ .Description }} | {{ .Severity }} | {{ .Message }} <br> See {{ .PrimaryURL }} | {{ if .CauseMetadata.Code.Lines }} <pre> {{- range .CauseMetadata.Code.Lines }}{{ .Number }} {{ .Content }} <br> {{- end }} </pre> {{- end }}|
{{- end }}
{{- end }}

{{- if .Secrets }}
### Secrets
| RuleID | Category | Severity | Title |
|--------|----------|----------|-------|
{{- range .Secrets }}
|{{ .RuleID }}| {{ .Category }} |{{ .Severity }} | {{ .Title }} |
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- else }}
### Trivy Returned Empty Report
{{- end }}`
