package user_model

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name         string `json:"name" validate:"nonzero"`
	Description  string `json:"description" validate:"min=10"`
	SocialNumber string `json:"social_number" validate:"min=8, regexp=^[0-9]"`
}

func ValidateUserData(user *UserModel) error {
	if err := validator.Validate(user); err != nil {
		return err
	}
	return nil
}

var UsersList []UserModel = []UserModel{
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
