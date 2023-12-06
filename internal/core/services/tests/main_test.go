package services_tests

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/vault/api"

	"github.com/rarimo/issuer-node/internal/config"
	"github.com/rarimo/issuer-node/internal/db"
	"github.com/rarimo/issuer-node/internal/db/tests"
	"github.com/rarimo/issuer-node/internal/kms"
	"github.com/rarimo/issuer-node/internal/loader"
	"github.com/rarimo/issuer-node/internal/log"
	"github.com/rarimo/issuer-node/internal/providers"
	"github.com/rarimo/issuer-node/pkg/cache"
)

var (
	storage        *db.Storage
	vaultCli       *api.Client
	bjjKeyProvider kms.KeyProvider
	keyStore       *kms.KMS
	cachex         cache.Cache
	docLoader      loader.DocumentLoader
	cfg            config.Configuration
)

const ipfsGatewayURL = "http://127.0.0.1:8080"

const ipfsGateway = "http://localhost:8080"

func TestMain(m *testing.M) {
	ctx := context.Background()
	conn := lookupPostgresURL()
	if conn == "" {
		conn = "postgres://postgres:postgres@localhost:5435"
	}

	cfgForTesting := config.Configuration{
		Database: config.Database{
			URL: conn,
		},
		KeyStore: config.VaultTest(),
	}
	s, teardown, err := tests.NewTestStorage(&cfgForTesting)
	defer teardown()
	if err != nil {
		log.Error(ctx, "failed to acquire test database", "err", err)
		os.Exit(1)
	}
	storage = s

	vaultCli, err = providers.VaultClient(ctx, providers.Config{
		Address: cfgForTesting.KeyStore.Address,
		Token:   cfgForTesting.KeyStore.Token,
	})
	if err != nil {
		log.Error(ctx, "failed to acquire vault client", "err", err)
		os.Exit(1)
	}

	bjjKeyProvider, err = kms.NewVaultPluginIden3KeyProvider(vaultCli, cfgForTesting.KeyStore.PluginIden3MountPath, kms.KeyTypeBabyJubJub)
	if err != nil {
		log.Error(ctx, "failed to create Iden3 Key Provider", "err", err)
		os.Exit(1)
	}

	ethKeyProvider, err := kms.NewVaultPluginIden3KeyProvider(vaultCli, cfgForTesting.KeyStore.PluginIden3MountPath, kms.KeyTypeEthereum)
	if err != nil {
		log.Error(ctx, "failed to create Iden3 Key Provider", "err", err)
		os.Exit(1)
	}

	keyStore = kms.NewKMS()
	err = keyStore.RegisterKeyProvider(kms.KeyTypeBabyJubJub, bjjKeyProvider)
	if err != nil {
		log.Error(ctx, "failed to register Key Provider", "err", err)
		os.Exit(1)
	}

	err = keyStore.RegisterKeyProvider(kms.KeyTypeEthereum, ethKeyProvider)
	if err != nil {
		log.Error(ctx, "failed to register eth Key Provider", "err", err)
		os.Exit(1)
	}

	cachex = cache.NewMemoryCache()

	docLoader = loader.NewDocumentLoader(ipfsGatewayURL)

	cfg.CredentialStatus = config.CredentialStatus{
		RHSMode: "None",
		DirectStatus: config.DirectStatus{
			URL: "http://localhost:3001",
		},
	}
	m.Run()
}

func lookupPostgresURL() string {
	con, ok := os.LookupEnv("POSTGRES_TEST_DATABASE")
	if !ok {
		return ""
	}
	return con
}
