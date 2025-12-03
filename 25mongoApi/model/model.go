package model

import (

)

type Netflix struct{
	ID int `json:"id"`
	Movie string `json:"movie"`
	Watched bool `json:"watched"`
}