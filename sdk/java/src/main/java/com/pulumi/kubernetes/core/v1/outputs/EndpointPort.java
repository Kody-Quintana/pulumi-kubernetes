// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetes.core.v1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class EndpointPort {
    /**
     * @return The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.
     * 
     */
    private @Nullable String appProtocol;
    /**
     * @return The name of this port.  This must match the &#39;name&#39; field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.
     * 
     */
    private @Nullable String name;
    /**
     * @return The port number of the endpoint.
     * 
     */
    private Integer port;
    /**
     * @return The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
     * 
     */
    private @Nullable String protocol;

    private EndpointPort() {}
    /**
     * @return The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.
     * 
     */
    public Optional<String> appProtocol() {
        return Optional.ofNullable(this.appProtocol);
    }
    /**
     * @return The name of this port.  This must match the &#39;name&#39; field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The port number of the endpoint.
     * 
     */
    public Integer port() {
        return this.port;
    }
    /**
     * @return The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
     * 
     */
    public Optional<String> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(EndpointPort defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String appProtocol;
        private @Nullable String name;
        private Integer port;
        private @Nullable String protocol;
        public Builder() {}
        public Builder(EndpointPort defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.appProtocol = defaults.appProtocol;
    	      this.name = defaults.name;
    	      this.port = defaults.port;
    	      this.protocol = defaults.protocol;
        }

        @CustomType.Setter
        public Builder appProtocol(@Nullable String appProtocol) {
            this.appProtocol = appProtocol;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        @CustomType.Setter
        public Builder protocol(@Nullable String protocol) {
            this.protocol = protocol;
            return this;
        }
        public EndpointPort build() {
            final var o = new EndpointPort();
            o.appProtocol = appProtocol;
            o.name = name;
            o.port = port;
            o.protocol = protocol;
            return o;
        }
    }
}