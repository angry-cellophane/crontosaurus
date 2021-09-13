# What's this? 

crontosaurus

## Why?
Write a command line application or script which parses a cron string and expands each field
to show the times at which it will run. You may use whichever language you feel most
comfortable with.

You should only consider the standard cron format with five time fields (minute, hour, day of
month, month, and day of week) plus a command, and you do not need to handle the special
time strings such as `"@yearly"`. 

The input will be on a single line.
The cron string will be passed to your application as a single argument.
`~$ your-program "*/15 0 1,15 * 1-5 /usr/bin/find"`

The output should be formatted as a table with the field name taking the first 14 columns and
the times as a space-separated list following it.

For example, the following input argument:

`*/15 0 1,15 * 1-5 /usr/bin/find`

Should yield the following output:

```
minute 0 15 30 45
hour 0
day of month 1 15
month 1 2 3 4 5 6 7 8 9 10 11 12
day of week 1 2 3 4 5
command /usr/bin/find
```

README and instructions for how to run your project in a clean OS X/Linux
environment.

## Getting Binaries

There's a github action job that uploads binaries on every build.

The job creates binaries for linux and mac.

The best way to get the binaries is to check the latest build's artifacts 
or download them directly from [this link](https://github.com/angry-cellophane/crontosaurus/actions/runs/1222762277).

## Usage

`crontosaurus <<cron expression>>`

where `<<cron expression>>` is a cron expression.

The accepted format of cron expressions: 
1. 6 columns separated by spaces or tabs: minutes, hours, days of the month, months, days of the week, command.
2. Format of every column [described here](https://github.com/robfig/cron/blob/master/README.md)

## Building locally

You'll need `golang 1.16+`.

1. Run `go test` in the project's directory to run tests.
2. Run `go build` in the project's directory. 
It will create a file - `crontosaurus` - in the same directory.
3. Run `chmod +x ./crontosaurus` to make it executable
4. Run `./crontosaurus --help` to see its full power.

## How it works?

It basically uses the [robfig/cron](https://github.com/robfig/cron) lib to parse cron expressions and a bit of formatting of the result in console.

The only trick is that this `CLI` doesn't accept seconds, so they are added automatically (`0` value).

## How to make a change?

You don't want to. This is a test project.

## Wait, you didn't write the parser?

Yeah, it wasn't the goal of this exercise. The goal was to write a cli.

## Alright, but I still want to see a parser
Alright, it's [here](https://github.com/angry-cellophane/cron-cli) but it's in java.
