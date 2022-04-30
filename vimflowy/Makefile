launch:	has_volume pull_image
	docker run -d \
             -e VIMFLOWY_PASSWORD=supersecretpassword \
             --name vimflowy \
             --mount source=vimflowy-db,target=/app/db \
             -p 3000:3000 \
             --restart unless-stopped \
             vimflowy/vimflowy

stop:
	docker stop vimflowy && docker rm vimflowy

pull_image:
	docker pull vimflowy/vimflowy

has_volume:
	docker volume create vimflowy-db


.PHONY:	launch stop pull_image has_volume
