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

func TestAccLoggingLogView_loggingLogViewBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       acctest.GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckLoggingLogViewDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccLoggingLogView_loggingLogViewBasicExample(context),
			},
			{
				ResourceName:            "google_logging_log_view.logging_log_view",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "parent", "location", "bucket"},
			},
		},
	})
}

func testAccLoggingLogView_loggingLogViewBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_logging_project_bucket_config" "logging_log_view" {
    project        = "%{project}"
    location       = "global"
    retention_days = 30
    bucket_id      = "_Default"
}

resource "google_logging_log_view" "logging_log_view" {
  name        = "tf-test-my-view%{random_suffix}"
  bucket      = google_logging_project_bucket_config.logging_log_view.id
  description = "A logging view configured with Terraform"
  filter      = "SOURCE(\"projects/myproject\") AND resource.type = \"gce_instance\" AND LOG_ID(\"stdout\")"
}
`, context)
}

func TestAccLoggingLogView_loggingLogViewLongNameExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       acctest.GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckLoggingLogViewDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccLoggingLogView_loggingLogViewLongNameExample(context),
			},
			{
				ResourceName:            "google_logging_log_view.logging_log_view",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "parent", "location", "bucket"},
			},
		},
	})
}

func testAccLoggingLogView_loggingLogViewLongNameExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_logging_project_bucket_config" "logging_log_view" {
    project        = "%{project}"
    location       = "global"
    retention_days = 30
    bucket_id      = "_Default"
}

resource "google_logging_log_view" "logging_log_view" {
  name        = "projects/%{project}/locations/global/buckets/_Default/views/tf-test-view%{random_suffix}"
  bucket      = google_logging_project_bucket_config.logging_log_view.id
  description = "A logging view configured with Terraform"
  filter      = "SOURCE(\"projects/myproject\") AND resource.type = \"gce_instance\" AND LOG_ID(\"stdout\")"
}
`, context)
}

func testAccCheckLoggingLogViewDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_logging_log_view" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := acctest.ReplaceVarsForTest(config, rs, "{{LoggingBasePath}}{{parent}}/locations/{{location}}/buckets/{{bucket}}/views/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("LoggingLogView still exists at %s", url)
			}
		}

		return nil
	}
}
