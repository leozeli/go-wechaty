package _interface

import (
	wechatyPuppet "github.com/leozeli/go-wechaty/wechaty-puppet"
)

// IAccessory accessory interface
type IAccessory interface {
  GetPuppet() wechatyPuppet.IPuppetAbstract

  GetWechaty() IWechaty
}
