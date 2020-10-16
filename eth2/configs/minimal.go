package configs

import (
	"github.com/protolambda/zrnt/eth2/beacon"
)

var Minimal = &beacon.Spec{
	CONFIG_NAME: "minimal",
	Phase0Config: beacon.Phase0Config{
		MAX_COMMITTEES_PER_SLOT:               4,
		TARGET_COMMITTEE_SIZE:                 4,
		MAX_VALIDATORS_PER_COMMITTEE:          2048,
		MIN_PER_EPOCH_CHURN_LIMIT:             4,
		CHURN_LIMIT_QUOTIENT:                  0x10000,
		SHUFFLE_ROUND_COUNT:                   10,
		MIN_GENESIS_ACTIVE_VALIDATOR_COUNT:    64,
		MIN_GENESIS_TIME:                      0x5e0e8400,
		HYSTERESIS_QUOTIENT:                   4,
		HYSTERESIS_DOWNWARD_MULTIPLIER:        1,
		HYSTERESIS_UPWARD_MULTIPLIER:          5,
		PROPORTIONAL_SLASHING_MULTIPLIER:      3,
		SAFE_SLOTS_TO_UPDATE_JUSTIFIED:        2,
		ETH1_FOLLOW_DISTANCE:                  16,
		TARGET_AGGREGATORS_PER_COMMITTEE:      16,
		RANDOM_SUBNETS_PER_VALIDATOR:          1,
		EPOCHS_PER_RANDOM_SUBNET_SUBSCRIPTION: 256,
		SECONDS_PER_ETH1_BLOCK:                14,
		DEPOSIT_CHAIN_ID:                      5,
		DEPOSIT_NETWORK_ID:                    5,
		DEPOSIT_CONTRACT_ADDRESS:              [20]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90},
		MIN_DEPOSIT_AMOUNT:                    0x3b9aca00,
		MAX_EFFECTIVE_BALANCE:                 0x773594000,
		EJECTION_BALANCE:                      0x3b9aca000,
		EFFECTIVE_BALANCE_INCREMENT:           0x3b9aca00,
		GENESIS_FORK_VERSION:                  beacon.Version{0x00, 0x00, 0x00, 0x01},
		BLS_WITHDRAWAL_PREFIX:                 [1]byte{0x00},
		GENESIS_DELAY:                         300,
		SECONDS_PER_SLOT:                      6,
		MIN_ATTESTATION_INCLUSION_DELAY:       1,
		SLOTS_PER_EPOCH:                       8,
		MIN_SEED_LOOKAHEAD:                    1,
		MAX_SEED_LOOKAHEAD:                    4,
		EPOCHS_PER_ETH1_VOTING_PERIOD:         4,
		SLOTS_PER_HISTORICAL_ROOT:             64,
		MIN_VALIDATOR_WITHDRAWABILITY_DELAY:   256,
		SHARD_COMMITTEE_PERIOD:                64,
		MIN_EPOCHS_TO_INACTIVITY_PENALTY:      4,
		EPOCHS_PER_HISTORICAL_VECTOR:          64,
		EPOCHS_PER_SLASHINGS_VECTOR:           64,
		HISTORICAL_ROOTS_LIMIT:                0x1000000,
		VALIDATOR_REGISTRY_LIMIT:              0x10000000000,
		BASE_REWARD_FACTOR:                    64,
		WHISTLEBLOWER_REWARD_QUOTIENT:         512,
		PROPOSER_REWARD_QUOTIENT:              8,
		INACTIVITY_PENALTY_QUOTIENT:           0x1000000,
		MIN_SLASHING_PENALTY_QUOTIENT:         32,
		MAX_PROPOSER_SLASHINGS:                16,
		MAX_ATTESTER_SLASHINGS:                2,
		MAX_ATTESTATIONS:                      128,
		MAX_DEPOSITS:                          16,
		MAX_VOLUNTARY_EXITS:                   16,
		DOMAIN_BEACON_PROPOSER:                beacon.BLSDomainType{0x00, 0x00, 0x00, 0x00},
		DOMAIN_BEACON_ATTESTER:                beacon.BLSDomainType{0x01, 0x00, 0x00, 0x00},
		DOMAIN_RANDAO:                         beacon.BLSDomainType{0x02, 0x00, 0x00, 0x00},
		DOMAIN_DEPOSIT:                        beacon.BLSDomainType{0x03, 0x00, 0x00, 0x00},
		DOMAIN_VOLUNTARY_EXIT:                 beacon.BLSDomainType{0x04, 0x00, 0x00, 0x00},
		DOMAIN_SELECTION_PROOF:                beacon.BLSDomainType{0x05, 0x00, 0x00, 0x00},
		DOMAIN_AGGREGATE_AND_PROOF:            beacon.BLSDomainType{0x06, 0x00, 0x00, 0x00},
	},
	Phase1Config: beacon.Phase1Config{
		PHASE_1_FORK_VERSION:                             beacon.Version{0x01, 0x00, 0x00, 0x01},
		PHASE_1_FORK_SLOT:                                0,
		INITIAL_ACTIVE_SHARDS:                            2,
		MAX_SHARDS:                                       8,
		LIGHT_CLIENT_COMMITTEE_SIZE:                      128,
		GASPRICE_ADJUSTMENT_COEFFICIENT:                  8,
		MAX_SHARD_BLOCK_SIZE:                             0x100000,
		TARGET_SHARD_BLOCK_SIZE:                          0x40000,
		SHARD_BLOCK_OFFSETS:                              []uint64{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233},
		MAX_SHARD_BLOCKS_PER_ATTESTATION:                 12,
		BYTES_PER_CUSTODY_CHUNK:                          4096,
		CUSTODY_RESPONSE_DEPTH:                           8,
		MAX_GASPRICE:                                     0x4000,
		MIN_GASPRICE:                                     8,
		ONLINE_PERIOD:                                    8,
		LIGHT_CLIENT_COMMITTEE_PERIOD:                    256,
		MAX_CUSTODY_CHUNK_CHALLENGE_RECORDS:              0x100000,
		DOMAIN_SHARD_PROPOSAL:                            beacon.BLSDomainType{0x80, 0x00, 0x00, 0x00},
		DOMAIN_SHARD_COMMITTEE:                           beacon.BLSDomainType{0x81, 0x00, 0x00, 0x00},
		DOMAIN_LIGHT_CLIENT:                              beacon.BLSDomainType{0x82, 0x00, 0x00, 0x00},
		DOMAIN_CUSTODY_BIT_SLASHING:                      beacon.BLSDomainType{0x83, 0x00, 0x00, 0x00},
		DOMAIN_LIGHT_SELECTION_PROOF:                     beacon.BLSDomainType{0x84, 0x00, 0x00, 0x00},
		DOMAIN_LIGHT_AGGREGATE_AND_PROOF:                 beacon.BLSDomainType{0x85, 0x00, 0x00, 0x00},
		RANDAO_PENALTY_EPOCHS:                            2,
		EARLY_DERIVED_SECRET_PENALTY_MAX_FUTURE_EPOCHS:   64,
		EPOCHS_PER_CUSTODY_PERIOD:                        32,
		CUSTODY_PERIOD_TO_RANDAO_PADDING:                 8,
		MAX_CHUNK_CHALLENGE_DELAY:                        64,
		CUSTODY_PRIME:                                    mustBigInt("115792089237316195423570985008687907853269984665640564039457584007913129639747"),
		CUSTODY_SECRETS:                                  3,
		BYTES_PER_CUSTODY_ATOM:                           32,
		CUSTODY_PROBABILITY_EXPONENT:                     2,
		MAX_CUSTODY_KEY_REVEALS:                          256,
		MAX_EARLY_DERIVED_SECRET_REVEALS:                 1,
		MAX_CUSTODY_CHUNK_CHALLENGES:                     2,
		MAX_CUSTODY_CHUNK_CHALLENGE_RESP:                 8,
		MAX_CUSTODY_SLASHINGS:                            1,
		EARLY_DERIVED_SECRET_REVEAL_SLOT_REWARD_MULTIPLE: 2,
		MINOR_REWARD_QUOTIENT:                            256,
	},
}
