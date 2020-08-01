package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ClosedRange(t *testing.T) {

	t.Run("下端点と上端点を変数に持つ。", func(t *testing.T) {
		t.Run("下端点と上端点を変数に持つ構造体を作成し、下端点と上端点を取得できる", func(t *testing.T) {
			cr, _ := NewClosedRange(1, 2)
			assert.NotEmpty(t, cr.lowerEndpoint)
			assert.NotEmpty(t, cr.upperEndpoint)
		})
	})

	t.Run("下端点と上端点変数の区間を表す文字列表現も返せる。", func(t *testing.T) {
		t.Run("下端点=1と上端点=2とした場合に文字列[1,2]を返す。", func(t *testing.T) {
			cr, _ := NewClosedRange(1, 2)
			assert.Equal(t, "[1,2]", cr.StringRange())
		})

		t.Run("下端点=3と上端点=8とした場合に文字列[3,8]を返す。", func(t *testing.T) {
			cr, _ := NewClosedRange(3, 8)
			assert.Equal(t, "[3,8]", cr.StringRange())
		})

		t.Run("下端点=10と上端点=10とした場合に文字列[10,10]を返す。", func(t *testing.T) {
			cr, _ := NewClosedRange(10, 10)
			assert.Equal(t, "[10,10]", cr.StringRange())
		})
	})

	t.Run("上端点より下端点が大きい閉区間を作ることはできない。", func(t *testing.T) {
		t.Run("下端点が8, 上端点が3のときはエラーを返す。", func(t *testing.T) {
			cr, err := NewClosedRange(8, 3)
			assert.Equal(t, "下端点は上端点以下にしてください。", err.Error())
			assert.Nil(t, cr)
		})
		t.Run("下端点が4, 上端点が9のときは正しいClosedRange構造体を取得できること。", func(t *testing.T) {
			cr, err := NewClosedRange(4, 9)
			assert.Equal(t, &ClosedRange{
				lowerEndpoint: 4,
				upperEndpoint: 9,
			}, cr)
			assert.Nil(t, err)
		})
	})

	t.Run(" 整数の閉区間は指定した整数を含むかどうかを判定できる。", func(t *testing.T) {
		t.Run("[2,8]の閉区間内に対して、3を指定するとtrueを返す。", func(t *testing.T) {
			cr, _ := NewClosedRange(2, 8)
			assert.True(t, cr.ContainsNumber(3))
		})
		t.Run("[2,8]の閉区間内に対して、2を指定するとtrueを返す。", func(t *testing.T) {
			cr, _ := NewClosedRange(2, 8)
			assert.True(t, cr.ContainsNumber(2))
		})
		t.Run("[2,8]の閉区間内に対して、8を指定するとtrueを返す。", func(t *testing.T) {
			cr, _ := NewClosedRange(2, 8)
			assert.True(t, cr.ContainsNumber(8))
		})
		t.Run("[2,8]の閉区間内に対して、1を指定するとfalseを返す。", func(t *testing.T) {
			cr, _ := NewClosedRange(2, 8)
			assert.False(t, cr.ContainsNumber(1))
		})
		t.Run("[2,8]の閉区間内に対して、9を指定するとfalseを返す。", func(t *testing.T) {
			cr, _ := NewClosedRange(2, 8)
			assert.False(t, cr.ContainsNumber(9))
		})
	})

	t.Run("別の閉区間と等価かどうか", func(t *testing.T) {
		t.Run("閉区間[3,8]と閉区間[3,8]を比較したとき、等価であるためTrueを返す。", func(t *testing.T) {
			cr1, _ := NewClosedRange(3, 8)
			cr2, _ := NewClosedRange(3, 8)
			assert.True(t, cr1.Equals(*cr2))
		})
		t.Run("閉区間[3,8]と閉区間[9,11]を比較したとき、等価でないためFalseを返す。", func(t *testing.T) {
			cr1, _ := NewClosedRange(3, 8)
			cr2, _ := NewClosedRange(9, 11)
			assert.False(t, cr1.Equals(*cr2))
		})
		t.Run("閉区間[3,8]と閉区間[1,2]を比較したとき、等価でないためFalseを返す。", func(t *testing.T) {
			cr1, _ := NewClosedRange(3, 8)
			cr2, _ := NewClosedRange(1, 2)
			assert.False(t, cr1.Equals(*cr2))
		})
	})

	t.Run("閉区間[3,8]は閉区間[4,6]を閉区間内に含んでいるため、Trueを返す。", func(t *testing.T) {
		cr1, _ := NewClosedRange(3, 8)
		cr2, _ := NewClosedRange(4, 6)
		assert.True(t, cr1.ContainsRange(*cr2))
	})
	t.Run("閉区間[3,8]は閉区間[1,2]を閉区間内に含んでいないため、Falseを返す。", func(t *testing.T) {
		cr1, _ := NewClosedRange(3, 8)
		cr2, _ := NewClosedRange(1, 2)
		assert.False(t, cr1.ContainsRange(*cr2))
	})
	t.Run("	閉区間[3,8]は閉区間[1,5]を閉区間内に含んでいないため、Falseを返す。", func(t *testing.T) {
		cr1, _ := NewClosedRange(3, 8)
		cr2, _ := NewClosedRange(1, 5)
		assert.False(t, cr1.ContainsRange(*cr2))
	})
	t.Run("	閉区間[3,8]は閉区間[5,11]を閉区間内に含んでいないため、Falseを返す。", func(t *testing.T) {
		cr1, _ := NewClosedRange(3, 8)
		cr2, _ := NewClosedRange(5, 11)
		assert.False(t, cr1.ContainsRange(*cr2))
	})
}
