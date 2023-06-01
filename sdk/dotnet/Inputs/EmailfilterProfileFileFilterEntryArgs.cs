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

    public sealed class EmailfilterProfileFileFilterEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action taken for matched file. Valid values: `log`, `block`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("fileTypes")]
        private InputList<Inputs.EmailfilterProfileFileFilterEntryFileTypeArgs>? _fileTypes;

        /// <summary>
        /// Select file type. The structure of `file_type` block is documented below.
        /// </summary>
        public InputList<Inputs.EmailfilterProfileFileFilterEntryFileTypeArgs> FileTypes
        {
            get => _fileTypes ?? (_fileTypes = new InputList<Inputs.EmailfilterProfileFileFilterEntryFileTypeArgs>());
            set => _fileTypes = value;
        }

        /// <summary>
        /// Add a file filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// Match password-protected files. Valid values: `yes`, `any`.
        /// </summary>
        [Input("passwordProtected")]
        public Input<string>? PasswordProtected { get; set; }

        /// <summary>
        /// Protocols to apply with. Valid values: `smtp`, `imap`, `pop3`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public EmailfilterProfileFileFilterEntryArgs()
        {
        }
        public static new EmailfilterProfileFileFilterEntryArgs Empty => new EmailfilterProfileFileFilterEntryArgs();
    }
}