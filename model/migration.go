package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	_ = DB.AutoMigrate(&User{})
	_ = DB.AutoMigrate(&Content{})
	_ = DB.AutoMigrate(&Psychological{})
	_ = DB.AutoMigrate(&Craft{})
	_ = DB.AutoMigrate(&Tag{})
	_ = DB.AutoMigrate(&Luntan{})
	_ = DB.AutoMigrate(&TestResult{})
	_ = DB.AutoMigrate(&TestBasic{})
	_ = DB.AutoMigrate(&KeyWordScene{})
	_ = DB.AutoMigrate(&Object{})
}
