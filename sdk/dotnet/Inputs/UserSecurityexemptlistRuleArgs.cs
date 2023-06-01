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

    public sealed class UserSecurityexemptlistRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("devices")]
        private InputList<Inputs.UserSecurityexemptlistRuleDeviceArgs>? _devices;

        /// <summary>
        /// Devices or device groups. The structure of `devices` block is documented below.
        /// </summary>
        public InputList<Inputs.UserSecurityexemptlistRuleDeviceArgs> Devices
        {
            get => _devices ?? (_devices = new InputList<Inputs.UserSecurityexemptlistRuleDeviceArgs>());
            set => _devices = value;
        }

        [Input("dstaddrs")]
        private InputList<Inputs.UserSecurityexemptlistRuleDstaddrArgs>? _dstaddrs;

        /// <summary>
        /// Destination addresses or address groups. The structure of `dstaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.UserSecurityexemptlistRuleDstaddrArgs> Dstaddrs
        {
            get => _dstaddrs ?? (_dstaddrs = new InputList<Inputs.UserSecurityexemptlistRuleDstaddrArgs>());
            set => _dstaddrs = value;
        }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("services")]
        private InputList<Inputs.UserSecurityexemptlistRuleServiceArgs>? _services;

        /// <summary>
        /// Destination services. The structure of `service` block is documented below.
        /// </summary>
        public InputList<Inputs.UserSecurityexemptlistRuleServiceArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.UserSecurityexemptlistRuleServiceArgs>());
            set => _services = value;
        }

        [Input("srcaddrs")]
        private InputList<Inputs.UserSecurityexemptlistRuleSrcaddrArgs>? _srcaddrs;

        /// <summary>
        /// Source addresses or address groups. The structure of `srcaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.UserSecurityexemptlistRuleSrcaddrArgs> Srcaddrs
        {
            get => _srcaddrs ?? (_srcaddrs = new InputList<Inputs.UserSecurityexemptlistRuleSrcaddrArgs>());
            set => _srcaddrs = value;
        }

        public UserSecurityexemptlistRuleArgs()
        {
        }
        public static new UserSecurityexemptlistRuleArgs Empty => new UserSecurityexemptlistRuleArgs();
    }
}