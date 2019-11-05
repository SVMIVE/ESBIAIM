package dosa

func SQLDosaTodas() string {
	return `SELECT
			cli.cod_cliente,
			cli.nombre_cliente,
			cli.is_por_lote,
			cli.cod_forma_pago,
			co.cod_tipo_dosa,
			co.id_dosa_anulada,
			co.cod_tipo_moneda --,
			--a.id_movimiento_aeronave,
			--a.fecha_apertura as fecha_lleg_vuelo,
			--vdh.id_tipo_vuelo,
			--tv.tipo_vuelo,
			--tv.descripcion as nom_tip_vuelo,
			--vdh.hora_real,
			--vdh.numero_vuelo as numero_vuelo_lleg,
			--vdh1.numero_vuelo as numero_vuelo_sal,
			--vdh.hora_real as hora_real_lleg,
			--vdh1.hora_real as hora_real_sal,
			--b.nacionalidad,
			--co.id_cobro,
			--b.matricula,
			--co.cod_facturacion,
			--co.cod_cierre,
			--co.cod_apertura,
			--co.cod_status,
			--co.fecha_apertura,
			--mod.modelo_aeronave,
			--is_especial,
			--is_anulada,
			--dosa_duplicada,
			  -- a.id_vuelos_diarios_salida,
			--co.peso_max_tonelada,
			--co.mixta,
			--co.correlativo_dm,
			--cli.impresion_dosa --,
			  -- a.id_vuelos_diarios_llegada,
			  -- fecha_inicio_asignacion as ini_asig_correa,
		  	-- fecha_fin_asignacion as fin_asig_correa,
			  -- (SELECT SUM(dcd.monto_cobro) as monto_total_dosa FROM dosas.cobros_detalles dcd WHERE dcd.id_cobro=co.id_cobro)
			FROM fids.vuelos_movimientos_aeronaves a
		INNER JOIN fids.manten_aeronaves b ON a.id_aeronave = b.id_aeronave
		LEFT JOIN fids.manten_tipo_aeronave mod ON b.id_modelo = mod.id_modelo
		LEFT JOIN dosas.cobros co ON a.id_movimiento_aeronave = co.id_movimiento_aeronave
		LEFT JOIN config.manten_clientes cli ON b.cod_cliente = cli.cod_cliente
		INNER JOIN (SELECT vd.id_movimiento_aeronave, vd.id_vuelos_diarios,
			vd.numero_vuelo, vd.id_tipo_vuelo, vd.cod_terminal, vd.id_pista,
			vd.cod_status, vd.id_aerolinea, vd.hora_itine,  vd.hora_esti,
			vd.hora_real, vd.origen_fila, vd.cod_accion, vd.id_aeropuerto_escala_2,
			vd.id_modelo, vd.id_aeropuerto_origen_destino, vd.id_aeropuerto_escala_1,
			vd.matricula, vd.is_delete, vd.cod_estacionamiento, vd.obser_extra,
			vd.id_correa, vd.id_puerta,false as is_historico FROM fids.vuelos_vuelos_diarios vd
			UNION  ALL SELECT vdh.id_movimiento_aeronave, vdh.id_vuelos_diarios,
			vdh.numero_vuelo, vdh.id_tipo_vuelo,  vdh.cod_terminal, vdh.id_pista,
			vdh.cod_status, vdh.id_aerolinea, vdh.hora_itine, vdh.hora_esti,
			vdh.hora_real, vdh.origen_fila, vdh.cod_accion, vdh.id_aeropuerto_escala_2,
			vdh.id_modelo, vdh.id_aeropuerto_origen_destino, vdh.id_aeropuerto_escala_1,
			vdh.matricula, vdh.is_delete, vdh.cod_estacionamiento, vdh.obser_extra,
			vdh.id_correa, vdh.id_puerta,vdh.is_historico FROM fids.vuelos_vuelos_diarios_historico vdh ) vdh on
			(vdh.id_vuelos_diarios=a.id_vuelos_diarios_llegada)
		LEFT JOIN (SELECT vd.id_movimiento_aeronave, vd.id_vuelos_diarios, vd.numero_vuelo,
			vd.id_tipo_vuelo, vd.cod_terminal, vd.id_pista,
			vd.cod_status, vd.id_aerolinea, vd.hora_itine,  vd.hora_esti, vd.hora_real,
			vd.origen_fila, vd.cod_accion, vd.id_aeropuerto_escala_2, vd.id_modelo,
			vd.id_aeropuerto_origen_destino, vd.id_aeropuerto_escala_1, vd.matricula,
			vd.is_delete, vd.cod_estacionamiento, vd.obser_extra, vd.id_correa,
			vd.id_puerta,false as is_historico FROM fids.vuelos_vuelos_diarios vd
			UNION  ALL SELECT vdh.id_movimiento_aeronave, vdh.id_vuelos_diarios,
			vdh.numero_vuelo, vdh.id_tipo_vuelo,  vdh.cod_terminal, vdh.id_pista,
			vdh.cod_status, vdh.id_aerolinea, vdh.hora_itine, vdh.hora_esti,
			vdh.hora_real, vdh.origen_fila, vdh.cod_accion, vdh.id_aeropuerto_escala_2,
			vdh.id_modelo, vdh.id_aeropuerto_origen_destino, vdh.id_aeropuerto_escala_1,
			 vdh.matricula, vdh.is_delete, vdh.cod_estacionamiento, vdh.obser_extra,
			 vdh.id_correa, vdh.id_puerta,vdh.is_historico FROM fids.vuelos_vuelos_diarios_historico vdh ) vdh1 on
			 (vdh1.id_vuelos_diarios=a.id_vuelos_diarios_salida)
		LEFT JOIN fids.manten_tipo_vuelo tv ON tv.id_tipo_vuelo = vdh.id_tipo_vuelo
		left JOIN fids.vuelos_asignacion_correa correa ON correa.id_vuelos_diarios = vdh.id_vuelos_diarios
		--WHERE id_cobro='39617'
		--ORDER BY id_movimiento_aeronave DESC
		LIMIT 10`
}

func SQLDosa() string {
	return ``
}
