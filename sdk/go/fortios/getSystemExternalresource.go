// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system externalresource
func LookupSystemExternalresource(ctx *pulumi.Context, args *LookupSystemExternalresourceArgs, opts ...pulumi.InvokeOption) (*LookupSystemExternalresourceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemExternalresourceResult
	err := ctx.Invoke("fortios:index/getSystemExternalresource:getSystemExternalresource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemExternalresource.
type LookupSystemExternalresourceArgs struct {
	// Specify the name of the desired system externalresource.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemExternalresource.
type LookupSystemExternalresourceResult struct {
	// User resource category.
	Category int `pulumi:"category"`
	// Comment.
	Comments string `pulumi:"comments"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specify outgoing interface to reach server.
	Interface string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server.
	InterfaceSelectMethod string `pulumi:"interfaceSelectMethod"`
	// External resource name.
	Name string `pulumi:"name"`
	// HTTP basic authentication password.
	Password string `pulumi:"password"`
	// Time interval to refresh external resource (1 - 43200 min, default = 5 min).
	RefreshRate int `pulumi:"refreshRate"`
	// URI of external resource.
	Resource string `pulumi:"resource"`
	// Source IPv4 address used to communicate with server.
	SourceIp string `pulumi:"sourceIp"`
	// Enable/disable user resource.
	Status string `pulumi:"status"`
	// User resource type.
	Type string `pulumi:"type"`
	// External resource update method.
	UpdateMethod string `pulumi:"updateMethod"`
	// Override HTTP User-Agent header used when retrieving this external resource.
	UserAgent string `pulumi:"userAgent"`
	// HTTP basic authentication user name.
	Username string `pulumi:"username"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid      string  `pulumi:"uuid"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemExternalresourceOutput(ctx *pulumi.Context, args LookupSystemExternalresourceOutputArgs, opts ...pulumi.InvokeOption) LookupSystemExternalresourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemExternalresourceResult, error) {
			args := v.(LookupSystemExternalresourceArgs)
			r, err := LookupSystemExternalresource(ctx, &args, opts...)
			var s LookupSystemExternalresourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemExternalresourceResultOutput)
}

// A collection of arguments for invoking getSystemExternalresource.
type LookupSystemExternalresourceOutputArgs struct {
	// Specify the name of the desired system externalresource.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemExternalresourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemExternalresourceArgs)(nil)).Elem()
}

// A collection of values returned by getSystemExternalresource.
type LookupSystemExternalresourceResultOutput struct{ *pulumi.OutputState }

func (LookupSystemExternalresourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemExternalresourceResult)(nil)).Elem()
}

func (o LookupSystemExternalresourceResultOutput) ToLookupSystemExternalresourceResultOutput() LookupSystemExternalresourceResultOutput {
	return o
}

func (o LookupSystemExternalresourceResultOutput) ToLookupSystemExternalresourceResultOutputWithContext(ctx context.Context) LookupSystemExternalresourceResultOutput {
	return o
}

// User resource category.
func (o LookupSystemExternalresourceResultOutput) Category() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) int { return v.Category }).(pulumi.IntOutput)
}

// Comment.
func (o LookupSystemExternalresourceResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) string { return v.Comments }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemExternalresourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specify outgoing interface to reach server.
func (o LookupSystemExternalresourceResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) string { return v.Interface }).(pulumi.StringOutput)
}

// Specify how to select outgoing interface to reach server.
func (o LookupSystemExternalresourceResultOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) string { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

// External resource name.
func (o LookupSystemExternalresourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// HTTP basic authentication password.
func (o LookupSystemExternalresourceResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) string { return v.Password }).(pulumi.StringOutput)
}

// Time interval to refresh external resource (1 - 43200 min, default = 5 min).
func (o LookupSystemExternalresourceResultOutput) RefreshRate() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) int { return v.RefreshRate }).(pulumi.IntOutput)
}

// URI of external resource.
func (o LookupSystemExternalresourceResultOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) string { return v.Resource }).(pulumi.StringOutput)
}

// Source IPv4 address used to communicate with server.
func (o LookupSystemExternalresourceResultOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) string { return v.SourceIp }).(pulumi.StringOutput)
}

// Enable/disable user resource.
func (o LookupSystemExternalresourceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) string { return v.Status }).(pulumi.StringOutput)
}

// User resource type.
func (o LookupSystemExternalresourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) string { return v.Type }).(pulumi.StringOutput)
}

// External resource update method.
func (o LookupSystemExternalresourceResultOutput) UpdateMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) string { return v.UpdateMethod }).(pulumi.StringOutput)
}

// Override HTTP User-Agent header used when retrieving this external resource.
func (o LookupSystemExternalresourceResultOutput) UserAgent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) string { return v.UserAgent }).(pulumi.StringOutput)
}

// HTTP basic authentication user name.
func (o LookupSystemExternalresourceResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) string { return v.Username }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o LookupSystemExternalresourceResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupSystemExternalresourceResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemExternalresourceResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemExternalresourceResultOutput{})
}