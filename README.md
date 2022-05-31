# Crawler_USP
Projeto desenvolvido para capturar todos os TCC, Mestrados, Doutorados e Livre Docências existentes na biblioteca da Universidade de São Paulo (USP) que versem sobre ```DIREITO```.
 
## Dependências

Para o esse crawler somos dependentes do projeto [selenium](https://github.com/tebeka/selenium#readme), sendo necessário a pré-instalação do [ChromeDriver](https://sites.google.com/a/chromium.org/chromedriver/) e do [selenium-server-standalone](https://selenium-release.storage.googleapis.com/index.html?path=3.5/). Nesse projeto também utilizamos o [htmlquery](https://github.com/antchfx/htmlquery).

## Run
Inicie o selenium server no terminal e depois basta 

```
java -jar selenium-server-standalone.jar

go run main.go

```

## Instalar - MacOS
```
Instalar o Google Chrome
Instalar o JAVA
Instalar o Brew

Baixar o selenium server stand alone\
Permitir o acesso ao arquivo (configurações de segurança)

brew install chrome river
Permitir o chromedriver no painel de segurança
```
