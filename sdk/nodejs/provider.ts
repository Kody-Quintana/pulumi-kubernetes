// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

/**
 * The provider type for the kubernetes package.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'kubernetes';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }


    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            inputs["cluster"] = args ? args.cluster : undefined;
            inputs["context"] = args ? args.context : undefined;
            inputs["enableDryRun"] = pulumi.output((args ? args.enableDryRun : undefined) ?? <any>utilities.getEnvBoolean("PULUMI_K8S_ENABLE_DRY_RUN")).apply(JSON.stringify);
            inputs["helmDriver"] = (args ? args.helmDriver : undefined) ?? utilities.getEnv("PULUMI_K8S_HELM_DRIVER");
            inputs["helmPluginsPath"] = (args ? args.helmPluginsPath : undefined) ?? utilities.getEnv("PULUMI_K8S_HELM_PLUGINS_PATH");
            inputs["helmRegistryConfigPath"] = (args ? args.helmRegistryConfigPath : undefined) ?? utilities.getEnv("PULUMI_K8S_HELM_REGISTRY_CONFIG_PATH");
            inputs["helmRepositoryCache"] = (args ? args.helmRepositoryCache : undefined) ?? utilities.getEnv("PULUMI_K8s_HELM_REPOSITORY_CACHE");
            inputs["helmRepositoryConfigPath"] = (args ? args.helmRepositoryConfigPath : undefined) ?? utilities.getEnv("PULUMI_K8S_HELM_REPOSITORY_CONFIG_PATH");
            inputs["kubeClientSettings"] = pulumi.output(args ? args.kubeClientSettings : undefined).apply(JSON.stringify);
            inputs["kubeconfig"] = (args ? args.kubeconfig : undefined) ?? utilities.getEnv("KUBECONFIG");
            inputs["namespace"] = args ? args.namespace : undefined;
            inputs["renderYamlToDirectory"] = args ? args.renderYamlToDirectory : undefined;
            inputs["suppressDeprecationWarnings"] = pulumi.output((args ? args.suppressDeprecationWarnings : undefined) ?? <any>utilities.getEnvBoolean("PULUMI_K8S_SUPPRESS_DEPRECATION_WARNINGS")).apply(JSON.stringify);
            inputs["suppressHelmHookWarnings"] = pulumi.output((args ? args.suppressHelmHookWarnings : undefined) ?? <any>utilities.getEnvBoolean("PULUMI_K8S_SUPPRESS_HELM_HOOK_WARNINGS")).apply(JSON.stringify);
            inputs["suppressHelmReleaseBetaWarning"] = pulumi.output((args ? args.suppressHelmReleaseBetaWarning : undefined) ?? <any>utilities.getEnvBoolean("PULUMI_K8S_SUPPRESS_HELM_RELEASE_BETA_WARNING")).apply(JSON.stringify);
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Provider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * If present, the name of the kubeconfig cluster to use.
     */
    cluster?: pulumi.Input<string>;
    /**
     * If present, the name of the kubeconfig context to use.
     */
    context?: pulumi.Input<string>;
    /**
     * BETA FEATURE - If present and set to true, enable server-side diff calculations.
     * This feature is in developer preview, and is disabled by default.
     */
    enableDryRun?: pulumi.Input<boolean>;
    /**
     * BETA FEATURE - Used for supporting Helm Release resource (Beta). The backend storage driver for Helm. Values are: configmap, secret, memory, sql.
     */
    helmDriver?: pulumi.Input<string>;
    /**
     * BETA FEATURE - Used for supporting Helm Release resource (Beta). The path to the helm plugins directory.
     */
    helmPluginsPath?: pulumi.Input<string>;
    /**
     * BETA FEATURE - Used for supporting Helm Release resource (Beta). The path to the registry config file.
     */
    helmRegistryConfigPath?: pulumi.Input<string>;
    /**
     * BETA FEATURE - Used for supporting Helm Release resource (Beta). The path to the file containing cached repository indexes.
     */
    helmRepositoryCache?: pulumi.Input<string>;
    /**
     * BETA FEATURE - Used for supporting Helm Release resource (Beta). The path to the file containing repository names and URLs.
     */
    helmRepositoryConfigPath?: pulumi.Input<string>;
    /**
     * Options for tuning the Kubernetes client used by a Provider.
     */
    kubeClientSettings?: pulumi.Input<inputs.KubeClientSettings>;
    /**
     * The contents of a kubeconfig file or the path to a kubeconfig file.
     */
    kubeconfig?: pulumi.Input<string>;
    /**
     * If present, the default namespace to use. This flag is ignored for cluster-scoped resources.
     *
     * A namespace can be specified in multiple places, and the precedence is as follows:
     * 1. `.metadata.namespace` set on the resource.
     * 2. This `namespace` parameter.
     * 3. `namespace` set for the active context in the kubeconfig.
     */
    namespace?: pulumi.Input<string>;
    /**
     * BETA FEATURE - If present, render resource manifests to this directory. In this mode, resources will not
     * be created on a Kubernetes cluster, but the rendered manifests will be kept in sync with changes
     * to the Pulumi program. This feature is in developer preview, and is disabled by default.
     *
     * Note that some computed Outputs such as status fields will not be populated
     * since the resources are not created on a Kubernetes cluster. These Output values will remain undefined,
     * and may result in an error if they are referenced by other resources. Also note that any secret values
     * used in these resources will be rendered in plaintext to the resulting YAML.
     */
    renderYamlToDirectory?: pulumi.Input<string>;
    /**
     * If present and set to true, suppress apiVersion deprecation warnings from the CLI.
     */
    suppressDeprecationWarnings?: pulumi.Input<boolean>;
    /**
     * If present and set to true, suppress unsupported Helm hook warnings from the CLI.
     */
    suppressHelmHookWarnings?: pulumi.Input<boolean>;
    /**
     * While Helm Release provider is in beta, by default 'pulumi up' will log a warning if the resource is used. If present and set to "true", this warning is omitted.
     */
    suppressHelmReleaseBetaWarning?: pulumi.Input<boolean>;
}
