package disk

import (
	"context"
	"testing"

	"github.com/spiffe/spire/pkg/server/plugin/keymanager/keymanagertest"
	"github.com/spiffe/spire/proto/common/plugin"
	"github.com/spiffe/spire/proto/server/keymanager"
	"github.com/stretchr/testify/require"
)

func TestKeyManager(t *testing.T) {
	keymanagertest.Run(t, makeKeyManager)
}

func makeKeyManager(t *testing.T) keymanager.Plugin {
	m := New()
	resp, err := m.Configure(context.Background(), &plugin.ConfigureRequest{})
	require.NoError(t, err)
	require.Equal(t, &plugin.ConfigureResponse{}, resp)
	return m
}
