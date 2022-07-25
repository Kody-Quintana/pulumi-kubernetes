// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.core.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.core.v1.outputs.NodeSelectorPatch;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VolumeNodeAffinityPatch {
    /**
     * @return required specifies hard node constraints that must be met.
     * 
     */
    private @Nullable NodeSelectorPatch required;

    private VolumeNodeAffinityPatch() {}
    /**
     * @return required specifies hard node constraints that must be met.
     * 
     */
    public Optional<NodeSelectorPatch> required() {
        return Optional.ofNullable(this.required);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VolumeNodeAffinityPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable NodeSelectorPatch required;
        public Builder() {}
        public Builder(VolumeNodeAffinityPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.required = defaults.required;
        }

        @CustomType.Setter
        public Builder required(@Nullable NodeSelectorPatch required) {
            this.required = required;
            return this;
        }
        public VolumeNodeAffinityPatch build() {
            final var o = new VolumeNodeAffinityPatch();
            o.required = required;
            return o;
        }
    }
}