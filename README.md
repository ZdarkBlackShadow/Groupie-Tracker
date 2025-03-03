# Groupie Tracker

## Description

This project is a website dedicated to Genshin Impact, offering a comprehensive and interactive database about the game. It allows users to browse and manage collections of information about characters, weapons, artifacts, and much more.

## Features

### 🎮 Authentication System

- User registration and login
- Personalized profile management
- Password recovery system

### 📚 API Interaction

- Characters
- Weapons
- Artifacts
- Elements
- Bosses
- Food
- Potions
- Domains
- Enemies

### 🔍 Search and Filtering Features

- Global search
- Advanced filters for each category
- Customizable sorting
- Paginated results

### 📑 Personal Collections

- Add items to collections
- Manage personal collections
- Track collected items

### 💫 User Interface

- Smooth animations
- Intuitive navigation
- Genshin Impact-inspired theme

## Technologies Used

### Frontend

- HTML5 🧱
- CSS3 🎨
- JavaScript ⚙️

### Backend

- Go 📘

### Database

- JSON 🟨

## Installation

### ⚙️ Prerequisites

- Git ([Download Git](https://git-scm.com/downloads))
- Goland (1.23.3) ([Download Goland](https://go.dev/dl))

### 🛠️ Download the project

1. Clone the repository

```bash
git clone https://github.com/ZdarkBlackShadow/Groupie-Tracker.git
```

### 🚀 Run the project

```bash
go run main.go
```

```bash
http://localhost:8080/home
```

## 💾 Details

###  🗂️ Struture of the project

```
└── groupie-tracker/
    ├── README.md
    ├── go.mod
    ├── main.go
    ├── assets/
    │   ├── css/
    │   │   ├── about.css
    │   │   ├── artifactDetail.css
    │   │   ├── artifacts.css
    │   │   ├── boss.css
    │   │   ├── bossDetails.css
    │   │   ├── character.css
    │   │   ├── characterDetails.css
    │   │   ├── collections.css
    │   │   ├── domainDetails.css
    │   │   ├── domains.css
    │   │   ├── elementDetails.css
    │   │   ├── elements.css
    │   │   ├── enemies.css
    │   │   ├── enemiesDetails.css
    │   │   ├── errorCode.css
    │   │   ├── food.css
    │   │   ├── foodDetails.css
    │   │   ├── forgotPassword.css
    │   │   ├── home.css
    │   │   ├── loading.css
    │   │   ├── login.css
    │   │   ├── potionDetails.css
    │   │   ├── potions.css
    │   │   ├── profil.css
    │   │   ├── searchResult.css
    │   │   ├── style.css
    │   │   ├── weaponDetails.css
    │   │   └── weapons.css
    │   ├── image/
    │   │   ├── Domain_Card.webp
    │   │   ├── GenshinImpactPotions.avif
    │   │   ├── NoImageAvaliable.webp
    │   │   ├── artifact.avif
    │   │   ├── background.webp
    │   │   ├── backgroundWeapon.webp
    │   │   ├── domain.webp
    │   │   ├── element.webp
    │   │   ├── elements.avif
    │   │   ├── food.avif
    │   │   ├── logo.webp
    │   │   ├── weapon.avif
    │   │   └── Exception/
    │   │       └── Glacier and Snowfield/
    │   │           └── flower.webp
    │   ├── js/
    │   │   ├── artifacts.js
    │   │   ├── characterDetails.js
    │   │   ├── characters.js
    │   │   ├── elements.js
    │   │   ├── food.js
    │   │   ├── loading.js
    │   │   ├── login.js
    │   │   ├── searchResult.js
    │   │   └── weapons.js
    │   └── ttf/
    │       └── zh-cn.ttf
    ├── controllers/
    │   ├── about.go
    │   ├── artifacts.go
    │   ├── boss.go
    │   ├── character.go
    │   ├── collection.go
    │   ├── domains.go
    │   ├── elements.go
    │   ├── enemies.go
    │   ├── error.go
    │   ├── food.go
    │   ├── home.go
    │   ├── init.go
    │   ├── loadData.go
    │   ├── login.go
    │   ├── potions.go
    │   ├── profil.go
    │   ├── search.go
    │   └── weapons.go
    ├── data/
    │   └── data.json
    ├── server/
    │   └── server.go
    ├── service/
    │   ├── artifacts.go
    │   ├── boss.go
    │   ├── characters.go
    │   ├── collections.go
    │   ├── domains.go
    │   ├── elements.go
    │   ├── enemies.go
    │   ├── food.go
    │   ├── potions.go
    │   ├── service.go
    │   ├── structure.go
    │   └── weapon.go
    ├── templates/
    │   ├── about.html
    │   ├── artifactDetail.html
    │   ├── artifacts.html
    │   ├── boss.html
    │   ├── bossDetails.html
    │   ├── characterDetails.html
    │   ├── characters.html
    │   ├── collections.html
    │   ├── domainDetails.html
    │   ├── domains.html
    │   ├── elementDetails.html
    │   ├── elements.html
    │   ├── enemies.html
    │   ├── enemiesDetails.html
    │   ├── errorCode.html
    │   ├── food.html
    │   ├── foodDetails.html
    │   ├── home.html
    │   ├── loadData.html
    │   ├── login.html
    │   ├── passwordforgot.html
    │   ├── potiondetail.html
    │   ├── potions.html
    │   ├── profil.html
    │   ├── searchResult.html
    │   ├── templates.html
    │   ├── weapon.html
    │   └── weaponsDetails.html
    └── utils/
        ├── artifacts.go
        ├── boss.go
        ├── character.go
        ├── collection.go
        ├── domains.go
        ├── error.go
        ├── food.go
        ├── loadData.go
        ├── login.go
        ├── potions.go
        ├── search.go
        └── weapons.go

```

## 🌐 Application Routes

### Main Routes :

| Method |             Route             |                                        Description                                        |
| :-----: | :---------------------------: | :---------------------------------------------------------------------------------------: |
|   GET   |            `/home`            |              The home page where you can choose which entity you want to see              |
|   GET   |           `/login`            |               Login page where you can chosse betwen login or new register                |
|   GET   |   `/login/password-forgot`    |            Forgot password where you can change your passwor if you forget it             |
|   GET   |        `/collections`         |            Collections page, if you're connected, you can see your collections            |
|   GET   |         `/artifacts`          |      Page displaying all Genshin artifacts. Supports `filter` and `page` parameters.      |
|   GET   |     `/artifacts/details`      |   Artifact details page. Requires an id parameter to display specific artifact details.   |
|   GET   |            `/boss`            |             Page listing all bosses. Supports `filter` and `page` parameters.             |
|   GET   |        `/boss/details`        |      Boss details page. Requires an `id` parameter to display specific boss details.      |
|   GET   |         `/characters`         |         Page displaying all characters. Supports `filter` and `page` parameters.          |
|   GET   |     `/characters/details`     | Character details page. Requires an `id` parameter to display specific character details. |
|   GET   |          `/domains`           |            Page listing all domains. Supports `filter` and `page` parameters.             |
|   GET   |      `/domains/details`       |    Domain details page. Requires an `id` parameter to display specific domain details.    |
|   GET   |          `/elements`          |          Page displaying all elements. Supports `filter` and `page` parameters.           |
|   GET   |      `/elements/details`      |   Element details page. Requires an `id` parameter to display specific element details.   |
|   GET   |          `/enemies`           |            Page listing all enemies. Supports `filter` and `page` parameters.             |
|   GET   |      `/enemies/details`       |     Enemy details page. Requires an `id` parameter to display specific enemy details.     |
|   GET   |          `/weapons`           |           Page displaying all weapons. Supports `filter` and `page` parameters.           |
|   GET   |      `/weapons/details`       |    Weapon details page. Requires an `id` parameter to display specific weapon details.    |
|   GET   |           `/foods`            |           Page listing all food items. Supports `filter` and `page` parameters.           |
|   GET   |       `/foods/details`        |      Food details page. Requires an `id` parameter to display specific food details.      |
|   GET   |          `/potions`           |           Page displaying all potions. Supports `filter` and `page` parameters.           |
|   GET   |      `/potions/details`       |    Potion details page. Requires an `id` parameter to display specific potion details.    |
|   GET   |           `/search`           |             Search page who can take `filter` parameter and `page` parameter              |
|   GET   |           `/profil`           |                                        Profil page                                        |
|   GET   |           `/about`            |                            About page who show the about page                             |
|   GET   |           `/logout`           |                             Logout who logout the actual user                             |
|   GET   |     `/add-to-collection`      |        Add to collections who take in parameter `type` and the `id` of the entiti         |
|   GET   | `/remove-from-the-collection` |      Remove from collections who take in parameter `type` and the `id` of the entiti      |
|   GET   |          `/loading`           |       Loading page who show a progress bar who represent the chargement of all data       |
|   GET   |          `/progress`          |                 Progress page who send the number for actual progression                  |

### Routes for data-related actions (`POST`):

| Méthode |             Route             |                  Description                  |
| :-----: | :---------------------------: | :-------------------------------------------: |
|  POST   |     `/login/newregister`      |             Registers a new user.             |
|  POST   |       `/login/register`       |       Authenticates and logs in a user.       |
|  POST   | `/login/password-forgot/form` | Submits a form to reset a forgotten password. |

## API Endpoint Used

### API link

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
