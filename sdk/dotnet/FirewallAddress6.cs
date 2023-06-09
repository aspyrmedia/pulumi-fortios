// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    /// <summary>
    /// Configure IPv6 firewall addresses.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.FirewallAddress6("trname", new()
    ///     {
    ///         CacheTtl = 0,
    ///         Color = 0,
    ///         EndIp = "::",
    ///         Host = "fdff:ffff::",
    ///         HostType = "any",
    ///         Ip6 = "fdff:ffff::/120",
    ///         StartIp = "fdff:ffff::",
    ///         Type = "ipprefix",
    ///         Visibility = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall Address6 can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/firewallAddress6:FirewallAddress6 labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/firewallAddress6:FirewallAddress6 labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/firewallAddress6:FirewallAddress6")]
    public partial class FirewallAddress6 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Minimal TTL of individual IPv6 addresses in FQDN cache.
        /// </summary>
        [Output("cacheTtl")]
        public Output<int> CacheTtl { get; private set; } = null!;

        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        /// </summary>
        [Output("color")]
        public Output<int> Color { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// IPv6 addresses associated to a specific country.
        /// </summary>
        [Output("country")]
        public Output<string> Country { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        /// </summary>
        [Output("endIp")]
        public Output<string> EndIp { get; private set; } = null!;

        /// <summary>
        /// Last MAC address in the range.
        /// </summary>
        [Output("endMac")]
        public Output<string> EndMac { get; private set; } = null!;

        /// <summary>
        /// Endpoint group name.
        /// </summary>
        [Output("epgName")]
        public Output<string> EpgName { get; private set; } = null!;

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fabricObject")]
        public Output<string> FabricObject { get; private set; } = null!;

        /// <summary>
        /// Fully qualified domain name.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// Host Address.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// Host type. Valid values: `any`, `specific`.
        /// </summary>
        [Output("hostType")]
        public Output<string> HostType { get; private set; } = null!;

        /// <summary>
        /// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
        /// </summary>
        [Output("ip6")]
        public Output<string> Ip6 { get; private set; } = null!;

        /// <summary>
        /// IP address list. The structure of `list` block is documented below.
        /// </summary>
        [Output("lists")]
        public Output<ImmutableArray<Outputs.FirewallAddress6List>> Lists { get; private set; } = null!;

        /// <summary>
        /// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
        /// </summary>
        [Output("macaddrs")]
        public Output<ImmutableArray<Outputs.FirewallAddress6Macaddr>> Macaddrs { get; private set; } = null!;

        /// <summary>
        /// Address name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Object ID for NSX.
        /// </summary>
        [Output("objId")]
        public Output<string?> ObjId { get; private set; } = null!;

        /// <summary>
        /// SDN.
        /// </summary>
        [Output("sdn")]
        public Output<string> Sdn { get; private set; } = null!;

        /// <summary>
        /// SDN Tag.
        /// </summary>
        [Output("sdnTag")]
        public Output<string> SdnTag { get; private set; } = null!;

        /// <summary>
        /// First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        /// </summary>
        [Output("startIp")]
        public Output<string> StartIp { get; private set; } = null!;

        /// <summary>
        /// First MAC address in the range.
        /// </summary>
        [Output("startMac")]
        public Output<string> StartMac { get; private set; } = null!;

        /// <summary>
        /// IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
        /// </summary>
        [Output("subnetSegments")]
        public Output<ImmutableArray<Outputs.FirewallAddress6SubnetSegment>> SubnetSegments { get; private set; } = null!;

        /// <summary>
        /// Config object tagging The structure of `tagging` block is documented below.
        /// </summary>
        [Output("taggings")]
        public Output<ImmutableArray<Outputs.FirewallAddress6Tagging>> Taggings { get; private set; } = null!;

        /// <summary>
        /// IPv6 address template.
        /// </summary>
        [Output("template")]
        public Output<string> Template { get; private set; } = null!;

        /// <summary>
        /// Tenant.
        /// </summary>
        [Output("tenant")]
        public Output<string> Tenant { get; private set; } = null!;

        /// <summary>
        /// Type of IPv6 address object (default = ipprefix).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Enable/disable the visibility of the object in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallAddress6 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallAddress6(string name, FirewallAddress6Args? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallAddress6:FirewallAddress6", name, args ?? new FirewallAddress6Args(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallAddress6(string name, Input<string> id, FirewallAddress6State? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallAddress6:FirewallAddress6", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FirewallAddress6 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallAddress6 Get(string name, Input<string> id, FirewallAddress6State? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallAddress6(name, id, state, options);
        }
    }

    public sealed class FirewallAddress6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Minimal TTL of individual IPv6 addresses in FQDN cache.
        /// </summary>
        [Input("cacheTtl")]
        public Input<int>? CacheTtl { get; set; }

        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// IPv6 addresses associated to a specific country.
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        /// </summary>
        [Input("endIp")]
        public Input<string>? EndIp { get; set; }

        /// <summary>
        /// Last MAC address in the range.
        /// </summary>
        [Input("endMac")]
        public Input<string>? EndMac { get; set; }

        /// <summary>
        /// Endpoint group name.
        /// </summary>
        [Input("epgName")]
        public Input<string>? EpgName { get; set; }

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fabricObject")]
        public Input<string>? FabricObject { get; set; }

        /// <summary>
        /// Fully qualified domain name.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// Host Address.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Host type. Valid values: `any`, `specific`.
        /// </summary>
        [Input("hostType")]
        public Input<string>? HostType { get; set; }

        /// <summary>
        /// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
        /// </summary>
        [Input("ip6")]
        public Input<string>? Ip6 { get; set; }

        [Input("lists")]
        private InputList<Inputs.FirewallAddress6ListArgs>? _lists;

        /// <summary>
        /// IP address list. The structure of `list` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallAddress6ListArgs> Lists
        {
            get => _lists ?? (_lists = new InputList<Inputs.FirewallAddress6ListArgs>());
            set => _lists = value;
        }

        [Input("macaddrs")]
        private InputList<Inputs.FirewallAddress6MacaddrArgs>? _macaddrs;

        /// <summary>
        /// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallAddress6MacaddrArgs> Macaddrs
        {
            get => _macaddrs ?? (_macaddrs = new InputList<Inputs.FirewallAddress6MacaddrArgs>());
            set => _macaddrs = value;
        }

        /// <summary>
        /// Address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Object ID for NSX.
        /// </summary>
        [Input("objId")]
        public Input<string>? ObjId { get; set; }

        /// <summary>
        /// SDN.
        /// </summary>
        [Input("sdn")]
        public Input<string>? Sdn { get; set; }

        /// <summary>
        /// SDN Tag.
        /// </summary>
        [Input("sdnTag")]
        public Input<string>? SdnTag { get; set; }

        /// <summary>
        /// First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        /// </summary>
        [Input("startIp")]
        public Input<string>? StartIp { get; set; }

        /// <summary>
        /// First MAC address in the range.
        /// </summary>
        [Input("startMac")]
        public Input<string>? StartMac { get; set; }

        [Input("subnetSegments")]
        private InputList<Inputs.FirewallAddress6SubnetSegmentArgs>? _subnetSegments;

        /// <summary>
        /// IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallAddress6SubnetSegmentArgs> SubnetSegments
        {
            get => _subnetSegments ?? (_subnetSegments = new InputList<Inputs.FirewallAddress6SubnetSegmentArgs>());
            set => _subnetSegments = value;
        }

        [Input("taggings")]
        private InputList<Inputs.FirewallAddress6TaggingArgs>? _taggings;

        /// <summary>
        /// Config object tagging The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallAddress6TaggingArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.FirewallAddress6TaggingArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// IPv6 address template.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        /// <summary>
        /// Tenant.
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        /// <summary>
        /// Type of IPv6 address object (default = ipprefix).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable the visibility of the object in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public FirewallAddress6Args()
        {
        }
        public static new FirewallAddress6Args Empty => new FirewallAddress6Args();
    }

    public sealed class FirewallAddress6State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Minimal TTL of individual IPv6 addresses in FQDN cache.
        /// </summary>
        [Input("cacheTtl")]
        public Input<int>? CacheTtl { get; set; }

        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// IPv6 addresses associated to a specific country.
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        /// </summary>
        [Input("endIp")]
        public Input<string>? EndIp { get; set; }

        /// <summary>
        /// Last MAC address in the range.
        /// </summary>
        [Input("endMac")]
        public Input<string>? EndMac { get; set; }

        /// <summary>
        /// Endpoint group name.
        /// </summary>
        [Input("epgName")]
        public Input<string>? EpgName { get; set; }

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fabricObject")]
        public Input<string>? FabricObject { get; set; }

        /// <summary>
        /// Fully qualified domain name.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// Host Address.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Host type. Valid values: `any`, `specific`.
        /// </summary>
        [Input("hostType")]
        public Input<string>? HostType { get; set; }

        /// <summary>
        /// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
        /// </summary>
        [Input("ip6")]
        public Input<string>? Ip6 { get; set; }

        [Input("lists")]
        private InputList<Inputs.FirewallAddress6ListGetArgs>? _lists;

        /// <summary>
        /// IP address list. The structure of `list` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallAddress6ListGetArgs> Lists
        {
            get => _lists ?? (_lists = new InputList<Inputs.FirewallAddress6ListGetArgs>());
            set => _lists = value;
        }

        [Input("macaddrs")]
        private InputList<Inputs.FirewallAddress6MacaddrGetArgs>? _macaddrs;

        /// <summary>
        /// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallAddress6MacaddrGetArgs> Macaddrs
        {
            get => _macaddrs ?? (_macaddrs = new InputList<Inputs.FirewallAddress6MacaddrGetArgs>());
            set => _macaddrs = value;
        }

        /// <summary>
        /// Address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Object ID for NSX.
        /// </summary>
        [Input("objId")]
        public Input<string>? ObjId { get; set; }

        /// <summary>
        /// SDN.
        /// </summary>
        [Input("sdn")]
        public Input<string>? Sdn { get; set; }

        /// <summary>
        /// SDN Tag.
        /// </summary>
        [Input("sdnTag")]
        public Input<string>? SdnTag { get; set; }

        /// <summary>
        /// First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        /// </summary>
        [Input("startIp")]
        public Input<string>? StartIp { get; set; }

        /// <summary>
        /// First MAC address in the range.
        /// </summary>
        [Input("startMac")]
        public Input<string>? StartMac { get; set; }

        [Input("subnetSegments")]
        private InputList<Inputs.FirewallAddress6SubnetSegmentGetArgs>? _subnetSegments;

        /// <summary>
        /// IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallAddress6SubnetSegmentGetArgs> SubnetSegments
        {
            get => _subnetSegments ?? (_subnetSegments = new InputList<Inputs.FirewallAddress6SubnetSegmentGetArgs>());
            set => _subnetSegments = value;
        }

        [Input("taggings")]
        private InputList<Inputs.FirewallAddress6TaggingGetArgs>? _taggings;

        /// <summary>
        /// Config object tagging The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallAddress6TaggingGetArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.FirewallAddress6TaggingGetArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// IPv6 address template.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        /// <summary>
        /// Tenant.
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        /// <summary>
        /// Type of IPv6 address object (default = ipprefix).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable the visibility of the object in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public FirewallAddress6State()
        {
        }
        public static new FirewallAddress6State Empty => new FirewallAddress6State();
    }
}
