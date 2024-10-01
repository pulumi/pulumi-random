---
title: Random Provider
meta_desc: Provides an overview on how to configure the Pulumi Random provider.
layout: package
---
## Installation

The random provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/random`](https://www.npmjs.com/package/@pulumi/random)
* Python: [`pulumi-random`](https://pypi.org/project/pulumi-random/)
* Go: [`github.com/pulumi/pulumi-random/sdk/v4/go/random`](https://github.com/pulumi/pulumi-random)
* .NET: [`Pulumi.Random`](https://www.nuget.org/packages/Pulumi.Random)
* Java: [`com.pulumi/random`](https://central.sonatype.com/artifact/com.pulumi/random)
## Overview

The "random" provider allows the use of randomness within Pulumi
configurations. This is a *logical provider*, which means that it works
entirely within Pulumi's logic, and doesn't interact with any other
services.

Unconstrained randomness within a Pulumi configuration would not be very
useful, since Pulumi's goal is to converge on a fixed configuration by
applying a diff. Because of this, the "random" provider provides an idea of
*managed randomness*: it provides resources that generate random values during
their creation and then hold those values steady until the inputs are changed.

Even with these resources, it is advisable to keep the use of randomness within
Pulumi configuration to a minimum, and retain it for special cases only;
Pulumi works best when the configuration is well-defined, since its behavior
can then be more readily predicted.

Unless otherwise stated within the documentation of a specific resource, this
provider's results are **not** sufficiently random for cryptographic use.

For more information on the specific resources available, see the links in the
navigation bar. Read on for information on the general patterns that apply
to this provider's resources.
## Resource "Keepers"

As noted above, the random resources generate randomness only when they are
created; the results produced are stored in the Pulumi state and re-used
until the inputs change, prompting the resource to be recreated.

The resources all provide a map argument called `keepers` that can be populated
with arbitrary key/value pairs that should be selected such that they remain
the same until new random values are desired.

For example:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as random from "@pulumi/random";

const server = new random.RandomId("server", {
    keepers: {
        ami_id: amiId,
    },
    byteLength: 8,
});
const serverInstance = new aws.ec2.Instance("server", {
    tags: {
        Name: pulumi.interpolate`web-server ${server.hex}`,
    },
    ami: server.keepers.apply(keepers => keepers?.amiId),
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_aws as aws
import pulumi_random as random

server = random.RandomId("server",
    keepers={
        "ami_id": ami_id,
    },
    byte_length=8)
server_instance = aws.ec2.Instance("server",
    tags={
        "Name": server.hex.apply(lambda hex: f"web-server {hex}"),
    },
    ami=server.keepers["amiId"])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;
using Random = Pulumi.Random;

return await Deployment.RunAsync(() =>
{
    var server = new Random.RandomId("server", new()
    {
        Keepers =
        {
            { "ami_id", amiId },
        },
        ByteLength = 8,
    });

    var serverInstance = new Aws.Ec2.Instance("server", new()
    {
        Tags =
        {
            { "Name", server.Hex.Apply(hex => $"web-server {hex}") },
        },
        Ami = server.Keepers.Apply(keepers => keepers?.AmiId),
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		server, err := random.NewRandomId(ctx, "server", &random.RandomIdArgs{
			Keepers: pulumi.StringMap{
				"ami_id": pulumi.Any(amiId),
			},
			ByteLength: pulumi.Int(8),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewInstance(ctx, "server", &ec2.InstanceArgs{
			Tags: pulumi.StringMap{
				"Name": server.Hex.ApplyT(func(hex string) (string, error) {
					return fmt.Sprintf("web-server %v", hex), nil
				}).(pulumi.StringOutput),
			},
			Ami: pulumi.String(server.Keepers.ApplyT(func(keepers map[string]string) (*string, error) {
				return &keepers.AmiId, nil
			}).(pulumi.StringPtrOutput)),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
resources:
  server:
    type: random:RandomId
    properties:
      keepers:
        ami_id: ${amiId}
      byteLength: 8
  serverInstance:
    type: aws:ec2:Instance
    name: server
    properties:
      tags:
        Name: web-server ${server.hex}
      ami: ${server.keepers.amiId}
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.random.RandomId;
import com.pulumi.random.RandomIdArgs;
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.ec2.InstanceArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var server = new RandomId("server", RandomIdArgs.builder()
            .keepers(Map.of("ami_id", amiId))
            .byteLength(8)
            .build());

        var serverInstance = new Instance("serverInstance", InstanceArgs.builder()
            .tags(Map.of("Name", server.hex().applyValue(hex -> String.format("web-server %s", hex))))
            .ami(server.keepers().applyValue(keepers -> keepers.amiId()))
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Resource "keepers" are optional. The other arguments to each resource must
*also* remain constant in order to retain a random result.

`keepers` are *not* treated as sensitive attributes; a value used for `keepers` will be displayed in Pulumi UI output as plaintext.