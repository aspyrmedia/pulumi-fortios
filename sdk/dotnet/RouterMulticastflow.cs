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
    /// Configure multicast-flow.
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
    ///     var trname = new Fortios.RouterMulticastflow("trname", new()
    ///     {
    ///         Flows = new[]
    ///         {
    ///             new Fortios.Inputs.RouterMulticastflowFlowArgs
    ///             {
    ///                 GroupAddr = "224.252.0.0",
    ///                 SourceAddr = "224.112.0.0",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Router MulticastFlow can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/routerMulticastflow:RouterMulticastflow labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/routerMulticastflow:RouterMulticastflow labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/routerMulticastflow:RouterMulticastflow")]
    public partial class RouterMulticastflow : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string> Comments { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Multicast-flow entries. The structure of `flows` block is documented below.
        /// </summary>
        [Output("flows")]
        public Output<ImmutableArray<Outputs.RouterMulticastflowFlow>> Flows { get; private set; } = null!;

        /// <summary>
        /// Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a RouterMulticastflow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouterMulticastflow(string name, RouterMulticastflowArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/routerMulticastflow:RouterMulticastflow", name, args ?? new RouterMulticastflowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouterMulticastflow(string name, Input<string> id, RouterMulticastflowState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/routerMulticastflow:RouterMulticastflow", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouterMulticastflow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouterMulticastflow Get(string name, Input<string> id, RouterMulticastflowState? state = null, CustomResourceOptions? options = null)
        {
            return new RouterMulticastflow(name, id, state, options);
        }
    }

    public sealed class RouterMulticastflowArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("flows")]
        private InputList<Inputs.RouterMulticastflowFlowArgs>? _flows;

        /// <summary>
        /// Multicast-flow entries. The structure of `flows` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterMulticastflowFlowArgs> Flows
        {
            get => _flows ?? (_flows = new InputList<Inputs.RouterMulticastflowFlowArgs>());
            set => _flows = value;
        }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public RouterMulticastflowArgs()
        {
        }
        public static new RouterMulticastflowArgs Empty => new RouterMulticastflowArgs();
    }

    public sealed class RouterMulticastflowState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("flows")]
        private InputList<Inputs.RouterMulticastflowFlowGetArgs>? _flows;

        /// <summary>
        /// Multicast-flow entries. The structure of `flows` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterMulticastflowFlowGetArgs> Flows
        {
            get => _flows ?? (_flows = new InputList<Inputs.RouterMulticastflowFlowGetArgs>());
            set => _flows = value;
        }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public RouterMulticastflowState()
        {
        }
        public static new RouterMulticastflowState Empty => new RouterMulticastflowState();
    }
}
