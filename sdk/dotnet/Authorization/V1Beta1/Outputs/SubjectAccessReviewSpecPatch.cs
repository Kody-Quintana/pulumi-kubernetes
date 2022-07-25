// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Authorization.V1Beta1
{

    /// <summary>
    /// SubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
    /// </summary>
    [OutputType]
    public sealed class SubjectAccessReviewSpecPatch
    {
        /// <summary>
        /// Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here.
        /// </summary>
        public readonly ImmutableDictionary<string, ImmutableArray<string>> Extra;
        /// <summary>
        /// Groups is the groups you're testing for.
        /// </summary>
        public readonly ImmutableArray<string> Group;
        /// <summary>
        /// NonResourceAttributes describes information for a non-resource access request
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Authorization.V1Beta1.NonResourceAttributesPatch NonResourceAttributes;
        /// <summary>
        /// ResourceAuthorizationAttributes describes information for a resource access request
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Authorization.V1Beta1.ResourceAttributesPatch ResourceAttributes;
        /// <summary>
        /// UID information about the requesting user.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// User is the user you're testing for. If you specify "User" but not "Group", then is it interpreted as "What if User were not a member of any groups
        /// </summary>
        public readonly string User;

        [OutputConstructor]
        private SubjectAccessReviewSpecPatch(
            ImmutableDictionary<string, ImmutableArray<string>> extra,

            ImmutableArray<string> group,

            Pulumi.Kubernetes.Types.Outputs.Authorization.V1Beta1.NonResourceAttributesPatch nonResourceAttributes,

            Pulumi.Kubernetes.Types.Outputs.Authorization.V1Beta1.ResourceAttributesPatch resourceAttributes,

            string uid,

            string user)
        {
            Extra = extra;
            Group = group;
            NonResourceAttributes = nonResourceAttributes;
            ResourceAttributes = resourceAttributes;
            Uid = uid;
            User = user;
        }
    }
}