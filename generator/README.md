# Control Struct

```
type AutoGenerated []struct {
	Name        string `json:"name"`
	Number      string `json:"number"`
	Description string `json:"description"`
	Why         string `json:"why"`
	Subcontrols []struct {
		Name        string `json:"name"`
		Number      string `json:"number"`
		Description string `json:"description"`
		Header      struct {
			AssetType            string   `json:"assetType"`
			SecurityFunction     string   `json:"securityFunction"`
			ImplementationGroups []string `json:"implementationGroups"`
		} `json:"header"`
		Dependencies []string `json:"dependencies"`
		Inputs       []struct {
			Name string `json:"name"`
		} `json:"inputs"`
		Operations []struct {
			Name  string      `json:"name"`
			Steps interface{} `json:"Steps"`
		} `json:"operations"`
		Measures []string `json:"measures"`
		Metrics  []struct {
			Name        string `json:"name"`
			Metric      string `json:"metric"`
			Calculation string `json:"calculation"`
		} `json:"metrics"`
	} `json:"subcontrols"`
}
```