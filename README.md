[![Actions Status](https://github.com/pulumi/pulumi-random/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-random/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Frandom.svg)](https://npmjs.com/package/@pulumi/random)
[![NuGet version](https://badge.fury.io/nu/pulumi.random.svg)](https://badge.fury.io/nu/pulumi.random)
[![Python version](https://badge.fury.io/py/pulumi-random.svg)](https://pypi.org/project/pulumi-random)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-random/sdk/v4/go)](https://pkg.go.dev/github.com/pulumi/pulumi-random/sdk/v4/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-random/blob/master/LICENSE)

# Random Provider

The random provider allows the safe use of randomness in a Pulumi program. This allows you to generate resource
properties, such as names, that contain randomness in a way that works with Pulumi's goal state oriented approach.
Using randomness as usual would not work well with Pulumi, because by definition, each time the program is evaluated,
a new random state would be produced, necessitating re-convergence on the goal state. This provider understands
how to work with the Pulumi resource lifecycle to accomplish randomness safely and in a way that works as desired.

## Example

For example, to generate a random password, allocate a `RandomPassword` resource
and then use its `result` output property (of type `Output<string>`) to pass
to another resource.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as random from "@pulumi/random";

const password = new random.RandomPassword("password", {
    length: 16,
    overrideSpecial: "_%@",
    special: true,
});
const example = new aws.rds.Instance("example", {
    password: password.result,
});
```

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/random

or `yarn`:

    $ yarn add @pulumi/random

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_random

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-random/sdk/v4/go/...

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Random

## Reference

For further information, please visit [the random provider docs](https://www.pulumi.com/docs/intro/cloud-providers/random) or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/random).
