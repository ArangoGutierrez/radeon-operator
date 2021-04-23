/*
Copyright Eduardo Arango 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"errors"

	riv1 "github.com/ArangoGutierrez/radeon-operator/api/v1alpha1"
)

// RI holds the needed information to watch from the Controller
// short for Radeon Instance
type RI struct {
	resources []Resources
	controls  []controlFunc
	rec       *RadeonInstanceReconciler
	ins       *riv1.RadeonInstance
	idx       int
}

func (n *RI) addState(path string) {
	res, ctrl := addResourcesControls(path)
	n.controls = append(n.controls, ctrl)
	n.resources = append(n.resources, res)
}

func (n *RI) init(
	r *RadeonInstanceReconciler,
	i *riv1.RadeonInstance,
) {
	n.rec = r
	n.ins = i
	n.idx = 0
	if len(n.controls) == 0 {
		n.addState("/opt/device-plugin")
	}
}

func (n *RI) step() error {
	for _, fs := range n.controls[n.idx] {
		stat, err := fs(*n)
		if err != nil {
			return err
		}
		if stat != Ready {
			return errors.New("ResourceNotReady")
		}
	}
	n.idx = n.idx + 1
	return nil
}

func (n *RI) last() bool {
	return n.idx == len(n.controls)
}
