package model

type BaseModel struct {
	Id int
}

func (model BaseModel) getId() int{
	return model.Id
}

func (model BaseModel) setId(id int) {
	model.Id = id
}