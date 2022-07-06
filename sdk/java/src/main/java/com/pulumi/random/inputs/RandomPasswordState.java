// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.random.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RandomPasswordState extends com.pulumi.resources.ResourceArgs {

    public static final RandomPasswordState Empty = new RandomPasswordState();

    /**
     * A bcrypt hash of the generated random string.
     * 
     */
    @Import(name="bcryptHash")
    private @Nullable Output<String> bcryptHash;

    /**
     * @return A bcrypt hash of the generated random string.
     * 
     */
    public Optional<Output<String>> bcryptHash() {
        return Optional.ofNullable(this.bcryptHash);
    }

    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
     * documentation](../index.html) for more information.
     * 
     */
    @Import(name="keepers")
    private @Nullable Output<Map<String,Object>> keepers;

    /**
     * @return Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
     * documentation](../index.html) for more information.
     * 
     */
    public Optional<Output<Map<String,Object>>> keepers() {
        return Optional.ofNullable(this.keepers);
    }

    /**
     * The length of the string desired. The minimum value for length is 1 and, length must also be &gt;= (`min_upper` +
     * `min_lower` + `min_numeric` + `min_special`).
     * 
     */
    @Import(name="length")
    private @Nullable Output<Integer> length;

    /**
     * @return The length of the string desired. The minimum value for length is 1 and, length must also be &gt;= (`min_upper` +
     * `min_lower` + `min_numeric` + `min_special`).
     * 
     */
    public Optional<Output<Integer>> length() {
        return Optional.ofNullable(this.length);
    }

    /**
     * Include lowercase alphabet characters in the result. Default value is `true`.
     * 
     */
    @Import(name="lower")
    private @Nullable Output<Boolean> lower;

    /**
     * @return Include lowercase alphabet characters in the result. Default value is `true`.
     * 
     */
    public Optional<Output<Boolean>> lower() {
        return Optional.ofNullable(this.lower);
    }

    /**
     * Minimum number of lowercase alphabet characters in the result. Default value is `0`.
     * 
     */
    @Import(name="minLower")
    private @Nullable Output<Integer> minLower;

    /**
     * @return Minimum number of lowercase alphabet characters in the result. Default value is `0`.
     * 
     */
    public Optional<Output<Integer>> minLower() {
        return Optional.ofNullable(this.minLower);
    }

    /**
     * Minimum number of numeric characters in the result. Default value is `0`.
     * 
     */
    @Import(name="minNumeric")
    private @Nullable Output<Integer> minNumeric;

    /**
     * @return Minimum number of numeric characters in the result. Default value is `0`.
     * 
     */
    public Optional<Output<Integer>> minNumeric() {
        return Optional.ofNullable(this.minNumeric);
    }

    /**
     * Minimum number of special characters in the result. Default value is `0`.
     * 
     */
    @Import(name="minSpecial")
    private @Nullable Output<Integer> minSpecial;

    /**
     * @return Minimum number of special characters in the result. Default value is `0`.
     * 
     */
    public Optional<Output<Integer>> minSpecial() {
        return Optional.ofNullable(this.minSpecial);
    }

    /**
     * Minimum number of uppercase alphabet characters in the result. Default value is `0`.
     * 
     */
    @Import(name="minUpper")
    private @Nullable Output<Integer> minUpper;

    /**
     * @return Minimum number of uppercase alphabet characters in the result. Default value is `0`.
     * 
     */
    public Optional<Output<Integer>> minUpper() {
        return Optional.ofNullable(this.minUpper);
    }

    /**
     * Include numeric characters in the result. Default value is `true`.
     * 
     */
    @Import(name="number")
    private @Nullable Output<Boolean> number;

    /**
     * @return Include numeric characters in the result. Default value is `true`.
     * 
     */
    public Optional<Output<Boolean>> number() {
        return Optional.ofNullable(this.number);
    }

    /**
     * Supply your own list of special characters to use for string generation. This overrides the default character list in
     * the special argument. The `special` argument must still be set to true for any overwritten characters to be used in
     * generation.
     * 
     */
    @Import(name="overrideSpecial")
    private @Nullable Output<String> overrideSpecial;

    /**
     * @return Supply your own list of special characters to use for string generation. This overrides the default character list in
     * the special argument. The `special` argument must still be set to true for any overwritten characters to be used in
     * generation.
     * 
     */
    public Optional<Output<String>> overrideSpecial() {
        return Optional.ofNullable(this.overrideSpecial);
    }

    /**
     * The generated random string.
     * 
     */
    @Import(name="result")
    private @Nullable Output<String> result;

    /**
     * @return The generated random string.
     * 
     */
    public Optional<Output<String>> result() {
        return Optional.ofNullable(this.result);
    }

    /**
     * Include special characters in the result. These are `!@#$%&amp;*()-_=+[]{}&lt;&gt;:?`. Default value is `true`.
     * 
     */
    @Import(name="special")
    private @Nullable Output<Boolean> special;

    /**
     * @return Include special characters in the result. These are `!@#$%&amp;*()-_=+[]{}&lt;&gt;:?`. Default value is `true`.
     * 
     */
    public Optional<Output<Boolean>> special() {
        return Optional.ofNullable(this.special);
    }

    /**
     * Include uppercase alphabet characters in the result. Default value is `true`.
     * 
     */
    @Import(name="upper")
    private @Nullable Output<Boolean> upper;

    /**
     * @return Include uppercase alphabet characters in the result. Default value is `true`.
     * 
     */
    public Optional<Output<Boolean>> upper() {
        return Optional.ofNullable(this.upper);
    }

    private RandomPasswordState() {}

    private RandomPasswordState(RandomPasswordState $) {
        this.bcryptHash = $.bcryptHash;
        this.keepers = $.keepers;
        this.length = $.length;
        this.lower = $.lower;
        this.minLower = $.minLower;
        this.minNumeric = $.minNumeric;
        this.minSpecial = $.minSpecial;
        this.minUpper = $.minUpper;
        this.number = $.number;
        this.overrideSpecial = $.overrideSpecial;
        this.result = $.result;
        this.special = $.special;
        this.upper = $.upper;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RandomPasswordState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RandomPasswordState $;

        public Builder() {
            $ = new RandomPasswordState();
        }

        public Builder(RandomPasswordState defaults) {
            $ = new RandomPasswordState(Objects.requireNonNull(defaults));
        }

        /**
         * @param bcryptHash A bcrypt hash of the generated random string.
         * 
         * @return builder
         * 
         */
        public Builder bcryptHash(@Nullable Output<String> bcryptHash) {
            $.bcryptHash = bcryptHash;
            return this;
        }

        /**
         * @param bcryptHash A bcrypt hash of the generated random string.
         * 
         * @return builder
         * 
         */
        public Builder bcryptHash(String bcryptHash) {
            return bcryptHash(Output.of(bcryptHash));
        }

        /**
         * @param keepers Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
         * documentation](../index.html) for more information.
         * 
         * @return builder
         * 
         */
        public Builder keepers(@Nullable Output<Map<String,Object>> keepers) {
            $.keepers = keepers;
            return this;
        }

        /**
         * @param keepers Arbitrary map of values that, when changed, will trigger recreation of resource. See [the main provider
         * documentation](../index.html) for more information.
         * 
         * @return builder
         * 
         */
        public Builder keepers(Map<String,Object> keepers) {
            return keepers(Output.of(keepers));
        }

        /**
         * @param length The length of the string desired. The minimum value for length is 1 and, length must also be &gt;= (`min_upper` +
         * `min_lower` + `min_numeric` + `min_special`).
         * 
         * @return builder
         * 
         */
        public Builder length(@Nullable Output<Integer> length) {
            $.length = length;
            return this;
        }

        /**
         * @param length The length of the string desired. The minimum value for length is 1 and, length must also be &gt;= (`min_upper` +
         * `min_lower` + `min_numeric` + `min_special`).
         * 
         * @return builder
         * 
         */
        public Builder length(Integer length) {
            return length(Output.of(length));
        }

        /**
         * @param lower Include lowercase alphabet characters in the result. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder lower(@Nullable Output<Boolean> lower) {
            $.lower = lower;
            return this;
        }

        /**
         * @param lower Include lowercase alphabet characters in the result. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder lower(Boolean lower) {
            return lower(Output.of(lower));
        }

        /**
         * @param minLower Minimum number of lowercase alphabet characters in the result. Default value is `0`.
         * 
         * @return builder
         * 
         */
        public Builder minLower(@Nullable Output<Integer> minLower) {
            $.minLower = minLower;
            return this;
        }

        /**
         * @param minLower Minimum number of lowercase alphabet characters in the result. Default value is `0`.
         * 
         * @return builder
         * 
         */
        public Builder minLower(Integer minLower) {
            return minLower(Output.of(minLower));
        }

        /**
         * @param minNumeric Minimum number of numeric characters in the result. Default value is `0`.
         * 
         * @return builder
         * 
         */
        public Builder minNumeric(@Nullable Output<Integer> minNumeric) {
            $.minNumeric = minNumeric;
            return this;
        }

        /**
         * @param minNumeric Minimum number of numeric characters in the result. Default value is `0`.
         * 
         * @return builder
         * 
         */
        public Builder minNumeric(Integer minNumeric) {
            return minNumeric(Output.of(minNumeric));
        }

        /**
         * @param minSpecial Minimum number of special characters in the result. Default value is `0`.
         * 
         * @return builder
         * 
         */
        public Builder minSpecial(@Nullable Output<Integer> minSpecial) {
            $.minSpecial = minSpecial;
            return this;
        }

        /**
         * @param minSpecial Minimum number of special characters in the result. Default value is `0`.
         * 
         * @return builder
         * 
         */
        public Builder minSpecial(Integer minSpecial) {
            return minSpecial(Output.of(minSpecial));
        }

        /**
         * @param minUpper Minimum number of uppercase alphabet characters in the result. Default value is `0`.
         * 
         * @return builder
         * 
         */
        public Builder minUpper(@Nullable Output<Integer> minUpper) {
            $.minUpper = minUpper;
            return this;
        }

        /**
         * @param minUpper Minimum number of uppercase alphabet characters in the result. Default value is `0`.
         * 
         * @return builder
         * 
         */
        public Builder minUpper(Integer minUpper) {
            return minUpper(Output.of(minUpper));
        }

        /**
         * @param number Include numeric characters in the result. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder number(@Nullable Output<Boolean> number) {
            $.number = number;
            return this;
        }

        /**
         * @param number Include numeric characters in the result. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder number(Boolean number) {
            return number(Output.of(number));
        }

        /**
         * @param overrideSpecial Supply your own list of special characters to use for string generation. This overrides the default character list in
         * the special argument. The `special` argument must still be set to true for any overwritten characters to be used in
         * generation.
         * 
         * @return builder
         * 
         */
        public Builder overrideSpecial(@Nullable Output<String> overrideSpecial) {
            $.overrideSpecial = overrideSpecial;
            return this;
        }

        /**
         * @param overrideSpecial Supply your own list of special characters to use for string generation. This overrides the default character list in
         * the special argument. The `special` argument must still be set to true for any overwritten characters to be used in
         * generation.
         * 
         * @return builder
         * 
         */
        public Builder overrideSpecial(String overrideSpecial) {
            return overrideSpecial(Output.of(overrideSpecial));
        }

        /**
         * @param result The generated random string.
         * 
         * @return builder
         * 
         */
        public Builder result(@Nullable Output<String> result) {
            $.result = result;
            return this;
        }

        /**
         * @param result The generated random string.
         * 
         * @return builder
         * 
         */
        public Builder result(String result) {
            return result(Output.of(result));
        }

        /**
         * @param special Include special characters in the result. These are `!@#$%&amp;*()-_=+[]{}&lt;&gt;:?`. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder special(@Nullable Output<Boolean> special) {
            $.special = special;
            return this;
        }

        /**
         * @param special Include special characters in the result. These are `!@#$%&amp;*()-_=+[]{}&lt;&gt;:?`. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder special(Boolean special) {
            return special(Output.of(special));
        }

        /**
         * @param upper Include uppercase alphabet characters in the result. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder upper(@Nullable Output<Boolean> upper) {
            $.upper = upper;
            return this;
        }

        /**
         * @param upper Include uppercase alphabet characters in the result. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder upper(Boolean upper) {
            return upper(Output.of(upper));
        }

        public RandomPasswordState build() {
            return $;
        }
    }

}
