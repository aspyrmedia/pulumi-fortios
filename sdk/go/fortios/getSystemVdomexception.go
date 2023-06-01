// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system vdomexception
func LookupSystemVdomexception(ctx *pulumi.Context, args *LookupSystemVdomexceptionArgs, opts ...pulumi.InvokeOption) (*LookupSystemVdomexceptionResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemVdomexceptionResult
	err := ctx.Invoke("fortios:index/getSystemVdomexception:getSystemVdomexception", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemVdomexception.
type LookupSystemVdomexceptionArgs struct {
	// Specify the fosid of the desired system vdomexception.
	Fosid int `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemVdomexception.
type LookupSystemVdomexceptionResult struct {
	// Index <1-4096>.
	Fosid int `pulumi:"fosid"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the configuration object that can be configured independently for all VDOMs.
	Object string `pulumi:"object"`
	// Object ID.
	Oid int `pulumi:"oid"`
	// Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration.
	Scope     string  `pulumi:"scope"`
	Vdomparam *string `pulumi:"vdomparam"`
	// Names of the VDOMs. The structure of `vdom` block is documented below.
	Vdoms []GetSystemVdomexceptionVdom `pulumi:"vdoms"`
}

func LookupSystemVdomexceptionOutput(ctx *pulumi.Context, args LookupSystemVdomexceptionOutputArgs, opts ...pulumi.InvokeOption) LookupSystemVdomexceptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemVdomexceptionResult, error) {
			args := v.(LookupSystemVdomexceptionArgs)
			r, err := LookupSystemVdomexception(ctx, &args, opts...)
			var s LookupSystemVdomexceptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemVdomexceptionResultOutput)
}

// A collection of arguments for invoking getSystemVdomexception.
type LookupSystemVdomexceptionOutputArgs struct {
	// Specify the fosid of the desired system vdomexception.
	Fosid pulumi.IntInput `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemVdomexceptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemVdomexceptionArgs)(nil)).Elem()
}

// A collection of values returned by getSystemVdomexception.
type LookupSystemVdomexceptionResultOutput struct{ *pulumi.OutputState }

func (LookupSystemVdomexceptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemVdomexceptionResult)(nil)).Elem()
}

func (o LookupSystemVdomexceptionResultOutput) ToLookupSystemVdomexceptionResultOutput() LookupSystemVdomexceptionResultOutput {
	return o
}

func (o LookupSystemVdomexceptionResultOutput) ToLookupSystemVdomexceptionResultOutputWithContext(ctx context.Context) LookupSystemVdomexceptionResultOutput {
	return o
}

// Index <1-4096>.
func (o LookupSystemVdomexceptionResultOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemVdomexceptionResult) int { return v.Fosid }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemVdomexceptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVdomexceptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the configuration object that can be configured independently for all VDOMs.
func (o LookupSystemVdomexceptionResultOutput) Object() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVdomexceptionResult) string { return v.Object }).(pulumi.StringOutput)
}

// Object ID.
func (o LookupSystemVdomexceptionResultOutput) Oid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemVdomexceptionResult) int { return v.Oid }).(pulumi.IntOutput)
}

// Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration.
func (o LookupSystemVdomexceptionResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVdomexceptionResult) string { return v.Scope }).(pulumi.StringOutput)
}

func (o LookupSystemVdomexceptionResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemVdomexceptionResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Names of the VDOMs. The structure of `vdom` block is documented below.
func (o LookupSystemVdomexceptionResultOutput) Vdoms() GetSystemVdomexceptionVdomArrayOutput {
	return o.ApplyT(func(v LookupSystemVdomexceptionResult) []GetSystemVdomexceptionVdom { return v.Vdoms }).(GetSystemVdomexceptionVdomArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemVdomexceptionResultOutput{})
}