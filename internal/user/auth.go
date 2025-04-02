package user

import "log"

func Login(email, password string) {
	log.Default().Print(email, password)
}
