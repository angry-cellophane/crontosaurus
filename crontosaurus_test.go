package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Test_toCronAndCommand(t *testing.T) {
	type args struct {
		fields []string
	}
	tests := []struct {
		name  string
		args  args
		expression  string
		command string
	}{
		{"*/15 0 1,15 * 1-5 /sbin/shutdown", args{[]string{"*/15", "0", "1,15", "*", "1-5", "/sbin/shutdown"}},
			"0 */15 0 1,15 * 1-5", "/sbin/shutdown"},
		{"* * * * * /sbin/shutdown", args{[]string{"*", "*", "*", "*", "*", "/sbin/shutdown"}},
			"0 * * * * *", "/sbin/shutdown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := toCronAndCommand(tt.args.fields)
			if got != tt.expression {
				t.Errorf("toCronAndCommand() got = %v, want %v", got, tt.expression)
			}
			if got1 != tt.command {
				t.Errorf("toCronAndCommand() got1 = %v, want %v", got1, tt.command)
			}
		})
	}
}

func Test_splitByFields(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name   string
		args   args
		fields []string
	}{
		{"valid, separated by spaces", args{"*/15 0 1,15 * 1-5 /sbin/shutdown"},
			[]string{"*/15", "0",  "1,15", "*", "1-5", "/sbin/shutdown"}},
		{"valid, separated by tabs", args{"*	0	1	*	1-5	/sbin/shutdown"},
			[]string{"*", "0",  "1", "*", "1-5", "/sbin/shutdown"}},
		{"valid, mixed tabs and spaces", args{"*	0 1	*	1-5 /sbin/shutdown"},
			[]string{"*", "0",  "1", "*", "1-5", "/sbin/shutdown"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitByFields(tt.args.expression)
			if !reflect.DeepEqual(got, tt.fields) {
				t.Errorf("splitByFields() got = %v, want %v", got, tt.fields)
			}
			if got1 != nil  {
				t.Errorf("toCronAndCommand() unexpected error got1 = %v", got1)
			}
		})
	}
}

func Test_splitByFields_errors(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name   string
		args   args
	}{
		{"len(columns) < 6", args{"*/15 0 * 1-5 /sbin/shutdown"}},
		{"len(columns) > 6", args{"*/15 0 * 1-5 12 1 2 /sbin/shutdown"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := splitByFields(tt.args.expression)
			if got1 == nil  {
				t.Errorf("toCronAndCommand() expected error but no error returned")
			}
		})
	}
}

func Test_toHumanReadableExplanation(t *testing.T) {
	type args struct {
		expression string
		command string
	}
	tests := []struct {
		name   string
		args   args
		output string
	}{
		{"all week days", args{"0 */15 0 4-5 3,4 *", "/command"},
			toTestTable("0 15 30 45", "0", "4 5", "3 4", "0 1 2 3 4 5 6", "/command")},
		{"all minutes", args{"0 * 1 1 1 1", "/command"},
			toTestTable(createRange(0, 59), "1", "1", "1", "1", "/command")},
		{"all hours", args{"0 1 * 1 1 1", "/command"},
			toTestTable("1", createRange(0, 23), "1", "1", "1", "/command")},
		{"all doms", args{"0 1 1 * 1 1", "/command"},
			toTestTable("1", "1", createRange(1, 31), "1", "1", "/command")},
		{"all months", args{"0 1 1 1 * 1", "/command"},
			toTestTable("1", "1", "1", createRange(1, 12), "1", "/command")},
		{"all dows", args{"0 1 1 1 1 *", "/command"},
			toTestTable("1", "1", "1", "1", createRange(0, 6), "/command")},
		{"sample example", args{"0 */15 0 1,15 * 1-5", "/usr/bin/find"},
			toTestTable("0 15 30 45", "0", "1 15", createRange(1, 12), createRange(1,5), "/usr/bin/find")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := toHumanReadableExplanation(tt.args.expression, tt.args.command)
			if got != tt.output {
				t.Errorf("toHumanReadableExplanation() got = \n%v, want \n%v", got, tt.output)
			}
			if got1 != nil  {
				t.Errorf("toHumanReadableExplanation() expected not to return error but got %v", got1)
			}
		})
	}
}

func toTestTable(minute, hour, dom, month, dow, command string) string {
	return fmt.Sprintf(
		"minute        %s\n" +
			"hour          %s\n" +
			"day of month  %s\n" +
			"month         %s\n" +
			"day of week   %s\n" +
			"command       %s\n",
			minute,
			hour,
			dom,
			month,
			dow,
			command,
	)
}

func createRange(from int, to int) string {
	s := strings.Builder{}
	s.Grow(to - from)
	s.WriteString(fmt.Sprintf("%d", from))
	for i:=from+1; i<=to; i++ {
		s.WriteString(" ")
		s.WriteString(fmt.Sprintf("%d", i))
	}

	return s.String()
}