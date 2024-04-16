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

// The resource `RandomString` generates a random permutation of alphanumeric characters and optionally special characters.
//
// This resource *does* use a cryptographic random number generator.
//
// Historically this resource's intended usage has been ambiguous as the original example used it in a password. For backwards compatibility it will continue to exist. For unique ids please use random_id, for sensitive random values please use random_password.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := random.NewRandomString(ctx, "random", &random.RandomStringArgs{
//				Length:          pulumi.Int(16),
//				Special:         pulumi.Bool(true),
//				OverrideSpecial: pulumi.String("/@£$"),
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
// You can import external strings into your Pulumi programs as RandomString resources as follows:
type RandomString struct {
	pulumi.CustomResourceState

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
	// Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
	//
	// Deprecated: **NOTE**: This is deprecated, use `numeric` instead.
	Number pulumi.BoolOutput `pulumi:"number"`
	// Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
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

// NewRandomString registers a new resource with the given unique name, arguments, and options.
func NewRandomString(ctx *pulumi.Context,
	name string, args *RandomStringArgs, opts ...pulumi.ResourceOption) (*RandomString, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Length == nil {
		return nil, errors.New("invalid value for required argument 'Length'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RandomString
	err := ctx.RegisterResource("random:index/randomString:RandomString", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandomString gets an existing RandomString resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomString(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomStringState, opts ...pulumi.ResourceOption) (*RandomString, error) {
	var resource RandomString
	err := ctx.ReadResource("random:index/randomString:RandomString", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RandomString resources.
type randomStringState struct {
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
	// Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
	//
	// Deprecated: **NOTE**: This is deprecated, use `numeric` instead.
	Number *bool `pulumi:"number"`
	// Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
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

type RandomStringState struct {
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
	// Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
	//
	// Deprecated: **NOTE**: This is deprecated, use `numeric` instead.
	Number pulumi.BoolPtrInput
	// Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
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

func (RandomStringState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomStringState)(nil)).Elem()
}

type randomStringArgs struct {
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
	// Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
	//
	// Deprecated: **NOTE**: This is deprecated, use `numeric` instead.
	Number *bool `pulumi:"number"`
	// Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
	Numeric *bool `pulumi:"numeric"`
	// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
	OverrideSpecial *string `pulumi:"overrideSpecial"`
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
	Special *bool `pulumi:"special"`
	// Include uppercase alphabet characters in the result. Default value is `true`.
	Upper *bool `pulumi:"upper"`
}

// The set of arguments for constructing a RandomString resource.
type RandomStringArgs struct {
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
	// Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
	//
	// Deprecated: **NOTE**: This is deprecated, use `numeric` instead.
	Number pulumi.BoolPtrInput
	// Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
	Numeric pulumi.BoolPtrInput
	// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
	OverrideSpecial pulumi.StringPtrInput
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
	Special pulumi.BoolPtrInput
	// Include uppercase alphabet characters in the result. Default value is `true`.
	Upper pulumi.BoolPtrInput
}

func (RandomStringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomStringArgs)(nil)).Elem()
}

type RandomStringInput interface {
	pulumi.Input

	ToRandomStringOutput() RandomStringOutput
	ToRandomStringOutputWithContext(ctx context.Context) RandomStringOutput
}

func (*RandomString) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomString)(nil)).Elem()
}

func (i *RandomString) ToRandomStringOutput() RandomStringOutput {
	return i.ToRandomStringOutputWithContext(context.Background())
}

func (i *RandomString) ToRandomStringOutputWithContext(ctx context.Context) RandomStringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomStringOutput)
}

// RandomStringArrayInput is an input type that accepts RandomStringArray and RandomStringArrayOutput values.
// You can construct a concrete instance of `RandomStringArrayInput` via:
//
//	RandomStringArray{ RandomStringArgs{...} }
type RandomStringArrayInput interface {
	pulumi.Input

	ToRandomStringArrayOutput() RandomStringArrayOutput
	ToRandomStringArrayOutputWithContext(context.Context) RandomStringArrayOutput
}

type RandomStringArray []RandomStringInput

func (RandomStringArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RandomString)(nil)).Elem()
}

func (i RandomStringArray) ToRandomStringArrayOutput() RandomStringArrayOutput {
	return i.ToRandomStringArrayOutputWithContext(context.Background())
}

func (i RandomStringArray) ToRandomStringArrayOutputWithContext(ctx context.Context) RandomStringArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomStringArrayOutput)
}

// RandomStringMapInput is an input type that accepts RandomStringMap and RandomStringMapOutput values.
// You can construct a concrete instance of `RandomStringMapInput` via:
//
//	RandomStringMap{ "key": RandomStringArgs{...} }
type RandomStringMapInput interface {
	pulumi.Input

	ToRandomStringMapOutput() RandomStringMapOutput
	ToRandomStringMapOutputWithContext(context.Context) RandomStringMapOutput
}

type RandomStringMap map[string]RandomStringInput

func (RandomStringMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RandomString)(nil)).Elem()
}

func (i RandomStringMap) ToRandomStringMapOutput() RandomStringMapOutput {
	return i.ToRandomStringMapOutputWithContext(context.Background())
}

func (i RandomStringMap) ToRandomStringMapOutputWithContext(ctx context.Context) RandomStringMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomStringMapOutput)
}

type RandomStringOutput struct{ *pulumi.OutputState }

func (RandomStringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomString)(nil)).Elem()
}

func (o RandomStringOutput) ToRandomStringOutput() RandomStringOutput {
	return o
}

func (o RandomStringOutput) ToRandomStringOutputWithContext(ctx context.Context) RandomStringOutput {
	return o
}

// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
func (o RandomStringOutput) Keepers() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RandomString) pulumi.StringMapOutput { return v.Keepers }).(pulumi.StringMapOutput)
}

// The length of the string desired. The minimum value for length is 1 and, length must also be >= (`minUpper` + `minLower` + `minNumeric` + `minSpecial`).
func (o RandomStringOutput) Length() pulumi.IntOutput {
	return o.ApplyT(func(v *RandomString) pulumi.IntOutput { return v.Length }).(pulumi.IntOutput)
}

// Include lowercase alphabet characters in the result. Default value is `true`.
func (o RandomStringOutput) Lower() pulumi.BoolOutput {
	return o.ApplyT(func(v *RandomString) pulumi.BoolOutput { return v.Lower }).(pulumi.BoolOutput)
}

// Minimum number of lowercase alphabet characters in the result. Default value is `0`.
func (o RandomStringOutput) MinLower() pulumi.IntOutput {
	return o.ApplyT(func(v *RandomString) pulumi.IntOutput { return v.MinLower }).(pulumi.IntOutput)
}

// Minimum number of numeric characters in the result. Default value is `0`.
func (o RandomStringOutput) MinNumeric() pulumi.IntOutput {
	return o.ApplyT(func(v *RandomString) pulumi.IntOutput { return v.MinNumeric }).(pulumi.IntOutput)
}

// Minimum number of special characters in the result. Default value is `0`.
func (o RandomStringOutput) MinSpecial() pulumi.IntOutput {
	return o.ApplyT(func(v *RandomString) pulumi.IntOutput { return v.MinSpecial }).(pulumi.IntOutput)
}

// Minimum number of uppercase alphabet characters in the result. Default value is `0`.
func (o RandomStringOutput) MinUpper() pulumi.IntOutput {
	return o.ApplyT(func(v *RandomString) pulumi.IntOutput { return v.MinUpper }).(pulumi.IntOutput)
}

// Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
//
// Deprecated: **NOTE**: This is deprecated, use `numeric` instead.
func (o RandomStringOutput) Number() pulumi.BoolOutput {
	return o.ApplyT(func(v *RandomString) pulumi.BoolOutput { return v.Number }).(pulumi.BoolOutput)
}

// Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
func (o RandomStringOutput) Numeric() pulumi.BoolOutput {
	return o.ApplyT(func(v *RandomString) pulumi.BoolOutput { return v.Numeric }).(pulumi.BoolOutput)
}

// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
func (o RandomStringOutput) OverrideSpecial() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RandomString) pulumi.StringPtrOutput { return v.OverrideSpecial }).(pulumi.StringPtrOutput)
}

// The generated random string.
func (o RandomStringOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v *RandomString) pulumi.StringOutput { return v.Result }).(pulumi.StringOutput)
}

// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
func (o RandomStringOutput) Special() pulumi.BoolOutput {
	return o.ApplyT(func(v *RandomString) pulumi.BoolOutput { return v.Special }).(pulumi.BoolOutput)
}

// Include uppercase alphabet characters in the result. Default value is `true`.
func (o RandomStringOutput) Upper() pulumi.BoolOutput {
	return o.ApplyT(func(v *RandomString) pulumi.BoolOutput { return v.Upper }).(pulumi.BoolOutput)
}

type RandomStringArrayOutput struct{ *pulumi.OutputState }

func (RandomStringArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RandomString)(nil)).Elem()
}

func (o RandomStringArrayOutput) ToRandomStringArrayOutput() RandomStringArrayOutput {
	return o
}

func (o RandomStringArrayOutput) ToRandomStringArrayOutputWithContext(ctx context.Context) RandomStringArrayOutput {
	return o
}

func (o RandomStringArrayOutput) Index(i pulumi.IntInput) RandomStringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RandomString {
		return vs[0].([]*RandomString)[vs[1].(int)]
	}).(RandomStringOutput)
}

type RandomStringMapOutput struct{ *pulumi.OutputState }

func (RandomStringMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RandomString)(nil)).Elem()
}

func (o RandomStringMapOutput) ToRandomStringMapOutput() RandomStringMapOutput {
	return o
}

func (o RandomStringMapOutput) ToRandomStringMapOutputWithContext(ctx context.Context) RandomStringMapOutput {
	return o
}

func (o RandomStringMapOutput) MapIndex(k pulumi.StringInput) RandomStringOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RandomString {
		return vs[0].(map[string]*RandomString)[vs[1].(string)]
	}).(RandomStringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RandomStringInput)(nil)).Elem(), &RandomString{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomStringArrayInput)(nil)).Elem(), RandomStringArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomStringMapInput)(nil)).Elem(), RandomStringMap{})
	pulumi.RegisterOutputType(RandomStringOutput{})
	pulumi.RegisterOutputType(RandomStringArrayOutput{})
	pulumi.RegisterOutputType(RandomStringMapOutput{})
}
