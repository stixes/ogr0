# Description

Based on the work done by [https://github.com/influx6/go0r], this is a Docker wrapped version, with alot of optimization for running securely in a docker/container environment.

Significant changes from the original:

- Removed configurability. Use port mappings and services for configuration.
- Reduced log output. local address is irrelevant in a container setting.
- Run as unprivileged user in an empty container. Just in case.


