### Starting server


        make run


### Starting client

        docker run  --network host -it --entrypoint  /bin/bash  solsson/websocat

From container shell:

        websocat ws://0.0.0.0:$SERVER_PORT/ws