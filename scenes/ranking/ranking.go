package ranking

import (
	"encoding/json"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/krile136/mineSweeper/enum/route"
	"github.com/krile136/mineSweeper/internal/text"
	"github.com/krile136/mineSweeper/store"
)

const routeType route.RouteType = route.Ranking

type Ranking struct {
}

func (r *Ranking) Update() error {
	return nil
}

type Player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type Response struct {
	MyScore int      `json:"my_score"`
	Player  []Player `json:"player"`
}

func (r *Ranking) Draw(screen *ebiten.Image) {
	var jsonStr string = `{"my_score":999999,"player":[{"name":"hoge","score":1000},{"name":"fuga","score":1000}]}`
	var resp Response
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		fmt.Println(err)
	}

	var centerX float64 = float64(store.Data.Layout.OutsideWidth) / 2
	var hiscore string = fmt.Sprintf("Your high score : %d", resp.MyScore)
	text.DrawTextAtCenter(screen, hiscore, int(centerX), 50, "L", store.Data.Color.White)

	text.DrawTextAtCenter(screen, "Ranking", int(centerX), 100, "L", store.Data.Color.White)
	for i, r := range resp.Player {
		var name string = fmt.Sprintf("%s ", r.Name)
		var name_l int = text.Length(name, "L")
		var y_position int = 150 + 30*i
		text.DrawText(screen, name, int(centerX)-name_l, y_position, "L", store.Data.Color.White)
		text.DrawText(screen, fmt.Sprintf(" %d", r.Score), int(centerX), y_position, "L", store.Data.Color.White)
	}
}

func (r *Ranking) GetRouteType() route.RouteType {
	return routeType
}
