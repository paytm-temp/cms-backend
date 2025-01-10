package services

import (
    "time"
    "github.com/paytm-temp/cms-backend/pkg/database/mock"
    "github.com/paytm-temp/cms-backend/pkg/models/case"
)

type CaseService struct {
    store *mock.Store
}

func NewCaseService() *CaseService {
    return &CaseService{
        store: mock.NewStore(),
    }
}

func (s *CaseService) GetAllCases() ([]cases.Case, error) {
    return s.store.GetAllCases()
}

func (s *CaseService) GetCaseByID(id string) (*cases.Case, error) {
    return s.store.GetCaseByID(id)
}

func (s *CaseService) CreateCase(c cases.Case) (*cases.Case, error) {
    c.CreatedAt = time.Now()
    c.UpdatedAt = time.Now()
    return s.store.CreateCase(c)
}

func (s *CaseService) UpdateCase(id string, c cases.Case) (*cases.Case, error) {
    c.UpdatedAt = time.Now()
    return s.store.UpdateCase(id, c)
}

func (s *CaseService) DeleteCase(id string) error {
    return s.store.DeleteCase(id)
}

func (s *CaseService) GetStats() (*cases.Stats, error) {
    allCases, err := s.store.GetAllCases()
    if err != nil {
        return nil, err
    }

    stats := &cases.Stats{}
    now := time.Now()

    for _, c := range allCases {
        switch c.Status {
        case cases.StatusOpen:
            stats.Open++
        case cases.StatusOnHold:
            stats.OnHold++
        case cases.StatusClosed:
            stats.Closed++
        }

        if c.AssignedTo == "" {
            stats.Unassigned++
        }

        // Check for overdue cases (more than 48 hours old and not closed)
        if c.Status != cases.StatusClosed && now.Sub(c.CreatedAt) > 48*time.Hour {
            stats.Overdue++
        }
    }

    return stats, nil
}
