package config

type Binary struct {
	Go        string `default:"${BINARY_GO=go}" json:"go,omitempty"`
	Upx       string `default:"${BINARY_UPX=go}" json:"upx,omitempty"`
	Alignment string `default:"${BINARY_ALIGNMENT=fieldalignment}" json:"alignment,omitempty"`
}
