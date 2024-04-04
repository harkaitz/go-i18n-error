GO I18N ERROR
=============

Go library for internationalizable (i18n) errors.

## Go documentation

    package i18ne // import "github.com/harkaitz/go-i18n-error"
    
    func CheckError(t *testing.T, e error, code string) error
    func CheckErrors(t *testing.T, el []error, code string) error
    func InvalidArgumentError() (err error)
    func JoinErrors(errl []error) (err error)
    func UnimplementedError() (err error)

## Collaborating

For making bug reports, feature requests and donations visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
