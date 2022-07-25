// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.core.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.kubernetes.core.v1.outputs.PersistentVolumeClaimConditionPatch;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PersistentVolumeClaimStatusPatch {
    /**
     * @return accessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
     * 
     */
    private @Nullable List<String> accessModes;
    /**
     * @return allocatedResources is the storage resource within AllocatedResources tracks the capacity allocated to a PVC. It may be larger than the actual capacity when a volume expansion operation is requested. For storage quota, the larger value from allocatedResources and PVC.spec.resources is used. If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation. If a volume expansion capacity request is lowered, allocatedResources is only lowered if there are no expansion operations in progress and if the actual volume capacity is equal or lower than the requested capacity. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.
     * 
     */
    private @Nullable Map<String,String> allocatedResources;
    /**
     * @return capacity represents the actual resources of the underlying volume.
     * 
     */
    private @Nullable Map<String,String> capacity;
    /**
     * @return conditions is the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to &#39;ResizeStarted&#39;.
     * 
     */
    private @Nullable List<PersistentVolumeClaimConditionPatch> conditions;
    /**
     * @return phase represents the current phase of PersistentVolumeClaim.
     * 
     */
    private @Nullable String phase;
    /**
     * @return resizeStatus stores status of resize operation. ResizeStatus is not set by default but when expansion is complete resizeStatus is set to empty string by resize controller or kubelet. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.
     * 
     */
    private @Nullable String resizeStatus;

    private PersistentVolumeClaimStatusPatch() {}
    /**
     * @return accessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
     * 
     */
    public List<String> accessModes() {
        return this.accessModes == null ? List.of() : this.accessModes;
    }
    /**
     * @return allocatedResources is the storage resource within AllocatedResources tracks the capacity allocated to a PVC. It may be larger than the actual capacity when a volume expansion operation is requested. For storage quota, the larger value from allocatedResources and PVC.spec.resources is used. If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation. If a volume expansion capacity request is lowered, allocatedResources is only lowered if there are no expansion operations in progress and if the actual volume capacity is equal or lower than the requested capacity. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.
     * 
     */
    public Map<String,String> allocatedResources() {
        return this.allocatedResources == null ? Map.of() : this.allocatedResources;
    }
    /**
     * @return capacity represents the actual resources of the underlying volume.
     * 
     */
    public Map<String,String> capacity() {
        return this.capacity == null ? Map.of() : this.capacity;
    }
    /**
     * @return conditions is the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to &#39;ResizeStarted&#39;.
     * 
     */
    public List<PersistentVolumeClaimConditionPatch> conditions() {
        return this.conditions == null ? List.of() : this.conditions;
    }
    /**
     * @return phase represents the current phase of PersistentVolumeClaim.
     * 
     */
    public Optional<String> phase() {
        return Optional.ofNullable(this.phase);
    }
    /**
     * @return resizeStatus stores status of resize operation. ResizeStatus is not set by default but when expansion is complete resizeStatus is set to empty string by resize controller or kubelet. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.
     * 
     */
    public Optional<String> resizeStatus() {
        return Optional.ofNullable(this.resizeStatus);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PersistentVolumeClaimStatusPatch defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> accessModes;
        private @Nullable Map<String,String> allocatedResources;
        private @Nullable Map<String,String> capacity;
        private @Nullable List<PersistentVolumeClaimConditionPatch> conditions;
        private @Nullable String phase;
        private @Nullable String resizeStatus;
        public Builder() {}
        public Builder(PersistentVolumeClaimStatusPatch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessModes = defaults.accessModes;
    	      this.allocatedResources = defaults.allocatedResources;
    	      this.capacity = defaults.capacity;
    	      this.conditions = defaults.conditions;
    	      this.phase = defaults.phase;
    	      this.resizeStatus = defaults.resizeStatus;
        }

        @CustomType.Setter
        public Builder accessModes(@Nullable List<String> accessModes) {
            this.accessModes = accessModes;
            return this;
        }
        public Builder accessModes(String... accessModes) {
            return accessModes(List.of(accessModes));
        }
        @CustomType.Setter
        public Builder allocatedResources(@Nullable Map<String,String> allocatedResources) {
            this.allocatedResources = allocatedResources;
            return this;
        }
        @CustomType.Setter
        public Builder capacity(@Nullable Map<String,String> capacity) {
            this.capacity = capacity;
            return this;
        }
        @CustomType.Setter
        public Builder conditions(@Nullable List<PersistentVolumeClaimConditionPatch> conditions) {
            this.conditions = conditions;
            return this;
        }
        public Builder conditions(PersistentVolumeClaimConditionPatch... conditions) {
            return conditions(List.of(conditions));
        }
        @CustomType.Setter
        public Builder phase(@Nullable String phase) {
            this.phase = phase;
            return this;
        }
        @CustomType.Setter
        public Builder resizeStatus(@Nullable String resizeStatus) {
            this.resizeStatus = resizeStatus;
            return this;
        }
        public PersistentVolumeClaimStatusPatch build() {
            final var o = new PersistentVolumeClaimStatusPatch();
            o.accessModes = accessModes;
            o.allocatedResources = allocatedResources;
            o.capacity = capacity;
            o.conditions = conditions;
            o.phase = phase;
            o.resizeStatus = resizeStatus;
            return o;
        }
    }
}