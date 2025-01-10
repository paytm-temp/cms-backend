package services

import (
    "fmt"
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
    // Validate required fields
    if c.SubCategory == "" || c.Subject == "" || c.Description == "" {
        return nil, fmt.Errorf("missing required fields: subcategory, subject, and description are required")
    }

    // Set default values if not provided
    if c.Status == "" {
        c.Status = cases.StatusOpen
    }
    if c.ResolverStatus == "" {
        c.ResolverStatus = cases.ResolverStatusNotStarted
    }
    if c.Priority == "" {
        c.Priority = cases.PriorityMedium
    }

    // Set timestamps
    now := time.Now()
    c.CreatedAt = now
    c.UpdatedAt = now

    // Create the case
    createdCase, err := s.store.CreateCase(c)
    if err != nil {
        return nil, fmt.Errorf("failed to create case: %w", err)
    }

    return createdCase, nil
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
        // Count cases by status
        switch c.Status {
        case cases.StatusOpen:
            stats.Open++
        case cases.StatusOnHold:
            stats.OnHold++
        case cases.StatusClosed:
            stats.Closed++
        default:
            // If status is not set, count as open
            stats.Open++
        }

        // Count unassigned cases
        if c.AssignedTo == "" {
            stats.Unassigned++
        }

        // Check for overdue cases
        // Cases are considered overdue if:
        // 1. They are not closed AND
        // 2. They are more than 48 hours old AND
        // 3. They are either Open or On Hold
        if c.Status != cases.StatusClosed && 
           now.Sub(c.CreatedAt) > 48*time.Hour &&
           (c.Status == cases.StatusOpen || c.Status == cases.StatusOnHold) {
            stats.Overdue++
        }
    }

    return stats, nil
}
