# Dev Test Host Container

This directory contains the necessary files to build and run a development test host container using Docker. The container is based on Debian Bookworm and includes SSH server functionality.

## Prerequisites

* Copy or link your SSH key to `dev/testhost/authorized_keys`
   ```
   ln -s ~/.ssh/id_rsa.pub dev/testhost/authorized_keys
   # or
   cp ~/.ssh/id_rsa.pub dev/testhost/authorized_keys
   ```


## Usage

To use this development test host container, follow these steps:

1. Build the Docker image:
   ```
   make build
   ```

2. Run the container:
   ```
   make run
   ```

3. Connect to the container via SSH:
   ```
   make connect
   ```

4. To stop and remove the container:
   ```
   make stop
   ```

These commands are defined in the Makefile and provide an easy way to manage the container lifecycle.

