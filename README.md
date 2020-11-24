## What is it

Jira API integrations for managing issues and worklogs through a cli

## Setup

The jira api requires a api token. Generate one by following the instructions
here: https://bit.ly/39jQgMa

Store the resulting api in a config file at `~/.config/jirago.yaml` with the
following format:

```yaml
username: {your jira username}
password: {newly generates api key}
```

## Install

In the root of the project, Build the tool 

```
go build .
```

[More info on building go binaries for different targets.](https://do.co/3fvYLF7)

Use as is 

```
./jirago -h
```

Or Move the binary to a location in your PATH

```
mv ./jirago /usr/local/bin/
```

## Commands

#### _jirago transition_

WIP

#### _jirago time_

*args*

1. issue id
1. duration
1. optional natural language date

*examples*

```
# add time to the  ticket sarting now

jirago time IOT-6062 90m

# add time to the ticket on the date and time quoted

jirago time IOT-6062 90m "tuesday the 17th at 2:00 pm"
jirago time IOT-6062 90m "last tuesday at 2:00 pm"
jirago time IOT-5999 6h "Nov 6th 2020 at 2:00 pm"
```
