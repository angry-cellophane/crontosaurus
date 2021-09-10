package main

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var cmd = &cobra.Command{
	Use:     "crontosaurus",
	Example: "crontosaurus '*/15 0 1,15 * 1-5 /usr/bin/find'",
	Short:   "CLI to explain cron expressions",
	Args:    cobra.ExactArgs(1),
	RunE:    runCommand,
}

func main() {
	if e := cmd.Execute(); e != nil {
		os.Exit(1)
	}
}

func runCommand(cmd *cobra.Command, args []string) error {
	var fields []string
	fields, e := splitByFields(args[0])
	if e != nil {
		return e
	}

	expression, command := toCronAndCommand(fields)
	table, e := toHumanReadableExplanation(expression, command)
	if e != nil {
		return e
	}

	fmt.Println(table)
	return nil
}

func toHumanReadableExplanation(expression, command string) (string, error) {
	schedule, e := cron.Parse(expression)
	if e != nil {
		return "", e
	}

	table := toTable(schedule.(*cron.SpecSchedule), command)
	return table, nil
}

func splitByFields(expression string) ([]string, error) {
	fields := strings.Fields(expression)
	if len(fields) != 6 {
		return nil, fmt.Errorf("expected 6 fields in cron expression but got %d in %s\n" +
			"Format: min hour dayOfMonth month dayOfWeek command",
			len(fields), expression)
	}

	return fields, nil
}

func toCronAndCommand(fields []string) (string, string) {
	// the input expression doesn't have seconds but the parser expects it to be
	// adding a fake 0 second to make parser happy
	cronExpression := "0 " + strings.Join(fields[:len(fields)-1], " ")
	command := fields[len(fields)-1]
	return cronExpression, command
}

/*
	first 14 chars on field name + padding, then a list of value separated by space
 */
func toTable(schedule *cron.SpecSchedule, command string) string {
	return fmt.Sprintf(
		"minute        %s\n" +
		"hour          %s\n" +
		"day of month  %s\n" +
		"month         %s\n" +
		"day of week   %s\n" +
		"command       %s\n",
		getValues(schedule.Minute, 59),
		getValues(schedule.Hour, 23),
		getValues(schedule.Dom, 31),
		getValues(schedule.Month,12),
		getValues(schedule.Dow, 6),
		strings.TrimSpace(command),
	)
}


/*
	SpecSchedule stores fields in bitsets, uint64.
	every bit represents a value, up to 64 distinct value.
	minute uses only 60 bits
	hour uses 24 bits
	etc
 */
func getValues(value uint64, bitLimit int) string {
	var mask uint64 = 1

	s := strings.Builder{}
	s.Grow(bitLimit * 2)
	isFirstChar := true
	for i:=0; i<= bitLimit; i++ {
		masked := value & mask
		value = value >> 1
		if masked == 0 {
			continue
		}

		if isFirstChar {
			isFirstChar = false
		} else {
			s.WriteString(" ")
		}

		s.WriteString(fmt.Sprintf("%d", i))
	}

	return s.String()
}