package Vehicles

/**
We declare the interface, anything that implement GetStartMethod()
can be used as Vehicle.
*/
type Vehicle interface {
	GetStartMethod() string
}

type abstractVehicle struct {
	startMethod string
}

func (vehicle abstractVehicle) GetStartMethod() string {
	return vehicle.startMethod
}

/**
A car struct. Note this is NOT exported (lowercase name), this will
prevent the user from creating a car without using the constructor method NewCar().
*/
type car struct {
	abstractVehicle // Embedding: a car will have all the members and methods of abstractVehicle.
}

/**
Constructor method for a car, here we initialise the car members,
this is the only way to get an instance of a car.
*/
func NewCar() *car {
	return &car{abstractVehicle{"Ignition key"}}
}

/**
As above for a bicycle.
*/
type bicycle struct {
	abstractVehicle
}

func NewBicycle() *bicycle {
	return &bicycle{abstractVehicle{"Pedal"}}
}
