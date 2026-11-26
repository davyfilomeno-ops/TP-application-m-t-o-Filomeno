# Station météo

Cette application Go affiche les prévisions météo d'une commune française.
## Lancer le projet

terminal de commande :
go run .
```

Ouvrez ensuite `http://localhost:8000/communes` dans le navigateur.

## Fonctionnement

- `routes/` associe chaque adresse du site à un contrôleur.
- `controllers/` reçoit la recherche et prépare la réponse HTTP.
- `services/` récupère la commune, sa météo et l'adresse de la carte.
- `api/` appelle l'API des communes et l'API Open-Meteo.
- `models/` décrit les données utilisées dans l'application.
- `helpers/` traduit les codes WMO en icônes et libellés météo.
- `templates/` génère les pages HTML côté serveur, sans JavaScript.
