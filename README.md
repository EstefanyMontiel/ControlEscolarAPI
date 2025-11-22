# 🎓 API de Control Escolar

API REST desarrollada en Go para la gestión de estudiantes, materias y calificaciones en un sistema escolar. Utiliza Gin Framework, GORM y MySQL.

## 📋 Características

- ✅ CRUD completo de estudiantes, materias y calificaciones
- ✅ Validación de datos con reglas de negocio
- ✅ Relaciones entre entidades con llaves foráneas
- ✅ Respuestas en formato JSON
- ✅ Manejo apropiado de códigos HTTP
- ✅ Documentación con Swagger/OpenAPI
- ✅ Base de datos MySQL con GORM

## 🛠️ Tecnologías

- **Lenguaje**: Go 1.21+
- **Framework Web**: Gin
- **ORM**: GORM
- **Base de Datos**: MySQL
- **Documentación**: Swagger (swaggo)

## 📦 Instalación

### Prerrequisitos

- Go 1.21 o superior
- MySQL 8.0 o superior
- Git

### Pasos de instalación

1. **Clonar el repositorio**
```bash
git clone https://github.com/EstefanyMontiel/ControlEscolarAPI.git
cd ControlEscolarAPI
```

2. **Instalar dependencias**
```bash
go mod download
```

3. **Configurar la base de datos**

Crear la base de datos en MySQL:
```sql
CREATE DATABASE control_escolar CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

4. **Configurar variables de entorno**

Crear un archivo `.env` en la raíz del proyecto:
```env
DB_USER=root
DB_PASSWORD=tu_contraseña
DB_HOST=localhost
DB_PORT=3306
DB_NAME=control_escolar
PORT=8082
```

5. **Ejecutar la aplicación**
```bash
go run main.go
```

La API estará disponible en `http://localhost:8082`

## 📚 Documentación de la API

### Swagger UI
Una vez iniciada la aplicación, accede a la documentación interactiva:
```
http://localhost:8082/swagger/index.html
```

### Estructura de Base URL
```
http://localhost:8082/api
```

---

## 🚀 Rutas de la API

### 👨‍🎓 Estudiantes

#### 1. Crear un estudiante
- **Método**: `POST`
- **Ruta**: `/api/students`
- **Descripción**: Registra un nuevo estudiante en el sistema

**Ejemplo con curl:**
```bash
curl -X POST http://localhost:8082/api/students \
  -H "Content-Type: application/json" \
  -d '{
    "name": "María García",
    "group": "5A",
    "email": "maria.garcia@escuela.com"
  }'
```

**Respuesta exitosa (201):**
```json
{
  "message": "Estudiante creado exitosamente",
  "data": {
    "student_id": 1,
    "name": "María García",
    "group": "5A",
    "email": "maria.garcia@escuela.com"
  }
}
```

#### 2. Listar todos los estudiantes
- **Método**: `GET`
- **Ruta**: `/api/students`

**Ejemplo con curl:**
```bash
curl http://localhost:8082/api/students
```

#### 3. Obtener un estudiante por ID
- **Método**: `GET`
- **Ruta**: `/api/students/:student_id`

**Ejemplo con curl:**
```bash
curl http://localhost:8082/api/students/1
```

#### 4. Actualizar un estudiante
- **Método**: `PUT`
- **Ruta**: `/api/students/:student_id`

**Ejemplo con curl:**
```bash
curl -X PUT http://localhost:8082/api/students/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "María García López",
    "group": "5B",
    "email": "maria.garcia@escuela.com"
  }'
```

#### 5. Eliminar un estudiante
- **Método**: `DELETE`
- **Ruta**: `/api/students/:student_id`

**Ejemplo con curl:**
```bash
curl -X DELETE http://localhost:8082/api/students/1
```

---

### 📚 Materias

#### 1. Crear una materia
- **Método**: `POST`
- **Ruta**: `/api/subjects`

**Ejemplo con curl:**
```bash
curl -X POST http://localhost:8082/api/subjects \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Matemáticas"
  }'
```

**Respuesta exitosa (201):**
```json
{
  "message": "Materia creada exitosamente",
  "data": {
    "subject_id": 1,
    "name": "Matemáticas"
  }
}
```

#### 2. Obtener una materia por ID
- **Método**: `GET`
- **Ruta**: `/api/subjects/:subject_id`

**Ejemplo con curl:**
```bash
curl http://localhost:8082/api/subjects/1
```

#### 3. Actualizar una materia
- **Método**: `PUT`
- **Ruta**: `/api/subjects/:subject_id`

**Ejemplo con curl:**
```bash
curl -X PUT http://localhost:8082/api/subjects/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Matemáticas Avanzadas"
  }'
```

#### 4. Eliminar una materia
- **Método**: `DELETE`
- **Ruta**: `/api/subjects/:subject_id`

**Ejemplo con curl:**
```bash
curl -X DELETE http://localhost:8082/api/subjects/1
```

---

### 📝 Calificaciones

#### 1. Crear una calificación
- **Método**: `POST`
- **Ruta**: `/api/grades`
- **Descripción**: Registra una calificación para un estudiante en una materia

**Ejemplo con curl:**
```bash
curl -X POST http://localhost:8082/api/grades \
  -H "Content-Type: application/json" \
  -d '{
    "student_id": 1,
    "subject_id": 1,
    "grade": 95.5
  }'
```

**Respuesta exitosa (201):**
```json
{
  "message": "Calificación creada exitosamente",
  "data": {
    "grade_id": 1,
    "student_id": 1,
    "subject_id": 1,
    "grade": 95.5,
    "student": {
      "student_id": 1,
      "name": "María García",
      "group": "5A",
      "email": "maria.garcia@escuela.com"
    },
    "subject": {
      "subject_id": 1,
      "name": "Matemáticas"
    }
  }
}
```

#### 2. Actualizar una calificación
- **Método**: `PUT`
- **Ruta**: `/api/grades/:grade_id`

**Ejemplo con curl:**
```bash
curl -X PUT http://localhost:8082/api/grades/1 \
  -H "Content-Type: application/json" \
  -d '{
    "grade": 98.0
  }'
```

#### 3. Eliminar una calificación
- **Método**: `DELETE`
- **Ruta**: `/api/grades/:grade_id`

**Ejemplo con curl:**
```bash
curl -X DELETE http://localhost:8082/api/grades/1
```

#### 4. Obtener calificación específica
- **Método**: `GET`
- **Ruta**: `/api/grades/:grade_id/student/:student_id`

**Ejemplo con curl:**
```bash
curl http://localhost:8082/api/grades/1/student/1
```

#### 5. Obtener todas las calificaciones de un estudiante
- **Método**: `GET`
- **Ruta**: `/api/grades/student/:student_id`

**Ejemplo con curl:**
```bash
curl http://localhost:8082/api/grades/student/1
```

---

## 📊 Ejemplos con Postman

### Importar colección

Puedes crear una colección en Postman con las siguientes peticiones:

1. **Crear Estudiante**
   - Method: POST
   - URL: `http://localhost:8082/api/students`
   - Body (JSON):
   ```json
   {
     "name": "Juan Pérez",
     "group": "3A",
     "email": "juan.perez@escuela.com"
   }
   ```

2. **Crear Materia**
   - Method: POST
   - URL: `http://localhost:8082/api/subjects`
   - Body (JSON):
   ```json
   {
     "name": "Historia"
   }
   ```

3. **Crear Calificación**
   - Method: POST
   - URL: `http://localhost:8082/api/grades`
   - Body (JSON):
   ```json
   {
     "student_id": 1,
     "subject_id": 1,
     "grade": 85.5
   }
   ```

---

## 🗃️ Estructura del Proyecto

```
ControlEscolarAPI/
├── config/           # Configuración de base de datos
│   └── database.go
├── docs/            # Documentación Swagger generada
├── handlers/        # Controladores de las rutas
│   ├── grade_handler.go
│   ├── student_handler.go
│   └── subject_handler.go
├── models/          # Modelos de datos
│   ├── dto.go
│   ├── grade.go
│   ├── student.go
│   └── subject.go
├── routes/          # Definición de rutas
│   └── routes.go
├── utils/           # Utilidades
│   └── response.go
├── .env             # Variables de entorno
├── go.mod           # Dependencias
├── go.sum
└── main.go          # Punto de entrada
```

---

## 🔍 Validaciones Implementadas

### Estudiantes
- **Nombre**: Requerido, entre 2 y 100 caracteres
- **Grupo**: Requerido, entre 1 y 10 caracteres
- **Email**: Requerido, formato válido de email, único

### Materias
- **Nombre**: Requerido, entre 2 y 100 caracteres, único

### Calificaciones
- **student_id**: Requerido, mínimo 1, debe existir en la BD
- **subject_id**: Requerido, mínimo 1, debe existir en la BD
- **grade**: Requerido, entre 0 y 100

---

## 🔐 Llaves Foráneas

Las relaciones están protegidas con constraints de base de datos:

```sql
grades.student_id → students.student_id (ON DELETE CASCADE)
grades.subject_id → subjects.subject_id (ON DELETE CASCADE)
```

Esto asegura la integridad referencial: al eliminar un estudiante o materia, sus calificaciones asociadas también se eliminan automáticamente.

---

## 🚦 Códigos de Estado HTTP

| Código | Significado | Uso |
|--------|------------|-----|
| 200 | OK | Consulta o actualización exitosa |
| 201 | Created | Recurso creado exitosamente |
| 400 | Bad Request | Datos inválidos o faltantes |
| 404 | Not Found | Recurso no encontrado |
| 500 | Internal Server Error | Error del servidor |

---

## 🧪 Probar la API

### Ejemplo de flujo completo

```bash
# 1. Crear un estudiante
curl -X POST http://localhost:8082/api/students \
  -H "Content-Type: application/json" \
  -d '{"name": "Ana López", "group": "4B", "email": "ana@escuela.com"}'

# 2. Crear una materia
curl -X POST http://localhost:8082/api/subjects \
  -H "Content-Type: application/json" \
  -d '{"name": "Física"}'

# 3. Registrar una calificación
curl -X POST http://localhost:8082/api/grades \
  -H "Content-Type: application/json" \
  -d '{"student_id": 1, "subject_id": 1, "grade": 92.0}'

# 4. Consultar calificaciones del estudiante
curl http://localhost:8082/api/grades/student/1
```

---

## 👤 Autor

**Estefany Montiel**
- GitHub: [@EstefanyMontiel](https://github.com/EstefanyMontiel)

---

## 📄 Licencia

Este proyecto está bajo la Licencia MIT.

---