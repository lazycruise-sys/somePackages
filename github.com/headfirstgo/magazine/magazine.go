package magazine

type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	Address
}

type Employee struct {
	Name        string
	Salary      float64
	Address
	History
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

type History struct {
	School string
	Work string
}
