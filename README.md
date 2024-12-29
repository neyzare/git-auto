git-auto
git-auto est un outil de ligne de commande écrit en Go qui simplifie les interactions avec Git. Il permet d'exécuter une série de commandes Git (add, commit, et push) avec un seul appel en passant un message de commit directement en argument.

Fonctionnalités
Ajoute tous les fichiers (git add .) au commit.
Crée un commit signé avec un message personnalisé.
Pousse automatiquement les modifications sur la branche principale (main).
Installation
1. Cloner le dépôt
Clonez le projet sur votre machine :

bash
Copier le code
git clone https://github.com/votre-utilisateur/git-auto.git
cd git-auto
2. Initialiser le module Go
Si ce n'est pas déjà fait, initialisez le module Go :

bash
Copier le code
go mod init git-auto
3. Compiler le programme
Générez un exécutable pour votre système d'exploitation :

bash
Copier le code
go build -o git-auto
4. Ajouter au PATH
Déplacez le fichier exécutable dans un répertoire inclus dans votre PATH pour l'utiliser partout :

bash
Copier le code
mv git-auto /usr/local/bin/
Utilisation
Exécutez la commande dans votre terminal avec la syntaxe suivante :

bash
Copier le code
git-auto commit "Votre message de commit"
Exemple :
bash
Copier le code
git-auto commit "Ajout des fichiers de configuration"
Cela exécutera les commandes suivantes dans cet ordre :

git add .
git commit -S -m "Ajout des fichiers de configuration"
git push origin main
Dépendances
Git : Assurez-vous que Git est installé et configuré sur votre machine.
GPG (optionnel) : Si vous utilisez des commits signés, configurez correctement votre clé GPG. Si vous ne souhaitez pas signer vos commits, désactivez cette option avec :
bash
Copier le code
git config --global commit.gpgsign false
Personnalisation
Si vous travaillez sur une branche différente de main, vous pouvez modifier le code dans la fonction launchCommand :

go
Copier le code
commands := fmt.Sprintf("git add . && git commit -S -m '%s' && git push origin <votre-branche>", message)
Remplacez <votre-branche> par le nom de votre branche.

Contribuer
Les contributions sont les bienvenues ! Pour contribuer :

Forkez ce dépôt.
Créez une nouvelle branche pour votre fonctionnalité ou correction de bug :
bash
Copier le code
git checkout -b ma-nouvelle-fonctionnalite
Effectuez vos modifications et committez-les :
bash
Copier le code
git commit -m "Ajout d'une nouvelle fonctionnalité"
Poussez les modifications :
bash
Copier le code
git push origin ma-nouvelle-fonctionnalite
Ouvrez une Pull Request sur GitHub.
Licence
Ce projet est sous licence MIT. Consultez le fichier LICENSE pour plus d'informations.

Auteur
Créé avec ❤️ par Lucas.
Pour toute question ou suggestion, n'hésitez pas à me contacter.

Améliorations futures
Ajouter la prise en charge de branches multiples.
Support pour les options supplémentaires (--amend, --no-verify, etc.).
Gestion des erreurs plus détaillée.
Ce README fournit toutes les informations nécessaires pour utiliser et contribuer à votre projet. 😊






Vous avez dit :
tu peux me donner tout en meme que je puisse le copier coller
ChatGPT a dit :
ChatGPT
Voici le README complet que vous pouvez copier-coller :

