// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The resource `random.RandomInteger` generates random values from a given range, described by the `min` and `max` attributes of a given resource.
 *
 * This resource can be used in conjunction with resources that have
 * the `createBeforeDestroy` lifecycle flag set, to avoid conflicts with
 * unique names during the brief period where both the old and new resources
 * exist concurrently.
 *
 * ## Example Usage
 *
 * The following example shows how to generate a random priority between 1 and 50000 for
 * a `awsAlbListenerRule` resource:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as random from "@pulumi/random";
 *
 * const priority = new random.RandomInteger("priority", {
 *     keepers: {
 *         // Generate a new integer each time we switch to a new listener ARN
 *         listener_arn: var_listener_arn,
 *     },
 *     max: 50000,
 *     min: 1,
 * });
 * const main = new aws.alb.ListenerRule("main", {
 *     actions: [{
 *         targetGroupArn: var_target_group_arn,
 *         type: "forward",
 *     }],
 *     listenerArn: var_listener_arn,
 *     priority: priority.result,
 * });
 * ```
 *
 * The result of the above will set a random priority.
 *
 * ## Import
 *
 * Random integers can be imported using the `result`, `min`, and `max`, with an optional `seed`. This can be used to replace a config value with a value interpolated from the random provider without experiencing diffs. Example (values are separated by a `,`)
 *
 * ```sh
 *  $ pulumi import random:index/randomInteger:RandomInteger priority 15390,1,50000
 * ```
 */
export class RandomInteger extends pulumi.CustomResource {
    /**
     * Get an existing RandomInteger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RandomIntegerState, opts?: pulumi.CustomResourceOptions): RandomInteger {
        return new RandomInteger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'random:index/randomInteger:RandomInteger';

    /**
     * Returns true if the given object is an instance of RandomInteger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RandomInteger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RandomInteger.__pulumiType;
    }

    /**
     * Arbitrary map of values that, when changed, will
     * trigger a new id to be generated. See
     * the main provider documentation for more information.
     */
    public readonly keepers!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The maximum inclusive value of the range.
     */
    public readonly max!: pulumi.Output<number>;
    /**
     * The minimum inclusive value of the range.
     */
    public readonly min!: pulumi.Output<number>;
    /**
     * (int) The random Integer result.
     */
    public /*out*/ readonly result!: pulumi.Output<number>;
    /**
     * A custom seed to always produce the same value.
     */
    public readonly seed!: pulumi.Output<string | undefined>;

    /**
     * Create a RandomInteger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RandomIntegerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RandomIntegerArgs | RandomIntegerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RandomIntegerState | undefined;
            resourceInputs["keepers"] = state ? state.keepers : undefined;
            resourceInputs["max"] = state ? state.max : undefined;
            resourceInputs["min"] = state ? state.min : undefined;
            resourceInputs["result"] = state ? state.result : undefined;
            resourceInputs["seed"] = state ? state.seed : undefined;
        } else {
            const args = argsOrState as RandomIntegerArgs | undefined;
            if ((!args || args.max === undefined) && !opts.urn) {
                throw new Error("Missing required property 'max'");
            }
            if ((!args || args.min === undefined) && !opts.urn) {
                throw new Error("Missing required property 'min'");
            }
            resourceInputs["keepers"] = args ? args.keepers : undefined;
            resourceInputs["max"] = args ? args.max : undefined;
            resourceInputs["min"] = args ? args.min : undefined;
            resourceInputs["seed"] = args ? args.seed : undefined;
            resourceInputs["result"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RandomInteger.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RandomInteger resources.
 */
export interface RandomIntegerState {
    /**
     * Arbitrary map of values that, when changed, will
     * trigger a new id to be generated. See
     * the main provider documentation for more information.
     */
    keepers?: pulumi.Input<{[key: string]: any}>;
    /**
     * The maximum inclusive value of the range.
     */
    max?: pulumi.Input<number>;
    /**
     * The minimum inclusive value of the range.
     */
    min?: pulumi.Input<number>;
    /**
     * (int) The random Integer result.
     */
    result?: pulumi.Input<number>;
    /**
     * A custom seed to always produce the same value.
     */
    seed?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RandomInteger resource.
 */
export interface RandomIntegerArgs {
    /**
     * Arbitrary map of values that, when changed, will
     * trigger a new id to be generated. See
     * the main provider documentation for more information.
     */
    keepers?: pulumi.Input<{[key: string]: any}>;
    /**
     * The maximum inclusive value of the range.
     */
    max: pulumi.Input<number>;
    /**
     * The minimum inclusive value of the range.
     */
    min: pulumi.Input<number>;
    /**
     * A custom seed to always produce the same value.
     */
    seed?: pulumi.Input<string>;
}
