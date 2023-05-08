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

func ResourceCloudIdentityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudIdentityGroupCreate,
		Read:   resourceCloudIdentityGroupRead,
		Update: resourceCloudIdentityGroupUpdate,
		Delete: resourceCloudIdentityGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudIdentityGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"group_key": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: `EntityKey of the Group.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `The ID of the entity.

For Google-managed entities, the id must be the email address of an existing
group or user.

For external-identity-mapped entities, the id must be a string conforming
to the Identity Source's requirements.

Must be unique within a namespace.`,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `The namespace in which the entity exists.

If not specified, the EntityKey represents a Google-managed entity
such as a Google user or a Google Group.

If specified, the EntityKey represents an external-identity-mapped group.
The namespace must correspond to an identity source created in Admin Console
and must be in the form of 'identitysources/{identity_source_id}'.`,
						},
					},
				},
			},
			"labels": {
				Type:     schema.TypeMap,
				Required: true,
				Description: `One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value.

Google Groups are the default type of group and have a label with a key of cloudidentity.googleapis.com/groups.discussion_forum and an empty value.

Existing Google Groups can have an additional label with a key of cloudidentity.googleapis.com/groups.security and an empty value added to them. This is an immutable change and the security label cannot be removed once added.

Dynamic groups have a label with a key of cloudidentity.googleapis.com/groups.dynamic.

Identity-mapped groups for Cloud Search have a label with a key of system/groups/external and an empty value.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"parent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The resource name of the entity under which this Group resides in the
Cloud Identity resource hierarchy.

Must be of the form identitysources/{identity_source_id} for external-identity-mapped
groups or customers/{customer_id} for Google Groups.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `An extended description to help users determine the purpose of a Group.
Must not be longer than 4,096 characters.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The display name of the Group.`,
			},
			"initial_group_config": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"INITIAL_GROUP_CONFIG_UNSPECIFIED", "WITH_INITIAL_OWNER", "EMPTY", ""}),
				Description: `The initial configuration options for creating a Group.

See the
[API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/create#initialgroupconfig)
for possible values. Default value: "EMPTY" Possible values: ["INITIAL_GROUP_CONFIG_UNSPECIFIED", "WITH_INITIAL_OWNER", "EMPTY"]`,
				Default: "EMPTY",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the Group was created.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Resource name of the Group in the format: groups/{group_id}, where group_id
is the unique ID assigned to the Group.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the Group was last updated.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceCloudIdentityGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	groupKeyProp, err := expandCloudIdentityGroupGroupKey(d.Get("group_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("group_key"); !isEmptyValue(reflect.ValueOf(groupKeyProp)) && (ok || !reflect.DeepEqual(v, groupKeyProp)) {
		obj["groupKey"] = groupKeyProp
	}
	parentProp, err := expandCloudIdentityGroupParent(d.Get("parent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	displayNameProp, err := expandCloudIdentityGroupDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandCloudIdentityGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandCloudIdentityGroupLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := ReplaceVars(d, config, "{{CloudIdentityBasePath}}groups?initialGroupConfig={{initial_group_config}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Group: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Group: %s", err)
	}
	if err := d.Set("name", flattenCloudIdentityGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		respBody, ok := res["response"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}

		name, ok = respBody.(map[string]interface{})["name"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}
	}
	if err := d.Set("name", name.(string)); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name.(string))

	err = PollingWaitTime(resourceCloudIdentityGroupPollRead(d, meta), PollCheckForExistenceWith403, "Creating Group", d.Timeout(schema.TimeoutCreate), 10)
	if err != nil {
		return fmt.Errorf("Error waiting to create Group: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Group %q: %#v", d.Id(), res)

	return resourceCloudIdentityGroupRead(d, meta)
}

func resourceCloudIdentityGroupPollRead(d *schema.ResourceData, meta interface{}) PollReadFunc {
	return func() (map[string]interface{}, error) {
		config := meta.(*transport_tpg.Config)

		url, err := ReplaceVars(d, config, "{{CloudIdentityBasePath}}{{name}}")
		if err != nil {
			return nil, err
		}

		billingProject := ""

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		userAgent, err := generateUserAgentString(d, config.UserAgent)
		if err != nil {
			return nil, err
		}

		res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
		if err != nil {
			return res, err
		}
		return res, nil
	}
}

func resourceCloudIdentityGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{CloudIdentityBasePath}}{{name}}")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("CloudIdentityGroup %q", d.Id()))
	}

	if err := d.Set("name", flattenCloudIdentityGroupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("group_key", flattenCloudIdentityGroupGroupKey(res["groupKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("parent", flattenCloudIdentityGroupParent(res["parent"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("display_name", flattenCloudIdentityGroupDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("description", flattenCloudIdentityGroupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("create_time", flattenCloudIdentityGroupCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("update_time", flattenCloudIdentityGroupUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}
	if err := d.Set("labels", flattenCloudIdentityGroupLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Group: %s", err)
	}

	return nil
}

func resourceCloudIdentityGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	displayNameProp, err := expandCloudIdentityGroupDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandCloudIdentityGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandCloudIdentityGroupLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := ReplaceVars(d, config, "{{CloudIdentityBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Group %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
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
		return fmt.Errorf("Error updating Group %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Group %q: %#v", d.Id(), res)
	}

	err = PollingWaitTime(resourceCloudIdentityGroupPollRead(d, meta), PollCheckForExistenceWith403, "Updating Group", d.Timeout(schema.TimeoutUpdate), 10)
	if err != nil {
		return err
	}

	return resourceCloudIdentityGroupRead(d, meta)
}

func resourceCloudIdentityGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := ReplaceVars(d, config, "{{CloudIdentityBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Group %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Group")
	}

	err = PollingWaitTime(resourceCloudIdentityGroupPollRead(d, meta), PollCheckForAbsenceWith403, "Deleting Group", d.Timeout(schema.TimeoutCreate), 10)
	if err != nil {
		return fmt.Errorf("Error waiting to delete Group: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting Group %q: %#v", d.Id(), res)
	return nil
}

func resourceCloudIdentityGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	name := d.Get("name").(string)

	if d.Get("initial_group_config") == nil {
		d.Set("initial_group_config", "EMPTY")
	}

	if err := d.Set("name", name); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name)
	return []*schema.ResourceData{d}, nil
}

func flattenCloudIdentityGroupName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupGroupKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["id"] =
		flattenCloudIdentityGroupGroupKeyId(original["id"], d, config)
	transformed["namespace"] =
		flattenCloudIdentityGroupGroupKeyNamespace(original["namespace"], d, config)
	return []interface{}{transformed}
}
func flattenCloudIdentityGroupGroupKeyId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupGroupKeyNamespace(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupParent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenCloudIdentityGroupLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandCloudIdentityGroupGroupKey(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedId, err := expandCloudIdentityGroupGroupKeyId(original["id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedId); val.IsValid() && !isEmptyValue(val) {
		transformed["id"] = transformedId
	}

	transformedNamespace, err := expandCloudIdentityGroupGroupKeyNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !isEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	return transformed, nil
}

func expandCloudIdentityGroupGroupKeyId(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupGroupKeyNamespace(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupParent(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupDisplayName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdentityGroupLabels(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
