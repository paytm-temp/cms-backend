package cases

import (
    "time"
    "fmt"
)

// Case represents a case in the system
type Status string
type Priority string
type ResolverStatus string

const (
    StatusOpen    Status = "OPEN"
    StatusClosed  Status = "CLOSED"
    StatusOnHold  Status = "ON_HOLD"

    PriorityUrgent Priority = "URGENT"
    PriorityHigh   Priority = "HIGH"
    PriorityMedium Priority = "MEDIUM"
    PriorityLow    Priority = "LOW"

    ResolverStatusNotStarted ResolverStatus = "NOT_STARTED"
    ResolverStatusInProgress ResolverStatus = "IN_PROGRESS"
    ResolverStatusDone       ResolverStatus = "DONE"
)

type Case struct {
    ID             string         `json:"id"`
    CaseNumber     string         `json:"caseNumber"`
    ResolverStatus ResolverStatus `json:"resolverStatus"`
    SubCategory    string         `json:"subCategory"`
    Contact        Contact        `json:"contact"`
    Priority       Priority       `json:"priority"`
    Status         Status         `json:"status"`
    Subject        string         `json:"subject"`
    Description    string         `json:"description"`
    AssignedTo     string         `json:"assignedTo"`
    CreatedAt      time.Time      `json:"createdAt"`
    UpdatedAt      time.Time      `json:"updatedAt"`
}

// Contact represents contact information
type Contact struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Stats represents case statistics
type Stats struct {
    Overdue    int `json:"overdue"`
    Open       int `json:"open"`
    OnHold     int `json:"onHold"`
    Closed     int `json:"closed"`
    Unassigned int `json:"unassigned"`
}
