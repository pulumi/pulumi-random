// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.random;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RandomIntegerArgs extends com.pulumi.resources.ResourceArgs {

    public static final RandomIntegerArgs Empty = new RandomIntegerArgs();

    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     * 
     */
    @Import(name="keepers")
    private @Nullable Output<Map<String,String>> keepers;

    /**
     * @return Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     * 
     */
    public Optional<Output<Map<String,String>>> keepers() {
        return Optional.ofNullable(this.keepers);
    }

    /**
     * The maximum inclusive value of the range.
     * 
     */
    @Import(name="max", required=true)
    private Output<Integer> max;

    /**
     * @return The maximum inclusive value of the range.
     * 
     */
    public Output<Integer> max() {
        return this.max;
    }

    /**
     * The minimum inclusive value of the range.
     * 
     */
    @Import(name="min", required=true)
    private Output<Integer> min;

    /**
     * @return The minimum inclusive value of the range.
     * 
     */
    public Output<Integer> min() {
        return this.min;
    }

    /**
     * A custom seed to always produce the same value.
     * 
     */
    @Import(name="seed")
    private @Nullable Output<String> seed;

    /**
     * @return A custom seed to always produce the same value.
     * 
     */
    public Optional<Output<String>> seed() {
        return Optional.ofNullable(this.seed);
    }

    private RandomIntegerArgs() {}

    private RandomIntegerArgs(RandomIntegerArgs $) {
        this.keepers = $.keepers;
        this.max = $.max;
        this.min = $.min;
        this.seed = $.seed;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RandomIntegerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RandomIntegerArgs $;

        public Builder() {
            $ = new RandomIntegerArgs();
        }

        public Builder(RandomIntegerArgs defaults) {
            $ = new RandomIntegerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param keepers Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
         * 
         * @return builder
         * 
         */
        public Builder keepers(@Nullable Output<Map<String,String>> keepers) {
            $.keepers = keepers;
            return this;
        }

        /**
         * @param keepers Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
         * 
         * @return builder
         * 
         */
        public Builder keepers(Map<String,String> keepers) {
            return keepers(Output.of(keepers));
        }

        /**
         * @param max The maximum inclusive value of the range.
         * 
         * @return builder
         * 
         */
        public Builder max(Output<Integer> max) {
            $.max = max;
            return this;
        }

        /**
         * @param max The maximum inclusive value of the range.
         * 
         * @return builder
         * 
         */
        public Builder max(Integer max) {
            return max(Output.of(max));
        }

        /**
         * @param min The minimum inclusive value of the range.
         * 
         * @return builder
         * 
         */
        public Builder min(Output<Integer> min) {
            $.min = min;
            return this;
        }

        /**
         * @param min The minimum inclusive value of the range.
         * 
         * @return builder
         * 
         */
        public Builder min(Integer min) {
            return min(Output.of(min));
        }

        /**
         * @param seed A custom seed to always produce the same value.
         * 
         * @return builder
         * 
         */
        public Builder seed(@Nullable Output<String> seed) {
            $.seed = seed;
            return this;
        }

        /**
         * @param seed A custom seed to always produce the same value.
         * 
         * @return builder
         * 
         */
        public Builder seed(String seed) {
            return seed(Output.of(seed));
        }

        public RandomIntegerArgs build() {
            $.max = Objects.requireNonNull($.max, "expected parameter 'max' to be non-null");
            $.min = Objects.requireNonNull($.min, "expected parameter 'min' to be non-null");
            return $;
        }
    }

}
