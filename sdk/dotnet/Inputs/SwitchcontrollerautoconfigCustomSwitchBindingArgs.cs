// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class SwitchcontrollerautoconfigCustomSwitchBindingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom auto-config policy.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// Switch name.
        /// </summary>
        [Input("switchId")]
        public Input<string>? SwitchId { get; set; }

        public SwitchcontrollerautoconfigCustomSwitchBindingArgs()
        {
        }
        public static new SwitchcontrollerautoconfigCustomSwitchBindingArgs Empty => new SwitchcontrollerautoconfigCustomSwitchBindingArgs();
    }
}
