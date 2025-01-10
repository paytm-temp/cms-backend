package mock

import (
    "errors"
    "fmt"
    "sync"
    "github.com/paytm-temp/cms-backend/pkg/models/case"
)

// Store represents a mock data store
type Store struct {
    cases []cases.Case
    mu    sync.RWMutex
}

// NewStore creates a new instance of Store with initial data
func NewStore() *Store {
    return &Store{
        cases: InitialCases(),
    }
}

// GetAllCases returns all cases
func (s *Store) GetAllCases() ([]cases.Case, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    return s.cases, nil
}

// GetCaseByID returns a case by its ID
func (s *Store) GetCaseByID(id string) (*cases.Case, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    for _, c := range s.cases {
        if c.ID == id {
            return &c, nil
        }
    }
    return nil, errors.New("case not found")
}

// CreateCase creates a new case
func (s *Store) CreateCase(c cases.Case) (*cases.Case, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Generate case ID and number
    nextID := len(s.cases) + 1
    c.ID = fmt.Sprintf("CASE-%03d", nextID)
    c.CaseNumber = fmt.Sprintf("CMS-2024-%03d", nextID)

    // Set default values if not provided
    if c.Status == "" {
        c.Status = cases.StatusOpen
    }
    if c.ResolverStatus == "" {
        c.ResolverStatus = cases.ResolverStatusNotStarted
    }
    if c.AssignedTo == "" {
        c.AssignedTo = "" // Unassigned by default
    }

    // Validate required fields
    if c.SubCategory == "" || c.Subject == "" || c.Description == "" {
        return nil, errors.New("missing required fields")
    }

    // Add to cases list
    s.cases = append(s.cases, c)
    return &c, nil
}

// UpdateCase updates an existing case
func (s *Store) UpdateCase(id string, updatedCase cases.Case) (*cases.Case, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    for i, c := range s.cases {
        if c.ID == id {
            s.cases[i] = updatedCase
            return &updatedCase, nil
        }
    }
    return nil, errors.New("case not found")
}

// DeleteCase deletes a case by its ID
func (s *Store) DeleteCase(id string) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    for i, c := range s.cases {
        if c.ID == id {
            s.cases = append(s.cases[:i], s.cases[i+1:]...)
            return nil
        }
    }
    return errors.New("case not found")
}
