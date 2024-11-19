## Installation

### Prerequisites:
- Docker


### Starting the server
This will start the server in the background.

```
git clone <this repo>
cd cpdBoy
docker compose up --detach
```

To test if the server is running an reachable paste
```
curl -X POST http://127.0.0.1:3000/message -d '{"text": "cpd00058,cpd00059"}' -H "Content-Type: application/json"
```
into your terminal. It should return
```
cpd00058,Cu2+
cpd00059,L-Ascorbate
```

### R-Package


### Bash script
To get the bash script run and make sure that `/usr/local/bin is in your $PATH`
```
sudo wget -O ~/usr/local/bin/cpd.sh <this repo>/bash_script/cpd.sh
```

Running the cpd script like this
```
cpd cpd00058,cpd00059
```
should again give you the same return as the curl command above
