// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios router multicast6
func LookupRouterMulticast6(ctx *pulumi.Context, args *LookupRouterMulticast6Args, opts ...pulumi.InvokeOption) (*LookupRouterMulticast6Result, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterMulticast6Result
	err := ctx.Invoke("fortios:index/getRouterMulticast6:getRouterMulticast6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouterMulticast6.
type LookupRouterMulticast6Args struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getRouterMulticast6.
type LookupRouterMulticast6Result struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
	Interfaces []GetRouterMulticast6Interface `pulumi:"interfaces"`
	// Enable/disable PMTU for IPv6 multicast.
	MulticastPmtu string `pulumi:"multicastPmtu"`
	// Enable/disable IPv6 multicast routing.
	MulticastRouting string `pulumi:"multicastRouting"`
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobals []GetRouterMulticast6PimSmGlobal `pulumi:"pimSmGlobals"`
	Vdomparam    *string                          `pulumi:"vdomparam"`
}

func LookupRouterMulticast6Output(ctx *pulumi.Context, args LookupRouterMulticast6OutputArgs, opts ...pulumi.InvokeOption) LookupRouterMulticast6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterMulticast6Result, error) {
			args := v.(LookupRouterMulticast6Args)
			r, err := LookupRouterMulticast6(ctx, &args, opts...)
			var s LookupRouterMulticast6Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterMulticast6ResultOutput)
}

// A collection of arguments for invoking getRouterMulticast6.
type LookupRouterMulticast6OutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterMulticast6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterMulticast6Args)(nil)).Elem()
}

// A collection of values returned by getRouterMulticast6.
type LookupRouterMulticast6ResultOutput struct{ *pulumi.OutputState }

func (LookupRouterMulticast6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterMulticast6Result)(nil)).Elem()
}

func (o LookupRouterMulticast6ResultOutput) ToLookupRouterMulticast6ResultOutput() LookupRouterMulticast6ResultOutput {
	return o
}

func (o LookupRouterMulticast6ResultOutput) ToLookupRouterMulticast6ResultOutputWithContext(ctx context.Context) LookupRouterMulticast6ResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterMulticast6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterMulticast6Result) string { return v.Id }).(pulumi.StringOutput)
}

// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
func (o LookupRouterMulticast6ResultOutput) Interfaces() GetRouterMulticast6InterfaceArrayOutput {
	return o.ApplyT(func(v LookupRouterMulticast6Result) []GetRouterMulticast6Interface { return v.Interfaces }).(GetRouterMulticast6InterfaceArrayOutput)
}

// Enable/disable PMTU for IPv6 multicast.
func (o LookupRouterMulticast6ResultOutput) MulticastPmtu() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterMulticast6Result) string { return v.MulticastPmtu }).(pulumi.StringOutput)
}

// Enable/disable IPv6 multicast routing.
func (o LookupRouterMulticast6ResultOutput) MulticastRouting() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterMulticast6Result) string { return v.MulticastRouting }).(pulumi.StringOutput)
}

// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
func (o LookupRouterMulticast6ResultOutput) PimSmGlobals() GetRouterMulticast6PimSmGlobalArrayOutput {
	return o.ApplyT(func(v LookupRouterMulticast6Result) []GetRouterMulticast6PimSmGlobal { return v.PimSmGlobals }).(GetRouterMulticast6PimSmGlobalArrayOutput)
}

func (o LookupRouterMulticast6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterMulticast6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterMulticast6ResultOutput{})
}
