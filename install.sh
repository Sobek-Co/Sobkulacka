#!/bin/bash

sudo curl https://git.sleepat.fit/api/packages/SobekCo/debian/repository.key -o /etc/apt/keyrings/forgejo-SobekCo.asc
echo "deb [signed-by=/etc/apt/keyrings/forgejo-SobekCo.asc] https://git.sleepat.fit/api/packages/SobekCo/debian stable main" | sudo tee -a /etc/apt/sources.list.d/forgejo.list
sudo apt update

sudo apt install sobkulacka
