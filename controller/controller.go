package Controller

type Controller struct{}

type Alex struct {
	Name string
	Age  int
}

func (c *Controller) Index() Alex {
	return Alex{"Alex Merced", 34}
}
