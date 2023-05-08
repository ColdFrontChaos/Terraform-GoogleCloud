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

func ResourceAccessContextManagerAuthorizedOrgsDesc() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessContextManagerAuthorizedOrgsDescCreate,
		Read:   resourceAccessContextManagerAuthorizedOrgsDescRead,
		Update: resourceAccessContextManagerAuthorizedOrgsDescUpdate,
		Delete: resourceAccessContextManagerAuthorizedOrgsDescDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessContextManagerAuthorizedOrgsDescImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Resource name for the 'AuthorizedOrgsDesc'. Format:
'accessPolicies/{access_policy}/authorizedOrgsDescs/{authorized_orgs_desc}'.
The 'authorized_orgs_desc' component must begin with a letter, followed by
alphanumeric characters or '_'.
After you create an 'AuthorizedOrgsDesc', you cannot change its 'name'.`,
			},
			"parent": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Required. Resource name for the access policy which owns this 'AuthorizedOrgsDesc'.`,
			},
			"asset_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ASSET_TYPE_DEVICE", "ASSET_TYPE_CREDENTIAL_STRENGTH", ""}),
				Description: `The type of entities that need to use the authorization relationship during
evaluation, such as a device. Valid values are "ASSET_TYPE_DEVICE" and
"ASSET_TYPE_CREDENTIAL_STRENGTH". Possible values: ["ASSET_TYPE_DEVICE", "ASSET_TYPE_CREDENTIAL_STRENGTH"]`,
			},
			"authorization_direction": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"AUTHORIZATION_DIRECTION_TO", "AUTHORIZATION_DIRECTION_FROM", ""}),
				Description: `The direction of the authorization relationship between this organization
and the organizations listed in the "orgs" field. The valid values for this
field include the following:

AUTHORIZATION_DIRECTION_FROM: Allows this organization to evaluate traffic
in the organizations listed in the 'orgs' field.

AUTHORIZATION_DIRECTION_TO: Allows the organizations listed in the 'orgs'
field to evaluate the traffic in this organization.

For the authorization relationship to take effect, all of the organizations
must authorize and specify the appropriate relationship direction. For
example, if organization A authorized organization B and C to evaluate its
traffic, by specifying "AUTHORIZATION_DIRECTION_TO" as the authorization
direction, organizations B and C must specify
"AUTHORIZATION_DIRECTION_FROM" as the authorization direction in their
"AuthorizedOrgsDesc" resource. Possible values: ["AUTHORIZATION_DIRECTION_TO", "AUTHORIZATION_DIRECTION_FROM"]`,
			},
			"authorization_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"AUTHORIZATION_TYPE_TRUST", ""}),
				Description:  `A granular control type for authorization levels. Valid value is "AUTHORIZATION_TYPE_TRUST". Possible values: ["AUTHORIZATION_TYPE_TRUST"]`,
			},
			"orgs": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `The list of organization ids in this AuthorizedOrgsDesc.
Format: 'organizations/<org_number>'
Example: 'organizations/123456'`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the AuthorizedOrgsDesc was created in UTC.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the AuthorizedOrgsDesc was updated in UTC.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceAccessContextManagerAuthorizedOrgsDescCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	parentProp, err := expandAccessContextManagerAuthorizedOrgsDescParent(d.Get("parent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	nameProp, err := expandAccessContextManagerAuthorizedOrgsDescName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	orgsProp, err := expandAccessContextManagerAuthorizedOrgsDescOrgs(d.Get("orgs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("orgs"); !isEmptyValue(reflect.ValueOf(orgsProp)) && (ok || !reflect.DeepEqual(v, orgsProp)) {
		obj["orgs"] = orgsProp
	}
	assetTypeProp, err := expandAccessContextManagerAuthorizedOrgsDescAssetType(d.Get("asset_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("asset_type"); !isEmptyValue(reflect.ValueOf(assetTypeProp)) && (ok || !reflect.DeepEqual(v, assetTypeProp)) {
		obj["assetType"] = assetTypeProp
	}
	authorizationDirectionProp, err := expandAccessContextManagerAuthorizedOrgsDescAuthorizationDirection(d.Get("authorization_direction"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authorization_direction"); !isEmptyValue(reflect.ValueOf(authorizationDirectionProp)) && (ok || !reflect.DeepEqual(v, authorizationDirectionProp)) {
		obj["authorizationDirection"] = authorizationDirectionProp
	}
	authorizationTypeProp, err := expandAccessContextManagerAuthorizedOrgsDescAuthorizationType(d.Get("authorization_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authorization_type"); !isEmptyValue(reflect.ValueOf(authorizationTypeProp)) && (ok || !reflect.DeepEqual(v, authorizationTypeProp)) {
		obj["authorizationType"] = authorizationTypeProp
	}

	obj, err = resourceAccessContextManagerAuthorizedOrgsDescEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{AccessContextManagerBasePath}}{{parent}}/authorizedOrgsDescs")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AuthorizedOrgsDesc: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating AuthorizedOrgsDesc: %s", err)
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
	err = AccessContextManagerOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating AuthorizedOrgsDesc", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create AuthorizedOrgsDesc: %s", err)
	}

	if err := d.Set("name", flattenAccessContextManagerAuthorizedOrgsDescName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// This is useful if the resource in question doesn't have a perfectly consistent API
	// That is, the Operation for Create might return before the Get operation shows the
	// completed state of the resource.
	time.Sleep(2 * time.Minute)

	log.Printf("[DEBUG] Finished creating AuthorizedOrgsDesc %q: %#v", d.Id(), res)

	return resourceAccessContextManagerAuthorizedOrgsDescRead(d, meta)
}

func resourceAccessContextManagerAuthorizedOrgsDescRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{AccessContextManagerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("AccessContextManagerAuthorizedOrgsDesc %q", d.Id()))
	}

	if err := d.Set("create_time", flattenAccessContextManagerAuthorizedOrgsDescCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading AuthorizedOrgsDesc: %s", err)
	}
	if err := d.Set("update_time", flattenAccessContextManagerAuthorizedOrgsDescUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading AuthorizedOrgsDesc: %s", err)
	}
	if err := d.Set("name", flattenAccessContextManagerAuthorizedOrgsDescName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading AuthorizedOrgsDesc: %s", err)
	}
	if err := d.Set("orgs", flattenAccessContextManagerAuthorizedOrgsDescOrgs(res["orgs"], d, config)); err != nil {
		return fmt.Errorf("Error reading AuthorizedOrgsDesc: %s", err)
	}
	if err := d.Set("asset_type", flattenAccessContextManagerAuthorizedOrgsDescAssetType(res["assetType"], d, config)); err != nil {
		return fmt.Errorf("Error reading AuthorizedOrgsDesc: %s", err)
	}
	if err := d.Set("authorization_direction", flattenAccessContextManagerAuthorizedOrgsDescAuthorizationDirection(res["authorizationDirection"], d, config)); err != nil {
		return fmt.Errorf("Error reading AuthorizedOrgsDesc: %s", err)
	}
	if err := d.Set("authorization_type", flattenAccessContextManagerAuthorizedOrgsDescAuthorizationType(res["authorizationType"], d, config)); err != nil {
		return fmt.Errorf("Error reading AuthorizedOrgsDesc: %s", err)
	}

	return nil
}

func resourceAccessContextManagerAuthorizedOrgsDescUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	orgsProp, err := expandAccessContextManagerAuthorizedOrgsDescOrgs(d.Get("orgs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("orgs"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, orgsProp)) {
		obj["orgs"] = orgsProp
	}

	obj, err = resourceAccessContextManagerAuthorizedOrgsDescEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{AccessContextManagerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AuthorizedOrgsDesc %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("orgs") {
		updateMask = append(updateMask, "orgs")
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
		return fmt.Errorf("Error updating AuthorizedOrgsDesc %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating AuthorizedOrgsDesc %q: %#v", d.Id(), res)
	}

	err = AccessContextManagerOperationWaitTime(
		config, res, "Updating AuthorizedOrgsDesc", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceAccessContextManagerAuthorizedOrgsDescRead(d, meta)
}

func resourceAccessContextManagerAuthorizedOrgsDescDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := ReplaceVars(d, config, "{{AccessContextManagerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting AuthorizedOrgsDesc %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "AuthorizedOrgsDesc")
	}

	err = AccessContextManagerOperationWaitTime(
		config, res, "Deleting AuthorizedOrgsDesc", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting AuthorizedOrgsDesc %q: %#v", d.Id(), res)
	return nil
}

func resourceAccessContextManagerAuthorizedOrgsDescImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}
	stringParts := strings.Split(d.Get("name").(string), "/")
	if len(stringParts) < 2 {
		return nil, fmt.Errorf("Error parsing parent name. Should be in form accessPolicies/{{policy_id}}/authorizedOrgsDescs/{{short_name}}")
	}
	if err := d.Set("parent", fmt.Sprintf("%s/%s", stringParts[0], stringParts[1])); err != nil {
		return nil, fmt.Errorf("Error setting parent, %s", err)
	}
	return []*schema.ResourceData{d}, nil
}

func flattenAccessContextManagerAuthorizedOrgsDescCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessContextManagerAuthorizedOrgsDescUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessContextManagerAuthorizedOrgsDescName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessContextManagerAuthorizedOrgsDescOrgs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessContextManagerAuthorizedOrgsDescAssetType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessContextManagerAuthorizedOrgsDescAuthorizationDirection(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAccessContextManagerAuthorizedOrgsDescAuthorizationType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandAccessContextManagerAuthorizedOrgsDescParent(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAuthorizedOrgsDescName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAuthorizedOrgsDescOrgs(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAuthorizedOrgsDescAssetType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAuthorizedOrgsDescAuthorizationDirection(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAuthorizedOrgsDescAuthorizationType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceAccessContextManagerAuthorizedOrgsDescEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "parent")
	return obj, nil
}
