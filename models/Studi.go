package models

type Studi struct {
	Id     uint   `gorm:"primaryKey" json:"id"`
	IdUser uint   `json:"id_user"`
	S01    string `gorm:"column:S01" json:"S01"` //jenjang pendidikan melanjutkan
	S02    string `gorm:"column:S02" json:"S02"` //Nama Perguruan Tinggi
	S03    string `gorm:"column:S03" json:"S03"` //Program Studi
	S04    string `gorm:"column:S04" json:"S04"` // Alasan Melanjutkan
}
