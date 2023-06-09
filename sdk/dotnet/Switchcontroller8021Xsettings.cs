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
    /// Configure global 802.1X settings.
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
    ///     var trname = new Fortios.Switchcontroller8021Xsettings("trname", new()
    ///     {
    ///         LinkDownAuth = "set-unauth",
    ///         MaxReauthAttempt = 3,
    ///         ReauthPeriod = 12,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SwitchController 8021XSettings can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontroller8021Xsettings:Switchcontroller8021Xsettings labelname SwitchController8021XSettings
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontroller8021Xsettings:Switchcontroller8021Xsettings labelname SwitchController8021XSettings
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/switchcontroller8021Xsettings:Switchcontroller8021Xsettings")]
    public partial class Switchcontroller8021Xsettings : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Interface-reauthentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.
        /// </summary>
        [Output("linkDownAuth")]
        public Output<string> LinkDownAuth { get; private set; } = null!;

        /// <summary>
        /// Enable/disable MAB re-authentication. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("mabReauth")]
        public Output<string> MabReauth { get; private set; } = null!;

        /// <summary>
        /// Maximum number of authentication attempts (0 - 15, default = 3).
        /// </summary>
        [Output("maxReauthAttempt")]
        public Output<int> MaxReauthAttempt { get; private set; } = null!;

        /// <summary>
        /// Period of time to allow for reauthentication (1 - 1440 sec, default = 60, 0 = disable reauthentication).
        /// </summary>
        [Output("reauthPeriod")]
        public Output<int> ReauthPeriod { get; private set; } = null!;

        /// <summary>
        /// 802.1X Tx period (seconds, default=30).
        /// </summary>
        [Output("txPeriod")]
        public Output<int> TxPeriod { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Switchcontroller8021Xsettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Switchcontroller8021Xsettings(string name, Switchcontroller8021XsettingsArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontroller8021Xsettings:Switchcontroller8021Xsettings", name, args ?? new Switchcontroller8021XsettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Switchcontroller8021Xsettings(string name, Input<string> id, Switchcontroller8021XsettingsState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontroller8021Xsettings:Switchcontroller8021Xsettings", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Switchcontroller8021Xsettings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Switchcontroller8021Xsettings Get(string name, Input<string> id, Switchcontroller8021XsettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new Switchcontroller8021Xsettings(name, id, state, options);
        }
    }

    public sealed class Switchcontroller8021XsettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interface-reauthentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.
        /// </summary>
        [Input("linkDownAuth")]
        public Input<string>? LinkDownAuth { get; set; }

        /// <summary>
        /// Enable/disable MAB re-authentication. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("mabReauth")]
        public Input<string>? MabReauth { get; set; }

        /// <summary>
        /// Maximum number of authentication attempts (0 - 15, default = 3).
        /// </summary>
        [Input("maxReauthAttempt")]
        public Input<int>? MaxReauthAttempt { get; set; }

        /// <summary>
        /// Period of time to allow for reauthentication (1 - 1440 sec, default = 60, 0 = disable reauthentication).
        /// </summary>
        [Input("reauthPeriod")]
        public Input<int>? ReauthPeriod { get; set; }

        /// <summary>
        /// 802.1X Tx period (seconds, default=30).
        /// </summary>
        [Input("txPeriod")]
        public Input<int>? TxPeriod { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Switchcontroller8021XsettingsArgs()
        {
        }
        public static new Switchcontroller8021XsettingsArgs Empty => new Switchcontroller8021XsettingsArgs();
    }

    public sealed class Switchcontroller8021XsettingsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interface-reauthentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.
        /// </summary>
        [Input("linkDownAuth")]
        public Input<string>? LinkDownAuth { get; set; }

        /// <summary>
        /// Enable/disable MAB re-authentication. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("mabReauth")]
        public Input<string>? MabReauth { get; set; }

        /// <summary>
        /// Maximum number of authentication attempts (0 - 15, default = 3).
        /// </summary>
        [Input("maxReauthAttempt")]
        public Input<int>? MaxReauthAttempt { get; set; }

        /// <summary>
        /// Period of time to allow for reauthentication (1 - 1440 sec, default = 60, 0 = disable reauthentication).
        /// </summary>
        [Input("reauthPeriod")]
        public Input<int>? ReauthPeriod { get; set; }

        /// <summary>
        /// 802.1X Tx period (seconds, default=30).
        /// </summary>
        [Input("txPeriod")]
        public Input<int>? TxPeriod { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Switchcontroller8021XsettingsState()
        {
        }
        public static new Switchcontroller8021XsettingsState Empty => new Switchcontroller8021XsettingsState();
    }
}
