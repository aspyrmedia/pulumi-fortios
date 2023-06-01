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
    /// Configure FortiSwitch Auto-Config custom QoS policy.
    /// 
    /// ## Import
    /// 
    /// SwitchControllerAutoConfig Custom can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerautoconfigCustom:SwitchcontrollerautoconfigCustom labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerautoconfigCustom:SwitchcontrollerautoconfigCustom labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/switchcontrollerautoconfigCustom:SwitchcontrollerautoconfigCustom")]
    public partial class SwitchcontrollerautoconfigCustom : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Auto-Config FortiLink or ISL/ICL interface name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Switch binding list. The structure of `switch_binding` block is documented below.
        /// </summary>
        [Output("switchBindings")]
        public Output<ImmutableArray<Outputs.SwitchcontrollerautoconfigCustomSwitchBinding>> SwitchBindings { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SwitchcontrollerautoconfigCustom resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SwitchcontrollerautoconfigCustom(string name, SwitchcontrollerautoconfigCustomArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerautoconfigCustom:SwitchcontrollerautoconfigCustom", name, args ?? new SwitchcontrollerautoconfigCustomArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SwitchcontrollerautoconfigCustom(string name, Input<string> id, SwitchcontrollerautoconfigCustomState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerautoconfigCustom:SwitchcontrollerautoconfigCustom", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SwitchcontrollerautoconfigCustom resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SwitchcontrollerautoconfigCustom Get(string name, Input<string> id, SwitchcontrollerautoconfigCustomState? state = null, CustomResourceOptions? options = null)
        {
            return new SwitchcontrollerautoconfigCustom(name, id, state, options);
        }
    }

    public sealed class SwitchcontrollerautoconfigCustomArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Auto-Config FortiLink or ISL/ICL interface name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("switchBindings")]
        private InputList<Inputs.SwitchcontrollerautoconfigCustomSwitchBindingArgs>? _switchBindings;

        /// <summary>
        /// Switch binding list. The structure of `switch_binding` block is documented below.
        /// </summary>
        public InputList<Inputs.SwitchcontrollerautoconfigCustomSwitchBindingArgs> SwitchBindings
        {
            get => _switchBindings ?? (_switchBindings = new InputList<Inputs.SwitchcontrollerautoconfigCustomSwitchBindingArgs>());
            set => _switchBindings = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchcontrollerautoconfigCustomArgs()
        {
        }
        public static new SwitchcontrollerautoconfigCustomArgs Empty => new SwitchcontrollerautoconfigCustomArgs();
    }

    public sealed class SwitchcontrollerautoconfigCustomState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Auto-Config FortiLink or ISL/ICL interface name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("switchBindings")]
        private InputList<Inputs.SwitchcontrollerautoconfigCustomSwitchBindingGetArgs>? _switchBindings;

        /// <summary>
        /// Switch binding list. The structure of `switch_binding` block is documented below.
        /// </summary>
        public InputList<Inputs.SwitchcontrollerautoconfigCustomSwitchBindingGetArgs> SwitchBindings
        {
            get => _switchBindings ?? (_switchBindings = new InputList<Inputs.SwitchcontrollerautoconfigCustomSwitchBindingGetArgs>());
            set => _switchBindings = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchcontrollerautoconfigCustomState()
        {
        }
        public static new SwitchcontrollerautoconfigCustomState Empty => new SwitchcontrollerautoconfigCustomState();
    }
}