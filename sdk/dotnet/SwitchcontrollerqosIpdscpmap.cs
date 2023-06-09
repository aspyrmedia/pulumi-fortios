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
    /// Configure FortiSwitch QoS IP precedence/DSCP.
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
    ///     var trname = new Fortios.SwitchcontrollerqosIpdscpmap("trname", new()
    ///     {
    ///         Description = "DEIW",
    ///         Maps = new[]
    ///         {
    ///             new Fortios.Inputs.SwitchcontrollerqosIpdscpmapMapArgs
    ///             {
    ///                 CosQueue = 3,
    ///                 Diffserv = "CS0 CS1 AF11",
    ///                 Name = "1",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SwitchControllerQos IpDscpMap can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerqosIpdscpmap:SwitchcontrollerqosIpdscpmap labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerqosIpdscpmap:SwitchcontrollerqosIpdscpmap labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/switchcontrollerqosIpdscpmap:SwitchcontrollerqosIpdscpmap")]
    public partial class SwitchcontrollerqosIpdscpmap : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the ip-dscp map name.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
        /// </summary>
        [Output("maps")]
        public Output<ImmutableArray<Outputs.SwitchcontrollerqosIpdscpmapMap>> Maps { get; private set; } = null!;

        /// <summary>
        /// Dscp map name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SwitchcontrollerqosIpdscpmap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SwitchcontrollerqosIpdscpmap(string name, SwitchcontrollerqosIpdscpmapArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerqosIpdscpmap:SwitchcontrollerqosIpdscpmap", name, args ?? new SwitchcontrollerqosIpdscpmapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SwitchcontrollerqosIpdscpmap(string name, Input<string> id, SwitchcontrollerqosIpdscpmapState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerqosIpdscpmap:SwitchcontrollerqosIpdscpmap", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SwitchcontrollerqosIpdscpmap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SwitchcontrollerqosIpdscpmap Get(string name, Input<string> id, SwitchcontrollerqosIpdscpmapState? state = null, CustomResourceOptions? options = null)
        {
            return new SwitchcontrollerqosIpdscpmap(name, id, state, options);
        }
    }

    public sealed class SwitchcontrollerqosIpdscpmapArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the ip-dscp map name.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("maps")]
        private InputList<Inputs.SwitchcontrollerqosIpdscpmapMapArgs>? _maps;

        /// <summary>
        /// Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
        /// </summary>
        public InputList<Inputs.SwitchcontrollerqosIpdscpmapMapArgs> Maps
        {
            get => _maps ?? (_maps = new InputList<Inputs.SwitchcontrollerqosIpdscpmapMapArgs>());
            set => _maps = value;
        }

        /// <summary>
        /// Dscp map name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchcontrollerqosIpdscpmapArgs()
        {
        }
        public static new SwitchcontrollerqosIpdscpmapArgs Empty => new SwitchcontrollerqosIpdscpmapArgs();
    }

    public sealed class SwitchcontrollerqosIpdscpmapState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the ip-dscp map name.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("maps")]
        private InputList<Inputs.SwitchcontrollerqosIpdscpmapMapGetArgs>? _maps;

        /// <summary>
        /// Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
        /// </summary>
        public InputList<Inputs.SwitchcontrollerqosIpdscpmapMapGetArgs> Maps
        {
            get => _maps ?? (_maps = new InputList<Inputs.SwitchcontrollerqosIpdscpmapMapGetArgs>());
            set => _maps = value;
        }

        /// <summary>
        /// Dscp map name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchcontrollerqosIpdscpmapState()
        {
        }
        public static new SwitchcontrollerqosIpdscpmapState Empty => new SwitchcontrollerqosIpdscpmapState();
    }
}
