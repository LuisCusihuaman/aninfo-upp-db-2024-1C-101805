CREATE TABLE `tbl_sede` (
                            `codigo_sede` varchar(255) PRIMARY KEY,
                            `nombre` varchar(255),
                            `direccion` varchar(255),
                            `pais` varchar(255),
                            `telefono` varchar(255)
);

CREATE TABLE `tbl_aula` (
                            `numero_aula` varchar(255) PRIMARY KEY,
                            `capacidad` int,
                            `codigo_sede` varchar(255),
                            FOREIGN KEY (`codigo_sede`) REFERENCES `tbl_sede` (`codigo_sede`)
);

CREATE TABLE `tbl_cuatrimestre` (
                                    `codigo_cuatrimestre` varchar(255) PRIMARY KEY,
                                    `fecha_de_inicio` date,
                                    `fecha_de_fin` date,
                                    `periodo_de_inscripcion` date,
                                    `periodo_de_examenes` date,
                                    `examenes_finales` date
);

CREATE TABLE `tbl_docente` (
                               `numero_de_empleado` varchar(255) PRIMARY KEY,
                               `dni` varchar(255),
                               `apellido` varchar(255),
                               `nombre` varchar(255),
                               `direccion` varchar(255),
                               `telefono` varchar(255),
                               `fecha_de_nacimiento` date,
                               `fecha_de_ingreso` date,
                               `fecha_de_egreso` date,
                               `categoria` varchar(255),
                               `titulos` varchar(255)
);

CREATE TABLE `tbl_materia` (
                               `codigo_de_materia` varchar(255) PRIMARY KEY,
                               `nombre` varchar(255),
                               `contenidos` text
);

CREATE TABLE `tbl_plan_de_estudio` (
                                       `codigo_de_plan` varchar(255) PRIMARY KEY,
                                       `fecha_de_entrada_en_vigencia` date,
                                       `fecha_de_salida_de_vigencia` date
);

CREATE TABLE `tbl_carrera` (
                               `codigo_de_carrera` varchar(255) PRIMARY KEY,
                               `nombre` varchar(255)
);

CREATE TABLE `tbl_alumno` (
                              `numero_de_matricula` varchar(255) PRIMARY KEY,
                              `dni` varchar(255),
                              `apellido` varchar(255),
                              `nombre` varchar(255),
                              `direccion` varchar(255),
                              `telefono` varchar(255),
                              `fecha_de_nacimiento` date,
                              `fecha_de_ingreso` date,
                              `fecha_de_egreso` date
);

CREATE TABLE `tbl_curso` (
                             `codigo_curso` varchar(255) PRIMARY KEY,
                             `cantidad_maxima_de_alumnos` int
);

CREATE TABLE `tbl_horario` (
                               `id_horario` int PRIMARY KEY,
                               `dia` varchar(255),
                               `hora_de_inicio` time,
                               `hora_de_fin` time,
                               `numero_aula` varchar(255),
                               FOREIGN KEY (`numero_aula`) REFERENCES `tbl_aula` (`numero_aula`)
);

CREATE TABLE `tbl_transaccion` (
                                   `identificador` int PRIMARY KEY,
                                   `fecha_de_operacion` date,
                                   `importe_abonado` decimal,
                                   `numero_de_matricula` varchar(255),
                                   FOREIGN KEY (`numero_de_matricula`) REFERENCES `tbl_alumno` (`numero_de_matricula`)
);

CREATE TABLE `tbl_factura` (
                               `identificador` int PRIMARY KEY,
                               `fecha_de_emision` date,
                               `fecha_de_vencimiento` date,
                               `importe_total` decimal,
                               `numero_de_matricula` varchar(255),
                               FOREIGN KEY (`numero_de_matricula`) REFERENCES `tbl_alumno` (`numero_de_matricula`)
);

CREATE TABLE `tbl_item_factura` (
                                    `id_item_factura` int PRIMARY KEY,
                                    `recargo` decimal,
                                    `meses_faltantes` int,
                                    `identificador` int,
                                    FOREIGN KEY (`identificador`) REFERENCES `tbl_factura` (`identificador`)
);

CREATE TABLE `tbl_cuota` (
                             `identificador_cuota` int PRIMARY KEY,
                             `anio_y_mes_de_vencimiento` date,
                             `arancel` decimal
);

CREATE TABLE `tbl_inscripcion` (
                                   `identificador_inscripcion` int PRIMARY KEY,
                                   `fecha` date,
                                   `hora` time,
                                   `numero_de_matricula` varchar(255),
                                   `codigo_curso` varchar(255),
                                   `codigo_cuatrimestre` varchar(255),
                                   FOREIGN KEY (`numero_de_matricula`) REFERENCES `tbl_alumno` (`numero_de_matricula`),
                                   FOREIGN KEY (`codigo_curso`) REFERENCES `tbl_curso` (`codigo_curso`),
                                   FOREIGN KEY (`codigo_cuatrimestre`) REFERENCES `tbl_cuatrimestre` (`codigo_cuatrimestre`)
);

CREATE TABLE `tbl_acta` (
                            `numero_correlativo` int PRIMARY KEY,
                            `notas` text,
                            `fecha_de_emision` date,
                            `estado` varchar(255),
                            `codigo_curso` varchar(255),
                            `numero_de_empleado` varchar(255),
                            FOREIGN KEY (`codigo_curso`) REFERENCES `tbl_curso` (`codigo_curso`),
                            FOREIGN KEY (`numero_de_empleado`) REFERENCES `tbl_docente` (`numero_de_empleado`)
);

CREATE TABLE `tbl_materia_correlativa` (
                                           `id` int PRIMARY KEY,
                                           `codigo_de_materia` varchar(255),
                                           `codigo_de_materia_correlativa` varchar(255),
                                           FOREIGN KEY (`codigo_de_materia`) REFERENCES `tbl_materia` (`codigo_de_materia`),
                                           FOREIGN KEY (`codigo_de_materia_correlativa`) REFERENCES `tbl_materia` (`codigo_de_materia`)
);

CREATE TABLE `tbl_plan_de_estudio_materia` (
                                               `id` int PRIMARY KEY,
                                               `codigo_de_plan` varchar(255),
                                               `codigo_de_materia` varchar(255),
                                               FOREIGN KEY (`codigo_de_plan`) REFERENCES `tbl_plan_de_estudio` (`codigo_de_plan`),
                                               FOREIGN KEY (`codigo_de_materia`) REFERENCES `tbl_materia` (`codigo_de_materia`)
);

CREATE TABLE `tbl_carrera_plan_de_estudio` (
                                               `id` int PRIMARY KEY,
                                               `codigo_de_carrera` varchar(255),
                                               `codigo_de_plan` varchar(255),
                                               FOREIGN KEY (`codigo_de_carrera`) REFERENCES `tbl_carrera` (`codigo_de_carrera`),
                                               FOREIGN KEY (`codigo_de_plan`) REFERENCES `tbl_plan_de_estudio` (`codigo_de_plan`)
);

CREATE TABLE `tbl_alumno_carrera` (
                                      `id` int PRIMARY KEY,
                                      `numero_de_matricula` varchar(255),
                                      `codigo_de_carrera` varchar(255),
                                      FOREIGN KEY (`numero_de_matricula`) REFERENCES `tbl_alumno` (`numero_de_matricula`),
                                      FOREIGN KEY (`codigo_de_carrera`) REFERENCES `tbl_carrera` (`codigo_de_carrera`)
);

CREATE TABLE `tbl_alumno_plan_de_estudio` (
                                              `id` int PRIMARY KEY,
                                              `numero_de_matricula` varchar(255),
                                              `codigo_de_plan` varchar(255),
                                              FOREIGN KEY (`numero_de_matricula`) REFERENCES `tbl_alumno` (`numero_de_matricula`),
                                              FOREIGN KEY (`codigo_de_plan`) REFERENCES `tbl_plan_de_estudio` (`codigo_de_plan`)
);

CREATE TABLE `tbl_alumno_materia` (
                                      `id` int PRIMARY KEY,
                                      `numero_de_matricula` varchar(255),
                                      `codigo_de_materia` varchar(255),
                                      FOREIGN KEY (`numero_de_matricula`) REFERENCES `tbl_alumno` (`numero_de_matricula`),
                                      FOREIGN KEY (`codigo_de_materia`) REFERENCES `tbl_materia` (`codigo_de_materia`)
);

CREATE TABLE `tbl_curso_horario` (
                                     `id` int PRIMARY KEY,
                                     `codigo_curso` varchar(255),
                                     `id_horario` int,
                                     FOREIGN KEY (`codigo_curso`) REFERENCES `tbl_curso` (`codigo_curso`),
                                     FOREIGN KEY (`id_horario`) REFERENCES `tbl_horario` (`id_horario`)
);

CREATE TABLE `tbl_curso_docente` (
                                     `id` int PRIMARY KEY,
                                     `codigo_curso` varchar(255),
                                     `numero_de_empleado` varchar(255),
                                     FOREIGN KEY (`codigo_curso`) REFERENCES `tbl_curso` (`codigo_curso`),
                                     FOREIGN KEY (`numero_de_empleado`) REFERENCES `tbl_docente` (`numero_de_empleado`)
);

CREATE TABLE `tbl_curso_cuatrimestre` (
                                          `id` int PRIMARY KEY,
                                          `codigo_curso` varchar(255),
                                          `codigo_cuatrimestre` varchar(255),
                                          FOREIGN KEY (`codigo_curso`) REFERENCES `tbl_curso` (`codigo_curso`),
                                          FOREIGN KEY (`codigo_cuatrimestre`) REFERENCES `tbl_cuatrimestre` (`codigo_cuatrimestre`)
);
