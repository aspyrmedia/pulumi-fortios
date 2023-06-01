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
    /// Configure debug URL addresses.
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
    ///     var trname = new Fortios.WebproxyDebugurl("trname", new()
    ///     {
    ///         Exact = "enable",
    ///         Status = "enable",
    ///         UrlPattern = "/examples/servlet/*Servlet",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// WebProxy DebugUrl can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/webproxyDebugurl:WebproxyDebugurl labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/webproxyDebugurl:WebproxyDebugurl labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/webproxyDebugurl:WebproxyDebugurl")]
    public partial class WebproxyDebugurl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable matching the exact path. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("exact")]
        public Output<string> Exact { get; private set; } = null!;

        /// <summary>
        /// Debug URL name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this URL exemption. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// URL exemption pattern.
        /// </summary>
        [Output("urlPattern")]
        public Output<string> UrlPattern { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a WebproxyDebugurl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebproxyDebugurl(string name, WebproxyDebugurlArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/webproxyDebugurl:WebproxyDebugurl", name, args ?? new WebproxyDebugurlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebproxyDebugurl(string name, Input<string> id, WebproxyDebugurlState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/webproxyDebugurl:WebproxyDebugurl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WebproxyDebugurl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebproxyDebugurl Get(string name, Input<string> id, WebproxyDebugurlState? state = null, CustomResourceOptions? options = null)
        {
            return new WebproxyDebugurl(name, id, state, options);
        }
    }

    public sealed class WebproxyDebugurlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable matching the exact path. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("exact")]
        public Input<string>? Exact { get; set; }

        /// <summary>
        /// Debug URL name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable this URL exemption. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// URL exemption pattern.
        /// </summary>
        [Input("urlPattern", required: true)]
        public Input<string> UrlPattern { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public WebproxyDebugurlArgs()
        {
        }
        public static new WebproxyDebugurlArgs Empty => new WebproxyDebugurlArgs();
    }

    public sealed class WebproxyDebugurlState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable matching the exact path. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("exact")]
        public Input<string>? Exact { get; set; }

        /// <summary>
        /// Debug URL name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable this URL exemption. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// URL exemption pattern.
        /// </summary>
        [Input("urlPattern")]
        public Input<string>? UrlPattern { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public WebproxyDebugurlState()
        {
        }
        public static new WebproxyDebugurlState Empty => new WebproxyDebugurlState();
    }
}