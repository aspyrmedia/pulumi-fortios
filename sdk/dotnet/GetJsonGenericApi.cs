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
    public static class GetJsonGenericApi
    {
        /// <summary>
        /// Provides a FortiAPI Generic Interface data source.
        /// </summary>
        public static Task<GetJsonGenericApiResult> InvokeAsync(GetJsonGenericApiArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetJsonGenericApiResult>("fortios:index/getJsonGenericApi:getJsonGenericApi", args ?? new GetJsonGenericApiArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a FortiAPI Generic Interface data source.
        /// </summary>
        public static Output<GetJsonGenericApiResult> Invoke(GetJsonGenericApiInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetJsonGenericApiResult>("fortios:index/getJsonGenericApi:getJsonGenericApi", args ?? new GetJsonGenericApiInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetJsonGenericApiArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// FortiAPI URL path.
        /// </summary>
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        /// <summary>
        /// URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..
        /// </summary>
        [Input("specialparams")]
        public string? Specialparams { get; set; }

        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetJsonGenericApiArgs()
        {
        }
        public static new GetJsonGenericApiArgs Empty => new GetJsonGenericApiArgs();
    }

    public sealed class GetJsonGenericApiInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// FortiAPI URL path.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..
        /// </summary>
        [Input("specialparams")]
        public Input<string>? Specialparams { get; set; }

        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetJsonGenericApiInvokeArgs()
        {
        }
        public static new GetJsonGenericApiInvokeArgs Empty => new GetJsonGenericApiInvokeArgs();
    }


    [OutputType]
    public sealed class GetJsonGenericApiResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// FortiAPI URL path.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// FortiAPI returns results.
        /// </summary>
        public readonly string Response;
        /// <summary>
        /// URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..
        /// </summary>
        public readonly string? Specialparams;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetJsonGenericApiResult(
            string id,

            string path,

            string response,

            string? specialparams,

            string? vdomparam)
        {
            Id = id;
            Path = path;
            Response = response;
            Specialparams = specialparams;
            Vdomparam = vdomparam;
        }
    }
}
