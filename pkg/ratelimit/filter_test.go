// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
//
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

package ratelimit_test

import (
	"net"
	"testing"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/config"
	"go.thethings.network/lorawan-stack/v3/pkg/ratelimit"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func subnet(s string) net.IPNet {
	_, n, _ := net.ParseCIDR(s)
	return *n
}

func TestRateLimitFilter(t *testing.T) {
	for _, tc := range []struct {
		name      string
		config    config.RateLimiting
		maxPerMin map[ratelimit.Resource]int
	}{
		{
			name: "blacklist_and_whitelist",
			config: config.RateLimiting{
				Profiles: []config.RateLimitingProfile{
					{
						MaxPerMin:      20,
						Associations:   []string{"blacklist_and_whitelist"},
						OnlySubnets:    []net.IPNet{subnet("10.10.0.0/16")},
						ExcludeSubnets: []net.IPNet{subnet("10.10.10.0/24")},
					},
				},
			},
			maxPerMin: map[ratelimit.Resource]int{
				ratelimit.NewCustomResource("no_ip", "blacklist_and_whitelist"):                                          20,
				ratelimit.NewCustomResourceWithIP("ip_out_of_whitelist", net.IP{10, 1, 1, 1}, "blacklist_and_whitelist"): 0,
				ratelimit.NewCustomResourceWithIP("ip_in_blacklist", net.IP{10, 10, 1, 1}, "blacklist_and_whitelist"):    20,
				ratelimit.NewCustomResourceWithIP("ip_in_whitelist", net.IP{10, 10, 10, 1}, "blacklist_and_whitelist"):   0,
			},
		},
		{
			name: "subnet_best_match",
			config: config.RateLimiting{
				Profiles: []config.RateLimitingProfile{
					{
						MaxPerMin:    20,
						Associations: []string{"subnet_best_match"},
					},
					{
						MaxPerMin:    25,
						Associations: []string{"subnet_best_match"},
						OnlySubnets:  []net.IPNet{subnet("10.10.0.0/16")},
					},
					{
						MaxPerMin:    30,
						Associations: []string{"subnet_best_match"},
						OnlySubnets:  []net.IPNet{subnet("10.10.10.0/24")},
					},
				},
			},
			maxPerMin: map[ratelimit.Resource]int{
				ratelimit.NewCustomResource("no_ip", "subnet_best_match"):                                      20,
				ratelimit.NewCustomResourceWithIP("ip_out_of_range", net.IP{10, 0, 1, 2}, "subnet_best_match"): 20,
				ratelimit.NewCustomResourceWithIP("ip_in_16", net.IP{10, 10, 1, 2}, "subnet_best_match"):       25,
				ratelimit.NewCustomResourceWithIP("ip_in_24", net.IP{10, 10, 10, 2}, "subnet_best_match"):      30,
			},
		},
		{
			name: "blacklist_only",
			config: config.RateLimiting{
				Profiles: []config.RateLimitingProfile{
					{
						MaxPerMin:      20,
						Associations:   []string{"blacklist_only"},
						ExcludeSubnets: []net.IPNet{subnet("10.10.10.0/24")},
					},
				},
			},
			maxPerMin: map[ratelimit.Resource]int{
				ratelimit.NewCustomResource("no_ip", "blacklist_only"):                                       20,
				ratelimit.NewCustomResourceWithIP("ip_out_of_range", net.IP{10, 10, 1, 1}, "blacklist_only"): 20,
				ratelimit.NewCustomResourceWithIP("ip_in_range", net.IP{10, 10, 10, 1}, "blacklist_only"):    0,
			},
		},
		{
			name: "whitelist_only",
			config: config.RateLimiting{
				Profiles: []config.RateLimitingProfile{
					{
						MaxPerMin:    20,
						Associations: []string{"whitelist_only"},
						OnlySubnets:  []net.IPNet{subnet("10.10.10.0/24")},
					},
				},
			},
			maxPerMin: map[ratelimit.Resource]int{
				ratelimit.NewCustomResource("no_ip", "whitelist_only"):                                       20,
				ratelimit.NewCustomResourceWithIP("ip_out_of_range", net.IP{10, 10, 1, 1}, "whitelist_only"): 0,
				ratelimit.NewCustomResourceWithIP("ip_in_range", net.IP{10, 10, 10, 1}, "whitelist_only"):    20,
			},
		},
		{
			name: "blacklist_priority",
			config: config.RateLimiting{
				Profiles: []config.RateLimitingProfile{
					{
						MaxPerMin:      20,
						Associations:   []string{"blacklist_priority"},
						OnlySubnets:    []net.IPNet{subnet("10.10.10.0/24")},
						ExcludeSubnets: []net.IPNet{subnet("10.10.0.0/16")},
					},
				},
			},
			maxPerMin: map[ratelimit.Resource]int{
				ratelimit.NewCustomResource("no_ip", "blacklist_priority"):                                        20,
				ratelimit.NewCustomResourceWithIP("ip_out_of_range", net.IP{10, 1, 1, 1}, "blacklist_priority"):   0,
				ratelimit.NewCustomResourceWithIP("ip_in_blacklist", net.IP{10, 10, 1, 1}, "blacklist_priority"):  0,
				ratelimit.NewCustomResourceWithIP("ip_in_whitelist", net.IP{10, 10, 10, 1}, "blacklist_priority"): 0,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for resource, expectedLimit := range tc.maxPerMin {
				t.Run(resource.Key(), func(t *testing.T) {
					a := assertions.New(t)
					limiter, err := ratelimit.New(test.Context(), tc.config, config.BlobConfig{})
					if !a.So(err, should.BeNil) {
						t.FailNow()
					}
					_, r := limiter.RateLimit(resource)
					a.So(r.Limit, should.Resemble, expectedLimit)
				})
			}
		})
	}
}
