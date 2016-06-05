package klondike

import (
	"errors"
	"math"
)

var (
	InvalidMovement       = errors.New("invalid movement")
	StockIsEmpty          = errors.New("stock is empty")
	StockIsNotEmpty       = errors.New("stock is not empty")
	CardIsNotExist        = errors.New("card is not exist")
	CardIsNotOpen         = errors.New("card is not open")
	CardIsNotLastCol      = errors.New("card is not last col")
	CanNotPutInFoundation = errors.New("can not put in foundation")
	CanNotPutInTableau    = errors.New("can not put in tableau")
)

func (k *Klondike) Move() error {
	fromRow, toRow := k.Selected.Row, k.Cursor.Row
	err := InvalidMovement

	switch {
	case isStock(fromRow) && isStock(toRow) && len(k.Table[stock]) > 0:
		// 山札をめくる
		err = k.Draw()
	case isStock(fromRow) && isStock(toRow) && len(k.Table[stock]) <= 0:
		// 捨て札を山札に
		err = k.WasteToStock()
	case isWaste(fromRow) && isFoundation(toRow):
		// 捨て札を組み札に
		err = k.MoveToFoundation(k.Selected, k.Cursor)
	case isWaste(fromRow) && isColumn(toRow):
		// 捨て札を場札に
		err = k.MoveToColumn(k.Selected, k.Cursor)
	case k.Selected.Equal(k.Cursor) && (isWaste(fromRow) || isColumn(fromRow)):
		// 組札に出せたら出す
		err = k.SearchAndMoveToFoundation(k.Selected, k.Cursor)
	case isColumn(fromRow) && isColumn(toRow):
		// 場札から場札に
		err = k.MoveToColumn(k.Selected, k.Cursor)
	case isColumn(fromRow) && isFoundation(toRow):
		// 場札から組み札に
		err = k.MoveToFoundation(k.Selected, k.Cursor)
	}

	k.CursorReset()
	k.Selected = nil

	return err
}

func isStock(i int) bool {
	return i == stock
}

func isWaste(i int) bool {
	return i == waste
}

func isFoundation(i int) bool {
	return foundation1 <= i && i <= foundation4
}

func isColumn(i int) bool {
	return column1 <= i && i <= column7
}

func (k *Klondike) Draw() error {
	if len(k.Table[stock]) <= 0 {
		return StockIsEmpty
	}
	last := len(k.Table[stock]) - 1
	k.Table[stock][last].Open = true
	k.Table[waste] = append(k.Table[waste], k.Table[stock][last])
	k.Table[stock] = k.Table[stock][:last]
	return nil
}

func (k *Klondike) WasteToStock() error {
	if len(k.Table[stock]) > 0 {
		return StockIsNotEmpty
	}
	k.Table[stock], k.Table[waste] = k.Table[waste], k.Table[stock]
	for _, c := range k.Table[stock] {
		c.Open = false
	}
	// スコア反映
	k.Score = int(math.Max(0, float64(k.Score-100)))
	return nil
}

func (k *Klondike) MoveToFoundation(from, to *Position) error {
	// from が不正
	if !isWaste(from.Row) && !isColumn(from.Row) {
		return InvalidMovement
	}
	// to が不正
	if !isFoundation(to.Row) {
		return InvalidMovement
	}
	// 一番上のカードしか移動できない
	if from.Col != k.LastCol(from.Row) || to.Col != k.LastCol(to.Row) {
		return CardIsNotLastCol
	}
	// カード取得
	fromCard, toCard := k.GetCard(from), k.GetCard(to)
	// オープンしているカードだけ
	if !fromCard.Open {
		return CardIsNotOpen
	}
	// A もしくは同じスートの +1 のみ置ける
	if !fromCard.CanPutInFoundation(toCard) {
		return CanNotPutInFoundation
	}
	// 移動実行
	k.Table[to.Row] = append(k.Table[to.Row], fromCard)
	k.Table[from.Row] = k.Table[from.Row][:from.Col]
	// スコア反映
	k.Score += 10
	// オープン
	if len(k.Table[from.Row]) > 0 {
		k.Table[from.Row][from.Col-1].Open = true
	}
	return nil
}

func (k *Klondike) MoveToColumn(from, to *Position) error {
	// from が不正
	if !isWaste(from.Row) && !isColumn(from.Row) {
		return InvalidMovement
	}
	// to が不正
	if !isColumn(to.Row) {
		return InvalidMovement
	}
	// 捨て札は一番上のカードしか移動できない
	if isWaste(from.Row) && from.Col != k.LastCol(from.Row) {
		return CardIsNotLastCol
	}
	// 移動先は一番上のカード
	if to.Col != k.LastCol(to.Row) {
		return CardIsNotLastCol
	}
	// カード取得
	fromCard, toCard := k.GetCard(from), k.GetCard(to)
	// カードがある？
	if fromCard == nil {
		return CardIsNotExist
	}
	// オープンか？
	if !fromCard.Open || (toCard != nil && !toCard.Open) {
		return CardIsNotOpen
	}
	// 13 か色違いで -1 のみ置ける
	if !fromCard.CanPutInTableau(toCard) {
		return CanNotPutInTableau
	}
	// 移動実行
	k.Table[to.Row] = append(k.Table[to.Row], k.Table[from.Row][from.Col:]...)
	k.Table[from.Row] = k.Table[from.Row][:from.Col]
	// スコア反映
	k.Score += 5
	// オープン
	if len(k.Table[from.Row]) > 0 {
		k.Table[from.Row][from.Col-1].Open = true
	}
	return nil
}

func (k *Klondike) SearchAndMoveToFoundation(p1, p2 *Position) error {
	// 同じカード？
	if !p1.Equal(p2) {
		return InvalidMovement
	}
	// 捨て札か場札
	if !isWaste(p1.Row) && !isColumn(p1.Row) {
		return InvalidMovement
	}
	// 一番上？
	if p1.Col != k.LastCol(p1.Row) {
		return CardIsNotLastCol
	}
	// カード取得
	c := k.GetCard(p1)
	if c == nil {
		return CardIsNotExist
	}
	// オープンか？
	if !c.Open {
		return CardIsNotOpen
	}
	// 移動先組札の検索
	for i := foundation1; i <= foundation4; i++ {
		to := &Position{i, k.LastCol(i)}
		c2 := k.GetCard(to)
		if c.CanPutInFoundation(c2) {
			return k.MoveToFoundation(p1, to)
		}
	}
	return nil
}
