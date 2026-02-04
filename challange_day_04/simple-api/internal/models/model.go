package models

import "time"

type BaseModel struct {
	CreatedDate  time.Time `gorm:"column:created_date;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_date"`
	ModifiedDate time.Time `gorm:"column:modified_date;type:timestamp;default:CURRENT_TIMESTAMP" json:"modified_date"`
}
type Region struct {
	RegionID   uint   `gorm:"primaryKey;autoIncrement" json:"region_id"`
	RegionName string `gorm:"type:varchar(25)" json:"region_name"`
	// tambahkan field Countries bertipe slice []Country
	Countries []Country `gorm:"foreignKey:RegionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"countries"`
}

func (Region) TableName() string { return "hr.regions" }

type Country struct {
	CountryID   string `gorm:"type:char(2);primaryKey" json:"country_id"`
	CountryName string `gorm:"type:varchar(40)" json:"country_name"`
	RegionID    uint   `gorm:"not null" json:"region_id"`
	// remove foreign key & reference agar tidak terjadi sirkular nested json & hindari error saat create table
	Region Region `gorm:"Region;" json:"region"`
	BaseModel
}

func (Country) TableName() string { return "hr.countries" }
