// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package order


type OrderItemsCoffee struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hashicups/0.3.1/docs/resources/order#id Order#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"required" json:"id" yaml:"id"`
}

