// Copyright © 2020 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package multinode

import (
	"context"

	eth2client "github.com/attestantio/go-eth2-client"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	zerologger "github.com/rs/zerolog/log"
)

// Service is the provider for beacon block proposals.
type Service struct {
	processConcurrency                    int64
	beaconBlockSubmitters                 map[string]eth2client.BeaconBlockSubmitter
	attestationSubmitters                 map[string]eth2client.AttestationSubmitter
	aggregateAttestationsSubmitters       map[string]eth2client.AggregateAttestationsSubmitter
	beaconCommitteeSubscriptionSubmitters map[string]eth2client.BeaconCommitteeSubscriptionsSubmitter
}

// module-wide log.
var log zerolog.Logger

// New creates a new beacon block propsal strategy.
func New(ctx context.Context, params ...Parameter) (*Service, error) {
	parameters, err := parseAndCheckParameters(params...)
	if err != nil {
		return nil, errors.Wrap(err, "problem with parameters")
	}

	// Set logging.
	log = zerologger.With().Str("strategy", "submitter").Str("impl", "all").Logger()
	if parameters.logLevel != log.GetLevel() {
		log = log.Level(parameters.logLevel)
	}

	s := &Service{
		processConcurrency:                    parameters.processConcurrency,
		beaconBlockSubmitters:                 parameters.beaconBlockSubmitters,
		attestationSubmitters:                 parameters.attestationSubmitters,
		aggregateAttestationsSubmitters:       parameters.aggregateAttestationsSubmitters,
		beaconCommitteeSubscriptionSubmitters: parameters.beaconCommitteeSubscriptionsSubmitters,
	}

	return s, nil
}
