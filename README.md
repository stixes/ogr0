# Description

This image container a low interactions ssh honeypot. This is designed for deployment in clusters or other environments where a concentration of honeypots are needed to detect possible intrusions. 

This honeypot works by immitating the handshake and initial login of the ssh protocol, however to valid logins are possible, and connection attepts are logged, along with source ip addresses and username/password combinations.

Based on the work done by [https://github.com/influx6/go0r], this is a Docker wrapped version, with alot of optimization for running securely in a docker/container environment.

Security of the original code is already considerable, however in a container environment it is possible to reduce complexity and surface even further.

Significant changes from the original:

- Removed configurability. Use port mappings and services for configuration.
- Reduced log output. local address is irrelevant in a container setting.
- Run as unprivileged user in an empty container. Just in case.


# Running

To increase security when running in a shared/container environment, no configuration inputs has been provided.

Running a test instance is as simple as:

		docker run --rm -i -p 2222:2222/tcp stixes/ogr0:latest

For production run, a compose example is provided:

		version: "2.4"
		services:
			ogr0:
				image: stixes/ogr0:latest
				restart: always
				ports:
					- 2222:2222/tcp
		