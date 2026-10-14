

# GO-RELOADED

## Présentation du projet

Il s'agit d'un projet de correction de texte dont l'objectif est de découvrir la manipulation des chaînes de caractères et des nombres.

## Fonctionnalités
-Lecture de fichiers : récupère et lit un fichier fourni en entrée.
-Analyse du contenu : parcourt le texte pour identifier les éléments à traiter.
-Correction ciblée : corrige ou modifie certains éléments en fonction des indications présentes dans le texte.
-Génération du résultat : produit un fichier de sortie contenant le texte modifié.
-Automatisation du traitement : effectue les corrections sans nécessiter de modification manuelle de chaque élément.

## Technologies
-GO
-GITHUB

## Prérequis
-une version de go plus récente.
-un IDE par exemple visual studio code

## Installation
-télécharger le fichier ZIP du projet.
-Extrayez son contenu dans un dossier.
-ouvrez le dossier du projet dans un IDE compatiple avec Go.
-assurez-vous que Go est installé sur votre ordinateur.
-Exécutez le programme depuis votre IDE ou fournissez un fichier d'entrée, selon les arguments attendus par le programme.

## Lancement
Après avoir exécuté le programme depuis votre IDE, ouvrez le terminal de l'IDE et lancez les commandes suivantes :
cat sample.txt : affiche le contenu du fichier d'origine.
go run . sample.txt result.txt : lit le fichier d'entrée, effectue les corrections et écrit le résultat dans un nouveau fichier.
cat result.txt : affiche le contenu du fichier après correction.
Remarque : remplacez sample.txt par le nom du fichier que vous souhaitez traiter. Veillez à utiliser le nom correct du fichier d'entrée afin que le programme puisse fonctionner correctement.

## Tests
Le programme contient des tests unitaires pour chaque fonction principale afin de vérifier leur bon fonctionnement. Pour les exécuter, il suffit de saisir la commande suivante dans le terminal :
go test ./...
Vous avez également la possibilité de modifier les entrées et les résultats attendus des tests afin d'expérimenter, de mieux comprendre leur fonctionnement et de découvrir les limites du code.

## Achitecture du projet

l'aboraissance du projet est :
KouameThieryGo-reaload
    modifan
        modifan.go
    modifan_test
        modifan_test.go
    modifBin
        modifBin.go
    modifBin_test
        modifBin_test.go
    modifCap
        modifCap.go
    modifcap_test
        modifcap_Nombre_test.go
        modifcap_test.go
    modifHex
        modifHex.go
    modifhex_test
        modifhex_test.go
    modifLow
        modifLow.go
    modifLow_test
        modifLow_test.go
        modifLowNombre_test.go
    modifponctuation
        modifponctuation.go
    modifponctuation_test
        modifponctuation_test.go
    modifUp
        modifUp.go
    modifUp_test
        modifUp_test.go
        modifUpNombre_test.go
    go.mod
    main.go
    README.md
    fichier S
    fichier E



