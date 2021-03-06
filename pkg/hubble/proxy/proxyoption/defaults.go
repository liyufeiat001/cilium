// Copyright 2020 Authors of Cilium
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

package proxyoption

import (
	"fmt"
	"time"

	"github.com/cilium/cilium/pkg/defaults"
	hubbledefaults "github.com/cilium/cilium/pkg/hubble/defaults"
)

// Default is the reference point for default values.
var Default = Options{
	HubbleTarget:  "unix://" + defaults.HubbleSockPath,
	DialTimeout:   5 * time.Second,
	ListenAddress: fmt.Sprintf(":%d", hubbledefaults.ProxyPort),
}
