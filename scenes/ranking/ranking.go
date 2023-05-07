package ranking

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/krile136/mineSweeper/enum/route"
	"github.com/krile136/mineSweeper/internal/text"
	"github.com/krile136/mineSweeper/scenes/scene"
	"github.com/krile136/mineSweeper/store"
)

const routeType route.RouteType = route.Ranking

type Ranking struct {
	getRankingIndexCh       chan error
	isFailedGetRankingIndex bool
	resp                    *Response
}

type Player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type Response struct {
	MyScore int      `json:"my_score"`
	Players []Player `json:"players"`
}

func (r *Ranking) init() {
	r.isFailedGetRankingIndex = false
	r.resp = nil
}

func (r *Ranking) Update() error {
	if scene.Is_just_changed {
		r.init()
		r.getRankingIndexCh = make(chan error)
		go r.getRanking()
	}

	select {
	case err := <-r.getRankingIndexCh:
		if err != nil {
			r.isFailedGetRankingIndex = true
		}
	default:
	}

	if r.isFailedGetRankingIndex == false && inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		// エンターキーでタイトルに戻る
		scene.RouteType = route.Title
	}
	return nil
}

func (r *Ranking) Draw(screen *ebiten.Image) {
	if r.isFailedGetRankingIndex {
		ebitenutil.DebugPrint(screen, "failed to get")
		return
	}

	if r.resp == nil {
		ebitenutil.DebugPrint(screen, "loading...")
		return
	}

	var centerX float64 = float64(store.Data.Layout.OutsideWidth) / 2
	var hiscore string = fmt.Sprintf("Your high score : %d", r.resp.MyScore)
	text.DrawTextAtCenter(screen, hiscore, int(centerX), 50, "L", store.Data.Color.White)

	text.DrawTextAtCenter(screen, "Ranking", int(centerX), 100, "L", store.Data.Color.White)

	const name_base_y int = 150

	const name_column string = "name"
	var n_c_l int = text.Length(name_column, "L")
	text.DrawText(screen, name_column, int(centerX)/2-n_c_l/2, name_base_y, "L", store.Data.Color.White)

	const score_column string = "column"
	var s_c_l int = text.Length(score_column, "L")
	text.DrawText(screen, "score", int(centerX*1.5)-s_c_l/2, name_base_y, "L", store.Data.Color.White)

	for i, r := range r.resp.Players {
		var name string
		name = fmt.Sprintf("%s ", r.Name)
		var name_l int = text.Length(name, "L")
		var y_position int = name_base_y + 10 + 30*(i+1)
		text.DrawText(screen, name, int(centerX)/2-name_l/2, y_position, "L", store.Data.Color.White)
		var score_text string = fmt.Sprintf("%d", r.Score)
		var score_text_l int = text.Length(score_text, "L")
		text.DrawText(screen, score_text, int(centerX*1.5)-score_text_l/2, y_position, "L", store.Data.Color.White)
	}

	text.DrawTextAtCenter(screen, "Push enter to return title", int(centerX), 450, "L", store.Data.Color.White)

}

func (r *Ranking) GetRouteType() route.RouteType {
	return routeType
}
