package userProfile

import (
	"log"
)

const (
	//By every 10 ARS spent user wins 1 experience point
	ExpRateByAmount = 10.0
)

//Levels and experience for the next one
var levels = map[int]LevelConfig{
	1: {
		experienceForLevelUp: 100,
		discount:             0.0,
	},
	2: {
		experienceForLevelUp: 200,
		discount:             0.10,
	},
	3: {
		experienceForLevelUp: 300,
		discount:             0.15,
	},
	4: {
		experienceForLevelUp: 500,
		discount:             0.20,
	},
}

type Service interface {
	CreateNewProfile(userID string) *Profile
	GetProfile(userID string) *Profile
	UpdateProfile(userID string, amountSpent float64) bool
}

type ProfileRepo interface {
	SaveNewProfile(newProfile *Profile) error
	FindByID(userID string) (*Profile, error)
	UpdateProfile(userID string, profile *Profile) error
}

type service struct {
	profileRepo ProfileRepo
}
/**
 * @api {get} /v1/loyalty/userProfile Consultar Profile
 * @apiName Consultar un profile
 * @apiGroup Loyalty
 * @apiParam Authorization{bearer-token}
 *
 * @apiDescription Consulta un profile existente
 * @apiSuccessExample {json} Respuesta
 * HTTP/1.1 200 OK
 *  {
 *  "userID": "{userID}",
 *  "userLevel" : "{userLevel}" ,
 *  "experience" :    "{experiencia actual}",
 *  "currentDiscount" : "{descuento actual}"
 *	}
 *  @apiUse ParamValidationErrors
 *  @apiUse OtherErrors
 */
func (s service) GetProfile(userID string) *Profile {
	profile, err := s.profileRepo.FindByID(userID)
	if err != nil {
		log.Panic(err)
	}
	return profile
}
/**
 * @api {post} /v1/loyalty/userProfile Crear Profile
 * @apiName Crear Profile
 * @apiGroup Loyalty
 *
 * @apiDescription Crea y asocia un profile a un nuevo usuario.
 * @apiParam Body
 *    {
 *      "userID": "{userID}"
 *    }
 *
 * @apiSuccessExample {string} Body
 *    HTTP/1.1 200 Ok
 *  @apiUse ParamValidationErrors
 *  @apiUse OtherErrors
 **/
func (s service) UpdateProfile(userID string, amountSpent float64) bool {
	var hasLevelUp bool
	uProfile, err := s.profileRepo.FindByID(userID)
	if err != nil {
		panic(err)
	}
	currentExp := uProfile.Experience
	uProfile.Experience = currentExp + int(amountSpent/ExpRateByAmount)
	currentLevel := uProfile.UserLevel
	for levels[currentLevel].experienceForLevelUp < uProfile.Experience {
		currentLevel++
		hasLevelUp = true
	}
	uProfile.UserLevel = currentLevel
	uProfile.CurrentDiscount = levels[currentLevel].discount
	err = s.profileRepo.UpdateProfile(userID, uProfile)
	if err != nil {
		panic(err)
	}

	return hasLevelUp
}

func (s service) CreateNewProfile(userID string) *Profile {
	profile := &Profile{
		UserID:          userID,
		UserLevel:       1,
		Experience:      0,
		CurrentDiscount: levels[1].discount,
	}
	err := s.profileRepo.SaveNewProfile(profile)
	if err != nil {
		log.Panic(err)
	}
	return profile
}

func NewService(profile ProfileRepo) Service {
	return &service{profileRepo: profile}
}
