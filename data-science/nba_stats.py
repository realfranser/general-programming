import requests

URL = 'https://stats.nba.com/stats/leagueLeaders?LeagueID=00&PerMode=PerGame&Scope=Rookies&Season=2021-22&SeasonType=Regular+Season&StatCategory=EFF'

headers = {'user_agent':'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36'}

response = requests.get(URL, headers=headers)

#print(response.body)
print(response.json())


