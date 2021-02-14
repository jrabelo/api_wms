package Model

import (
	"api_wms/src/database"
	"log"
)

type Usuarios struct {
	Usuarios []*Usuarios `json:"usuarios"`
}

type Usuario struct {
	ID        int    `json:"id"`
	ID_FILIAL int    `json:"id_filial"`
	Filial    string `json:"filial"`
	Nome      string `json:"nome"`
	Login     string `json:"login"`
	Senha     string `json:"senha"`
}

type RequestLogin struct {
	ID        int    `json:"id"`
	ID_FILIAL int    `json:"id_filial"`
	Filial    string `json:"filial"`
	Nome      string `json:"nome"`
	Login     string `json:"login"`
	Pass      string `json:"pass"`
}

type RequestApp struct {
	ID        int    `json:"id_user"`
	ID_FILIAL int    `json:"id_empresa"`
	Jwt       string `json:"jwt"`
}

func AutenticarUsuarios(login, pass string) *Usuario {
	db := database.Connect()
	usuario := Usuario{}

	strSql := "SELECT ID, ID_FILIAL, NOME_SITE, NOME, LOGIN, SENHA "
	strSql += "FROM ARQT5008 "
	strSql += "INNER JOIN ARQT100 ON NCODI_EMP = ID_FILIAL "
	strSql += "WHERE LOGIN = :0 AND SENHA = :1 AND STATUS = 1"
	row := db.QueryRow(strSql, login, pass).Scan(
		&usuario.ID,
		&usuario.ID_FILIAL,
		&usuario.Filial,
		&usuario.Nome,
		&usuario.Login,
		&usuario.Senha,
	)

	if row != nil {
		log.Fatal(row)
	}

	defer db.Close()

	return &usuario
}
