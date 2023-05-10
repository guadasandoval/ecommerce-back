package query

const JoinTipoDocumentoPF = "left join tipos_documentos on tipos_documentos.Id = personas_fortalecimiento.IdTipoDocumento"
const JoinProvinciaPF = "left join provincias on provincias.pro_id = personas_fortalecimiento.IdProvincia"
const JoinLocalidadPF = "left join localidades on localidades.loc_id = personas_fortalecimiento.IdLocalidad"
const JoinMunicipioPF = "left join municipios on municipios.mun_id = personas_fortalecimiento.IdMunicipio"
const JoinFuenteIngresoPF = "LEFT JOIN (SELECT pfi.IdPersonaFortalecimiento as IdPersona, (CASE WHEN pfi.IdFuente = fuente_ingreso.Id AND pfi.IdPersonaCupo IS NOT NULL THEN CONCAT('Registro aspirantes cupo TTT, ', fuente_ingreso.Fuente) WHEN pfi.IdFuente = fuente_ingreso.Id THEN fuente_ingreso.Fuente ELSE '' END) AS Fuente FROM `personas_fuente_ingreso` AS pfi LEFT JOIN fuente_ingreso on fuente_ingreso.Id = pfi.IdFuente WHERE pfi.IdPersonaFortalecimiento IS NOT NULL) AS fuenteI ON personas_fortalecimiento.Id = fuenteI.IdPersona"

const SELECTPERSONAF = "personas_fortalecimiento.Id as ID,personas_fortalecimiento.IdPersona as IdPersona, personas_fortalecimiento.Orden, personas_fortalecimiento.Apellido, personas_fortalecimiento.NombreAutopercibido, personas_fortalecimiento.FechaNac, tipos_documentos.TipoDocumento, personas_fortalecimiento.NroDocumento, personas_fortalecimiento.Email, provincias.pro_nombre as Provincia, localidades.loc_nombre as Localidad, municipios.mun_nombre as Municipio" +", "