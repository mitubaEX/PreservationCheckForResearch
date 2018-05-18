import requests
import time
import sys


def convertHexData(row_birthmark):
    hex_row = ""
    for j in row_birthmark:
        for i in j.split(' '):
            try:
                hex_row += hex(int(i)).replace('0x', '')
            except:
                return row_birthmark
        hex_row += ','
    return hex_row


import os
os.makedirs('search_result', exist_ok=True)

# python3 <filename> postFile length birthmark
postFile = sys.argv[1]
length = sys.argv[2]
birthmark = sys.argv[3]
with open(postFile, 'r') as f:
    import csv
    for row in csv.reader(f):
        if len(row[3:]) <= int(length):
            continue
        if birthmark == 'uc':
            postData = ','.join(row[3:])
        else:
            postData = convertHexData(row[3:])
        # print(postData)

        if birthmark == 'uc':
            payload = {'indent': 'on', 'q': 'data:'+postData, 'sort': 'score desc',
                       'wt': 'json', 'rows': '1000', 'fl': '*,score'}
        else:
            payload = {'indent': 'on', 'q': 'encode_data:'+postData, 'sort': 'score desc',
                       'wt': 'json', 'rows': '1000', 'fl': '*,score'}

        start = time.time()
        sumQtime = 0

        r = requests.post(
                'http://localhost:8983/solr/' + birthmark + '/query?fl=data,score,place,barthmark,data&rows=1000&sort=score%20desc&wt=json',
                json={'query': 'encode_data: '+ postData})
        # print(r.json())
        maxScore = float(r.json()['response']['maxScore'])
        # print(maxScore)
        # print(float(r.json()['response']['docs'][-1]['score']))
        starts = 1000
        while True:
            print(maxScore)
            with open('search_result/'+row[0]+birthmark, 'a') as write_file:
                write_file.write(','.join(row) + '\n')
                for result in r.json()['response']['docs']:
                    if float(result['score']) / maxScore < 0.25:
                        break
                    write_file.write('{0},{1},{2},{3}\n'.format(
                        result['data'], result['score'], result['barthmark'], result['data'].replace('quot;', '')))
            if float(float(r.json()['response']['docs'][-1]['score']) / maxScore) < 0.25:
                break
            if birthmark == 'uc':
                payload = {'indent': 'on', 'q': 'data:'+postData,
                           'wt': 'json', 'rows': '1000', 'fl': '*,score', 'start': starts}
            else:
                payload = {'indent': 'on', 'q': 'encode_data:'+postData,
                           'wt': 'json', 'rows': '1000', 'fl': '*,score', 'start': starts}
            r = requests.post(
                    'http://localhost:8983/solr/' + birthmark + '/query?fl=data,score,place,barthmark,data&rows=1000&sort=score%20desc&wt=json&start=' + str(starts),
                    json={'query': 'encode_data: '+ postData})
            # r = requests.get(
            #     'http://localhost:8983/solr/' + birthmark + '/select', params=payload)
            starts += 1000
            # qtime
            sumQtime += r.json()['responseHeader']['QTime']
            # print('{0}, {1}'.format(maxScore, float(
            #     r.json()['response']['docs'][-1]['score'])))
            # print(sumQtime)
            elapsed_time = time.time() - start
            # print("elapsed_time:{0}".format(elapsed_time) + "[sec]")

        elapsed_time = time.time() - start
        # print("elapsed_time:{0}".format(elapsed_time) + "[sec]")
