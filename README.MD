# Delivery Code Challenge

**This Service provides a input for sending delivery guys to their job**

At the moment, once the service receives a destination for a delivery job , at first, it tries to send a first free
delivery
guy at main controller.

if any delivery guy doesn't exist at there . main controller sends the nearest free delivery guy

**moving to every unit for a delivery guy takes 1 second**

### Getting started

In order to run the service follow the below instructions.

1- Install docker ([installation manual](https://docs.docker.com/engine/install/))

2-
```shell
sudo docker build -t delivery .
```
3-
```shell
sudo docker run --rm -it delivery /bin/bash

```

After that you are in docker container 

```shell
/app/minidelivery
```

Now you can enter a destination in blow format
```shell
-1 5
```
And the delivery guy goes for doing his job .
you can send another one.

**for testing, you can use this command at docker container**

```shell
go test ./... 
```

Good luck :)
