# Sobkulačka

Davidova kalkulačka

## Jak přispět/interagovat s tímto repositářem

Vysoce doporučuji nejdříve přečíst celý návod, než začnete cokoliv dělat.

### Založení účtu na Sleep@FIT

Pro interakci s tímto repositářem (nebo jinými repositáři hostovanými na tomto git serveru) je zapotřebí účet u služby Sleep@FIT. Registrace je otevřená veřejnosti a lze se na stránku příhlášení/registrace dostat například zmáčknutím tlačítka [Sign in](https://git.sleepat.fit/user/login?redirect_to=%2fexplore%2frepos) v pravém horním rohu uživatelského rozhraní [tohoto git serveru](https://git.sleepat.fit) a poté zmáčknutím na tlačítko "Sign in with Sleep@FIT", které vás přesměruje na přihlášení/registraci do služby.

### Přidání SSH klíče

Doporučenou metodou interagování s repositáři na tomto git serveru je SSH. K přístupu přes SSH je nutno si zaregistrovat svůj SSH klíč u git serveru, aby vašemu zařízení důvěřoval. K tomu je prvně zapotřebí mít nějaký SSH klíč. Velké množství moderních operačních systémů generuje uživateli při registraci klíče, ale nemusí tomu tak vždy být. Pro ověření zda máte již vygenerované klíče bych doporučil [tento článek od GitHubu](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/checking-for-existing-ssh-keys) a v případě že klíče zatím nemáte, tak můžete postupovat podle [tohoto návodu také od GitHubu](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent) (zapotřebí by mělo být jen vygenerovat klíč, další kroky uvedené v návodu, jako třeba přidání klíče do `ssh-agent` by nemělo být v tomto případě zapotřebí).

Jakmile máte vlastní SSH klíč, je ho potřeba nahrát do git serveru. Důležité zmínit, že nahrávat budete pouze **veřejný** klíč, soukromý by jste měli za každou cenu udržet v bezpečí a s nikým nesdílet. Nahrání klíče můžete provést na stránce [nastavení v záložce SSH / GPG klíču](https://git.sleepat.fit/user/settings/keys). Po zmáčknutí na tlačítko pro přídání klíče se vám objeví dvě textová pole, do vrchního zadejte nějaký název který bude sloužit k identifikaci zařízení (můžete použít úplně cokoliv), do druhého pole vložte textový obsah souboru veřejného klíče, který jste vygenerovli (v textovém poli je placeholder, který říká čím by klíč měl začínat, aby jste poznali, že se jedná o ten pravý).

Po přidání klíče je ještě možné klíč ověřit, není to nutné k používání, ale doporučil bych to. Instrukce se zobrazí po zmáčknutí na tlačítko ověřit vedle nově přidaného klíče.

### Naklonování repositáře

Repositář lze naklonovat použitím SSH adresy repositáře (`git@git.sleepat.fit:SobekCo/Sobkulacka.git`) a kterýmkoli nástrojem na klonování, ať už z příkazové řádky (pomocí příkazu pod tímto odstavcem), nebo lze repositář otevřít přímo v editoru za použití tlačíkta "More operations" nacházejícího se vpravo od adresy HTTPS/SSH tohoto repositáře.

```bash
git clone git@git.sleepat.fit:SobekCo/Sobkulacka.git
```

### Příprava gitu

Před zaslaním změn přes git, je zapotřebí nastavit dva hlavní parametry potřebné k identifikaci vás jakožto autora změn. To lze učinit jednoduše těmito příkazy:

```bash
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"
```

Pokud již máte nastavené jméno a email globálně ale chtěli by jste identifikaci změnit čistě pro tento konkrétní repositář, je možné příkazy spustit opět ale bez `--global` v pracovní složce repositáře.

Jelikož se jedná o školní projekt a vyučující specificky požádali o možnost jednoduše identifikovat autory, poprosil bych o použití celého jména místo přezdívky.

### Vytváření větví

Pro tento repositář bylo vytvořeno pravidlo chránící hlavní větev `master` před přímým nahráním změn. Místo toho, když chcete přidat své změny do hlavní větve, je zapotřebí nekolika kroků. Prvním z nich je vytvořit si větev pro své změny. Doporučuji tento krok provést **před tím, než začnete upravovat kód**.

Pokud využíváte příkazové řádky (vysoce doporučuji se s ní naučit), můžete vytvářet větve a přepínat mezi nimi například takto:

```bash
git checkout master # Prepnuti na vetev master
git checkout -b feature/some-feature # Vytvoreni nove vetve s nazvem "feature/some-feature"
git branch # Vypise vsechny vetve nachazejici se **na vasem zareizeni**
git fetch # Stazeni prehledu dostupnych vetvi od serveru
git switch other-feature # Dalsi zpusob prepnuti na jinou vetev
```

Pokud využíváte integrované vývojové prostředí, tak bych doporučil najít si instrukce pro konkrétní prostředí na internetu. Například pro Visual Studio Code jsou instrukce dostupné [zde](https://code.visualstudio.com/docs/sourcecontrol/branches-worktrees)

### Vytvoření žádosti o příjetí

Jakmile jste vytvořili svou novou větev, přidali do ní změny a nahráli ji na server (všechno z prezentací tady opakovat nebudu), tak je zapotřebí vytvořit takzvanou žádost o přijetí změn. Jedná se o proces ověření změn a po schválení dalšími vývojáři, sloučení do hlavní větve. Proces je poměrně jednoduchý a instrukce (i s obrázky) můžete najít [zde](https://forgejo.org/docs/next/user/pull-requests-and-git-flow/). Pro tento repositář byla nastavena povinnost alespoň jednoho ověření před tím, než může být větev sloučena. Ověření musí provést **jiný člen projektu**, nemůžete si ověřit vlastní změny.

Prostředí
---------

Ubuntu 64bit

Autoři
------

Sobek & Co.

- xpalacj00 Jakub Palacký
- xsobekd00 David Sobek
- xstovim00 Martin Šťovíček
- xvitekv00 Vojtěch Vítek

Licence
-------

Veškeré části tohoto projektu včetně jeho zdrojového kódu (pokud není v hlavičce konkrétního souboru uvedeno jinak) jsou distribuovány pod licencí [GNU GPL 3.0](LICENSE)
