package model

type Category struct {
	id      string
	name    string
	creator User
}

func NewCategory(id, name string, creator User) *Category {
	return &Category{
		id:      id,
		name:    name,
		creator: creator,
	}
}

func (c *Category) GetId() string {
	return c.id
}

func (c *Category) GetName() string {
	return c.name
}

func (c *Category) GetCreator() User {
	return c.creator
}

func (c *Category) SetName(name string) {
	c.name = name
}

func (c *Category) SetCreator(creator User) {
	c.creator = creator
}

func (c *Category) SetId(id string) {
	c.id = id
}
