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

    public sealed class FirewallSecuritypolicyInternetService6CustomGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallSecuritypolicyInternetService6CustomGetArgs()
        {
        }
        public static new FirewallSecuritypolicyInternetService6CustomGetArgs Empty => new FirewallSecuritypolicyInternetService6CustomGetArgs();
    }
}
