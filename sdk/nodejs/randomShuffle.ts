// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The resource `random.RandomShuffle` generates a random permutation of a list of strings given as an argument.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as random from "@pulumi/random";
 *
 * const az = new random.RandomShuffle("az", {
 *     inputs: [
 *         "us-west-1a",
 *         "us-west-1c",
 *         "us-west-1d",
 *         "us-west-1e",
 *     ],
 *     resultCount: 2,
 * });
 * const example = new aws.elb/loadBalancer.LoadBalancer("example", {availabilityZones: az.results});
 * ```
 */
export class RandomShuffle extends pulumi.CustomResource {
    /**
     * Get an existing RandomShuffle resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RandomShuffleState, opts?: pulumi.CustomResourceOptions): RandomShuffle {
        return new RandomShuffle(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'random:index/randomShuffle:RandomShuffle';

    /**
     * Returns true if the given object is an instance of RandomShuffle.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RandomShuffle {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RandomShuffle.__pulumiType;
    }

    /**
     * The list of strings to shuffle.
     */
    public readonly inputs!: pulumi.Output<string[]>;
    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     */
    public readonly keepers!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.
     */
    public readonly resultCount!: pulumi.Output<number | undefined>;
    /**
     * Random permutation of the list of strings given in `input`. The number of elements is determined by `resultCount` if
     * set, or the number of elements in `input`.
     */
    public /*out*/ readonly results!: pulumi.Output<string[]>;
    /**
     * Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.
     */
    public readonly seed!: pulumi.Output<string | undefined>;

    /**
     * Create a RandomShuffle resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RandomShuffleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RandomShuffleArgs | RandomShuffleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RandomShuffleState | undefined;
            resourceInputs["inputs"] = state ? state.inputs : undefined;
            resourceInputs["keepers"] = state ? state.keepers : undefined;
            resourceInputs["resultCount"] = state ? state.resultCount : undefined;
            resourceInputs["results"] = state ? state.results : undefined;
            resourceInputs["seed"] = state ? state.seed : undefined;
        } else {
            const args = argsOrState as RandomShuffleArgs | undefined;
            if ((!args || args.inputs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'inputs'");
            }
            resourceInputs["inputs"] = args ? args.inputs : undefined;
            resourceInputs["keepers"] = args ? args.keepers : undefined;
            resourceInputs["resultCount"] = args ? args.resultCount : undefined;
            resourceInputs["seed"] = args ? args.seed : undefined;
            resourceInputs["results"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RandomShuffle.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RandomShuffle resources.
 */
export interface RandomShuffleState {
    /**
     * The list of strings to shuffle.
     */
    inputs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     */
    keepers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.
     */
    resultCount?: pulumi.Input<number>;
    /**
     * Random permutation of the list of strings given in `input`. The number of elements is determined by `resultCount` if
     * set, or the number of elements in `input`.
     */
    results?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.
     */
    seed?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RandomShuffle resource.
 */
export interface RandomShuffleArgs {
    /**
     * The list of strings to shuffle.
     */
    inputs: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     */
    keepers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The number of results to return. Defaults to the number of items in the `input` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.
     */
    resultCount?: pulumi.Input<number>;
    /**
     * Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.
     */
    seed?: pulumi.Input<string>;
}
