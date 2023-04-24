package services

type AiClient interface {
	AskAi(*string) (string, error)
}
