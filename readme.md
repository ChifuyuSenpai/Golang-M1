
# TP-HTTP-API - Service de Health Check

API REST en Go fournissant des informations de santé système (health check) avec métriques CPU, mémoire et uptime.

## 📋 Prérequis

- Go 1.23 ou supérieur
- Certificats SSL (pour HTTPS)
- Système Linux (pour les métriques CPU via `/proc/stat`)

## 🚀 Installation

1. **Cloner le repository**
```bash
git clone https://github.com/ChifuyuSenpai/Golang-M1.git
cd TP-HTTP-API
```

2. **Initialiser le module Go**
```bash
go mod init tp-http-api
```

3. **Installer les dépendances**
```bash
go get github.com/c9s/goprocinfo/linux
go mod tidy
```

## 🔐 Configuration SSL

Générez vos certificats SSL pour localhost :

```bash
mksert localhost
```

Les fichiers requis :
- `localhost.pem` : Certificat SSL
- `localhost-key.pem` : Clé privée

## ▶️ Utilisation

**Démarrer le serveur :**
```bash
go run main.go
```

Le service démarrera sur le port **8443** en HTTPS.

**Accéder au health check :**
```bash
curl -k https://localhost:8443/health
```

## 📊 Endpoint API

### GET /health

Retourne les informations de santé du système au format JSON.

**Réponse exemple :**
```json
{
  "time": "2025-01-15 14:30:45",
  "hostname": "LINUX-01",
  "pid": 12345,
  "status": "OK",
  "go_version": "go1.23",
  "uptime": "2h15m30s",
  "memory_usage_mb": 5,
  "memory_alloc_mb": 10,
  "memory_total_mb": 256,
  "cpu_usage_percent": 15.5,
  "cpu_cores": 8,
  "cpu_user": 1234567,
  "cpu_system": 234567,
  "cpu_idle": 8901234
}
```

**Champs de la réponse :**

| Champ | Type | Description |
|-------|------|-------------|
| `time` | string | Horodatage de la requête |
| `hostname` | string | Nom de la machine |
| `pid` | int | Process ID de l'application |
| `status` | string | État du service (`OK`) |
| `go_version` | string | Version de Go utilisée |
| `uptime` | string | Temps depuis le démarrage du service |
| `memory_usage_mb` | uint64 | Mémoire actuellement allouée (MB) |
| `memory_alloc_mb` | uint64 | Total cumulé de mémoire allouée (MB) |
| `memory_total_mb` | uint64 | Mémoire système totale (MB) |
| `cpu_usage_percent` | float64 | Pourcentage d'utilisation CPU |
| `cpu_cores` | int | Nombre de cœurs CPU |
| `cpu_user` | uint64 | Temps CPU en mode utilisateur |
| `cpu_system` | uint64 | Temps CPU en mode système |
| `cpu_idle` | uint64 | Temps CPU inactif |

## 🛠️ Technologies

- **Go 1.23** - Langage de programmation
- **net/http** - Serveur HTTP/HTTPS natif
- **runtime** - Métriques mémoire et informations système Go
- **goprocinfo** - Lecture des statistiques CPU Linux (`/proc/stat`)

## 📦 Dépendances

```go
github.com/c9s/goprocinfo v0.0.0-20210130143923-c95fcf8c64a8
```

## 🔧 Build

**Compiler l'application :**
```bash
go build -o health-api main.go
```

**Exécuter le binaire :**
```bash
./health-api
```

## 📝 Logs

Au démarrage, le service affiche :
```
Service UP » Listening on port 8443 !
```

Les erreurs de lecture CPU sont loguées sans bloquer le service.

## ⚠️ Notes importantes

- **Linux uniquement** : Les métriques CPU utilisent `/proc/stat` (non disponible sur Windows/macOS ou alors via WSL)
- **HTTPS obligatoire** : Le service nécessite des certificats SSL valides
- **Certificats auto-signés** : Les navigateurs afficheront un avertissement de sécurité

## 👤 Auteur

**Kyllian R**
- GitHub: [@ChifuyuSenpai](https://github.com/ChifuyuSenpai)

## 📄 Licence

Ce projet est à usage éducatif dans le cadre du cours de Golang M1.

## 🌿 Branches

- `main` - Version stable
- `dev` - Développement en cours
```
