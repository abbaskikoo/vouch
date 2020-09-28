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

package prometheus

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func (s *Service) setupControllerMetrics() error {
	controllerStartTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "vouch",
		Name:      "start_time_secs",
		Help:      "The timestamp at which vouch started.",
	})
	if err := prometheus.Register(controllerStartTime); err != nil {
		return err
	}
	controllerStartTime.SetToCurrentTime()

	s.epochsProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "vouch",
		Name:      "epochs_processed_total",
		Help:      "The number of epochs vouch has processed.",
	})
	if err := prometheus.Register(s.epochsProcessed); err != nil {
		return err
	}

	s.blockReceiptDelay =
		prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: "vouch",
			Name:      "block_receipt_delay_seconds",
			Help:      "The delay between the start of a slot and the time vouch receives it.",
			Buckets: []float64{
				0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0,
				1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0,
				2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7, 2.8, 2.9, 3.0,
				3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 3.8, 3.9, 4.0,
			},
		})
	if err := prometheus.Register(s.blockReceiptDelay); err != nil {
		return err
	}

	return nil
}

// NewEpoch is called when vouch starts processing a new epoch.
func (s *Service) NewEpoch() {
	s.epochsProcessed.Inc()
}

// BlockDelay provides the delay between the start of a slot and vouch receiving its block.
func (s *Service) BlockDelay(delay time.Duration) {
	s.blockReceiptDelay.Observe(delay.Seconds())
}