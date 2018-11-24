package main

import (
	"encoding/json"
	"errors"

	"github.com/gobuffalo/packr"
)

// Jewel is struct for Jewel
type Jewel struct {
	Name           string            `json:"name"`
	ItemID         uint32            `json:"itemId"`
	EquippedItemID uint32            `json:"equippedItemId"`
	Max            int               `json:"max"`
	Locales        map[string]string `json:"locales"`
}

// JewelList is list of jewels
type JewelList []Jewel

// NewJewelList is constructor
func NewJewelList() (JewelList, error) {
	box := packr.NewBox("unify-jewel-info/dist")
	b, err := box.Find("decorations.json")
	if err != nil {
		return nil, err
	}
	var jewelList JewelList
	if err := json.Unmarshal(b, &jewelList); err != nil {
		return nil, err
	}
	return jewelList, nil
}

// FindJewelByEquippedItemID finds Jewel by EquippedItemID
func (jl JewelList) FindJewelByEquippedItemID(equippedItemID uint32) (*Jewel, error) {
	for _, jewel := range jl {
		if jewel.EquippedItemID == equippedItemID {
			return &jewel, nil
		}
	}
	return nil, errors.New("Not found")
}

// FindJewelByItemID finds Jewel by ItemID
func (jl JewelList) FindJewelByItemID(itemID uint32) (*Jewel, error) {
	for _, jewel := range jl {
		if jewel.ItemID == itemID {
			return &jewel, nil
		}
	}
	return nil, errors.New("Not found")
}
