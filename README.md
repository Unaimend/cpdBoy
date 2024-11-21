You can self-host the server on your own machine. For that you have to change the ip addresses in the r-scrip and bash file.

## Installation
A small server + apis for different programming languages to convert cpds into more useful information

### Prerequisites:
- Docker


### Starting the server
This will start the server in the background.

```
git clone https://github.com/Unaimend/cpdBoy.git
cd cpdBoy
docker build -t cpd-boy .
docker compose up --detach
```

To test if the server is running and in a reachable state. Paste
```
curl -X POST http://49.12.211.202:3000/message -d '{"text": "cpd00058,cpd00059"}' -H "Content-Type: application/json"
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
sudo wget -O ~/usr/local/bin/cpd https://github.com/Unaimend/cpdBoy/blob/main/bash_script/cpd.sh && chmod +x  ~/usr/local/bin/cpd
``` 

Running the cpd script like this
```
cpd cpd00058,cpd00059
```
should again give you the same return as the curl command above

## Development

### Run tests
To execute the tests run 

`go tests ./...` 
