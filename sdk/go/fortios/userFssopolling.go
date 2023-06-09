// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FSSO active directory servers for polling mode.
//
// ## Import
//
// # User FssoPolling can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/userFssopolling:UserFssopolling labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/userFssopolling:UserFssopolling labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type UserFssopolling struct {
	pulumi.CustomResourceState

	// LDAP Group Info. The structure of `adgrp` block is documented below.
	Adgrps UserFssopollingAdgrpArrayOutput `pulumi:"adgrps"`
	// Default domain managed by this Active Directory server.
	DefaultDomain pulumi.StringOutput `pulumi:"defaultDomain"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Active Directory server ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// LDAP server name used in LDAP connection strings.
	LdapServer pulumi.StringOutput `pulumi:"ldapServer"`
	// Number of hours of logon history to keep, 0 means keep all history.
	LogonHistory pulumi.IntOutput `pulumi:"logonHistory"`
	// Password required to log into this Active Directory server
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Polling frequency (every 1 to 30 seconds).
	PollingFrequency pulumi.IntOutput `pulumi:"pollingFrequency"`
	// Port to communicate with this Active Directory server.
	Port pulumi.IntOutput `pulumi:"port"`
	// Host name or IP address of the Active Directory server.
	Server pulumi.StringOutput `pulumi:"server"`
	// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
	SmbNtlmv1Auth pulumi.StringOutput `pulumi:"smbNtlmv1Auth"`
	// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
	Smbv1 pulumi.StringOutput `pulumi:"smbv1"`
	// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// User name required to log into this Active Directory server.
	User pulumi.StringOutput `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewUserFssopolling registers a new resource with the given unique name, arguments, and options.
func NewUserFssopolling(ctx *pulumi.Context,
	name string, args *UserFssopollingArgs, opts ...pulumi.ResourceOption) (*UserFssopolling, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LdapServer == nil {
		return nil, errors.New("invalid value for required argument 'LdapServer'")
	}
	if args.Server == nil {
		return nil, errors.New("invalid value for required argument 'Server'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource UserFssopolling
	err := ctx.RegisterResource("fortios:index/userFssopolling:UserFssopolling", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserFssopolling gets an existing UserFssopolling resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserFssopolling(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserFssopollingState, opts ...pulumi.ResourceOption) (*UserFssopolling, error) {
	var resource UserFssopolling
	err := ctx.ReadResource("fortios:index/userFssopolling:UserFssopolling", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserFssopolling resources.
type userFssopollingState struct {
	// LDAP Group Info. The structure of `adgrp` block is documented below.
	Adgrps []UserFssopollingAdgrp `pulumi:"adgrps"`
	// Default domain managed by this Active Directory server.
	DefaultDomain *string `pulumi:"defaultDomain"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Active Directory server ID.
	Fosid *int `pulumi:"fosid"`
	// LDAP server name used in LDAP connection strings.
	LdapServer *string `pulumi:"ldapServer"`
	// Number of hours of logon history to keep, 0 means keep all history.
	LogonHistory *int `pulumi:"logonHistory"`
	// Password required to log into this Active Directory server
	Password *string `pulumi:"password"`
	// Polling frequency (every 1 to 30 seconds).
	PollingFrequency *int `pulumi:"pollingFrequency"`
	// Port to communicate with this Active Directory server.
	Port *int `pulumi:"port"`
	// Host name or IP address of the Active Directory server.
	Server *string `pulumi:"server"`
	// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
	SmbNtlmv1Auth *string `pulumi:"smbNtlmv1Auth"`
	// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
	Smbv1 *string `pulumi:"smbv1"`
	// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// User name required to log into this Active Directory server.
	User *string `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type UserFssopollingState struct {
	// LDAP Group Info. The structure of `adgrp` block is documented below.
	Adgrps UserFssopollingAdgrpArrayInput
	// Default domain managed by this Active Directory server.
	DefaultDomain pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Active Directory server ID.
	Fosid pulumi.IntPtrInput
	// LDAP server name used in LDAP connection strings.
	LdapServer pulumi.StringPtrInput
	// Number of hours of logon history to keep, 0 means keep all history.
	LogonHistory pulumi.IntPtrInput
	// Password required to log into this Active Directory server
	Password pulumi.StringPtrInput
	// Polling frequency (every 1 to 30 seconds).
	PollingFrequency pulumi.IntPtrInput
	// Port to communicate with this Active Directory server.
	Port pulumi.IntPtrInput
	// Host name or IP address of the Active Directory server.
	Server pulumi.StringPtrInput
	// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
	SmbNtlmv1Auth pulumi.StringPtrInput
	// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
	Smbv1 pulumi.StringPtrInput
	// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// User name required to log into this Active Directory server.
	User pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserFssopollingState) ElementType() reflect.Type {
	return reflect.TypeOf((*userFssopollingState)(nil)).Elem()
}

type userFssopollingArgs struct {
	// LDAP Group Info. The structure of `adgrp` block is documented below.
	Adgrps []UserFssopollingAdgrp `pulumi:"adgrps"`
	// Default domain managed by this Active Directory server.
	DefaultDomain *string `pulumi:"defaultDomain"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Active Directory server ID.
	Fosid *int `pulumi:"fosid"`
	// LDAP server name used in LDAP connection strings.
	LdapServer string `pulumi:"ldapServer"`
	// Number of hours of logon history to keep, 0 means keep all history.
	LogonHistory *int `pulumi:"logonHistory"`
	// Password required to log into this Active Directory server
	Password *string `pulumi:"password"`
	// Polling frequency (every 1 to 30 seconds).
	PollingFrequency *int `pulumi:"pollingFrequency"`
	// Port to communicate with this Active Directory server.
	Port *int `pulumi:"port"`
	// Host name or IP address of the Active Directory server.
	Server string `pulumi:"server"`
	// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
	SmbNtlmv1Auth *string `pulumi:"smbNtlmv1Auth"`
	// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
	Smbv1 *string `pulumi:"smbv1"`
	// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// User name required to log into this Active Directory server.
	User string `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserFssopolling resource.
type UserFssopollingArgs struct {
	// LDAP Group Info. The structure of `adgrp` block is documented below.
	Adgrps UserFssopollingAdgrpArrayInput
	// Default domain managed by this Active Directory server.
	DefaultDomain pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Active Directory server ID.
	Fosid pulumi.IntPtrInput
	// LDAP server name used in LDAP connection strings.
	LdapServer pulumi.StringInput
	// Number of hours of logon history to keep, 0 means keep all history.
	LogonHistory pulumi.IntPtrInput
	// Password required to log into this Active Directory server
	Password pulumi.StringPtrInput
	// Polling frequency (every 1 to 30 seconds).
	PollingFrequency pulumi.IntPtrInput
	// Port to communicate with this Active Directory server.
	Port pulumi.IntPtrInput
	// Host name or IP address of the Active Directory server.
	Server pulumi.StringInput
	// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
	SmbNtlmv1Auth pulumi.StringPtrInput
	// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
	Smbv1 pulumi.StringPtrInput
	// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// User name required to log into this Active Directory server.
	User pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserFssopollingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userFssopollingArgs)(nil)).Elem()
}

type UserFssopollingInput interface {
	pulumi.Input

	ToUserFssopollingOutput() UserFssopollingOutput
	ToUserFssopollingOutputWithContext(ctx context.Context) UserFssopollingOutput
}

func (*UserFssopolling) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFssopolling)(nil)).Elem()
}

func (i *UserFssopolling) ToUserFssopollingOutput() UserFssopollingOutput {
	return i.ToUserFssopollingOutputWithContext(context.Background())
}

func (i *UserFssopolling) ToUserFssopollingOutputWithContext(ctx context.Context) UserFssopollingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFssopollingOutput)
}

// UserFssopollingArrayInput is an input type that accepts UserFssopollingArray and UserFssopollingArrayOutput values.
// You can construct a concrete instance of `UserFssopollingArrayInput` via:
//
//	UserFssopollingArray{ UserFssopollingArgs{...} }
type UserFssopollingArrayInput interface {
	pulumi.Input

	ToUserFssopollingArrayOutput() UserFssopollingArrayOutput
	ToUserFssopollingArrayOutputWithContext(context.Context) UserFssopollingArrayOutput
}

type UserFssopollingArray []UserFssopollingInput

func (UserFssopollingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserFssopolling)(nil)).Elem()
}

func (i UserFssopollingArray) ToUserFssopollingArrayOutput() UserFssopollingArrayOutput {
	return i.ToUserFssopollingArrayOutputWithContext(context.Background())
}

func (i UserFssopollingArray) ToUserFssopollingArrayOutputWithContext(ctx context.Context) UserFssopollingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFssopollingArrayOutput)
}

// UserFssopollingMapInput is an input type that accepts UserFssopollingMap and UserFssopollingMapOutput values.
// You can construct a concrete instance of `UserFssopollingMapInput` via:
//
//	UserFssopollingMap{ "key": UserFssopollingArgs{...} }
type UserFssopollingMapInput interface {
	pulumi.Input

	ToUserFssopollingMapOutput() UserFssopollingMapOutput
	ToUserFssopollingMapOutputWithContext(context.Context) UserFssopollingMapOutput
}

type UserFssopollingMap map[string]UserFssopollingInput

func (UserFssopollingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserFssopolling)(nil)).Elem()
}

func (i UserFssopollingMap) ToUserFssopollingMapOutput() UserFssopollingMapOutput {
	return i.ToUserFssopollingMapOutputWithContext(context.Background())
}

func (i UserFssopollingMap) ToUserFssopollingMapOutputWithContext(ctx context.Context) UserFssopollingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFssopollingMapOutput)
}

type UserFssopollingOutput struct{ *pulumi.OutputState }

func (UserFssopollingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFssopolling)(nil)).Elem()
}

func (o UserFssopollingOutput) ToUserFssopollingOutput() UserFssopollingOutput {
	return o
}

func (o UserFssopollingOutput) ToUserFssopollingOutputWithContext(ctx context.Context) UserFssopollingOutput {
	return o
}

// LDAP Group Info. The structure of `adgrp` block is documented below.
func (o UserFssopollingOutput) Adgrps() UserFssopollingAdgrpArrayOutput {
	return o.ApplyT(func(v *UserFssopolling) UserFssopollingAdgrpArrayOutput { return v.Adgrps }).(UserFssopollingAdgrpArrayOutput)
}

// Default domain managed by this Active Directory server.
func (o UserFssopollingOutput) DefaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFssopolling) pulumi.StringOutput { return v.DefaultDomain }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o UserFssopollingOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserFssopolling) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Active Directory server ID.
func (o UserFssopollingOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *UserFssopolling) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// LDAP server name used in LDAP connection strings.
func (o UserFssopollingOutput) LdapServer() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFssopolling) pulumi.StringOutput { return v.LdapServer }).(pulumi.StringOutput)
}

// Number of hours of logon history to keep, 0 means keep all history.
func (o UserFssopollingOutput) LogonHistory() pulumi.IntOutput {
	return o.ApplyT(func(v *UserFssopolling) pulumi.IntOutput { return v.LogonHistory }).(pulumi.IntOutput)
}

// Password required to log into this Active Directory server
func (o UserFssopollingOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserFssopolling) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Polling frequency (every 1 to 30 seconds).
func (o UserFssopollingOutput) PollingFrequency() pulumi.IntOutput {
	return o.ApplyT(func(v *UserFssopolling) pulumi.IntOutput { return v.PollingFrequency }).(pulumi.IntOutput)
}

// Port to communicate with this Active Directory server.
func (o UserFssopollingOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *UserFssopolling) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Host name or IP address of the Active Directory server.
func (o UserFssopollingOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFssopolling) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
func (o UserFssopollingOutput) SmbNtlmv1Auth() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFssopolling) pulumi.StringOutput { return v.SmbNtlmv1Auth }).(pulumi.StringOutput)
}

// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
func (o UserFssopollingOutput) Smbv1() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFssopolling) pulumi.StringOutput { return v.Smbv1 }).(pulumi.StringOutput)
}

// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
func (o UserFssopollingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFssopolling) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// User name required to log into this Active Directory server.
func (o UserFssopollingOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *UserFssopolling) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o UserFssopollingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserFssopolling) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type UserFssopollingArrayOutput struct{ *pulumi.OutputState }

func (UserFssopollingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserFssopolling)(nil)).Elem()
}

func (o UserFssopollingArrayOutput) ToUserFssopollingArrayOutput() UserFssopollingArrayOutput {
	return o
}

func (o UserFssopollingArrayOutput) ToUserFssopollingArrayOutputWithContext(ctx context.Context) UserFssopollingArrayOutput {
	return o
}

func (o UserFssopollingArrayOutput) Index(i pulumi.IntInput) UserFssopollingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserFssopolling {
		return vs[0].([]*UserFssopolling)[vs[1].(int)]
	}).(UserFssopollingOutput)
}

type UserFssopollingMapOutput struct{ *pulumi.OutputState }

func (UserFssopollingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserFssopolling)(nil)).Elem()
}

func (o UserFssopollingMapOutput) ToUserFssopollingMapOutput() UserFssopollingMapOutput {
	return o
}

func (o UserFssopollingMapOutput) ToUserFssopollingMapOutputWithContext(ctx context.Context) UserFssopollingMapOutput {
	return o
}

func (o UserFssopollingMapOutput) MapIndex(k pulumi.StringInput) UserFssopollingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserFssopolling {
		return vs[0].(map[string]*UserFssopolling)[vs[1].(string)]
	}).(UserFssopollingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserFssopollingInput)(nil)).Elem(), &UserFssopolling{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserFssopollingArrayInput)(nil)).Elem(), UserFssopollingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserFssopollingMapInput)(nil)).Elem(), UserFssopollingMap{})
	pulumi.RegisterOutputType(UserFssopollingOutput{})
	pulumi.RegisterOutputType(UserFssopollingArrayOutput{})
	pulumi.RegisterOutputType(UserFssopollingMapOutput{})
}
