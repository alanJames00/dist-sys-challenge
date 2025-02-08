## Globally Unique ID generation

GUID(Globally Unique ID) generation is the process of creating unique identifiers across
multiple nodes in a distributed system without a centralized coordination

### Some Standard Methods

1. Snowflake from twitter
2. ObjectID from mongodb: `4 byte timestamp from unix epoch` + `5 byte random value` + `3 byte incrementing counter` => `12 bytes`

### My custom method

`8 byte timestamp nanosecs from unix epoch` + `8 byte cryptographically secure random number`
