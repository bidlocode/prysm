package direct

import (
	"context"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
	keystorev4 "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"

	"github.com/prysmaticlabs/prysm/shared/bls"
)

var log = logrus.WithField("prefix", "keymanager-v2")

// Config for a direct keymanager.
type Config struct{}

// Keymanager implementation for direct keystores.
type Keymanager struct{}

// DefaultConfig for a direct keymanager implementation.
func DefaultConfig() *Config {
	return &Config{}
}

// NewKeymanager instantiates a new direct keymanager from configuration options.
func NewKeymanager(ctx context.Context, cfg *Config) *Keymanager {
	return &Keymanager{}
}

// CreateAccount for a direct keymanager implementation.
func (dr *Keymanager) CreateAccount(ctx context.Context, password string) error {
	if password == "" {
		return errors.New("account password must not be empty")
	}
	encryptor := keystorev4.New()
	signingKey := bls.RandKey()
	// Withdrawal key will be shown by using a mnemonic.
	withdrawalKey := bls.RandKey()
	keystore, err := encryptor.Encrypt(signingKey.Marshal(), []byte(password))
	if err != nil {
		return errors.Wrap(err, "could not encrypt signing key")
	}
	return errors.New("unimplemented")
}

// ConfigFile returns a marshaled configuration file for a direct keymanager.
func (dr *Keymanager) ConfigFile(ctx context.Context) ([]byte, error) {
	return nil, nil
}