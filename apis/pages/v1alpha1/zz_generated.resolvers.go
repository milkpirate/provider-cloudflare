/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/milkpirate/provider-cloudflare/apis/account/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Domain.
func (mg *Domain) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AccountIDRef,
		Selector:     mg.Spec.ForProvider.AccountIDSelector,
		To: reference.To{
			List:    &v1alpha1.AccountList{},
			Managed: &v1alpha1.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccountID")
	}
	mg.Spec.ForProvider.AccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccountIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProjectName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProjectNameRef,
		Selector:     mg.Spec.ForProvider.ProjectNameSelector,
		To: reference.To{
			List:    &ProjectList{},
			Managed: &Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProjectName")
	}
	mg.Spec.ForProvider.ProjectName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Project.
func (mg *Project) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AccountIDRef,
		Selector:     mg.Spec.ForProvider.AccountIDSelector,
		To: reference.To{
			List:    &v1alpha1.AccountList{},
			Managed: &v1alpha1.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccountID")
	}
	mg.Spec.ForProvider.AccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccountIDRef = rsp.ResolvedReference

	return nil
}
