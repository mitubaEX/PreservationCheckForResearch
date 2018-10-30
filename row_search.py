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

# python3 <output> postFile length birthmark
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

        if birthmark == 'uc':
            payload = {'query': 'data: ' + postData}
        else:
            payload = {'query': 'encode_data: ' + postData}

        start = time.time()
        sumQtime = 0

        r = requests.post(
            'http://localhost:8983/solr/' + birthmark +
            '/query?fl=output,score,place,barthmark,data&rows=1000000&sort=score%20desc&wt=json',
            json=payload)
        # print(r.json())
        maxScore = float(r.json()['response']['maxScore'])
        starts = 1000
        with open('search_result/'+row[0]+birthmark, 'a') as write_file:
            write_file.write(','.join(row) + '\n')
            try:
                if len(list(r.json()['response']['docs'])) == 0:
                    sys.exit(1)
            except:
                sys.exit(1)
            for result in r.json()['response']['docs']:
                write_file.write('{0},{1},{2},{3}\n'.format(
                    result['output'], float(result['score'])/maxScore, result['barthmark'], result['data'].replace('quot;', '')))

        # qtime
        sumQtime += r.json()['responseHeader']['QTime']
        elapsed_time = time.time() - start

        elapsed_time = time.time() - start
        print("elapsed_time:{0}".format(elapsed_time) + "[sec]")
