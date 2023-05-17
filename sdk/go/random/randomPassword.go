// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Identical to random_string.
//
// This resource *does* use a cryptographic random number generator.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/rds"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			password, err := random.NewRandomPassword(ctx, "password", &random.RandomPasswordArgs{
//				Length:          pulumi.Int(16),
//				Special:         pulumi.Bool(true),
//				OverrideSpecial: pulumi.String("!#$%&*()-_=+[]{}<>:?"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rds.NewInstance(ctx, "example", &rds.InstanceArgs{
//				InstanceClass:    pulumi.String("db.t3.micro"),
//				AllocatedStorage: pulumi.Int(64),
//				Engine:           pulumi.String("mysql"),
//				Username:         pulumi.String("someone"),
//				Password:         password.Result,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ### Avoiding Replacement
//
// ```sh
//
//	$ pulumi import random:index/randomPassword:RandomPassword If the resource were imported using `random_password.password securepassword`,
//
// ```
//
//	replacement could be avoided by using1. Attribute values that match the imported ID and defaults:
//
//	terraform
//
//	resource "random_password" "password" {
//
//	length = 14
//
//	lower
//
// = true
//
//	} 2. Attribute values that match the imported ID and omit the attributes with defaults:
//
//	terraform
//
//	resource "random_password" "password" {
//
//	length = 14
//
//	} 3. `ignore_changes` specifying the attributes to ignore:
//
//	terraform
//
//	resource "random_password" "password" {
//
//	length = 16
//
//	lower
//
// = false
//
//	lifecycle {
//
//	ignore_changes = [
//
//	length,
//
//	lower,
//
//	]
//
//	}
//
//	}
//
//	**NOTE** `ignore_changes` is only required until the resource is recreated after import,
//
//	after which it will use the configuration values specified.
type RandomPassword struct {
	pulumi.CustomResourceState

	// A bcrypt hash of the generated random string.
	BcryptHash pulumi.StringOutput `pulumi:"bcryptHash"`
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers pulumi.StringMapOutput `pulumi:"keepers"`
	// The length of the string desired. The minimum value for length is 1 and, length must also be >= (`minUpper` + `minLower` + `minNumeric` + `minSpecial`).
	Length pulumi.IntOutput `pulumi:"length"`
	// Include lowercase alphabet characters in the result. Default value is `true`.
	Lower pulumi.BoolOutput `pulumi:"lower"`
	// Minimum number of lowercase alphabet characters in the result. Default value is `0`.
	MinLower pulumi.IntOutput `pulumi:"minLower"`
	// Minimum number of numeric characters in the result. Default value is `0`.
	MinNumeric pulumi.IntOutput `pulumi:"minNumeric"`
	// Minimum number of special characters in the result. Default value is `0`.
	MinSpecial pulumi.IntOutput `pulumi:"minSpecial"`
	// Minimum number of uppercase alphabet characters in the result. Default value is `0`.
	MinUpper pulumi.IntOutput `pulumi:"minUpper"`
	// Include numeric characters in the result. Default value is `true`. **NOTE**: This is deprecated, use `numeric` instead.
	//
	// Deprecated: **NOTE**: This is deprecated, use `numeric` instead.
	Number pulumi.BoolOutput `pulumi:"number"`
	// Include numeric characters in the result. Default value is `true`.
	Numeric pulumi.BoolOutput `pulumi:"numeric"`
	// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
	OverrideSpecial pulumi.StringPtrOutput `pulumi:"overrideSpecial"`
	// The generated random string.
	Result pulumi.StringOutput `pulumi:"result"`
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
	Special pulumi.BoolOutput `pulumi:"special"`
	// Include uppercase alphabet characters in the result. Default value is `true`.
	Upper pulumi.BoolOutput `pulumi:"upper"`
}

// NewRandomPassword registers a new resource with the given unique name, arguments, and options.
func NewRandomPassword(ctx *pulumi.Context,
	name string, args *RandomPasswordArgs, opts ...pulumi.ResourceOption) (*RandomPassword, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Length == nil {
		return nil, errors.New("invalid value for required argument 'Length'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"bcryptHash",
		"result",
	})
	opts = append(opts, secrets)
	var resource RandomPassword
	err := ctx.RegisterResource("random:index/randomPassword:RandomPassword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandomPassword gets an existing RandomPassword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomPassword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomPasswordState, opts ...pulumi.ResourceOption) (*RandomPassword, error) {
	var resource RandomPassword
	err := ctx.ReadResource("random:index/randomPassword:RandomPassword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RandomPassword resources.
type randomPasswordState struct {
	// A bcrypt hash of the generated random string.
	BcryptHash *string `pulumi:"bcryptHash"`
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers map[string]string `pulumi:"keepers"`
	// The length of the string desired. The minimum value for length is 1 and, length must also be >= (`minUpper` + `minLower` + `minNumeric` + `minSpecial`).
	Length *int `pulumi:"length"`
	// Include lowercase alphabet characters in the result. Default value is `true`.
	Lower *bool `pulumi:"lower"`
	// Minimum number of lowercase alphabet characters in the result. Default value is `0`.
	MinLower *int `pulumi:"minLower"`
	// Minimum number of numeric characters in the result. Default value is `0`.
	MinNumeric *int `pulumi:"minNumeric"`
	// Minimum number of special characters in the result. Default value is `0`.
	MinSpecial *int `pulumi:"minSpecial"`
	// Minimum number of uppercase alphabet characters in the result. Default value is `0`.
	MinUpper *int `pulumi:"minUpper"`
	// Include numeric characters in the result. Default value is `true`. **NOTE**: This is deprecated, use `numeric` instead.
	//
	// Deprecated: **NOTE**: This is deprecated, use `numeric` instead.
	Number *bool `pulumi:"number"`
	// Include numeric characters in the result. Default value is `true`.
	Numeric *bool `pulumi:"numeric"`
	// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
	OverrideSpecial *string `pulumi:"overrideSpecial"`
	// The generated random string.
	Result *string `pulumi:"result"`
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
	Special *bool `pulumi:"special"`
	// Include uppercase alphabet characters in the result. Default value is `true`.
	Upper *bool `pulumi:"upper"`
}

type RandomPasswordState struct {
	// A bcrypt hash of the generated random string.
	BcryptHash pulumi.StringPtrInput
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers pulumi.StringMapInput
	// The length of the string desired. The minimum value for length is 1 and, length must also be >= (`minUpper` + `minLower` + `minNumeric` + `minSpecial`).
	Length pulumi.IntPtrInput
	// Include lowercase alphabet characters in the result. Default value is `true`.
	Lower pulumi.BoolPtrInput
	// Minimum number of lowercase alphabet characters in the result. Default value is `0`.
	MinLower pulumi.IntPtrInput
	// Minimum number of numeric characters in the result. Default value is `0`.
	MinNumeric pulumi.IntPtrInput
	// Minimum number of special characters in the result. Default value is `0`.
	MinSpecial pulumi.IntPtrInput
	// Minimum number of uppercase alphabet characters in the result. Default value is `0`.
	MinUpper pulumi.IntPtrInput
	// Include numeric characters in the result. Default value is `true`. **NOTE**: This is deprecated, use `numeric` instead.
	//
	// Deprecated: **NOTE**: This is deprecated, use `numeric` instead.
	Number pulumi.BoolPtrInput
	// Include numeric characters in the result. Default value is `true`.
	Numeric pulumi.BoolPtrInput
	// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
	OverrideSpecial pulumi.StringPtrInput
	// The generated random string.
	Result pulumi.StringPtrInput
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
	Special pulumi.BoolPtrInput
	// Include uppercase alphabet characters in the result. Default value is `true`.
	Upper pulumi.BoolPtrInput
}

func (RandomPasswordState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomPasswordState)(nil)).Elem()
}

type randomPasswordArgs struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers map[string]string `pulumi:"keepers"`
	// The length of the string desired. The minimum value for length is 1 and, length must also be >= (`minUpper` + `minLower` + `minNumeric` + `minSpecial`).
	Length int `pulumi:"length"`
	// Include lowercase alphabet characters in the result. Default value is `true`.
	Lower *bool `pulumi:"lower"`
	// Minimum number of lowercase alphabet characters in the result. Default value is `0`.
	MinLower *int `pulumi:"minLower"`
	// Minimum number of numeric characters in the result. Default value is `0`.
	MinNumeric *int `pulumi:"minNumeric"`
	// Minimum number of special characters in the result. Default value is `0`.
	MinSpecial *int `pulumi:"minSpecial"`
	// Minimum number of uppercase alphabet characters in the result. Default value is `0`.
	MinUpper *int `pulumi:"minUpper"`
	// Include numeric characters in the result. Default value is `true`. **NOTE**: This is deprecated, use `numeric` instead.
	//
	// Deprecated: **NOTE**: This is deprecated, use `numeric` instead.
	Number *bool `pulumi:"number"`
	// Include numeric characters in the result. Default value is `true`.
	Numeric *bool `pulumi:"numeric"`
	// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
	OverrideSpecial *string `pulumi:"overrideSpecial"`
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
	Special *bool `pulumi:"special"`
	// Include uppercase alphabet characters in the result. Default value is `true`.
	Upper *bool `pulumi:"upper"`
}

// The set of arguments for constructing a RandomPassword resource.
type RandomPasswordArgs struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers pulumi.StringMapInput
	// The length of the string desired. The minimum value for length is 1 and, length must also be >= (`minUpper` + `minLower` + `minNumeric` + `minSpecial`).
	Length pulumi.IntInput
	// Include lowercase alphabet characters in the result. Default value is `true`.
	Lower pulumi.BoolPtrInput
	// Minimum number of lowercase alphabet characters in the result. Default value is `0`.
	MinLower pulumi.IntPtrInput
	// Minimum number of numeric characters in the result. Default value is `0`.
	MinNumeric pulumi.IntPtrInput
	// Minimum number of special characters in the result. Default value is `0`.
	MinSpecial pulumi.IntPtrInput
	// Minimum number of uppercase alphabet characters in the result. Default value is `0`.
	MinUpper pulumi.IntPtrInput
	// Include numeric characters in the result. Default value is `true`. **NOTE**: This is deprecated, use `numeric` instead.
	//
	// Deprecated: **NOTE**: This is deprecated, use `numeric` instead.
	Number pulumi.BoolPtrInput
	// Include numeric characters in the result. Default value is `true`.
	Numeric pulumi.BoolPtrInput
	// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
	OverrideSpecial pulumi.StringPtrInput
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
	Special pulumi.BoolPtrInput
	// Include uppercase alphabet characters in the result. Default value is `true`.
	Upper pulumi.BoolPtrInput
}

func (RandomPasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomPasswordArgs)(nil)).Elem()
}

type RandomPasswordInput interface {
	pulumi.Input

	ToRandomPasswordOutput() RandomPasswordOutput
	ToRandomPasswordOutputWithContext(ctx context.Context) RandomPasswordOutput
}

func (*RandomPassword) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomPassword)(nil)).Elem()
}

func (i *RandomPassword) ToRandomPasswordOutput() RandomPasswordOutput {
	return i.ToRandomPasswordOutputWithContext(context.Background())
}

func (i *RandomPassword) ToRandomPasswordOutputWithContext(ctx context.Context) RandomPasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomPasswordOutput)
}

// RandomPasswordArrayInput is an input type that accepts RandomPasswordArray and RandomPasswordArrayOutput values.
// You can construct a concrete instance of `RandomPasswordArrayInput` via:
//
//	RandomPasswordArray{ RandomPasswordArgs{...} }
type RandomPasswordArrayInput interface {
	pulumi.Input

	ToRandomPasswordArrayOutput() RandomPasswordArrayOutput
	ToRandomPasswordArrayOutputWithContext(context.Context) RandomPasswordArrayOutput
}

type RandomPasswordArray []RandomPasswordInput

func (RandomPasswordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RandomPassword)(nil)).Elem()
}

func (i RandomPasswordArray) ToRandomPasswordArrayOutput() RandomPasswordArrayOutput {
	return i.ToRandomPasswordArrayOutputWithContext(context.Background())
}

func (i RandomPasswordArray) ToRandomPasswordArrayOutputWithContext(ctx context.Context) RandomPasswordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomPasswordArrayOutput)
}

// RandomPasswordMapInput is an input type that accepts RandomPasswordMap and RandomPasswordMapOutput values.
// You can construct a concrete instance of `RandomPasswordMapInput` via:
//
//	RandomPasswordMap{ "key": RandomPasswordArgs{...} }
type RandomPasswordMapInput interface {
	pulumi.Input

	ToRandomPasswordMapOutput() RandomPasswordMapOutput
	ToRandomPasswordMapOutputWithContext(context.Context) RandomPasswordMapOutput
}

type RandomPasswordMap map[string]RandomPasswordInput

func (RandomPasswordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RandomPassword)(nil)).Elem()
}

func (i RandomPasswordMap) ToRandomPasswordMapOutput() RandomPasswordMapOutput {
	return i.ToRandomPasswordMapOutputWithContext(context.Background())
}

func (i RandomPasswordMap) ToRandomPasswordMapOutputWithContext(ctx context.Context) RandomPasswordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomPasswordMapOutput)
}

type RandomPasswordOutput struct{ *pulumi.OutputState }

func (RandomPasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomPassword)(nil)).Elem()
}

func (o RandomPasswordOutput) ToRandomPasswordOutput() RandomPasswordOutput {
	return o
}

func (o RandomPasswordOutput) ToRandomPasswordOutputWithContext(ctx context.Context) RandomPasswordOutput {
	return o
}

// A bcrypt hash of the generated random string.
func (o RandomPasswordOutput) BcryptHash() pulumi.StringOutput {
	return o.ApplyT(func(v *RandomPassword) pulumi.StringOutput { return v.BcryptHash }).(pulumi.StringOutput)
}

// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
func (o RandomPasswordOutput) Keepers() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RandomPassword) pulumi.StringMapOutput { return v.Keepers }).(pulumi.StringMapOutput)
}

// The length of the string desired. The minimum value for length is 1 and, length must also be >= (`minUpper` + `minLower` + `minNumeric` + `minSpecial`).
func (o RandomPasswordOutput) Length() pulumi.IntOutput {
	return o.ApplyT(func(v *RandomPassword) pulumi.IntOutput { return v.Length }).(pulumi.IntOutput)
}

// Include lowercase alphabet characters in the result. Default value is `true`.
func (o RandomPasswordOutput) Lower() pulumi.BoolOutput {
	return o.ApplyT(func(v *RandomPassword) pulumi.BoolOutput { return v.Lower }).(pulumi.BoolOutput)
}

// Minimum number of lowercase alphabet characters in the result. Default value is `0`.
func (o RandomPasswordOutput) MinLower() pulumi.IntOutput {
	return o.ApplyT(func(v *RandomPassword) pulumi.IntOutput { return v.MinLower }).(pulumi.IntOutput)
}

// Minimum number of numeric characters in the result. Default value is `0`.
func (o RandomPasswordOutput) MinNumeric() pulumi.IntOutput {
	return o.ApplyT(func(v *RandomPassword) pulumi.IntOutput { return v.MinNumeric }).(pulumi.IntOutput)
}

// Minimum number of special characters in the result. Default value is `0`.
func (o RandomPasswordOutput) MinSpecial() pulumi.IntOutput {
	return o.ApplyT(func(v *RandomPassword) pulumi.IntOutput { return v.MinSpecial }).(pulumi.IntOutput)
}

// Minimum number of uppercase alphabet characters in the result. Default value is `0`.
func (o RandomPasswordOutput) MinUpper() pulumi.IntOutput {
	return o.ApplyT(func(v *RandomPassword) pulumi.IntOutput { return v.MinUpper }).(pulumi.IntOutput)
}

// Include numeric characters in the result. Default value is `true`. **NOTE**: This is deprecated, use `numeric` instead.
//
// Deprecated: **NOTE**: This is deprecated, use `numeric` instead.
func (o RandomPasswordOutput) Number() pulumi.BoolOutput {
	return o.ApplyT(func(v *RandomPassword) pulumi.BoolOutput { return v.Number }).(pulumi.BoolOutput)
}

// Include numeric characters in the result. Default value is `true`.
func (o RandomPasswordOutput) Numeric() pulumi.BoolOutput {
	return o.ApplyT(func(v *RandomPassword) pulumi.BoolOutput { return v.Numeric }).(pulumi.BoolOutput)
}

// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
func (o RandomPasswordOutput) OverrideSpecial() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RandomPassword) pulumi.StringPtrOutput { return v.OverrideSpecial }).(pulumi.StringPtrOutput)
}

// The generated random string.
func (o RandomPasswordOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v *RandomPassword) pulumi.StringOutput { return v.Result }).(pulumi.StringOutput)
}

// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
func (o RandomPasswordOutput) Special() pulumi.BoolOutput {
	return o.ApplyT(func(v *RandomPassword) pulumi.BoolOutput { return v.Special }).(pulumi.BoolOutput)
}

// Include uppercase alphabet characters in the result. Default value is `true`.
func (o RandomPasswordOutput) Upper() pulumi.BoolOutput {
	return o.ApplyT(func(v *RandomPassword) pulumi.BoolOutput { return v.Upper }).(pulumi.BoolOutput)
}

type RandomPasswordArrayOutput struct{ *pulumi.OutputState }

func (RandomPasswordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RandomPassword)(nil)).Elem()
}

func (o RandomPasswordArrayOutput) ToRandomPasswordArrayOutput() RandomPasswordArrayOutput {
	return o
}

func (o RandomPasswordArrayOutput) ToRandomPasswordArrayOutputWithContext(ctx context.Context) RandomPasswordArrayOutput {
	return o
}

func (o RandomPasswordArrayOutput) Index(i pulumi.IntInput) RandomPasswordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RandomPassword {
		return vs[0].([]*RandomPassword)[vs[1].(int)]
	}).(RandomPasswordOutput)
}

type RandomPasswordMapOutput struct{ *pulumi.OutputState }

func (RandomPasswordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RandomPassword)(nil)).Elem()
}

func (o RandomPasswordMapOutput) ToRandomPasswordMapOutput() RandomPasswordMapOutput {
	return o
}

func (o RandomPasswordMapOutput) ToRandomPasswordMapOutputWithContext(ctx context.Context) RandomPasswordMapOutput {
	return o
}

func (o RandomPasswordMapOutput) MapIndex(k pulumi.StringInput) RandomPasswordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RandomPassword {
		return vs[0].(map[string]*RandomPassword)[vs[1].(string)]
	}).(RandomPasswordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RandomPasswordInput)(nil)).Elem(), &RandomPassword{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomPasswordArrayInput)(nil)).Elem(), RandomPasswordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomPasswordMapInput)(nil)).Elem(), RandomPasswordMap{})
	pulumi.RegisterOutputType(RandomPasswordOutput{})
	pulumi.RegisterOutputType(RandomPasswordArrayOutput{})
	pulumi.RegisterOutputType(RandomPasswordMapOutput{})
}
