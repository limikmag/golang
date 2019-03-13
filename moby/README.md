# Moby project (https://www.github.com/moby/moby)

## How to build binaries? (Linux CentOS 7.6)

Prerequisites:
- git

```bash
yum install git
```
- docker CE (go through this link)
 https://docs.docker.com/install/linux/docker-ce/centos/

```bash
$ git clone https://github.com/moby/moby.git
```

We are going to moby/ folder and run:
```bash
$ make
```

After run this command we will have all binaries in directory __moby/bundles/binary-daemon__:

```
containerd              dockerd-dev.md5             rootlesskit
containerd.md5          dockerd-dev.sha256          rootlesskit.md5
containerd.sha256       dockerd-rootless.sh         rootlesskit.sha256
containerd-shim         dockerd-rootless.sh.md5     runc
containerd-shim.md5     dockerd-rootless.sh.sha256  runc.md5
containerd-shim.sha256  docker-init                 runc.sha256
ctr                     docker-init.md5             vpnkit
ctr.md5                 docker-init.sha256          vpnkit.md5
ctr.sha256              docker-proxy                vpnkit.sha256
dockerd                 docker-proxy.md5
dockerd-dev             docker-proxy.sha256

```

1. __containerd__ - containerd is available as a daemon for Linux and Windows. It manages the complete container lifecycle of its host system, from image transfer and storage to container execution and supervision to low-level storage to network attachments and beyond.
   

``` bash
$ ./containerd --help

NAME:
   containerd - 
                    __        _                     __
  _________  ____  / /_____ _(_)___  ___  _________/ /
 / ___/ __ \/ __ \/ __/ __ `/ / __ \/ _ \/ ___/ __  /
/ /__/ /_/ / / / / /_/ /_/ / / / / /  __/ /  / /_/ /
\___/\____/_/ /_/\__/\__,_/_/_/ /_/\___/_/   \__,_/

high performance container runtime


USAGE:
   containerd [global options] command [command options] [arguments...]

VERSION:
   v1.2.4
```

2. __containerd-shim__ - containerd calls containerd-shim that uses runC to run the container. Then, containerd-shim allows the runtime (runC in this case) to exit after it starts the container : This way we can run daemon-less containers because we are not having to have the long running runtime processes for containers.

3. __ctr__

```bash
./ctr --help
NAME:
   ctr - 
        __
  _____/ /______
 / ___/ __/ ___/
/ /__/ /_/ /
\___/\__/_/

containerd CLI


USAGE:
   ctr [global options] command [command options] [arguments...]

VERSION:
   v1.2.4
```


4. __dockerd__ - symlink to docker-dev binary - it is docker deamon which share API server for us.


5. runc

```bash
./runc --help
NAME:
   runc - Open Container Initiative runtime

runc is a command line client for running applications packaged according to
the Open Container Initiative (OCI) format and is a compliant implementation of the
Open Container Initiative specification.

runc integrates well with existing process supervisors to provide a production
container runtime environment for applications. It can be used with your
existing process monitoring tools and the container will be spawned as a
direct child of the process supervisor.

Containers are configured using bundles. A bundle for a container is a directory
that includes a specification file named "config.json" and a root filesystem.
The root filesystem contains the contents of the container.

To start a new instance of a container:

    # runc run [ -b bundle ] <container-id>

Where "<container-id>" is your name for the instance of the container that you
are starting. The name you provide for the container instance must be unique on
your host. Providing the bundle directory using "-b" is optional. The default
value for "bundle" is the current directory.

USAGE:
   runc [global options] command [command options] [arguments...]
   
VERSION:
   1.0.0-rc6+dev
```


## Running dockerd from made binaries

```bash
$ sudo su
$ systemctl stop docker
$ export XDG_RUNTIME_DIR=/var/run - (place where docker.sock will be put)
$ ./dockerd --experimental


WARN[2019-03-11T22:21:22.211562586+01:00] Running experimental build                   
WARN[2019-03-11T22:21:22.211691366+01:00] Running in rootless mode. Cgroups, AppArmor, and CRIU are disabled. 
INFO[2019-03-11T22:21:22.279959292+01:00] parsed scheme: "unix"                         module=grpc
INFO[2019-03-11T22:21:22.280166592+01:00] scheme "unix" not registered, fallback to default scheme  module=grpc
INFO[2019-03-11T22:21:22.280418962+01:00] parsed scheme: "unix"                         module=grpc
INFO[2019-03-11T22:21:22.280551302+01:00] scheme "unix" not registered, fallback to default scheme  module=grpc
INFO[2019-03-11T22:21:22.280824142+01:00] ccResolverWrapper: sending new addresses to cc: [{unix:///var/run/containerd/containerd.sock 0  <nil>}]  module=grpc
INFO[2019-03-11T22:21:22.280919592+01:00] ClientConn switching balancer to "pick_first"  module=grpc
INFO[2019-03-11T22:21:22.281173642+01:00] pickfirstBalancer: HandleSubConnStateChange: 0xc000a10120, CONNECTING  module=grpc
INFO[2019-03-11T22:21:22.282035313+01:00] ccResolverWrapper: sending new addresses to cc: [{unix:///var/run/containerd/containerd.sock 0  <nil>}]  module=grpc
INFO[2019-03-11T22:21:22.282154473+01:00] ClientConn switching balancer to "pick_first"  module=grpc
INFO[2019-03-11T22:21:22.282551973+01:00] pickfirstBalancer: HandleSubConnStateChange: 0xc000ab6010, CONNECTING  module=grpc
INFO[2019-03-11T22:21:22.282619473+01:00] blockingPicker: the picked transport is not ready, loop back to repick  module=grpc
INFO[2019-03-11T22:21:22.283149603+01:00] pickfirstBalancer: HandleSubConnStateChange: 0xc000a10120, READY  module=grpc
INFO[2019-03-11T22:21:22.283599263+01:00] pickfirstBalancer: HandleSubConnStateChange: 0xc000ab6010, READY  module=grpc
INFO[2019-03-11T22:21:22.300801080+01:00] [graphdriver] using prior storage driver: overlay2 
INFO[2019-03-11T22:21:22.305728162+01:00] Loading containers: start.                   
INFO[2019-03-11T22:21:22.453358958+01:00] Default bridge (docker0) is assigned with an IP address 172.17.0.0/16. Daemon option --bip can be used to set a preferred IP address 
INFO[2019-03-11T22:21:22.497245655+01:00] Loading containers: done.                    
INFO[2019-03-11T22:21:22.542159792+01:00] Docker daemon                                 commit=33c3200 graphdriver(s)=overlay2 version=dev
INFO[2019-03-11T22:21:22.542287182+01:00] Daemon has completed initialization          
INFO[2019-03-11T22:21:22.555503567+01:00] API listen on /var/run/docker.sock

```



## Running mobicli

If you run dockerd from previous step, you can run basic cli tool for moby deamon:

```bash
$ mobicli --help
Usage:
  mobicli [command]

Available Commands:
  help        Help about any command
  ps          list all running containers
  start       run container
  stop        stop container with specific name

Flags:
  -h, --help   help for mobicli
  ```

