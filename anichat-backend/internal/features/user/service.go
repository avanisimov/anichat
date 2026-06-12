package user

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUserProfile(ctx context.Context, userID string) (*UsersMeData, error) {
	usersMeDTO, err := s.repo.GetProfileByUserID(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}
	var profile *UserProfile
	if usersMeDTO.ProfileCreated {
		profile = &UserProfile{
			FirstName: *usersMeDTO.FirstName,
			LastName:  *usersMeDTO.LastName,
			AvatarURL: usersMeDTO.AvatarURL,
		}
	}

	resp := &UsersMeData{
		ID:             usersMeDTO.ID,
		Email:          usersMeDTO.Email,
		ProfileCreated: usersMeDTO.ProfileCreated,
		Profile:        profile,
	}

	return resp, nil
}

func (s *Service) UpsertProfile(
	ctx context.Context,
	userID string,
	req UpsertProfileRequest,
) error {

	return s.repo.UpsertProfile(
		ctx,
		userID,
		req.FirstName,
		req.LastName,
	)
}