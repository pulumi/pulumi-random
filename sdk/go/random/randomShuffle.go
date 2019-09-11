// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The resource `.RandomShuffle` generates a random permutation of a list
// of strings given as an argument.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-random/blob/master/website/docs/r/shuffle.html.markdown.
type RandomShuffle struct {
	s *pulumi.ResourceState
}

// NewRandomShuffle registers a new resource with the given unique name, arguments, and options.
func NewRandomShuffle(ctx *pulumi.Context,
	name string, args *RandomShuffleArgs, opts ...pulumi.ResourceOpt) (*RandomShuffle, error) {
	if args == nil || args.Inputs == nil {
		return nil, errors.New("missing required argument 'Inputs'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["inputs"] = nil
		inputs["keepers"] = nil
		inputs["resultCount"] = nil
		inputs["seed"] = nil
	} else {
		inputs["inputs"] = args.Inputs
		inputs["keepers"] = args.Keepers
		inputs["resultCount"] = args.ResultCount
		inputs["seed"] = args.Seed
	}
	inputs["results"] = nil
	s, err := ctx.RegisterResource("random:index/randomShuffle:RandomShuffle", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RandomShuffle{s: s}, nil
}

// GetRandomShuffle gets an existing RandomShuffle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomShuffle(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RandomShuffleState, opts ...pulumi.ResourceOpt) (*RandomShuffle, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["inputs"] = state.Inputs
		inputs["keepers"] = state.Keepers
		inputs["results"] = state.Results
		inputs["resultCount"] = state.ResultCount
		inputs["seed"] = state.Seed
	}
	s, err := ctx.ReadResource("random:index/randomShuffle:RandomShuffle", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RandomShuffle{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RandomShuffle) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RandomShuffle) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The list of strings to shuffle.
func (r *RandomShuffle) Inputs() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["inputs"])
}

// Arbitrary map of values that, when changed, will
// trigger a new id to be generated. See
// the main provider documentation for more information.
func (r *RandomShuffle) Keepers() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["keepers"])
}

// Random permutation of the list of strings given in `input`.
func (r *RandomShuffle) Results() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["results"])
}

// The number of results to return. Defaults to
// the number of items in the `input` list. If fewer items are requested,
// some elements will be excluded from the result. If more items are requested,
// items will be repeated in the result but not more frequently than the number
// of items in the input list.
func (r *RandomShuffle) ResultCount() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["resultCount"])
}

func (r *RandomShuffle) Seed() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["seed"])
}

// Input properties used for looking up and filtering RandomShuffle resources.
type RandomShuffleState struct {
	// The list of strings to shuffle.
	Inputs interface{}
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated. See
	// the main provider documentation for more information.
	Keepers interface{}
	// Random permutation of the list of strings given in `input`.
	Results interface{}
	// The number of results to return. Defaults to
	// the number of items in the `input` list. If fewer items are requested,
	// some elements will be excluded from the result. If more items are requested,
	// items will be repeated in the result but not more frequently than the number
	// of items in the input list.
	ResultCount interface{}
	Seed interface{}
}

// The set of arguments for constructing a RandomShuffle resource.
type RandomShuffleArgs struct {
	// The list of strings to shuffle.
	Inputs interface{}
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated. See
	// the main provider documentation for more information.
	Keepers interface{}
	// The number of results to return. Defaults to
	// the number of items in the `input` list. If fewer items are requested,
	// some elements will be excluded from the result. If more items are requested,
	// items will be repeated in the result but not more frequently than the number
	// of items in the input list.
	ResultCount interface{}
	Seed interface{}
}
