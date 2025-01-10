package mock

import (
    "time"
    "github.com/paytm-temp/cms-backend/pkg/models/case"
)

// InitialCases returns a slice of initial mock cases
func InitialCases() []cases.Case {
    return []cases.Case{
        {
            ID:             "CASE-001",
            CaseNumber:     "CMS-2024-001",
            ResolverStatus: cases.ResolverStatusInProgress,
            SubCategory:    "Customer Complaint",
            Contact: cases.Contact{
                Name:  "John Doe",
                Email: "john.doe@example.com",
            },
            Priority:    cases.PriorityHigh,
            Status:      cases.StatusOpen,
            Subject:     "Payment Failed",
            Description: "Unable to complete payment transaction",
            AssignedTo:  "agent1",
            CreatedAt:   time.Now().Add(-24 * time.Hour),
            UpdatedAt:   time.Now(),
        },
        {
            ID:             "CASE-002",
            CaseNumber:     "CMS-2024-002",
            ResolverStatus: cases.ResolverStatusDone,
            SubCategory:    "Internal Request",
            Contact: cases.Contact{
                Name:  "Jane Smith",
                Email: "jane.smith@example.com",
            },
            Priority:    cases.PriorityMedium,
            Status:      cases.StatusClosed,
            Subject:     "Access Request",
            Description: "Need access to reporting dashboard",
            AssignedTo:  "agent2",
            CreatedAt:   time.Now().Add(-48 * time.Hour),
            UpdatedAt:   time.Now().Add(-24 * time.Hour),
        },
        {
            ID:             "CASE-003",
            CaseNumber:     "CMS-2024-003",
            ResolverStatus: cases.ResolverStatusNotStarted,
            SubCategory:    "Lender Request",
            Contact: cases.Contact{
                Name:  "Alice Johnson",
                Email: "alice.johnson@example.com",
            },
            Priority:    cases.PriorityUrgent,
            Status:      cases.StatusOnHold,
            Subject:     "Integration Issue",
            Description: "API integration failing with timeout",
            AssignedTo:  "agent3",
            CreatedAt:   time.Now().Add(-72 * time.Hour),
            UpdatedAt:   time.Now().Add(-48 * time.Hour),
        },
    }
}
