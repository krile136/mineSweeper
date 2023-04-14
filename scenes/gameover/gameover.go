package gameover

import (
	"encoding/json"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/krile136/mineSweeper/enum/route"
	"github.com/krile136/mineSweeper/internal/text"
	"github.com/krile136/mineSweeper/scenes/scene"
	"github.com/krile136/mineSweeper/store"
)

const routeType route.RouteType = route.GameOver

type Gameover struct {
	postScoreCh chan error
}

func NewGameover() *Gameover {
	return &Gameover{
		postScoreCh: make(chan error),
	}
}

type Player struct {
	Name   string `json:"name"`
	Score  int    `json:"score"`
	RankIn bool   `json:"rank_in"`
}

type Response struct {
	MyScore int      `json:"my_score"`
	Player  []Player `json:"player"`
}

func (g *Gameover) Update() error {
	if scene.Is_just_changed {
		fmt.Println("start post")
		g.postScoreCh = make(chan error)
		go g.postScore()
	}

	return nil
}

func (g *Gameover) Draw(screen *ebiten.Image) {

	var centerX float64 = float64(store.Data.Layout.OutsideWidth) / 2
	text.DrawTextAtCenter(screen, "G A M E   O V E R", int(centerX), 30, "L", store.Data.Color.White)

	var jsonStr string = `{"my_score":999999,"player":[{"name":"hoge","score":1000,"rank_in":false},{"name":"you","score":1000, "rank_in":true}]}`

	var resp Response
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		fmt.Println(err)
	}

	if store.Data.CurrentScore > resp.MyScore {
		var current_score string = fmt.Sprintf("New Record!! : %d", store.Data.CurrentScore)
		text.DrawTextAtCenter(screen, current_score, int(centerX), 70, "L", store.Data.Color.White)
	} else {
		var current_score string = fmt.Sprintf("Score : %d", store.Data.CurrentScore)
		text.DrawTextAtCenter(screen, current_score, int(centerX), 70, "L", store.Data.Color.White)
	}

	var past_score = fmt.Sprintf("Your High Score : %d", resp.MyScore)
	text.DrawTextAtCenter(screen, past_score, int(centerX), 110, "L", store.Data.Color.White)

	text.DrawTextAtCenter(screen, "Ranking", int(centerX), 170, "L", store.Data.Color.White)
	for i, r := range resp.Player {

		var name string
		if r.RankIn {
			name = fmt.Sprintf("Rank In! %s ", r.Name)
		} else {
			name = fmt.Sprintf("%s ", r.Name)
		}
		var name_l int = text.Length(name, "L")
		var y_position int = 220 + 30*i
		text.DrawText(screen, name, int(centerX)-name_l, y_position, "L", store.Data.Color.White)
		text.DrawText(screen, fmt.Sprintf(" %d", r.Score), int(centerX), y_position, "L", store.Data.Color.White)
	}
}

func (g *Gameover) GetRouteType() route.RouteType {
	return routeType
}
