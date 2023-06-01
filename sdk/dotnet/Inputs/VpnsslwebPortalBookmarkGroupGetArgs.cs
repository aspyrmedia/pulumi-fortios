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

    public sealed class VpnsslwebPortalBookmarkGroupGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("bookmarks")]
        private InputList<Inputs.VpnsslwebPortalBookmarkGroupBookmarkGetArgs>? _bookmarks;

        /// <summary>
        /// Bookmark table. The structure of `bookmarks` block is documented below.
        /// </summary>
        public InputList<Inputs.VpnsslwebPortalBookmarkGroupBookmarkGetArgs> Bookmarks
        {
            get => _bookmarks ?? (_bookmarks = new InputList<Inputs.VpnsslwebPortalBookmarkGroupBookmarkGetArgs>());
            set => _bookmarks = value;
        }

        /// <summary>
        /// Bookmark group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public VpnsslwebPortalBookmarkGroupGetArgs()
        {
        }
        public static new VpnsslwebPortalBookmarkGroupGetArgs Empty => new VpnsslwebPortalBookmarkGroupGetArgs();
    }
}