# Getting started

- [Getting started](#getting-started)
  - [1. First run](#1-first-run)
    - [1.1. Use executable file](#11-use-executable-file)
    - [1.2. Use Docker image](#12-use-docker-image)
  - [2. Flag](#2-flag)
  - [3. Configuration](#3-configuration)
    - [3.1. Global configuration](#31-global-configuration)
  - [3.2. Etcd configuration](#32-etcd-configuration)

## 1. First run

* Clone the repository & cd into it:

```bash
$ git clone github.com/ntk148v/faythe
$ cd faythe
```

* [Prepare an Etcd cluster](https://github.com/etcd-io/etcd). This is a must-have component.

### 1.1. Use executable file

```bash
# Build it
$ go build -mod vendor -o /path/to/executable/faythe cmd/faythe/main.go
# Create a config file / Simply make a clone
$ cp examples/faythe.yml /path/to/config/dir/config.yml
# Modify /path/to/config/dir/config.yml
# Run it
$ /path/to/executable/faythe --config.file /path/to/config/dir/config.yml
```

### 1.2. Use Docker image

- Build Docker image (use git tag/git branch as Docker image tag).

```bash
$ make build
```

- Run container from built image.

```bash
$ make run
```

- For more details & options please check [Makefile](../Makefile).

### 1.3. Use docker-compose

- Update [sample config file](../examples/faythe.yml).

- Run docker-compose, it will build faythe image, start faythe and etcd container.

```
$ docker-compose up -d
```

## 2. Flag

```bash
usage: main [<flags>]

The Faythe server

Flags:
  -h, --help                Show context-sensitive help (also try --help-long and --help-man).
      --config.file="/etc/faythe/config.yml"
                            Faythe configuration file path.
      --listen-address="0.0.0.0:8600"
                            Address to listen on for API.
      --external-url=<URL>  The URL under which Faythe is externally reachable.
      --log.level=info      Only log messages with the given severity or above. One of: [debug, info, warn, error]
      --log.format=logfmt   Output format of log messages. One of: [logfmt, json]
```

## 3. Configuration

For more information, please check [config/module](../config).

```yaml
# Global configuration
global:
  # Example:
  # "www.example.com"
  # "([a-z]+).domain.com"
  # The regex has to follow Golang regex convention
  remote_host_pattern: "192.168.(128|129).*"
  basic_auth:
    username: "admin"
    password: "notverysecurepassword"

# Etcd configuration
etcd:
  endpoints: ["192.168.1.2:2379", "192.168.1.3:2379", "192.168.1.4:2379"]
  username: "etcdadmin"
  password: "etcdsuperpassword"
  auto_sync_interval: 30s
  dial_timeout: 2s
  dial_keep_alive_time: 2s
  dial_keep_alive_timeout: 6s
```

### 3.1. Global configuration

- `remote_host_pattern`: define an optional regexp host pattern to be matched. Faythe accepts requests from every hosts by default, no restrict. Please check Golang regex cheatsheet for more details.

  ```yaml
  # Example
  # Allow requests from every hosts.
  remote_host_pattern: ".*"
  # Allow requests from host whose ip address starts with 192.168.128 or 192.168.129
  remote_host_pattern: "192.168.(128|129).*"
  ```

- `basic_auth`: HTTP basic authentication with `username` & `password`. If you don't want to enable basic authentication, just remove this section.

## 3.2. Etcd configuration

Etcd is a Faythe's heart. You can figure out etcd configuration [here](https://github.com/etcd-io/etcd/blob/master/clientv3/config.go). The configuration in Faythe is just a mimic.
