package config

type Binary struct {
	Go        string `default:"${BINARY_GO=go}" json:"go,omitempty"`
	Lint      string `default:"${BINARY_LINT=golangci-lint}" json:"lint,omitempty"`
	Upx       string `default:"${BINARY_UPX=upx}" json:"upx,omitempty"`
	Alignment string `default:"${BINARY_ALIGNMENT=fieldalignment}" json:"alignment,omitempty"`
}
