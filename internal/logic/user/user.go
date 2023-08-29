package user

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf-demo-user/v2/internal/dao"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/model/do"
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
	"github.com/gogf/gf-demo-user/v2/internal/service"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// Create creates user account.
func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (err error) {
	// If Nickname is not specified, it then uses Passport as its default Nickname.
	if in.Nickname == "" {
		in.Nickname = in.Passport
	}
	var (
		available bool
	)
	available, err = s.VerifyCodeCheck(ctx, in.VerifyCode, in.PhoneNumber)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`VerificationCode "%s" is wrong`, in.VerifyCode)
	}
	// Passport checks.
	available, err = s.IsPassportAvailable(ctx, in.Passport)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Passport "%s" is already token by others`, in.Passport)
	}
	// Nickname checks.
	available, err = s.IsNicknameAvailable(ctx, in.Nickname)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Nickname "%s" is already token by others`, in.Nickname)
	}
	// Phonenumber checks.
	available, err = s.IsPhonenumberAvailable(ctx, in.Passport)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Passport "%s" is already token by others`, in.Passport)
	}
	return dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(do.User{
			Passport:    in.Passport,
			Password:    in.Password,
			Nickname:    in.Nickname,
			Phonenumber: in.PhoneNumber,
		}).Insert()
		return err
	})
}

func (s *sUser) AdminCreate(ctx context.Context, in model.AdminSignUp) (err error) {
	// If Nickname is not specified, it then uses Passport as its default Nickname.
	if in.Nickname == "" {
		in.Nickname = in.Passport
	}
	var (
		available bool
	)

	// Passport checks.
	available, err = s.IsPassportAvailable(ctx, in.Passport)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Passport "%s" is already token by others`, in.Passport)
	}
	// Nickname checks.
	available, err = s.IsNicknameAvailable(ctx, in.Nickname)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Nickname "%s" is already token by others`, in.Nickname)
	}

	return dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(do.User{
			Passport: in.Passport,
			Password: in.Password,
			Nickname: in.Nickname,
			IsAdmin:  1,
		}).Insert()
		return err
	})
}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *sUser) IsSignedIn(ctx context.Context) bool {
	if v := service.BizCtx().Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// SignIn creates session for given user account.
func (s *sUser) SignIn(ctx context.Context, in model.UserSignInInput) (token string, err error) {
	var user *entity.User
	err = dao.User.Ctx(ctx).Where(do.User{
		Passport: in.Passport,
		Password: in.Password,
	}).Scan(&user)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", gerror.New(`Passport or Password not correct`)
	}
	if token, err = service.Session().SetUser(ctx, user); err != nil {
		return "", err
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Passport: user.Passport,
		Nickname: user.Nickname,
	})
	return token, nil
}

// SignOut removes the session for current signed-in user.
func (s *sUser) SignOut(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}

// IsPassportAvailable checks and returns given passport is available for signing up.
func (s *sUser) IsPassportAvailable(ctx context.Context, passport string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Passport: passport,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsNicknameAvailable checks and returns given nickname is available for signing up.
func (s *sUser) IsNicknameAvailable(ctx context.Context, nickname string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Nickname: nickname,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (s *sUser) IsPhonenumberAvailable(ctx context.Context, phonenumber string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Phonenumber: phonenumber,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// GetProfile retrieves and returns current user info in session.
func (s *sUser) GetProfile(ctx context.Context) *entity.User {
	return service.Session().GetUser(ctx)
}

func (s *sUser) VerifyCodeSend(ctx context.Context, PhoneNumber string) (err error) {
	vcode := CreateRandomNumber(6)

	//联系第三方向指定手机号发送短信
	return dao.Verificationcode.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Verificationcode.Ctx(ctx).Data(do.Verificationcode{
			Phonenumber:      PhoneNumber,
			Verificationcode: vcode,
		}).Insert()
		_, err = dao.Verificationcode.Ctx(ctx).Where("create_at<?", gtime.Now().Add(time.Duration(-5)*time.Minute)).Delete()
		return err
	})
}

func (s *sUser) VerifyCodeCheck(ctx context.Context, checkverifycode string, phonenumber string) (bool, error) {
	count, err := dao.Verificationcode.Ctx(ctx).Where(dao.Verificationcode.Columns().Phonenumber, phonenumber).Where(dao.Verificationcode.Columns().Verificationcode, checkverifycode).Count()
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

func CreateRandomNumber(len int) string {
	var numbers = []byte{1, 2, 3, 4, 5, 7, 8, 9}
	var container string
	length := bytes.NewReader(numbers).Len()

	for i := 1; i <= len; i++ {
		random, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
		}
		container += fmt.Sprintf("%d", numbers[random.Int64()])
	}
	return container
}
