# Django - Painel Administrativo

## 🏃 Executar localmente

1. Na pasta do projeto execute:

```bash
$ docker compose up -d
```

2. Entre no contêiner criado:

```bash
$ docker compose exec app bash
```

3. Ative o virtualenv:

```bash
$ pipenv shell
```

4. Instale as dependências:

```bash
$ pipenv install
```

5. Execute as migrações e carrege os dados iniciais:

```bash
$ python manage.py makemigrations app && python manage.py migrate && python manage.py loaddata initial_data
```

6. Execute a aplicação:

```bash
$ python manage.py runserver 0.0.0.0:8000
```

2. Clique [aqui](http:localhost:8000/admin) para acessar a aplicação no navegador.