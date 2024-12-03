go build --ldflags "-s -w" -o best_wish .
nohup ./best_wish >/dev/null 2>&1 &