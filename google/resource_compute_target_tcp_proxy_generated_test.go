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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccComputeTargetTcpProxy_targetTcpProxyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeTargetTcpProxyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeTargetTcpProxy_targetTcpProxyBasicExample(context),
			},
			{
				ResourceName:            "google_compute_target_tcp_proxy.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service"},
			},
		},
	})
}

func testAccComputeTargetTcpProxy_targetTcpProxyBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_target_tcp_proxy" "default" {
  name            = "tf-test-test-proxy%{random_suffix}"
  backend_service = google_compute_backend_service.default.id
}

resource "google_compute_backend_service" "default" {
  name        = "tf-test-backend-service%{random_suffix}"
  protocol    = "TCP"
  timeout_sec = 10

  health_checks = [google_compute_health_check.default.id]
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  timeout_sec        = 1
  check_interval_sec = 1

  tcp_health_check {
    port = "443"
  }
}
`, context)
}

func testAccCheckComputeTargetTcpProxyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_target_tcp_proxy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := acctest.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/targetTcpProxies/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeTargetTcpProxy still exists at %s", url)
			}
		}

		return nil
	}
}
