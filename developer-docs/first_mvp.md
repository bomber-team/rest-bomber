### Definitions

#### Schema

In scheme we have three section:

* headers - common http headers, which can be set up from user 
or generated from generators by indicating fields

* body - common http body, parameters in body can be any deep and parameters
 set up from user or generated
 
* request-params - request params from route, like as `v1/route/?param1=?`


For example:
```json
{
  "headers": {
    "method": "GET",
    "authority": "bla"
  },
  "body": {
    "Name": "WordGenerator",
    "Password": "PasswordGenerator",
    "Nas-Ip": "IpGenerator",
    "body": {
      "test": "Test"
    }
  },
  "request-params": {
    "Test": "test"
  }
}
```

#### Script

In script we have different parameters for attack:

* server-address - address of server where we will send packets
* buffer-size - how many packets we generated before attack
* time-between-attacks - how many we wait before each attack
* amount-attacks - if we have 100 attacks we will send 100*100 packets
* log-level - level of system logging

For example:
```json
{
"server-address": "ip:port/v1/route",
"buffer-size": 100,
"time-between-attacks": 1000,
"amount_attacks": 100,
"log-level": "debug"
}
```

#### Metrics

Metrics will be sent on route with specified attack's id

Format of metrics:
```json
{
 "times": [123,213,234],
 "date": "24.05.16 18.30"
 }
``` 
can be changed

#### Receiving a signal from server

Each module has a mini server to receive script and schema for attack

There are three states:

1) Module is ready for new attack
2) Script and scheme received
3) Module is ready to start attack


 