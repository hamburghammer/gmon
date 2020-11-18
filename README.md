# gmon
[![Build Status](https://cloud.drone.io/api/badges/hamburghammer/gmon/status.svg)](https://cloud.drone.io/hamburghammer/gmon)

Analyse data from [gsave](https://github.com/hamburghammer/gsave) and send a notification to [Gotify](https://gotify.net).

## Configuration
There are tow necessary configuration files for this service to work.

- config.toml
- rules.toml

It will look for this files inside the current directory. To specify another path use the `--config` and `--rules` arguments with the path and the file name.

### config.toml
Example configuration:
```
# The interval in which it should check for new data.
interval = 1

# Configuration for the gsave endpoint.
[stats]
endpoint = "http://localhost:8080"
hostname = "foobar" # The hostname of the host you want to monitor.
token = "foo"

# Configuration for the gotify notification endpoint.
[gotify]
endpoint = "http://localhost:80"
token = "AzCkehMSkHFlphf"
```

### rules.toml
Example configuration:
```
[[CPU]]
Name = "Unexpected CPU usage"
Description = "More than 50% of CPU utilization"
Compare = ">"
Warning = 50.0
Alert = 100.0
Deactivated = false

[[Disk]]
Name = "Unexpected disk usage"
Description = "More than 50 GB of the disk are used"
Compare = ">"
Warning = 50000
Alert = 70000
Deactivated = true

[[RAM]]
Name = "Unexpected RAM usage"
Description = "More than 5 GB of the RAM are in use"
Compare = ">"
Warning = 5000
Alert = 7000
Deactivated = false
```
