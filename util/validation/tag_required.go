package validation

import (
	"EduCRM/util/logrus_log"
	"errors"
	"reflect"
	"strconv"
)

var (
	requiredTag   = "required"
	lengthMinTag  = "lenMin"
	lengthMaxTag  = "lenMax"
	amountMinTag  = "amountMin"
	amountMaxTag  = "amountMax"
	regexTag      = "regex"
	loginRegex    = "login"
	emailRegex    = "email"
	numberRegex   = "number"
	passwordRegex = "password"
	hexColorRegex = "hexColor"
	phoneRegex    = "phone"
)

func ValidationStructTag(logger *logrus_log.Logger, model interface{}) error {
	structType := reflect.TypeOf(model)
	structValue := reflect.ValueOf(model)
	for i := 0; i < structType.NumField(); i++ {
		fieldName := structType.Field(i).Name
		fieldValue := reflect.Indirect(structValue).FieldByName(fieldName)
		requiredTagValue, requiredTagExist := structType.Field(i).Tag.Lookup(requiredTag)
		if requiredTagExist {
			err := TagRequired(requiredTagValue, fieldValue.String(),
				fieldName, logger)
			if err != nil {
				logger.Error(err)
				return err
			}
		}
		lenMinTagValue, lenMinTagExist := structType.Field(i).Tag.Lookup(
			lengthMinTag)
		lenMaxTagValue, lenMaxTagExist := structType.Field(i).Tag.Lookup(
			lengthMaxTag)
		if lenMinTagExist && lenMaxTagExist {
			err := LengthRequired(fieldValue.String(), fieldName,
				lenMinTagValue, lenMaxTagValue, logger)
			if err != nil {
				logger.Error(err)
				return err
			}
		}
		amountMinTagValue, amountMinTagExist := structType.Field(i).Tag.
			Lookup(amountMinTag)
		amountMaxTagValue, amountMaxTagExist := structType.Field(i).Tag.
			Lookup(amountMaxTag)
		if amountMinTagExist && amountMaxTagExist {
			err := AmountRequired(int(fieldValue.Int()), fieldName,
				amountMinTagValue, amountMaxTagValue, logger)
			if err != nil {
				logger.Error(err)
				return err
			}
		}
		regexValue, regexExist := structType.Field(i).Tag.Lookup(regexTag)
		if regexExist {
			err := RegexRequired(regexValue, fieldValue, fieldName, logger)
			if err != nil {
				logger.Error(err)
				return err
			}
		}
	}
	return nil
}

func RegexRequired(regexValue string, fieldValue reflect.Value,
	fieldName string, logger *logrus_log.Logger) error {
	switch regexValue {
	case loginRegex:
		err := ValidationLogin(fieldValue.String(), fieldName)
		if err != nil {
			logger.Error(err)
			return err
		}
	case emailRegex:
		err := ValidationEmail(fieldValue.String(), fieldName)
		if err != nil {
			logger.Error(err)
			return err
		}
	case numberRegex:
		err := ValidatorNumber(strconv.FormatInt(fieldValue.Int(), 10), fieldName)
		if err != nil {
			logger.Error(err)
			return err
		}
	case passwordRegex:
		err := ValidatePassword(fieldValue.String(), fieldName)
		if err != nil {
			logger.Error(err)
			return err
		}
	case hexColorRegex:
		err := ValidatorHexColor(fieldValue.String(), fieldName)
		if err != nil {
			logger.Error(err)
			return err
		}
	case phoneRegex:
		err := ValidatorPhone(fieldName, fieldValue.String())
		if err != nil {
			logger.Error(err)
			return err
		}
	}

	return nil
}

func TagRequired(requiredTagValue, fieldValue, fieldName string,
	logger *logrus_log.Logger) error {
	requiredTagBoolValue, err := strconv.ParseBool(requiredTagValue)
	if err != nil {
		return errors.New(err.Error())
	}
	if requiredTagBoolValue {
		if len(fieldValue) == 0 {
			return errors.New(fieldName + " must be required")
		}
	}
	return nil
}

func LengthRequired(fieldValue, fieldName,
	lenMinTagValue, lenMaxTagValue string, logger *logrus_log.Logger) error {
	lenMinTagIntValue, err := strconv.ParseInt(lenMinTagValue, 10, 64)
	if err != nil {
		logger.Error(err)
		return errors.New(err.Error())
	}
	lenMaxTagIntValue, err := strconv.ParseInt(lenMaxTagValue, 10, 64)
	if err != nil {
		logger.Error(fieldName)
		logger.Error(lenMaxTagValue)
		logger.Error(err)
		return errors.New(err.Error())
	}
	if lenMinTagIntValue >= int64(len(fieldValue)) || int64(len(
		fieldValue)) >= lenMaxTagIntValue {
		return errors.New(fieldName + " must be " + strconv.Itoa(int(
			lenMinTagIntValue)) + " and " + strconv.Itoa(int(
			lenMaxTagIntValue)) + " length")
	}
	return nil
}

func AmountRequired(fieldValue int, fieldName,
	amountMinTagValue, amountMaxTagValue string,
	logger *logrus_log.Logger) error {
	amountMinTagIntValue, err := strconv.Atoi(amountMinTagValue)
	if err != nil {
		logger.Error(err)
		return errors.New(err.Error())
	}
	amountMaxTagIntValue, err := strconv.Atoi(amountMaxTagValue)
	if err != nil {
		logger.Error(err)
		return errors.New(err.Error())
	}
	if amountMinTagIntValue >= fieldValue || fieldValue >= amountMaxTagIntValue {
		return errors.New(fieldName + " must be " + strconv.Itoa(amountMinTagIntValue) + " and " + strconv.Itoa(amountMaxTagIntValue) + " length")
	}
	return nil
}
