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
    public static class GetSystemReplacemsggroup
    {
        /// <summary>
        /// Use this data source to get information on an fortios system replacemsggroup
        /// </summary>
        public static Task<GetSystemReplacemsggroupResult> InvokeAsync(GetSystemReplacemsggroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemReplacemsggroupResult>("fortios:index/getSystemReplacemsggroup:getSystemReplacemsggroup", args ?? new GetSystemReplacemsggroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system replacemsggroup
        /// </summary>
        public static Output<GetSystemReplacemsggroupResult> Invoke(GetSystemReplacemsggroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemReplacemsggroupResult>("fortios:index/getSystemReplacemsggroup:getSystemReplacemsggroup", args ?? new GetSystemReplacemsggroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemReplacemsggroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system replacemsggroup.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemReplacemsggroupArgs()
        {
        }
        public static new GetSystemReplacemsggroupArgs Empty => new GetSystemReplacemsggroupArgs();
    }

    public sealed class GetSystemReplacemsggroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system replacemsggroup.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemReplacemsggroupInvokeArgs()
        {
        }
        public static new GetSystemReplacemsggroupInvokeArgs Empty => new GetSystemReplacemsggroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemReplacemsggroupResult
    {
        /// <summary>
        /// Replacement message table entries. The structure of `admin` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupAdminResult> Admins;
        /// <summary>
        /// Replacement message table entries. The structure of `alertmail` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupAlertmailResult> Alertmails;
        /// <summary>
        /// Replacement message table entries. The structure of `auth` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupAuthResult> Auths;
        /// <summary>
        /// Replacement message table entries. The structure of `automation` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupAutomationResult> Automations;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// Replacement message table entries. The structure of `custom_message` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupCustomMessageResult> CustomMessages;
        /// <summary>
        /// Replacement message table entries. The structure of `device_detection_portal` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupDeviceDetectionPortalResult> DeviceDetectionPortals;
        /// <summary>
        /// Replacement message table entries. The structure of `ec` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupEcResult> Ecs;
        /// <summary>
        /// Replacement message table entries. The structure of `fortiguard_wf` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupFortiguardWfResult> FortiguardWfs;
        /// <summary>
        /// Replacement message table entries. The structure of `ftp` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupFtpResult> Ftps;
        /// <summary>
        /// Group type.
        /// </summary>
        public readonly string GroupType;
        /// <summary>
        /// Replacement message table entries. The structure of `http` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupHttpResult> Https;
        /// <summary>
        /// Replacement message table entries. The structure of `icap` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupIcapResult> Icaps;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Replacement message table entries. The structure of `mail` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupMailResult> Mails;
        /// <summary>
        /// Replacement message table entries. The structure of `nac_quar` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupNacQuarResult> NacQuars;
        /// <summary>
        /// Group name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Replacement message table entries. The structure of `nntp` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupNntpResult> Nntps;
        /// <summary>
        /// Replacement message table entries. The structure of `spam` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupSpamResult> Spams;
        /// <summary>
        /// Replacement message table entries. The structure of `sslvpn` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupSslvpnResult> Sslvpns;
        /// <summary>
        /// Replacement message table entries. The structure of `traffic_quota` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupTrafficQuotaResult> TrafficQuotas;
        /// <summary>
        /// Replacement message table entries. The structure of `utm` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupUtmResult> Utms;
        public readonly string? Vdomparam;
        /// <summary>
        /// Replacement message table entries. The structure of `webproxy` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemReplacemsggroupWebproxyResult> Webproxies;

        [OutputConstructor]
        private GetSystemReplacemsggroupResult(
            ImmutableArray<Outputs.GetSystemReplacemsggroupAdminResult> admins,

            ImmutableArray<Outputs.GetSystemReplacemsggroupAlertmailResult> alertmails,

            ImmutableArray<Outputs.GetSystemReplacemsggroupAuthResult> auths,

            ImmutableArray<Outputs.GetSystemReplacemsggroupAutomationResult> automations,

            string comment,

            ImmutableArray<Outputs.GetSystemReplacemsggroupCustomMessageResult> customMessages,

            ImmutableArray<Outputs.GetSystemReplacemsggroupDeviceDetectionPortalResult> deviceDetectionPortals,

            ImmutableArray<Outputs.GetSystemReplacemsggroupEcResult> ecs,

            ImmutableArray<Outputs.GetSystemReplacemsggroupFortiguardWfResult> fortiguardWfs,

            ImmutableArray<Outputs.GetSystemReplacemsggroupFtpResult> ftps,

            string groupType,

            ImmutableArray<Outputs.GetSystemReplacemsggroupHttpResult> https,

            ImmutableArray<Outputs.GetSystemReplacemsggroupIcapResult> icaps,

            string id,

            ImmutableArray<Outputs.GetSystemReplacemsggroupMailResult> mails,

            ImmutableArray<Outputs.GetSystemReplacemsggroupNacQuarResult> nacQuars,

            string name,

            ImmutableArray<Outputs.GetSystemReplacemsggroupNntpResult> nntps,

            ImmutableArray<Outputs.GetSystemReplacemsggroupSpamResult> spams,

            ImmutableArray<Outputs.GetSystemReplacemsggroupSslvpnResult> sslvpns,

            ImmutableArray<Outputs.GetSystemReplacemsggroupTrafficQuotaResult> trafficQuotas,

            ImmutableArray<Outputs.GetSystemReplacemsggroupUtmResult> utms,

            string? vdomparam,

            ImmutableArray<Outputs.GetSystemReplacemsggroupWebproxyResult> webproxies)
        {
            Admins = admins;
            Alertmails = alertmails;
            Auths = auths;
            Automations = automations;
            Comment = comment;
            CustomMessages = customMessages;
            DeviceDetectionPortals = deviceDetectionPortals;
            Ecs = ecs;
            FortiguardWfs = fortiguardWfs;
            Ftps = ftps;
            GroupType = groupType;
            Https = https;
            Icaps = icaps;
            Id = id;
            Mails = mails;
            NacQuars = nacQuars;
            Name = name;
            Nntps = nntps;
            Spams = spams;
            Sslvpns = sslvpns;
            TrafficQuotas = trafficQuotas;
            Utms = utms;
            Vdomparam = vdomparam;
            Webproxies = webproxies;
        }
    }
}
