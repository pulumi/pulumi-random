You can import external passwords into your Pulumi programs as follows:

```sh<break>
$ import random:index/randomPassword:RandomPassword newPassword supersecret
<break>```

This command will encode the `supersecret` token in Pulumi state and generate a code suggestion to
include a new RandomPassword resource in your Pulumi program. Include the suggested code and do a
`pulumi up`. Your secret password is now securely stored in Pulumi, and you can reference it in your
Pulumi program as `newPassword.result`.
