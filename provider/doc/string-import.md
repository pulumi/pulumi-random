You can import external strings into your Pulumi programs as RandomString resources as follows:

```bash<break><break>
import random:index/randomString:RandomString newString myspecialdata
<break><break>```

This command will encode the `myspecialdata` token in Pulumi state and generate a code suggestion to
include a new RandomString resource in your Pulumi program. Include the suggested code and do a
`pulumi up`. Your data is now stored in Pulumi, and you can reference it in your Pulumi program as
`newString.result`.

If the data needs to be stored securily as a secret, consider using the RandomPassword resource
instead.

