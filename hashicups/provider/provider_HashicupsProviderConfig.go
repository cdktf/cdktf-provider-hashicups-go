package provider


type HashicupsProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hashicups#alias HashicupsProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hashicups#host HashicupsProvider#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hashicups#password HashicupsProvider#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/hashicups#username HashicupsProvider#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

