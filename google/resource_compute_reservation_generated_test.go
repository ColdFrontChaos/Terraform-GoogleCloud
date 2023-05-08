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

func TestAccComputeReservation_reservationBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeReservationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeReservation_reservationBasicExample(context),
			},
			{
				ResourceName:            "google_compute_reservation.gce_reservation",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"share_settings", "zone"},
			},
		},
	})
}

func testAccComputeReservation_reservationBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_reservation" "gce_reservation" {
  name = "tf-test-gce-reservation%{random_suffix}"
  zone = "us-central1-a"

  specific_reservation {
    count = 1
    instance_properties {
      min_cpu_platform = "Intel Cascade Lake"
      machine_type     = "n2-standard-2"
    }
  }
}
`, context)
}

func TestAccComputeReservation_sharedReservationBasicExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"project":         acctest.GetTestProjectFromEnv(),
		"org_id":          acctest.GetTestOrgFromEnv(t),
		"billing_account": acctest.GetTestBillingAccountFromEnv(t),
		"random_suffix":   RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeReservationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeReservation_sharedReservationBasicExample(context),
			},
			{
				ResourceName:            "google_compute_reservation.gce_reservation",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"share_settings", "zone"},
			},
		},
	})
}

func testAccComputeReservation_sharedReservationBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "owner_project" {
  project_id      = "tf-test%{random_suffix}"
  name            = "tf-test%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}


resource "google_project_service" "compute" {
  project = google_project.owner_project.project_id
  service = "compute.googleapis.com"
  disable_on_destroy = false
}

resource "google_project" "guest_project" {
  project_id      = "tf-test-2%{random_suffix}"
  name            = "tf-test-2%{random_suffix}"
  org_id          = "%{org_id}"
}

resource "google_organization_policy" "shared_reservation_org_policy" {
  org_id     = "%{org_id}"
  constraint = "constraints/compute.sharedReservationsOwnerProjects"
  list_policy {
    allow {
      values = ["projects/${google_project.owner_project.number}"]
    }
  }
}

resource "google_compute_reservation" "gce_reservation" {
  project = google_project.owner_project.project_id
  name = "tf-test-gce-shared-reservation%{random_suffix}"
  zone = "us-central1-a"

  specific_reservation {
    count = 1
    instance_properties {
      min_cpu_platform = "Intel Cascade Lake"
      machine_type     = "n2-standard-2"
    }
  }
  share_settings {
    share_type = "SPECIFIC_PROJECTS"
    project_map {
      id = google_project.guest_project.project_id
      project_id = google_project.guest_project.project_id
    }
  }
  depends_on = [google_organization_policy.shared_reservation_org_policy,google_project_service.compute]
}
`, context)
}

func testAccCheckComputeReservationDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_reservation" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := acctest.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/reservations/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeReservation still exists at %s", url)
			}
		}

		return nil
	}
}
