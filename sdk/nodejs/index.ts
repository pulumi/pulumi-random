import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { RandomBytesArgs, RandomBytesState } from "./randomBytes";
export type RandomBytes = import("./randomBytes").RandomBytes;
export const RandomBytes: typeof import("./randomBytes").RandomBytes = null as any;
utilities.lazyLoad(exports, ["RandomBytes"], () => require("./randomBytes"));

export { RandomIdArgs, RandomIdState } from "./randomId";
export type RandomId = import("./randomId").RandomId;
export const RandomId: typeof import("./randomId").RandomId = null as any;
utilities.lazyLoad(exports, ["RandomId"], () => require("./randomId"));

export { RandomIntegerArgs, RandomIntegerState } from "./randomInteger";
export type RandomInteger = import("./randomInteger").RandomInteger;
export const RandomInteger: typeof import("./randomInteger").RandomInteger = null as any;
utilities.lazyLoad(exports, ["RandomInteger"], () => require("./randomInteger"));

export { RandomPasswordArgs, RandomPasswordState } from "./randomPassword";
export type RandomPassword = import("./randomPassword").RandomPassword;
export const RandomPassword: typeof import("./randomPassword").RandomPassword = null as any;
utilities.lazyLoad(exports, ["RandomPassword"], () => require("./randomPassword"));

export { RandomPetArgs, RandomPetState } from "./randomPet";
export type RandomPet = import("./randomPet").RandomPet;
export const RandomPet: typeof import("./randomPet").RandomPet = null as any;
utilities.lazyLoad(exports, ["RandomPet"], () => require("./randomPet"));

export { RandomShuffleArgs, RandomShuffleState } from "./randomShuffle";
export type RandomShuffle = import("./randomShuffle").RandomShuffle;
export const RandomShuffle: typeof import("./randomShuffle").RandomShuffle = null as any;
utilities.lazyLoad(exports, ["RandomShuffle"], () => require("./randomShuffle"));

export { RandomStringArgs, RandomStringState } from "./randomString";
export type RandomString = import("./randomString").RandomString;
export const RandomString: typeof import("./randomString").RandomString = null as any;
utilities.lazyLoad(exports, ["RandomString"], () => require("./randomString"));

export { RandomUuidArgs, RandomUuidState } from "./randomUuid";
export type RandomUuid = import("./randomUuid").RandomUuid;
export const RandomUuid: typeof import("./randomUuid").RandomUuid = null as any;
utilities.lazyLoad(exports, ["RandomUuid"], () => require("./randomUuid"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "random:index/randomBytes:RandomBytes":
                return new RandomBytes(name, <any>undefined, { urn })
            case "random:index/randomId:RandomId":
                return new RandomId(name, <any>undefined, { urn })
            case "random:index/randomInteger:RandomInteger":
                return new RandomInteger(name, <any>undefined, { urn })
            case "random:index/randomPassword:RandomPassword":
                return new RandomPassword(name, <any>undefined, { urn })
            case "random:index/randomPet:RandomPet":
                return new RandomPet(name, <any>undefined, { urn })
            case "random:index/randomShuffle:RandomShuffle":
                return new RandomShuffle(name, <any>undefined, { urn })
            case "random:index/randomString:RandomString":
                return new RandomString(name, <any>undefined, { urn })
            case "random:index/randomUuid:RandomUuid":
                return new RandomUuid(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("random", "index/randomBytes", _module)
pulumi.runtime.registerResourceModule("random", "index/randomId", _module)
pulumi.runtime.registerResourceModule("random", "index/randomInteger", _module)
pulumi.runtime.registerResourceModule("random", "index/randomPassword", _module)
pulumi.runtime.registerResourceModule("random", "index/randomPet", _module)
pulumi.runtime.registerResourceModule("random", "index/randomShuffle", _module)
pulumi.runtime.registerResourceModule("random", "index/randomString", _module)
pulumi.runtime.registerResourceModule("random", "index/randomUuid", _module)
pulumi.runtime.registerResourcePackage("random", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:random") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
