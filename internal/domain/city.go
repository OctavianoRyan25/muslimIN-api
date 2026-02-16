package domain

type City struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}
