package idf

import "github.com/mchmarny/idf/internal/formatter"

// WithBase64Encoding base64 with URL encodes the input string.
func WithBase64Encoding() func(*IDFormatter) {
	return func(f *IDFormatter) {
		f.formatters = append(f.formatters, &formatter.Base64EncodingFormatter{})
	}
}

// WithDatetime formats the time value and appends it to the id.
func WithDatetime(format, separator string, utc bool) func(*IDFormatter) {
	return func(f *IDFormatter) {
		f.formatters = append(f.formatters, &formatter.DatetimeFormatter{
			BaseFormat: format,
			Separator:  separator,
			UTC:        utc,
		})
	}
}

// WithPadding will pad the input string with the given character to the given length.
func WithPadding(char string, length int) func(*IDFormatter) {
	return func(f *IDFormatter) {
		f.formatters = append(f.formatters, &formatter.PaddingFormatter{
			PadChar: char,
			Length:  length,
		})
	}
}

// WithPrefix normalizes the input string to lowercase and trim spaces.
func WithPrefix(prefix string) func(*IDFormatter) {
	return func(f *IDFormatter) {
		f.formatters = append(f.formatters, &formatter.PrefixFormater{
			Prefix: prefix,
		})
	}
}

// WithSHA256Encoding sha256 encodes the input string.
func WithSHA256Encoding() func(*IDFormatter) {
	return func(f *IDFormatter) {
		f.formatters = append(f.formatters, &formatter.SHA256EncodingFormatter{})
	}
}

// WithPrefix prepends the prefix to the input string.
func WithSuffix(suffix string) func(*IDFormatter) {
	return func(f *IDFormatter) {
		f.formatters = append(f.formatters, &formatter.SuffixFormatter{
			Suffix: suffix,
		})
	}
}
