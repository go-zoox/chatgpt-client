package chatgptclient

import openaiclient "github.com/go-zoox/openai-client"

func (c *client) ImageGeneration(cfg *openaiclient.ImageGenerationRequest) (*openaiclient.ImageGenerationResponse, error) {
	return c.core.ImageGeneration(cfg)
}
