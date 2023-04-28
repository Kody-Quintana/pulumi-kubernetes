// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Resource.V1Alpha2
{

    /// <summary>
    /// AllocationResult contains attributes of an allocated resource.
    /// </summary>
    public class AllocationResultArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This field will get set by the resource driver after it has allocated the resource to inform the scheduler where it can schedule Pods using the ResourceClaim.
        /// 
        /// Setting this field is optional. If null, the resource is available everywhere.
        /// </summary>
        [Input("availableOnNodes")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeSelectorArgs>? AvailableOnNodes { get; set; }

        [Input("resourceHandles")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Resource.V1Alpha2.ResourceHandleArgs>? _resourceHandles;

        /// <summary>
        /// ResourceHandles contain the state associated with an allocation that should be maintained throughout the lifetime of a claim. Each ResourceHandle contains data that should be passed to a specific kubelet plugin once it lands on a node. This data is returned by the driver after a successful allocation and is opaque to Kubernetes. Driver documentation may explain to users how to interpret this data if needed.
        /// 
        /// Setting this field is optional. It has a maximum size of 32 entries. If null (or empty), it is assumed this allocation will be processed by a single kubelet plugin with no ResourceHandle data attached. The name of the kubelet plugin invoked will match the DriverName set in the ResourceClaimStatus this AllocationResult is embedded in.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Resource.V1Alpha2.ResourceHandleArgs> ResourceHandles
        {
            get => _resourceHandles ?? (_resourceHandles = new InputList<Pulumi.Kubernetes.Types.Inputs.Resource.V1Alpha2.ResourceHandleArgs>());
            set => _resourceHandles = value;
        }

        /// <summary>
        /// Shareable determines whether the resource supports more than one consumer at a time.
        /// </summary>
        [Input("shareable")]
        public Input<bool>? Shareable { get; set; }

        public AllocationResultArgs()
        {
        }
        public static new AllocationResultArgs Empty => new AllocationResultArgs();
    }
}