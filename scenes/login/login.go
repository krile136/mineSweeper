package login

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/krile136/mineSweeper/enum/route"
	"github.com/krile136/mineSweeper/scenes/scene"
	"github.com/krile136/mineSweeper/store"
)

const routeType route.RouteType = route.Login

type Login struct {
	getApiTokenCh chan error
	getXrsfTokenCh chan error
}

var (
	isFailedToGetApiToken        = false
	tick                  int    = 0
	dot                   string = ""
	isExecuteLogin               = false
)

func NewLogin() *Login {
	return &Login{
		getApiTokenCh: make(chan error),
		getXrsfTokenCh: make(chan error),
	}
}

func (l *Login) Update() error {
	if scene.Is_just_changed {
		// 最初にチャネルを初期化する
		l.getApiTokenCh = make(chan error)
	}

	if store.Data.Env.UserId != -999 && store.Data.Env.OneTimeToken != "" && isExecuteLogin == false {
		// jsで必要なデータが渡ってきていたら一度だけログインAPIを叩く
		log.Print("start login")
		isExecuteLogin = true
		go l.login()
	}

	select {
	case err := <-l.getApiTokenCh:
		if err != nil {
			log.Printf("%s", err)
			isFailedToGetApiToken = true
		}
	default:
	}

	if store.Data.Env.ApiToken == "" {
		if tick%10 == 0 {
			dot += "."
		}
		tick += 1
	} else {
		// ApiToken取れたらタイトルへ遷移
		scene.RouteType = route.Title
	}

	return nil
}

func (l *Login) Draw(screen *ebiten.Image) {
	if isFailedToGetApiToken {
		ebitenutil.DebugPrint(screen, "failed to login. Please Reload")
	} else {
		ebitenutil.DebugPrint(screen, "log in"+dot)
	}
}

func (l *Login) GetRouteType() route.RouteType {
	return routeType
}
