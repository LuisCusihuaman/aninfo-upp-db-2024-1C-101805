package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/jaswdr/faker"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func generateUniqueIntID(usedIntIDs map[int]bool) int {
	id := rand.Intn(1000) + 1 // Generate a random ID between 1 and 1000
	for usedIntIDs[id] {
		id = rand.Intn(1000) + 1
	}
	usedIntIDs[id] = true
	return id
}

func getRandomElement(slice []string) string {
	if len(slice) == 0 {
		return ""
	}
	return slice[rand.Intn(len(slice))]
}

func getRandomIntElement(slice []string) int {
	if len(slice) == 0 {
		return 0
	}
	val, _ := strconv.Atoi(slice[rand.Intn(len(slice))])
	return val
}

func seedInitialTables(db *gorm.DB, fake faker.Faker, primaryKeys map[string][]string, usedIntIDs map[string]map[int]bool) {
	// Seed Alumnos
	for i := 0; i < 10; i++ {
		id := fake.UUID().V4()
		alumno := TblAlumno{
			NumeroDeMatricula: id,
			Dni:               fake.Person().SSN(),
			Apellido:          fake.Person().LastName(),
			Nombre:            fake.Person().FirstName(),
			Direccion:         fake.Address().StreetAddress(),
			Telefono:          fmt.Sprintf("%s", fake.Phone().Number()),
			FechaDeNacimiento: fake.Time().Time(time.Now().AddDate(-25, 0, 0)),
			FechaDeIngreso:    fake.Time().Time(time.Now().AddDate(-4, 0, 0)),
			FechaDeEgreso:     fake.Time().Time(time.Now()),
		}
		db.Create(&alumno)
		primaryKeys["TblAlumno"] = append(primaryKeys["TblAlumno"], id)
	}

	// Seed Carreras
	for i := 0; i < 3; i++ {
		id := fake.UUID().V4()
		carrera := TblCarrera{
			CodigoDeCarrera: id,
			Nombre:          fake.Company().Name(),
		}
		db.Create(&carrera)
		primaryKeys["TblCarrera"] = append(primaryKeys["TblCarrera"], id)
	}

	// Seed Materias
	for i := 0; i < 5; i++ {
		id := fake.UUID().V4()
		materia := TblMateria{
			CodigoDeMateria: id,
			Nombre:          fake.Lorem().Word(),
			Contenidos:      fake.Lorem().Sentence(10),
		}
		db.Create(&materia)
		primaryKeys["TblMateria"] = append(primaryKeys["TblMateria"], id)
	}

	// Seed PlanDeEstudios
	for i := 0; i < 2; i++ {
		id := fake.UUID().V4()
		planDeEstudio := TblPlanDeEstudio{
			CodigoDePlan:             id,
			FechaDeEntradaEnVigencia: fake.Time().Time(time.Now().AddDate(-10, 0, 0)),
			FechaDeSalidaDeVigencia:  fake.Time().Time(time.Now().AddDate(10, 0, 0)),
		}
		db.Create(&planDeEstudio)
		primaryKeys["TblPlanDeEstudio"] = append(primaryKeys["TblPlanDeEstudio"], id)
	}

	// Seed Docentes
	for i := 0; i < 5; i++ {
		id := fake.UUID().V4()
		docente := TblDocente{
			NumeroDeEmpleado:  id,
			Dni:               fake.Person().SSN(),
			Apellido:          fake.Person().LastName(),
			Nombre:            fake.Person().FirstName(),
			Direccion:         fake.Address().StreetAddress(),
			Telefono:          fmt.Sprintf("%s", fake.Phone().Number()),
			FechaDeNacimiento: fake.Time().Time(time.Now().AddDate(-35, 0, 0)),
			FechaDeIngreso:    fake.Time().Time(time.Now().AddDate(-10, 0, 0)),
			FechaDeEgreso:     fake.Time().Time(time.Now()),
			Categoria:         fake.Company().Name(),
			Titulos:           fake.Company().CatchPhrase(),
		}
		db.Create(&docente)
		primaryKeys["TblDocente"] = append(primaryKeys["TblDocente"], id)
	}

	// Seed Sedes
	for i := 0; i < 2; i++ {
		id := fake.UUID().V4()
		sede := TblSede{
			CodigoSede: id,
			Nombre:     fake.Company().Name(),
			Direccion:  fake.Address().StreetAddress(),
			Pais:       fake.Address().Country(),
			Telefono:   fmt.Sprintf("%s", fake.Phone().Number()),
		}
		db.Create(&sede)
		primaryKeys["TblSede"] = append(primaryKeys["TblSede"], id)
	}

	// Seed Aulas
	for i := 0; i < 5; i++ {
		id := fake.UUID().V4()
		aula := TblAula{
			NumeroAula: id,
			Capacidad:  fake.IntBetween(10, 50),
			CodigoSede: getRandomElement(primaryKeys["TblSede"]),
		}
		db.Create(&aula)
		primaryKeys["TblAula"] = append(primaryKeys["TblAula"], id)
	}

	// Seed Cursos
	for i := 0; i < 5; i++ {
		id := fake.UUID().V4()
		curso := TblCurso{
			CodigoCurso:             id,
			CantidadMaximaDeAlumnos: fake.IntBetween(20, 100),
		}
		db.Create(&curso)
		primaryKeys["TblCurso"] = append(primaryKeys["TblCurso"], id)
	}

	// Seed Cuatrimestres
	for i := 0; i < 3; i++ {
		id := fake.UUID().V4()
		cuatrimestre := TblCuatrimestre{
			CodigoCuatrimestre:   id,
			FechaDeInicio:        fake.Time().Time(time.Now().AddDate(-1, 0, 0)),
			FechaDeFin:           fake.Time().Time(time.Now().AddDate(0, 3, 0)),
			PeriodoDeInscripcion: fake.Time().Time(time.Now().AddDate(-1, 0, 0)),
			PeriodoDeExamenes:    fake.Time().Time(time.Now().AddDate(0, 3, 0)),
			ExamenesFinales:      fake.Time().Time(time.Now().AddDate(0, 3, 0)),
		}
		db.Create(&cuatrimestre)
		primaryKeys["TblCuatrimestre"] = append(primaryKeys["TblCuatrimestre"], id)
	}

	// Seed Horarios
	usedIntIDs["TblHorario"] = make(map[int]bool)
	for i := 0; i < 10; i++ {
		id := generateUniqueIntID(usedIntIDs["TblHorario"])
		horario := TblHorario{
			IDHorario:    id,
			Dia:          fake.Time().DayOfWeek().String(),
			HoraDeInicio: fake.Time().Time(time.Now().Add(-time.Hour)),
			HoraDeFin:    fake.Time().Time(time.Now()),
			NumeroAula:   getRandomElement(primaryKeys["TblAula"]),
		}
		db.Create(&horario)
		primaryKeys["TblHorario"] = append(primaryKeys["TblHorario"], strconv.Itoa(id))
	}

	// Seed Facturas
	usedIntIDs["TblFactura"] = make(map[int]bool)
	for i := 0; i < 10; i++ {
		id := generateUniqueIntID(usedIntIDs["TblFactura"])
		factura := TblFactura{
			IDentificador:      id,
			FechaDeEmision:     fake.Time().Time(time.Now().AddDate(-1, 0, 0)),
			FechaDeVencimiento: fake.Time().Time(time.Now().AddDate(0, 1, 0)),
			ImporteTotal:       fake.Float64(2, 1000, 5000),
			NumeroDeMatricula:  getRandomElement(primaryKeys["TblAlumno"]),
		}
		db.Create(&factura)
		primaryKeys["TblFactura"] = append(primaryKeys["TblFactura"], strconv.Itoa(id))
	}

	// Seed Cuotas
	usedIntIDs["TblCuota"] = make(map[int]bool)
	for i := 0; i < 10; i++ {
		id := generateUniqueIntID(usedIntIDs["TblCuota"])
		cuota := TblCuota{
			IDentificadorCuota:   id,
			AñoYMesDeVencimiento: fake.Time().Time(time.Now().AddDate(0, 1, 0)),
			Arancel:              fake.Float64(2, 100, 1000),
		}
		db.Create(&cuota)
		primaryKeys["TblCuota"] = append(primaryKeys["TblCuota"], strconv.Itoa(id))
	}
}

func seedRelatedTables(db *gorm.DB, fake faker.Faker, primaryKeys map[string][]string, usedIntIDs map[string]map[int]bool) {
	// Seed ItemFactura (asignar items a facturas)
	for i := 0; i < 10; i++ {
		id := generateUniqueIntID(usedIntIDs["TblItemFactura"])
		itemFactura := TblItemFactura{
			IDItemFactura:  id,
			Recargo:        fake.Float64(2, 10, 100),
			MesesFaltantes: fake.IntBetween(1, 12),
			IDentificador:  getRandomIntElement(primaryKeys["TblFactura"]),
		}
		db.Create(&itemFactura)
	}

	// Seed CursoHorario (asignar horarios a cursos)
	for i := 0; i < 10; i++ {
		id := generateUniqueIntID(usedIntIDs["TblCursoHorario"])
		cursoHorario := TblCursoHorario{
			ID:          id,
			CodigoCurso: getRandomElement(primaryKeys["TblCurso"]),
			IDHorario:   getRandomIntElement(primaryKeys["TblHorario"]),
		}
		db.Create(&cursoHorario)
	}

	// Seed CarreraPlanDeEstudio (asignar planes de estudio a carreras)
	for i := 0; i < 10; i++ {
		id := generateUniqueIntID(usedIntIDs["TblCarreraPlanDeEstudio"])
		carreraPlanDeEstudio := TblCarreraPlanDeEstudio{
			ID:              id,
			CodigoDeCarrera: getRandomElement(primaryKeys["TblCarrera"]),
			CodigoDePlan:    getRandomElement(primaryKeys["TblPlanDeEstudio"]),
		}
		db.Create(&carreraPlanDeEstudio)
	}

	// Seed AlumnoPlanDeEstudio (asignar planes de estudio a alumnos)
	for i := 0; i < 10; i++ {
		id := generateUniqueIntID(usedIntIDs["TblAlumnoPlanDeEstudio"])
		alumnoPlanDeEstudio := TblAlumnoPlanDeEstudio{
			ID:                id,
			NumeroDeMatricula: getRandomElement(primaryKeys["TblAlumno"]),
			CodigoDePlan:      getRandomElement(primaryKeys["TblPlanDeEstudio"]),
		}
		db.Create(&alumnoPlanDeEstudio)
	}

	// Seed AlumnoMateria (asignar materias a alumnos)
	for i := 0; i < 10; i++ {
		id := generateUniqueIntID(usedIntIDs["TblAlumnoMateria"])
		alumnoMateria := TblAlumnoMateria{
			ID:                id,
			NumeroDeMatricula: getRandomElement(primaryKeys["TblAlumno"]),
			CodigoDeMateria:   getRandomElement(primaryKeys["TblMateria"]),
		}
		db.Create(&alumnoMateria)
	}

	// Seed AlumnoCarrera (asignar carreras a alumnos)
	for i := 0; i < 10; i++ {
		id := generateUniqueIntID(usedIntIDs["TblAlumnoCarrera"])
		alumnoCarrera := TblAlumnoCarrera{
			ID:                id,
			NumeroDeMatricula: getRandomElement(primaryKeys["TblAlumno"]),
			CodigoDeCarrera:   getRandomElement(primaryKeys["TblCarrera"]),
		}
		db.Create(&alumnoCarrera)
	}

	// Seed CursoDocente (asignar docentes a cursos)
	for i := 0; i < 10; i++ {
		id := generateUniqueIntID(usedIntIDs["TblCursoDocente"])
		cursoDocente := TblCursoDocente{
			ID:               id,
			CodigoCurso:      getRandomElement(primaryKeys["TblCurso"]),
			NumeroDeEmpleado: getRandomElement(primaryKeys["TblDocente"]),
		}
		db.Create(&cursoDocente)
	}

	// Seed CursoCuatrimestre (asignar cuatrimestres a cursos)
	for i := 0; i < 10; i++ {
		id := generateUniqueIntID(usedIntIDs["TblCursoCuatrimestre"])
		cursoCuatrimestre := TblCursoCuatrimestre{
			ID:                 id,
			CodigoCurso:        getRandomElement(primaryKeys["TblCurso"]),
			CodigoCuatrimestre: getRandomElement(primaryKeys["TblCuatrimestre"]),
		}
		db.Create(&cursoCuatrimestre)
	}

	// Seed MateriaCorrelativa (asignar materias correlativas a materias)
	for i := 0; i < 10; i++ {
		id := generateUniqueIntID(usedIntIDs["TblMateriaCorrelativa"])
		materiaCorrelativa := TblMateriaCorrelativa{
			ID:                         id,
			CodigoDeMateria:            getRandomElement(primaryKeys["TblMateria"]),
			CodigoDeMateriaCorrelativa: getRandomElement(primaryKeys["TblMateria"]),
		}
		db.Create(&materiaCorrelativa)
	}

	// Seed Inscripcion (asignar inscripciones de alumnos a cursos en un cuatrimestre)
	for i := 0; i < 10; i++ {
		id := generateUniqueIntID(usedIntIDs["TblInscripcion"])
		inscripcion := TblInscripcion{
			IDentificadorInscripcion: id,
			Fecha:                    fake.Time().Time(time.Now().AddDate(-1, 0, 0)),
			Hora:                     fake.Time().Time(time.Now()),
			NumeroDeMatricula:        getRandomElement(primaryKeys["TblAlumno"]),
			CodigoCurso:              getRandomElement(primaryKeys["TblCurso"]),
			CodigoCuatrimestre:       getRandomElement(primaryKeys["TblCuatrimestre"]),
		}
		db.Create(&inscripcion)
	}

	// Seed PlanDeEstudioMateria (asignar materias a planes de estudio)
	for i := 0; i < 10; i++ {
		id := generateUniqueIntID(usedIntIDs["TblPlanDeEstudioMateria"])
		planDeEstudioMateria := TblPlanDeEstudioMateria{
			ID:              id,
			CodigoDePlan:    getRandomElement(primaryKeys["TblPlanDeEstudio"]),
			CodigoDeMateria: getRandomElement(primaryKeys["TblMateria"]),
		}
		db.Create(&planDeEstudioMateria)
	}

	// Seed Transaccion (asignar transacciones a alumnos)
	for i := 0; i < 10; i++ {
		id := generateUniqueIntID(usedIntIDs["TblTransaccion"])
		transaccion := TblTransaccion{
			IDentificador:     id,
			FechaDeOperacion:  fake.Time().Time(time.Now().AddDate(-1, 0, 0)),
			ImporteAbonado:    fake.Float64(2, 100, 1000),
			NumeroDeMatricula: getRandomElement(primaryKeys["TblAlumno"]),
		}
		db.Create(&transaccion)
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {
	dbHost := getEnv("DB_HOST", "127.0.0.1")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "")
	dbName := getEnv("DB_NAME", "psa")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fake := faker.New()
	rand.New(rand.NewSource(time.Now().UnixNano()))

	primaryKeys := make(map[string][]string)
	usedIntIDs := make(map[string]map[int]bool)

	for _, tableName := range []string{"TblHorario", "TblFactura", "TblCuota", "TblItemFactura", "TblCursoHorario", "TblCarreraPlanDeEstudio", "TblAlumnoPlanDeEstudio", "TblAlumnoMateria", "TblAlumnoCarrera", "TblCursoDocente", "TblCursoCuatrimestre", "TblMateriaCorrelativa", "TblInscripcion", "TblPlanDeEstudioMateria", "TblTransaccion"} {
		usedIntIDs[tableName] = make(map[int]bool)
	}

	seedInitialTables(db, fake, primaryKeys, usedIntIDs)
	seedRelatedTables(db, fake, primaryKeys, usedIntIDs)
}
