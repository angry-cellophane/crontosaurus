# What's this? 
crontosaurus

# Why?
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