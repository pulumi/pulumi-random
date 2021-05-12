from pulumi import export
import pulumi_random as random

random_id = random.RandomId("id", byte_length=4)
export("id_hex", random_id.hex)