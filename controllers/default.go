package controllers

import (
	"sort"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

//controller functions
type MainController struct {
	beego.Controller
}

//controller structures
type Lista struct {
	Id      int    `form:"-"`
	Dados   string `form:"dados"`
	Tamanho int    `form:"tamanho"`
	Chave   int    `form:"chave"`
}

//request functions
func (c *MainController) Get() {
	var dados string
	var tamanho string
	var chave int
	//pega os dados por querystring
	c.Ctx.Input.Bind(&dados, "dados")     //id ==123
	c.Ctx.Input.Bind(&tamanho, "tamanho") //id ==123
	c.Ctx.Input.Bind(&chave, "chave")     //id ==123
	//variaveis do template
	listaCompleta := ConvStringArray(dados, tamanho)
	c.Layout = "helloWeb.html"
	c.TplName = "helloWeb.html"
	c.Data["ListaCompleta"] = listaCompleta
	c.Data["ChaveBuscada"] = BinarySearch(chave, listaCompleta)
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Style"] = "style.html"
}
func (c *MainController) Post() {
	l := Lista{}

	if err := c.ParseForm(&l); err != nil {
		//handle error
	}

}

//processing functions
func ConvStringArray(A string, N string) []int { //converte string em arrays
	a := strings.Split(A, " ")
	n, _ := strconv.Atoi(N) // int 32bit
	b := make([]int, n)
	if len(a) <= n {
		for i, v := range a {
			b[i], _ = strconv.Atoi(v)
		}
	} else {
		return []int{}
	}
	return b
}

func BinarySearch(agulha int, palheiro []int) bool { //busca binaria
	minimo := 0
	maximo := len(palheiro) - 1
	//sorting array
	sort.Ints(palheiro)

	//loop de busca
	for minimo <= maximo {
		media := (minimo + maximo) / 2

		if palheiro[media] < agulha {
			minimo = media + 1
		} else {
			maximo = media - 1
		}
	}

	if minimo == len(palheiro) || palheiro[minimo] != agulha {
		return false
	}

	return true
}
