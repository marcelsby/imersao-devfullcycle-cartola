from django.urls import path
from .api import update_my_team_players

urlpatterns = [
    path('my-teams/<uuid:my_team_uuid>/players', update_my_team_players)
]
