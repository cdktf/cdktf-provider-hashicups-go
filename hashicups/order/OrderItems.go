package order


type OrderItems struct {
	// coffee block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hashicups/0.3.1/docs/resources/order#coffee Order#coffee}
	Coffee *OrderItemsCoffee `field:"required" json:"coffee" yaml:"coffee"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hashicups/0.3.1/docs/resources/order#quantity Order#quantity}.
	Quantity *float64 `field:"required" json:"quantity" yaml:"quantity"`
}

