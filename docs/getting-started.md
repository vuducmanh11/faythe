# Getting started

- [Getting started](#getting-started)
  - [1. Installation](#1-installation)
    - [1.1. Use executable file](#11-use-executable-file)
    - [1.2. Use Docker image](#12-use-docker-image)
  - [2. Flag](#2-flag)
  - [3. Configuration](#3-configuration)
    - [3.1. Faythe Server Config](#31-faythe-server-config)
    - [3.2. Faythe OpenStack Config](#32-faythe-openstack-config)
    - [3.3. Faythe StackStorm Config](#33-faythe-stackstorm-config)
  - [4. Endpoints](#4-endpoints)
    - [4.1. Basic endpoints](#41-basic-endpoints)
    - [4.2. OpenStack endpoints](#42-openstack-endpoints)
    - [4.3. StackStorm endpoints](#43-stackstorm-endpoints)

## 1. Installation

### 1.1. Use executable file

```bash
# Modify etc/config.yml file
$ vim etc/config.yml
# Move config file to config directory
$ cp etc/config.yml /path/to/config/dir
# Build it
$ go build -mod vendor -o /path/to/executable/faythe
# Run it
$ /path/to/executable/faythe -conf /path/to/config/dir/config.yml
```

### 1.2. Use Docker image

- Build Docker image (use git tag/git branch as Docker image tag):

```bash
$ make build
```

- Run container from built image (For more details & options please check [Makefile](./Makefile)):

```
$ make run
```

## 2. Flag

```bash
Usage:
-conf string
  config file path. (default "/etc/faythe/config.yml")
-listen-addr string
  server listen address. (default ":8600")
```

## 3. Configuration

For more options, please check [config module](../config/config.go).

### 3.1. Faythe Server Config

`server_config` section - values that are used to config Faythe HTTP Server.

```yaml
# Sample
server_config:
  # Example:
  # "www.example.com"
  # "([a-z]+).domain.com"
  remote_host_pattern: "192.168.(128|129).*"
  basic_auth:
    username: "admin"
    password: "notverysecurepassword"
  log_dir: "/tmp/faythe-logs"
```

- `remote_host_pattern`: define an optional regexp host pattern to be matched. Faythe accepts requests from every hosts by default, no restrict. Please check Golang regex cheatsheet for more details.

  ```yaml
  # Example
  # Allow requests from every hosts.
  remote_host_pattern: ".*"
  # Allow requests from host whose ip address starts with 192.168.128 or 192.168.129
  remote_host_pattern: "192.168.(128|129).*"
  ```

- `basic_auth`: HTTP basic authentication with `username` & `password`. If you don't want to enable basic authentication, just remove this section.
- `log_dir`: logging directory, by default it is `/var/log/faythe`.

### 3.2. Faythe OpenStack Config

```yaml
# OpenStackConfiguration.
openstack_configs:
  openstack-1f:
    region_name: "RegionOne"
    domain_name: "Default"
    auth_url: "http://openstackhost1:5000"
    username: "admin"
    password: "password"
    project_name: "tenantName"

  openstack-2f:
    region_name: "RegionOne"
    domain_name: "Default"
    auth_url: "http://openstackhost2:5000"
    username: "admin"
    password: "password"
    project_name: "tenantName"
```

`openstack_configs` is an map of multiple OpenStack configuration. Each OpenStack configuration contains:

- `auth_url`: specifies the HTTP endpoint that is required to work with the Identity API of the appropriate version. While it's ultimately needed by all of the identity services, it will often be populated by a provider-level function.
- `region_name`: region name.
- `username` & `user_id`: Username is required if using Identity V2 API. Consult with your provider's control panel to discover your account's username. In Identity V3, either UserID or a combination of Username & DomainID or DomainName are needed.
- `domain_name` & `domain_id`: At most one of DomainID and DomainName must be provided if using Username with Identity V3. Otherwise, either are optional.
- `project_name` & `project_id`: The ProjectID and ProjectName fields are optional for the Identity V2 API. The same fields are known as project_id and project_name in the Identity V3 API, but are collected as ProjectID and ProjectName here in both cases. Some providers allow you to specify a ProjectName instead of the ProjectId. Some require both. Your provider's authentication policies will determine how these fields influence authentication.
- Note that, the name of OpenStack configuration block is customizable. It means you don't have to name these block as `openstack-1f` or `openstack-2f`.

### 3.3. Faythe StackStorm Config

```yaml
# StackStormConfiguration
stackstorm_configs:
  stackstorm-1f:
    host: "stackstormhost"
    api_key: "fakestackstomrapikey"
```

`stackstorm_configs` is a map of StackStorm configuration. StackStorm configuration contains:

- `host`: The ip address or hostname of StackStorm host.
- `api_key`: StackStorm API key, please check [StackStorm API for more details](https://api.stackstorm.com/).

## 4. Endpoints

### 4.1. Basic endpoints

- `/`: the home endpoint, returns a welcome message only.
- `/healthz`: returns Faythe HTTP server uptime.
- Note that, the name of StackStorm configuration block is customizable. It means you don't have to name these block as `stackstorm-1f`.

### 4.2. OpenStack endpoints

- `/openstack/autoscaling/{ops-name}`: receives alerts Prometheus Alertmanager, processes and handles scale action. `ops-name` variable is the name of OpenStack configuration block in config file. **It has to be match each other**.

* Only POST request with Prometheus Alert body is supported.

### 4.3. StackStorm endpoints

- `/stackstorm/{st-name}/{st-rule}`: handles POST request and forwards this one to StackStorm host.
- `/stackstorm/alertmanager/{st-name}/{st-rule}`: handles POST request with Prometheus Alert body, processes & forwards request to StackStorm host.
- **Note that the `st-name` has to be match with StackStorm configuration block name in config file**. You have to define a rule with name `st-rule` in StackStorm side before.
