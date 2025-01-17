// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	application "github.com/upbound/provider-azure/internal/controller/iotcentral/application"
	applicationnetworkruleset "github.com/upbound/provider-azure/internal/controller/iotcentral/applicationnetworkruleset"
)

// Setup_iotcentral creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_iotcentral(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		applicationnetworkruleset.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
