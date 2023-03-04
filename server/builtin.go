package server

import (
	"github.com/polygomic/polygomic-edge/consensus"
	consensusDev "github.com/polygomic/polygomic-edge/consensus/dev"
	consensusDummy "github.com/polygomic/polygomic-edge/consensus/dummy"
	consensusIBFT "github.com/polygomic/polygomic-edge/consensus/ibft"
	"github.com/polygomic/polygomic-edge/secrets"
	"github.com/polygomic/polygomic-edge/secrets/awsssm"
	"github.com/polygomic/polygomic-edge/secrets/gcpssm"
	"github.com/polygomic/polygomic-edge/secrets/hashicorpvault"
	"github.com/polygomic/polygomic-edge/secrets/local"
)

type ConsensusType string

const (
	DevConsensus   ConsensusType = "dev"
	IBFTConsensus  ConsensusType = "ibft"
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:   consensusDev.Factory,
	IBFTConsensus:  consensusIBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
