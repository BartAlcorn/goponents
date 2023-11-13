#!/usr/local/bin/zsh
echo "starting air" &
air -c ./.air.toml &
npx tailwindcss -i ./web/index.css -o ./static/index.css --watch &
# npx browser-sync start \
#   --files 'web/**/*.gohtml, static/**/*.css' \
#   --port 8080 \
#   --proxy 'localhost:3000' \
#   --middleware 'function(req, res, next) { \
#     res.setHeader("Cache-Control", "no-cache, no-store, must-revalidate"); \
#     return next(); \
#   }'
