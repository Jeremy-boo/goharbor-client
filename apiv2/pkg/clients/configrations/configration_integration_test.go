package configrations

import (
	"context"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAPIGetConfig(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	resp, err := c.GetConfigurationsInfo(ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestAPIUpdateConfig(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	resp, err := c.GetConfigurationsInfo(ctx)
	require.NoError(t, err)

	resp.AuthMode.Value = "oidc_auth"
	authMode, _ := resp.AuthMode.MarshalBinary()
	str := string(authMode)
	updateParams := &model.Configurations{
		AuthMode: &str,
	}
	updateErr := c.UpdateConfigurationsInfo(ctx, updateParams)
	require.NoError(t, updateErr)
}
