package user_model

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name         string `json:"name"`
	Description  string `json:"description"`
	SocialNumber string `json:"social_number"`
}

var Users []UserModel = []UserModel{
	{
		Name:         "Raveline",
		Description:  "oasdnal ajsldhj asdlja hdjlahs dlkas",
		SocialNumber: "156455",
	},
	{
		Name:         "Marina",
		Description:  "sdha sajdh ajshdkja hdjahs kladhiasgd lkas",
		SocialNumber: "113131",
	},
	{
		Name:         "Luciene",
		Description:  "dlkasd lahsld halsd halsd alksd",
		SocialNumber: "2313513",
	}, {
		Name:         "Lorem Ipsib",
		Description:  " aksldkla jdklasjd lkajsdlkj alskdj als",
		SocialNumber: "45445",
	},
}
