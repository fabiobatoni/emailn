package campaign

import "emailn/internal/contract"

type Service struct {
	Repository Repository
}

func (s *Service) Create(NewCampaign contract.NewCampaign) error {
	return nil
}
