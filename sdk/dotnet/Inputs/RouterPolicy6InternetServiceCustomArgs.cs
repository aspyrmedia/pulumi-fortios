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

    public sealed class RouterPolicy6InternetServiceCustomArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom Destination Internet Service name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public RouterPolicy6InternetServiceCustomArgs()
        {
        }
        public static new RouterPolicy6InternetServiceCustomArgs Empty => new RouterPolicy6InternetServiceCustomArgs();
    }
}
