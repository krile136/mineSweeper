package game

import (
	"math/rand"
	"strconv"
	"syscall/js"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/krile136/mineSweeper/internal/draw"
	"github.com/krile136/mineSweeper/scenes/login"
	"github.com/krile136/mineSweeper/scenes/scene"
	"github.com/krile136/mineSweeper/store"
)

type Game struct {
	resourceLoadedCh chan error
}

func NewGame() (*Game, error) {
	// 初期画面としてtitle画面を設定
	scene.Display = &login.Login{}
	scene.Next = &login.Login{}

	rand.Seed(time.Now().UnixNano())

	// シーン間共通変数を初期化
	store.Data.Init()

	game := &Game{
		resourceLoadedCh: make(chan error),
	}

	// レイアウト設定
	game.Layout(store.Data.Layout.OutsideWidth, store.Data.Layout.OutsideHeight)

	// 画像リソース読み込み
	go func() {
		err := draw.LoadImages()
		if err != nil {
			game.resourceLoadedCh <- err
			return
		}

		close(game.resourceLoadedCh)
	}()

	js.Global().Set("setUserId", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		store.Data.Env.UserId, _ = strconv.Atoi(args[0].String())
		store.Data.Env.OneTimeToken = args[1].String()
		return nil
	}))

	return game, nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	if scene.Display.GetRouteType() != scene.RouteType {
		scene.Display = routeMap[scene.RouteType]
		scene.Is_just_changed = true
	}

	scene.Display.Update()

	if scene.Is_just_changed {
		scene.Is_just_changed = false
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	scene.Display.Draw(screen)
}
