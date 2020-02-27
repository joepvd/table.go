package table

// Options contains the runtime configuration
type Options struct {
	Style string `short:"s" long:"style" default:"md"`
	FS    string `short:"F" long:"field-seperator" default:"[ \t]+"`
}
