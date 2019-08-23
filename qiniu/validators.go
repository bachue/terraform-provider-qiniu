package qiniu

import (
	"fmt"
	"net/url"
	"regexp"
)

var validBucketNameRegex *regexp.Regexp

func init() {
	validBucketNameRegex = regexp.MustCompile("^[a-zA-Z0-9\\-_]+$")
}

func validateBucketName(v interface{}, attributeName string) (warns []string, errs []error) {
	bucketName := v.(string)
	if len(bucketName) == 0 {
		errs = append(errs, fmt.Errorf("%q must not be empty", attributeName))
		return
	}
	if len(bucketName) > 63 {
		errs = append(errs, fmt.Errorf("%q must not be longer than 63 characters", attributeName))
		return
	}
	if !validBucketNameRegex.MatchString(bucketName) {
		errs = append(errs, fmt.Errorf("%q must not contain invalid characters", attributeName))
		return
	}
	return
}

func validateRegionID(v interface{}, attributeName string) (warns []string, errs []error) {
	regionId := v.(string)
	switch regionId {
	case "z0", "z1", "z2", "na0", "as0":
		return
	default:
		errs = append(errs, fmt.Errorf("%q is invalid", attributeName))
		return
	}
}

func validatePositiveInt(v interface{}, attributeName string) (warns []string, errs []error) {
	if v.(int) <= 0 {
		errs = append(errs, fmt.Errorf("%q must be positive", attributeName))
	}
	return
}

func validateURL(v interface{}, attributeName string) (warns []string, errs []error) {
	if u, err := url.ParseRequestURI(v.(string)); err != nil {
		errs = append(errs, fmt.Errorf("%q must be valid url", attributeName))
	} else if u.Scheme != "http" && u.Scheme != "https" {
		errs = append(errs, fmt.Errorf("%q should be http or https protocol", attributeName))
	}
	return
}

func validateHost(v interface{}, attributeName string) (warns []string, errs []error) {
	const r = "^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])$"
	if !regexp.MustCompile(r).MatchString(v.(string)) {
		errs = append(errs, fmt.Errorf("%q must be valid host", attributeName))
	}
	return
}
