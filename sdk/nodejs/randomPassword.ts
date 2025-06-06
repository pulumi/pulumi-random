// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as random from "@pulumi/random";
 *
 * const password = new random.RandomPassword("password", {
 *     length: 16,
 *     special: true,
 *     overrideSpecial: "!#$%&*()-_=+[]{}<>:?",
 * });
 * const example = new aws.rds.Instance("example", {
 *     instanceClass: aws.rds.InstanceType.T3_Micro,
 *     allocatedStorage: 64,
 *     engine: "mysql",
 *     username: "someone",
 *     password: password.result,
 * });
 * ```
 *
 * ## Import
 *
 * You can import external passwords into your Pulumi programs as follows:
 */
export class RandomPassword extends pulumi.CustomResource {
    /**
     * Get an existing RandomPassword resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RandomPasswordState, opts?: pulumi.CustomResourceOptions): RandomPassword {
        return new RandomPassword(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'random:index/randomPassword:RandomPassword';

    /**
     * Returns true if the given object is an instance of RandomPassword.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RandomPassword {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RandomPassword.__pulumiType;
    }

    /**
     * A bcrypt hash of the generated random string. **NOTE**: If the generated random string is greater than 72 bytes in length, `bcryptHash` will contain a hash of the first 72 bytes.
     */
    public /*out*/ readonly bcryptHash!: pulumi.Output<string>;
    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     */
    public readonly keepers!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The length of the string desired. The minimum value for length is 1 and, length must also be >= (`minUpper` + `minLower` + `minNumeric` + `minSpecial`).
     */
    public readonly length!: pulumi.Output<number>;
    /**
     * Include lowercase alphabet characters in the result. Default value is `true`.
     */
    public readonly lower!: pulumi.Output<boolean>;
    /**
     * Minimum number of lowercase alphabet characters in the result. Default value is `0`.
     */
    public readonly minLower!: pulumi.Output<number>;
    /**
     * Minimum number of numeric characters in the result. Default value is `0`.
     */
    public readonly minNumeric!: pulumi.Output<number>;
    /**
     * Minimum number of special characters in the result. Default value is `0`.
     */
    public readonly minSpecial!: pulumi.Output<number>;
    /**
     * Minimum number of uppercase alphabet characters in the result. Default value is `0`.
     */
    public readonly minUpper!: pulumi.Output<number>;
    /**
     * Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
     *
     * @deprecated **NOTE**: This is deprecated, use `numeric` instead.
     */
    public readonly number!: pulumi.Output<boolean>;
    /**
     * Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
     */
    public readonly numeric!: pulumi.Output<boolean>;
    /**
     * Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
     */
    public readonly overrideSpecial!: pulumi.Output<string | undefined>;
    /**
     * The generated random string.
     */
    public /*out*/ readonly result!: pulumi.Output<string>;
    /**
     * Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
     */
    public readonly special!: pulumi.Output<boolean>;
    /**
     * Include uppercase alphabet characters in the result. Default value is `true`.
     */
    public readonly upper!: pulumi.Output<boolean>;

    /**
     * Create a RandomPassword resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RandomPasswordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RandomPasswordArgs | RandomPasswordState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RandomPasswordState | undefined;
            resourceInputs["bcryptHash"] = state ? state.bcryptHash : undefined;
            resourceInputs["keepers"] = state ? state.keepers : undefined;
            resourceInputs["length"] = state ? state.length : undefined;
            resourceInputs["lower"] = state ? state.lower : undefined;
            resourceInputs["minLower"] = state ? state.minLower : undefined;
            resourceInputs["minNumeric"] = state ? state.minNumeric : undefined;
            resourceInputs["minSpecial"] = state ? state.minSpecial : undefined;
            resourceInputs["minUpper"] = state ? state.minUpper : undefined;
            resourceInputs["number"] = state ? state.number : undefined;
            resourceInputs["numeric"] = state ? state.numeric : undefined;
            resourceInputs["overrideSpecial"] = state ? state.overrideSpecial : undefined;
            resourceInputs["result"] = state ? state.result : undefined;
            resourceInputs["special"] = state ? state.special : undefined;
            resourceInputs["upper"] = state ? state.upper : undefined;
        } else {
            const args = argsOrState as RandomPasswordArgs | undefined;
            if ((!args || args.length === undefined) && !opts.urn) {
                throw new Error("Missing required property 'length'");
            }
            resourceInputs["keepers"] = args ? args.keepers : undefined;
            resourceInputs["length"] = args ? args.length : undefined;
            resourceInputs["lower"] = args ? args.lower : undefined;
            resourceInputs["minLower"] = args ? args.minLower : undefined;
            resourceInputs["minNumeric"] = args ? args.minNumeric : undefined;
            resourceInputs["minSpecial"] = args ? args.minSpecial : undefined;
            resourceInputs["minUpper"] = args ? args.minUpper : undefined;
            resourceInputs["number"] = args ? args.number : undefined;
            resourceInputs["numeric"] = args ? args.numeric : undefined;
            resourceInputs["overrideSpecial"] = args ? args.overrideSpecial : undefined;
            resourceInputs["special"] = args ? args.special : undefined;
            resourceInputs["upper"] = args ? args.upper : undefined;
            resourceInputs["bcryptHash"] = undefined /*out*/;
            resourceInputs["result"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["bcryptHash", "result"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(RandomPassword.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RandomPassword resources.
 */
export interface RandomPasswordState {
    /**
     * A bcrypt hash of the generated random string. **NOTE**: If the generated random string is greater than 72 bytes in length, `bcryptHash` will contain a hash of the first 72 bytes.
     */
    bcryptHash?: pulumi.Input<string>;
    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     */
    keepers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The length of the string desired. The minimum value for length is 1 and, length must also be >= (`minUpper` + `minLower` + `minNumeric` + `minSpecial`).
     */
    length?: pulumi.Input<number>;
    /**
     * Include lowercase alphabet characters in the result. Default value is `true`.
     */
    lower?: pulumi.Input<boolean>;
    /**
     * Minimum number of lowercase alphabet characters in the result. Default value is `0`.
     */
    minLower?: pulumi.Input<number>;
    /**
     * Minimum number of numeric characters in the result. Default value is `0`.
     */
    minNumeric?: pulumi.Input<number>;
    /**
     * Minimum number of special characters in the result. Default value is `0`.
     */
    minSpecial?: pulumi.Input<number>;
    /**
     * Minimum number of uppercase alphabet characters in the result. Default value is `0`.
     */
    minUpper?: pulumi.Input<number>;
    /**
     * Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
     *
     * @deprecated **NOTE**: This is deprecated, use `numeric` instead.
     */
    number?: pulumi.Input<boolean>;
    /**
     * Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
     */
    numeric?: pulumi.Input<boolean>;
    /**
     * Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
     */
    overrideSpecial?: pulumi.Input<string>;
    /**
     * The generated random string.
     */
    result?: pulumi.Input<string>;
    /**
     * Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
     */
    special?: pulumi.Input<boolean>;
    /**
     * Include uppercase alphabet characters in the result. Default value is `true`.
     */
    upper?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a RandomPassword resource.
 */
export interface RandomPasswordArgs {
    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     */
    keepers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The length of the string desired. The minimum value for length is 1 and, length must also be >= (`minUpper` + `minLower` + `minNumeric` + `minSpecial`).
     */
    length: pulumi.Input<number>;
    /**
     * Include lowercase alphabet characters in the result. Default value is `true`.
     */
    lower?: pulumi.Input<boolean>;
    /**
     * Minimum number of lowercase alphabet characters in the result. Default value is `0`.
     */
    minLower?: pulumi.Input<number>;
    /**
     * Minimum number of numeric characters in the result. Default value is `0`.
     */
    minNumeric?: pulumi.Input<number>;
    /**
     * Minimum number of special characters in the result. Default value is `0`.
     */
    minSpecial?: pulumi.Input<number>;
    /**
     * Minimum number of uppercase alphabet characters in the result. Default value is `0`.
     */
    minUpper?: pulumi.Input<number>;
    /**
     * Include numeric characters in the result. Default value is `true`. If `number`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`. **NOTE**: This is deprecated, use `numeric` instead.
     *
     * @deprecated **NOTE**: This is deprecated, use `numeric` instead.
     */
    number?: pulumi.Input<boolean>;
    /**
     * Include numeric characters in the result. Default value is `true`. If `numeric`, `upper`, `lower`, and `special` are all configured, at least one of them must be set to `true`.
     */
    numeric?: pulumi.Input<boolean>;
    /**
     * Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
     */
    overrideSpecial?: pulumi.Input<string>;
    /**
     * Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`. Default value is `true`.
     */
    special?: pulumi.Input<boolean>;
    /**
     * Include uppercase alphabet characters in the result. Default value is `true`.
     */
    upper?: pulumi.Input<boolean>;
}
