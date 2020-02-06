// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.Terraform":       schema_pkg_apis_tf_v1alpha1_Terraform(ref),
		"github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.TerraformSpec":   schema_pkg_apis_tf_v1alpha1_TerraformSpec(ref),
		"github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.TerraformStatus": schema_pkg_apis_tf_v1alpha1_TerraformStatus(ref),
	}
}

func schema_pkg_apis_tf_v1alpha1_Terraform(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Terraform is the Schema for the terraforms API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.TerraformSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.TerraformStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.TerraformSpec", "github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.TerraformStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_tf_v1alpha1_TerraformSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TerraformSpec defines the desired state of Terraform",
				Properties: map[string]spec.Schema{
					"stack": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Ref:         ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.TerraformStack"),
						},
					},
					"config": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.TerraformConfig"),
						},
					},
					"sshProxy": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.ProxyOpts"),
						},
					},
					"sshKeySecretRefs": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.SSHKeySecretRef"),
									},
								},
							},
						},
					},
					"tokenSecretRefs": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.TokenSecretRef"),
									},
								},
							},
						},
					},
				},
				Required: []string{"stack", "config"},
			},
		},
		Dependencies: []string{
			"github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.ProxyOpts", "github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.SSHKeySecretRef", "github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.TerraformConfig", "github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.TerraformStack", "github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha1.TokenSecretRef"},
	}
}

func schema_pkg_apis_tf_v1alpha1_TerraformStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TerraformStatus defines the observed state of Terraform",
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"lastGeneration": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
				},
				Required: []string{"phase", "lastGeneration"},
			},
		},
		Dependencies: []string{},
	}
}
