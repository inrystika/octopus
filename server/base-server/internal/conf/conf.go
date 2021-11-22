package conf

import (
	"regexp"
	"server/common/errors"
	"server/common/utils"
	"strings"
)

const (
	ConfigKeyTypeInput    string = "input"
	ConfigKeyTypeRadio    string = "radio"
	ConfigKeyTypeCheckBox string = "checkbox"
)

func (k *ConfigKey) ValidateValue(v string) error {
	if k.Required && v == "" {
		return errors.Errorf(nil, errors.ErrorConfigValueValidateFailed)
	}

	if (k.Type == ConfigKeyTypeRadio || k.Type == ConfigKeyTypeCheckBox) && v != "" &&
		!utils.StringSliceContainsValue(strings.Split(k.Options, ","), v) {
		return errors.Errorf(nil, errors.ErrorConfigValueValidateFailed)
	}

	if k.Regexp != "" {
		matched, err := regexp.MatchString(k.Regexp, v)
		if err != nil || !matched {
			return errors.Errorf(nil, errors.ErrorConfigValueValidateFailed)
		}
	}

	return nil
}
