// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system zone
func LookupSystemZone(ctx *pulumi.Context, args *LookupSystemZoneArgs, opts ...pulumi.InvokeOption) (*LookupSystemZoneResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemZoneResult
	err := ctx.Invoke("fortios:index/getSystemZone:getSystemZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemZone.
type LookupSystemZoneArgs struct {
	// Specify the name of the desired system zone.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemZone.
type LookupSystemZoneResult struct {
	// Description.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
	Interfaces []GetSystemZoneInterface `pulumi:"interfaces"`
	// Allow or deny traffic routing between different interfaces in the same zone (default = deny).
	Intrazone string `pulumi:"intrazone"`
	// Tag name.
	Name string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings  []GetSystemZoneTagging `pulumi:"taggings"`
	Vdomparam *string                `pulumi:"vdomparam"`
}

func LookupSystemZoneOutput(ctx *pulumi.Context, args LookupSystemZoneOutputArgs, opts ...pulumi.InvokeOption) LookupSystemZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemZoneResult, error) {
			args := v.(LookupSystemZoneArgs)
			r, err := LookupSystemZone(ctx, &args, opts...)
			var s LookupSystemZoneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemZoneResultOutput)
}

// A collection of arguments for invoking getSystemZone.
type LookupSystemZoneOutputArgs struct {
	// Specify the name of the desired system zone.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemZoneArgs)(nil)).Elem()
}

// A collection of values returned by getSystemZone.
type LookupSystemZoneResultOutput struct{ *pulumi.OutputState }

func (LookupSystemZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemZoneResult)(nil)).Elem()
}

func (o LookupSystemZoneResultOutput) ToLookupSystemZoneResultOutput() LookupSystemZoneResultOutput {
	return o
}

func (o LookupSystemZoneResultOutput) ToLookupSystemZoneResultOutputWithContext(ctx context.Context) LookupSystemZoneResultOutput {
	return o
}

// Description.
func (o LookupSystemZoneResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemZoneResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

// Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
func (o LookupSystemZoneResultOutput) Interfaces() GetSystemZoneInterfaceArrayOutput {
	return o.ApplyT(func(v LookupSystemZoneResult) []GetSystemZoneInterface { return v.Interfaces }).(GetSystemZoneInterfaceArrayOutput)
}

// Allow or deny traffic routing between different interfaces in the same zone (default = deny).
func (o LookupSystemZoneResultOutput) Intrazone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemZoneResult) string { return v.Intrazone }).(pulumi.StringOutput)
}

// Tag name.
func (o LookupSystemZoneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemZoneResult) string { return v.Name }).(pulumi.StringOutput)
}

// Config object tagging. The structure of `tagging` block is documented below.
func (o LookupSystemZoneResultOutput) Taggings() GetSystemZoneTaggingArrayOutput {
	return o.ApplyT(func(v LookupSystemZoneResult) []GetSystemZoneTagging { return v.Taggings }).(GetSystemZoneTaggingArrayOutput)
}

func (o LookupSystemZoneResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemZoneResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemZoneResultOutput{})
}