#!/bin/bash 
curl -s https://zone01normandie.org/assets/superhero/all.json | jq ".[] | select( .id == $HERO_ID ).connections.relatives" | tr -d \"
#!/bin/bash
curl -s https://01.alem.school/assets/superhero/all.json | jq -c '.[] | select(.id == $id)' --argjson id $HERO_ID | jq -c '.connections.relatives' | tr -d '"/"' | tr -d '\"' | tr -d "\'"