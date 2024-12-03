go build --ldflags "-s -w" -o best_wish .
ps -ef | grep best_wish | grep -v grep | awk '{print $2}' | xargs kill
nohup ./best_wish >/dev/null 2>&1 &