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
    /// Configure FM. Applies to FortiOS Version `&lt;= 7.0.1`.
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
    ///     var trname = new Fortios.SystemFm("trname", new()
    ///     {
    ///         AutoBackup = "disable",
    ///         Ip = "0.0.0.0",
    ///         Ipsec = "disable",
    ///         ScheduledConfigRestore = "disable",
    ///         Status = "disable",
    ///         Vdom = "root",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System Fm can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemFm:SystemFm labelname SystemFm
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemFm:SystemFm labelname SystemFm
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemFm:SystemFm")]
    public partial class SystemFm : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable automatic backup. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("autoBackup")]
        public Output<string> AutoBackup { get; private set; } = null!;

        /// <summary>
        /// ID.
        /// </summary>
        [Output("fosid")]
        public Output<string> Fosid { get; private set; } = null!;

        /// <summary>
        /// IP address.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// Enable/disable IPsec. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ipsec")]
        public Output<string> Ipsec { get; private set; } = null!;

        /// <summary>
        /// Enable/disable scheduled configuration restore. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("scheduledConfigRestore")]
        public Output<string> ScheduledConfigRestore { get; private set; } = null!;

        /// <summary>
        /// Enable/disable FM. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// VDOM.
        /// </summary>
        [Output("vdom")]
        public Output<string> Vdom { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SystemFm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemFm(string name, SystemFmArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemFm:SystemFm", name, args ?? new SystemFmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemFm(string name, Input<string> id, SystemFmState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemFm:SystemFm", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemFm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemFm Get(string name, Input<string> id, SystemFmState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemFm(name, id, state, options);
        }
    }

    public sealed class SystemFmArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable automatic backup. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autoBackup")]
        public Input<string>? AutoBackup { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("fosid")]
        public Input<string>? Fosid { get; set; }

        /// <summary>
        /// IP address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Enable/disable IPsec. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsec")]
        public Input<string>? Ipsec { get; set; }

        /// <summary>
        /// Enable/disable scheduled configuration restore. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("scheduledConfigRestore")]
        public Input<string>? ScheduledConfigRestore { get; set; }

        /// <summary>
        /// Enable/disable FM. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// VDOM.
        /// </summary>
        [Input("vdom")]
        public Input<string>? Vdom { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemFmArgs()
        {
        }
        public static new SystemFmArgs Empty => new SystemFmArgs();
    }

    public sealed class SystemFmState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable automatic backup. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autoBackup")]
        public Input<string>? AutoBackup { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("fosid")]
        public Input<string>? Fosid { get; set; }

        /// <summary>
        /// IP address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Enable/disable IPsec. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsec")]
        public Input<string>? Ipsec { get; set; }

        /// <summary>
        /// Enable/disable scheduled configuration restore. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("scheduledConfigRestore")]
        public Input<string>? ScheduledConfigRestore { get; set; }

        /// <summary>
        /// Enable/disable FM. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// VDOM.
        /// </summary>
        [Input("vdom")]
        public Input<string>? Vdom { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemFmState()
        {
        }
        public static new SystemFmState Empty => new SystemFmState();
    }
}
