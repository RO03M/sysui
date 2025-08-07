serve:
	$(MAKE) backend & $(MAKE) frontend
	wait
backend:
	cd server && go run *.go
frontend:
	cd client && yarn dev