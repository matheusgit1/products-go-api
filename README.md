# Api de Produtos

Este projeto é uma parte de um projeto maior, um ecommerce baseado em microserviços.

![Arquitetura](https://github.com/matheusgit1/products-go-api/blob/main/assets/Captura%20de%20tela%202024-01-31%20163051.png)

Este repositório se refere api de produtos

Este serviço visa tratar dados referentes aos produtos.

Os demais serviços que compoem o projeto podem ser baixados em cada um dos respectivos repositórios, sendo eles:

Frontend:
Api de compras:
Api de pagamentos:

## Preparação

### Requisitos

Estem alguns requisots básicos para a execuçao do projeto em ambiente local. Uma delas é o docker para execução dos containers,
que pode ser baixado gratuitamente em https://www.docker.com/

Uma vez intalado e configurado o docker, execute os seguintes comandos dentro da pasta raíz do serviço

```bash
docker-compose up --build
```

isso irá subir um banco de daos mysql para execução da aplicação.
Para validar se o banco está de pé, execute os camandos em sequencia

```bash
docker compose exec mysql bash
```

```bash
mysql exec mysql bash
```

```bash
mysql -uroot -proot
```

```bash
show databases;
```

procure pelo banco "commerce" ou pelo banco que você definiu nas variaveis de ambiente

## Stack utilizada

Go lang, Docker, sql, MySql

## Testes unitário

Use o arquivo api.http para execução das rotas no vscode.

## Sobre o Autor

Eu sou uma pessoa desenvolvedora full-stack, técnico em administração, engenheiro de automação em formação, e cientista de dados em formação. Sempre busco por excelência e entregar o máximo com a maior qualidade, sem claro, deixar de lado boas práticas.

Atualmente sou desenvolvedor full stack júnior da área de desenvolvimento de softwares, mirando senioridades cada vez mais altas.

Tenho habilidades com as stacks mais modernas, como :nodeJS, typeScript, css, html, nestJs, NextJs, aws-cloud, bancos de dados não relacionais como mongodb e redis, bancos de dados relacionais como MySql, postgres, docker, testes unitários e de integração. Atuando também em diferentes setores, como educação e telecomunicação.

## Quer entrar em contato com o desenvolvedor?

🪜 Instagram (sempre respondo): @ap_matheus

📱 Telefone e whatsapp: 55 27 997822665

📫 Email: pereira.matheusalves@gmail.com

🔗 Linkedin: https://www.linkedin.com/in/matheus-alves-pereira-4b3781222/
