// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The resource `RandomUuid` generates random uuid string that is intended to be used as unique identifiers for other resources.
//
// This resource uses [hashicorp/go-uuid](https://github.com/hashicorp/go-uuid) to generate a UUID-formatted string for use with services needed a unique string identifier.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := random.NewRandomUuid(ctx, "testRandomUuid", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = core.NewResourceGroup(ctx, "testResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("Central US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// # Random UUID's can be imported. This can be used to replace a config # value with a value interpolated from the random provider without # experiencing diffs.
//
// ```sh
//  $ pulumi import random:index/randomUuid:RandomUuid main aabbccdd-eeff-0011-2233-445566778899
// ```
type RandomUuid struct {
	pulumi.CustomResourceState

	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers pulumi.MapOutput `pulumi:"keepers"`
	// The generated uuid presented in string format.
	Result pulumi.StringOutput `pulumi:"result"`
}

// NewRandomUuid registers a new resource with the given unique name, arguments, and options.
func NewRandomUuid(ctx *pulumi.Context,
	name string, args *RandomUuidArgs, opts ...pulumi.ResourceOption) (*RandomUuid, error) {
	if args == nil {
		args = &RandomUuidArgs{}
	}

	var resource RandomUuid
	err := ctx.RegisterResource("random:index/randomUuid:RandomUuid", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandomUuid gets an existing RandomUuid resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomUuid(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomUuidState, opts ...pulumi.ResourceOption) (*RandomUuid, error) {
	var resource RandomUuid
	err := ctx.ReadResource("random:index/randomUuid:RandomUuid", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RandomUuid resources.
type randomUuidState struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The generated uuid presented in string format.
	Result *string `pulumi:"result"`
}

type RandomUuidState struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers pulumi.MapInput
	// The generated uuid presented in string format.
	Result pulumi.StringPtrInput
}

func (RandomUuidState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomUuidState)(nil)).Elem()
}

type randomUuidArgs struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers map[string]interface{} `pulumi:"keepers"`
}

// The set of arguments for constructing a RandomUuid resource.
type RandomUuidArgs struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers pulumi.MapInput
}

func (RandomUuidArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomUuidArgs)(nil)).Elem()
}

type RandomUuidInput interface {
	pulumi.Input

	ToRandomUuidOutput() RandomUuidOutput
	ToRandomUuidOutputWithContext(ctx context.Context) RandomUuidOutput
}

func (*RandomUuid) ElementType() reflect.Type {
	return reflect.TypeOf((*RandomUuid)(nil))
}

func (i *RandomUuid) ToRandomUuidOutput() RandomUuidOutput {
	return i.ToRandomUuidOutputWithContext(context.Background())
}

func (i *RandomUuid) ToRandomUuidOutputWithContext(ctx context.Context) RandomUuidOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomUuidOutput)
}

func (i *RandomUuid) ToRandomUuidPtrOutput() RandomUuidPtrOutput {
	return i.ToRandomUuidPtrOutputWithContext(context.Background())
}

func (i *RandomUuid) ToRandomUuidPtrOutputWithContext(ctx context.Context) RandomUuidPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomUuidPtrOutput)
}

type RandomUuidPtrInput interface {
	pulumi.Input

	ToRandomUuidPtrOutput() RandomUuidPtrOutput
	ToRandomUuidPtrOutputWithContext(ctx context.Context) RandomUuidPtrOutput
}

type randomUuidPtrType RandomUuidArgs

func (*randomUuidPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomUuid)(nil))
}

func (i *randomUuidPtrType) ToRandomUuidPtrOutput() RandomUuidPtrOutput {
	return i.ToRandomUuidPtrOutputWithContext(context.Background())
}

func (i *randomUuidPtrType) ToRandomUuidPtrOutputWithContext(ctx context.Context) RandomUuidPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomUuidPtrOutput)
}

// RandomUuidArrayInput is an input type that accepts RandomUuidArray and RandomUuidArrayOutput values.
// You can construct a concrete instance of `RandomUuidArrayInput` via:
//
//          RandomUuidArray{ RandomUuidArgs{...} }
type RandomUuidArrayInput interface {
	pulumi.Input

	ToRandomUuidArrayOutput() RandomUuidArrayOutput
	ToRandomUuidArrayOutputWithContext(context.Context) RandomUuidArrayOutput
}

type RandomUuidArray []RandomUuidInput

func (RandomUuidArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RandomUuid)(nil))
}

func (i RandomUuidArray) ToRandomUuidArrayOutput() RandomUuidArrayOutput {
	return i.ToRandomUuidArrayOutputWithContext(context.Background())
}

func (i RandomUuidArray) ToRandomUuidArrayOutputWithContext(ctx context.Context) RandomUuidArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomUuidArrayOutput)
}

// RandomUuidMapInput is an input type that accepts RandomUuidMap and RandomUuidMapOutput values.
// You can construct a concrete instance of `RandomUuidMapInput` via:
//
//          RandomUuidMap{ "key": RandomUuidArgs{...} }
type RandomUuidMapInput interface {
	pulumi.Input

	ToRandomUuidMapOutput() RandomUuidMapOutput
	ToRandomUuidMapOutputWithContext(context.Context) RandomUuidMapOutput
}

type RandomUuidMap map[string]RandomUuidInput

func (RandomUuidMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RandomUuid)(nil))
}

func (i RandomUuidMap) ToRandomUuidMapOutput() RandomUuidMapOutput {
	return i.ToRandomUuidMapOutputWithContext(context.Background())
}

func (i RandomUuidMap) ToRandomUuidMapOutputWithContext(ctx context.Context) RandomUuidMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomUuidMapOutput)
}

type RandomUuidOutput struct {
	*pulumi.OutputState
}

func (RandomUuidOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RandomUuid)(nil))
}

func (o RandomUuidOutput) ToRandomUuidOutput() RandomUuidOutput {
	return o
}

func (o RandomUuidOutput) ToRandomUuidOutputWithContext(ctx context.Context) RandomUuidOutput {
	return o
}

func (o RandomUuidOutput) ToRandomUuidPtrOutput() RandomUuidPtrOutput {
	return o.ToRandomUuidPtrOutputWithContext(context.Background())
}

func (o RandomUuidOutput) ToRandomUuidPtrOutputWithContext(ctx context.Context) RandomUuidPtrOutput {
	return o.ApplyT(func(v RandomUuid) *RandomUuid {
		return &v
	}).(RandomUuidPtrOutput)
}

type RandomUuidPtrOutput struct {
	*pulumi.OutputState
}

func (RandomUuidPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomUuid)(nil))
}

func (o RandomUuidPtrOutput) ToRandomUuidPtrOutput() RandomUuidPtrOutput {
	return o
}

func (o RandomUuidPtrOutput) ToRandomUuidPtrOutputWithContext(ctx context.Context) RandomUuidPtrOutput {
	return o
}

type RandomUuidArrayOutput struct{ *pulumi.OutputState }

func (RandomUuidArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RandomUuid)(nil))
}

func (o RandomUuidArrayOutput) ToRandomUuidArrayOutput() RandomUuidArrayOutput {
	return o
}

func (o RandomUuidArrayOutput) ToRandomUuidArrayOutputWithContext(ctx context.Context) RandomUuidArrayOutput {
	return o
}

func (o RandomUuidArrayOutput) Index(i pulumi.IntInput) RandomUuidOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RandomUuid {
		return vs[0].([]RandomUuid)[vs[1].(int)]
	}).(RandomUuidOutput)
}

type RandomUuidMapOutput struct{ *pulumi.OutputState }

func (RandomUuidMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RandomUuid)(nil))
}

func (o RandomUuidMapOutput) ToRandomUuidMapOutput() RandomUuidMapOutput {
	return o
}

func (o RandomUuidMapOutput) ToRandomUuidMapOutputWithContext(ctx context.Context) RandomUuidMapOutput {
	return o
}

func (o RandomUuidMapOutput) MapIndex(k pulumi.StringInput) RandomUuidOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RandomUuid {
		return vs[0].(map[string]RandomUuid)[vs[1].(string)]
	}).(RandomUuidOutput)
}

func init() {
	pulumi.RegisterOutputType(RandomUuidOutput{})
	pulumi.RegisterOutputType(RandomUuidPtrOutput{})
	pulumi.RegisterOutputType(RandomUuidArrayOutput{})
	pulumi.RegisterOutputType(RandomUuidMapOutput{})
}
