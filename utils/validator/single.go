package validator

func SingleFieldValidate() {
	v := validate

	var err error

	var b bool
	err = v.Var(b, "boolean")
	outRes("boolean", &err)

	var b1 = 100
	err = v.Var(b1, "number")
	outRes("boolean", &err)
}
