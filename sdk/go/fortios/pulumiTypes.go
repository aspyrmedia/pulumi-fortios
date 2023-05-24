// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetFirewallDoSpolicyAnomaly struct {
	Action           string `pulumi:"action"`
	Log              string `pulumi:"log"`
	Name             string `pulumi:"name"`
	Quarantine       string `pulumi:"quarantine"`
	QuarantineExpiry string `pulumi:"quarantineExpiry"`
	QuarantineLog    string `pulumi:"quarantineLog"`
	Status           string `pulumi:"status"`
	Threshold        int    `pulumi:"threshold"`
	Thresholddefault int    `pulumi:"thresholddefault"`
}

// GetFirewallDoSpolicyAnomalyInput is an input type that accepts GetFirewallDoSpolicyAnomalyArgs and GetFirewallDoSpolicyAnomalyOutput values.
// You can construct a concrete instance of `GetFirewallDoSpolicyAnomalyInput` via:
//
//	GetFirewallDoSpolicyAnomalyArgs{...}
type GetFirewallDoSpolicyAnomalyInput interface {
	pulumi.Input

	ToGetFirewallDoSpolicyAnomalyOutput() GetFirewallDoSpolicyAnomalyOutput
	ToGetFirewallDoSpolicyAnomalyOutputWithContext(context.Context) GetFirewallDoSpolicyAnomalyOutput
}

type GetFirewallDoSpolicyAnomalyArgs struct {
	Action           pulumi.StringInput `pulumi:"action"`
	Log              pulumi.StringInput `pulumi:"log"`
	Name             pulumi.StringInput `pulumi:"name"`
	Quarantine       pulumi.StringInput `pulumi:"quarantine"`
	QuarantineExpiry pulumi.StringInput `pulumi:"quarantineExpiry"`
	QuarantineLog    pulumi.StringInput `pulumi:"quarantineLog"`
	Status           pulumi.StringInput `pulumi:"status"`
	Threshold        pulumi.IntInput    `pulumi:"threshold"`
	Thresholddefault pulumi.IntInput    `pulumi:"thresholddefault"`
}

func (GetFirewallDoSpolicyAnomalyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallDoSpolicyAnomaly)(nil)).Elem()
}

func (i GetFirewallDoSpolicyAnomalyArgs) ToGetFirewallDoSpolicyAnomalyOutput() GetFirewallDoSpolicyAnomalyOutput {
	return i.ToGetFirewallDoSpolicyAnomalyOutputWithContext(context.Background())
}

func (i GetFirewallDoSpolicyAnomalyArgs) ToGetFirewallDoSpolicyAnomalyOutputWithContext(ctx context.Context) GetFirewallDoSpolicyAnomalyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetFirewallDoSpolicyAnomalyOutput)
}

// GetFirewallDoSpolicyAnomalyArrayInput is an input type that accepts GetFirewallDoSpolicyAnomalyArray and GetFirewallDoSpolicyAnomalyArrayOutput values.
// You can construct a concrete instance of `GetFirewallDoSpolicyAnomalyArrayInput` via:
//
//	GetFirewallDoSpolicyAnomalyArray{ GetFirewallDoSpolicyAnomalyArgs{...} }
type GetFirewallDoSpolicyAnomalyArrayInput interface {
	pulumi.Input

	ToGetFirewallDoSpolicyAnomalyArrayOutput() GetFirewallDoSpolicyAnomalyArrayOutput
	ToGetFirewallDoSpolicyAnomalyArrayOutputWithContext(context.Context) GetFirewallDoSpolicyAnomalyArrayOutput
}

type GetFirewallDoSpolicyAnomalyArray []GetFirewallDoSpolicyAnomalyInput

func (GetFirewallDoSpolicyAnomalyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetFirewallDoSpolicyAnomaly)(nil)).Elem()
}

func (i GetFirewallDoSpolicyAnomalyArray) ToGetFirewallDoSpolicyAnomalyArrayOutput() GetFirewallDoSpolicyAnomalyArrayOutput {
	return i.ToGetFirewallDoSpolicyAnomalyArrayOutputWithContext(context.Background())
}

func (i GetFirewallDoSpolicyAnomalyArray) ToGetFirewallDoSpolicyAnomalyArrayOutputWithContext(ctx context.Context) GetFirewallDoSpolicyAnomalyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetFirewallDoSpolicyAnomalyArrayOutput)
}

type GetFirewallDoSpolicyAnomalyOutput struct{ *pulumi.OutputState }

func (GetFirewallDoSpolicyAnomalyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallDoSpolicyAnomaly)(nil)).Elem()
}

func (o GetFirewallDoSpolicyAnomalyOutput) ToGetFirewallDoSpolicyAnomalyOutput() GetFirewallDoSpolicyAnomalyOutput {
	return o
}

func (o GetFirewallDoSpolicyAnomalyOutput) ToGetFirewallDoSpolicyAnomalyOutputWithContext(ctx context.Context) GetFirewallDoSpolicyAnomalyOutput {
	return o
}

func (o GetFirewallDoSpolicyAnomalyOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallDoSpolicyAnomaly) string { return v.Action }).(pulumi.StringOutput)
}

func (o GetFirewallDoSpolicyAnomalyOutput) Log() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallDoSpolicyAnomaly) string { return v.Log }).(pulumi.StringOutput)
}

func (o GetFirewallDoSpolicyAnomalyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallDoSpolicyAnomaly) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetFirewallDoSpolicyAnomalyOutput) Quarantine() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallDoSpolicyAnomaly) string { return v.Quarantine }).(pulumi.StringOutput)
}

func (o GetFirewallDoSpolicyAnomalyOutput) QuarantineExpiry() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallDoSpolicyAnomaly) string { return v.QuarantineExpiry }).(pulumi.StringOutput)
}

func (o GetFirewallDoSpolicyAnomalyOutput) QuarantineLog() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallDoSpolicyAnomaly) string { return v.QuarantineLog }).(pulumi.StringOutput)
}

func (o GetFirewallDoSpolicyAnomalyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallDoSpolicyAnomaly) string { return v.Status }).(pulumi.StringOutput)
}

func (o GetFirewallDoSpolicyAnomalyOutput) Threshold() pulumi.IntOutput {
	return o.ApplyT(func(v GetFirewallDoSpolicyAnomaly) int { return v.Threshold }).(pulumi.IntOutput)
}

func (o GetFirewallDoSpolicyAnomalyOutput) Thresholddefault() pulumi.IntOutput {
	return o.ApplyT(func(v GetFirewallDoSpolicyAnomaly) int { return v.Thresholddefault }).(pulumi.IntOutput)
}

type GetFirewallDoSpolicyAnomalyArrayOutput struct{ *pulumi.OutputState }

func (GetFirewallDoSpolicyAnomalyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetFirewallDoSpolicyAnomaly)(nil)).Elem()
}

func (o GetFirewallDoSpolicyAnomalyArrayOutput) ToGetFirewallDoSpolicyAnomalyArrayOutput() GetFirewallDoSpolicyAnomalyArrayOutput {
	return o
}

func (o GetFirewallDoSpolicyAnomalyArrayOutput) ToGetFirewallDoSpolicyAnomalyArrayOutputWithContext(ctx context.Context) GetFirewallDoSpolicyAnomalyArrayOutput {
	return o
}

func (o GetFirewallDoSpolicyAnomalyArrayOutput) Index(i pulumi.IntInput) GetFirewallDoSpolicyAnomalyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetFirewallDoSpolicyAnomaly {
		return vs[0].([]GetFirewallDoSpolicyAnomaly)[vs[1].(int)]
	}).(GetFirewallDoSpolicyAnomalyOutput)
}

type GetFirewallDoSpolicyDstaddr struct {
	Name string `pulumi:"name"`
}

// GetFirewallDoSpolicyDstaddrInput is an input type that accepts GetFirewallDoSpolicyDstaddrArgs and GetFirewallDoSpolicyDstaddrOutput values.
// You can construct a concrete instance of `GetFirewallDoSpolicyDstaddrInput` via:
//
//	GetFirewallDoSpolicyDstaddrArgs{...}
type GetFirewallDoSpolicyDstaddrInput interface {
	pulumi.Input

	ToGetFirewallDoSpolicyDstaddrOutput() GetFirewallDoSpolicyDstaddrOutput
	ToGetFirewallDoSpolicyDstaddrOutputWithContext(context.Context) GetFirewallDoSpolicyDstaddrOutput
}

type GetFirewallDoSpolicyDstaddrArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetFirewallDoSpolicyDstaddrArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallDoSpolicyDstaddr)(nil)).Elem()
}

func (i GetFirewallDoSpolicyDstaddrArgs) ToGetFirewallDoSpolicyDstaddrOutput() GetFirewallDoSpolicyDstaddrOutput {
	return i.ToGetFirewallDoSpolicyDstaddrOutputWithContext(context.Background())
}

func (i GetFirewallDoSpolicyDstaddrArgs) ToGetFirewallDoSpolicyDstaddrOutputWithContext(ctx context.Context) GetFirewallDoSpolicyDstaddrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetFirewallDoSpolicyDstaddrOutput)
}

// GetFirewallDoSpolicyDstaddrArrayInput is an input type that accepts GetFirewallDoSpolicyDstaddrArray and GetFirewallDoSpolicyDstaddrArrayOutput values.
// You can construct a concrete instance of `GetFirewallDoSpolicyDstaddrArrayInput` via:
//
//	GetFirewallDoSpolicyDstaddrArray{ GetFirewallDoSpolicyDstaddrArgs{...} }
type GetFirewallDoSpolicyDstaddrArrayInput interface {
	pulumi.Input

	ToGetFirewallDoSpolicyDstaddrArrayOutput() GetFirewallDoSpolicyDstaddrArrayOutput
	ToGetFirewallDoSpolicyDstaddrArrayOutputWithContext(context.Context) GetFirewallDoSpolicyDstaddrArrayOutput
}

type GetFirewallDoSpolicyDstaddrArray []GetFirewallDoSpolicyDstaddrInput

func (GetFirewallDoSpolicyDstaddrArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetFirewallDoSpolicyDstaddr)(nil)).Elem()
}

func (i GetFirewallDoSpolicyDstaddrArray) ToGetFirewallDoSpolicyDstaddrArrayOutput() GetFirewallDoSpolicyDstaddrArrayOutput {
	return i.ToGetFirewallDoSpolicyDstaddrArrayOutputWithContext(context.Background())
}

func (i GetFirewallDoSpolicyDstaddrArray) ToGetFirewallDoSpolicyDstaddrArrayOutputWithContext(ctx context.Context) GetFirewallDoSpolicyDstaddrArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetFirewallDoSpolicyDstaddrArrayOutput)
}

type GetFirewallDoSpolicyDstaddrOutput struct{ *pulumi.OutputState }

func (GetFirewallDoSpolicyDstaddrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallDoSpolicyDstaddr)(nil)).Elem()
}

func (o GetFirewallDoSpolicyDstaddrOutput) ToGetFirewallDoSpolicyDstaddrOutput() GetFirewallDoSpolicyDstaddrOutput {
	return o
}

func (o GetFirewallDoSpolicyDstaddrOutput) ToGetFirewallDoSpolicyDstaddrOutputWithContext(ctx context.Context) GetFirewallDoSpolicyDstaddrOutput {
	return o
}

func (o GetFirewallDoSpolicyDstaddrOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallDoSpolicyDstaddr) string { return v.Name }).(pulumi.StringOutput)
}

type GetFirewallDoSpolicyDstaddrArrayOutput struct{ *pulumi.OutputState }

func (GetFirewallDoSpolicyDstaddrArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetFirewallDoSpolicyDstaddr)(nil)).Elem()
}

func (o GetFirewallDoSpolicyDstaddrArrayOutput) ToGetFirewallDoSpolicyDstaddrArrayOutput() GetFirewallDoSpolicyDstaddrArrayOutput {
	return o
}

func (o GetFirewallDoSpolicyDstaddrArrayOutput) ToGetFirewallDoSpolicyDstaddrArrayOutputWithContext(ctx context.Context) GetFirewallDoSpolicyDstaddrArrayOutput {
	return o
}

func (o GetFirewallDoSpolicyDstaddrArrayOutput) Index(i pulumi.IntInput) GetFirewallDoSpolicyDstaddrOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetFirewallDoSpolicyDstaddr {
		return vs[0].([]GetFirewallDoSpolicyDstaddr)[vs[1].(int)]
	}).(GetFirewallDoSpolicyDstaddrOutput)
}

type GetFirewallDoSpolicyService struct {
	Name string `pulumi:"name"`
}

// GetFirewallDoSpolicyServiceInput is an input type that accepts GetFirewallDoSpolicyServiceArgs and GetFirewallDoSpolicyServiceOutput values.
// You can construct a concrete instance of `GetFirewallDoSpolicyServiceInput` via:
//
//	GetFirewallDoSpolicyServiceArgs{...}
type GetFirewallDoSpolicyServiceInput interface {
	pulumi.Input

	ToGetFirewallDoSpolicyServiceOutput() GetFirewallDoSpolicyServiceOutput
	ToGetFirewallDoSpolicyServiceOutputWithContext(context.Context) GetFirewallDoSpolicyServiceOutput
}

type GetFirewallDoSpolicyServiceArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetFirewallDoSpolicyServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallDoSpolicyService)(nil)).Elem()
}

func (i GetFirewallDoSpolicyServiceArgs) ToGetFirewallDoSpolicyServiceOutput() GetFirewallDoSpolicyServiceOutput {
	return i.ToGetFirewallDoSpolicyServiceOutputWithContext(context.Background())
}

func (i GetFirewallDoSpolicyServiceArgs) ToGetFirewallDoSpolicyServiceOutputWithContext(ctx context.Context) GetFirewallDoSpolicyServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetFirewallDoSpolicyServiceOutput)
}

// GetFirewallDoSpolicyServiceArrayInput is an input type that accepts GetFirewallDoSpolicyServiceArray and GetFirewallDoSpolicyServiceArrayOutput values.
// You can construct a concrete instance of `GetFirewallDoSpolicyServiceArrayInput` via:
//
//	GetFirewallDoSpolicyServiceArray{ GetFirewallDoSpolicyServiceArgs{...} }
type GetFirewallDoSpolicyServiceArrayInput interface {
	pulumi.Input

	ToGetFirewallDoSpolicyServiceArrayOutput() GetFirewallDoSpolicyServiceArrayOutput
	ToGetFirewallDoSpolicyServiceArrayOutputWithContext(context.Context) GetFirewallDoSpolicyServiceArrayOutput
}

type GetFirewallDoSpolicyServiceArray []GetFirewallDoSpolicyServiceInput

func (GetFirewallDoSpolicyServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetFirewallDoSpolicyService)(nil)).Elem()
}

func (i GetFirewallDoSpolicyServiceArray) ToGetFirewallDoSpolicyServiceArrayOutput() GetFirewallDoSpolicyServiceArrayOutput {
	return i.ToGetFirewallDoSpolicyServiceArrayOutputWithContext(context.Background())
}

func (i GetFirewallDoSpolicyServiceArray) ToGetFirewallDoSpolicyServiceArrayOutputWithContext(ctx context.Context) GetFirewallDoSpolicyServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetFirewallDoSpolicyServiceArrayOutput)
}

type GetFirewallDoSpolicyServiceOutput struct{ *pulumi.OutputState }

func (GetFirewallDoSpolicyServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallDoSpolicyService)(nil)).Elem()
}

func (o GetFirewallDoSpolicyServiceOutput) ToGetFirewallDoSpolicyServiceOutput() GetFirewallDoSpolicyServiceOutput {
	return o
}

func (o GetFirewallDoSpolicyServiceOutput) ToGetFirewallDoSpolicyServiceOutputWithContext(ctx context.Context) GetFirewallDoSpolicyServiceOutput {
	return o
}

func (o GetFirewallDoSpolicyServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallDoSpolicyService) string { return v.Name }).(pulumi.StringOutput)
}

type GetFirewallDoSpolicyServiceArrayOutput struct{ *pulumi.OutputState }

func (GetFirewallDoSpolicyServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetFirewallDoSpolicyService)(nil)).Elem()
}

func (o GetFirewallDoSpolicyServiceArrayOutput) ToGetFirewallDoSpolicyServiceArrayOutput() GetFirewallDoSpolicyServiceArrayOutput {
	return o
}

func (o GetFirewallDoSpolicyServiceArrayOutput) ToGetFirewallDoSpolicyServiceArrayOutputWithContext(ctx context.Context) GetFirewallDoSpolicyServiceArrayOutput {
	return o
}

func (o GetFirewallDoSpolicyServiceArrayOutput) Index(i pulumi.IntInput) GetFirewallDoSpolicyServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetFirewallDoSpolicyService {
		return vs[0].([]GetFirewallDoSpolicyService)[vs[1].(int)]
	}).(GetFirewallDoSpolicyServiceOutput)
}

type GetFirewallDoSpolicySrcaddr struct {
	Name string `pulumi:"name"`
}

// GetFirewallDoSpolicySrcaddrInput is an input type that accepts GetFirewallDoSpolicySrcaddrArgs and GetFirewallDoSpolicySrcaddrOutput values.
// You can construct a concrete instance of `GetFirewallDoSpolicySrcaddrInput` via:
//
//	GetFirewallDoSpolicySrcaddrArgs{...}
type GetFirewallDoSpolicySrcaddrInput interface {
	pulumi.Input

	ToGetFirewallDoSpolicySrcaddrOutput() GetFirewallDoSpolicySrcaddrOutput
	ToGetFirewallDoSpolicySrcaddrOutputWithContext(context.Context) GetFirewallDoSpolicySrcaddrOutput
}

type GetFirewallDoSpolicySrcaddrArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetFirewallDoSpolicySrcaddrArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallDoSpolicySrcaddr)(nil)).Elem()
}

func (i GetFirewallDoSpolicySrcaddrArgs) ToGetFirewallDoSpolicySrcaddrOutput() GetFirewallDoSpolicySrcaddrOutput {
	return i.ToGetFirewallDoSpolicySrcaddrOutputWithContext(context.Background())
}

func (i GetFirewallDoSpolicySrcaddrArgs) ToGetFirewallDoSpolicySrcaddrOutputWithContext(ctx context.Context) GetFirewallDoSpolicySrcaddrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetFirewallDoSpolicySrcaddrOutput)
}

// GetFirewallDoSpolicySrcaddrArrayInput is an input type that accepts GetFirewallDoSpolicySrcaddrArray and GetFirewallDoSpolicySrcaddrArrayOutput values.
// You can construct a concrete instance of `GetFirewallDoSpolicySrcaddrArrayInput` via:
//
//	GetFirewallDoSpolicySrcaddrArray{ GetFirewallDoSpolicySrcaddrArgs{...} }
type GetFirewallDoSpolicySrcaddrArrayInput interface {
	pulumi.Input

	ToGetFirewallDoSpolicySrcaddrArrayOutput() GetFirewallDoSpolicySrcaddrArrayOutput
	ToGetFirewallDoSpolicySrcaddrArrayOutputWithContext(context.Context) GetFirewallDoSpolicySrcaddrArrayOutput
}

type GetFirewallDoSpolicySrcaddrArray []GetFirewallDoSpolicySrcaddrInput

func (GetFirewallDoSpolicySrcaddrArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetFirewallDoSpolicySrcaddr)(nil)).Elem()
}

func (i GetFirewallDoSpolicySrcaddrArray) ToGetFirewallDoSpolicySrcaddrArrayOutput() GetFirewallDoSpolicySrcaddrArrayOutput {
	return i.ToGetFirewallDoSpolicySrcaddrArrayOutputWithContext(context.Background())
}

func (i GetFirewallDoSpolicySrcaddrArray) ToGetFirewallDoSpolicySrcaddrArrayOutputWithContext(ctx context.Context) GetFirewallDoSpolicySrcaddrArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetFirewallDoSpolicySrcaddrArrayOutput)
}

type GetFirewallDoSpolicySrcaddrOutput struct{ *pulumi.OutputState }

func (GetFirewallDoSpolicySrcaddrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallDoSpolicySrcaddr)(nil)).Elem()
}

func (o GetFirewallDoSpolicySrcaddrOutput) ToGetFirewallDoSpolicySrcaddrOutput() GetFirewallDoSpolicySrcaddrOutput {
	return o
}

func (o GetFirewallDoSpolicySrcaddrOutput) ToGetFirewallDoSpolicySrcaddrOutputWithContext(ctx context.Context) GetFirewallDoSpolicySrcaddrOutput {
	return o
}

func (o GetFirewallDoSpolicySrcaddrOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallDoSpolicySrcaddr) string { return v.Name }).(pulumi.StringOutput)
}

type GetFirewallDoSpolicySrcaddrArrayOutput struct{ *pulumi.OutputState }

func (GetFirewallDoSpolicySrcaddrArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetFirewallDoSpolicySrcaddr)(nil)).Elem()
}

func (o GetFirewallDoSpolicySrcaddrArrayOutput) ToGetFirewallDoSpolicySrcaddrArrayOutput() GetFirewallDoSpolicySrcaddrArrayOutput {
	return o
}

func (o GetFirewallDoSpolicySrcaddrArrayOutput) ToGetFirewallDoSpolicySrcaddrArrayOutputWithContext(ctx context.Context) GetFirewallDoSpolicySrcaddrArrayOutput {
	return o
}

func (o GetFirewallDoSpolicySrcaddrArrayOutput) Index(i pulumi.IntInput) GetFirewallDoSpolicySrcaddrOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetFirewallDoSpolicySrcaddr {
		return vs[0].([]GetFirewallDoSpolicySrcaddr)[vs[1].(int)]
	}).(GetFirewallDoSpolicySrcaddrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetFirewallDoSpolicyAnomalyInput)(nil)).Elem(), GetFirewallDoSpolicyAnomalyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetFirewallDoSpolicyAnomalyArrayInput)(nil)).Elem(), GetFirewallDoSpolicyAnomalyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetFirewallDoSpolicyDstaddrInput)(nil)).Elem(), GetFirewallDoSpolicyDstaddrArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetFirewallDoSpolicyDstaddrArrayInput)(nil)).Elem(), GetFirewallDoSpolicyDstaddrArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetFirewallDoSpolicyServiceInput)(nil)).Elem(), GetFirewallDoSpolicyServiceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetFirewallDoSpolicyServiceArrayInput)(nil)).Elem(), GetFirewallDoSpolicyServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetFirewallDoSpolicySrcaddrInput)(nil)).Elem(), GetFirewallDoSpolicySrcaddrArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetFirewallDoSpolicySrcaddrArrayInput)(nil)).Elem(), GetFirewallDoSpolicySrcaddrArray{})
	pulumi.RegisterOutputType(GetFirewallDoSpolicyAnomalyOutput{})
	pulumi.RegisterOutputType(GetFirewallDoSpolicyAnomalyArrayOutput{})
	pulumi.RegisterOutputType(GetFirewallDoSpolicyDstaddrOutput{})
	pulumi.RegisterOutputType(GetFirewallDoSpolicyDstaddrArrayOutput{})
	pulumi.RegisterOutputType(GetFirewallDoSpolicyServiceOutput{})
	pulumi.RegisterOutputType(GetFirewallDoSpolicyServiceArrayOutput{})
	pulumi.RegisterOutputType(GetFirewallDoSpolicySrcaddrOutput{})
	pulumi.RegisterOutputType(GetFirewallDoSpolicySrcaddrArrayOutput{})
}
