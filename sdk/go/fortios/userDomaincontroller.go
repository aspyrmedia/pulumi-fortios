// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure domain controller entries.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			trname1, err := fortios.NewUserLdap(ctx, "trname1", &fortios.UserLdapArgs{
//				AccountKeyFilter:      pulumi.String("(&(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))"),
//				AccountKeyProcessing:  pulumi.String("same"),
//				Cnid:                  pulumi.String("cn"),
//				Dn:                    pulumi.String("EIWNCIEW"),
//				GroupMemberCheck:      pulumi.String("user-attr"),
//				GroupObjectFilter:     pulumi.String("(&(objectcategory=group)(member=*))"),
//				MemberAttr:            pulumi.String("memberOf"),
//				PasswordExpiryWarning: pulumi.String("disable"),
//				PasswordRenewal:       pulumi.String("disable"),
//				Port:                  pulumi.Int(389),
//				Secure:                pulumi.String("disable"),
//				Server:                pulumi.String("1.1.1.1"),
//				ServerIdentityCheck:   pulumi.String("disable"),
//				SourceIp:              pulumi.String("0.0.0.0"),
//				SslMinProtoVersion:    pulumi.String("default"),
//				Type:                  pulumi.String("simple"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = fortios.NewUserDomaincontroller(ctx, "trname", &fortios.UserDomaincontrollerArgs{
//				DomainName: pulumi.String("s.com"),
//				IpAddress:  pulumi.String("1.1.1.1"),
//				LdapServer: trname1.Name,
//				Port:       pulumi.Int(445),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # User DomainController can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/userDomaincontroller:UserDomaincontroller labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/userDomaincontroller:UserDomaincontroller labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type UserDomaincontroller struct {
	pulumi.CustomResourceState

	// Set Active Directory mode. Valid values: `none`, `ds`, `lds`.
	AdMode pulumi.StringOutput `pulumi:"adMode"`
	// AD LDS distinguished name.
	AdldsDn pulumi.StringOutput `pulumi:"adldsDn"`
	// AD LDS IPv6 address.
	AdldsIp6 pulumi.StringOutput `pulumi:"adldsIp6"`
	// AD LDS IPv4 address.
	AdldsIpAddress pulumi.StringOutput `pulumi:"adldsIpAddress"`
	// Port number of AD LDS service (default = 389).
	AdldsPort pulumi.IntOutput `pulumi:"adldsPort"`
	// Enable/disable DNS service lookup. Valid values: `enable`, `disable`.
	DnsSrvLookup pulumi.StringOutput `pulumi:"dnsSrvLookup"`
	// Domain DNS name.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// extra servers. The structure of `extraServer` block is documented below.
	ExtraServers UserDomaincontrollerExtraServerArrayOutput `pulumi:"extraServers"`
	// Hostname of the server to connect to.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// Domain controller IPv6 address.
	Ip6 pulumi.StringOutput `pulumi:"ip6"`
	// Domain controller IP address.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// LDAP server name.
	LdapServer pulumi.StringOutput `pulumi:"ldapServer"`
	// Domain controller entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password for specified username.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Port to be used for communication with the domain controller (default = 445).
	Port pulumi.IntOutput `pulumi:"port"`
	// Port to be used for communication with the domain controller for replication service. Port number 0 indicates automatic discovery.
	ReplicationPort pulumi.IntOutput `pulumi:"replicationPort"`
	// FortiGate IPv6 address to be used for communication with the domain controller.
	SourceIp6 pulumi.StringOutput `pulumi:"sourceIp6"`
	// FortiGate IPv4 address to be used for communication with the domain controller.
	SourceIpAddress pulumi.StringOutput `pulumi:"sourceIpAddress"`
	// Source port to be used for communication with the domain controller.
	SourcePort pulumi.IntOutput `pulumi:"sourcePort"`
	// User name to sign in with. Must have proper permissions for service.
	Username pulumi.StringOutput `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewUserDomaincontroller registers a new resource with the given unique name, arguments, and options.
func NewUserDomaincontroller(ctx *pulumi.Context,
	name string, args *UserDomaincontrollerArgs, opts ...pulumi.ResourceOption) (*UserDomaincontroller, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpAddress == nil {
		return nil, errors.New("invalid value for required argument 'IpAddress'")
	}
	if args.LdapServer == nil {
		return nil, errors.New("invalid value for required argument 'LdapServer'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource UserDomaincontroller
	err := ctx.RegisterResource("fortios:index/userDomaincontroller:UserDomaincontroller", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserDomaincontroller gets an existing UserDomaincontroller resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserDomaincontroller(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserDomaincontrollerState, opts ...pulumi.ResourceOption) (*UserDomaincontroller, error) {
	var resource UserDomaincontroller
	err := ctx.ReadResource("fortios:index/userDomaincontroller:UserDomaincontroller", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserDomaincontroller resources.
type userDomaincontrollerState struct {
	// Set Active Directory mode. Valid values: `none`, `ds`, `lds`.
	AdMode *string `pulumi:"adMode"`
	// AD LDS distinguished name.
	AdldsDn *string `pulumi:"adldsDn"`
	// AD LDS IPv6 address.
	AdldsIp6 *string `pulumi:"adldsIp6"`
	// AD LDS IPv4 address.
	AdldsIpAddress *string `pulumi:"adldsIpAddress"`
	// Port number of AD LDS service (default = 389).
	AdldsPort *int `pulumi:"adldsPort"`
	// Enable/disable DNS service lookup. Valid values: `enable`, `disable`.
	DnsSrvLookup *string `pulumi:"dnsSrvLookup"`
	// Domain DNS name.
	DomainName *string `pulumi:"domainName"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// extra servers. The structure of `extraServer` block is documented below.
	ExtraServers []UserDomaincontrollerExtraServer `pulumi:"extraServers"`
	// Hostname of the server to connect to.
	Hostname *string `pulumi:"hostname"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Domain controller IPv6 address.
	Ip6 *string `pulumi:"ip6"`
	// Domain controller IP address.
	IpAddress *string `pulumi:"ipAddress"`
	// LDAP server name.
	LdapServer *string `pulumi:"ldapServer"`
	// Domain controller entry name.
	Name *string `pulumi:"name"`
	// Password for specified username.
	Password *string `pulumi:"password"`
	// Port to be used for communication with the domain controller (default = 445).
	Port *int `pulumi:"port"`
	// Port to be used for communication with the domain controller for replication service. Port number 0 indicates automatic discovery.
	ReplicationPort *int `pulumi:"replicationPort"`
	// FortiGate IPv6 address to be used for communication with the domain controller.
	SourceIp6 *string `pulumi:"sourceIp6"`
	// FortiGate IPv4 address to be used for communication with the domain controller.
	SourceIpAddress *string `pulumi:"sourceIpAddress"`
	// Source port to be used for communication with the domain controller.
	SourcePort *int `pulumi:"sourcePort"`
	// User name to sign in with. Must have proper permissions for service.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type UserDomaincontrollerState struct {
	// Set Active Directory mode. Valid values: `none`, `ds`, `lds`.
	AdMode pulumi.StringPtrInput
	// AD LDS distinguished name.
	AdldsDn pulumi.StringPtrInput
	// AD LDS IPv6 address.
	AdldsIp6 pulumi.StringPtrInput
	// AD LDS IPv4 address.
	AdldsIpAddress pulumi.StringPtrInput
	// Port number of AD LDS service (default = 389).
	AdldsPort pulumi.IntPtrInput
	// Enable/disable DNS service lookup. Valid values: `enable`, `disable`.
	DnsSrvLookup pulumi.StringPtrInput
	// Domain DNS name.
	DomainName pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// extra servers. The structure of `extraServer` block is documented below.
	ExtraServers UserDomaincontrollerExtraServerArrayInput
	// Hostname of the server to connect to.
	Hostname pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Domain controller IPv6 address.
	Ip6 pulumi.StringPtrInput
	// Domain controller IP address.
	IpAddress pulumi.StringPtrInput
	// LDAP server name.
	LdapServer pulumi.StringPtrInput
	// Domain controller entry name.
	Name pulumi.StringPtrInput
	// Password for specified username.
	Password pulumi.StringPtrInput
	// Port to be used for communication with the domain controller (default = 445).
	Port pulumi.IntPtrInput
	// Port to be used for communication with the domain controller for replication service. Port number 0 indicates automatic discovery.
	ReplicationPort pulumi.IntPtrInput
	// FortiGate IPv6 address to be used for communication with the domain controller.
	SourceIp6 pulumi.StringPtrInput
	// FortiGate IPv4 address to be used for communication with the domain controller.
	SourceIpAddress pulumi.StringPtrInput
	// Source port to be used for communication with the domain controller.
	SourcePort pulumi.IntPtrInput
	// User name to sign in with. Must have proper permissions for service.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserDomaincontrollerState) ElementType() reflect.Type {
	return reflect.TypeOf((*userDomaincontrollerState)(nil)).Elem()
}

type userDomaincontrollerArgs struct {
	// Set Active Directory mode. Valid values: `none`, `ds`, `lds`.
	AdMode *string `pulumi:"adMode"`
	// AD LDS distinguished name.
	AdldsDn *string `pulumi:"adldsDn"`
	// AD LDS IPv6 address.
	AdldsIp6 *string `pulumi:"adldsIp6"`
	// AD LDS IPv4 address.
	AdldsIpAddress *string `pulumi:"adldsIpAddress"`
	// Port number of AD LDS service (default = 389).
	AdldsPort *int `pulumi:"adldsPort"`
	// Enable/disable DNS service lookup. Valid values: `enable`, `disable`.
	DnsSrvLookup *string `pulumi:"dnsSrvLookup"`
	// Domain DNS name.
	DomainName *string `pulumi:"domainName"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// extra servers. The structure of `extraServer` block is documented below.
	ExtraServers []UserDomaincontrollerExtraServer `pulumi:"extraServers"`
	// Hostname of the server to connect to.
	Hostname *string `pulumi:"hostname"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Domain controller IPv6 address.
	Ip6 *string `pulumi:"ip6"`
	// Domain controller IP address.
	IpAddress string `pulumi:"ipAddress"`
	// LDAP server name.
	LdapServer string `pulumi:"ldapServer"`
	// Domain controller entry name.
	Name *string `pulumi:"name"`
	// Password for specified username.
	Password *string `pulumi:"password"`
	// Port to be used for communication with the domain controller (default = 445).
	Port *int `pulumi:"port"`
	// Port to be used for communication with the domain controller for replication service. Port number 0 indicates automatic discovery.
	ReplicationPort *int `pulumi:"replicationPort"`
	// FortiGate IPv6 address to be used for communication with the domain controller.
	SourceIp6 *string `pulumi:"sourceIp6"`
	// FortiGate IPv4 address to be used for communication with the domain controller.
	SourceIpAddress *string `pulumi:"sourceIpAddress"`
	// Source port to be used for communication with the domain controller.
	SourcePort *int `pulumi:"sourcePort"`
	// User name to sign in with. Must have proper permissions for service.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserDomaincontroller resource.
type UserDomaincontrollerArgs struct {
	// Set Active Directory mode. Valid values: `none`, `ds`, `lds`.
	AdMode pulumi.StringPtrInput
	// AD LDS distinguished name.
	AdldsDn pulumi.StringPtrInput
	// AD LDS IPv6 address.
	AdldsIp6 pulumi.StringPtrInput
	// AD LDS IPv4 address.
	AdldsIpAddress pulumi.StringPtrInput
	// Port number of AD LDS service (default = 389).
	AdldsPort pulumi.IntPtrInput
	// Enable/disable DNS service lookup. Valid values: `enable`, `disable`.
	DnsSrvLookup pulumi.StringPtrInput
	// Domain DNS name.
	DomainName pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// extra servers. The structure of `extraServer` block is documented below.
	ExtraServers UserDomaincontrollerExtraServerArrayInput
	// Hostname of the server to connect to.
	Hostname pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Domain controller IPv6 address.
	Ip6 pulumi.StringPtrInput
	// Domain controller IP address.
	IpAddress pulumi.StringInput
	// LDAP server name.
	LdapServer pulumi.StringInput
	// Domain controller entry name.
	Name pulumi.StringPtrInput
	// Password for specified username.
	Password pulumi.StringPtrInput
	// Port to be used for communication with the domain controller (default = 445).
	Port pulumi.IntPtrInput
	// Port to be used for communication with the domain controller for replication service. Port number 0 indicates automatic discovery.
	ReplicationPort pulumi.IntPtrInput
	// FortiGate IPv6 address to be used for communication with the domain controller.
	SourceIp6 pulumi.StringPtrInput
	// FortiGate IPv4 address to be used for communication with the domain controller.
	SourceIpAddress pulumi.StringPtrInput
	// Source port to be used for communication with the domain controller.
	SourcePort pulumi.IntPtrInput
	// User name to sign in with. Must have proper permissions for service.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserDomaincontrollerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userDomaincontrollerArgs)(nil)).Elem()
}

type UserDomaincontrollerInput interface {
	pulumi.Input

	ToUserDomaincontrollerOutput() UserDomaincontrollerOutput
	ToUserDomaincontrollerOutputWithContext(ctx context.Context) UserDomaincontrollerOutput
}

func (*UserDomaincontroller) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDomaincontroller)(nil)).Elem()
}

func (i *UserDomaincontroller) ToUserDomaincontrollerOutput() UserDomaincontrollerOutput {
	return i.ToUserDomaincontrollerOutputWithContext(context.Background())
}

func (i *UserDomaincontroller) ToUserDomaincontrollerOutputWithContext(ctx context.Context) UserDomaincontrollerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDomaincontrollerOutput)
}

// UserDomaincontrollerArrayInput is an input type that accepts UserDomaincontrollerArray and UserDomaincontrollerArrayOutput values.
// You can construct a concrete instance of `UserDomaincontrollerArrayInput` via:
//
//	UserDomaincontrollerArray{ UserDomaincontrollerArgs{...} }
type UserDomaincontrollerArrayInput interface {
	pulumi.Input

	ToUserDomaincontrollerArrayOutput() UserDomaincontrollerArrayOutput
	ToUserDomaincontrollerArrayOutputWithContext(context.Context) UserDomaincontrollerArrayOutput
}

type UserDomaincontrollerArray []UserDomaincontrollerInput

func (UserDomaincontrollerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserDomaincontroller)(nil)).Elem()
}

func (i UserDomaincontrollerArray) ToUserDomaincontrollerArrayOutput() UserDomaincontrollerArrayOutput {
	return i.ToUserDomaincontrollerArrayOutputWithContext(context.Background())
}

func (i UserDomaincontrollerArray) ToUserDomaincontrollerArrayOutputWithContext(ctx context.Context) UserDomaincontrollerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDomaincontrollerArrayOutput)
}

// UserDomaincontrollerMapInput is an input type that accepts UserDomaincontrollerMap and UserDomaincontrollerMapOutput values.
// You can construct a concrete instance of `UserDomaincontrollerMapInput` via:
//
//	UserDomaincontrollerMap{ "key": UserDomaincontrollerArgs{...} }
type UserDomaincontrollerMapInput interface {
	pulumi.Input

	ToUserDomaincontrollerMapOutput() UserDomaincontrollerMapOutput
	ToUserDomaincontrollerMapOutputWithContext(context.Context) UserDomaincontrollerMapOutput
}

type UserDomaincontrollerMap map[string]UserDomaincontrollerInput

func (UserDomaincontrollerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserDomaincontroller)(nil)).Elem()
}

func (i UserDomaincontrollerMap) ToUserDomaincontrollerMapOutput() UserDomaincontrollerMapOutput {
	return i.ToUserDomaincontrollerMapOutputWithContext(context.Background())
}

func (i UserDomaincontrollerMap) ToUserDomaincontrollerMapOutputWithContext(ctx context.Context) UserDomaincontrollerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDomaincontrollerMapOutput)
}

type UserDomaincontrollerOutput struct{ *pulumi.OutputState }

func (UserDomaincontrollerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDomaincontroller)(nil)).Elem()
}

func (o UserDomaincontrollerOutput) ToUserDomaincontrollerOutput() UserDomaincontrollerOutput {
	return o
}

func (o UserDomaincontrollerOutput) ToUserDomaincontrollerOutputWithContext(ctx context.Context) UserDomaincontrollerOutput {
	return o
}

// Set Active Directory mode. Valid values: `none`, `ds`, `lds`.
func (o UserDomaincontrollerOutput) AdMode() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringOutput { return v.AdMode }).(pulumi.StringOutput)
}

// AD LDS distinguished name.
func (o UserDomaincontrollerOutput) AdldsDn() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringOutput { return v.AdldsDn }).(pulumi.StringOutput)
}

// AD LDS IPv6 address.
func (o UserDomaincontrollerOutput) AdldsIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringOutput { return v.AdldsIp6 }).(pulumi.StringOutput)
}

// AD LDS IPv4 address.
func (o UserDomaincontrollerOutput) AdldsIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringOutput { return v.AdldsIpAddress }).(pulumi.StringOutput)
}

// Port number of AD LDS service (default = 389).
func (o UserDomaincontrollerOutput) AdldsPort() pulumi.IntOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.IntOutput { return v.AdldsPort }).(pulumi.IntOutput)
}

// Enable/disable DNS service lookup. Valid values: `enable`, `disable`.
func (o UserDomaincontrollerOutput) DnsSrvLookup() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringOutput { return v.DnsSrvLookup }).(pulumi.StringOutput)
}

// Domain DNS name.
func (o UserDomaincontrollerOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o UserDomaincontrollerOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// extra servers. The structure of `extraServer` block is documented below.
func (o UserDomaincontrollerOutput) ExtraServers() UserDomaincontrollerExtraServerArrayOutput {
	return o.ApplyT(func(v *UserDomaincontroller) UserDomaincontrollerExtraServerArrayOutput { return v.ExtraServers }).(UserDomaincontrollerExtraServerArrayOutput)
}

// Hostname of the server to connect to.
func (o UserDomaincontrollerOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// Specify outgoing interface to reach server.
func (o UserDomaincontrollerOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
func (o UserDomaincontrollerOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

// Domain controller IPv6 address.
func (o UserDomaincontrollerOutput) Ip6() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringOutput { return v.Ip6 }).(pulumi.StringOutput)
}

// Domain controller IP address.
func (o UserDomaincontrollerOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// LDAP server name.
func (o UserDomaincontrollerOutput) LdapServer() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringOutput { return v.LdapServer }).(pulumi.StringOutput)
}

// Domain controller entry name.
func (o UserDomaincontrollerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password for specified username.
func (o UserDomaincontrollerOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Port to be used for communication with the domain controller (default = 445).
func (o UserDomaincontrollerOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Port to be used for communication with the domain controller for replication service. Port number 0 indicates automatic discovery.
func (o UserDomaincontrollerOutput) ReplicationPort() pulumi.IntOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.IntOutput { return v.ReplicationPort }).(pulumi.IntOutput)
}

// FortiGate IPv6 address to be used for communication with the domain controller.
func (o UserDomaincontrollerOutput) SourceIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringOutput { return v.SourceIp6 }).(pulumi.StringOutput)
}

// FortiGate IPv4 address to be used for communication with the domain controller.
func (o UserDomaincontrollerOutput) SourceIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringOutput { return v.SourceIpAddress }).(pulumi.StringOutput)
}

// Source port to be used for communication with the domain controller.
func (o UserDomaincontrollerOutput) SourcePort() pulumi.IntOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.IntOutput { return v.SourcePort }).(pulumi.IntOutput)
}

// User name to sign in with. Must have proper permissions for service.
func (o UserDomaincontrollerOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o UserDomaincontrollerOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDomaincontroller) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type UserDomaincontrollerArrayOutput struct{ *pulumi.OutputState }

func (UserDomaincontrollerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserDomaincontroller)(nil)).Elem()
}

func (o UserDomaincontrollerArrayOutput) ToUserDomaincontrollerArrayOutput() UserDomaincontrollerArrayOutput {
	return o
}

func (o UserDomaincontrollerArrayOutput) ToUserDomaincontrollerArrayOutputWithContext(ctx context.Context) UserDomaincontrollerArrayOutput {
	return o
}

func (o UserDomaincontrollerArrayOutput) Index(i pulumi.IntInput) UserDomaincontrollerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserDomaincontroller {
		return vs[0].([]*UserDomaincontroller)[vs[1].(int)]
	}).(UserDomaincontrollerOutput)
}

type UserDomaincontrollerMapOutput struct{ *pulumi.OutputState }

func (UserDomaincontrollerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserDomaincontroller)(nil)).Elem()
}

func (o UserDomaincontrollerMapOutput) ToUserDomaincontrollerMapOutput() UserDomaincontrollerMapOutput {
	return o
}

func (o UserDomaincontrollerMapOutput) ToUserDomaincontrollerMapOutputWithContext(ctx context.Context) UserDomaincontrollerMapOutput {
	return o
}

func (o UserDomaincontrollerMapOutput) MapIndex(k pulumi.StringInput) UserDomaincontrollerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserDomaincontroller {
		return vs[0].(map[string]*UserDomaincontroller)[vs[1].(string)]
	}).(UserDomaincontrollerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserDomaincontrollerInput)(nil)).Elem(), &UserDomaincontroller{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserDomaincontrollerArrayInput)(nil)).Elem(), UserDomaincontrollerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserDomaincontrollerMapInput)(nil)).Elem(), UserDomaincontrollerMap{})
	pulumi.RegisterOutputType(UserDomaincontrollerOutput{})
	pulumi.RegisterOutputType(UserDomaincontrollerArrayOutput{})
	pulumi.RegisterOutputType(UserDomaincontrollerMapOutput{})
}
