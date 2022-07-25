// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Storage.V1Alpha1
{

    /// <summary>
    /// VolumeAttachmentSpec is the specification of a VolumeAttachment request.
    /// </summary>
    [OutputType]
    public sealed class VolumeAttachmentSpecPatch
    {
        /// <summary>
        /// Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
        /// </summary>
        public readonly string Attacher;
        /// <summary>
        /// The node that the volume should be attached to.
        /// </summary>
        public readonly string NodeName;
        /// <summary>
        /// Source represents the volume that should be attached.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Storage.V1Alpha1.VolumeAttachmentSourcePatch Source;

        [OutputConstructor]
        private VolumeAttachmentSpecPatch(
            string attacher,

            string nodeName,

            Pulumi.Kubernetes.Types.Outputs.Storage.V1Alpha1.VolumeAttachmentSourcePatch source)
        {
            Attacher = attacher;
            NodeName = nodeName;
            Source = source;
        }
    }
}