package config

type Binary struct {
	Go        string `default:"${BINARY_GO=go}" json:"go,omitempty"`
	Alignment string `default:"${BINARY_ALIGNMENT=fieldalignment}" json:"alignment,omitempty"`
}
