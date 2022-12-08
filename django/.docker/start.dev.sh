#!/bin/bash

# Instala as dependências
pipenv install

# Executa as migrations
pipenv run python manage.py migrate

# Carrega os dados iniciais
pipenv run python manage.py loaddata initial_data

tail -f /dev/null