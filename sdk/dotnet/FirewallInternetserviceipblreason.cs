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
    /// IP blacklist reason. Applies to FortiOS Version `&gt;= 6.2.4`.
    /// 
    /// ## Import
    /// 
    /// Firewall InternetServiceIpblReason can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/firewallInternetserviceipblreason:FirewallInternetserviceipblreason labelname {{fosid}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/firewallInternetserviceipblreason:FirewallInternetserviceipblreason labelname {{fosid}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/firewallInternetserviceipblreason:FirewallInternetserviceipblreason")]
    public partial class FirewallInternetserviceipblreason : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IP blacklist reason ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// IP blacklist reason name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallInternetserviceipblreason resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallInternetserviceipblreason(string name, FirewallInternetserviceipblreasonArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallInternetserviceipblreason:FirewallInternetserviceipblreason", name, args ?? new FirewallInternetserviceipblreasonArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallInternetserviceipblreason(string name, Input<string> id, FirewallInternetserviceipblreasonState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallInternetserviceipblreason:FirewallInternetserviceipblreason", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallInternetserviceipblreason resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallInternetserviceipblreason Get(string name, Input<string> id, FirewallInternetserviceipblreasonState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallInternetserviceipblreason(name, id, state, options);
        }
    }

    public sealed class FirewallInternetserviceipblreasonArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP blacklist reason ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// IP blacklist reason name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FirewallInternetserviceipblreasonArgs()
        {
        }
        public static new FirewallInternetserviceipblreasonArgs Empty => new FirewallInternetserviceipblreasonArgs();
    }

    public sealed class FirewallInternetserviceipblreasonState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP blacklist reason ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// IP blacklist reason name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FirewallInternetserviceipblreasonState()
        {
        }
        public static new FirewallInternetserviceipblreasonState Empty => new FirewallInternetserviceipblreasonState();
    }
}
