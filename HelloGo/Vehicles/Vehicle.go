package Vehicles

/**
We declare the interface, anything that implement GetStartMethod()
can be used as Vehicle.
*/
type Vehicle interface {
	GetStartMethod() string
}

/**
A car struct. Note this is NOT exported (lowercase name), this will
prevent the user from creating a car without using the constructor method.
*/
type car struct {
	startMethod string
}

/**
Constructor method for a car, here we initialise the car members,
this is the only way to get an instance of a car.
*/
func NewCar() *car {
	return &car{startMethod: "Ignition key"}
}

func (car car) GetStartMethod() string {
	return car.startMethod
}

/**
As above for a bicycle.
*/
type bicycle struct {
	startMethod string
}

func NewBicycle() *bicycle {
	return &bicycle{startMethod: "Pedal"}
}

func (bicycle bicycle) GetStartMethod() string {
	return bicycle.startMethod
}

func FindOutHowToStart(vehicle Vehicle) string {
	return vehicle.GetStartMethod()
}
