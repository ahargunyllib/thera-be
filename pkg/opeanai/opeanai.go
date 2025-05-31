package openai

import (
	"context"

	"github.com/ahargunyllib/thera-be/internal/infra/env"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type Message struct {
	Content string
	Role    string
}

type CustomOpenAIInterface interface {
	Chat(ctx context.Context, messages []Message) (*openai.ChatCompletion, error)
}

type CustomOpenAIStruct struct {
	client *openai.Client
}

func getOpenAI() CustomOpenAIInterface {
	client := openai.NewClient(
		option.WithAPIKey(env.AppEnv.OpenAIKey),
	)

	return &CustomOpenAIStruct{
		client: &client,
	}
}

var OpenAI = getOpenAI()

func (o *CustomOpenAIStruct) Chat(ctx context.Context, messages []Message) (*openai.ChatCompletion, error) {
	req := openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT4o,
	}

	openAIMessages := make([]openai.ChatCompletionMessageParamUnion, len(messages))
	for idx, msg := range messages {
		if msg.Role == "user" {
			openAIMessages[idx] = openai.UserMessage(msg.Content)
			continue
		}

		if msg.Role == "assistant" {
			openAIMessages[idx] = openai.AssistantMessage(msg.Content)
			continue
		}
	}

	res, err := o.client.Chat.Completions.New(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
