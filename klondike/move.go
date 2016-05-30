package klondike

import (
	"errors"
	"log"
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
	fromRow, fromCol := k.selected.Row, k.selected.Col
	toRow, toCol := k.cursor.Row, k.cursor.Col
	fromCard, toCard := k.table[fromRow][fromCol], k.table[toRow][toCol]
	log.Println("START MOVE: ", k.selected, k.cursor, fromCard, toCard)

	switch {
	case isStock(fromRow) && isStock(toRow) && len(k.table[stock]) > 0:
		// 山札をめくる
		return k.Draw()
	case isStock(fromRow) && isStock(toRow) && len(k.table[stock]) <= 0:
		// 捨て札を山札に
		return k.WasteToStock()
	case isWaste(fromRow) && isFoundation(toRow):
		// 捨て札を組み札に
		return k.MoveToFoundation(k.selected, k.cursor)
	case isWaste(fromRow) && isColumn(toRow):
		// 捨て札を場札に
		return k.MoveToColumn(k.selected, k.cursor)
	case isColumn(fromRow) && isColumn(toRow):
		// 場札から場札に
		return k.MoveToColumn(k.selected, k.cursor)
	case isColumn(fromRow) && isFoundation(toRow):
		// 場札から組み札に
		return k.MoveToFoundation(k.selected, k.cursor)
	}

	k.selected = nil

	return InvalidMovement
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
	if len(k.table[stock]) <= 0 {
		return StockIsEmpty
	}
	last := len(k.table[stock]) - 1
	k.table[stock][last].Open = true
	k.table[waste] = append(k.table[waste], k.table[stock][last])
	k.table[stock] = k.table[stock][:last]
	return nil
}

func (k *Klondike) WasteToStock() error {
	if len(k.table[stock]) > 0 {
		return StockIsNotEmpty
	}
	k.table[stock], k.table[waste] = k.table[waste], k.table[stock]
	for _, c := range k.table[stock] {
		c.Open = false
	}
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
	k.table[to.Row] = append(k.table[to.Row], fromCard)
	k.table[from.Row] = k.table[from.Row][:from.Col]
	// オープン
	if len(k.table[from.Row]) > 0 {
		k.table[from.Row][from.Col-1].Open = true
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
	k.table[to.Row] = append(k.table[to.Row], k.table[from.Row][from.Col:]...)
	k.table[from.Row] = k.table[from.Row][:from.Col]
	// オープン
	if len(k.table[from.Row]) > 0 {
		k.table[from.Row][from.Col-1].Open = true
	}
	return nil
}
