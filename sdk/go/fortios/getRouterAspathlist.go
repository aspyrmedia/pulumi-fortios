// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios router aspathlist
func LookupRouterAspathlist(ctx *pulumi.Context, args *LookupRouterAspathlistArgs, opts ...pulumi.InvokeOption) (*LookupRouterAspathlistResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterAspathlistResult
	err := ctx.Invoke("fortios:index/getRouterAspathlist:getRouterAspathlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouterAspathlist.
type LookupRouterAspathlistArgs struct {
	// Specify the name of the desired router aspathlist.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getRouterAspathlist.
type LookupRouterAspathlistResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// AS path list name.
	Name string `pulumi:"name"`
	// AS path list rule. The structure of `rule` block is documented below.
	Rules     []GetRouterAspathlistRule `pulumi:"rules"`
	Vdomparam *string                   `pulumi:"vdomparam"`
}

func LookupRouterAspathlistOutput(ctx *pulumi.Context, args LookupRouterAspathlistOutputArgs, opts ...pulumi.InvokeOption) LookupRouterAspathlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterAspathlistResult, error) {
			args := v.(LookupRouterAspathlistArgs)
			r, err := LookupRouterAspathlist(ctx, &args, opts...)
			var s LookupRouterAspathlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterAspathlistResultOutput)
}

// A collection of arguments for invoking getRouterAspathlist.
type LookupRouterAspathlistOutputArgs struct {
	// Specify the name of the desired router aspathlist.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterAspathlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterAspathlistArgs)(nil)).Elem()
}

// A collection of values returned by getRouterAspathlist.
type LookupRouterAspathlistResultOutput struct{ *pulumi.OutputState }

func (LookupRouterAspathlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterAspathlistResult)(nil)).Elem()
}

func (o LookupRouterAspathlistResultOutput) ToLookupRouterAspathlistResultOutput() LookupRouterAspathlistResultOutput {
	return o
}

func (o LookupRouterAspathlistResultOutput) ToLookupRouterAspathlistResultOutputWithContext(ctx context.Context) LookupRouterAspathlistResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterAspathlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterAspathlistResult) string { return v.Id }).(pulumi.StringOutput)
}

// AS path list name.
func (o LookupRouterAspathlistResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterAspathlistResult) string { return v.Name }).(pulumi.StringOutput)
}

// AS path list rule. The structure of `rule` block is documented below.
func (o LookupRouterAspathlistResultOutput) Rules() GetRouterAspathlistRuleArrayOutput {
	return o.ApplyT(func(v LookupRouterAspathlistResult) []GetRouterAspathlistRule { return v.Rules }).(GetRouterAspathlistRuleArrayOutput)
}

func (o LookupRouterAspathlistResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterAspathlistResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterAspathlistResultOutput{})
}