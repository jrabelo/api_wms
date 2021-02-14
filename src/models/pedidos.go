package Model

import (
	"api_wms/src/database"
	"log"
)

type Pedido struct {
	PED_CLI   string `json:"PED_CLI"`
	ORDER_SEP string `json:"ORDER_SEP"`
}

func CarregaTodosPedidos(filial int) *Pedido {
	db := database.Connect()

	pedido := Pedido{}

	strSql := `SELECT ARQT221.OBS_EXPED AS PED_CLI
					, ARQT221.DOCU_EXPED AS ORDER_SEP
				 FROM ARQT221
			LEFT JOIN ARQT5035 ON ARQT5035.NR_PEDIDO = ARQT221.PED_EXPED AND ARQT5035.FILIAL = ARQT221.FILI_EXPED
				WHERE ARQT221.FILI_EXPED = :0
				  AND ARQT221.STAT_EXPED = 0
				  AND ARQT221.TIPO_EXPED = 2
				  AND ARQT5035.DT_IN_SEPARACAO IS NULL
				  AND ARQT5035.DT_FIM_SEPARACAO IS NULL
			 ORDER BY ARQT221.DOCU_EXPED ASC`

	row := db.QueryRow(strSql, filial).Scan(
		&pedido.PED_CLI,
		&pedido.ORDER_SEP,
	)

	if row != nil {
		log.Fatal(row)
	}

	defer db.Close()

	return &pedido
}
