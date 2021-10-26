//go:build integration

package project

import (
	"context"
	"fmt"
	"testing"

	clienterrors "github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"

	"github.com/stretchr/testify/require"
)

var (
	storageLimitPositive int64 = 1
	storageLimitNegative int64 = -1
)

func TestAPIProjectNew(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.NewProject(ctx, name, &storageLimitPositive)
	require.NoError(t, err)

	p, err := c.GetProject(ctx, name)
	require.NoError(t, err)

	defer c.DeleteProject(ctx, name)

	require.NoError(t, err)
	require.Equal(t, name, p.Name)
	require.False(t, p.Deleted)
}

func TestAPIProjectNew_UnlimitedStorage(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.NewProject(ctx, name, &storageLimitPositive)
	require.NoError(t, err)

	p, err := c.GetProject(ctx, name)
	require.NoError(t, err)
	defer c.DeleteProject(ctx, name)

	require.NoError(t, err)
	require.Equal(t, name, p.Name)
	require.False(t, p.Deleted)
}

func TestAPIProjectGet(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.NewProject(ctx, name, &storageLimitPositive)
	require.NoError(t, err)

	p, err := c.GetProject(ctx, name)
	require.NoError(t, err)
	require.NoError(t, err)
	defer c.DeleteProject(ctx, name)

	p2, err := c.GetProject(ctx, name)
	require.NoError(t, err)
	require.Equal(t, p, p2)
}

func TestAPIProjectDelete(t *testing.T) {
	name := "test-project"
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.NewProject(ctx, name, &storageLimitPositive)
	require.NoError(t, err)

	_, err = c.GetProject(ctx, name)
	require.NoError(t, err)

	err = c.DeleteProject(ctx, name)
	require.NoError(t, err)

	_, err = c.GetProject(ctx, name)
	require.Error(t, err)
	require.ErrorIs(t, err, &clienterrors.ErrProjectNotFound{})
}

func TestAPIProjectList(t *testing.T) {
	namePrefix := "test-project"
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("%s-%d", namePrefix, i)
		err := c.NewProject(ctx, name, &storageLimitPositive)
		require.NoError(t, err)

		_, err = c.GetProject(ctx, name)
		require.NoError(t, err)
		require.NoError(t, err)
		defer func() {
			defer c.DeleteProject(ctx, name)
			if err != nil {
				panic("error in cleanup routine: " + err.Error())
			}
		}()
	}

	projects, err := c.ListProjects(ctx, namePrefix)
	require.NoError(t, err)
	require.Len(t, projects, 10)
	for _, v := range projects {
		require.Contains(t, v.Name, namePrefix)
	}
}

func TestAPIProjectUpdate(t *testing.T) {
	name := "test-project"
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.NewProject(ctx, name, &storageLimitPositive)
	require.NoError(t, err)

	p, err := c.GetProject(ctx, name)
	require.NoError(t, err)
	require.NoError(t, err)

	defer c.DeleteProject(ctx, name)

	require.Equal(t, false, p.Togglable)

	p.Togglable = true

	err = c.UpdateProject(ctx, p, &storageLimitPositive)
	require.NoError(t, err)
	p2, err := c.GetProject(ctx, name)
	require.NoError(t, err)

	require.NotEqual(t, p, p2)
}