package main

import (
	"time"
)

// TblActa [...]
type TblActa struct {
	NumeroCorrelativo int        `gorm:"primaryKey;column:numero_correlativo"`
	Notas             string     `gorm:"column:notas"`
	FechaDeEmision    time.Time  `gorm:"column:fecha_de_emision"`
	Estado            string     `gorm:"column:estado"`
	CodigoCurso       string     `gorm:"column:codigo_curso"`
	TblCurso          TblCurso   `gorm:"joinForeignKey:codigo_curso;foreignKey:codigo_curso;references:CodigoCurso"`
	NumeroDeEmpleado  string     `gorm:"column:numero_de_empleado"`
	TblDocente        TblDocente `gorm:"joinForeignKey:numero_de_empleado;foreignKey:numero_de_empleado;references:NumeroDeEmpleado"`
}

// TblAlumno [...]
type TblAlumno struct {
	NumeroDeMatricula string    `gorm:"primaryKey;column:numero_de_matricula"`
	Dni               string    `gorm:"column:dni"`
	Apellido          string    `gorm:"column:apellido"`
	Nombre            string    `gorm:"column:nombre"`
	Direccion         string    `gorm:"column:direccion"`
	Telefono          string    `gorm:"column:telefono"`
	FechaDeNacimiento time.Time `gorm:"column:fecha_de_nacimiento"`
	FechaDeIngreso    time.Time `gorm:"column:fecha_de_ingreso"`
	FechaDeEgreso     time.Time `gorm:"column:fecha_de_egreso"`
}

// TblAlumnoCarrera [...]
type TblAlumnoCarrera struct {
	ID                int        `gorm:"primaryKey;column:id"`
	NumeroDeMatricula string     `gorm:"column:numero_de_matricula"`
	TblAlumno         TblAlumno  `gorm:"joinForeignKey:numero_de_matricula;foreignKey:numero_de_matricula;references:NumeroDeMatricula"`
	CodigoDeCarrera   string     `gorm:"column:codigo_de_carrera"`
	TblCarrera        TblCarrera `gorm:"joinForeignKey:codigo_de_carrera;foreignKey:codigo_de_carrera;references:CodigoDeCarrera"`
}

// TblAlumnoMateria [...]
type TblAlumnoMateria struct {
	ID                int        `gorm:"primaryKey;column:id"`
	NumeroDeMatricula string     `gorm:"column:numero_de_matricula"`
	TblAlumno         TblAlumno  `gorm:"joinForeignKey:numero_de_matricula;foreignKey:numero_de_matricula;references:NumeroDeMatricula"`
	CodigoDeMateria   string     `gorm:"column:codigo_de_materia"`
	TblMateria        TblMateria `gorm:"joinForeignKey:codigo_de_materia;foreignKey:codigo_de_materia;references:CodigoDeMateria"`
}

// TblAlumnoPlanDeEstudio [...]
type TblAlumnoPlanDeEstudio struct {
	ID                int              `gorm:"primaryKey;column:id"`
	NumeroDeMatricula string           `gorm:"column:numero_de_matricula"`
	TblAlumno         TblAlumno        `gorm:"joinForeignKey:numero_de_matricula;foreignKey:numero_de_matricula;references:NumeroDeMatricula"`
	CodigoDePlan      string           `gorm:"column:codigo_de_plan"`
	TblPlanDeEstudio  TblPlanDeEstudio `gorm:"joinForeignKey:codigo_de_plan;foreignKey:codigo_de_plan;references:CodigoDePlan"`
}

// TblAula [...]
type TblAula struct {
	NumeroAula string  `gorm:"primaryKey;column:numero_aula"`
	Capacidad  int     `gorm:"column:capacidad"`
	CodigoSede string  `gorm:"column:codigo_sede"`
	TblSede    TblSede `gorm:"joinForeignKey:codigo_sede;foreignKey:codigo_sede;references:CodigoSede"`
}

// TblCarrera [...]
type TblCarrera struct {
	CodigoDeCarrera string `gorm:"primaryKey;column:codigo_de_carrera"`
	Nombre          string `gorm:"column:nombre"`
}

// TblCarreraPlanDeEstudio [...]
type TblCarreraPlanDeEstudio struct {
	ID               int              `gorm:"primaryKey;column:id"`
	CodigoDeCarrera  string           `gorm:"column:codigo_de_carrera"`
	TblCarrera       TblCarrera       `gorm:"joinForeignKey:codigo_de_carrera;foreignKey:codigo_de_carrera;references:CodigoDeCarrera"`
	CodigoDePlan     string           `gorm:"column:codigo_de_plan"`
	TblPlanDeEstudio TblPlanDeEstudio `gorm:"joinForeignKey:codigo_de_plan;foreignKey:codigo_de_plan;references:CodigoDePlan"`
}

// TblCuatrimestre [...]
type TblCuatrimestre struct {
	CodigoCuatrimestre   string    `gorm:"primaryKey;column:codigo_cuatrimestre"`
	FechaDeInicio        time.Time `gorm:"column:fecha_de_inicio"`
	FechaDeFin           time.Time `gorm:"column:fecha_de_fin"`
	PeriodoDeInscripcion time.Time `gorm:"column:periodo_de_inscripcion"`
	PeriodoDeExamenes    time.Time `gorm:"column:periodo_de_examenes"`
	ExamenesFinales      time.Time `gorm:"column:examenes_finales"`
}

// TblCuota [...]
type TblCuota struct {
	IDentificadorCuota    int       `gorm:"primaryKey;column:identificador_cuota"`
	AnioYMesDeVencimiento time.Time `gorm:"column:anio_y_mes_de_vencimiento"`
	Arancel               float64   `gorm:"column:arancel"`
}

// TblCurso [...]
type TblCurso struct {
	CodigoCurso             string `gorm:"primaryKey;column:codigo_curso"`
	CantidadMaximaDeAlumnos int    `gorm:"column:cantidad_maxima_de_alumnos"`
}

// TblCursoCuatrimestre [...]
type TblCursoCuatrimestre struct {
	ID                 int             `gorm:"primaryKey;column:id"`
	CodigoCurso        string          `gorm:"column:codigo_curso"`
	TblCurso           TblCurso        `gorm:"joinForeignKey:codigo_curso;foreignKey:codigo_curso;references:CodigoCurso"`
	CodigoCuatrimestre string          `gorm:"column:codigo_cuatrimestre"`
	TblCuatrimestre    TblCuatrimestre `gorm:"joinForeignKey:codigo_cuatrimestre;foreignKey:codigo_cuatrimestre;references:CodigoCuatrimestre"`
}

// TblCursoDocente [...]
type TblCursoDocente struct {
	ID               int        `gorm:"primaryKey;column:id"`
	CodigoCurso      string     `gorm:"column:codigo_curso"`
	TblCurso         TblCurso   `gorm:"joinForeignKey:codigo_curso;foreignKey:codigo_curso;references:CodigoCurso"`
	NumeroDeEmpleado string     `gorm:"column:numero_de_empleado"`
	TblDocente       TblDocente `gorm:"joinForeignKey:numero_de_empleado;foreignKey:numero_de_empleado;references:NumeroDeEmpleado"`
}

// TblCursoHorario [...]
type TblCursoHorario struct {
	ID          int        `gorm:"primaryKey;column:id"`
	CodigoCurso string     `gorm:"column:codigo_curso"`
	TblCurso    TblCurso   `gorm:"joinForeignKey:codigo_curso;foreignKey:codigo_curso;references:CodigoCurso"`
	IDHorario   int        `gorm:"column:id_horario"`
	TblHorario  TblHorario `gorm:"joinForeignKey:id_horario;foreignKey:id_horario;references:IDHorario"`
}

type TblDocente struct {
	NumeroDeEmpleado  string    `gorm:"primaryKey;column:numero_de_empleado"`
	Dni               string    `gorm:"column:dni"`
	Apellido          string    `gorm:"column:apellido"`
	Nombre            string    `gorm:"column:nombre"`
	Direccion         string    `gorm:"column:direccion"`
	Telefono          string    `gorm:"column:telefono"`
	FechaDeNacimiento time.Time `gorm:"column:fecha_de_nacimiento"`
	FechaDeIngreso    time.Time `gorm:"column:fecha_de_ingreso"`
	FechaDeEgreso     time.Time `gorm:"column:fecha_de_egreso"`
	Categoria         string    `gorm:"column:categoria"`
	Titulos           string    `gorm:"column:titulos"`
}

// TblFactura [...]
type TblFactura struct {
	IDentificador      int       `gorm:"primaryKey;column:identificador"`
	FechaDeEmision     time.Time `gorm:"column:fecha_de_emision"`
	FechaDeVencimiento time.Time `gorm:"column:fecha_de_vencimiento"`
	ImporteTotal       float64   `gorm:"column:importe_total"`
	NumeroDeMatricula  string    `gorm:"column:numero_de_matricula"`
	TblAlumno          TblAlumno `gorm:"joinForeignKey:numero_de_matricula;foreignKey:numero_de_matricula;references:NumeroDeMatricula"`
}

// TblHorario [...]
type TblHorario struct {
	IDHorario    int       `gorm:"primaryKey;column:id_horario"`
	Dia          string    `gorm:"column:dia"`
	HoraDeInicio time.Time `gorm:"column:hora_de_inicio"`
	HoraDeFin    time.Time `gorm:"column:hora_de_fin"`
	NumeroAula   string    `gorm:"column:numero_aula"`
	TblAula      TblAula   `gorm:"joinForeignKey:numero_aula;foreignKey:numero_aula;references:NumeroAula"`
}

// TblInscripcion [...]
type TblInscripcion struct {
	IDentificadorInscripcion int             `gorm:"primaryKey;column:identificador_inscripcion"`
	Fecha                    time.Time       `gorm:"column:fecha"`
	Hora                     time.Time       `gorm:"column:hora"`
	NumeroDeMatricula        string          `gorm:"column:numero_de_matricula"`
	TblAlumno                TblAlumno       `gorm:"joinForeignKey:numero_de_matricula;foreignKey:numero_de_matricula;references:NumeroDeMatricula"`
	CodigoCurso              string          `gorm:"column:codigo_curso"`
	TblCurso                 TblCurso        `gorm:"joinForeignKey:codigo_curso;foreignKey:codigo_curso;references:CodigoCurso"`
	CodigoCuatrimestre       string          `gorm:"column:codigo_cuatrimestre"`
	TblCuatrimestre          TblCuatrimestre `gorm:"joinForeignKey:codigo_cuatrimestre;foreignKey:codigo_cuatrimestre;references:CodigoCuatrimestre"`
}

type TblItemFactura struct {
	IDItemFactura  int        `gorm:"primaryKey;column:id_item_factura"`
	Recargo        float64    `gorm:"column:recargo"`
	MesesFaltantes int        `gorm:"column:meses_faltantes"`
	IDentificador  int        `gorm:"column:identificador"`
	TblFactura     TblFactura `gorm:"joinForeignKey:identificador;foreignKey:identificador;references:IDentificador"`
}

// TblMateria [...]
type TblMateria struct {
	CodigoDeMateria string `gorm:"primaryKey;column:codigo_de_materia"`
	Nombre          string `gorm:"column:nombre"`
	Contenidos      string `gorm:"column:contenidos"`
}

// TblMateriaCorrelativa [...]
type TblMateriaCorrelativa struct {
	ID                         int        `gorm:"primaryKey;column:id"`
	CodigoDeMateria            string     `gorm:"column:codigo_de_materia"`
	TblMateria                 TblMateria `gorm:"joinForeignKey:codigo_de_materia;foreignKey:codigo_de_materia;references:CodigoDeMateria"`
	CodigoDeMateriaCorrelativa string     `gorm:"column:codigo_de_materia_correlativa"`
}

// TblPlanDeEstudio [...]
type TblPlanDeEstudio struct {
	CodigoDePlan             string    `gorm:"primaryKey;column:codigo_de_plan"`
	FechaDeEntradaEnVigencia time.Time `gorm:"column:fecha_de_entrada_en_vigencia"`
	FechaDeSalidaDeVigencia  time.Time `gorm:"column:fecha_de_salida_de_vigencia"`
}

// TblPlanDeEstudioMateria [...]
type TblPlanDeEstudioMateria struct {
	ID               int              `gorm:"primaryKey;column:id"`
	CodigoDePlan     string           `gorm:"column:codigo_de_plan"`
	TblPlanDeEstudio TblPlanDeEstudio `gorm:"joinForeignKey:codigo_de_plan;foreignKey:codigo_de_plan;references:CodigoDePlan"`
	CodigoDeMateria  string           `gorm:"column:codigo_de_materia"`
	TblMateria       TblMateria       `gorm:"joinForeignKey:codigo_de_materia;foreignKey:codigo_de_materia;references:CodigoDeMateria"`
}

// TblSede [...]
type TblSede struct {
	CodigoSede string `gorm:"primaryKey;column:codigo_sede"`
	Nombre     string `gorm:"column:nombre"`
	Direccion  string `gorm:"column:direccion"`
	Pais       string `gorm:"column:pais"`
	Telefono   string `gorm:"column:telefono"`
}

// TblTransaccion [...]
type TblTransaccion struct {
	IDentificador     int       `gorm:"primaryKey;column:identificador"`
	FechaDeOperacion  time.Time `gorm:"column:fecha_de_operacion"`
	ImporteAbonado    float64   `gorm:"column:importe_abonado"`
	NumeroDeMatricula string    `gorm:"column:numero_de_matricula"`
	TblAlumno         TblAlumno `gorm:"joinForeignKey:numero_de_matricula;foreignKey:numero_de_matricula;references:NumeroDeMatricula"`
}
