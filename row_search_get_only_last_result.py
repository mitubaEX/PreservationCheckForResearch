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
os.makedirs('compare_result', exist_ok=True)

# python3 <output> postFile length birthmark
postFile = sys.argv[1]
length = sys.argv[2]
birthmark = sys.argv[3]
threshold_list = [0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9]
# threshold = sys.argv[4]
port = sys.argv[4]
file_count = 0
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
            'http://localhost:'+port+'/solr/' + birthmark +
            '/query?fl=output,score,place,barthmark,data&rows=1000000&sort=score%20desc&wt=json',
            json=payload)
        # print(r.json())
        try:
            maxScore = float(r.json()['response']['maxScore'])
        except:
            continue
        starts = 1000
        # qtime
        sumQtime += r.json()['responseHeader']['QTime']
        elapsed_time = time.time() - start

        elapsed_time = time.time() - start
        print("elapsed_time:{0}".format(elapsed_time) + "[sec]")
        import subprocess
        for threshold in threshold_list:
            with open('search_result/a.csv', 'w') as write_file:
                write_file.write(','.join(row) + '\n')

            with open('search_result/b.csv', 'w') as write_file:
                write_file.write(','.join(row) + '\n')
                try:
                    if len(list(r.json()['response']['docs'])) == 0:
                        continue
                except:
                    continue
                narrow_count = 0
                for result in r.json()['response']['docs']:
                    if float(result['score'] / maxScore) >= float(threshold):
                        narrow_count += 1
                        write_file.write('{1},{2},{3},{4}\n'.format(
                            result['output'], float(result['score'])/maxScore, result['barthmark'], result['data'].replace('quot;', '')))

                print('{0},narrow_count: {1}'.format(threshold, narrow_count))


            subprocess.run(["sh", "pochicomp.sh"])

            with open('compare_result/comp_result.csv', 'r') as comp_result:
                count = 0
                correct_count = 0
                for comp in csv.reader(comp_result):
                    if len(comp) == 1:
                        print(comp[0])
                        continue
                    if float(comp[2]) >= 0.75:
                        correct_count += 1
                    count += 1
            print('{0},correct_count: {1}, count: {2}'.format(threshold, correct_count, count))
        # with open('search_result/'+row[0]+birthmark, 'a') as write_file:
        #     write_file.write(','.join(row) + '\n')
        #     try:
        #         if len(list(r.json()['response']['docs'])) == 0:
        #             sys.exit(1)
        #     except:
        #         sys.exit(1)
        #     for result in r.json()['response']['docs']:
        #         write_file.write('{0},{1},{2},{3}\n'.format(
        #             result['output'], float(result['score'])/maxScore, result['barthmark'], result['data'].replace('quot;', '')))


