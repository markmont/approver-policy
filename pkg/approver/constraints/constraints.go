/*
Copyright 2021 The cert-manager Authors.

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

package constraints

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/spf13/pflag"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	policyapi "github.com/cert-manager/approver-policy/pkg/apis/policy/v1alpha1"
	"github.com/cert-manager/approver-policy/pkg/approver"
	"github.com/cert-manager/approver-policy/pkg/registry"
)

// Load the Constraints approver.
func init() {
	registry.Shared.Store(Constraints{})
}

// Constraints is a base approver-policy Approver that is responsible for
// ensuring incoming requests satisfy the Constraints defined on
// CertificateRequestPolicies. It is expected that constraints must _always_ be
// registered for all approver-policy builds.
type Constraints struct{}

// Name of Approver is "constraints"
func (c Constraints) Name() string {
	return "constraints"
}

// RegisterFlags is a no-op, constraints doesn't need any flags.
func (c Constraints) RegisterFlags(_ *pflag.FlagSet) {
	return
}

// Prepare is a no-op, constraints doesn't need to prepare anything.
func (c Constraints) Prepare(_ context.Context, _ logr.Logger, _ manager.Manager) error {
	return nil
}

// Ready always returns ready, constraints doesn't have any dependencies to
// block readiness.
func (c Constraints) Ready(_ context.Context, _ *policyapi.CertificateRequestPolicy) (approver.ReconcilerReadyResponse, error) {
	return approver.ReconcilerReadyResponse{Ready: true}, nil
}

// Constraints never needs to manually enqueue policies.
func (c Constraints) EnqueueChan() <-chan string {
	return nil
}
