package valueobjects

//BankAccOwnerType describes type for bank account owner
type BankAccOwnerType int

//List of bank occ owner type
const (
	AccPersonal BankAccOwnerType = iota + 1
	AccOrganization
)

func (ty BankAccOwnerType) String() string {
	if ty == AccOrganization {
		return "Organization"
	}
	return "Personal"
}
