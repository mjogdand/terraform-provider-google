// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/iam3/ProjectsPolicyBinding.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package iam3

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceIAM3ProjectsPolicyBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceIAM3ProjectsPolicyBindingCreate,
		Read:   resourceIAM3ProjectsPolicyBindingRead,
		Update: resourceIAM3ProjectsPolicyBindingUpdate,
		Delete: resourceIAM3ProjectsPolicyBindingDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIAM3ProjectsPolicyBindingImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetAnnotationsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location of the Policy Binding`,
			},
			"policy": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Required. Immutable. The resource name of the policy to be bound. The binding parent and policy must belong to the same Organization (or Project).`,
			},
			"policy_binding_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The Policy Binding ID.`,
			},
			"target": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Target is the full resource name of the resource to which the policy will be bound. Immutable once set.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"principal_set": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Required. Immutable. Full Resource Name of the principal set used for principal access boundary policy bindings.
Examples for each one of the following supported principal set types:
* Project:
  * '//cloudresourcemanager.googleapis.com/projects/PROJECT_NUMBER'
  * '//cloudresourcemanager.googleapis.com/projects/PROJECT_ID'
* Workload Identity Pool: '//iam.googleapis.com/projects/PROJECT_NUMBER/locations/LOCATION/workloadIdentityPools/WORKLOAD_POOL_ID'
It must be parent by the policy binding's parent (the project).`,
						},
					},
				},
			},
			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Optional. User defined annotations. See https://google.aip.dev/148#annotations for more details such as format and size limitations


**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
Please refer to the field 'effective_annotations' for all of the annotations present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"condition": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Represents a textual expression in the Common Expression Language
(CEL) syntax. CEL is a C-like expression language. The syntax and semantics of
CEL are documented at https://github.com/google/cel-spec.
Example (Comparison):
title: \"Summary size limit\"
description: \"Determines if a summary is less than 100 chars\"
expression: \"document.summary.size() < 100\"
Example
(Equality):
title: \"Requestor is owner\"
description: \"Determines if requestor is the document owner\"
expression: \"document.owner == request.auth.claims.email\"  Example
(Logic):
title: \"Public documents\"
description: \"Determine whether the document should be publicly visible\"
expression: \"document.type != 'private' && document.type != 'internal'\"
Example (Data Manipulation):
title: \"Notification string\"
description: \"Create a notification string with a timestamp.\"
expression: \"'New message received at ' + string(document.create_time)\"
The exact variables and functions that may be referenced within an expression are
determined by the service that evaluates it. See the service documentation for
additional information.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.`,
						},
						"expression": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Textual representation of an expression in Common Expression Language syntax.`,
						},
						"location": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.`,
						},
						"title": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.`,
						},
					},
				},
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Optional. The description of the policy binding. Must be less than or equal to 63 characters.`,
			},
			"policy_kind": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Immutable. The kind of the policy to attach in this binding. This
field must be one of the following:  - Left empty (will be automatically set
to the policy kind) - The input policy kind   Possible values:  POLICY_KIND_UNSPECIFIED PRINCIPAL_ACCESS_BOUNDARY ACCESS`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The time when the policy binding was created.`,
			},
			"effective_annotations": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Optional. The etag for the policy binding. If this is provided on update, it must match the server's etag.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the policy binding in the format '{binding_parent/locations/{location}/policyBindings/{policy_binding_id}'`,
			},
			"policy_uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The globally unique ID of the policy to be bound.`,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The globally unique ID of the policy binding. Assigned when the policy binding is created.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The time when the policy binding was most recently updated.`,
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

func resourceIAM3ProjectsPolicyBindingCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	etagProp, err := expandIAM3ProjectsPolicyBindingEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	displayNameProp, err := expandIAM3ProjectsPolicyBindingDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	targetProp, err := expandIAM3ProjectsPolicyBindingTarget(d.Get("target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	policyKindProp, err := expandIAM3ProjectsPolicyBindingPolicyKind(d.Get("policy_kind"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("policy_kind"); !tpgresource.IsEmptyValue(reflect.ValueOf(policyKindProp)) && (ok || !reflect.DeepEqual(v, policyKindProp)) {
		obj["policyKind"] = policyKindProp
	}
	policyProp, err := expandIAM3ProjectsPolicyBindingPolicy(d.Get("policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(policyProp)) && (ok || !reflect.DeepEqual(v, policyProp)) {
		obj["policy"] = policyProp
	}
	conditionProp, err := expandIAM3ProjectsPolicyBindingCondition(d.Get("condition"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("condition"); !tpgresource.IsEmptyValue(reflect.ValueOf(conditionProp)) && (ok || !reflect.DeepEqual(v, conditionProp)) {
		obj["condition"] = conditionProp
	}
	annotationsProp, err := expandIAM3ProjectsPolicyBindingEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IAM3BasePath}}projects/{{project}}/locations/{{location}}/policyBindings?policyBindingId={{policy_binding_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ProjectsPolicyBinding: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectsPolicyBinding: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating ProjectsPolicyBinding: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/policyBindings/{{policy_binding_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = IAM3OperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating ProjectsPolicyBinding", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create ProjectsPolicyBinding: %s", err)
	}

	if err := d.Set("name", flattenIAM3ProjectsPolicyBindingName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/policyBindings/{{policy_binding_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ProjectsPolicyBinding %q: %#v", d.Id(), res)

	return resourceIAM3ProjectsPolicyBindingRead(d, meta)
}

func resourceIAM3ProjectsPolicyBindingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IAM3BasePath}}projects/{{project}}/locations/{{location}}/policyBindings/{{policy_binding_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectsPolicyBinding: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IAM3ProjectsPolicyBinding %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ProjectsPolicyBinding: %s", err)
	}

	if err := d.Set("name", flattenIAM3ProjectsPolicyBindingName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectsPolicyBinding: %s", err)
	}
	if err := d.Set("uid", flattenIAM3ProjectsPolicyBindingUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectsPolicyBinding: %s", err)
	}
	if err := d.Set("etag", flattenIAM3ProjectsPolicyBindingEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectsPolicyBinding: %s", err)
	}
	if err := d.Set("display_name", flattenIAM3ProjectsPolicyBindingDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectsPolicyBinding: %s", err)
	}
	if err := d.Set("annotations", flattenIAM3ProjectsPolicyBindingAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectsPolicyBinding: %s", err)
	}
	if err := d.Set("target", flattenIAM3ProjectsPolicyBindingTarget(res["target"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectsPolicyBinding: %s", err)
	}
	if err := d.Set("policy_kind", flattenIAM3ProjectsPolicyBindingPolicyKind(res["policyKind"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectsPolicyBinding: %s", err)
	}
	if err := d.Set("policy", flattenIAM3ProjectsPolicyBindingPolicy(res["policy"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectsPolicyBinding: %s", err)
	}
	if err := d.Set("policy_uid", flattenIAM3ProjectsPolicyBindingPolicyUid(res["policyUid"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectsPolicyBinding: %s", err)
	}
	if err := d.Set("condition", flattenIAM3ProjectsPolicyBindingCondition(res["condition"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectsPolicyBinding: %s", err)
	}
	if err := d.Set("create_time", flattenIAM3ProjectsPolicyBindingCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectsPolicyBinding: %s", err)
	}
	if err := d.Set("update_time", flattenIAM3ProjectsPolicyBindingUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectsPolicyBinding: %s", err)
	}
	if err := d.Set("effective_annotations", flattenIAM3ProjectsPolicyBindingEffectiveAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectsPolicyBinding: %s", err)
	}

	return nil
}

func resourceIAM3ProjectsPolicyBindingUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectsPolicyBinding: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	etagProp, err := expandIAM3ProjectsPolicyBindingEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	displayNameProp, err := expandIAM3ProjectsPolicyBindingDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	targetProp, err := expandIAM3ProjectsPolicyBindingTarget(d.Get("target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	conditionProp, err := expandIAM3ProjectsPolicyBindingCondition(d.Get("condition"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("condition"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, conditionProp)) {
		obj["condition"] = conditionProp
	}
	annotationsProp, err := expandIAM3ProjectsPolicyBindingEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IAM3BasePath}}projects/{{project}}/locations/{{location}}/policyBindings/{{policy_binding_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ProjectsPolicyBinding %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("etag") {
		updateMask = append(updateMask, "etag")
	}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("target") {
		updateMask = append(updateMask, "target")
	}

	if d.HasChange("condition") {
		updateMask = append(updateMask, "condition")
	}

	if d.HasChange("effective_annotations") {
		updateMask = append(updateMask, "annotations")
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
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating ProjectsPolicyBinding %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ProjectsPolicyBinding %q: %#v", d.Id(), res)
		}

		err = IAM3OperationWaitTime(
			config, res, project, "Updating ProjectsPolicyBinding", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceIAM3ProjectsPolicyBindingRead(d, meta)
}

func resourceIAM3ProjectsPolicyBindingDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectsPolicyBinding: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{IAM3BasePath}}projects/{{project}}/locations/{{location}}/policyBindings/{{policy_binding_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting ProjectsPolicyBinding %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "ProjectsPolicyBinding")
	}

	err = IAM3OperationWaitTime(
		config, res, project, "Deleting ProjectsPolicyBinding", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}
	// This is useful if the resource in question doesn't have a perfectly consistent API
	// That is, if the deletion of a dependent resource has not propagated.
	time.Sleep(5 * time.Second)

	log.Printf("[DEBUG] Finished deleting ProjectsPolicyBinding %q: %#v", d.Id(), res)
	return nil
}

func resourceIAM3ProjectsPolicyBindingImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/policyBindings/(?P<policy_binding_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<policy_binding_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<policy_binding_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/policyBindings/{{policy_binding_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIAM3ProjectsPolicyBindingName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3ProjectsPolicyBindingUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3ProjectsPolicyBindingEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3ProjectsPolicyBindingDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3ProjectsPolicyBindingAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("annotations"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenIAM3ProjectsPolicyBindingTarget(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["principal_set"] =
		flattenIAM3ProjectsPolicyBindingTargetPrincipalSet(original["principalSet"], d, config)
	return []interface{}{transformed}
}
func flattenIAM3ProjectsPolicyBindingTargetPrincipalSet(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3ProjectsPolicyBindingPolicyKind(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3ProjectsPolicyBindingPolicy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3ProjectsPolicyBindingPolicyUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3ProjectsPolicyBindingCondition(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["expression"] =
		flattenIAM3ProjectsPolicyBindingConditionExpression(original["expression"], d, config)
	transformed["title"] =
		flattenIAM3ProjectsPolicyBindingConditionTitle(original["title"], d, config)
	transformed["description"] =
		flattenIAM3ProjectsPolicyBindingConditionDescription(original["description"], d, config)
	transformed["location"] =
		flattenIAM3ProjectsPolicyBindingConditionLocation(original["location"], d, config)
	return []interface{}{transformed}
}
func flattenIAM3ProjectsPolicyBindingConditionExpression(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3ProjectsPolicyBindingConditionTitle(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3ProjectsPolicyBindingConditionDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3ProjectsPolicyBindingConditionLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3ProjectsPolicyBindingCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3ProjectsPolicyBindingUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3ProjectsPolicyBindingEffectiveAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandIAM3ProjectsPolicyBindingEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3ProjectsPolicyBindingDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3ProjectsPolicyBindingTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPrincipalSet, err := expandIAM3ProjectsPolicyBindingTargetPrincipalSet(original["principal_set"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrincipalSet); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["principalSet"] = transformedPrincipalSet
	}

	return transformed, nil
}

func expandIAM3ProjectsPolicyBindingTargetPrincipalSet(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3ProjectsPolicyBindingPolicyKind(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3ProjectsPolicyBindingPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3ProjectsPolicyBindingCondition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandIAM3ProjectsPolicyBindingConditionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandIAM3ProjectsPolicyBindingConditionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandIAM3ProjectsPolicyBindingConditionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandIAM3ProjectsPolicyBindingConditionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandIAM3ProjectsPolicyBindingConditionExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3ProjectsPolicyBindingConditionTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3ProjectsPolicyBindingConditionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3ProjectsPolicyBindingConditionLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3ProjectsPolicyBindingEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
