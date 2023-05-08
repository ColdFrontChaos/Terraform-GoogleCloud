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
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceActiveDirectoryDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceActiveDirectoryDomainCreate,
		Read:   resourceActiveDirectoryDomainRead,
		Update: resourceActiveDirectoryDomainUpdate,
		Delete: resourceActiveDirectoryDomainDelete,

		Importer: &schema.ResourceImporter{
			State: resourceActiveDirectoryDomainImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateADDomainName(),
				Description: `The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.`,
			},
			"locations": {
				Type:     schema.TypeList,
				Required: true,
				Description: `Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"reserved_ip_range": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks`,
			},
			"admin": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The name of delegated administrator account used to perform Active Directory operations.
If not specified, setupadmin will be used.`,
				Default: "setupadmin",
			},
			"authorized_networks": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: `The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
If CIDR subnets overlap between networks, domain creation will fail.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Resource labels that can contain user-provided metadata`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"fqdn": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The fully-qualified domain name of the exposed domain used by clients to connect to the service.
Similar to what would be chosen for an Active Directory set up on an internal network.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The unique name of the domain using the format: 'projects/{project}/locations/global/domains/{domainName}'.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceActiveDirectoryDomainCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandActiveDirectoryDomainLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	authorizedNetworksProp, err := expandActiveDirectoryDomainAuthorizedNetworks(d.Get("authorized_networks"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authorized_networks"); !isEmptyValue(reflect.ValueOf(authorizedNetworksProp)) && (ok || !reflect.DeepEqual(v, authorizedNetworksProp)) {
		obj["authorizedNetworks"] = authorizedNetworksProp
	}
	reservedIpRangeProp, err := expandActiveDirectoryDomainReservedIpRange(d.Get("reserved_ip_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("reserved_ip_range"); !isEmptyValue(reflect.ValueOf(reservedIpRangeProp)) && (ok || !reflect.DeepEqual(v, reservedIpRangeProp)) {
		obj["reservedIpRange"] = reservedIpRangeProp
	}
	locationsProp, err := expandActiveDirectoryDomainLocations(d.Get("locations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("locations"); !isEmptyValue(reflect.ValueOf(locationsProp)) && (ok || !reflect.DeepEqual(v, locationsProp)) {
		obj["locations"] = locationsProp
	}
	adminProp, err := expandActiveDirectoryDomainAdmin(d.Get("admin"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("admin"); !isEmptyValue(reflect.ValueOf(adminProp)) && (ok || !reflect.DeepEqual(v, adminProp)) {
		obj["admin"] = adminProp
	}

	url, err := ReplaceVars(d, config, "{{ActiveDirectoryBasePath}}projects/{{project}}/locations/global/domains?domainName={{domain_name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Domain: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Domain: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Domain: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = ActiveDirectoryOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Domain", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Domain: %s", err)
	}

	if err := d.Set("name", flattenActiveDirectoryDomainName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Domain %q: %#v", d.Id(), res)

	return resourceActiveDirectoryDomainRead(d, meta)
}

func resourceActiveDirectoryDomainRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ActiveDirectoryBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Domain: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ActiveDirectoryDomain %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Domain: %s", err)
	}

	if err := d.Set("name", flattenActiveDirectoryDomainName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Domain: %s", err)
	}
	if err := d.Set("labels", flattenActiveDirectoryDomainLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Domain: %s", err)
	}
	if err := d.Set("authorized_networks", flattenActiveDirectoryDomainAuthorizedNetworks(res["authorizedNetworks"], d, config)); err != nil {
		return fmt.Errorf("Error reading Domain: %s", err)
	}
	if err := d.Set("reserved_ip_range", flattenActiveDirectoryDomainReservedIpRange(res["reservedIpRange"], d, config)); err != nil {
		return fmt.Errorf("Error reading Domain: %s", err)
	}
	if err := d.Set("locations", flattenActiveDirectoryDomainLocations(res["locations"], d, config)); err != nil {
		return fmt.Errorf("Error reading Domain: %s", err)
	}
	if err := d.Set("admin", flattenActiveDirectoryDomainAdmin(res["admin"], d, config)); err != nil {
		return fmt.Errorf("Error reading Domain: %s", err)
	}
	if err := d.Set("fqdn", flattenActiveDirectoryDomainFqdn(res["fqdn"], d, config)); err != nil {
		return fmt.Errorf("Error reading Domain: %s", err)
	}

	return nil
}

func resourceActiveDirectoryDomainUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Domain: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandActiveDirectoryDomainLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	authorizedNetworksProp, err := expandActiveDirectoryDomainAuthorizedNetworks(d.Get("authorized_networks"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authorized_networks"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, authorizedNetworksProp)) {
		obj["authorizedNetworks"] = authorizedNetworksProp
	}
	locationsProp, err := expandActiveDirectoryDomainLocations(d.Get("locations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("locations"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, locationsProp)) {
		obj["locations"] = locationsProp
	}

	url, err := ReplaceVars(d, config, "{{ActiveDirectoryBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Domain %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("authorized_networks") {
		updateMask = append(updateMask, "authorizedNetworks")
	}

	if d.HasChange("locations") {
		updateMask = append(updateMask, "locations")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Domain %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Domain %q: %#v", d.Id(), res)
	}

	err = ActiveDirectoryOperationWaitTime(
		config, res, project, "Updating Domain", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceActiveDirectoryDomainRead(d, meta)
}

func resourceActiveDirectoryDomainDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Domain: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{ActiveDirectoryBasePath}}projects/{{project}}/locations/global/domains/{{domain_name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Domain %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Domain")
	}

	err = ActiveDirectoryOperationWaitTime(
		config, res, project, "Deleting Domain", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Domain %q: %#v", d.Id(), res)
	return nil
}

func resourceActiveDirectoryDomainImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := ParseImportId([]string{"(?P<project>[^ ]+) (?P<name>[^ ]+)", "(?P<name>[^ ]+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenActiveDirectoryDomainName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenActiveDirectoryDomainLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenActiveDirectoryDomainAuthorizedNetworks(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenActiveDirectoryDomainReservedIpRange(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenActiveDirectoryDomainLocations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenActiveDirectoryDomainAdmin(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenActiveDirectoryDomainFqdn(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandActiveDirectoryDomainLabels(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandActiveDirectoryDomainAuthorizedNetworks(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandActiveDirectoryDomainReservedIpRange(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandActiveDirectoryDomainLocations(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandActiveDirectoryDomainAdmin(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
