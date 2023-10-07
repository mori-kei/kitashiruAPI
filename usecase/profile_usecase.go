package usecase

import (
	"kitashiruAPI/model"
	"kitashiruAPI/repository"
)

type IProfileUsecase interface {
	CreateProfile(profile model.Profile) (model.ProfileResponse, error)
	UpdateProfile(profile model.Profile, userId uint) (model.ProfileResponse, error)
	DeleteProfile(userId uint) error
	GetProfileByUserId(userId uint) (model.ProfileResponse, error)
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
func (pu *profileUsecase) UpdateProfile(profile model.Profile, userId uint) (model.ProfileResponse, error) {
	if err := pu.pr.UpdateProfile(&profile, userId); err != nil {
		return model.ProfileResponse{}, err
	}
	resProfile := model.ProfileResponse{
		ID:          profile.ID,
		Beuraucracy: profile.Beuraucracy,
		Family:      profile.Family,
		Innovation:  profile.Innovation,
		Market:      profile.Market,
	}
	return resProfile, nil
}
func (pu *profileUsecase) DeleteProfile(userId uint) error {
	if err := pu.pr.DeleteProfile(userId); err != nil {
		return err
	}
	return nil
}
func (pu *profileUsecase) GetProfileByUserId(userId uint) (model.ProfileResponse, error) {
	profile := model.Profile{}
	if err := pu.pr.GetProfileByUserId(&profile, userId); err != nil {
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
