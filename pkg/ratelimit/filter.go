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

package ratelimit

import (
	"net"
)

type rateLimiterWithSubnet struct {
	Interface
	whitelist []net.IPNet
	blacklist []net.IPNet
}

type rateLimiterChooser struct {
	pool []rateLimiterWithSubnet
}

// subnetListEqual checks if two subnet lists are equal.
// For simplicity, the string representation of the subnets is used for the equality check.
func subnetListEqual(A []net.IPNet, B []net.IPNet) bool {
	if len(A) != len(B) {
		return false
	}
	m := make(map[string]struct{}, len(A))
	for _, n := range A {
		m[n.String()] = struct{}{}
	}
	for _, n := range B {
		if _, ok := m[n.String()]; !ok {
			return false
		}
	}
	return true
}

// add a new limiter to the chooser. The limiter will replace an existing limiter
// in the list if they have the same whitelist and blacklist.
func (r *rateLimiterChooser) add(l rateLimiterWithSubnet) {
	for i := 0; i < len(r.pool); i++ {
		if subnetListEqual(l.blacklist, r.pool[i].blacklist) && subnetListEqual(l.whitelist, r.pool[i].whitelist) {
			r.pool[i] = l
			return
		}
	}
	r.pool = append(r.pool, l)
}

// ForResource returns the rate limiter for the given resource.
// If the resource does not have an IP address, the first limiter in the pool will be returned. Note that this will only occur if subnet filters are set for resources where that does not make sense, which is an error in the rate limiting configuration.
// If the resource has an IP address, the limiter with the most specific subnet will be returned.
// The blacklist filter takes priority over the whitelist.
// An empty whitelist has the same effect as a whitelist with a single 0.0.0.0/0 subnet (which matches any IP).
// Note for our future selves, not sure how this works with IPv6.
func (r *rateLimiterChooser) ForResource(resource Resource) (Interface, bool) {
	if len(r.pool) == 0 {
		return nil, false
	}
	ip, hasIP := resource.IPAddress()
	if !hasIP {
		return r.pool[0].Interface, true
	}
	bestMatch, bestMatchSubnetSize := -1, 0
outer:
	for idx, limiter := range r.pool {
		// Make sure resource IP does not match blacklist (if any)
		for _, n := range limiter.blacklist {
			if n.Contains(ip) {
				continue outer
			}
		}

		// Make sure resource IP matches the whitelist (if any)
		if len(limiter.whitelist) == 0 && bestMatchSubnetSize == 0 {
			bestMatch = idx
			continue
		}
		for _, n := range limiter.whitelist {
			if n.Contains(ip) {
				size, _ := n.Mask.Size()
				if size > bestMatchSubnetSize {
					bestMatch = idx
					bestMatchSubnetSize = size
				}
				break
			}
		}
	}
	if bestMatch == -1 {
		return nil, false
	}
	return r.pool[bestMatch].Interface, true
}
