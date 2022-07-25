// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.core.v1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


/**
 * NodeAddress contains information for the node&#39;s address.
 * 
 */
public final class NodeAddressArgs extends com.pulumi.resources.ResourceArgs {

    public static final NodeAddressArgs Empty = new NodeAddressArgs();

    /**
     * The node address.
     * 
     */
    @Import(name="address", required=true)
    private Output<String> address;

    /**
     * @return The node address.
     * 
     */
    public Output<String> address() {
        return this.address;
    }

    /**
     * Node address type, one of Hostname, ExternalIP or InternalIP.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Node address type, one of Hostname, ExternalIP or InternalIP.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private NodeAddressArgs() {}

    private NodeAddressArgs(NodeAddressArgs $) {
        this.address = $.address;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NodeAddressArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NodeAddressArgs $;

        public Builder() {
            $ = new NodeAddressArgs();
        }

        public Builder(NodeAddressArgs defaults) {
            $ = new NodeAddressArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address The node address.
         * 
         * @return builder
         * 
         */
        public Builder address(Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address The node address.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param type Node address type, one of Hostname, ExternalIP or InternalIP.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Node address type, one of Hostname, ExternalIP or InternalIP.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public NodeAddressArgs build() {
            $.address = Objects.requireNonNull($.address, "expected parameter 'address' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}