package models

type Data struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	FullName string `gorm:"size:100" json:"full_name"`
	Email    string `gorm:"size:100" json:"email"`
	Address  string `gorm:"size:255" json:"address"`
}

// NewData create new data to database
func NewData(data *Data) error {
	err := db.Create(data).Error
	return err
}

// GetDataByID return data by passing the ID
func GetDataByID(id uint) (Data, error) {
	var data Data
	err := db.Model(&Data{}).Where("id = ?", id).First(&data).Error
	return data, err
}

// GetAllData returns all data
func GetAllData() []Data {
	var datas []Data
	db.Model(&Data{}).Find(&datas)
	return datas
}
