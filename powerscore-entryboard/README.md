# Entryboard

## How to configure endpoint

In src/components/Entryboard.vue, in refreshTeams(), there is a comment that says `API endpoint here` next to where the endoint is set.

## How to get dummy data

Set API endpoint to "http://localhost:3001/teams" and run json-server with the included db.json file:

`npx json-server db.json -p 3001`

You will need to run `npm install json-server` if you've never installed it

# Editboard

Reachable through /editboard, which really means /#/editboard the way this server is configured. Not hooked up to anything, it's pretty much a demo
