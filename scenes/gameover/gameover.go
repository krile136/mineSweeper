package gameover

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

const routeType route.RouteType = route.GameOver

// respの型をポインタのResponse型にすることでnilの代入が可能になる
type Gameover struct {
	postScoreCh  chan error
	isFailedPost bool
	resp         *Response
}

func NewGameover() *Gameover {
	return &Gameover{
		postScoreCh:  make(chan error),
		isFailedPost: false,
		resp:         nil,
	}
}

type Player struct {
	Name   string `json:"name"`
	Score  int    `json:"score"`
	RankIn int    `json:"rank_in"`
}

type Response struct {
	MyScore int      `json:"my_score"`
	Players []Player `json:"players"`
}

func (g *Gameover) init() {
	g.isFailedPost = false
	g.resp = nil
}

func (g *Gameover) Update() error {
	if scene.Is_just_changed {
		g.init()
		fmt.Println("start post")
		g.postScoreCh = make(chan error)
		go g.postScore()
	}

	select {
	case err := <-g.postScoreCh:
		if err != nil {
			g.isFailedPost = true
		}
	default:
	}

	if g.isFailedPost == false && inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		// エンターキーでタイトルに戻る
		scene.RouteType = route.Title
	}

	return nil
}

func (g *Gameover) Draw(screen *ebiten.Image) {
	if g.isFailedPost {
		ebitenutil.DebugPrint(screen, "failed to post")
		return
	}

	if g.resp == nil {
		ebitenutil.DebugPrint(screen, "loading...")
	} else {
		var centerX float64 = float64(store.Data.Layout.OutsideWidth) / 2
		text.DrawTextAtCenter(screen, "G A M E   O V E R", int(centerX), 30, "L", store.Data.Color.White)

		// var jsonStr string = `{"my_score":999999,"players":[{"name":"hoge","score":1000,"rank_in":false},{"name":"you","score":1000, "rank_in":true}]}`

		// var resp Response
		// err := json.Unmarshal([]byte(jsonStr), &resp)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		if store.Data.CurrentScore > g.resp.MyScore {
			var current_score string = fmt.Sprintf("New Record!! : %d", store.Data.CurrentScore)
			text.DrawTextAtCenter(screen, current_score, int(centerX), 70, "L", store.Data.Color.White)
		} else {
			var current_score string = fmt.Sprintf("Score : %d", store.Data.CurrentScore)
			text.DrawTextAtCenter(screen, current_score, int(centerX), 70, "L", store.Data.Color.White)
		}

		var past_score = fmt.Sprintf("Your High Score : %d", g.resp.MyScore)
		text.DrawTextAtCenter(screen, past_score, int(centerX), 110, "L", store.Data.Color.White)

		text.DrawTextAtCenter(screen, "Ranking", int(centerX), 170, "L", store.Data.Color.White)

		const name_base_y int = 220

		const name_column string = "name"
		var n_c_l int = text.Length(name_column, "L")
		text.DrawText(screen, name_column, int(centerX)/2-n_c_l/2, name_base_y, "L", store.Data.Color.White)

		const score_column string = "column"
		var s_c_l int = text.Length(score_column, "L")
		text.DrawText(screen, "score", int(centerX*1.5)-s_c_l/2, name_base_y, "L", store.Data.Color.White)

		for i, r := range g.resp.Players {
			var name string
			// ランクイン表示は一旦やめとく
			// if r.RankIn == 1 {
			// 	name = fmt.Sprintf("Rank In! %s ", r.Name)
			// } else {
			// 	name = fmt.Sprintf("%s ", r.Name)
			// }
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

}

func (g *Gameover) GetRouteType() route.RouteType {
	return routeType
}
