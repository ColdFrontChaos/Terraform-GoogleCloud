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

func TestAccPrivatecaCertificateTemplateIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/privateca.templateUser",
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
				Config: testAccPrivatecaCertificateTemplateIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_privateca_certificate_template_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s roles/privateca.templateUser", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccPrivatecaCertificateTemplateIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_privateca_certificate_template_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s roles/privateca.templateUser", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPrivatecaCertificateTemplateIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/privateca.templateUser",
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
				Config: testAccPrivatecaCertificateTemplateIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_privateca_certificate_template_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s roles/privateca.templateUser user:admin@hashicorptest.com", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPrivatecaCertificateTemplateIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/privateca.templateUser",
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
				Config: testAccPrivatecaCertificateTemplateIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_privateca_certificate_template_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccPrivatecaCertificateTemplateIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_privateca_certificate_template_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPrivatecaCertificateTemplateIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/privateca.templateUser",
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
				Config: testAccPrivatecaCertificateTemplateIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_privateca_certificate_template_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s roles/privateca.templateUser %s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-template%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPrivatecaCertificateTemplateIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/privateca.templateUser",
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
				Config: testAccPrivatecaCertificateTemplateIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_privateca_certificate_template_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s roles/privateca.templateUser", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_privateca_certificate_template_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s roles/privateca.templateUser %s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-template%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_privateca_certificate_template_iam_binding.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s roles/privateca.templateUser %s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-template%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPrivatecaCertificateTemplateIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/privateca.templateUser",
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
				Config: testAccPrivatecaCertificateTemplateIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_privateca_certificate_template_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s roles/privateca.templateUser user:admin@hashicorptest.com %s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-template%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPrivatecaCertificateTemplateIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/privateca.templateUser",
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
				Config: testAccPrivatecaCertificateTemplateIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_privateca_certificate_template_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s roles/privateca.templateUser user:admin@hashicorptest.com", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_privateca_certificate_template_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s roles/privateca.templateUser user:admin@hashicorptest.com %s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-template%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_privateca_certificate_template_iam_member.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s roles/privateca.templateUser user:admin@hashicorptest.com %s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-template%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPrivatecaCertificateTemplateIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/privateca.templateUser",
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
				Config: testAccPrivatecaCertificateTemplateIamPolicy_withConditionGenerated(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO(SarahFrench) - uncomment once https://github.com/GoogleCloudPlatform/magic-modules/pull/6466 merged
					// resource.TestCheckResourceAttr("data.google_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttr("google_privateca_certificate_template_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", checkGoogleIamPolicy),
				),
			},
			{
				ResourceName:      "google_privateca_certificate_template_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/certificateTemplates/%s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-my-template%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPrivatecaCertificateTemplateIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_template" "default" {
  name = "tf-test-my-template%{random_suffix}"
  location = "us-central1"

  identity_constraints {
    allow_subject_alt_names_passthrough = true
    allow_subject_passthrough           = true

    cel_expression {
      description = "Always true"
      expression  = "true"
      location    = "any.file.anywhere"
      title       = "Sample expression"
    }
  }
}

resource "google_privateca_certificate_template_iam_member" "foo" {
  certificate_template = google_privateca_certificate_template.default.id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccPrivatecaCertificateTemplateIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_template" "default" {
  name = "tf-test-my-template%{random_suffix}"
  location = "us-central1"

  identity_constraints {
    allow_subject_alt_names_passthrough = true
    allow_subject_passthrough           = true

    cel_expression {
      description = "Always true"
      expression  = "true"
      location    = "any.file.anywhere"
      title       = "Sample expression"
    }
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_privateca_certificate_template_iam_policy" "foo" {
  certificate_template = google_privateca_certificate_template.default.id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccPrivatecaCertificateTemplateIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_template" "default" {
  name = "tf-test-my-template%{random_suffix}"
  location = "us-central1"

  identity_constraints {
    allow_subject_alt_names_passthrough = true
    allow_subject_passthrough           = true

    cel_expression {
      description = "Always true"
      expression  = "true"
      location    = "any.file.anywhere"
      title       = "Sample expression"
    }
  }
}

data "google_iam_policy" "foo" {
}

resource "google_privateca_certificate_template_iam_policy" "foo" {
  certificate_template = google_privateca_certificate_template.default.id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccPrivatecaCertificateTemplateIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_template" "default" {
  name = "tf-test-my-template%{random_suffix}"
  location = "us-central1"

  identity_constraints {
    allow_subject_alt_names_passthrough = true
    allow_subject_passthrough           = true

    cel_expression {
      description = "Always true"
      expression  = "true"
      location    = "any.file.anywhere"
      title       = "Sample expression"
    }
  }
}

resource "google_privateca_certificate_template_iam_binding" "foo" {
  certificate_template = google_privateca_certificate_template.default.id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccPrivatecaCertificateTemplateIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_template" "default" {
  name = "tf-test-my-template%{random_suffix}"
  location = "us-central1"

  identity_constraints {
    allow_subject_alt_names_passthrough = true
    allow_subject_passthrough           = true

    cel_expression {
      description = "Always true"
      expression  = "true"
      location    = "any.file.anywhere"
      title       = "Sample expression"
    }
  }
}

resource "google_privateca_certificate_template_iam_binding" "foo" {
  certificate_template = google_privateca_certificate_template.default.id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccPrivatecaCertificateTemplateIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_template" "default" {
  name = "tf-test-my-template%{random_suffix}"
  location = "us-central1"

  identity_constraints {
    allow_subject_alt_names_passthrough = true
    allow_subject_passthrough           = true

    cel_expression {
      description = "Always true"
      expression  = "true"
      location    = "any.file.anywhere"
      title       = "Sample expression"
    }
  }
}

resource "google_privateca_certificate_template_iam_binding" "foo" {
  certificate_template = google_privateca_certificate_template.default.id
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

func testAccPrivatecaCertificateTemplateIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_template" "default" {
  name = "tf-test-my-template%{random_suffix}"
  location = "us-central1"

  identity_constraints {
    allow_subject_alt_names_passthrough = true
    allow_subject_passthrough           = true

    cel_expression {
      description = "Always true"
      expression  = "true"
      location    = "any.file.anywhere"
      title       = "Sample expression"
    }
  }
}

resource "google_privateca_certificate_template_iam_binding" "foo" {
  certificate_template = google_privateca_certificate_template.default.id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_privateca_certificate_template_iam_binding" "foo2" {
  certificate_template = google_privateca_certificate_template.default.id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_privateca_certificate_template_iam_binding" "foo3" {
  certificate_template = google_privateca_certificate_template.default.id
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

func testAccPrivatecaCertificateTemplateIamMember_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_template" "default" {
  name = "tf-test-my-template%{random_suffix}"
  location = "us-central1"

  identity_constraints {
    allow_subject_alt_names_passthrough = true
    allow_subject_passthrough           = true

    cel_expression {
      description = "Always true"
      expression  = "true"
      location    = "any.file.anywhere"
      title       = "Sample expression"
    }
  }
}

resource "google_privateca_certificate_template_iam_member" "foo" {
  certificate_template = google_privateca_certificate_template.default.id
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

func testAccPrivatecaCertificateTemplateIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_template" "default" {
  name = "tf-test-my-template%{random_suffix}"
  location = "us-central1"

  identity_constraints {
    allow_subject_alt_names_passthrough = true
    allow_subject_passthrough           = true

    cel_expression {
      description = "Always true"
      expression  = "true"
      location    = "any.file.anywhere"
      title       = "Sample expression"
    }
  }
}

resource "google_privateca_certificate_template_iam_member" "foo" {
  certificate_template = google_privateca_certificate_template.default.id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_privateca_certificate_template_iam_member" "foo2" {
  certificate_template = google_privateca_certificate_template.default.id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_privateca_certificate_template_iam_member" "foo3" {
  certificate_template = google_privateca_certificate_template.default.id
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

func testAccPrivatecaCertificateTemplateIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_privateca_certificate_template" "default" {
  name = "tf-test-my-template%{random_suffix}"
  location = "us-central1"

  identity_constraints {
    allow_subject_alt_names_passthrough = true
    allow_subject_passthrough           = true

    cel_expression {
      description = "Always true"
      expression  = "true"
      location    = "any.file.anywhere"
      title       = "Sample expression"
    }
  }
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

resource "google_privateca_certificate_template_iam_policy" "foo" {
  certificate_template = google_privateca_certificate_template.default.id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
