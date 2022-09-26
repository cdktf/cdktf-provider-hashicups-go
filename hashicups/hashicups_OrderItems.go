// Prebuilt hashicups Provider for Terraform CDK (cdktf)
package hashicups


type OrderItems struct {
	// coffee block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hashicups/r/order#coffee Order#coffee}
	Coffee *OrderItemsCoffee `field:"required" json:"coffee" yaml:"coffee"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hashicups/r/order#quantity Order#quantity}.
	Quantity *float64 `field:"required" json:"quantity" yaml:"quantity"`
}

