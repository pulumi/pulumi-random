// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The resource `RandomBytes` generates random bytes that are intended to be used as a secret, or key. Use this in preference to `RandomId` when the output is considered sensitive, and should not be displayed in the CLI.
//
// This resource *does* use a cryptographic random number generator.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azurerm/sdk/v1/go/azurerm"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			jwtSecret, err := random.NewRandomBytes(ctx, "jwt_secret", &random.RandomBytesArgs{
//				Length: pulumi.Int(64),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = index.NewKeyVaultSecret(ctx, "jwt_secret", &index.KeyVaultSecretArgs{
//				KeyVaultId: "some-azure-key-vault-id",
//				Name:       "JwtSecret",
//				Value:      jwtSecret.Base64,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Random bytes can be imported by specifying the value as base64 string.
//
// ```sh
// $ pulumi import random:index/randomBytes:RandomBytes basic "8/fu3q+2DcgSJ19i0jZ5Cw=="
// ```
type RandomBytes struct {
	pulumi.CustomResourceState

	// The generated bytes presented in base64 string format.
	Base64 pulumi.StringOutput `pulumi:"base64"`
	// The generated bytes presented in lowercase hexadecimal string format. The length of the encoded string is exactly twice the `length` parameter.
	Hex pulumi.StringOutput `pulumi:"hex"`
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers pulumi.StringMapOutput `pulumi:"keepers"`
	// The number of bytes requested. The minimum value for length is 1.
	Length pulumi.IntOutput `pulumi:"length"`
}

// NewRandomBytes registers a new resource with the given unique name, arguments, and options.
func NewRandomBytes(ctx *pulumi.Context,
	name string, args *RandomBytesArgs, opts ...pulumi.ResourceOption) (*RandomBytes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Length == nil {
		return nil, errors.New("invalid value for required argument 'Length'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"base64",
		"hex",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RandomBytes
	err := ctx.RegisterResource("random:index/randomBytes:RandomBytes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandomBytes gets an existing RandomBytes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomBytes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomBytesState, opts ...pulumi.ResourceOption) (*RandomBytes, error) {
	var resource RandomBytes
	err := ctx.ReadResource("random:index/randomBytes:RandomBytes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RandomBytes resources.
type randomBytesState struct {
	// The generated bytes presented in base64 string format.
	Base64 *string `pulumi:"base64"`
	// The generated bytes presented in lowercase hexadecimal string format. The length of the encoded string is exactly twice the `length` parameter.
	Hex *string `pulumi:"hex"`
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers map[string]string `pulumi:"keepers"`
	// The number of bytes requested. The minimum value for length is 1.
	Length *int `pulumi:"length"`
}

type RandomBytesState struct {
	// The generated bytes presented in base64 string format.
	Base64 pulumi.StringPtrInput
	// The generated bytes presented in lowercase hexadecimal string format. The length of the encoded string is exactly twice the `length` parameter.
	Hex pulumi.StringPtrInput
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers pulumi.StringMapInput
	// The number of bytes requested. The minimum value for length is 1.
	Length pulumi.IntPtrInput
}

func (RandomBytesState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomBytesState)(nil)).Elem()
}

type randomBytesArgs struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers map[string]string `pulumi:"keepers"`
	// The number of bytes requested. The minimum value for length is 1.
	Length int `pulumi:"length"`
}

// The set of arguments for constructing a RandomBytes resource.
type RandomBytesArgs struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers pulumi.StringMapInput
	// The number of bytes requested. The minimum value for length is 1.
	Length pulumi.IntInput
}

func (RandomBytesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomBytesArgs)(nil)).Elem()
}

type RandomBytesInput interface {
	pulumi.Input

	ToRandomBytesOutput() RandomBytesOutput
	ToRandomBytesOutputWithContext(ctx context.Context) RandomBytesOutput
}

func (*RandomBytes) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomBytes)(nil)).Elem()
}

func (i *RandomBytes) ToRandomBytesOutput() RandomBytesOutput {
	return i.ToRandomBytesOutputWithContext(context.Background())
}

func (i *RandomBytes) ToRandomBytesOutputWithContext(ctx context.Context) RandomBytesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomBytesOutput)
}

// RandomBytesArrayInput is an input type that accepts RandomBytesArray and RandomBytesArrayOutput values.
// You can construct a concrete instance of `RandomBytesArrayInput` via:
//
//	RandomBytesArray{ RandomBytesArgs{...} }
type RandomBytesArrayInput interface {
	pulumi.Input

	ToRandomBytesArrayOutput() RandomBytesArrayOutput
	ToRandomBytesArrayOutputWithContext(context.Context) RandomBytesArrayOutput
}

type RandomBytesArray []RandomBytesInput

func (RandomBytesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RandomBytes)(nil)).Elem()
}

func (i RandomBytesArray) ToRandomBytesArrayOutput() RandomBytesArrayOutput {
	return i.ToRandomBytesArrayOutputWithContext(context.Background())
}

func (i RandomBytesArray) ToRandomBytesArrayOutputWithContext(ctx context.Context) RandomBytesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomBytesArrayOutput)
}

// RandomBytesMapInput is an input type that accepts RandomBytesMap and RandomBytesMapOutput values.
// You can construct a concrete instance of `RandomBytesMapInput` via:
//
//	RandomBytesMap{ "key": RandomBytesArgs{...} }
type RandomBytesMapInput interface {
	pulumi.Input

	ToRandomBytesMapOutput() RandomBytesMapOutput
	ToRandomBytesMapOutputWithContext(context.Context) RandomBytesMapOutput
}

type RandomBytesMap map[string]RandomBytesInput

func (RandomBytesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RandomBytes)(nil)).Elem()
}

func (i RandomBytesMap) ToRandomBytesMapOutput() RandomBytesMapOutput {
	return i.ToRandomBytesMapOutputWithContext(context.Background())
}

func (i RandomBytesMap) ToRandomBytesMapOutputWithContext(ctx context.Context) RandomBytesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomBytesMapOutput)
}

type RandomBytesOutput struct{ *pulumi.OutputState }

func (RandomBytesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomBytes)(nil)).Elem()
}

func (o RandomBytesOutput) ToRandomBytesOutput() RandomBytesOutput {
	return o
}

func (o RandomBytesOutput) ToRandomBytesOutputWithContext(ctx context.Context) RandomBytesOutput {
	return o
}

// The generated bytes presented in base64 string format.
func (o RandomBytesOutput) Base64() pulumi.StringOutput {
	return o.ApplyT(func(v *RandomBytes) pulumi.StringOutput { return v.Base64 }).(pulumi.StringOutput)
}

// The generated bytes presented in lowercase hexadecimal string format. The length of the encoded string is exactly twice the `length` parameter.
func (o RandomBytesOutput) Hex() pulumi.StringOutput {
	return o.ApplyT(func(v *RandomBytes) pulumi.StringOutput { return v.Hex }).(pulumi.StringOutput)
}

// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
func (o RandomBytesOutput) Keepers() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RandomBytes) pulumi.StringMapOutput { return v.Keepers }).(pulumi.StringMapOutput)
}

// The number of bytes requested. The minimum value for length is 1.
func (o RandomBytesOutput) Length() pulumi.IntOutput {
	return o.ApplyT(func(v *RandomBytes) pulumi.IntOutput { return v.Length }).(pulumi.IntOutput)
}

type RandomBytesArrayOutput struct{ *pulumi.OutputState }

func (RandomBytesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RandomBytes)(nil)).Elem()
}

func (o RandomBytesArrayOutput) ToRandomBytesArrayOutput() RandomBytesArrayOutput {
	return o
}

func (o RandomBytesArrayOutput) ToRandomBytesArrayOutputWithContext(ctx context.Context) RandomBytesArrayOutput {
	return o
}

func (o RandomBytesArrayOutput) Index(i pulumi.IntInput) RandomBytesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RandomBytes {
		return vs[0].([]*RandomBytes)[vs[1].(int)]
	}).(RandomBytesOutput)
}

type RandomBytesMapOutput struct{ *pulumi.OutputState }

func (RandomBytesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RandomBytes)(nil)).Elem()
}

func (o RandomBytesMapOutput) ToRandomBytesMapOutput() RandomBytesMapOutput {
	return o
}

func (o RandomBytesMapOutput) ToRandomBytesMapOutputWithContext(ctx context.Context) RandomBytesMapOutput {
	return o
}

func (o RandomBytesMapOutput) MapIndex(k pulumi.StringInput) RandomBytesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RandomBytes {
		return vs[0].(map[string]*RandomBytes)[vs[1].(string)]
	}).(RandomBytesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RandomBytesInput)(nil)).Elem(), &RandomBytes{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomBytesArrayInput)(nil)).Elem(), RandomBytesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomBytesMapInput)(nil)).Elem(), RandomBytesMap{})
	pulumi.RegisterOutputType(RandomBytesOutput{})
	pulumi.RegisterOutputType(RandomBytesArrayOutput{})
	pulumi.RegisterOutputType(RandomBytesMapOutput{})
}
