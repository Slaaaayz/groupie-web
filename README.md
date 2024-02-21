# <div align="center">GROUPIE-TRACKER (Symphony)

## SOMMAIRE

- [I. Comment installer le Groupie-Tracker](#i-comment-installer-le-groupie-tracker)
- [II. Hebergement de Symphony](#ii-hebergement-de-symphony)
- [III. Fonctionnement de Symphony](#iii-fonctionnement-de-symphony)
- [IV. Page Jeu](#iv-page-jeu)
- [V. Si le Groupie-Tracker ne fonctionne pas](#v-si-le-groupie-tracker-ne-fonctionne-pas)

## I. Comment installer le Groupie-Tracker

Pour installer le groupie-tracker, commencez par cloner le repository sur votre ordinateur en utilisant le terminal et la commande suivante :

```bash
git clone https://ytrack.learn.ynov.com/git/cmaxime/groupie-tracker.git
```
Ensuite, ouvrez le dossier groupie-tracker dans VsCode, puis dans un terminal, lancez le site via la commande suivante:

```go
go run ./serveur.go
```

Pour accéder au site, ouvrez un navigateur et entrez l'adresse suivante : http://localhost:8080/

## II. Hebergement de Symphony

Nous avons choisi d'héberger Symphony. Pour accéder à Symphony, veuillez vous connectez à l'adresse suivante : https://diane.etudiants.ynov-bordeaux.com/

## III. Fonctionnement de Symphony

Le Groupie Tracker (Symphony) est un site web conçu pour permettre aux utilisateurs de rechercher des informations sur l'artiste de son choix parmi ceux présent sur notre plateforme. Que ce soit la liste des membres, la date de création du groupe, les lieux de concerts ainsi que leurs dates, l'utilisateur a à sa disposition de nombreuses informations. Il peut également écouter la musique des artistes directement sur le site, ou encore jouer à un QUI EST-CE musical !

## IV. Page Jeu

La page jeu est une page qui permet à l'utilisateur de jouer à un jeu de devinette. Le but du jeu est de deviner le nom de l'artiste à l'aide des indices suivants : la date de sortie de son premier album, puis l'année de son lancement, et enfin une image de l'artiste. Si l'utilisateur a deviné le nom de l'artiste, un message de félicitations s'affiche. Sinon, un message de défaite s'affiche.

## V. Si le Groupie-Tracker ne fonctionne pas

Si le Groupie-Tracker ne fonctionne pas en local, suivez ces étapes :

1. Vérifiez l'installation de Go sur votre ordinateur.

2. Assurez-vous que le repository du Groupie-Tracker est correctement installé.

3. Ouvrez le dossier Groupie-Tracker sur VsCode.

4. Si le problème persiste, supprimez le dossier Groupie-Tracker et recommencez l'installation.

5. Si aucune de ces solutions ne fonctionne, contactez les auteurs du Groupie-Tracker.

## <div align="right">Les auteurs du Groupie-Tracker

<div align="right">SAUTEREAU DU PART Diane  
<div align="right">CHAMOULEAU Julien
<div align="right">CHORT Maxime
