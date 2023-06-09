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
    public static class GetSystemFm
    {
        /// <summary>
        /// Use this data source to get information on fortios system fm
        /// </summary>
        public static Task<GetSystemFmResult> InvokeAsync(GetSystemFmArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemFmResult>("fortios:index/getSystemFm:getSystemFm", args ?? new GetSystemFmArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system fm
        /// </summary>
        public static Output<GetSystemFmResult> Invoke(GetSystemFmInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemFmResult>("fortios:index/getSystemFm:getSystemFm", args ?? new GetSystemFmInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemFmArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemFmArgs()
        {
        }
        public static new GetSystemFmArgs Empty => new GetSystemFmArgs();
    }

    public sealed class GetSystemFmInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemFmInvokeArgs()
        {
        }
        public static new GetSystemFmInvokeArgs Empty => new GetSystemFmInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemFmResult
    {
        /// <summary>
        /// Enable/disable automatic backup.
        /// </summary>
        public readonly string AutoBackup;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly string Fosid;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IP address.
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// Enable/disable IPsec.
        /// </summary>
        public readonly string Ipsec;
        /// <summary>
        /// Enable/disable scheduled configuration restore.
        /// </summary>
        public readonly string ScheduledConfigRestore;
        /// <summary>
        /// Enable/disable FM.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// VDOM.
        /// </summary>
        public readonly string Vdom;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemFmResult(
            string autoBackup,

            string fosid,

            string id,

            string ip,

            string ipsec,

            string scheduledConfigRestore,

            string status,

            string vdom,

            string? vdomparam)
        {
            AutoBackup = autoBackup;
            Fosid = fosid;
            Id = id;
            Ip = ip;
            Ipsec = ipsec;
            ScheduledConfigRestore = scheduledConfigRestore;
            Status = status;
            Vdom = vdom;
            Vdomparam = vdomparam;
        }
    }
}
