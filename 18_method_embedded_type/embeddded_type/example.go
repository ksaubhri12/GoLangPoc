package embeddded_type

type Vehicle struct {
	Seats ,Color, Engine string
}

type Car struct {
	Vehicle *Vehicle
	Wheels string
	Doors string
}

func (vehicle *Vehicle)SetEnginePower(enginePower string)  {
	vehicle.Engine = enginePower
}

func (vehicle *Vehicle)SetColor(colour string)  {
	vehicle.Color = colour
}