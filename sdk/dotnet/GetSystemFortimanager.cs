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
    public static class GetSystemFortimanager
    {
        /// <summary>
        /// Use this data source to get information on fortios system fortimanager
        /// </summary>
        public static Task<GetSystemFortimanagerResult> InvokeAsync(GetSystemFortimanagerArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemFortimanagerResult>("fortios:index/getSystemFortimanager:getSystemFortimanager", args ?? new GetSystemFortimanagerArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system fortimanager
        /// </summary>
        public static Output<GetSystemFortimanagerResult> Invoke(GetSystemFortimanagerInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemFortimanagerResult>("fortios:index/getSystemFortimanager:getSystemFortimanager", args ?? new GetSystemFortimanagerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemFortimanagerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemFortimanagerArgs()
        {
        }
        public static new GetSystemFortimanagerArgs Empty => new GetSystemFortimanagerArgs();
    }

    public sealed class GetSystemFortimanagerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemFortimanagerInvokeArgs()
        {
        }
        public static new GetSystemFortimanagerInvokeArgs Empty => new GetSystemFortimanagerInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemFortimanagerResult
    {
        /// <summary>
        /// Enable/disable FortiManager central management.
        /// </summary>
        public readonly string CentralManagement;
        /// <summary>
        /// Enable/disable central management auto backup.
        /// </summary>
        public readonly string CentralMgmtAutoBackup;
        /// <summary>
        /// Enable/disable central management schedule config restore.
        /// </summary>
        public readonly string CentralMgmtScheduleConfigRestore;
        /// <summary>
        /// Enable/disable central management schedule script restore.
        /// </summary>
        public readonly string CentralMgmtScheduleScriptRestore;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IP address.
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// Enable/disable FortiManager IPsec tunnel.
        /// </summary>
        public readonly string Ipsec;
        /// <summary>
        /// Virtual domain name.
        /// </summary>
        public readonly string Vdom;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemFortimanagerResult(
            string centralManagement,

            string centralMgmtAutoBackup,

            string centralMgmtScheduleConfigRestore,

            string centralMgmtScheduleScriptRestore,

            string id,

            string ip,

            string ipsec,

            string vdom,

            string? vdomparam)
        {
            CentralManagement = centralManagement;
            CentralMgmtAutoBackup = centralMgmtAutoBackup;
            CentralMgmtScheduleConfigRestore = centralMgmtScheduleConfigRestore;
            CentralMgmtScheduleScriptRestore = centralMgmtScheduleScriptRestore;
            Id = id;
            Ip = ip;
            Ipsec = ipsec;
            Vdom = vdom;
            Vdomparam = vdomparam;
        }
    }
}