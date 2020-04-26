// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package yandexexport

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	Paths               []string          `mapstructure:"paths" required:"true" cty:"paths"`
	FolderID            *string           `mapstructure:"folder_id" required:"true" cty:"folder_id"`
	ServiceAccountID    *string           `mapstructure:"service_account_id" required:"true" cty:"service_account_id"`
	DiskSizeGb          *int              `mapstructure:"disk_size" required:"false" cty:"disk_size"`
	DiskType            *string           `mapstructure:"disk_type" required:"false" cty:"disk_type"`
	PlatformID          *string           `mapstructure:"platform_id" required:"false" cty:"platform_id"`
	SubnetID            *string           `mapstructure:"subnet_id" required:"false" cty:"subnet_id"`
	Zone                *string           `mapstructure:"zone" required:"false" cty:"zone"`
	Token               *string           `mapstructure:"token" required:"false" cty:"token"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"paths":                      &hcldec.AttrSpec{Name: "paths", Type: cty.List(cty.String), Required: false},
		"folder_id":                  &hcldec.AttrSpec{Name: "folder_id", Type: cty.String, Required: false},
		"service_account_id":         &hcldec.AttrSpec{Name: "service_account_id", Type: cty.String, Required: false},
		"disk_size":                  &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"disk_type":                  &hcldec.AttrSpec{Name: "disk_type", Type: cty.String, Required: false},
		"platform_id":                &hcldec.AttrSpec{Name: "platform_id", Type: cty.String, Required: false},
		"subnet_id":                  &hcldec.AttrSpec{Name: "subnet_id", Type: cty.String, Required: false},
		"zone":                       &hcldec.AttrSpec{Name: "zone", Type: cty.String, Required: false},
		"token":                      &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
	}
	return s
}
