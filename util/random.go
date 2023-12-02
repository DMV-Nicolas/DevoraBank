package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates a random int number between min and max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// RandomOwner generates a really random owner name
func RandomOwner() string {
	monsters := []string{
		"Caballo",
		"Omnitorrinco",
		"Avion",
		"Hijo de puta",
		"Avi Nicolas",
		"Luis Cortes",
		"Rodrigo Rivera",
		"Andres Franco",
		"Santiago Zamora",
		"Juan Diego",
		"Nicolas Moreno",
		"Cristiansito",
		"Juliansito",
		"Valentinita",
	}
	actions := []string{
		"violador de",
		"abusador de",
		"terapeuta de",
		"simp de",
		"desarmador de",
		"dominado por",
		"esclavizado por",
		"sexualmente abusado por",
		"atraido por",
	}
	victims := []string{
		"abuelas",
		"feministas",
		"comunistas",
		"capitalistas",
		"langostas",
		"hombres",
		"jirafas",
		"penes",
		"duendes",
	}
	str := monsters[rand.Intn(len(monsters))] + " "
	str += actions[rand.Intn(len(actions))] + " "
	str += victims[rand.Intn(len(victims))]
	return str
}

func RandomCurrency() string {
	currencies := []string{
		COP, USD, EUR,
		ARS, MXN, UYU,
		CLP, PEN, BRL,
	}
	return currencies[rand.Intn(len(currencies))]
}

func RandomMoney() float64 {
	return float64(RandomInt(0, 5000))
}

func RandomID() int64 {
	return int64(RandomInt(1, 1000))
}
