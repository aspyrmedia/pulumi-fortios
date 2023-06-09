// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiGuard - AntiSpam. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// # Emailfilter Fortishield can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/emailfilterFortishield:EmailfilterFortishield labelname EmailfilterFortishield
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/emailfilterFortishield:EmailfilterFortishield labelname EmailfilterFortishield
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type EmailfilterFortishield struct {
	pulumi.CustomResourceState

	// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
	SpamSubmitForce pulumi.StringOutput `pulumi:"spamSubmitForce"`
	// Hostname of the spam submission server.
	SpamSubmitSrv pulumi.StringOutput `pulumi:"spamSubmitSrv"`
	// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
	SpamSubmitTxt2htm pulumi.StringOutput `pulumi:"spamSubmitTxt2htm"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewEmailfilterFortishield registers a new resource with the given unique name, arguments, and options.
func NewEmailfilterFortishield(ctx *pulumi.Context,
	name string, args *EmailfilterFortishieldArgs, opts ...pulumi.ResourceOption) (*EmailfilterFortishield, error) {
	if args == nil {
		args = &EmailfilterFortishieldArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource EmailfilterFortishield
	err := ctx.RegisterResource("fortios:index/emailfilterFortishield:EmailfilterFortishield", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailfilterFortishield gets an existing EmailfilterFortishield resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailfilterFortishield(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailfilterFortishieldState, opts ...pulumi.ResourceOption) (*EmailfilterFortishield, error) {
	var resource EmailfilterFortishield
	err := ctx.ReadResource("fortios:index/emailfilterFortishield:EmailfilterFortishield", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailfilterFortishield resources.
type emailfilterFortishieldState struct {
	// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
	SpamSubmitForce *string `pulumi:"spamSubmitForce"`
	// Hostname of the spam submission server.
	SpamSubmitSrv *string `pulumi:"spamSubmitSrv"`
	// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
	SpamSubmitTxt2htm *string `pulumi:"spamSubmitTxt2htm"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type EmailfilterFortishieldState struct {
	// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
	SpamSubmitForce pulumi.StringPtrInput
	// Hostname of the spam submission server.
	SpamSubmitSrv pulumi.StringPtrInput
	// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
	SpamSubmitTxt2htm pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EmailfilterFortishieldState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailfilterFortishieldState)(nil)).Elem()
}

type emailfilterFortishieldArgs struct {
	// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
	SpamSubmitForce *string `pulumi:"spamSubmitForce"`
	// Hostname of the spam submission server.
	SpamSubmitSrv *string `pulumi:"spamSubmitSrv"`
	// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
	SpamSubmitTxt2htm *string `pulumi:"spamSubmitTxt2htm"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a EmailfilterFortishield resource.
type EmailfilterFortishieldArgs struct {
	// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
	SpamSubmitForce pulumi.StringPtrInput
	// Hostname of the spam submission server.
	SpamSubmitSrv pulumi.StringPtrInput
	// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
	SpamSubmitTxt2htm pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EmailfilterFortishieldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailfilterFortishieldArgs)(nil)).Elem()
}

type EmailfilterFortishieldInput interface {
	pulumi.Input

	ToEmailfilterFortishieldOutput() EmailfilterFortishieldOutput
	ToEmailfilterFortishieldOutputWithContext(ctx context.Context) EmailfilterFortishieldOutput
}

func (*EmailfilterFortishield) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailfilterFortishield)(nil)).Elem()
}

func (i *EmailfilterFortishield) ToEmailfilterFortishieldOutput() EmailfilterFortishieldOutput {
	return i.ToEmailfilterFortishieldOutputWithContext(context.Background())
}

func (i *EmailfilterFortishield) ToEmailfilterFortishieldOutputWithContext(ctx context.Context) EmailfilterFortishieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterFortishieldOutput)
}

// EmailfilterFortishieldArrayInput is an input type that accepts EmailfilterFortishieldArray and EmailfilterFortishieldArrayOutput values.
// You can construct a concrete instance of `EmailfilterFortishieldArrayInput` via:
//
//	EmailfilterFortishieldArray{ EmailfilterFortishieldArgs{...} }
type EmailfilterFortishieldArrayInput interface {
	pulumi.Input

	ToEmailfilterFortishieldArrayOutput() EmailfilterFortishieldArrayOutput
	ToEmailfilterFortishieldArrayOutputWithContext(context.Context) EmailfilterFortishieldArrayOutput
}

type EmailfilterFortishieldArray []EmailfilterFortishieldInput

func (EmailfilterFortishieldArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailfilterFortishield)(nil)).Elem()
}

func (i EmailfilterFortishieldArray) ToEmailfilterFortishieldArrayOutput() EmailfilterFortishieldArrayOutput {
	return i.ToEmailfilterFortishieldArrayOutputWithContext(context.Background())
}

func (i EmailfilterFortishieldArray) ToEmailfilterFortishieldArrayOutputWithContext(ctx context.Context) EmailfilterFortishieldArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterFortishieldArrayOutput)
}

// EmailfilterFortishieldMapInput is an input type that accepts EmailfilterFortishieldMap and EmailfilterFortishieldMapOutput values.
// You can construct a concrete instance of `EmailfilterFortishieldMapInput` via:
//
//	EmailfilterFortishieldMap{ "key": EmailfilterFortishieldArgs{...} }
type EmailfilterFortishieldMapInput interface {
	pulumi.Input

	ToEmailfilterFortishieldMapOutput() EmailfilterFortishieldMapOutput
	ToEmailfilterFortishieldMapOutputWithContext(context.Context) EmailfilterFortishieldMapOutput
}

type EmailfilterFortishieldMap map[string]EmailfilterFortishieldInput

func (EmailfilterFortishieldMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailfilterFortishield)(nil)).Elem()
}

func (i EmailfilterFortishieldMap) ToEmailfilterFortishieldMapOutput() EmailfilterFortishieldMapOutput {
	return i.ToEmailfilterFortishieldMapOutputWithContext(context.Background())
}

func (i EmailfilterFortishieldMap) ToEmailfilterFortishieldMapOutputWithContext(ctx context.Context) EmailfilterFortishieldMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterFortishieldMapOutput)
}

type EmailfilterFortishieldOutput struct{ *pulumi.OutputState }

func (EmailfilterFortishieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailfilterFortishield)(nil)).Elem()
}

func (o EmailfilterFortishieldOutput) ToEmailfilterFortishieldOutput() EmailfilterFortishieldOutput {
	return o
}

func (o EmailfilterFortishieldOutput) ToEmailfilterFortishieldOutputWithContext(ctx context.Context) EmailfilterFortishieldOutput {
	return o
}

// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
func (o EmailfilterFortishieldOutput) SpamSubmitForce() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailfilterFortishield) pulumi.StringOutput { return v.SpamSubmitForce }).(pulumi.StringOutput)
}

// Hostname of the spam submission server.
func (o EmailfilterFortishieldOutput) SpamSubmitSrv() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailfilterFortishield) pulumi.StringOutput { return v.SpamSubmitSrv }).(pulumi.StringOutput)
}

// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
func (o EmailfilterFortishieldOutput) SpamSubmitTxt2htm() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailfilterFortishield) pulumi.StringOutput { return v.SpamSubmitTxt2htm }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o EmailfilterFortishieldOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailfilterFortishield) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type EmailfilterFortishieldArrayOutput struct{ *pulumi.OutputState }

func (EmailfilterFortishieldArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailfilterFortishield)(nil)).Elem()
}

func (o EmailfilterFortishieldArrayOutput) ToEmailfilterFortishieldArrayOutput() EmailfilterFortishieldArrayOutput {
	return o
}

func (o EmailfilterFortishieldArrayOutput) ToEmailfilterFortishieldArrayOutputWithContext(ctx context.Context) EmailfilterFortishieldArrayOutput {
	return o
}

func (o EmailfilterFortishieldArrayOutput) Index(i pulumi.IntInput) EmailfilterFortishieldOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EmailfilterFortishield {
		return vs[0].([]*EmailfilterFortishield)[vs[1].(int)]
	}).(EmailfilterFortishieldOutput)
}

type EmailfilterFortishieldMapOutput struct{ *pulumi.OutputState }

func (EmailfilterFortishieldMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailfilterFortishield)(nil)).Elem()
}

func (o EmailfilterFortishieldMapOutput) ToEmailfilterFortishieldMapOutput() EmailfilterFortishieldMapOutput {
	return o
}

func (o EmailfilterFortishieldMapOutput) ToEmailfilterFortishieldMapOutputWithContext(ctx context.Context) EmailfilterFortishieldMapOutput {
	return o
}

func (o EmailfilterFortishieldMapOutput) MapIndex(k pulumi.StringInput) EmailfilterFortishieldOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EmailfilterFortishield {
		return vs[0].(map[string]*EmailfilterFortishield)[vs[1].(string)]
	}).(EmailfilterFortishieldOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterFortishieldInput)(nil)).Elem(), &EmailfilterFortishield{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterFortishieldArrayInput)(nil)).Elem(), EmailfilterFortishieldArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterFortishieldMapInput)(nil)).Elem(), EmailfilterFortishieldMap{})
	pulumi.RegisterOutputType(EmailfilterFortishieldOutput{})
	pulumi.RegisterOutputType(EmailfilterFortishieldArrayOutput{})
	pulumi.RegisterOutputType(EmailfilterFortishieldMapOutput{})
}
