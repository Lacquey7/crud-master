# CRUD Master

CRUD Master est un projet d’exemple en architecture micro‑services écrit en Go, destiné à illustrer un CRUD de films (Inventory) et une application Billing, derrière un API Gateway. L’infrastructure de dev s’appuie sur Vagrant pour provisionner 3 VM Ubuntu, PM2 pour superviser les binaires Go, et Docker pour lancer RabbitMQ et PostgreSQL.

- API Gateway (Go) — centralise l’accès aux services Inventory et Billing
- Inventory App (Go) — CRUD de films, persistant en PostgreSQL
- Billing App (Go) — service de facturation, PostgreSQL + RabbitMQ

Le dépôt fournit des scripts de provisioning et un Vagrantfile prêts à l’emploi pour démarrer rapidement un environnement multi‑VM.

## Sommaire
- Prérequis
- Architecture & Réseau
- Démarrage rapide (Vagrant)
- Configuration (variables d’environnement)
- Build & Run (PM2)
- Endpoints
- Développement local (optionnel)
- Dépannage

## Prérequis
- Vagrant 2.3+ et un provider (VirtualBox, VMware, etc.)
- Connexion Internet (pour installer Go, Node/PM2, Docker, images)
- 4 à 6 Go de RAM disponibles pour les 3 VM (selon votre machine)

Les scripts d’install téléchargent / installent automatiquement :
- Go 1.25.1 (ARM64 dans les scripts — adapté aux box Ubuntu utilisées)
- Node.js (LTS) + PM2
- Docker (et images PostgreSQL, RabbitMQ)

## Architecture & Réseau
Trois VMs sont définies dans le Vagrantfile :
- gateway — 192.168.56.10
  - Fwd port hôte→VM: 8080→8080 (API Gateway)
- inventory — 192.168.56.11
  - Fwd port: 8083→8080 (service Inventory)
  - Docker: PostgreSQL exposé en 5432
- billing — 192.168.56.12
  - Fwd ports: 8082→8080 (service Billing), 5672→5672 (AMQP), 15672→15672 (RabbitMQ UI)
  - Docker: PostgreSQL exposé en 5431, RabbitMQ en 5672/15672

L’API Gateway appelle :
- Inventory sur 192.168.56.11:8080
- Billing sur 192.168.56.12:8080
- RabbitMQ sur 192.168.56.12:5672

## Démarrage rapide (Vagrant)
1) Depuis la racine du projet :
- vagrant up

Vagrant :
- copie le code de chaque service dans la VM correspondante
- installe Go/Node/PM2/Docker
- lance les conteneurs nécessaires (PostgreSQL, RabbitMQ)
- build chaque service Go et le démarre via PM2

2) Accès rapides
- API Gateway: http://localhost:8080
- Inventory direct: http://localhost:8083
- Billing direct: http://localhost:8082
- RabbitMQ UI (billing VM): http://localhost:15672 (guest/guest)

Conseil: la première exécution peut être longue (téléchargements d’images/packages).

## Configuration (.env)
Chaque service possède son fichier .env dans srcs/<service> :

- srcs/api-gateway/.env
  - PORT_SERVER=8080
  - ADDRESS_SERVER=0.0.0.0
  - RabbitMQ: ADDRESS_MQ=192.168.56.12, PORT_MQ=5672, USERNAME_MQ=guest, PASSWORD_MQ=guest
  - Services: BILLING_APP_SRV=192.168.56.12:8080, INVENTORY_APP_SRV=192.168.56.11:8080

- srcs/inventory-app/.env
  - PORT_SERVER=8080, ADDRESS_SERVER=0.0.0.0
  - PostgreSQL: ADDRESS_DB=localhost, PORT_DB=5432, USERNAME_DB=username, PASSWORD_DB=Azerty12@, NAME_DB=postgres

- srcs/billing-app/.env
  - PORT_SERVER=8080, ADDRESS_SERVER=0.0.0.0
  - PostgreSQL: ADDRESS_DB=localhost, PORT_DB=5431, USERNAME_DB=username, PASSWORD_DB=Azerty12@, NAME_DB=postgres
  - RabbitMQ: ADDRESS_MQ=0.0.0.0, PORT_MQ=5672, USERNAME_MQ=guest, PASSWORD_MQ=guest

Les scripts Vagrant alignent les ports/IPS ci‑dessus via Docker et la config réseau privée.

## Build & Run (PM2)
Chaque service inclut un fichier ecosystem.config.js pour PM2 et build le binaire Go dans tmp/main.

- PM2 démarre automatiquement les services lors du provisioning et les relance au reboot.
- Logs PM2 : ~/.pm2/logs/* dans chaque VM (ex. api-gateway-out.log / error.log).

Commandes utiles (dans la VM cible, utilisateur vagrant) :
- pm2 ls
- pm2 logs <nom-app>
- pm2 restart ecosystem.config.js --update-env
- pm2 stop <nom-app> && pm2 delete <nom-app>

Noms d’apps PM2 :
- api-gateway (VM gateway)
- inventory-app (VM inventory)
- billing-app (VM billing)

## Endpoints
Voici un aperçu minimal des routes exposées par les services :

- API Gateway (http://localhost:8080)
  - GET /billing — proxy vers Billing
  - GET /inventory — proxy vers Inventory

- Inventory (http://localhost:8083 ou 192.168.56.11:8080)
  - Base: /api/movies
  - GET    /api/movies             — liste tous les films (filtre possible par nom dans le titre)
  - POST   /api/movies             — crée un film
  - DELETE /api/movies             — supprime tous les films
  - GET    /api/movies/{id}        — récupère un film par id
  - PUT    /api/movies/{id}        — met à jour un film par id
  - DELETE /api/movies/{id}        — supprime un film par id

- Billing (http://localhost:8082 ou 192.168.56.12:8080)
  - Base: /api/billing
  - GET /api/billing/{id}          — récupère une facture par id

Selon l’implémentation, certaines routes de Gateway ajoutent de la logique (ex: RabbitMQ) avant d’appeler les services backend.

## Développement local (optionnel)
Vous pouvez exécuter un service en dehors de Vagrant, mais veillez à aligner les variables d’environnement (.env) et les dépendances (PostgreSQL, RabbitMQ). Exemple pour Inventory :

- Installer Go 1.21+ sur votre machine
- Copier/adapter srcs/inventory-app/.env
- cd srcs/inventory-app && go build -o tmp/main . && ./tmp/main

Pour un run supervisé, installez Node.js + PM2 et utilisez le ecosystem.config.js du service.

## Dépannage
- vagrant status — vérifier l’état des VM
- vagrant reload <vm> — redémarrer une VM
- vagrant destroy <vm> -f && vagrant up <vm> — recréer une VM propre
- Sur une VM, vérifier Docker : docker ps, docker logs <container>
- Sur une VM, vérifier PM2 : pm2 ls, pm2 logs
- Connexions DB :
  - Inventory : PostgreSQL écoute sur 5432 (VM inventory)
  - Billing : PostgreSQL écoute sur 5431 (VM billing)
- RabbitMQ UI : http://localhost:15672 (guest/guest)

Si un build Go échoue, consultez les fichiers tmp/build-errors.log (si présents) ou les logs PM2.