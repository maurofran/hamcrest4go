package matcher

import (
	"fmt"
	"reflect"
	"strings"
)

// StringWriter is the interface implemented
type StringWriter interface {
	WriteString(string) (int, error)
	String() string
}

type nullWriter struct {
}

func (nullWriter) WriteString(string) (int, error) {
	return 0, nil
}

func (nullWriter) String() string {
	return ""
}

// SelfDescribing is implemented by matchers who need to describe themselves.
type SelfDescribing interface {
	DescribeTo(description Description)
}

func toString(value SelfDescribing) string {
	description := StringDescription()
	description.AppendDescriptionOf(value)
	return description.String()
}

// Description is the representation of a matcher error description.
type Description struct {
	writer StringWriter
	err    error
}

// NewDescription creates a new Description instance from provided StringWriter instance.
// If the StringWriter instance is null, it creates a NullDescription instance.
func NewDescription(writer StringWriter) Description {
	if writer == nil {
		writer = nullWriter{}
	}
	return Description{writer: writer}
}

// NullDescription gets a new Description instance that doesn't write anything.
func NullDescription() Description {
	return NewDescription(nil)
}

// StringDescription gets a new Description instance that write to a string.
func StringDescription() Description {
	writer := new(strings.Builder)
	return NewDescription(writer)
}

func (d Description) append(text string) {
	if d.err == nil {
		_, err := d.writer.WriteString(text)
		d.err = err
	}
}

// Err returns the first error encountered while writing description.
func (d Description) Err() error {
	return d.err
}

// AppendText appends some plain text to the description.
func (d Description) AppendText(text string) {
	d.append(text)
}

// AppendDescriptionOf appends the description of a SelfDescribing value to this description.
func (d Description) AppendDescriptionOf(value SelfDescribing) {
	value.DescribeTo(d)
}

// AppendValue appends an arbitrary value to the description.
func (d Description) AppendValue(value interface{}) {
	if value == nil {
		d.append("nil")
	} else if str, ok := value.(string); ok {
		d.append(toGoSyntax(str))
	} else if r, ok := value.(rune); ok {
		d.append(toGoSyntax(string([]rune{r})))
	} else if reflect.TypeOf(value).Kind() == reflect.Slice || reflect.TypeOf(value).Kind() == reflect.Array {
		slice := reflect.ValueOf(value)
		d.appendSlice("[", ", ", "]", slice)
	} else {
		d.append("<")
		d.append(descriptionOf(value))
		d.append(">")
	}
}

// AppendValueList appends a list of values to the description.
func (d Description) AppendValueList(start, separator, end string, values ...interface{}) {
	d.appendSlice(start, separator, end, reflect.ValueOf(values))
}

func (d Description) appendSlice(start, separator, end string, values reflect.Value) {
	length := values.Len()
	d.append(start)
	for i := 0; i < length; i++ {
		if i > 0 {
			d.append(separator)
		}
		item := values.Index(i).Interface()
		if sd, ok := item.(SelfDescribing); ok {
			d.AppendDescriptionOf(sd)
		} else {
			d.AppendValue(item)
		}
	}
	d.append(end)
}

func (d Description) String() string {
	if d.err != nil {
		return ""
	}
	return d.writer.String()
}

func descriptionOf(value interface{}) string {
	return fmt.Sprintf("%s", value)
}

func toGoSyntax(value string) string {
	value = strings.NewReplacer("\\", "\\\\", "\"", "\\\"", "\n", "\\n", "\r", "\\r", "\t", "\\t").Replace(value)
	return fmt.Sprintf("\"%s\"", value)
}
