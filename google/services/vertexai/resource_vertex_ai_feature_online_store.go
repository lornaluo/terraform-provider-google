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

package vertexai

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceVertexAIFeatureOnlineStore() *schema.Resource {
	return &schema.Resource{
		Create: resourceVertexAIFeatureOnlineStoreCreate,
		Read:   resourceVertexAIFeatureOnlineStoreRead,
		Update: resourceVertexAIFeatureOnlineStoreUpdate,
		Delete: resourceVertexAIFeatureOnlineStoreDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVertexAIFeatureOnlineStoreImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource name of the Feature Online Store.`,
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The region of feature online store. eg us-central1`,
			},
			"bigtable": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Settings for Cloud Bigtable instance that will be created to serve featureValues for all FeatureViews under this FeatureOnlineStore.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_scaling": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `Autoscaling config applied to Bigtable Instance.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_node_count": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: `The maximum number of nodes to scale up to. Must be greater than or equal to minNodeCount, and less than or equal to 10 times of 'minNodeCount'.`,
									},
									"min_node_count": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: `The minimum number of nodes to scale down to. Must be greater than or equal to 1.`,
									},
									"cpu_utilization_target": {
										Type:        schema.TypeInt,
										Computed:    true,
										Optional:    true,
										Description: `A percentage of the cluster's CPU capacity. Can be from 10% to 80%. When a cluster's CPU utilization exceeds the target that you have set, Bigtable immediately adds nodes to the cluster. When CPU utilization is substantially lower than the target, Bigtable removes nodes. If not set will default to 50%.`,
									},
								},
							},
						},
					},
				},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `The labels with user-defined metadata to organize your feature online stores.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the feature online store was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Used to perform consistent read-modify-write updates.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The state of the Feature Online Store. See the possible states in [this link](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featureOnlineStores#state).`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the feature online store was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
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

func resourceVertexAIFeatureOnlineStoreCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandVertexAIFeatureOnlineStoreName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	bigtableProp, err := expandVertexAIFeatureOnlineStoreBigtable(d.Get("bigtable"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bigtable"); !tpgresource.IsEmptyValue(reflect.ValueOf(bigtableProp)) && (ok || !reflect.DeepEqual(v, bigtableProp)) {
		obj["bigtable"] = bigtableProp
	}
	labelsProp, err := expandVertexAIFeatureOnlineStoreEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featureOnlineStores?featureOnlineStoreId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FeatureOnlineStore: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for FeatureOnlineStore: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating FeatureOnlineStore: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = VertexAIOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating FeatureOnlineStore", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create FeatureOnlineStore: %s", err)
	}

	if err := d.Set("name", flattenVertexAIFeatureOnlineStoreName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating FeatureOnlineStore %q: %#v", d.Id(), res)

	return resourceVertexAIFeatureOnlineStoreRead(d, meta)
}

func resourceVertexAIFeatureOnlineStoreRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for FeatureOnlineStore: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("VertexAIFeatureOnlineStore %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}

	if err := d.Set("name", flattenVertexAIFeatureOnlineStoreName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}
	if err := d.Set("create_time", flattenVertexAIFeatureOnlineStoreCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}
	if err := d.Set("update_time", flattenVertexAIFeatureOnlineStoreUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}
	if err := d.Set("labels", flattenVertexAIFeatureOnlineStoreLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}
	if err := d.Set("state", flattenVertexAIFeatureOnlineStoreState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}
	if err := d.Set("bigtable", flattenVertexAIFeatureOnlineStoreBigtable(res["bigtable"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}
	if err := d.Set("terraform_labels", flattenVertexAIFeatureOnlineStoreTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}
	if err := d.Set("effective_labels", flattenVertexAIFeatureOnlineStoreEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeatureOnlineStore: %s", err)
	}

	return nil
}

func resourceVertexAIFeatureOnlineStoreUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for FeatureOnlineStore: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	bigtableProp, err := expandVertexAIFeatureOnlineStoreBigtable(d.Get("bigtable"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bigtable"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, bigtableProp)) {
		obj["bigtable"] = bigtableProp
	}
	labelsProp, err := expandVertexAIFeatureOnlineStoreEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating FeatureOnlineStore %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("bigtable") {
		updateMask = append(updateMask, "bigtable")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})

		if err != nil {
			return fmt.Errorf("Error updating FeatureOnlineStore %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating FeatureOnlineStore %q: %#v", d.Id(), res)
		}

		err = VertexAIOperationWaitTime(
			config, res, project, "Updating FeatureOnlineStore", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceVertexAIFeatureOnlineStoreRead(d, meta)
}

func resourceVertexAIFeatureOnlineStoreDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for FeatureOnlineStore: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting FeatureOnlineStore %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "FeatureOnlineStore")
	}

	err = VertexAIOperationWaitTime(
		config, res, project, "Deleting FeatureOnlineStore", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting FeatureOnlineStore %q: %#v", d.Id(), res)
	return nil
}

func resourceVertexAIFeatureOnlineStoreImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/featureOnlineStores/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenVertexAIFeatureOnlineStoreName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenVertexAIFeatureOnlineStoreCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeatureOnlineStoreUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeatureOnlineStoreLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenVertexAIFeatureOnlineStoreState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeatureOnlineStoreBigtable(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["auto_scaling"] =
		flattenVertexAIFeatureOnlineStoreBigtableAutoScaling(original["autoScaling"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeatureOnlineStoreBigtableAutoScaling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["min_node_count"] =
		flattenVertexAIFeatureOnlineStoreBigtableAutoScalingMinNodeCount(original["minNodeCount"], d, config)
	transformed["max_node_count"] =
		flattenVertexAIFeatureOnlineStoreBigtableAutoScalingMaxNodeCount(original["maxNodeCount"], d, config)
	transformed["cpu_utilization_target"] =
		flattenVertexAIFeatureOnlineStoreBigtableAutoScalingCpuUtilizationTarget(original["cpuUtilizationTarget"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIFeatureOnlineStoreBigtableAutoScalingMinNodeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenVertexAIFeatureOnlineStoreBigtableAutoScalingMaxNodeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenVertexAIFeatureOnlineStoreBigtableAutoScalingCpuUtilizationTarget(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenVertexAIFeatureOnlineStoreTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenVertexAIFeatureOnlineStoreEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandVertexAIFeatureOnlineStoreName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureOnlineStoreBigtable(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAutoScaling, err := expandVertexAIFeatureOnlineStoreBigtableAutoScaling(original["auto_scaling"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAutoScaling); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["autoScaling"] = transformedAutoScaling
	}

	return transformed, nil
}

func expandVertexAIFeatureOnlineStoreBigtableAutoScaling(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinNodeCount, err := expandVertexAIFeatureOnlineStoreBigtableAutoScalingMinNodeCount(original["min_node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minNodeCount"] = transformedMinNodeCount
	}

	transformedMaxNodeCount, err := expandVertexAIFeatureOnlineStoreBigtableAutoScalingMaxNodeCount(original["max_node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxNodeCount"] = transformedMaxNodeCount
	}

	transformedCpuUtilizationTarget, err := expandVertexAIFeatureOnlineStoreBigtableAutoScalingCpuUtilizationTarget(original["cpu_utilization_target"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCpuUtilizationTarget); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cpuUtilizationTarget"] = transformedCpuUtilizationTarget
	}

	return transformed, nil
}

func expandVertexAIFeatureOnlineStoreBigtableAutoScalingMinNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureOnlineStoreBigtableAutoScalingMaxNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureOnlineStoreBigtableAutoScalingCpuUtilizationTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeatureOnlineStoreEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}