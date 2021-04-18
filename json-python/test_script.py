# In this script I will test all the writing and reading from python to a json file
import json

###ATRIBUTES###
global data

with open('data.json') as file:
    data = json.load(file)

data['customers'] = []
data['customers'].append({
    'first_name': 'Sigrid',
    'last_name': 'Mannock',
    'age': 27,
    'amount': 7.17})
data['customers'].append({
    'first_name': 'Joe',
    'last_name': 'Hinners',
    'age': 31,
    'amount': [1.90, 5.50]})
data['customers'].append({
    'first_name': 'Theodoric',
    'last_name': 'Rivers',
    'age': 36,
    'amount': 1.11})

print(data)

with open('data.json', 'w') as file:
    json.dump(data, file, indent=4)

