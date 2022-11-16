package main

import (
	"context"
)

func GetDeploymentValuesDownloadUrl(repo string, configuration Configuration) (*string, error) {
	ctx := context.Background()
	return configuration.GetDownloadURL(ctx, configuration.GetOrg(), repo, "deployment.yaml", nil)
}
