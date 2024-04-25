# Projet en Go : Service de réservation en ligne

## Objectif

Développer un système de réservation de salles en ligne de commande pour gérer les réservations,les salles et les utilisateurs.

## Commandes pour lancer le Projet

docker compose up
go run ./main.go

## Description détaillée

#### Interaction avec l'utilisateur

Implémente la gestion des entrées utilisateur pour permettre la navigation et l'exécution des actions dans le menu.

#### Navigation

Permet à l'utilisateur de sélectionner les différentes fonctionnalités.

### Fonctionnalités du menu

#### Lister les salles disponibles

Fonction qui affiche les salles qui ne sont pas réservées pour une plage horaire spécifique.

#### Créer une réservation

Demande à l'utilisateur de fournir les détails nécessaires pour une réservation (taille du groupe, nom de la salle, date, heure) et vérifie si la salle est disponible avant de créer la réservation.

#### Annuler une réservation

Demande à l'utilisateur de fournir les détails nécessaires pour annuler une réservation (numéro de réservation) et vérifie si la réservation existe avant de l'annuler.

#### Visualiser les réservations

Affiche les réservations existantes pour une salle spécifique en permettant de filtrer pour une date spécifique.

### Base de Données

Utilisation d'une base de données dockerisée pour stocker les salles et les réservations. Implémentez des opérations pour ajouter, supprimer et récupérer des données.

### Export des réservations

Fonctionnalité pour exporter les réservations en JSON.
