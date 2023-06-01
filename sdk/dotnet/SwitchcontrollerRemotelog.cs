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
    /// Configure logging by FortiSwitch device to a remote syslog server. Applies to FortiOS Version `&gt;= 6.2.4`.
    /// 
    /// ## Import
    /// 
    /// SwitchController RemoteLog can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerRemotelog:SwitchcontrollerRemotelog labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerRemotelog:SwitchcontrollerRemotelog labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/switchcontrollerRemotelog:SwitchcontrollerRemotelog")]
    public partial class SwitchcontrollerRemotelog : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable comma-separated value (CSV) strings. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("csv")]
        public Output<string> Csv { get; private set; } = null!;

        /// <summary>
        /// Facility to log to remote syslog server. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
        /// </summary>
        [Output("facility")]
        public Output<string> Facility { get; private set; } = null!;

        /// <summary>
        /// Remote log name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Remote syslog server listening port.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// IPv4 address of the remote syslog server.
        /// </summary>
        [Output("server")]
        public Output<string> Server { get; private set; } = null!;

        /// <summary>
        /// Severity of logs to be transferred to remote log server. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        /// </summary>
        [Output("severity")]
        public Output<string> Severity { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging by FortiSwitch device to a remote syslog server. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SwitchcontrollerRemotelog resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SwitchcontrollerRemotelog(string name, SwitchcontrollerRemotelogArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerRemotelog:SwitchcontrollerRemotelog", name, args ?? new SwitchcontrollerRemotelogArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SwitchcontrollerRemotelog(string name, Input<string> id, SwitchcontrollerRemotelogState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerRemotelog:SwitchcontrollerRemotelog", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SwitchcontrollerRemotelog resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SwitchcontrollerRemotelog Get(string name, Input<string> id, SwitchcontrollerRemotelogState? state = null, CustomResourceOptions? options = null)
        {
            return new SwitchcontrollerRemotelog(name, id, state, options);
        }
    }

    public sealed class SwitchcontrollerRemotelogArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable comma-separated value (CSV) strings. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("csv")]
        public Input<string>? Csv { get; set; }

        /// <summary>
        /// Facility to log to remote syslog server. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
        /// </summary>
        [Input("facility")]
        public Input<string>? Facility { get; set; }

        /// <summary>
        /// Remote log name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Remote syslog server listening port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// IPv4 address of the remote syslog server.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Severity of logs to be transferred to remote log server. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// Enable/disable logging by FortiSwitch device to a remote syslog server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchcontrollerRemotelogArgs()
        {
        }
        public static new SwitchcontrollerRemotelogArgs Empty => new SwitchcontrollerRemotelogArgs();
    }

    public sealed class SwitchcontrollerRemotelogState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable comma-separated value (CSV) strings. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("csv")]
        public Input<string>? Csv { get; set; }

        /// <summary>
        /// Facility to log to remote syslog server. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
        /// </summary>
        [Input("facility")]
        public Input<string>? Facility { get; set; }

        /// <summary>
        /// Remote log name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Remote syslog server listening port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// IPv4 address of the remote syslog server.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Severity of logs to be transferred to remote log server. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// Enable/disable logging by FortiSwitch device to a remote syslog server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchcontrollerRemotelogState()
        {
        }
        public static new SwitchcontrollerRemotelogState Empty => new SwitchcontrollerRemotelogState();
    }
}