package valueobjects

//SupplierStatus Provides an alias type for Supplier Status
type SupplierStatus int

//Defines states of supplier
const (
	ActiveSupplier SupplierStatus = iota + 1
	InactiveSupplier
)

func (st SupplierStatus) String() string {
	if st == ActiveSupplier {
		return "Active"
	}
	return "Inactive"
}
