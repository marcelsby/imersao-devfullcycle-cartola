from rest_framework.request import Request
from rest_framework.response import Response
from rest_framework.decorators import api_view
from .use_cases import UpdateMyTeamPlayers


@api_view(['PUT'])
def update_my_team_players(request: Request, my_team_uuid):
    UpdateMyTeamPlayers().execute(my_team_uuid, request.data.get('players_uuid'))
    return Response(None, 204)
