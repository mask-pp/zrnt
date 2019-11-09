// +build preset_mainnet

package constant_presets

const PRESET_NAME string = "mainnet"

const MAX_COMMITTEES_PER_SLOT = 64

const TARGET_COMMITTEE_SIZE = 128

const MAX_VALIDATORS_PER_COMMITTEE = 2048

const MIN_PER_EPOCH_CHURN_LIMIT = 4

const CHURN_LIMIT_QUOTIENT = 65536

const SHUFFLE_ROUND_COUNT = 90

const MIN_GENESIS_ACTIVE_VALIDATOR_COUNT = 65536

const MIN_GENESIS_TIME = 1578009600

var DEPOSIT_CONTRACT_ADDRESS = [40]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90}

const MIN_DEPOSIT_AMOUNT = 1000000000

const MAX_EFFECTIVE_BALANCE = 32000000000

const EJECTION_BALANCE = 16000000000

const EFFECTIVE_BALANCE_INCREMENT = 1000000000

const GENESIS_SLOT = 0

const BLS_WITHDRAWAL_PREFIX = 0

const SECONDS_PER_SLOT = 12

const MIN_ATTESTATION_INCLUSION_DELAY = 1

const SLOTS_PER_EPOCH = 32

const MIN_SEED_LOOKAHEAD = 1

const MAX_SEED_LOOKAHEAD = 4

const SLOTS_PER_ETH1_VOTING_PERIOD = 1024

const SLOTS_PER_HISTORICAL_ROOT = 8192

const MIN_VALIDATOR_WITHDRAWABILITY_DELAY = 256

const PERSISTENT_COMMITTEE_PERIOD = 2048

const MAX_EPOCHS_PER_CROSSLINK = 64

const MIN_EPOCHS_TO_INACTIVITY_PENALTY = 4

const EARLY_DERIVED_SECRET_PENALTY_MAX_FUTURE_EPOCHS = 16384

const EPOCHS_PER_HISTORICAL_VECTOR = 65536

const EPOCHS_PER_SLASHINGS_VECTOR = 8192

const HISTORICAL_ROOTS_LIMIT = 16777216

const VALIDATOR_REGISTRY_LIMIT = 1099511627776

const BASE_REWARD_FACTOR = 64

const WHISTLEBLOWER_REWARD_QUOTIENT = 512

const PROPOSER_REWARD_QUOTIENT = 8

const INACTIVITY_PENALTY_QUOTIENT = 33554432

const MIN_SLASHING_PENALTY_QUOTIENT = 32

const MAX_PROPOSER_SLASHINGS = 16

const MAX_ATTESTER_SLASHINGS = 1

const MAX_ATTESTATIONS = 128

const MAX_DEPOSITS = 16

const MAX_VOLUNTARY_EXITS = 16

const DOMAIN_BEACON_PROPOSER = 0x00000000

const DOMAIN_BEACON_ATTESTER = 0x01000000

const DOMAIN_RANDAO = 0x02000000

const DOMAIN_DEPOSIT = 0x03000000

const DOMAIN_VOLUNTARY_EXIT = 0x04000000

const DOMAIN_CUSTODY_BIT_CHALLENGE = 0x06000000

const DOMAIN_SHARD_PROPOSER = 0x80000000

const DOMAIN_SHARD_ATTESTER = 0x81000000
