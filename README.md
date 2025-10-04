# Task Management API - Arquitectura Hexagonal

Una API REST para la gestión de tareas implementada en Go utilizando los principios de la **Arquitectura Hexagonal** (también conocida como **Ports and Adapters**).

## 📋 Tabla de Contenidos

- [Arquitectura](#-arquitectura)
- [Tecnologías](#-tecnologías)
- [Estructura del Proyecto](#-estructura-del-proyecto)
- [Instalación y Configuración](#-instalación-y-configuración)
- [Uso de la API](#-uso-de-la-api)
- [Documentación de la API](#-documentación-de-la-api)
- [Base de Datos](#-base-de-datos)

## 🏗️ Arquitectura

Este proyecto implementa la **Arquitectura Hexagonal** propuesta por Alistair Cockburn, que permite crear aplicaciones altamente desacopladas y testeable. La arquitectura se organiza en las siguientes capas:

### Capas de la Arquitectura Hexagonal

```
┌─────────────────────────────────────────────────────────────┐
│                    Adaptadores Externos                    │
│  ┌─────────────┐  ┌──────────────┐  ┌─────────────────┐   │
│  │   HTTP API  │  │  MongoDB     │  │  Documentación  │   │
│  │ (Gin/REST)  │  │ Repository   │  │   (OpenAPI)     │   │
│  └─────────────┘  └──────────────┘  └─────────────────┘   │
└─────────────────────────────────────────────────────────────┘
           │                 │                   │
           ▼                 ▼                   ▼
┌─────────────────────────────────────────────────────────────┐
│                        Puertos                             │
│    ┌─────────────┐         ┌──────────────────┐           │
│    │   Handler   │         │  TaskRepository  │           │
│    │ Interface   │         │   Interface      │           │
│    └─────────────┘         └──────────────────┘           │
└─────────────────────────────────────────────────────────────┘
           │                           │
           ▼                           ▼
┌─────────────────────────────────────────────────────────────┐
│                    Núcleo de Dominio                       │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────┐    │
│  │   Domain    │  │  Services   │  │     Ports       │    │
│  │   (Task)    │  │ (Use Cases) │  │ (Interfaces)    │    │
│  └─────────────┘  └─────────────┘  └─────────────────┘    │
└─────────────────────────────────────────────────────────────┘
```

### Componentes Principales

#### 1. **Dominio (Core)**

- **Entidades**: `Task` - Modelo de dominio con lógica de negocio
- **Puertos**: `TaskRepository` - Interfaces que define los contratos
- **Servicios**: Casos de uso de la aplicación (Create, Get, List, Update, Delete)

#### 2. **Adaptadores (Adapters)**

- **HTTP Adapter**: Maneja las peticiones REST usando Gin
- **MongoDB Adapter**: Implementa la persistencia en MongoDB
- **Documentation Adapter**: Genera documentación OpenAPI

#### 3. **Infraestructura (Shared)**

- **Configuración**: Manejo de variables de entorno
- **Base de Datos**: Conexión y configuración de MongoDB

## 🛠️ Tecnologías

### Backend

- **[Go](https://golang.org/)** v1.25.1 - Lenguaje de programación
- **[Gin](https://gin-gonic.com/)** v1.11.0 - Framework web HTTP
- **[UUID](https://github.com/google/uuid)** v1.6.0 - Generación de identificadores únicos

### Base de Datos

- **[MongoDB](https://www.mongodb.com/)** - Base de datos NoSQL
- **[MongoDB Go Driver](https://go.mongodb.org/mongo-driver)** v2.3.0 - Driver oficial de MongoDB para Go

### Configuración y Documentación

- **[GoDotEnv](https://github.com/joho/godotenv)** v1.5.1 - Carga de variables de entorno
- **[Go Scalar API Reference](https://github.com/MarceloPetrucio/go-scalar-api-reference)** - Documentación interactiva de la API

### Herramientas de Desarrollo

- **[Go Modules](https://blog.golang.org/using-go-modules)** - Gestión de dependencias
- **Context** - Manejo de timeouts y cancelación de operaciones

## 📁 Estructura del Proyecto

```
test_hex_architecture/
├── api/
│   └── openapi/                    # Especificaciones OpenAPI
├── cmd/
│   └── api/
│       └── main.go                 # Punto de entrada de la aplicación
├── internal/
│   ├── adapter/                    # Adaptadores externos
│   │   ├── http/                   # Adaptador HTTP (REST API)
│   │   │   ├── task_handler.go     # Controladores de tareas
│   │   │   └── docs/               # Documentación de la API
│   │   │       ├── components.go
│   │   │       ├── spec.go
│   │   │       └── paths/
│   │   │           └── task.go
│   │   └── repository/             # Adaptadores de persistencia
│   │       └── mongo/
│   │           └── task_repository.go # Implementación de MongoDB
│   ├── core/                       # Núcleo de la aplicación
│   │   ├── domain/                 # Entidades de dominio
│   │   │   └── task/
│   │   │       └── task.go         # Modelo de dominio Task
│   │   ├── port/                   # Puertos (interfaces)
│   │   │   └── task_repository.go  # Interface del repositorio
│   │   └── service/                # Servicios de aplicación (Use Cases)
│   │       └── task/
│   │           ├── create.go       # Caso de uso: Crear tarea
│   │           ├── delete.go       # Caso de uso: Eliminar tarea
│   │           ├── get.go          # Caso de uso: Obtener tarea
│   │           └── update.go       # Caso de uso: Actualizar tarea
│   └── shared/                     # Infraestructura compartida
│       ├── config/                 # Configuración
│       │   ├── env.go              # Carga de variables de entorno
│       │   └── vars.go             # Definición de variables
│       └── db/                     # Configuración de base de datos
│           └── mongo.go            # Conexión a MongoDB
├── go.mod                          # Definición de módulo Go
├── go.sum                          # Checksums de dependencias
└── README.md                       # Documentación del proyecto
```

### Descripción de Directorios

- **`cmd/`**: Contiene los puntos de entrada de la aplicación
- **`internal/adapter/`**: Implementaciones concretas de los puertos (HTTP handlers, repositories)
- **`internal/core/`**: Lógica de negocio pura, independiente de la infraestructura
- **`internal/shared/`**: Código de infraestructura compartido (configuración, DB, etc.)
- **`api/`**: Especificaciones y documentación de la API

## 🚀 Instalación y Configuración

### Prerrequisitos

- **Go** 1.25.1 o superior
- **MongoDB** 4.4 o superior
- **Git**

### 1. Clonar el Repositorio

```bash
git clone <repository-url>
cd test_hex_architecture
```

### 2. Instalar Dependencias

```bash
go mod download
```

### 3. Configurar Variables de Entorno

Crear un archivo `.env` en la raíz del proyecto:

```env
# Configuración del servidor HTTP
PORT=8080

# Configuración de MongoDB
MONGO_USER=admin
MONGO_PASS=password
MONGO_HOST=localhost
MONGO_PORT=27017
MONGODB_DB=tasks_db
```

### 4. Ejecutar la Aplicación

```bash
go run cmd/api/main.go
```

La aplicación estará disponible en `http://localhost:8080`

## 🔗 Uso de la API

### Endpoints Disponibles

| Método | Endpoint      | Descripción              |
| ------ | ------------- | ------------------------ |
| POST   | `/tasks`      | Crear una nueva tarea    |
| GET    | `/tasks/{id}` | Obtener una tarea por ID |
| GET    | `/tasks`      | Listar todas las tareas  |
| PUT    | `/tasks/{id}` | Actualizar una tarea     |
| DELETE | `/tasks/{id}` | Eliminar una tarea       |

### Ejemplos de Uso

#### Crear una Tarea

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Completar documentación",
    "description": "Escribir documentación completa del proyecto"
  }'
```

#### Obtener una Tarea

```bash
curl http://localhost:8080/tasks/{task-id}
```

#### Listar Todas las Tareas

```bash
curl http://localhost:8080/tasks
```

#### Actualizar una Tarea

```bash
curl -X PUT http://localhost:8080/tasks/{task-id} \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Título actualizado",
    "description": "Descripción actualizada",
    "done": true
  }'
```

#### Eliminar una Tarea

```bash
curl -X DELETE http://localhost:8080/tasks/{task-id}
```

## 📚 Documentación de la API

La aplicación incluye documentación interactiva de la API generada automáticamente usando OpenAPI 3.0.

Accede a la documentación en: `http://localhost:8080/docs`

La documentación incluye:

- Esquemas de datos
- Ejemplos de request/response
- Códigos de estado HTTP
- Interfaz interactiva para probar los endpoints

## 🗃️ Base de Datos

### MongoDB

El proyecto utiliza **MongoDB** como base de datos principal. La elección de MongoDB se debe a:

- **Flexibilidad de esquema**: Ideal para el desarrollo ágil
- **Escalabilidad horizontal**: Soporte para grandes volúmenes de datos
- **Consultas rich**: Soporte para consultas complejas
- **Alta disponibilidad**: Replicación y sharding nativos

### Modelo de Datos

#### Colección: `tasks`

```javascript
{
  "_id": "uuid-string",           // ID único de la tarea
  "title": "string",              // Título de la tarea (obligatorio)
  "description": "string",        // Descripción de la tarea
  "done": boolean,                // Estado de completado
  "created_at": int64,            // Timestamp de creación (Unix)
  "updated_at": int64             // Timestamp de última actualización (Unix)
}
```

### Configuración de Conexión

La conexión a MongoDB se configura a través de variables de entorno:

```env
MONGO_USER=admin
MONGO_PASS=password
MONGO_HOST=localhost
MONGO_PORT=27017
MONGODB_DB=tasks_db
```

La URI de conexión se construye automáticamente: `mongodb://user:pass@host:port`

### Gestión de Conexiones

- **Timeout de conexión**: 10 segundos
- **Context timeout**: 5 segundos para operaciones
- **Graceful shutdown**: Cierre controlado de conexiones
- **Pool de conexiones**: Manejado automáticamente por el driver

## 🏗️ Principios de Arquitectura Hexagonal Aplicados

### 1. **Inversión de Dependencias**

- El core no depende de detalles de implementación
- Los adaptadores dependen del core, no al revés

### 2. **Separación de Responsabilidades**

- **Dominio**: Lógica de negocio pura
- **Puertos**: Contratos de entrada y salida
- **Adaptadores**: Implementaciones específicas de tecnología

### 3. **Testabilidad**

- Los casos de uso pueden ser testeados sin infraestructura
- Los adaptadores pueden ser mockeados fácilmente

### 4. **Flexibilidad**

- Fácil cambio de base de datos (MongoDB → PostgreSQL)
- Fácil cambio de framework web (Gin → Echo)
- Fácil adición de nuevos adaptadores

## 🚦 Estado del Proyecto

Este proyecto es una implementación de referencia que demuestra:

✅ **Implementado:**

- Arquitectura Hexagonal completa
- CRUD completo para tareas
- Persistencia en MongoDB
- API REST con Gin
- Documentación OpenAPI
- Configuración por variables de entorno
- Graceful shutdown

🔄 **En desarrollo:**

- Tests unitarios e integración
- Validación de datos más robusta
- Logging estructurado
- Métricas y monitoreo

---
