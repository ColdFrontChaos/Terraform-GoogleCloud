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

	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccComputeSubnetworkIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.networkUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetworkIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s roles/compute.networkUser", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-test-subnetwork%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccComputeSubnetworkIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s roles/compute.networkUser", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-test-subnetwork%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeSubnetworkIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.networkUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccComputeSubnetworkIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s roles/compute.networkUser user:admin@hashicorptest.com", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-test-subnetwork%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeSubnetworkIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.networkUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetworkIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-test-subnetwork%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeSubnetworkIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-test-subnetwork%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeSubnetworkIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.networkUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetworkIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s roles/compute.networkUser %s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-test-subnetwork%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeSubnetworkIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.networkUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetworkIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s roles/compute.networkUser", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-test-subnetwork%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s roles/compute.networkUser %s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-test-subnetwork%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_binding.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s roles/compute.networkUser %s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-test-subnetwork%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeSubnetworkIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.networkUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetworkIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s roles/compute.networkUser user:admin@hashicorptest.com %s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-test-subnetwork%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeSubnetworkIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.networkUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetworkIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s roles/compute.networkUser user:admin@hashicorptest.com", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-test-subnetwork%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s roles/compute.networkUser user:admin@hashicorptest.com %s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-test-subnetwork%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_member.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s roles/compute.networkUser user:admin@hashicorptest.com %s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-test-subnetwork%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeSubnetworkIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/compute.networkUser",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	// Test should have 2 bindings: one with a description and one without. Any < chars are converted to a unicode character by the API.
	expectedPolicyData := Nprintf(`{"bindings":[{"condition":{"description":"%{condition_desc}","expression":"%{condition_expr}","title":"%{condition_title}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"},{"condition":{"expression":"%{condition_expr}","title":"%{condition_title}-no-description"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"}]}`, context)
	expectedPolicyData = strings.Replace(expectedPolicyData, "<", "\\u003c", -1)

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetworkIamPolicy_withConditionGenerated(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO(SarahFrench) - uncomment once https://github.com/GoogleCloudPlatform/magic-modules/pull/6466 merged
					// resource.TestCheckResourceAttr("data.google_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttr("google_compute_subnetwork_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", checkGoogleIamPolicy),
				),
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-test-subnetwork%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeSubnetworkIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "tf-test-test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-test-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork_iam_member" "foo" {
  project = google_compute_subnetwork.network-with-private-secondary-ip-ranges.project
  region = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  subnetwork = google_compute_subnetwork.network-with-private-secondary-ip-ranges.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccComputeSubnetworkIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "tf-test-test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-test-network%{random_suffix}"
  auto_create_subnetworks = false
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_compute_subnetwork_iam_policy" "foo" {
  project = google_compute_subnetwork.network-with-private-secondary-ip-ranges.project
  region = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  subnetwork = google_compute_subnetwork.network-with-private-secondary-ip-ranges.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeSubnetworkIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "tf-test-test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-test-network%{random_suffix}"
  auto_create_subnetworks = false
}

data "google_iam_policy" "foo" {
}

resource "google_compute_subnetwork_iam_policy" "foo" {
  project = google_compute_subnetwork.network-with-private-secondary-ip-ranges.project
  region = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  subnetwork = google_compute_subnetwork.network-with-private-secondary-ip-ranges.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeSubnetworkIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "tf-test-test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-test-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork_iam_binding" "foo" {
  project = google_compute_subnetwork.network-with-private-secondary-ip-ranges.project
  region = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  subnetwork = google_compute_subnetwork.network-with-private-secondary-ip-ranges.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccComputeSubnetworkIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "tf-test-test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-test-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork_iam_binding" "foo" {
  project = google_compute_subnetwork.network-with-private-secondary-ip-ranges.project
  region = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  subnetwork = google_compute_subnetwork.network-with-private-secondary-ip-ranges.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccComputeSubnetworkIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "tf-test-test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-test-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork_iam_binding" "foo" {
  project = google_compute_subnetwork.network-with-private-secondary-ip-ranges.project
  region = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  subnetwork = google_compute_subnetwork.network-with-private-secondary-ip-ranges.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccComputeSubnetworkIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "tf-test-test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-test-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork_iam_binding" "foo" {
  project = google_compute_subnetwork.network-with-private-secondary-ip-ranges.project
  region = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  subnetwork = google_compute_subnetwork.network-with-private-secondary-ip-ranges.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_compute_subnetwork_iam_binding" "foo2" {
  project = google_compute_subnetwork.network-with-private-secondary-ip-ranges.project
  region = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  subnetwork = google_compute_subnetwork.network-with-private-secondary-ip-ranges.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_compute_subnetwork_iam_binding" "foo3" {
  project = google_compute_subnetwork.network-with-private-secondary-ip-ranges.project
  region = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  subnetwork = google_compute_subnetwork.network-with-private-secondary-ip-ranges.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccComputeSubnetworkIamMember_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "tf-test-test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-test-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork_iam_member" "foo" {
  project = google_compute_subnetwork.network-with-private-secondary-ip-ranges.project
  region = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  subnetwork = google_compute_subnetwork.network-with-private-secondary-ip-ranges.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccComputeSubnetworkIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "tf-test-test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-test-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork_iam_member" "foo" {
  project = google_compute_subnetwork.network-with-private-secondary-ip-ranges.project
  region = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  subnetwork = google_compute_subnetwork.network-with-private-secondary-ip-ranges.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_compute_subnetwork_iam_member" "foo2" {
  project = google_compute_subnetwork.network-with-private-secondary-ip-ranges.project
  region = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  subnetwork = google_compute_subnetwork.network-with-private-secondary-ip-ranges.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_compute_subnetwork_iam_member" "foo3" {
  project = google_compute_subnetwork.network-with-private-secondary-ip-ranges.project
  region = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  subnetwork = google_compute_subnetwork.network-with-private-secondary-ip-ranges.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccComputeSubnetworkIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "tf-test-test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "tf-test-test-network%{random_suffix}"
  auto_create_subnetworks = false
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      # Check that lack of description doesn't cause any issues
      # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
      title       = "%{condition_title_no_desc}"
      expression  = "%{condition_expr_no_desc}"
    }
  }
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "%{condition_desc}"
      expression  = "%{condition_expr}"
    }
  }
}

resource "google_compute_subnetwork_iam_policy" "foo" {
  project = google_compute_subnetwork.network-with-private-secondary-ip-ranges.project
  region = google_compute_subnetwork.network-with-private-secondary-ip-ranges.region
  subnetwork = google_compute_subnetwork.network-with-private-secondary-ip-ranges.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
