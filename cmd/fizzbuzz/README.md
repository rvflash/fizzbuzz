# FizzBuzz REST API

FizzBuzz REST API is a simple HTTPS service to generate your own FizzBuzz.
It returns the result as JSON.
By default, it's available on https://localhost:4433.

See the options bellow to change the context of execution:

```
Usage of ./fizzbuzz:
  -env string
    	environment name (default "prod")
  -port int
    	service port (default 4433)
```

### Routes

It only exposes two endpoints on GET method:
* A [health check](#health-check) (GET /health) to verify its status.
* A landing page (GET /) to play with it.

The landing page takes the following GET parameters:

| parameters | required? | description                                                      |
|------------|:---------:|------------------------------------------------------------------|
| string1    | mandatory | all multiples of int1 are replaced by this string.               |
| string2    | mandatory | all multiples of int2 are replaced by this string.               |
| int1       | mandatory | first int value to use as multiple.                              |
| int2       | mandatory | second int value to use as multiple.                             |
| limit      |  optional | list the numbers to return from 1 to the given value, default 0. |

With `limit` as optional parameter, if it's omit (or inferior to 1), the result is `null`.

Sample call:
> https://localhost:4433/?string1=fizz&string2=buzz&int1=3&int2=5&limit=100


## Quick start

Ensure to have the latest version of this repository and get its dependencies with dep:

```bash
$ go get -u github.com/rvflash/fizzbuzz
$ cd $GOPATH/src/github.com/rvflash/fizzbuzz
$ dep ensure
```

Use your deployment tool to specify on building, the latest version available of this project (useful in logs).  

```bash
$ cd $GOPATH/src/github.com/rvflash/fizzbuzz/cmd/fizzbuzz
$ go build -ldflags "-X main.buildVersion=v0.0.1"
```

## Systemd unit file (sample)

> Paths must be changed to correspond to your environment.

To simplify its deployment and monitoring, a service file is also available.  
This file contains all required configuration to launch the API via `systemd` as daemon.
It must be copy to the server's systemd folder.

Then:
```bash
$ systemctl enable fizzbuzz.service
```

Once enabled, the service is defined to restart on server's reboot.
You can now start it and follow it when you want.

```bash
$ systemctl start fizzbuzz.service
$ systemctl status fizzbuzz.service
```

If the service goes wrong, `systemctl` will try to restart it after one second.  

It's also easy to follow on live its logs with the following command: 

```bash
$ journalctl -f --user-unit fizzbuzz
```

## Health check

The service exposes a health check API endpoint (GET /health) and so, can be followed by any others systems.
FizzBuzz REST API is very simple, so it only returns the running status of itself.

> If a cache is added to improve the response times, the status of its connections will need to be added.