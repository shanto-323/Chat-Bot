package main

type Message struct {
	Role    string      `json:"role"`
	Content string      `json:"content"`
	Refusal interface{} `json:"refusal"` // Optional field
}

type Choice struct {
	Logprobs           interface{} `json:"logprobs"` // Optional field
	FinishReason       string      `json:"finish_reason"`
	NativeFinishReason string      `json:"native_finish_reason"`
	Index              int         `json:"index"`
	Message            Message     `json:"message"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChatCompletionResponse struct {
	ID                string   `json:"id"`
	Provider          string   `json:"provider"`
	Model             string   `json:"model"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Choices           []Choice `json:"choices"`
	SystemFingerprint *string  `json:"system_fingerprint"` // Optional field
	Usage             Usage    `json:"usage"`
}

type Payload struct {
	Model    string     `json:"model"`
	Messages []Messages `json:"messages"`
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
