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
    /// Configure FortiToken.
    /// 
    /// ## Import
    /// 
    /// User Fortitoken can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/userFortitoken:UserFortitoken labelname {{serial_number}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/userFortitoken:UserFortitoken labelname {{serial_number}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/userFortitoken:UserFortitoken")]
    public partial class UserFortitoken : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Mobile token user activation-code.
        /// </summary>
        [Output("activationCode")]
        public Output<string> ActivationCode { get; private set; } = null!;

        /// <summary>
        /// Mobile token user activation-code expire time.
        /// </summary>
        [Output("activationExpire")]
        public Output<int> ActivationExpire { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        /// <summary>
        /// Mobile token license.
        /// </summary>
        [Output("license")]
        public Output<string> License { get; private set; } = null!;

        /// <summary>
        /// Device Mobile Version.
        /// </summary>
        [Output("osVer")]
        public Output<string> OsVer { get; private set; } = null!;

        /// <summary>
        /// Device Reg ID.
        /// </summary>
        [Output("regId")]
        public Output<string> RegId { get; private set; } = null!;

        /// <summary>
        /// Token seed.
        /// </summary>
        [Output("seed")]
        public Output<string> Seed { get; private set; } = null!;

        /// <summary>
        /// Serial number.
        /// </summary>
        [Output("serialNumber")]
        public Output<string> SerialNumber { get; private set; } = null!;

        /// <summary>
        /// Status Valid values: `active`, `lock`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a UserFortitoken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserFortitoken(string name, UserFortitokenArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/userFortitoken:UserFortitoken", name, args ?? new UserFortitokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserFortitoken(string name, Input<string> id, UserFortitokenState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/userFortitoken:UserFortitoken", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserFortitoken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserFortitoken Get(string name, Input<string> id, UserFortitokenState? state = null, CustomResourceOptions? options = null)
        {
            return new UserFortitoken(name, id, state, options);
        }
    }

    public sealed class UserFortitokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mobile token user activation-code.
        /// </summary>
        [Input("activationCode")]
        public Input<string>? ActivationCode { get; set; }

        /// <summary>
        /// Mobile token user activation-code expire time.
        /// </summary>
        [Input("activationExpire")]
        public Input<int>? ActivationExpire { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Mobile token license.
        /// </summary>
        [Input("license")]
        public Input<string>? License { get; set; }

        /// <summary>
        /// Device Mobile Version.
        /// </summary>
        [Input("osVer")]
        public Input<string>? OsVer { get; set; }

        /// <summary>
        /// Device Reg ID.
        /// </summary>
        [Input("regId")]
        public Input<string>? RegId { get; set; }

        /// <summary>
        /// Token seed.
        /// </summary>
        [Input("seed")]
        public Input<string>? Seed { get; set; }

        /// <summary>
        /// Serial number.
        /// </summary>
        [Input("serialNumber")]
        public Input<string>? SerialNumber { get; set; }

        /// <summary>
        /// Status Valid values: `active`, `lock`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public UserFortitokenArgs()
        {
        }
        public static new UserFortitokenArgs Empty => new UserFortitokenArgs();
    }

    public sealed class UserFortitokenState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mobile token user activation-code.
        /// </summary>
        [Input("activationCode")]
        public Input<string>? ActivationCode { get; set; }

        /// <summary>
        /// Mobile token user activation-code expire time.
        /// </summary>
        [Input("activationExpire")]
        public Input<int>? ActivationExpire { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Mobile token license.
        /// </summary>
        [Input("license")]
        public Input<string>? License { get; set; }

        /// <summary>
        /// Device Mobile Version.
        /// </summary>
        [Input("osVer")]
        public Input<string>? OsVer { get; set; }

        /// <summary>
        /// Device Reg ID.
        /// </summary>
        [Input("regId")]
        public Input<string>? RegId { get; set; }

        /// <summary>
        /// Token seed.
        /// </summary>
        [Input("seed")]
        public Input<string>? Seed { get; set; }

        /// <summary>
        /// Serial number.
        /// </summary>
        [Input("serialNumber")]
        public Input<string>? SerialNumber { get; set; }

        /// <summary>
        /// Status Valid values: `active`, `lock`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public UserFortitokenState()
        {
        }
        public static new UserFortitokenState Empty => new UserFortitokenState();
    }
}
