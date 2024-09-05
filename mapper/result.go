package mapper

type Result string

const (
	Approved            Result = "approved"
	Rejected            Result = "rejected"
	RejectedByProvider  Result = "rejected_by_provider"
	CallForAuth         Result = "call_for_auth"
	RejectedOtherReason Result = "rejected_other_reason"
	Contingency         Result = "contingency"
	NotMapped           Result = "not_mapped"
)
