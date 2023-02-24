package google

import (
	"context"
	"testing"

	"github.com/interflowrepo/interflow/interflow-wallet-api/configs"
	"github.com/interflowrepo/interflow/interflow-wallet-api/keys"
)

// Needs to be run manually with proper env configuration
// It's skipped during standard test execution
func TestGenerate(t *testing.T) {
	cfg := configs.ParseTestConfig(t)

	if cfg.DefaultKeyType != keys.AccountKeyTypeGoogleKMS {
		t.Skip("skipping since DefaultKeyType is not", keys.AccountKeyTypeGoogleKMS)
	}

	t.Run("key is generated", func(t *testing.T) {
		flowAccountKey, privateKey, err := Generate(cfg, context.Background(), 0, 1000)

		if err != nil {
			t.Fatal(err)
		}

		if flowAccountKey == nil {
			t.Fatal("Flow account key was not generated")
		}

		if privateKey == nil {
			t.Fatal("private key was not generated")
		}
	})
}
