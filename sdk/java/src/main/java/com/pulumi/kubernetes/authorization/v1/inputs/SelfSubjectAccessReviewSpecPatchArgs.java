// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.authorization.v1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetes.authorization.v1.inputs.NonResourceAttributesPatchArgs;
import com.pulumi.kubernetes.authorization.v1.inputs.ResourceAttributesPatchArgs;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * SelfSubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
 * 
 */
public final class SelfSubjectAccessReviewSpecPatchArgs extends com.pulumi.resources.ResourceArgs {

    public static final SelfSubjectAccessReviewSpecPatchArgs Empty = new SelfSubjectAccessReviewSpecPatchArgs();

    /**
     * NonResourceAttributes describes information for a non-resource access request
     * 
     */
    @Import(name="nonResourceAttributes")
    private @Nullable Output<NonResourceAttributesPatchArgs> nonResourceAttributes;

    /**
     * @return NonResourceAttributes describes information for a non-resource access request
     * 
     */
    public Optional<Output<NonResourceAttributesPatchArgs>> nonResourceAttributes() {
        return Optional.ofNullable(this.nonResourceAttributes);
    }

    /**
     * ResourceAuthorizationAttributes describes information for a resource access request
     * 
     */
    @Import(name="resourceAttributes")
    private @Nullable Output<ResourceAttributesPatchArgs> resourceAttributes;

    /**
     * @return ResourceAuthorizationAttributes describes information for a resource access request
     * 
     */
    public Optional<Output<ResourceAttributesPatchArgs>> resourceAttributes() {
        return Optional.ofNullable(this.resourceAttributes);
    }

    private SelfSubjectAccessReviewSpecPatchArgs() {}

    private SelfSubjectAccessReviewSpecPatchArgs(SelfSubjectAccessReviewSpecPatchArgs $) {
        this.nonResourceAttributes = $.nonResourceAttributes;
        this.resourceAttributes = $.resourceAttributes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SelfSubjectAccessReviewSpecPatchArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SelfSubjectAccessReviewSpecPatchArgs $;

        public Builder() {
            $ = new SelfSubjectAccessReviewSpecPatchArgs();
        }

        public Builder(SelfSubjectAccessReviewSpecPatchArgs defaults) {
            $ = new SelfSubjectAccessReviewSpecPatchArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param nonResourceAttributes NonResourceAttributes describes information for a non-resource access request
         * 
         * @return builder
         * 
         */
        public Builder nonResourceAttributes(@Nullable Output<NonResourceAttributesPatchArgs> nonResourceAttributes) {
            $.nonResourceAttributes = nonResourceAttributes;
            return this;
        }

        /**
         * @param nonResourceAttributes NonResourceAttributes describes information for a non-resource access request
         * 
         * @return builder
         * 
         */
        public Builder nonResourceAttributes(NonResourceAttributesPatchArgs nonResourceAttributes) {
            return nonResourceAttributes(Output.of(nonResourceAttributes));
        }

        /**
         * @param resourceAttributes ResourceAuthorizationAttributes describes information for a resource access request
         * 
         * @return builder
         * 
         */
        public Builder resourceAttributes(@Nullable Output<ResourceAttributesPatchArgs> resourceAttributes) {
            $.resourceAttributes = resourceAttributes;
            return this;
        }

        /**
         * @param resourceAttributes ResourceAuthorizationAttributes describes information for a resource access request
         * 
         * @return builder
         * 
         */
        public Builder resourceAttributes(ResourceAttributesPatchArgs resourceAttributes) {
            return resourceAttributes(Output.of(resourceAttributes));
        }

        public SelfSubjectAccessReviewSpecPatchArgs build() {
            return $;
        }
    }

}