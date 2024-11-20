_help:
  @just --list

# Start the FTL backend with hot-reloading
dev:
  ftl dev --allow-origins '*'

# Generate the mobile API client
gen-mobile:
  ftl schema generate mobile/templates/ online-boutique/mobile/lib/api --watch=5s

# Generate the web API client
gen-web:
  ftl schema generate web/templates/ online-boutique/web/src/api --watch=5s

# Start the example web frontend for online-boutique
web:
  cd web && npm install && npm run dev

# Run "go mod tidy" on all packages including tests
tidy:
  git ls-files | grep go.mod | grep -v '{{{{' | xargs -n1 dirname | xargs -I {} sh -c 'cd {} && echo {} && go mod tidy'
