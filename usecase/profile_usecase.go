package usecase

import (
	"kitashiruAPI/model"
	"kitashiruAPI/repository"
)

type IProfileUsecase interface {
	CreateProfile(profile model.Profile) (model.ProfileResponse, error)
	DeleteProfile(userId uint) error
}

type profileUsecase struct {
	pr repository.IProfileRepository
}

func NewProfileUsecase(pr repository.IProfileRepository) IProfileUsecase {
	return &profileUsecase{pr}
}

func (pu *profileUsecase) CreateProfile(profile model.Profile) (model.ProfileResponse, error) {
	if err := pu.pr.CreateProfile(&profile); err != nil {
		return model.ProfileResponse{}, err
	}
	resProfile := model.ProfileResponse{
		ID:          profile.ID,
		Beuraucracy: profile.Beuraucracy,
		Family:      profile.Family,
		Innovation:  profile.Innovation,
		Market:      profile.Market,
		CreatedAt:   profile.CreatedAt,
		UpdatedAt:   profile.UpdatedAt,
	}
	return resProfile, nil
}

func (pu *profileUsecase) DeleteProfile(userId uint) error {
	if err := pu.pr.DeleteProfile(userId); err != nil {
		return err
	}
	return nil
}
