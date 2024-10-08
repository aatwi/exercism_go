// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	resident := Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
	return &resident
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r.Name != "" && r.Address != nil && len(r.Address) != 0 {
		for add := range r.Address {
			if r.Address[add] == "" || add == "unknown key" {
				return false
			}
		}
		return true
	}
	return false
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Age = 0
	r.Name = ""
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	counter := 0
	for _, resident := range residents {
		if resident.HasRequiredInfo() == true {
			counter++
		}
	}
	return counter
}
