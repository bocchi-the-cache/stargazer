# StargazeR

English | [简体中文](./README.md)

![logo.png](./doc/logo.png)

## Overview

**StargazeR** is test suites for monitoring network targets, providing HTTP/HTTPS, Ping methods. Provides dynamic
configuration and visualization panels, implemented in Go language.

Project is currently under development. Breaking changes may occur.

## background

It is a common requirement to simulate user access in various places of the world.

- HTTP test, PING test, DNS test
- The test system is important, but not technical complex
- It is rare to see a suite that can quickly establish a distributed dial test system

### Target users

#### Individual developers/geeks:

- Provide a status page to observe the health status of your own VPS/NAS and other computing resources
- Receive an alarm message when the status is abnormal

#### Enterprise

- Provides convenient configuration to monitor targets
- Provide Prometheus Http metric for monitoring infrastructure capture
- Can be easily deployed anywhere as a probe

## Development Log

### v0.2 (Doing)

front end

- [ ] Provide a Web Status Page to check the health status overview
- [ ] provides a configuration page for probing targets

### v0.1 (Done)

- [x] pure backend and api implementation
- [x] Support HTTP, HTTPS, Ping test
- [x] Support configuration file
- [x] provide api to configure detection target
- [x] provide api to get detection target status and logs