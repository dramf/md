package formatter

type Output struct {
	Results []Results `json:"Results"`
}
type Results struct {
	Target            string             `json:"Target"`
	Type              string             `json:"Type"`
	Vulnerabilities   []Vulnerability    `json:"Vulnerabilities"`
	Misconfigurations []Misconfiguration `json:"Misconfigurations,omitempty"`
	Secrets           []Secret           `json:"Secrets,omitempty"`
}

type Vulnerability struct {
	Library          string   `json:"PkgName"`
	VulnerabilityID  string   `json:"VulnerabilityID"`
	Severity         string   `json:"Severity"`
	Status           string   `json:"Status"`
	InstalledVersion string   `json:"InstalledVersion"`
	FixedVersion     string   `json:"FixedVersion"`
	Title            string   `json:"Title"`
	References       []string `json:"References"`
}

type Misconfiguration struct {
	ID            string        `json:"ID"`
	Title         string        `json:"Title"`
	Description   string        `json:"Description"`
	Severity      string        `json:"Severity"`
	Message       string        `json:"Message"`
	PrimaryURL    string        `json:"PrimaryURL"`
	CauseMetadata CauseMetadata `json:"CauseMetadata"`
}
type CauseMetadata struct {
	Code struct {
		Lines []LineInfo `json:"Lines"`
	}
}

type LineInfo struct {
	Number  int    `json:"Number"`
	Content string `json:"Content"`
}

type Secret struct {
	RuleID   string `json:"RuleID"`
	Category string `json:"Category"`
	Severity string `json:"Severity"`
	Title    string `json:"Title"`
}
