using System.Collections.Generic;
using System.Threading.Tasks;

using Pulumi;
using Pulumi.Random;

class Program
{
    static Task<int> Main()
    {
        return Deployment.RunAsync(() => {

            var pet = new RandomPet("bucket", new RandomPetArgs
            {
            });
            return new Dictionary<string, object>
            {
                { "pet", pet.Id },
            };
        });
    }
}
