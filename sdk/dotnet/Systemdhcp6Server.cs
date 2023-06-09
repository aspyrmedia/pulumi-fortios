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
    /// Configure DHCPv6 servers.
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
    ///     var trname = new Fortios.Systemdhcp6Server("trname", new()
    ///     {
    ///         Fosid = 1,
    ///         Interface = "port3",
    ///         LeaseTime = 604800,
    ///         RapidCommit = "disable",
    ///         Status = "enable",
    ///         Subnet = "2001:db8:1234:113::/64",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SystemDhcp6 Server can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemdhcp6Server:Systemdhcp6Server labelname {{fosid}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemdhcp6Server:Systemdhcp6Server labelname {{fosid}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemdhcp6Server:Systemdhcp6Server")]
    public partial class Systemdhcp6Server : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IAID of obtained delegated-prefix from the upstream interface.
        /// </summary>
        [Output("delegatedPrefixIaid")]
        public Output<int> DelegatedPrefixIaid { get; private set; } = null!;

        /// <summary>
        /// DNS search list options. Valid values: `delegated`, `specify`.
        /// </summary>
        [Output("dnsSearchList")]
        public Output<string> DnsSearchList { get; private set; } = null!;

        /// <summary>
        /// DNS server 1.
        /// </summary>
        [Output("dnsServer1")]
        public Output<string> DnsServer1 { get; private set; } = null!;

        /// <summary>
        /// DNS server 2.
        /// </summary>
        [Output("dnsServer2")]
        public Output<string> DnsServer2 { get; private set; } = null!;

        /// <summary>
        /// DNS server 3.
        /// </summary>
        [Output("dnsServer3")]
        public Output<string> DnsServer3 { get; private set; } = null!;

        /// <summary>
        /// DNS server 4.
        /// </summary>
        [Output("dnsServer4")]
        public Output<string> DnsServer4 { get; private set; } = null!;

        /// <summary>
        /// Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
        /// </summary>
        [Output("dnsService")]
        public Output<string> DnsService { get; private set; } = null!;

        /// <summary>
        /// Domain name suffix for the IP addresses that the DHCP server assigns to clients.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// DHCP server can assign IP configurations to clients connected to this interface.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Method used to assign client IP. Valid values: `range`, `delegated`.
        /// </summary>
        [Output("ipMode")]
        public Output<string> IpMode { get; private set; } = null!;

        /// <summary>
        /// DHCP IP range configuration. The structure of `ip_range` block is documented below.
        /// </summary>
        [Output("ipRanges")]
        public Output<ImmutableArray<Outputs.Systemdhcp6ServerIpRange>> IpRanges { get; private set; } = null!;

        /// <summary>
        /// Lease time in seconds, 0 means unlimited.
        /// </summary>
        [Output("leaseTime")]
        public Output<int> LeaseTime { get; private set; } = null!;

        /// <summary>
        /// Option 1.
        /// </summary>
        [Output("option1")]
        public Output<string> Option1 { get; private set; } = null!;

        /// <summary>
        /// Option 2.
        /// </summary>
        [Output("option2")]
        public Output<string> Option2 { get; private set; } = null!;

        /// <summary>
        /// Option 3.
        /// </summary>
        [Output("option3")]
        public Output<string> Option3 { get; private set; } = null!;

        /// <summary>
        /// Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
        /// </summary>
        [Output("prefixMode")]
        public Output<string> PrefixMode { get; private set; } = null!;

        /// <summary>
        /// DHCP prefix configuration. The structure of `prefix_range` block is documented below.
        /// </summary>
        [Output("prefixRanges")]
        public Output<ImmutableArray<Outputs.Systemdhcp6ServerPrefixRange>> PrefixRanges { get; private set; } = null!;

        /// <summary>
        /// Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("rapidCommit")]
        public Output<string> RapidCommit { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Subnet or subnet-id if the IP mode is delegated.
        /// </summary>
        [Output("subnet")]
        public Output<string> Subnet { get; private set; } = null!;

        /// <summary>
        /// Interface name from where delegated information is provided.
        /// </summary>
        [Output("upstreamInterface")]
        public Output<string> UpstreamInterface { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Systemdhcp6Server resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Systemdhcp6Server(string name, Systemdhcp6ServerArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/systemdhcp6Server:Systemdhcp6Server", name, args ?? new Systemdhcp6ServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Systemdhcp6Server(string name, Input<string> id, Systemdhcp6ServerState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemdhcp6Server:Systemdhcp6Server", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Systemdhcp6Server resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Systemdhcp6Server Get(string name, Input<string> id, Systemdhcp6ServerState? state = null, CustomResourceOptions? options = null)
        {
            return new Systemdhcp6Server(name, id, state, options);
        }
    }

    public sealed class Systemdhcp6ServerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IAID of obtained delegated-prefix from the upstream interface.
        /// </summary>
        [Input("delegatedPrefixIaid")]
        public Input<int>? DelegatedPrefixIaid { get; set; }

        /// <summary>
        /// DNS search list options. Valid values: `delegated`, `specify`.
        /// </summary>
        [Input("dnsSearchList")]
        public Input<string>? DnsSearchList { get; set; }

        /// <summary>
        /// DNS server 1.
        /// </summary>
        [Input("dnsServer1")]
        public Input<string>? DnsServer1 { get; set; }

        /// <summary>
        /// DNS server 2.
        /// </summary>
        [Input("dnsServer2")]
        public Input<string>? DnsServer2 { get; set; }

        /// <summary>
        /// DNS server 3.
        /// </summary>
        [Input("dnsServer3")]
        public Input<string>? DnsServer3 { get; set; }

        /// <summary>
        /// DNS server 4.
        /// </summary>
        [Input("dnsServer4")]
        public Input<string>? DnsServer4 { get; set; }

        /// <summary>
        /// Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
        /// </summary>
        [Input("dnsService")]
        public Input<string>? DnsService { get; set; }

        /// <summary>
        /// Domain name suffix for the IP addresses that the DHCP server assigns to clients.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("fosid", required: true)]
        public Input<int> Fosid { get; set; } = null!;

        /// <summary>
        /// DHCP server can assign IP configurations to clients connected to this interface.
        /// </summary>
        [Input("interface", required: true)]
        public Input<string> Interface { get; set; } = null!;

        /// <summary>
        /// Method used to assign client IP. Valid values: `range`, `delegated`.
        /// </summary>
        [Input("ipMode")]
        public Input<string>? IpMode { get; set; }

        [Input("ipRanges")]
        private InputList<Inputs.Systemdhcp6ServerIpRangeArgs>? _ipRanges;

        /// <summary>
        /// DHCP IP range configuration. The structure of `ip_range` block is documented below.
        /// </summary>
        public InputList<Inputs.Systemdhcp6ServerIpRangeArgs> IpRanges
        {
            get => _ipRanges ?? (_ipRanges = new InputList<Inputs.Systemdhcp6ServerIpRangeArgs>());
            set => _ipRanges = value;
        }

        /// <summary>
        /// Lease time in seconds, 0 means unlimited.
        /// </summary>
        [Input("leaseTime")]
        public Input<int>? LeaseTime { get; set; }

        /// <summary>
        /// Option 1.
        /// </summary>
        [Input("option1")]
        public Input<string>? Option1 { get; set; }

        /// <summary>
        /// Option 2.
        /// </summary>
        [Input("option2")]
        public Input<string>? Option2 { get; set; }

        /// <summary>
        /// Option 3.
        /// </summary>
        [Input("option3")]
        public Input<string>? Option3 { get; set; }

        /// <summary>
        /// Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
        /// </summary>
        [Input("prefixMode")]
        public Input<string>? PrefixMode { get; set; }

        [Input("prefixRanges")]
        private InputList<Inputs.Systemdhcp6ServerPrefixRangeArgs>? _prefixRanges;

        /// <summary>
        /// DHCP prefix configuration. The structure of `prefix_range` block is documented below.
        /// </summary>
        public InputList<Inputs.Systemdhcp6ServerPrefixRangeArgs> PrefixRanges
        {
            get => _prefixRanges ?? (_prefixRanges = new InputList<Inputs.Systemdhcp6ServerPrefixRangeArgs>());
            set => _prefixRanges = value;
        }

        /// <summary>
        /// Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("rapidCommit")]
        public Input<string>? RapidCommit { get; set; }

        /// <summary>
        /// Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Subnet or subnet-id if the IP mode is delegated.
        /// </summary>
        [Input("subnet", required: true)]
        public Input<string> Subnet { get; set; } = null!;

        /// <summary>
        /// Interface name from where delegated information is provided.
        /// </summary>
        [Input("upstreamInterface")]
        public Input<string>? UpstreamInterface { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Systemdhcp6ServerArgs()
        {
        }
        public static new Systemdhcp6ServerArgs Empty => new Systemdhcp6ServerArgs();
    }

    public sealed class Systemdhcp6ServerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IAID of obtained delegated-prefix from the upstream interface.
        /// </summary>
        [Input("delegatedPrefixIaid")]
        public Input<int>? DelegatedPrefixIaid { get; set; }

        /// <summary>
        /// DNS search list options. Valid values: `delegated`, `specify`.
        /// </summary>
        [Input("dnsSearchList")]
        public Input<string>? DnsSearchList { get; set; }

        /// <summary>
        /// DNS server 1.
        /// </summary>
        [Input("dnsServer1")]
        public Input<string>? DnsServer1 { get; set; }

        /// <summary>
        /// DNS server 2.
        /// </summary>
        [Input("dnsServer2")]
        public Input<string>? DnsServer2 { get; set; }

        /// <summary>
        /// DNS server 3.
        /// </summary>
        [Input("dnsServer3")]
        public Input<string>? DnsServer3 { get; set; }

        /// <summary>
        /// DNS server 4.
        /// </summary>
        [Input("dnsServer4")]
        public Input<string>? DnsServer4 { get; set; }

        /// <summary>
        /// Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
        /// </summary>
        [Input("dnsService")]
        public Input<string>? DnsService { get; set; }

        /// <summary>
        /// Domain name suffix for the IP addresses that the DHCP server assigns to clients.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// DHCP server can assign IP configurations to clients connected to this interface.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Method used to assign client IP. Valid values: `range`, `delegated`.
        /// </summary>
        [Input("ipMode")]
        public Input<string>? IpMode { get; set; }

        [Input("ipRanges")]
        private InputList<Inputs.Systemdhcp6ServerIpRangeGetArgs>? _ipRanges;

        /// <summary>
        /// DHCP IP range configuration. The structure of `ip_range` block is documented below.
        /// </summary>
        public InputList<Inputs.Systemdhcp6ServerIpRangeGetArgs> IpRanges
        {
            get => _ipRanges ?? (_ipRanges = new InputList<Inputs.Systemdhcp6ServerIpRangeGetArgs>());
            set => _ipRanges = value;
        }

        /// <summary>
        /// Lease time in seconds, 0 means unlimited.
        /// </summary>
        [Input("leaseTime")]
        public Input<int>? LeaseTime { get; set; }

        /// <summary>
        /// Option 1.
        /// </summary>
        [Input("option1")]
        public Input<string>? Option1 { get; set; }

        /// <summary>
        /// Option 2.
        /// </summary>
        [Input("option2")]
        public Input<string>? Option2 { get; set; }

        /// <summary>
        /// Option 3.
        /// </summary>
        [Input("option3")]
        public Input<string>? Option3 { get; set; }

        /// <summary>
        /// Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
        /// </summary>
        [Input("prefixMode")]
        public Input<string>? PrefixMode { get; set; }

        [Input("prefixRanges")]
        private InputList<Inputs.Systemdhcp6ServerPrefixRangeGetArgs>? _prefixRanges;

        /// <summary>
        /// DHCP prefix configuration. The structure of `prefix_range` block is documented below.
        /// </summary>
        public InputList<Inputs.Systemdhcp6ServerPrefixRangeGetArgs> PrefixRanges
        {
            get => _prefixRanges ?? (_prefixRanges = new InputList<Inputs.Systemdhcp6ServerPrefixRangeGetArgs>());
            set => _prefixRanges = value;
        }

        /// <summary>
        /// Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("rapidCommit")]
        public Input<string>? RapidCommit { get; set; }

        /// <summary>
        /// Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Subnet or subnet-id if the IP mode is delegated.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// Interface name from where delegated information is provided.
        /// </summary>
        [Input("upstreamInterface")]
        public Input<string>? UpstreamInterface { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Systemdhcp6ServerState()
        {
        }
        public static new Systemdhcp6ServerState Empty => new Systemdhcp6ServerState();
    }
}
