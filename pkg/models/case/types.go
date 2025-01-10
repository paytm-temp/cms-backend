package cases

// Status represents the status of a case
type Status string

const (
    StatusOpen       Status = "OPEN"
    StatusClosed     Status = "CLOSED"
    StatusInProgress Status = "IN_PROGRESS"
    StatusOnHold     Status = "ON_HOLD"
)

// Priority represents the priority level of a case
type Priority string

const (
    PriorityUrgent Priority = "URGENT"
    PriorityHigh   Priority = "HIGH"
    PriorityMedium Priority = "MEDIUM"
    PriorityLow    Priority = "LOW"
)

// ResolverStatus represents the status set by resolver
type ResolverStatus string

const (
    ResolverStatusNotStarted ResolverStatus = "NOT_STARTED"
    ResolverStatusInProgress ResolverStatus = "IN_PROGRESS"
    ResolverStatusDone       ResolverStatus = "DONE"
)
