# TP Golang M1

## Description
Trois travaux pratiques (TP) distincts en Go, indépendants les uns des autres :
- `Initiates` — premiers exercices et exemples HTTP/JSON.
- `TP-HTTP-API` — mini API HTTPS exposant un endpoint de health.
- `TP-FINAL` — projet final : agent, serveur et client pour monitoring/métriques.

## Arborescence (racine)
```
__________________________________

-Initiates/ | Dossier des premiers pas en Go
├── cust/
│   └── cust.go
├── json/
│   └── names.json
├── srv/
│   └── http.go
└── main.go
__________________________________

-TP-HTTP-API/ | TP Cours sur une mini API HTTP/HTTPS
├── go.mod/
│   └── go.sum
└── main.go
__________________________________
-TP-FINAL/  | TP Final Noté du module
├── agent/
│   └── main.go
├── server/
│   └── main.go
├── client/
│   └── main.go
└── README.md
__________________________________
```

## Prérequis
- Go 1.23 ou supérieur
- Pour métriques CPU sur Windows : WSL recommandé (les outils lisent `/proc/stat`)
- Certificats SSL pour les TPs HTTPS

## Installation et dépendances
1. Cloner le dépôt :
```bash
git clone https://github.com/ChifuyuSenpai/Golang-M1.git
```
2. Depuis la racine, récupérer les dépendances si nécessaire (ex : `TP-FINAL/agent` ou autre) :
```bash
go mod tidy
```
3. Dépendance utile pour métriques Linux :
```bash
go get github.com/c9s/goprocinfo/linux
```

## Génération des certificats (localhost)
Utiliser `mkcert` ou votre script préféré :
```bash
mkcert -install
mkcert localhost
```
Placer `localhost.pem` et `localhost-key.pem` dans le dossier qui exécute le serveur HTTPS (`TP-HTTP-API/` ou `TP-FINAL/server/`).

## Exécution par TP

- `Initiates` :
```bash
cd TP-FINAL/.. # ou depuis la racine
go run Initiates/main.go
```

- `TP-HTTP-API` (HTTPS, port par défaut 8443) :
```bash
cd TP-HTTP-API
go run main.go
# tester :
curl -k https://localhost:8443/health
```

- `TP-FINAL` (agent / server / client) :
1. Lancer le serveur :
```bash
go run TP-FINAL/server
```
2. Lancer l'agent dans un autre terminal :
```bash
go run TP-FINAL/agent
```
3. Utiliser le client pour lister les agents :
```bash
go run TP-FINAL/client list
```
Endpoints exposés par le serveur :
- `POST /metrics` — recevoir métriques agents (JSON)
- `GET /agents` — lister agents (JSON)
- `GET /health` — health check simple

Sur Windows vous pouvez aussi builder :
```bash
go build -o server.exe ./TP-FINAL/server
.\server.exe
```

## Exemple de requêtes
Envoyer des métriques (exemple) :
```bash
curl -X POST -H "Content-Type: application/json" -d @metrics.json http://localhost:8080/metrics
```
Lister les agents :
```bash
curl http://localhost:8080/agents
```

## Bonnes pratiques et hardening
- Fermer les corps de requête (`defer r.Body.Close()`) et limiter la taille (`io.LimitReader`).
- Valider `AgentID` et les payloads JSON.
- Utiliser `http.Server` et `Shutdown` pour arrêt propre.
- Éviter de tenir le mutex pendant l'encodage JSON : copier les données sous verrou puis encoder hors verrou.

## Ressources utiles et utilisées
- `net/http` : https://pkg.go.dev/net/http
- `encoding/json` : https://pkg.go.dev/encoding/json
- `sync` : https://pkg.go.dev/sync
- `goprocinfo` : https://github.com/c9s/goprocinfo
- https://gobyexample.com/http-servers
- https://pkg.go.dev/net/http#pkg-constants
- https://golang.org/doc/effective_go

## Auteur
Kyllian R — GitHub: `@ChifuyuSenpai`

## Licence
Usage éducatif — voir le dépôt pour détails.
```
