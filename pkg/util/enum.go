package util

import (
	"errors"
	"fmt"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/pkg/errx"
	"github.com/kataras/iris/v12"
)

func ParsePersonalizationFrequenceSport(frequenceSport string) (domain.PersonalizationFrequenceSport, error) {
	switch frequenceSport {
	case "OncePerWeek":
		return domain.PersonalizationFrequenceSportOncePerWeek, nil
	case "OnceToThreePerWeek":
		return domain.PersonalizationFrequenceSportOnceToThreePerWeek, nil
	case "FourToFiveTimesPerWeek":
		return domain.PersonalizationFrequenceSportFourToFiveTimesPerWeek, nil
	case "FiveToSevenTimesPerWeek":
		return domain.PersonalizationFrequenceSportFiveToSevenTimesPerWeek, nil
	default:
		return "", errx.New().
			WithCode(iris.StatusBadRequest).
			WithMessage(fmt.Sprintf("invalid frequence sport: %s", frequenceSport)).
			WithError(errors.New("invalid frequence sport type"))
	}
}

func ParsePersonalizationGender(gender string) (domain.PersonalizationGender, error) {
	switch gender {
	case "Male":
		return domain.PersonalizationGenderMale, nil
	case "Female":
		return domain.PersonalizationGenderFemale, nil
	default:
		return "", errx.New().
			WithCode(iris.StatusBadRequest).
			WithMessage(fmt.Sprintf("invalid gender : %s", gender)).
			WithError(errors.New("invalid gender type"))
	}
}
