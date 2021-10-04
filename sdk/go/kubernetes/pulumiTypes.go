// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Options for tuning the Kubernetes client used by a Provider.
type KubeClientSettings struct {
	// Maximum burst for throttle. Default value is 10.
	Burst *int `pulumi:"burst"`
	// Maximum queries per second (QPS) to the API server from this client. Default value is 5.
	Qps *float64 `pulumi:"qps"`
}

// KubeClientSettingsInput is an input type that accepts KubeClientSettingsArgs and KubeClientSettingsOutput values.
// You can construct a concrete instance of `KubeClientSettingsInput` via:
//
//          KubeClientSettingsArgs{...}
type KubeClientSettingsInput interface {
	pulumi.Input

	ToKubeClientSettingsOutput() KubeClientSettingsOutput
	ToKubeClientSettingsOutputWithContext(context.Context) KubeClientSettingsOutput
}

// Options for tuning the Kubernetes client used by a Provider.
type KubeClientSettingsArgs struct {
	// Maximum burst for throttle. Default value is 10.
	Burst pulumi.IntPtrInput `pulumi:"burst"`
	// Maximum queries per second (QPS) to the API server from this client. Default value is 5.
	Qps pulumi.Float64PtrInput `pulumi:"qps"`
}

func (KubeClientSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeClientSettings)(nil)).Elem()
}

func (i KubeClientSettingsArgs) ToKubeClientSettingsOutput() KubeClientSettingsOutput {
	return i.ToKubeClientSettingsOutputWithContext(context.Background())
}

func (i KubeClientSettingsArgs) ToKubeClientSettingsOutputWithContext(ctx context.Context) KubeClientSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeClientSettingsOutput)
}

func (i KubeClientSettingsArgs) ToKubeClientSettingsPtrOutput() KubeClientSettingsPtrOutput {
	return i.ToKubeClientSettingsPtrOutputWithContext(context.Background())
}

func (i KubeClientSettingsArgs) ToKubeClientSettingsPtrOutputWithContext(ctx context.Context) KubeClientSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeClientSettingsOutput).ToKubeClientSettingsPtrOutputWithContext(ctx)
}

// KubeClientSettingsPtrInput is an input type that accepts KubeClientSettingsArgs, KubeClientSettingsPtr and KubeClientSettingsPtrOutput values.
// You can construct a concrete instance of `KubeClientSettingsPtrInput` via:
//
//          KubeClientSettingsArgs{...}
//
//  or:
//
//          nil
type KubeClientSettingsPtrInput interface {
	pulumi.Input

	ToKubeClientSettingsPtrOutput() KubeClientSettingsPtrOutput
	ToKubeClientSettingsPtrOutputWithContext(context.Context) KubeClientSettingsPtrOutput
}

type kubeClientSettingsPtrType KubeClientSettingsArgs

func KubeClientSettingsPtr(v *KubeClientSettingsArgs) KubeClientSettingsPtrInput {
	return (*kubeClientSettingsPtrType)(v)
}

func (*kubeClientSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeClientSettings)(nil)).Elem()
}

func (i *kubeClientSettingsPtrType) ToKubeClientSettingsPtrOutput() KubeClientSettingsPtrOutput {
	return i.ToKubeClientSettingsPtrOutputWithContext(context.Background())
}

func (i *kubeClientSettingsPtrType) ToKubeClientSettingsPtrOutputWithContext(ctx context.Context) KubeClientSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeClientSettingsPtrOutput)
}

// Options for tuning the Kubernetes client used by a Provider.
type KubeClientSettingsOutput struct{ *pulumi.OutputState }

func (KubeClientSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeClientSettings)(nil)).Elem()
}

func (o KubeClientSettingsOutput) ToKubeClientSettingsOutput() KubeClientSettingsOutput {
	return o
}

func (o KubeClientSettingsOutput) ToKubeClientSettingsOutputWithContext(ctx context.Context) KubeClientSettingsOutput {
	return o
}

func (o KubeClientSettingsOutput) ToKubeClientSettingsPtrOutput() KubeClientSettingsPtrOutput {
	return o.ToKubeClientSettingsPtrOutputWithContext(context.Background())
}

func (o KubeClientSettingsOutput) ToKubeClientSettingsPtrOutputWithContext(ctx context.Context) KubeClientSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubeClientSettings) *KubeClientSettings {
		return &v
	}).(KubeClientSettingsPtrOutput)
}

// Maximum burst for throttle. Default value is 10.
func (o KubeClientSettingsOutput) Burst() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KubeClientSettings) *int { return v.Burst }).(pulumi.IntPtrOutput)
}

// Maximum queries per second (QPS) to the API server from this client. Default value is 5.
func (o KubeClientSettingsOutput) Qps() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KubeClientSettings) *float64 { return v.Qps }).(pulumi.Float64PtrOutput)
}

type KubeClientSettingsPtrOutput struct{ *pulumi.OutputState }

func (KubeClientSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeClientSettings)(nil)).Elem()
}

func (o KubeClientSettingsPtrOutput) ToKubeClientSettingsPtrOutput() KubeClientSettingsPtrOutput {
	return o
}

func (o KubeClientSettingsPtrOutput) ToKubeClientSettingsPtrOutputWithContext(ctx context.Context) KubeClientSettingsPtrOutput {
	return o
}

func (o KubeClientSettingsPtrOutput) Elem() KubeClientSettingsOutput {
	return o.ApplyT(func(v *KubeClientSettings) KubeClientSettings {
		if v != nil {
			return *v
		}
		var ret KubeClientSettings
		return ret
	}).(KubeClientSettingsOutput)
}

// Maximum burst for throttle. Default value is 10.
func (o KubeClientSettingsPtrOutput) Burst() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KubeClientSettings) *int {
		if v == nil {
			return nil
		}
		return v.Burst
	}).(pulumi.IntPtrOutput)
}

// Maximum queries per second (QPS) to the API server from this client. Default value is 5.
func (o KubeClientSettingsPtrOutput) Qps() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KubeClientSettings) *float64 {
		if v == nil {
			return nil
		}
		return v.Qps
	}).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(KubeClientSettingsOutput{})
	pulumi.RegisterOutputType(KubeClientSettingsPtrOutput{})
}
