// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "tencentcloud:Ses/batchSendEmail:BatchSendEmail":
		r = &BatchSendEmail{}
	case "tencentcloud:Ses/blackListDelete:BlackListDelete":
		r = &BlackListDelete{}
	case "tencentcloud:Ses/domain:Domain":
		r = &Domain{}
	case "tencentcloud:Ses/emailAddress:EmailAddress":
		r = &EmailAddress{}
	case "tencentcloud:Ses/receiver:Receiver":
		r = &Receiver{}
	case "tencentcloud:Ses/sendEmail:SendEmail":
		r = &SendEmail{}
	case "tencentcloud:Ses/template:Template":
		r = &Template{}
	case "tencentcloud:Ses/verifyDomain:VerifyDomain":
		r = &VerifyDomain{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Ses/batchSendEmail",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Ses/blackListDelete",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Ses/domain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Ses/emailAddress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Ses/receiver",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Ses/sendEmail",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Ses/template",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Ses/verifyDomain",
		&module{version},
	)
}