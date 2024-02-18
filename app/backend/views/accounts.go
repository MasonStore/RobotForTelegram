package views

import (
	"fmt"
	"github.com/MasonStore/RobotForTelegram/app/backend/store"
	"github.com/MasonStore/RobotForTelegram/app/backend/store/model"
	"github.com/MasonStore/RobotForTelegram/app/backend/tgbot"
	"github.com/MasonStore/RobotForTelegram/app/backend/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type AccountsView struct {
}

type AddAccountResp struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

type AddAccountReq struct {
	Token string `json:"token"`
}

func (v *AccountsView) AddAccount(req *AddAccountReq) *AddAccountResp {
	_ = store.Instance.DB
	botAPI, err := tgbotapi.NewBotAPI(req.Token)
	if err != nil {
		return &AddAccountResp{
			Ok:    false,
			Error: "NEWBOTAPI_ERROR",
		}
	}
	me, err := botAPI.GetMe()
	if err != nil {
		return &AddAccountResp{
			Ok:    false,
			Error: "GETME_ERROR",
		}
	}
	photos, err := botAPI.GetUserProfilePhotos(tgbotapi.NewUserProfilePhotos(me.ID))
	if err != nil {
		return &AddAccountResp{
			Ok:    false,
			Error: "GETUSERPROFILEPHOTOS_ERROR",
		}
	}
	if len(photos.Photos) != 0 {
		photo := photos.Photos[0][0]
		photoFile, err := botAPI.GetFileDirectURL(photo.FileID)
		if err == nil {
			utils.LogPrintf("photoFile: %s\n", photoFile)
		}
	}
	account := &model.Account{
		Model: &gorm.Model{
			ID: uint(me.ID),
		},
		Token:     req.Token,
		FirstName: me.FirstName,
		LastName:  me.LastName,
		UserName:  me.UserName,
	}
	repeatCount := int64(0)
	store.Instance.DB.Model(&model.Account{}).Where("id = ?", me.ID).Count(&repeatCount)
	if repeatCount > 0 {
		return &AddAccountResp{
			Ok:    false,
			Error: "ACCOUNT_EXISTS",
		}
	}
	store.Instance.DB.Create(account)
	tgbot.CheckRobot()
	return &AddAccountResp{
		Ok: true,
	}
}

type GetAccountsResp struct {
	Ok       bool                      `json:"ok"`
	Error    string                    `json:"error"`
	Accounts []*GetAccountsRespAccount `json:"accounts"`
}
type GetAccountsRespAccount struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Token     string    `gorm:"column:token;type:text;not null;" json:"token"`
	FirstName string    `gorm:"column:first_name;type:text;not null;default:'';" json:"first_name"`
	LastName  string    `gorm:"column:last_name;type:text;not null;default:'';" json:"last_name"`
	UserName  string    `gorm:"column:username;type:text;not null;default:'';" json:"username"`
	Running   bool      `json:"running"`
}

func (v *AccountsView) GetAccounts() *GetAccountsResp {
	_ = store.Instance.DB
	var accounts []*model.Account
	store.Instance.DB.Find(&accounts)
	var respList []*GetAccountsRespAccount
	for _, account := range accounts {
		running, _ := model.GetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyRunning, account.ID))

		respList = append(respList, &GetAccountsRespAccount{
			ID:        account.ID,
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
			Token:     account.Token,
			FirstName: account.FirstName,
			LastName:  account.LastName,
			UserName:  account.UserName,
			Running:   running == "true",
		})
	}
	return &GetAccountsResp{
		Ok:       true,
		Accounts: respList,
	}
}

func (v *AccountsView) RemoveAccount(id uint) *GetAccountsResp {
	_ = store.Instance.DB
	store.Instance.DB.Delete(&model.Account{}, id)
	tgbot.CheckRobot()
	return v.GetAccounts()
}

type SelectOneAccountResp struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

func (v *AccountsView) SelectOneAccount(id uint) *SelectOneAccountResp {
	model.SetValueByKey(store.Instance.DB, model.KeySelectedAccountID, strconv.FormatUint(uint64(id), 10))
	return &SelectOneAccountResp{
		Ok: true,
	}
}

func NowSelectedAccount() (*model.Account, bool) {
	idStr, ok := model.GetValueByKey(store.Instance.DB, model.KeySelectedAccountID)
	if !ok {
		return nil, false
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, false
	}
	if id <= 0 {
		return nil, false
	}
	var account model.Account
	store.Instance.DB.First(&account, id)
	return &account, true
}

type GetSelectedAccountResp struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
	Id    uint   `json:"id"`
}

func (v *AccountsView) GetSelectedAccount() *GetSelectedAccountResp {
	account, ok := NowSelectedAccount()
	if !ok {
		return &GetSelectedAccountResp{
			Ok: true,
			Id: 0,
		}
	}
	return &GetSelectedAccountResp{
		Ok: true,
		Id: account.ID,
	}
}
