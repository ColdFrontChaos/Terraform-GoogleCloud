// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"time"

	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

type IAMWorkforcePoolOperationWaiter struct {
	Config    *transport_tpg.Config
	UserAgent string
	CommonOperationWaiter
}

func (w *IAMWorkforcePoolOperationWaiter) QueryOp() (interface{}, error) {
	if w == nil {
		return nil, fmt.Errorf("Cannot query operation, it's unset or nil.")
	}
	// Returns the proper get.
	url := fmt.Sprintf("%s%s", w.Config.IAMWorkforcePoolBasePath, w.CommonOperationWaiter.Op.Name)

	return transport_tpg.SendRequest(w.Config, "GET", "", url, w.UserAgent, nil)
}

func createIAMWorkforcePoolWaiter(config *transport_tpg.Config, op map[string]interface{}, activity, userAgent string) (*IAMWorkforcePoolOperationWaiter, error) {
	w := &IAMWorkforcePoolOperationWaiter{
		Config:    config,
		UserAgent: userAgent,
	}
	if err := w.CommonOperationWaiter.SetOp(op); err != nil {
		return nil, err
	}
	return w, nil
}

func IAMWorkforcePoolOperationWaitTime(config *transport_tpg.Config, op map[string]interface{}, activity, userAgent string, timeout time.Duration) error {
	if val, ok := op["name"]; !ok || val == "" {
		// This was a synchronous call - there is no operation to wait for.
		return nil
	}
	w, err := createIAMWorkforcePoolWaiter(config, op, activity, userAgent)
	if err != nil {
		// If w is nil, the op was synchronous.
		return err
	}
	return OperationWait(w, activity, timeout, config.PollInterval)
}
