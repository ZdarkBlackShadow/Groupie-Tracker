# Groupie tracker

## Description

Ce projet est un site web dédié à Genshin Impact, offrant une base de données complète et interactive sur le jeu. Il permet aux utilisateurs de consulter et de gérer des collections d'informations sur les personnages, les armes, les artefacts et plus encore.

## Fonctionnalités

### 🎮 Système d'Authentification

- Inscription et connexion des utilisateurs
- Gestion de profil personnalisé
- Système de récupération de mot de passe

### 📚 Base de Données Complète

- Personnages
- Armes
- Artefacts
- Éléments
- Boss
- Nourriture
- Potions
- Domaines
- Ennemis

### 🔍 Fonctionnalités de Recherche et Filtrage

- Recherche globale
- Filtres avancés pour chaque catégorie
- Tri personnalisable
- Pagination des résultats

### 📑 Collections Personnelles

- Ajout d'éléments aux collections
- Gestion des collections personnelles
- Suivi des objets collectés

### 💫 Interface Utilisateur

- Animations fluides
- Navigation intuitive
- Thème inspiré de Genshin Impact

## Technologies Utilisées

### Frontend

- HTML5 🧱
- CSS3 🎨
- JavaScript ⚙️

### Backend

- Go 📘

### Database

- Json 🟨

## Installation

### ⚙️ Prérequis

- Avoir Git ([Télécharger Git](https://git-scm.com/downloads))
- Avoir goland (1.23.3) ([Télécharger Goland](https://go.dev/dl))

### 🛠️ Download the project

1. Cloner le repository

```bash
git clone https://github.com/ZdarkBlackShadow/Groupie-Tracker.git
```

### 🚀 Lancer le projet

```bash
go run main.go
```

```bash
http://localhost:8080/home
```

## 🌐 Routes de l'application

### Routes principales :

| Méthode |             Route             |       Description       |
| :-----: | :---------------------------: | :---------------------: |
|   GET   |            `/home`            |      The home page      |
|   GET   |           `/login`            |       Login page        |
|   GET   |   `/login/password-forgot`    |     Forgot password     |
|   GET   |        `/collections`         |       Collections       |
|   GET   |         `/artifacts`          |     Artifacts page      |
|   GET   |     `/artifacts/details`      |    Artifacts details    |
|   GET   |            `/boss`            |        Boss page        |
|   GET   |        `/boss/details`        |      Boss details       |
|   GET   |         `/characters`         |     Characters page     |
|   GET   |     `/characters/details`     |    Character details    |
|   GET   |          `/domains`           |      Domains page       |
|   GET   |      `/domains/details`       |     Domain details      |
|   GET   |          `/elements`          |      Elements page      |
|   GET   |      `/elements/details`      |     Element details     |
|   GET   |          `/enemies`           |      Enemies page       |
|   GET   |      `/enemies/details`       |      Enemy details      |
|   GET   |          `/weapons`           |      Weapons page       |
|   GET   |      `/weapons/details`       |     Weapon details      |
|   GET   |           `/foods`            |        Food page        |
|   GET   |       `/foods/details`        |      Food details       |
|   GET   |          `/potions`           |      Potions page       |
|   GET   |      `/potions/details`       |     Potion details      |
|   GET   |           `/search`           |       Search page       |
|   GET   |          `/loading`           |      Loading page       |
|   GET   |           `/profil`           |       Profil page       |
|   GET   |           `/about`            |       About page        |
|   GET   |           `/logout`           |         Logout          |
|   GET   |     `/add-to-collection`      |   Add to collections    |
|   GET   | `/remove-from-the-collection` | Remove from collections |
|   GET   |          `/loading`           |      Loading page       |
|   GET   |          `/progress`          |      Progress page      |

### Routes pour actions liées aux données (`POST`) :

| Méthode |             Route             |   Description   |
| :-----: | :---------------------------: | :-------------: |
|  POST   |     `/login/newregister`      |    New user     |
|  POST   |       `/login/register`       | User connexion  |
|  POST   | `/login/password-forgot/form` | Forgot password |

## Endpoint utilisé de l'API

### Lien de l'API

- Link of the API hosted on https://genshin.jmp.blue
- Github : https://github.com/genshindev/api

### Endpoint

> **:information_source: Notice:** Please replace `<baseUrl>` with the endpoint you are trying to access.

| Endpoint                | Description                                                   | Example Usage (bash)                       |
| ----------------------- | ------------------------------------------------------------- | ------------------------------------------ |
| `/`                     | Returns a list of available entity types.                     | `curl <baseUrl>/`                          |
| `/:type`                | Returns a list of available entities for a specific type.     | `curl <baseUrl>/characters`                |
| `/:type/all`            | Returns detailed information about all entities of a type.    | `curl <baseUrl>/characters/all?lang=en`    |
| `/:type/:id`            | Returns detailed information about a single entity.           | `curl <baseUrl>/characters/albedo?lang=en` |
| `/:type/:id/list`       | Returns a list of available images for a specific entity.     | `curl <baseUrl>/characters/albedo/list`    |
| `/:type/:id/:imageType` | Returns the image of a specific type for a particular entity. | `curl <baseUrl>/characters/albedo/card`    |

## 👥 Auteur

- Adrien Lecomte : [Github](https://github.com/ZdarkBlackShadow)
