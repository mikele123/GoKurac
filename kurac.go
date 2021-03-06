package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

var kurcevi = []string{
	"CiAgICAgICAgICAgICAgICAgICAgICAgIF8sLSUvJXwKICAgICAgICAgICAgICAgICAgICBfLC0nICAgIFwvLyVcCiAgICAgICAgICAgICAgICBfLC0nICAgICAgICBcJS98JQogICAgICAgICAgICAgIC8gLyApICAgIF9fLC0tICAvJVwKICAgICAgICAgICAgICBcX18vXywtJyUoJSAgOyAgJSklCiAgICAgICAgICAgICAgICAgICAgICAlXCUsICAgJVwKICAgICAgICAgICAgICAgICAgICAgICAgJy0tJScKCg==",
	"ICAgICAgICAgICAgICBfX18KICAgICAgICAgICAgIC8vICA3CiAgICAgICAgICAgIChfLF8vXAogICAgICAgICAgICAgXCAgICBcCiAgICAgICAgICAgICAgXCAgICBcCiAgICAgICAgICAgICAgX1wgICAgXF9fCiAgICAgICAgICAgICAoICAgXCAgICAgKQogICAgICAgICAgICAgIFxfX19cX19fLwoK",
	"4paR4paR4paE4paE4paR4paR4paR4paR4paE4paR4paR4paR4paR4paE4paR4paR4paRIArilpHilpHilpHilpHilpHilpHilojiloTilpHilpHilpHilojilpHilpHilojiloDilpHilpHilpHilpEgCuKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWgOKWiOKWhOKWkeKWgOKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkSAK4paR4paR4paR4paR4paE4paR4paR4paR4paE4paE4paI4paI4paI4paE4paR4paR4paR4paE4paE4paRIArilpHilpHilpHiloDiloDilpHilpHiloTilojilpHilpHilojilpHilojiloTilpHilpHilpHiloDiloggCuKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWiOKWkeKWkeKWkeKWkeKWkeKWkeKWiOKWkeKWkeKWkeKWkeKWkSAK4paR4paR4paR4paR4paR4paR4paR4paI4paI4paI4paI4paI4paI4paI4paI4paR4paR4paR4paR4paRIArilpHilpHilpHilpHilpHilpHilpHilojiloTiloTilpHilpHilpHilpHilojilpHilpHilpHilpHilpEgCuKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWiOKWkeKWkeKWkeKWkeKWgOKWgOKWiOKWkeKWkeKWkeKWkeKWkSAK4paR4paR4paR4paR4paR4paR4paR4paI4paA4paA4paA4paR4paE4paE4paI4paR4paR4paR4paR4paRIArilpHilpHilpHilpHilpHilpHilpHilojilpHilpHilpHilpHiloDiloDilojilpHilpHilpHilpHilpEgCuKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWiOKWgOKWgOKWgOKWkeKWhOKWhOKWiOKWkeKWkeKWkeKWkeKWkQrilpHilpHilpHilpHilpHilpHilpHilojilpHilpHilpHilpHilpHilpHilojiloTilpHilpHilpHilpEgCuKWhOKWhOKWhOKWhOKWiOKWiOKWgOKWgOKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWgOKWiOKWiOKWkeKWkSAK4paR4paE4paI4paA4paR4paA4paR4paR4paR4paR4paE4paR4paR4paR4paR4paR4paR4paI4paE4paEIAriloDiloDilojiloTiloTiloTilpHilpHilpHiloTilojilojilpHilpHilpHilpHiloTilojilpHilpEgCuKWkeKWiOKWgOKWiOKWhOKWhOKWhOKWhOKWiOKWgOKWkeKWiOKWiOKWhOKWhOKWiOKWiOKWhOKWhOKWkSAK4paR4paR4paR4paR4paA4paR4paR4paR4paA4paR4paR4paR4paA4paR4paR4paR4paR4paR4paR4paRCg==",
	"LiAgICAgICAgICAgICAgICAgICAgICAg4paE4paE4paE4paE4paE4paE4paECi4gICAgICAgICAgICAgICAgICAgICDiloTilozilpHilpHilpHilpHilpHilpHilpHilpDiloQgCi4gICAgICAgICAgICAgICAgICAgIOKWkOKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWjAouICAgICAgICAgICAgICAgICAg4paE4paI4paT4paR4paR4paR4paR4paR4paR4paR4paR4paR4paR4paT4paI4paECi4gICAgICAgICAgICAgICAg4paE4paA4paR4paR4paQ4paR4paR4paR4paR4paR4paR4paR4paR4paR4paR4paM4paR4paS4paMCi4gICAgICAgICAgICAgICDilpDilpHilpHilpHilpHilpDilpHilpHilpHilpHilpHilpHilpHilpHilpHilpHilozilpHilpHilpHilowKLiAgICAgICAgICAgICAg4paQ4paR4paR4paR4paR4paR4paQ4paR4paR4paR4paR4paR4paR4paR4paR4paR4paR4paM4paR4paR4paR4paR4paQCi4gICAgICAgICAgICAgIOKWkOKWkuKWkeKWkeKWkeKWkeKWkOKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWjOKWkeKWkeKWkuKWkuKWkAouICAgICAgICAgICAgICDilpDilpLilpHilpHilpHilpHilpDilpHilpHilpHilpHilpHilpHilpHilpHilpHilpHilozilpHilpHilpLilpAKLiAgICAgICAgICAgICAgIOKWgOKWhOKWkuKWkuKWkuKWkOKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWjOKWhOKWgAouICAgICAgICAgICAgICAgICDiloDiloDiloDilpDilpHilpHilpHilpHilpHilpHilpHilpHilpHilpHilowKLiAgICAgICAgICAgICAgICAgICAg4paQ4paR4paR4paR4paR4paR4paR4paR4paR4paR4paR4paMCi4gICAgICAgICAgICAgICAgICAgIOKWkOKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWjAouICAgICAgICAgICAgICAgICAgICDilpDilpHilpHilpHilpHilpHilpHilpHilpHilpHilpHilowKLiAgICAgICAgICAgICAgICAgICAg4paQ4paR4paR4paR4paR4paR4paR4paR4paR4paR4paR4paMCi4gICAgICAgICAgICAgICAgICAgIOKWkOKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWjAouICAgICAgICAgICAgICAgICAgICDilpDilpHilpHilpHilpHilpHilpHilpHilpHilpHilpHilowKLiAgICAgICAgICAgICAgICAgICAg4paQ4paR4paR4paR4paR4paR4paR4paR4paR4paR4paR4paMCi4gICAgICAgICAgICAgICAgICAgIOKWkOKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWkeKWjAouICAgICAgICAgICAgICAgICAgIOKWkOKWhOKWgOKWgOKWgOKWgOKWgOKWgOKWgOKWgOKWgOKWgOKWhOKWjAouICAgICAgICAgICAgICAgICAg4paQ4paS4paS4paS4paS4paS4paS4paS4paS4paS4paS4paS4paS4paS4paS4paMCi4gICAgICAgICAgICAgICAgICDilpDilpLilpLilpLilpLilpLilpLilpLilpLilpLilpLilpLilpLilpLilpLilowKLiAgICAgICAgICAgICAgICAgIOKWkOKWkuKWkuKWkuKWkuKWkuKWkuKWkuKWkuKWkuKWkuKWkuKWkuKWkuKWjAouICAgICAgICAgICAgICAgICAgICDiloDilozilpLilpLilpLiloDilpLilpLilpLilpLilpDiloAgCi4gICAgICAgICAgICAgICAgICAgICAg4paA4paA4paA4paA4paA4paA4paA4paACg==",
}

// Main
func main() {
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(kurcevi)
	decoded, err := base64.StdEncoding.DecodeString(kurcevi[n])
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%s", decoded)
}
