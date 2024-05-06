// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Ruleset
func (mg *Ruleset) GetTerraformResourceType() string {
	return "cloudflare_ruleset"
}

// GetConnectionDetailsMapping for this Ruleset
func (tr *Ruleset) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Ruleset
func (tr *Ruleset) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Ruleset
func (tr *Ruleset) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Ruleset
func (tr *Ruleset) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Ruleset
func (tr *Ruleset) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Ruleset
func (tr *Ruleset) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this Ruleset
func (tr *Ruleset) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this Ruleset using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Ruleset) LateInitialize(attrs []byte) (bool, error) {
	params := &RulesetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}
	opts = append(opts, resource.WithNameFilter("AccountID"))
	opts = append(opts, resource.WithNameFilter("Rules.ActionParameters.CacheKey.CacheByDeviceType"))
	opts = append(opts, resource.WithNameFilter("Rules.ActionParameters.CacheKey.CustomKey.QueryString.Exclude"))
	opts = append(opts, resource.WithNameFilter("Rules.ActionParameters.CacheKey.CustomKey.QueryString.Include"))
	opts = append(opts, resource.WithNameFilter("Rules.ActionParameters.CacheKey.CustomKey.User.DeviceType"))
	opts = append(opts, resource.WithNameFilter("Rules.ActionParameters.EdgeTTL.StatusCodeTTL.StatusCode"))
	opts = append(opts, resource.WithNameFilter("Rules.ActionParameters.EdgeTTL.StatusCodeTTL.StatusCodeRange"))
	opts = append(opts, resource.WithNameFilter("Rules.ActionParameters.FromValue.TargetURL.Expression"))
	opts = append(opts, resource.WithNameFilter("Rules.ActionParameters.FromValue.TargetURL.Value"))
	opts = append(opts, resource.WithNameFilter("Rules.ActionParameters.Headers.Expression"))
	opts = append(opts, resource.WithNameFilter("Rules.ActionParameters.Headers.Value"))
	opts = append(opts, resource.WithNameFilter("ZoneID"))

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Ruleset) GetTerraformSchemaVersion() int {
	return 0
}
