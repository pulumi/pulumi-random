// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The resource `RandomPet` generates random pet names that are intended to be used as unique identifiers for other resources.
//
// This resource can be used in conjunction with resources that have the `createBeforeDestroy` lifecycle flag set, to avoid conflicts with unique names during the brief period where both the old and new resources exist concurrently.
type RandomPet struct {
	pulumi.CustomResourceState

	// Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
	// documentation](../index.html) for more information.
	Keepers pulumi.MapOutput `pulumi:"keepers"`
	// The length (in words) of the pet name.
	Length pulumi.IntPtrOutput `pulumi:"length"`
	// A string to prefix the name with.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
	// The character to separate words in the pet name.
	Separator pulumi.StringPtrOutput `pulumi:"separator"`
}

// NewRandomPet registers a new resource with the given unique name, arguments, and options.
func NewRandomPet(ctx *pulumi.Context,
	name string, args *RandomPetArgs, opts ...pulumi.ResourceOption) (*RandomPet, error) {
	if args == nil {
		args = &RandomPetArgs{}
	}

	var resource RandomPet
	err := ctx.RegisterResource("random:index/randomPet:RandomPet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandomPet gets an existing RandomPet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomPet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomPetState, opts ...pulumi.ResourceOption) (*RandomPet, error) {
	var resource RandomPet
	err := ctx.ReadResource("random:index/randomPet:RandomPet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RandomPet resources.
type randomPetState struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
	// documentation](../index.html) for more information.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The length (in words) of the pet name.
	Length *int `pulumi:"length"`
	// A string to prefix the name with.
	Prefix *string `pulumi:"prefix"`
	// The character to separate words in the pet name.
	Separator *string `pulumi:"separator"`
}

type RandomPetState struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
	// documentation](../index.html) for more information.
	Keepers pulumi.MapInput
	// The length (in words) of the pet name.
	Length pulumi.IntPtrInput
	// A string to prefix the name with.
	Prefix pulumi.StringPtrInput
	// The character to separate words in the pet name.
	Separator pulumi.StringPtrInput
}

func (RandomPetState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomPetState)(nil)).Elem()
}

type randomPetArgs struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
	// documentation](../index.html) for more information.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The length (in words) of the pet name.
	Length *int `pulumi:"length"`
	// A string to prefix the name with.
	Prefix *string `pulumi:"prefix"`
	// The character to separate words in the pet name.
	Separator *string `pulumi:"separator"`
}

// The set of arguments for constructing a RandomPet resource.
type RandomPetArgs struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
	// documentation](../index.html) for more information.
	Keepers pulumi.MapInput
	// The length (in words) of the pet name.
	Length pulumi.IntPtrInput
	// A string to prefix the name with.
	Prefix pulumi.StringPtrInput
	// The character to separate words in the pet name.
	Separator pulumi.StringPtrInput
}

func (RandomPetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomPetArgs)(nil)).Elem()
}

type RandomPetInput interface {
	pulumi.Input

	ToRandomPetOutput() RandomPetOutput
	ToRandomPetOutputWithContext(ctx context.Context) RandomPetOutput
}

func (*RandomPet) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomPet)(nil)).Elem()
}

func (i *RandomPet) ToRandomPetOutput() RandomPetOutput {
	return i.ToRandomPetOutputWithContext(context.Background())
}

func (i *RandomPet) ToRandomPetOutputWithContext(ctx context.Context) RandomPetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomPetOutput)
}

// RandomPetArrayInput is an input type that accepts RandomPetArray and RandomPetArrayOutput values.
// You can construct a concrete instance of `RandomPetArrayInput` via:
//
//          RandomPetArray{ RandomPetArgs{...} }
type RandomPetArrayInput interface {
	pulumi.Input

	ToRandomPetArrayOutput() RandomPetArrayOutput
	ToRandomPetArrayOutputWithContext(context.Context) RandomPetArrayOutput
}

type RandomPetArray []RandomPetInput

func (RandomPetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RandomPet)(nil)).Elem()
}

func (i RandomPetArray) ToRandomPetArrayOutput() RandomPetArrayOutput {
	return i.ToRandomPetArrayOutputWithContext(context.Background())
}

func (i RandomPetArray) ToRandomPetArrayOutputWithContext(ctx context.Context) RandomPetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomPetArrayOutput)
}

// RandomPetMapInput is an input type that accepts RandomPetMap and RandomPetMapOutput values.
// You can construct a concrete instance of `RandomPetMapInput` via:
//
//          RandomPetMap{ "key": RandomPetArgs{...} }
type RandomPetMapInput interface {
	pulumi.Input

	ToRandomPetMapOutput() RandomPetMapOutput
	ToRandomPetMapOutputWithContext(context.Context) RandomPetMapOutput
}

type RandomPetMap map[string]RandomPetInput

func (RandomPetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RandomPet)(nil)).Elem()
}

func (i RandomPetMap) ToRandomPetMapOutput() RandomPetMapOutput {
	return i.ToRandomPetMapOutputWithContext(context.Background())
}

func (i RandomPetMap) ToRandomPetMapOutputWithContext(ctx context.Context) RandomPetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomPetMapOutput)
}

type RandomPetOutput struct{ *pulumi.OutputState }

func (RandomPetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomPet)(nil)).Elem()
}

func (o RandomPetOutput) ToRandomPetOutput() RandomPetOutput {
	return o
}

func (o RandomPetOutput) ToRandomPetOutputWithContext(ctx context.Context) RandomPetOutput {
	return o
}

type RandomPetArrayOutput struct{ *pulumi.OutputState }

func (RandomPetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RandomPet)(nil)).Elem()
}

func (o RandomPetArrayOutput) ToRandomPetArrayOutput() RandomPetArrayOutput {
	return o
}

func (o RandomPetArrayOutput) ToRandomPetArrayOutputWithContext(ctx context.Context) RandomPetArrayOutput {
	return o
}

func (o RandomPetArrayOutput) Index(i pulumi.IntInput) RandomPetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RandomPet {
		return vs[0].([]*RandomPet)[vs[1].(int)]
	}).(RandomPetOutput)
}

type RandomPetMapOutput struct{ *pulumi.OutputState }

func (RandomPetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RandomPet)(nil)).Elem()
}

func (o RandomPetMapOutput) ToRandomPetMapOutput() RandomPetMapOutput {
	return o
}

func (o RandomPetMapOutput) ToRandomPetMapOutputWithContext(ctx context.Context) RandomPetMapOutput {
	return o
}

func (o RandomPetMapOutput) MapIndex(k pulumi.StringInput) RandomPetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RandomPet {
		return vs[0].(map[string]*RandomPet)[vs[1].(string)]
	}).(RandomPetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RandomPetInput)(nil)).Elem(), &RandomPet{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomPetArrayInput)(nil)).Elem(), RandomPetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomPetMapInput)(nil)).Elem(), RandomPetMap{})
	pulumi.RegisterOutputType(RandomPetOutput{})
	pulumi.RegisterOutputType(RandomPetArrayOutput{})
	pulumi.RegisterOutputType(RandomPetMapOutput{})
}
