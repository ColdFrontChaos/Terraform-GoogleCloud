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

var SecurityCenterSourceIamSchema = map[string]*schema.Schema{
	"organization": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"source": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type SecurityCenterSourceIamUpdater struct {
	organization string
	source       string
	d            TerraformResourceData
	Config       *transport_tpg.Config
}

func SecurityCenterSourceIamUpdaterProducer(d TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	if v, ok := d.GetOk("organization"); ok {
		values["organization"] = v.(string)
	}

	if v, ok := d.GetOk("source"); ok {
		values["source"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"organizations/(?P<organization>[^/]+)/sources/(?P<source>[^/]+)", "(?P<organization>[^/]+)/(?P<source>[^/]+)", "(?P<source>[^/]+)"}, d, config, d.Get("source").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &SecurityCenterSourceIamUpdater{
		organization: values["organization"],
		source:       values["source"],
		d:            d,
		Config:       config,
	}

	if err := d.Set("organization", u.organization); err != nil {
		return nil, fmt.Errorf("Error setting organization: %s", err)
	}
	if err := d.Set("source", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting source: %s", err)
	}

	return u, nil
}

func SecurityCenterSourceIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	m, err := getImportIdQualifiers([]string{"organizations/(?P<organization>[^/]+)/sources/(?P<source>[^/]+)", "(?P<organization>[^/]+)/(?P<source>[^/]+)", "(?P<source>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &SecurityCenterSourceIamUpdater{
		organization: values["organization"],
		source:       values["source"],
		d:            d,
		Config:       config,
	}
	if err := d.Set("source", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting source: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *SecurityCenterSourceIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifySourceUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	var obj map[string]interface{}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(u.Config, "POST", "", url, userAgent, obj)
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

func (u *SecurityCenterSourceIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifySourceUrl("setIamPolicy")
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequestWithTimeout(u.Config, "POST", "", url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *SecurityCenterSourceIamUpdater) qualifySourceUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{SecurityCenterBasePath}}%s:%s", fmt.Sprintf("organizations/%s/sources/%s", u.organization, u.source), methodIdentifier)
	url, err := ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *SecurityCenterSourceIamUpdater) GetResourceId() string {
	return fmt.Sprintf("organizations/%s/sources/%s", u.organization, u.source)
}

func (u *SecurityCenterSourceIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-securitycenter-source-%s", u.GetResourceId())
}

func (u *SecurityCenterSourceIamUpdater) DescribeResource() string {
	return fmt.Sprintf("securitycenter source %q", u.GetResourceId())
}
