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
)

func ResourceGameServicesGameServerDeploymentRollout() *schema.Resource {
	return &schema.Resource{
		Create: resourceGameServicesGameServerDeploymentRolloutCreate,
		Read:   resourceGameServicesGameServerDeploymentRolloutRead,
		Update: resourceGameServicesGameServerDeploymentRolloutUpdate,
		Delete: resourceGameServicesGameServerDeploymentRolloutDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGameServicesGameServerDeploymentRolloutImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"default_game_server_config": {
				Type:     schema.TypeString,
				Required: true,
				Description: `This field points to the game server config that is
applied by default to all realms and clusters. For example,

'projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config'.`,
			},
			"deployment_id": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The deployment to rollout the new config to. Only 1 rollout must be associated with each deployment.`,
			},
			"game_server_config_overrides": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `The game_server_config_overrides contains the per game server config
overrides. The overrides are processed in the order they are listed. As
soon as a match is found for a cluster, the rest of the list is not
processed.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config_version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Version of the configuration.`,
						},
						"realms_selector": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Selection by realms.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"realms": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `List of realms to match against.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource id of the game server deployment

eg: 'projects/my-project/locations/global/gameServerDeployments/my-deployment/rollout'.`,
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

func resourceGameServicesGameServerDeploymentRolloutCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/global/gameServerDeployments/{{deployment_id}}/rollout")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Creating GameServerDeploymentRollout %q: ", d.Id())

	err = resourceGameServicesGameServerDeploymentRolloutUpdate(d, meta)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error trying to create GameServerDeploymentRollout: %s", err)
	}

	return nil
}

func resourceGameServicesGameServerDeploymentRolloutRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{GameServicesBasePath}}projects/{{project}}/locations/global/gameServerDeployments/{{deployment_id}}/rollout")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GameServerDeploymentRollout: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("GameServicesGameServerDeploymentRollout %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading GameServerDeploymentRollout: %s", err)
	}

	if err := d.Set("name", flattenGameServicesGameServerDeploymentRolloutName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading GameServerDeploymentRollout: %s", err)
	}
	if err := d.Set("default_game_server_config", flattenGameServicesGameServerDeploymentRolloutDefaultGameServerConfig(res["defaultGameServerConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading GameServerDeploymentRollout: %s", err)
	}
	if err := d.Set("game_server_config_overrides", flattenGameServicesGameServerDeploymentRolloutGameServerConfigOverrides(res["gameServerConfigOverrides"], d, config)); err != nil {
		return fmt.Errorf("Error reading GameServerDeploymentRollout: %s", err)
	}

	return nil
}

func resourceGameServicesGameServerDeploymentRolloutUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GameServerDeploymentRollout: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	defaultGameServerConfigProp, err := expandGameServicesGameServerDeploymentRolloutDefaultGameServerConfig(d.Get("default_game_server_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_game_server_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, defaultGameServerConfigProp)) {
		obj["defaultGameServerConfig"] = defaultGameServerConfigProp
	}
	gameServerConfigOverridesProp, err := expandGameServicesGameServerDeploymentRolloutGameServerConfigOverrides(d.Get("game_server_config_overrides"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("game_server_config_overrides"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, gameServerConfigOverridesProp)) {
		obj["gameServerConfigOverrides"] = gameServerConfigOverridesProp
	}

	url, err := ReplaceVars(d, config, "{{GameServicesBasePath}}projects/{{project}}/locations/global/gameServerDeployments/{{deployment_id}}/rollout")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating GameServerDeploymentRollout %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("default_game_server_config") {
		updateMask = append(updateMask, "defaultGameServerConfig")
	}

	if d.HasChange("game_server_config_overrides") {
		updateMask = append(updateMask, "gameServerConfigOverrides")
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
		return fmt.Errorf("Error updating GameServerDeploymentRollout %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating GameServerDeploymentRollout %q: %#v", d.Id(), res)
	}

	err = GameServicesOperationWaitTime(
		config, res, project, "Updating GameServerDeploymentRollout", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceGameServicesGameServerDeploymentRolloutRead(d, meta)
}

func resourceGameServicesGameServerDeploymentRolloutDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GameServerDeploymentRollout: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{GameServicesBasePath}}projects/{{project}}/locations/global/gameServerDeployments/{{deployment_id}}/rollout?updateMask=defaultGameServerConfig")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting GameServerDeploymentRollout %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "GameServerDeploymentRollout")
	}

	err = GameServicesOperationWaitTime(
		config, res, project, "Deleting GameServerDeploymentRollout", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting GameServerDeploymentRollout %q: %#v", d.Id(), res)
	return nil
}

func resourceGameServicesGameServerDeploymentRolloutImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/global/gameServerDeployments/(?P<deployment_id>[^/]+)/rollout",
		"(?P<project>[^/]+)/(?P<deployment_id>[^/]+)",
		"(?P<deployment_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/locations/global/gameServerDeployments/{{deployment_id}}/rollout")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenGameServicesGameServerDeploymentRolloutName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGameServicesGameServerDeploymentRolloutDefaultGameServerConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGameServicesGameServerDeploymentRolloutGameServerConfigOverrides(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"realms_selector": flattenGameServicesGameServerDeploymentRolloutGameServerConfigOverridesRealmsSelector(original["realmsSelector"], d, config),
			"config_version":  flattenGameServicesGameServerDeploymentRolloutGameServerConfigOverridesConfigVersion(original["configVersion"], d, config),
		})
	}
	return transformed
}
func flattenGameServicesGameServerDeploymentRolloutGameServerConfigOverridesRealmsSelector(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["realms"] =
		flattenGameServicesGameServerDeploymentRolloutGameServerConfigOverridesRealmsSelectorRealms(original["realms"], d, config)
	return []interface{}{transformed}
}
func flattenGameServicesGameServerDeploymentRolloutGameServerConfigOverridesRealmsSelectorRealms(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGameServicesGameServerDeploymentRolloutGameServerConfigOverridesConfigVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandGameServicesGameServerDeploymentRolloutDefaultGameServerConfig(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerDeploymentRolloutGameServerConfigOverrides(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedRealmsSelector, err := expandGameServicesGameServerDeploymentRolloutGameServerConfigOverridesRealmsSelector(original["realms_selector"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRealmsSelector); val.IsValid() && !isEmptyValue(val) {
			transformed["realmsSelector"] = transformedRealmsSelector
		}

		transformedConfigVersion, err := expandGameServicesGameServerDeploymentRolloutGameServerConfigOverridesConfigVersion(original["config_version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedConfigVersion); val.IsValid() && !isEmptyValue(val) {
			transformed["configVersion"] = transformedConfigVersion
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGameServicesGameServerDeploymentRolloutGameServerConfigOverridesRealmsSelector(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRealms, err := expandGameServicesGameServerDeploymentRolloutGameServerConfigOverridesRealmsSelectorRealms(original["realms"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRealms); val.IsValid() && !isEmptyValue(val) {
		transformed["realms"] = transformedRealms
	}

	return transformed, nil
}

func expandGameServicesGameServerDeploymentRolloutGameServerConfigOverridesRealmsSelectorRealms(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerDeploymentRolloutGameServerConfigOverridesConfigVersion(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
