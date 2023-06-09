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
    public static class GetUserSaml
    {
        /// <summary>
        /// Use this data source to get information on an fortios user saml
        /// </summary>
        public static Task<GetUserSamlResult> InvokeAsync(GetUserSamlArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserSamlResult>("fortios:index/getUserSaml:getUserSaml", args ?? new GetUserSamlArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios user saml
        /// </summary>
        public static Output<GetUserSamlResult> Invoke(GetUserSamlInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserSamlResult>("fortios:index/getUserSaml:getUserSaml", args ?? new GetUserSamlInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserSamlArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired user saml.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetUserSamlArgs()
        {
        }
        public static new GetUserSamlArgs Empty => new GetUserSamlArgs();
    }

    public sealed class GetUserSamlInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired user saml.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetUserSamlInvokeArgs()
        {
        }
        public static new GetUserSamlInvokeArgs Empty => new GetUserSamlInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserSamlResult
    {
        /// <summary>
        /// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable).
        /// </summary>
        public readonly string AdfsClaim;
        /// <summary>
        /// URL to verify authentication.
        /// </summary>
        public readonly string AuthUrl;
        /// <summary>
        /// Certificate to sign SAML messages.
        /// </summary>
        public readonly string Cert;
        /// <summary>
        /// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
        /// </summary>
        public readonly int ClockTolerance;
        /// <summary>
        /// Digest Method Algorithm. (default = sha1).
        /// </summary>
        public readonly string DigestMethod;
        /// <summary>
        /// SP entity ID.
        /// </summary>
        public readonly string EntityId;
        /// <summary>
        /// Group claim in assertion statement.
        /// </summary>
        public readonly string GroupClaimType;
        /// <summary>
        /// Group name in assertion statement.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IDP Certificate name.
        /// </summary>
        public readonly string IdpCert;
        /// <summary>
        /// IDP entity ID.
        /// </summary>
        public readonly string IdpEntityId;
        /// <summary>
        /// IDP single logout url.
        /// </summary>
        public readonly string IdpSingleLogoutUrl;
        /// <summary>
        /// IDP single sign-on URL.
        /// </summary>
        public readonly string IdpSingleSignOnUrl;
        /// <summary>
        /// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes).
        /// </summary>
        public readonly string LimitRelaystate;
        /// <summary>
        /// SAML server entry name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// SP single logout URL.
        /// </summary>
        public readonly string SingleLogoutUrl;
        /// <summary>
        /// SP single sign-on URL.
        /// </summary>
        public readonly string SingleSignOnUrl;
        /// <summary>
        /// User name claim in assertion statement.
        /// </summary>
        public readonly string UserClaimType;
        /// <summary>
        /// User name in assertion statement.
        /// </summary>
        public readonly string UserName;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetUserSamlResult(
            string adfsClaim,

            string authUrl,

            string cert,

            int clockTolerance,

            string digestMethod,

            string entityId,

            string groupClaimType,

            string groupName,

            string id,

            string idpCert,

            string idpEntityId,

            string idpSingleLogoutUrl,

            string idpSingleSignOnUrl,

            string limitRelaystate,

            string name,

            string singleLogoutUrl,

            string singleSignOnUrl,

            string userClaimType,

            string userName,

            string? vdomparam)
        {
            AdfsClaim = adfsClaim;
            AuthUrl = authUrl;
            Cert = cert;
            ClockTolerance = clockTolerance;
            DigestMethod = digestMethod;
            EntityId = entityId;
            GroupClaimType = groupClaimType;
            GroupName = groupName;
            Id = id;
            IdpCert = idpCert;
            IdpEntityId = idpEntityId;
            IdpSingleLogoutUrl = idpSingleLogoutUrl;
            IdpSingleSignOnUrl = idpSingleSignOnUrl;
            LimitRelaystate = limitRelaystate;
            Name = name;
            SingleLogoutUrl = singleLogoutUrl;
            SingleSignOnUrl = singleSignOnUrl;
            UserClaimType = userClaimType;
            UserName = userName;
            Vdomparam = vdomparam;
        }
    }
}
