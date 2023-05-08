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

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

var BigQueryTableIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"dataset_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"table_id": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type BigQueryTableIamUpdater struct {
	project   string
	datasetId string
	tableId   string
	d         TerraformResourceData
	Config    *transport_tpg.Config
}

func BigQueryTableIamUpdaterProducer(d TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	if v, ok := d.GetOk("dataset_id"); ok {
		values["dataset_id"] = v.(string)
	}

	if v, ok := d.GetOk("table_id"); ok {
		values["table_id"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/datasets/(?P<dataset_id>[^/]+)/tables/(?P<table_id>[^/]+)", "(?P<project>[^/]+)/(?P<dataset_id>[^/]+)/(?P<table_id>[^/]+)", "(?P<dataset_id>[^/]+)/(?P<table_id>[^/]+)", "(?P<table_id>[^/]+)"}, d, config, d.Get("table_id").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &BigQueryTableIamUpdater{
		project:   values["project"],
		datasetId: values["dataset_id"],
		tableId:   values["table_id"],
		d:         d,
		Config:    config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("dataset_id", u.datasetId); err != nil {
		return nil, fmt.Errorf("Error setting dataset_id: %s", err)
	}
	if err := d.Set("table_id", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting table_id: %s", err)
	}

	return u, nil
}

func BigQueryTableIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/datasets/(?P<dataset_id>[^/]+)/tables/(?P<table_id>[^/]+)", "(?P<project>[^/]+)/(?P<dataset_id>[^/]+)/(?P<table_id>[^/]+)", "(?P<dataset_id>[^/]+)/(?P<table_id>[^/]+)", "(?P<table_id>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &BigQueryTableIamUpdater{
		project:   values["project"],
		datasetId: values["dataset_id"],
		tableId:   values["table_id"],
		d:         d,
		Config:    config,
	}
	if err := d.Set("table_id", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting table_id: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *BigQueryTableIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyTableUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}
	obj = map[string]interface{}{
		"options": map[string]interface{}{
			"requestedPolicyVersion": 1,
		},
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(u.Config, "POST", project, url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = tpgresource.Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *BigQueryTableIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	// This is an override of the existing version that might have been set in the resource_iam_member|policy|binding code
	json["version"] = 1
	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyTableUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequestWithTimeout(u.Config, "POST", project, url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *BigQueryTableIamUpdater) qualifyTableUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{BigQueryBasePath}}%s:%s", fmt.Sprintf("projects/%s/datasets/%s/tables/%s", u.project, u.datasetId, u.tableId), methodIdentifier)
	url, err := ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *BigQueryTableIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/datasets/%s/tables/%s", u.project, u.datasetId, u.tableId)
}

func (u *BigQueryTableIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-bigquery-table-%s", u.GetResourceId())
}

func (u *BigQueryTableIamUpdater) DescribeResource() string {
	return fmt.Sprintf("bigquery table %q", u.GetResourceId())
}
