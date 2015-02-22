# MultiDockerCli

MultiDockerCli is just a way to user docker on several hosts (you know if you still don't  use docker swarm...)

Development in progress

## Prerequisites
 - golang >= 1.4.1

## Installation
```
 $ git clone https://github.com/fsamin/MultiDockerCli.git
 $ cd MultiDockerCli`
 $ make && sudo make install
 ```

## How to use
###Edit a `multidocker.json` file.
```
{
  "nodes" : [
    {
      "alias" : "boot2docker",
      "host" : "tcp://192.168.59.103:2376",
      "cert" : {
        "caFile":   "/Users/fsamin/.boot2docker/certs/boot2docker-vm/ca.pem",
        "certFile": "/Users/fsamin/.boot2docker/certs/boot2docker-vm/cert.pem",
        "keyFile":  "/Users/fsamin/.boot2docker/certs/boot2docker-vm/key.pem"
      }
    },
    {
      "alias" : "host1",
      "host" : "tcp://192.168.0.1:2376",
    }
  ]
}
```
###Run 
 `$ multidocker ps`

```
2015/02/22 18:12:54 Connecting to docker node boot2docker::tcp://192.168.59.103:2376 (version 1.5.0)
2015/02/22 18:12:54 Connecting to docker node host1::tcp://192.168.0.1:2376 (version 1.5.0)
|--------------------------------------------------------------------------------------------------------------------------------------------------|
| NODE        | HOST                      | CONTAINER ID | IMAGE          | COMMAND                   | CREATED                       | STATUS     |
|--------------------------------------------------------------------------------------------------------------------------------------------------|
| boot2docker | tcp://192.168.59.103:2376 | 2684bc2ec255 | jenkins:latest | /usr/local/bin/jenkins.sh | 2015-02-22 11:52:06 +0100 CET | Up 6 hours |
| host1       | tcp://192.168.0.1:2376    | 6a54fdf741a4 | jenkins:latest | /usr/local/bin/jenkins.sh | 2015-02-22 11:50:37 +0100 CET | Up 6 hours |
|--------------------------------------------------------------------------------------------------------------------------------------------------|
```

 
