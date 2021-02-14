package Model

import (
	"api_wms/src/database"
	"log"
)

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

	strSql := `SELECT ARQT5008.ID
	                , ARQT5008.ID_FILIAL
					, ARQT100.NOME_SITE
					, ARQT5008.NOME
					, ARQT5008.LOGIN
					, ARQT5008.SENHA 
				 FROM ARQT5008
	       INNER JOIN ARQT100 ON NCODI_EMP = ID_FILIAL
	            WHERE LOGIN = :0 
				  AND SENHA = :1 
				  AND STATUS = 1`
				  
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
