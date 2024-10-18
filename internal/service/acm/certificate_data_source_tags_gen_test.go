// Code generated by internal/generate/tagstests/main.go; DO NOT EDIT.

package acm_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccACMCertificateDataSource_tags(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_acm_certificate.test"
	privateKeyPEM := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificatePEM := acctest.TLSRSAX509SelfSignedCertificatePEM(t, privateKeyPEM, acctest.RandomDomain().String())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ACMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/Certificate/data.tags/"),
				ConfigVariables: config.Variables{
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtKey1: config.StringVariable(acctest.CtValue1),
					}),
					acctest.CtCertificatePEM: config.StringVariable(certificatePEM),
					acctest.CtPrivateKeyPEM:  config.StringVariable(privateKeyPEM),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtKey1: knownvalue.StringExact(acctest.CtValue1),
					})),
				},
			},
		},
	})
}

func TestAccACMCertificateDataSource_tags_NullMap(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_acm_certificate.test"
	privateKeyPEM := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificatePEM := acctest.TLSRSAX509SelfSignedCertificatePEM(t, privateKeyPEM, acctest.RandomDomain().String())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ACMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/Certificate/data.tags/"),
				ConfigVariables: config.Variables{
					acctest.CtResourceTags:   nil,
					acctest.CtCertificatePEM: config.StringVariable(certificatePEM),
					acctest.CtPrivateKeyPEM:  config.StringVariable(privateKeyPEM),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{})),
				},
			},
		},
	})
}

func TestAccACMCertificateDataSource_tags_EmptyMap(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_acm_certificate.test"
	privateKeyPEM := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificatePEM := acctest.TLSRSAX509SelfSignedCertificatePEM(t, privateKeyPEM, acctest.RandomDomain().String())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.ACMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/Certificate/data.tags/"),
				ConfigVariables: config.Variables{
					acctest.CtResourceTags:   config.MapVariable(map[string]config.Variable{}),
					acctest.CtCertificatePEM: config.StringVariable(certificatePEM),
					acctest.CtPrivateKeyPEM:  config.StringVariable(privateKeyPEM),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{})),
				},
			},
		},
	})
}

func TestAccACMCertificateDataSource_tags_DefaultTags_nonOverlapping(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_acm_certificate.test"
	privateKeyPEM := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificatePEM := acctest.TLSRSAX509SelfSignedCertificatePEM(t, privateKeyPEM, acctest.RandomDomain().String())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(ctx, t) },
		ErrorCheck: acctest.ErrorCheck(t, names.ACMServiceID),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/Certificate/data.tags_defaults/"),
				ConfigVariables: config.Variables{
					acctest.CtProviderTags: config.MapVariable(map[string]config.Variable{
						acctest.CtProviderKey1: config.StringVariable(acctest.CtProviderValue1),
					}),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtResourceKey1: config.StringVariable(acctest.CtResourceValue1),
					}),
					acctest.CtCertificatePEM: config.StringVariable(certificatePEM),
					acctest.CtPrivateKeyPEM:  config.StringVariable(privateKeyPEM),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtProviderKey1: knownvalue.StringExact(acctest.CtProviderValue1),
						acctest.CtResourceKey1: knownvalue.StringExact(acctest.CtResourceValue1),
					})),
				},
			},
		},
	})
}

func TestAccACMCertificateDataSource_tags_IgnoreTags_Overlap_DefaultTag(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_acm_certificate.test"
	privateKeyPEM := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificatePEM := acctest.TLSRSAX509SelfSignedCertificatePEM(t, privateKeyPEM, acctest.RandomDomain().String())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(ctx, t) },
		ErrorCheck: acctest.ErrorCheck(t, names.ACMServiceID),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/Certificate/data.tags_ignore/"),
				ConfigVariables: config.Variables{
					acctest.CtProviderTags: config.MapVariable(map[string]config.Variable{
						acctest.CtProviderKey1: config.StringVariable(acctest.CtProviderValue1),
					}),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtResourceKey1: config.StringVariable(acctest.CtResourceValue1),
					}),
					"ignore_tag_keys": config.SetVariable(
						config.StringVariable(acctest.CtProviderKey1),
					),
					acctest.CtCertificatePEM: config.StringVariable(certificatePEM),
					acctest.CtPrivateKeyPEM:  config.StringVariable(privateKeyPEM),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtResourceKey1: knownvalue.StringExact(acctest.CtResourceValue1),
					})),
					expectFullDataSourceTags(dataSourceName, knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtProviderKey1: knownvalue.StringExact(acctest.CtProviderValue1),
						acctest.CtResourceKey1: knownvalue.StringExact(acctest.CtResourceValue1),
					})),
				},
			},
		},
	})
}

func TestAccACMCertificateDataSource_tags_IgnoreTags_Overlap_ResourceTag(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_acm_certificate.test"
	privateKeyPEM := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificatePEM := acctest.TLSRSAX509SelfSignedCertificatePEM(t, privateKeyPEM, acctest.RandomDomain().String())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(ctx, t) },
		ErrorCheck: acctest.ErrorCheck(t, names.ACMServiceID),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/Certificate/data.tags_ignore/"),
				ConfigVariables: config.Variables{
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtResourceKey1: config.StringVariable(acctest.CtResourceValue1),
					}),
					"ignore_tag_keys": config.SetVariable(
						config.StringVariable(acctest.CtResourceKey1),
					),
					acctest.CtCertificatePEM: config.StringVariable(certificatePEM),
					acctest.CtPrivateKeyPEM:  config.StringVariable(privateKeyPEM),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{})),
					expectFullDataSourceTags(dataSourceName, knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtResourceKey1: knownvalue.StringExact(acctest.CtResourceValue1),
					})),
				},
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
