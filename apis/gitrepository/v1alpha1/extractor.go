package v1alpha1

import (
	xpref "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource"
)

// ExtractResourceID extracts the value of `spec.atProvider.id`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractResourceID() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {

		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		return tr.GetID()
	}
}
