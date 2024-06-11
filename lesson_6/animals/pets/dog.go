package pets

type Dog struct {
	Animal
	Age      uint8
	Weight   uint8
	IsAsleep bool
}

func (d *Dog) Eat(amount uint8) (uint8, error) {
	if d.IsAsleep {
		return 0, &ActionError{Name: d.Name, Reason: "it is asleep"}
	}

	if amount > 15 {
		return 0, newError("Dog can't eat that much", nil)
	}
	return amount, nil
}

func (d *Dog) Walk() string {
	return "Dog is walking!"
}
