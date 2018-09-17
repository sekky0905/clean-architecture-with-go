package model

// プロパティの名称。
const (
	PropertyName = "Name"
)

// エラー系。
const (
	NameShouldBeMoreThanOneUnderTheTwenty = "Length of Name should be 0 < name < 21"
)

// エラー用の名称。
const (
	ErrorProperty = "Property"
	ErrorMessage  = "Message"
)

// モデル名。
const (
	ModelNameProgrammingLang = "ProgrammingLang"
)

// テスト用の定数。
const (
	TestName      = "testName"
	TestFeature   = "testFeature, testFeature, testFeature, testFeature, testFeature, testFeature, testFeature"
	TestDBSomeErr = "DB some error"
)

// DBの操作。
const (
	DBMethodCreate = "Create"
	DBMethodList   = "List"
	DBMethodRead   = "Read"
	DBMethodUpdate = "Update"
	DBMethodDelete = "Delete"
)
