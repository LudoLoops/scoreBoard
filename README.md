# Tableau de Bord & Gestion de Projet 🚀

## Description

Application web conçue pour la gestion de projets et le suivi des performances de l'équipe. Elle présente un tableau de bord affichant les noms de projets et les scores des membres de l'équipe, ainsi que la possibilité d'ajouter de nouveaux utilisateurs et projets via API. Les données sont chargées à partir de `equipage.json` situé dans le répertoire `src/content` et affichées sur la page d'accueil.

## Technologies

- **Svelte:** Framework frontend pour la construction d'interfaces utilisateur réactives. ⚡
- **SvelteKit:** Framework Svelte pour la construction d'applications full-stack. 🏗️
- **Données JSON:** Utilise `equipage.json` pour les données de projet et d'utilisateur. 💾
- **Routes API:** Implémentées à l'aide des routes API de SvelteKit pour l'ingestion et la récupération des données. 🔗

## Fonctionnalités

- **Tableau de Bord de Projets:** Affiche une liste de projets avec leurs scores. 📊
- **Affichage des Données Utilisateur:** Cliquer sur un projet révèle un tableau des utilisateurs associés à ce projet, y compris leurs scores. Le tableau est triable par score. 🔍
- **Points d'Extrémité API :**
  - `GET /api/ingest` : Récupère toutes les données de projet et d'utilisateur. 🔄
  - `POST /api/ingest` : Ajoute un nouvel utilisateur à un projet. Nécessite un payload JSON au format spécifié (voir ci-dessous). ➕
  - `GET /api/projet` : Récupère une liste de tous les noms de projets. 📝

## Source de Données

L'application utilise `equipage.json` situé dans le répertoire `src/content`. Ce fichier contient les données de projet et d'utilisateur de base. 📁

## Détails des Points d'Extrémité API

- **`GET /api/ingest`**:
  - **Objectif :** Récupère toutes les données de projet et d'utilisateur.
  - **Réponse :** Retourne un tableau JSON contenant toutes les données de projet et d'utilisateur.
  - **Exemple de Réponse :** (Illustratif - la structure réelle dépendra de `equipage.json`)
    ```json
    [
    	{
    		"name": "Versatile Fast-Charging Toy",
    		"slug": "versatile-fast-charging-toy",
    		"users": [
    			{
    				"userId": "17",
    				"nom": "ToyHuger",
    				"note": 39,
    				"bonus": 12,
    				"xp": 51,
    				"dateSoumission": "2025-07-27",
    				"tags": ["innovation", "proactivité"]
    			}
    		]
    	}
    	// ... autres projets et utilisateurs
    ]
    ```

- **`POST /api/ingest`**:
  - **Objectif :** Ajoute un nouvel utilisateur à un projet.
  - **Corps de la Requête (JSON) :**
    ```json
    {
    	"name": "Nouveau Nom de Projet",
    	"user": {
    		"userId": "18",
    		"nom": "NouveauUtilisateur",
    		"note": 10,
    		"bonus": 5,
    		"xp": 20,
    		"dateSoumission": "2025-07-28",
    		"tags": ["test", "développement"]
    	}
    }
    ```
  - **Réponse :** Retourne les données mises à jour (potentiellement un code de statut 201 Créé avec le JSON mis à jour).

- **`GET /api/projet`**:
  - **Objectif :** Récupère une liste de tous les noms de projets.
  - **Réponse :** Retourne un tableau JSON contenant tous les noms de projets.
  - **Exemple de Réponse :**
    ```json
    ["Versatile Fast-Charging Toy", "Autre Projet", "Troisième Projet"]
    ```

## Structure du Projet

- `src/content/equipage.json`: Contient les données de projet et d'utilisateur de base.
- `src/routes/+page.svelte`: Page d'accueil affichant le tableau de bord.
- `src/routes/api/+server.js`: Gère les routes API (GET /api/ingest, GET /api/projet).

## Installation et Configuration du Développement

pour installer :
`npm install`

puis pour le developpement:
`npm run dev`

## Programme Golang pour la génération de la liste utilisateur

J'ai également créé un programme en Go dans le dossier `goGenerator`. Ce programme permet de générer la liste des utilisateurs.

Vous pouvez utiliser les flags suivants :

- `-u <nombre_d'utilisateurs>`: Définit le nombre d'utilisateurs à générer.
- `-p <nombre_de_projets>`: Définit le nombre de projets à générer.

Exemple :

`go run goGenerator/generate_users.go -u 10 -p 5`

Ce command générera 10 utilisateurs et 5 projets.
