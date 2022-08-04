package main

type API interface {
	Say(name string) string
}

func NewAPI(name string) API {
	if name == "dog" {
		return &DogAPI{}
	}
	if name == "cat" {
		return &CatAPI{}
	}
	return nil
}

type DogAPI struct {
}

func (d DogAPI) Say(name string) string {
	return "dog say:" + name
}

type CatAPI struct {
}

func (c CatAPI) Say(name string) string {
	return "cat say:" + name
}
