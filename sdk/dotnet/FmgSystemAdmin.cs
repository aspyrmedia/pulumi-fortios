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
    /// This resource supports modifying system admin setting for FortiManager.
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
    ///     var test1 = new Fortios.FmgSystemAdmin("test1", new()
    ///     {
    ///         HttpPort = 80,
    ///         HttpsPort = 443,
    ///         IdleTimeout = 20,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:index/fmgSystemAdmin:FmgSystemAdmin")]
    public partial class FmgSystemAdmin : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Http port.
        /// </summary>
        [Output("httpPort")]
        public Output<int?> HttpPort { get; private set; } = null!;

        /// <summary>
        /// Https port.
        /// </summary>
        [Output("httpsPort")]
        public Output<int?> HttpsPort { get; private set; } = null!;

        /// <summary>
        /// Idle Timeout(1-480 minute).
        /// </summary>
        [Output("idleTimeout")]
        public Output<int?> IdleTimeout { get; private set; } = null!;


        /// <summary>
        /// Create a FmgSystemAdmin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FmgSystemAdmin(string name, FmgSystemAdminArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/fmgSystemAdmin:FmgSystemAdmin", name, args ?? new FmgSystemAdminArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FmgSystemAdmin(string name, Input<string> id, FmgSystemAdminState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/fmgSystemAdmin:FmgSystemAdmin", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FmgSystemAdmin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FmgSystemAdmin Get(string name, Input<string> id, FmgSystemAdminState? state = null, CustomResourceOptions? options = null)
        {
            return new FmgSystemAdmin(name, id, state, options);
        }
    }

    public sealed class FmgSystemAdminArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Http port.
        /// </summary>
        [Input("httpPort")]
        public Input<int>? HttpPort { get; set; }

        /// <summary>
        /// Https port.
        /// </summary>
        [Input("httpsPort")]
        public Input<int>? HttpsPort { get; set; }

        /// <summary>
        /// Idle Timeout(1-480 minute).
        /// </summary>
        [Input("idleTimeout")]
        public Input<int>? IdleTimeout { get; set; }

        public FmgSystemAdminArgs()
        {
        }
        public static new FmgSystemAdminArgs Empty => new FmgSystemAdminArgs();
    }

    public sealed class FmgSystemAdminState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Http port.
        /// </summary>
        [Input("httpPort")]
        public Input<int>? HttpPort { get; set; }

        /// <summary>
        /// Https port.
        /// </summary>
        [Input("httpsPort")]
        public Input<int>? HttpsPort { get; set; }

        /// <summary>
        /// Idle Timeout(1-480 minute).
        /// </summary>
        [Input("idleTimeout")]
        public Input<int>? IdleTimeout { get; set; }

        public FmgSystemAdminState()
        {
        }
        public static new FmgSystemAdminState Empty => new FmgSystemAdminState();
    }
}
