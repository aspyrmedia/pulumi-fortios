// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios user saml
func LookupUserSaml(ctx *pulumi.Context, args *LookupUserSamlArgs, opts ...pulumi.InvokeOption) (*LookupUserSamlResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupUserSamlResult
	err := ctx.Invoke("fortios:index/getUserSaml:getUserSaml", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUserSaml.
type LookupUserSamlArgs struct {
	// Specify the name of the desired user saml.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getUserSaml.
type LookupUserSamlResult struct {
	// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable).
	AdfsClaim string `pulumi:"adfsClaim"`
	// URL to verify authentication.
	AuthUrl string `pulumi:"authUrl"`
	// Certificate to sign SAML messages.
	Cert string `pulumi:"cert"`
	// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
	ClockTolerance int `pulumi:"clockTolerance"`
	// Digest Method Algorithm. (default = sha1).
	DigestMethod string `pulumi:"digestMethod"`
	// SP entity ID.
	EntityId string `pulumi:"entityId"`
	// Group claim in assertion statement.
	GroupClaimType string `pulumi:"groupClaimType"`
	// Group name in assertion statement.
	GroupName string `pulumi:"groupName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IDP Certificate name.
	IdpCert string `pulumi:"idpCert"`
	// IDP entity ID.
	IdpEntityId string `pulumi:"idpEntityId"`
	// IDP single logout url.
	IdpSingleLogoutUrl string `pulumi:"idpSingleLogoutUrl"`
	// IDP single sign-on URL.
	IdpSingleSignOnUrl string `pulumi:"idpSingleSignOnUrl"`
	// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes).
	LimitRelaystate string `pulumi:"limitRelaystate"`
	// SAML server entry name.
	Name string `pulumi:"name"`
	// SP single logout URL.
	SingleLogoutUrl string `pulumi:"singleLogoutUrl"`
	// SP single sign-on URL.
	SingleSignOnUrl string `pulumi:"singleSignOnUrl"`
	// User name claim in assertion statement.
	UserClaimType string `pulumi:"userClaimType"`
	// User name in assertion statement.
	UserName  string  `pulumi:"userName"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupUserSamlOutput(ctx *pulumi.Context, args LookupUserSamlOutputArgs, opts ...pulumi.InvokeOption) LookupUserSamlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserSamlResult, error) {
			args := v.(LookupUserSamlArgs)
			r, err := LookupUserSaml(ctx, &args, opts...)
			var s LookupUserSamlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserSamlResultOutput)
}

// A collection of arguments for invoking getUserSaml.
type LookupUserSamlOutputArgs struct {
	// Specify the name of the desired user saml.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupUserSamlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserSamlArgs)(nil)).Elem()
}

// A collection of values returned by getUserSaml.
type LookupUserSamlResultOutput struct{ *pulumi.OutputState }

func (LookupUserSamlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserSamlResult)(nil)).Elem()
}

func (o LookupUserSamlResultOutput) ToLookupUserSamlResultOutput() LookupUserSamlResultOutput {
	return o
}

func (o LookupUserSamlResultOutput) ToLookupUserSamlResultOutputWithContext(ctx context.Context) LookupUserSamlResultOutput {
	return o
}

// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable).
func (o LookupUserSamlResultOutput) AdfsClaim() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.AdfsClaim }).(pulumi.StringOutput)
}

// URL to verify authentication.
func (o LookupUserSamlResultOutput) AuthUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.AuthUrl }).(pulumi.StringOutput)
}

// Certificate to sign SAML messages.
func (o LookupUserSamlResultOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.Cert }).(pulumi.StringOutput)
}

// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
func (o LookupUserSamlResultOutput) ClockTolerance() pulumi.IntOutput {
	return o.ApplyT(func(v LookupUserSamlResult) int { return v.ClockTolerance }).(pulumi.IntOutput)
}

// Digest Method Algorithm. (default = sha1).
func (o LookupUserSamlResultOutput) DigestMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.DigestMethod }).(pulumi.StringOutput)
}

// SP entity ID.
func (o LookupUserSamlResultOutput) EntityId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.EntityId }).(pulumi.StringOutput)
}

// Group claim in assertion statement.
func (o LookupUserSamlResultOutput) GroupClaimType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.GroupClaimType }).(pulumi.StringOutput)
}

// Group name in assertion statement.
func (o LookupUserSamlResultOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.GroupName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupUserSamlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.Id }).(pulumi.StringOutput)
}

// IDP Certificate name.
func (o LookupUserSamlResultOutput) IdpCert() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.IdpCert }).(pulumi.StringOutput)
}

// IDP entity ID.
func (o LookupUserSamlResultOutput) IdpEntityId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.IdpEntityId }).(pulumi.StringOutput)
}

// IDP single logout url.
func (o LookupUserSamlResultOutput) IdpSingleLogoutUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.IdpSingleLogoutUrl }).(pulumi.StringOutput)
}

// IDP single sign-on URL.
func (o LookupUserSamlResultOutput) IdpSingleSignOnUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.IdpSingleSignOnUrl }).(pulumi.StringOutput)
}

// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes).
func (o LookupUserSamlResultOutput) LimitRelaystate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.LimitRelaystate }).(pulumi.StringOutput)
}

// SAML server entry name.
func (o LookupUserSamlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.Name }).(pulumi.StringOutput)
}

// SP single logout URL.
func (o LookupUserSamlResultOutput) SingleLogoutUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.SingleLogoutUrl }).(pulumi.StringOutput)
}

// SP single sign-on URL.
func (o LookupUserSamlResultOutput) SingleSignOnUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.SingleSignOnUrl }).(pulumi.StringOutput)
}

// User name claim in assertion statement.
func (o LookupUserSamlResultOutput) UserClaimType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.UserClaimType }).(pulumi.StringOutput)
}

// User name in assertion statement.
func (o LookupUserSamlResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSamlResult) string { return v.UserName }).(pulumi.StringOutput)
}

func (o LookupUserSamlResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserSamlResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserSamlResultOutput{})
}
