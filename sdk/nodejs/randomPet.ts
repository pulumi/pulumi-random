// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The resource `random.RandomPet` generates random pet names that are intended to be used as unique identifiers for other resources.
 *
 * This resource can be used in conjunction with resources that have the `createBeforeDestroy` lifecycle flag set, to avoid conflicts with unique names during the brief period where both the old and new resources exist concurrently.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as random from "@pulumi/random";
 *
 * // The following example shows how to generate a unique pet name
 * // for an AWS EC2 instance that changes each time a new AMI id is
 * // selected.
 * const server = new random.RandomPet("server", {keepers: {
 *     ami_id: amiId,
 * }});
 * const serverInstance = new aws.ec2.Instance("server", {
 *     tags: {
 *         Name: pulumi.interpolate`web-server-${server.id}`,
 *     },
 *     ami: server.keepers.apply(keepers => keepers?.amiId),
 * });
 * ```
 */
export class RandomPet extends pulumi.CustomResource {
    /**
     * Get an existing RandomPet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RandomPetState, opts?: pulumi.CustomResourceOptions): RandomPet {
        return new RandomPet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'random:index/randomPet:RandomPet';

    /**
     * Returns true if the given object is an instance of RandomPet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RandomPet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RandomPet.__pulumiType;
    }

    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     */
    public readonly keepers!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The length (in words) of the pet name. Defaults to 2
     */
    public readonly length!: pulumi.Output<number>;
    /**
     * A string to prefix the name with.
     */
    public readonly prefix!: pulumi.Output<string | undefined>;
    /**
     * The character to separate words in the pet name. Defaults to "-"
     */
    public readonly separator!: pulumi.Output<string>;

    /**
     * Create a RandomPet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RandomPetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RandomPetArgs | RandomPetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RandomPetState | undefined;
            resourceInputs["keepers"] = state ? state.keepers : undefined;
            resourceInputs["length"] = state ? state.length : undefined;
            resourceInputs["prefix"] = state ? state.prefix : undefined;
            resourceInputs["separator"] = state ? state.separator : undefined;
        } else {
            const args = argsOrState as RandomPetArgs | undefined;
            resourceInputs["keepers"] = args ? args.keepers : undefined;
            resourceInputs["length"] = args ? args.length : undefined;
            resourceInputs["prefix"] = args ? args.prefix : undefined;
            resourceInputs["separator"] = args ? args.separator : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RandomPet.__pulumiType, name, resourceInputs, opts, false /*remote*/);
    }
}

/**
 * Input properties used for looking up and filtering RandomPet resources.
 */
export interface RandomPetState {
    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     */
    keepers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The length (in words) of the pet name. Defaults to 2
     */
    length?: pulumi.Input<number>;
    /**
     * A string to prefix the name with.
     */
    prefix?: pulumi.Input<string>;
    /**
     * The character to separate words in the pet name. Defaults to "-"
     */
    separator?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RandomPet resource.
 */
export interface RandomPetArgs {
    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     */
    keepers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The length (in words) of the pet name. Defaults to 2
     */
    length?: pulumi.Input<number>;
    /**
     * A string to prefix the name with.
     */
    prefix?: pulumi.Input<string>;
    /**
     * The character to separate words in the pet name. Defaults to "-"
     */
    separator?: pulumi.Input<string>;
}
