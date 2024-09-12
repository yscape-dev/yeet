# yeet - cli tool for evaluating performance

`yeet` is a cli tool for better performance evaluation. It is a simple tool that allows you to setup a stable environment for performance evaluation, sync your code to that environment, and then run your benchmarks or profiling tools.

# Usage

```
yeet init # initialize yeet in the current directory
yeet run <command> # run a command in the yeet environment
```

# Commands

## init

## run

* Syncs the code to the yeet environment
* Runs the command in the yeet environment
* Copies the results back to the host


# TODO

* Add support for JMH example
* Add support for prof example
