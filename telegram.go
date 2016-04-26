package trib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/net/context"

	bot "gopkg.in/telegram-bot-api.v4"
)

const (
	telegramFrontName = "Telegram"
	telegramUpdate    = "telegram_update"
	telegramEcho      = "plan_echo"
)

type TelegramBot interface {
	GetMe() (bot.User, error)
	MakeRequest(endpoint string, params url.Values) (bot.APIResponse, error)
}

type Telegram struct {
	api TelegramBot
}

func (t *Telegram) Name() string {
	return telegramFrontName
}

func (t *Telegram) Plan(r *http.Request) (Plan, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	update := bot.Update{}
	err = json.Unmarshal(b, &update)
	if err != nil {
		return nil, err
	}
	if update.Message == nil {
		return nil, errors.New("no message")
	}
	if update.Message.IsCommand() {
	}
	return nil, nil
}

func (t *Telegram) newCommandPlan(u bot.Update) (Plan, error) {
	switch u.Message.Command() {
	case "echo":
		return t.planEcho(u)
	}
	return nil, errors.New("no plan")
}

func (t *Telegram) planEcho(u bot.Update) (Plan, error) {
	p := NewBasePlan(telegramEcho)
	s := NewBaseStage("echo")
	echo := NewBaseStep("echo_back", false, t.Echo)
	s.AddStep(echo)
	p.AddStage(s)
	return p, nil
}

func (t *Telegram) Commit(ctx context.Context) {
}

func (t *Telegram) Echo(ctx context.Context) error {
	aCtx := ActiveCtx(ctx)
	if aCtx == nil {
		return errors.New("can't find context")
	}
	up, ok := GetOk(ctx, telegramUpdate)
	if !ok {
		return errors.New(" no telegram update to work on")
	}
	update := up.(bot.Update)
	fmt.Println(update.Message.CommandArguments())
	return nil
}
