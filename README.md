# FTDS3status 🗃️📈

**FTDS3status** es una Lambda en Go pensada para registrar el estado de procesos distribuidos (batch y task) en AWS S3.  
Ideal para flujos asincrónicos donde necesitas trazabilidad sin usar bases de datos.

---

## ✨ ¿Qué hace?

- Guarda en S3 el estado de procesamiento de **batches** y **tasks** (pending, running, failed, success, completed).
- El formato de almacenamiento permite fácilmente:
  - Consultar el estado actual.
  - Tener un historial completo.
  - Usar AWS Athena para análisis y auditoría.

---

## 📁 Estructura de S3 esperada

```
FTD_CONTRACT_LOADER/
├── BATCHS/
│   └── {batch_id}/
│       └── status/
│           ├── pending/
│           ├── running/
│           ├── failed/
│           ├── success/
│           └── completed/
│
│       └── tasks/
│           └── {task_id}/
│               ├── pending/
│               ├── running/
│               ├── failed/
│               ├── success/
│               └── completed/
```

---

## 🚀 Cómo usar

### Build con Docker para AWS Lambda

```bash
docker build -t ftds3status .
```

### Probar localmente

```bash
docker run -p 9000:8080 ftds3status
```

Y luego:

```bash
curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" \
     -d @examples/add_batch_completed.json
```

### Despliegue

```bash
./deploy.sh  # sube imagen a ECR y actualiza la función Lambda
```

---

## 🧪 Testing y Mocks

- En `cmd/lambda/main.go` encontrarás **mocks** de los métodos que escriben en S3, ideales para testing local o unit tests.
- También hay funciones `mockAdd<Status>` para simular requests de prueba.

---

## 📦 Estructura del Proyecto

```
FTDS3status/
├── cmd/lambda/         # Entrypoint de la Lambda
├── src/
│   ├── handlers/       # Dispatcher de comandos
│   ├── models/         # Tipos y estructuras comunes
│   ├── providers/      # Lógica de escritura en S3
│   └── settings/       # Constantes globales
├── examples/           # JSON de prueba
├── Dockerfile
├── deploy.sh
└── go.mod / go.sum
```

---

## 📥 Ejemplo de Request JSON

```json
{
  "command": "add-task-status",
  "account": "acme_inc",
  "batch_id": "abc123",
  "task_id": "15123456-9",
  "status": "completed",
  "payload": {
    "rut": "15123456-9",
    "resultado": "todo bien"
  }
}
```

---

## ⚙️ Requisitos

- Go 1.22
- Docker
- AWS CLI configurado para ECR y Lambda

---

## 🧠 Futuro

- Añadir soporte para `get-status` por batch o task
- Consolidar estados a `summary.json` para cada lote
- Integración con SQS / Step Functions para automatizar traza

---

## 📄 Licencia

MIT © 2025 [@patriciojlg](https://github.com/patriciojlg)